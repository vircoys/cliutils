package diskcache

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
)

const (
	dataHeaderLen = 4
	eofHint       = 0xdeadbeef
)

var (
	ErrNoData                    = errors.New("no data")
	ErrUnexpectedReadSize        = errors.New("unexpected read size")
	ErrTooLargeData              = errors.New("too large data")
	ErrEOF                       = errors.New("EOF")
	ErrInvalidDataFileName       = errors.New("invalid datafile name")
	ErrInvalidDataFileNameSuffix = errors.New("invalid datafile name suffix")
	ErrBadHeader                 = errors.New("bad header")

	l = logger.DefaultSLogger("diskcache")

	defaultOpt = &Option{
		NoSync:      false,
		BatchSize:   64 * 1024 * 1024,
		MaxDataSize: 32 * 1024 * 1024,
		DirPerms:    0750,
		FilePerms:   0640,
	}
)

type DiskCache struct {
	path string

	dataFiles []string

	curWriteFile string
	curReadfile  string

	wfd *os.File // write fd
	rfd *os.File // read fd

	lastPut time.Time

	rotateCount  int
	droppedBatch int

	size         int64
	curBatchSize int64

	lock *sync.Mutex

	opt *Option
}

type Option struct {
	// Batch file size, default 64MB
	BatchSize int64

	// Max single data size, default 32MB
	MaxDataSize int64

	// Total disk capacity, default unlimited
	Capacity int64

	// NoSync if enabled, may cause data missing, default false
	NoSync bool

	// File permisions, default 0750/0640
	DirPerms, FilePerms os.FileMode
}

// Open init and create a new disk cache
func Open(path string, opt *Option) (*DiskCache, error) {

	l = logger.SLogger("diskcache")

	c := &DiskCache{
		path:         path,
		opt:          opt,
		curWriteFile: filepath.Join(path, "data"),
		lock:         &sync.Mutex{},
	}

	if c.opt == nil {
		c.opt = defaultOpt
	}

	if c.opt.DirPerms == 0 {
		opt.DirPerms = 0755
	}

	if c.opt.FilePerms == 0 {
		opt.FilePerms = 0640
	}

	if c.opt.BatchSize == 0 {
		c.opt.BatchSize = defaultOpt.BatchSize
	}

	if c.opt.MaxDataSize == 0 {
		c.opt.MaxDataSize = defaultOpt.MaxDataSize
	}

	if c.opt.MaxDataSize > c.opt.BatchSize {
		l.Warnf("reset MaxDataSize from %d to %d", c.opt.MaxDataSize, c.opt.BatchSize/2)
		c.opt.MaxDataSize = c.opt.BatchSize / 2
	}

	if err := os.MkdirAll(path, c.opt.DirPerms); err != nil {
		return nil, err
	}

	// write append fd, always write to the same-name file
	if err := c.openWriteFile(); err != nil {
		return nil, err
	}

	// list files under @path
	arr := []string{}
	if err := filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		c.size += fi.Size()

		arr = append(arr, path)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Strings(arr)
	if len(arr) > 1 && arr[0] == c.curWriteFile {
		c.dataFiles = arr[1:] // ignore first writing file, we do not read file `data` if data.000001/0000002/... exists
	}

	l.Infof("init datafiles: %+#v\n", c.dataFiles)

	return c, nil
}

// Reset clear current DiskCache.
func (c *DiskCache) Reset() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.rfd != nil {
		if err := c.rfd.Close(); err != nil {
			return err
		}
	}

	if c.wfd != nil {
		if err := c.wfd.Close(); err != nil {
			return err
		}
	}

	c.dataFiles = nil
	c.curWriteFile = ""
	c.curReadfile = ""

	c.rotateCount = 0
	c.droppedBatch = 0

	c.size = 0
	c.curBatchSize = 0

	if err := os.Remove(c.path); err != nil {
		return err
	}

	return nil
}

// Put write @data to disk cache, if reached batch size, a new batch is rotated.
func (c *DiskCache) Put(data []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.lastPut = time.Now()

	if c.opt.Capacity > 0 && c.size+int64(len(data)) > c.opt.Capacity {
		if err := c.dropBatch(); err != nil {
			return err
		}
	}

	if int64(len(data)) > c.opt.MaxDataSize {
		return ErrTooLargeData
	}

	hdr := make([]byte, dataHeaderLen)

	binary.LittleEndian.PutUint32(hdr, uint32(len(data)))
	if _, err := c.wfd.Write(hdr); err != nil {
		return err
	}

	if _, err := c.wfd.Write(data); err != nil {
		return err
	}

	if !c.opt.NoSync {
		if err := c.wfd.Sync(); err != nil {
			return err
		}
	}

	c.curBatchSize += int64(len(data) + dataHeaderLen)
	c.size += int64(len(data) + dataHeaderLen)

	// rotate new file
	if c.curBatchSize >= c.opt.BatchSize {
		if err := c.rotate(); err != nil {
			return err
		}
	}

	return nil
}

// Fn is the handler to eat cache from disk
type Fn func([]byte) error

// Get fetch new data from disk cache, then passing to @fn
// if any error occurred during call @fn, the reading data is
// ignored, and will not read again
func (c *DiskCache) Get(fn Fn) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	// wakeup sleeping write file, rotate it for successing reading!
	if time.Since(c.lastPut) > time.Second*3 && c.curBatchSize > 0 {
		l.Debugf("####################### wakeup %s(%d bytes), global size: %d\n", c.curWriteFile, c.curBatchSize, c.size)
		if err := c.rotate(); err != nil {
			return err
		}

		time.Sleep(time.Second * 1)
	}

	if c.rfd == nil {
		if err := c.switchNextFile(); err != nil {
			return err
		}
	}

retry:
	if c.rfd == nil {
		return ErrEOF
	}

	hdr := make([]byte, dataHeaderLen)
	n, err := c.rfd.Read(hdr)
	if err != nil {
		pos, _err := c.rfd.Seek(0, 1)
		if _err != nil {
			return fmt.Errorf("rfd.Seek: %w", _err)
		}
		return fmt.Errorf("rfd.Read(%s/pos: %d): %w", c.curReadfile, pos, err)
	}

	if n != dataHeaderLen {
		return ErrBadHeader
	}

	nbytes := binary.LittleEndian.Uint32(hdr[0:])

	if nbytes == eofHint { // EOF
		if err := c.removeCurrentReadingFile(); err != nil {
			return fmt.Errorf("removeCurrentReadingFile: %w", err)
		}

		// reopen next file to read
		if err := c.switchNextFile(); err != nil {
			return err
		}

		goto retry // read next new file
	}

	databuf := make([]byte, nbytes)

	n, err = c.rfd.Read(databuf)
	if err != nil {
		return err
	}

	if n != int(nbytes) {
		return ErrUnexpectedReadSize
	}

	// NOTE: if @fn failed, c.rfd never seek back, data dropped
	return fn(databuf)
}
