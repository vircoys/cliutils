package logger

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"go.uber.org/zap"
)

var (
	flagLogFile = flag.String("f", "/dev/null", "log file path")
)

func _init() {
	flag.Parse()
}

// test write on symlink
func TestLogger8(t *testing.T) {
	MaxSize = 1
	MaxBackups = 5
	logmsg := "1234567890qwertyuioplkjhgfdsazxcvbnm"
	_ = logmsg

	logf := "log.origin"
	logln := "/tmp/log.ln"

	_ = os.Remove(logln) // ignore

	if err := os.Symlink(logf, logln); err != nil {
		t.Fatal(err)
	}

	SetGlobalRootLogger(logln, DEBUG, OPT_DEFAULT)

	l := SLogger("TestLogger8")
	for {
		l.Debug(logmsg)
	}
}

func TestLogger7(t *testing.T) {
	l := DefaultSLogger("test-7")
	l.Info("info")
	l.Warn("warn")
	l.Error("err")
}

func TestLogger6(t *testing.T) {

	f := func(i int) {
		l := DefaultSLogger(fmt.Sprintf("logger-%d", i))
		l.Debugf("[%d] debug msg", i)
		l.Infof("[%d] info msg", i)
		l.Warnf("[%d] warn msg", i)
	}

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			f(i)
		}(i)
	}

	wg.Wait()

	fmt.Printf("stdout pirint...")
}

func TestLogger5(t *testing.T) {
	_init()

	SetStdoutRootLogger(DEBUG, OPT_DEFAULT)
	l1 := SLogger("test")
	l1.Info("haha")

	SetGlobalRootLogger("/tmp/x.log", DEBUG, OPT_DEFAULT)
	l2 := SLogger("test")
	l2.Info("haha")
}

func TestLogger4(t *testing.T) {
	_init()
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_DEFAULT)

	l := SLogger("test")
	l.Debug("this is debug msg")
	l.Info("this is info msg")
	l.Error("this is error msg")
}

func TestLogger3(t *testing.T) {
	_init()
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_ENC_CONSOLE|OPT_SHORT_CALLER|OPT_COLOR|OPT_RESERVED_LOGGER)

	l := SLogger("test")
	l.Debug("this is debug msg")
	l.Info("this is info msg")
	l.Error("this is error msg")
}

func TestLogger2(t *testing.T) {
	_init()
	SetGlobalRootLogger("/tmp/x.log", DEBUG, OPT_DEFAULT)

	l := SLogger("test")
	l.Debug("this is debug msg")
	l.Info("this is info msg")
	l.Error("this is error msg")
}

func TestRorate2(t *testing.T) {
	_init()
	MaxSize = 1
	MaxBackups = 5
	f := `/dev/stdout`
	l, err := newRotateRootLogger(f, DEBUG, OPT_SHORT_CALLER|OPT_ENC_CONSOLE|OPT_COLOR)
	if err != nil {
		t.Fatal(err)
	}

	l1 := getSugarLogger(l, "TestRorate2")

	fn := func() {
		l1.Debug("this is debug msg")
		l1.Error("this is error msg")
		l1.Warn("this is warn msg")
	}

	fn()
}

func TestRorate(t *testing.T) {
	_init()
	MaxSize = 1
	MaxBackups = 5

	f := ``
	l, err := newRotateRootLogger(f, DEBUG, OPT_SHORT_CALLER|OPT_ENC_CONSOLE)
	if err != nil {
		t.Fatal(err)
	}

	l1 := getSugarLogger(l, "TestRorate.1")
	l2 := getSugarLogger(l, "TestRorate.2")
	l3 := getSugarLogger(l, "TestRorate.3")
	l4 := getSugarLogger(l, "TestRorate.4")

	exit := make(chan interface{})

	go func() {
		for {
			l1.Debug("this is debug msg")
			l1.Info("this is info msg")
			l2.Info("this is info msg")
			l2.Debug("this is info msg")
			l3.Info("this is info msg")
			l3.Debug("this is info msg")
			l4.Info("this is info msg")
			l4.Debug("this is info msg")

			select {
			case <-exit:
				return
			default:
			}
		}
	}()

	time.Sleep(time.Second * 30)
	close(exit)
}

func TestLogger1(t *testing.T) {
	_init()
	base := 4
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_ENC_CONSOLE|OPT_SHORT_CALLER|OPT_COLOR)

	l1 := SLogger("test")
	l2 := SLogger("test")

	wg := sync.WaitGroup{}

	for j := 0; j < base; j++ {
		wg.Add(2)
		go func() {
			i := 0
			defer wg.Done()
			for {
				l1.Debugf("L1: %v", l1)
				i++

				if i%(base*8) == 0 {
					fmt.Printf("[%d]L1: %d\n", j, i)
				}

				if i > base*32 {
					return
				}
			}
		}()

		go func() {
			i := 0
			defer wg.Done()
			for {
				l2.Debugf("L2: %v", l2)
				i++

				if i%(base*8) == 0 {
					fmt.Printf("[%d]L2: %d\n", j, i) //nolint:govet
				}

				if i > base*32 {
					return
				}
			}
		}()
	}

	wg.Wait()
}

func TestColor(t *testing.T) {
	_init()
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_ENC_CONSOLE|OPT_SHORT_CALLER|OPT_COLOR)

	l := SLogger("test")
	l.Debug("this is debug message")
	l.Info("this is info message")
	l.Warn("this is warn message")
	l.Error("this is error message")
}

func TestStdoutGlobalLogger(t *testing.T) {
	_init()
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_ENC_CONSOLE|OPT_SHORT_CALLER)

	l := SLogger("test")
	l.Debug("this is debug message")
	l.Info("this is info message")
}

func TestWinGlobalLogger(t *testing.T) {
	_init()
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_STDOUT|OPT_ENC_CONSOLE|OPT_SHORT_CALLER)

	l := SLogger("test")

	l.Debug("this is debug message")
	l.Info("this is info message")
}

func TestGlobalLoggerNotSet(t *testing.T) {
	_init()
	sl := SLogger("sugar-module")
	sl.Debugf("sugar debug msg")
}

func TestGlobalLogger(t *testing.T) {
	_init()
	SetGlobalRootLogger(*flagLogFile, DEBUG, OPT_ENC_CONSOLE|OPT_SHORT_CALLER)

	sl := SLogger("sugar-module")
	sl.Debugf("sugar debug msg")

	l := getLogger(defaultRootLogger, "x-module")
	fmt.Printf("%+#v", l)

	f := zap.Duration("backoff", time.Second)
	fmt.Println(f)

	l.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "baidu.com"),
		zap.Int("attempt", 3),
		zap.Int("attempt", 4))
}

func TestLogger(t *testing.T) {
	_init()
	rl, err := newRootLogger(*flagLogFile, INFO, OPT_ENC_CONSOLE|OPT_SHORT_CALLER)
	if err != nil {
		panic(err)
	}

	sl := getSugarLogger(rl, "testing")
	sl.Debug("test message")
	sl.Info("this is info msg: ", "info msg")

	l := getLogger(rl, "debug")
	l.Debug("this is debug: ", zap.Int("int", 42))
}
