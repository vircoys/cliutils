// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package diskcache

import (
	"errors"
	"sync"
	"sync/atomic"
	T "testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkNosyncPutGet(b *T.B) {
	p := b.TempDir()
	c, err := Open(WithPath(p), WithNoSync(true), WithBatchSize(1024*1024*4), WithCapacity(4*1024*1024*1024))
	require.NoError(b, err)

	_1mb := make([]byte, 1024*1024)
	_1kb := make([]byte, 1024)
	_512kb := make([]byte, 1024*512)

	b.Run(`put-1mb`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Put(_1mb)
		}
	})

	b.Run(`put-1kb`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Put(_1kb)
		}
	})

	b.Run(`put-512kb`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Put(_512kb)
		}
	})

	b.Run(`get`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Get(func(_ []byte) error { return nil })
		}
	})

	m := c.Metrics()
	b.Logf(m.LineProto())

	b.Cleanup(func() {
		assert.NoError(b, c.Close())
	})
}

func BenchmarkPutGet(b *T.B) {
	p := b.TempDir()
	c, err := Open(WithPath(p), WithBatchSize(1024*1024*4), WithCapacity(4*1024*1024*1024))
	require.NoError(b, err)

	_1mb := make([]byte, 1024*1024)
	_1kb := make([]byte, 1024)
	_512kb := make([]byte, 1024*512)

	b.Run(`put-1mb`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Put(_1mb)
		}
	})

	b.Run(`put-1kb`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Put(_1kb)
		}
	})

	b.Run(`put-512kb`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Put(_512kb)
		}
	})

	b.Run(`get`, func(b *T.B) {
		for i := 0; i < b.N; i++ {
			c.Get(func(_ []byte) error { return nil })
		}
	})

	m := c.Metrics()
	b.Logf(m.LineProto())

	b.Cleanup(func() {
		assert.NoError(b, c.Close())
	})
}

func TestConcurrentPutGet(t *T.T) {
	var (
		p      = t.TempDir()
		mb     = int64(1024 * 1024)
		sample = make([]byte, 5*7351)
		eof    = 0
	)

	c, err := Open(WithPath(p), WithBatchSize(4*mb), WithCapacity(128*mb))
	assert.NoError(t, err)

	defer c.Close()

	wg := sync.WaitGroup{}
	concurrency := 4

	fnPut := func(idx int) {
		defer wg.Done()
		nput := 0
		exceed100ms := 0

		for {
			start := time.Now()
			assert.NoError(t, c.Put(sample))

			cost := time.Since(start)
			if cost > 100*time.Millisecond {
				exceed100ms++
			}

			nput++

			if nput > 1000 {
				t.Logf("[%d] Put exit", idx)
				return
			}
		}
	}

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go fnPut(i)
	}

	fnGet := func(idx int) {
		defer wg.Done()
		nget := 0
		readBytes := 0
		exceed100ms := 0

		for {
			nget++
			start := time.Now()
			if err := c.Get(func(x []byte) error {
				assert.Equal(t, sample, x)
				readBytes += len(x)

				cost := time.Since(start)
				if cost > 100*time.Millisecond {
					exceed100ms++
				}

				return nil
			}); err != nil {
				if errors.Is(err, ErrEOF) {
					time.Sleep(time.Second)
					eof++
					if eof >= 10 {
						break
					}
				} else {
					t.Logf("[%d]: %s", idx, err)
					time.Sleep(time.Second)
				}
			} else {
				eof = 0 // reset eof if Get ok
			}
		}
	}

	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go fnGet(i)
	}

	wg.Wait()

	t.Logf("metric: %s", c.Metrics().LineProto())
	t.Cleanup(func() {
		assert.NoError(t, c.Close())
	})
}

func TestPutOnCapacityReached(t *T.T) {
	t.Run(`reach-capacity-single-put`, func(t *T.T) {
		var (
			mb       = int64(1024 * 1024)
			p        = t.TempDir()
			capacity = 32 * mb
			large    = make([]byte, mb)
			small    = make([]byte, 1024*3)
			maxPut   = 4 * capacity
		)

		t.Logf("path: %s", p)

		c, err := Open(WithPath(p), WithCapacity(capacity), WithBatchSize(4*mb))
		assert.NoError(t, err)

		defer c.Close()

		n := 0
		for {
			switch n % 2 {
			case 0:
				c.Put(small)
			case 1:
				c.Put(large)
			}
			n++

			if c.putBytes > maxPut {
				break
			}
		}

		m := c.Metrics()
		t.Logf("metric: %s", m.LineProto())
		fs := m.Fields()
		assert.True(t, fs.Get([]byte(`dropped_batch`)).GetI() > 0)
	})

	t.Run(`reach-capacity-concurrent-put`, func(t *T.T) {
		var (
			mb       = int64(1024 * 1024)
			p        = t.TempDir()
			capacity = 128 * mb
			large    = make([]byte, mb)
			small    = make([]byte, 1024*3)
			maxPut   = 4 * capacity
			wg       sync.WaitGroup
		)

		t.Logf("path: %s", p)

		c, err := Open(WithPath(p), WithCapacity(capacity), WithBatchSize(4*mb))
		assert.NoError(t, err)

		defer c.Close()

		total := int64(0)

		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				n := 0
				for {
					switch n % 2 {
					case 0:
						c.Put(small)
						atomic.AddInt64(&total, 1024*3)
					case 1:
						c.Put(large)
						atomic.AddInt64(&total, mb)
					}
					n++

					if atomic.LoadInt64(&total) > maxPut {
						return
					}
				}
			}()
		}

		wg.Wait()

		m := c.Metrics()
		t.Logf("metric: %s", m.LineProto())
		fs := m.Fields()
		assert.True(t, fs.Get([]byte(`dropped_batch`)).GetI() > 0)
	})
}
