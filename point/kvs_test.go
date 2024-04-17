// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package point

import (
	"math"
	"sort"
	T "testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *T.T) {
	t.Run("trim-field", func(t *T.T) {
		var kvs KVs
		kvs = kvs.Add("f0", 1.23, false, false)
		kvs = kvs.AddTag("t1", "v1")
		kvs = kvs.Add("f1", -123, false, false)
		kvs = kvs.Add("f2", uint64(123), false, false)
		kvs = kvs.Add("f3", "hello", false, false)
		kvs = kvs.Add("f4", []byte("world"), false, false)
		kvs = kvs.Add("f5", false, false, false)

		kvs = kvs.TrimFields(1)

		assert.Lenf(t, kvs, 2, "go kvs: %s", kvs.Pretty())
		assert.NotNil(t, kvs.Get("f0"))
		assert.NotNil(t, kvs.Get("t1"))
	})

	t.Run("point-pool-kv-reuse", func(t *T.T) {
		pp := &fullPointPool{}
		SetPointPool(pp)
		defer ClearPointPool()

		i := 0
		var kvs KVs

		kvs = kvs.Add("f0", 1.23, false, false)
		kvs = kvs.AddTag("t1", "v1")

		for {
			kvs = kvs.TrimFields(0)

			assert.Lenf(t, kvs, 1, "go kvs: %s", kvs.Pretty())

			kvs = kvs.Add("f-1", 123, false, false)

			// kvs = kvs.Add("f-5", 123, false, false)

			// kvs = kvs.Add("f-8", 123, false, false)
			if pp.kvReused.Load() > 0 {
				t.Logf("[%d] %s", i, pp)
				break
			}
			i++
		}
	})

	t.Run("trim-field-under-point-pool", func(t *T.T) {
		pp := &fullPointPool{}
		SetPointPool(pp)
		defer ClearPointPool()

		for loop := 0; loop < 2; loop++ {
			var kvs KVs
			kvs = kvs.Add("f0", 1.23, false, false)
			kvs = kvs.AddTag("t1", "v1")
			kvs = kvs.Add("f1", -123, false, false)
			kvs = kvs.Add("f2", uint64(123), false, false)
			kvs = kvs.Add("f3", "hello", false, false)
			kvs = kvs.Add("f4", []byte("world"), false, false)
			kvs = kvs.Add("f5", false, false, false)

			kvs = kvs.TrimFields(2)

			assert.Lenf(t, kvs, 3 /* 2 fields  + 1 tag */, "go kvs: %s", kvs.Pretty())
			assert.NotNil(t, kvs.Get("f0"))
			assert.NotNil(t, kvs.Get("f1"))
			assert.NotNil(t, kvs.Get("t1"))

			kvs = kvs.Add("f-2", 123, false, false)
			kvs = kvs.Add("f-3", 123, false, false)
			kvs = kvs.Add("f-4", 123, false, false)
			_ = kvs.Add("f-5", 123, false, false)
		}

		// XXX: why set loop to 1, the kvReused == 0?
		assert.True(t, pp.kvReused.Load() > 0)

		t.Logf("point-pool: %s", pp)
	})

	t.Run("trim-tag", func(t *T.T) {
		var kvs KVs
		kvs = kvs.Add("f0", 1.23, false, false)
		kvs = kvs.AddTag("t1", "v1")
		kvs = kvs.AddTag("t2", "v1")
		kvs = kvs.AddTag("t3", "v1")
		kvs = kvs.Add("f1", -123, false, false)

		kvs = kvs.TrimTags(1)

		assert.Lenf(t, kvs, 3, "go kvs: %s", kvs.Pretty())
		assert.NotNil(t, kvs.Get("t1"))
	})

	t.Run("trim-tag-under-point-pool", func(t *T.T) {
		pp := &fullPointPool{}
		SetPointPool(pp)
		defer ClearPointPool()

		for loop := 0; loop < 2; loop++ {
			var kvs KVs

			kvs = kvs.Add("f0", 1.23, false, false)
			kvs = kvs.AddTag("t1", "v1")
			kvs = kvs.AddTag("t2", "v1")
			kvs = kvs.Add("f1", -123, false, false)

			kvs = kvs.TrimTags(1)

			assert.Lenf(t, kvs, 3 /* 2 fields  + 1 tag */, "go kvs: %s", kvs.Pretty())

			assert.NotNil(t, kvs.Get("f0"))
			assert.NotNil(t, kvs.Get("f1"))
			assert.NotNil(t, kvs.Get("f1"))

			kvs = kvs.Add("f-2", 123, false, false)
			kvs = kvs.Add("f-3", 123, false, false)
			kvs = kvs.Add("f-4", 123, false, false)
			_ = kvs.Add("f-5", 123, false, false)
		}

		// XXX: why set loop to 1, the kvReused == 0?
		assert.True(t, pp.kvReused.Load() > 0)

		t.Logf("point-pool: %s", pp)
	})
}

func BenchmarkKVsTrim(b *T.B) {
	b.Run("trim", func(b *T.B) {
		pp := NewPointPoolLevel3()
		SetPointPool(pp)
		defer func() {
			ClearPointPool()
			b.Logf("point pool: %s", pp)
		}()

		for i := 0; i < b.N; i++ {
			var kvs KVs
			kvs = kvs.Add("f0", 1.23, false, false)
			kvs = kvs.AddTag("t1", "v1")
			kvs = kvs.Add("f1", -123, false, false)
			kvs = kvs.Add("f2", uint64(123), false, false)

			kvs = kvs.TrimFields(1)

			kvs = kvs.Add("f-2", 123, false, false)
			kvs = kvs.Add("f-3", 123, false, false)
			kvs = kvs.Add("f-4", 123, false, false)
			_ = kvs.Add("f-5", 123, false, false)
		}
	})

	b.Run("del", func(b *T.B) {
		pp := NewPointPoolLevel3()
		SetPointPool(pp)
		defer func() {
			ClearPointPool()
			b.Logf("point pool: %s", pp)
		}()

		for i := 0; i < b.N; i++ {
			var kvs KVs
			kvs = kvs.Add("f0", 1.23, false, false)
			kvs = kvs.AddTag("t1", "v1")
			kvs = kvs.Add("f1", -123, false, false)
			kvs = kvs.Add("f2", uint64(123), false, false)

			kvs = kvs.Del("f2")

			kvs = kvs.Add("f-2", 123, false, false)
			kvs = kvs.Add("f-3", 123, false, false)
			kvs = kvs.Add("f-4", 123, false, false)
			_ = kvs.Add("f-5", 123, false, false)
		}
	})
}

func TestKVsAdd(t *T.T) {
	t.Run("basic", func(t *T.T) {
		var kvs KVs
		kvs.Add("f1", 123, false, false)

		assert.Len(t, kvs, 0)

		kvs = kvs.Add("f1", 123, false, false)
		assert.Len(t, kvs, 1)
	})

	t.Run("add-v2", func(t *T.T) {
		var kvs KVs
		kvs = kvs.AddV2("f1", 123, false, WithKVUnit("dollar"), WithKVType(GAUGE))
		kvs = kvs.AddV2("cap", 123, false, WithKVUnit("bytes"), WithKVType(COUNT))

		t.Logf("kvs: %s", kvs.Pretty())
	})
}

func TestKVsReset(t *T.T) {
	t.Run("reset", func(t *T.T) {
		var kvs KVs
		kvs = kvs.Add("f0", 1.23, false, false)
		kvs = kvs.Add("f1", -123, false, false)
		kvs = kvs.Add("f2", uint64(123), false, false)
		kvs = kvs.Add("f3", "hello", false, false)
		kvs = kvs.Add("f4", []byte("world"), false, false)
		kvs = kvs.Add("f5", false, false, false)

		kvs.ResetFull()

		assert.Equal(t, "", kvs[0].Key)
		assert.Equal(t, 0.0, kvs[0].Raw().(float64))

		assert.Equal(t, int64(0), kvs[1].Raw().(int64))
		assert.Equal(t, uint64(0), kvs[2].Raw().(uint64))
		assert.Equal(t, "", kvs[3].Raw().(string))
		assert.Len(t, kvs[4].Raw().([]byte), 0)
		assert.False(t, kvs[5].Raw().(bool))
	})
}

func TestKVs(t *T.T) {
	t.Run("add-tag", func(t *T.T) {
		kvs := NewKVs(map[string]any{"f1": 123})

		kvs = kvs.AddTag(`t1`, `v1`)
		assert.Equal(t, `v1`, kvs.Get(`t1`).GetS())
		assert.Equal(t, 1, kvs.TagCount())

		// add new tag t2
		kvs = kvs.Add(`t2`, `v2`, true, true)
		assert.Equal(t, `v2`, kvs.Get(`t2`).GetS())
		assert.Equal(t, 2, kvs.TagCount())

		// replace t2's value v3
		kvs = kvs.Add(`t2`, `v3`, true, true)
		assert.Equal(t, `v3`, kvs.Get(`t2`).GetS())
		assert.Equal(t, 2, kvs.TagCount())

		// invalid tag value(must be []byte/string), switch to field
		kvs = kvs.Add(`tag-as-field`, 123, true, true)
		assert.Equal(t, int64(123), kvs.Get(`tag-as-field`).GetI())
		assert.Equal(t, 2, kvs.TagCount())

		// invalid tag override exist
		kvs = kvs.Add(`t2`, false, true, true)
		assert.Equal(t, false, kvs.Get(`t2`).GetB())
		assert.Equal(t, 1, kvs.TagCount())
	})

	t.Run(`new-empty`, func(t *T.T) {
		kvs := NewKVs(nil)
		assert.Equal(t, 0, len(kvs))
	})

	t.Run(`new-invalid-float`, func(t *T.T) {
		kvs := NewKVs(map[string]any{
			"f1": math.NaN(),
			"f2": math.Inf(1),
		})

		assert.Equal(t, 2, len(kvs))
	})

	t.Run(`new-all-types`, func(t *T.T) {
		kvs := NewKVs(map[string]any{
			"f1": 123,
			"f2": uint64(123),
			"f3": 3.14,
			"f4": "hello",
			"f5": []byte(`world`),
			"f6": false,
			"f7": true,
		})
		assert.Equal(t, 7, len(kvs))

		assert.Equal(t, int64(123), kvs.Get(`f1`).GetI())
		assert.Equal(t, uint64(123), kvs.Get(`f2`).GetU())
		assert.Equal(t, 3.14, kvs.Get(`f3`).GetF())
		assert.Equal(t, `hello`, kvs.Get(`f4`).GetS())
		assert.Equal(t, []byte(`world`), kvs.Get(`f5`).GetD())
		assert.Equal(t, false, kvs.Get(`f6`).GetB())
		assert.Equal(t, true, kvs.Get(`f7`).GetB())

		t.Logf("kvs:\n%s", kvs.Pretty())
	})

	t.Run(`add-kv`, func(t *T.T) {
		kvs := NewKVs(nil)

		kvs = kvs.MustAddKV(NewKV(`t1`, false, WithKVTagSet(true))) // set tag failed on bool value
		kvs = kvs.MustAddKV(NewKV(`t2`, "v1", WithKVTagSet(true)))
		kvs = kvs.MustAddKV(NewKV(`t3`, []byte("v2"), WithKVTagSet(true)))

		kvs = kvs.MustAddKV(NewKV(`f1`, "foo"))
		kvs = kvs.MustAddKV(NewKV(`f2`, 123, WithKVUnit("MB"), WithKVType(COUNT)))
		kvs = kvs.MustAddKV(NewKV(`f3`, 3.14, WithKVUnit("some"), WithKVType(GAUGE)))

		assert.Equal(t, 6, len(kvs))

		t.Logf("kvs:\n%s", kvs.Pretty())
	})

	// any update to kvs should keep them sorted
	t.Run(`test-not-sorted`, func(t *T.T) {
		kvs := NewKVs(nil)

		assert.True(t, sort.IsSorted(kvs)) // empty kvs sorted

		kvs = kvs.Add(`f2`, false, false, false)
		kvs = kvs.Add(`f1`, 123, false, false)
		kvs = kvs.Add(`f0`, 123, false, false)
		kvs = kvs.MustAddTag(`t1`, "v1")

		assert.False(t, sort.IsSorted(kvs))

		kvs = kvs.Del(`f1`)
		assert.False(t, sort.IsSorted(kvs))

		kvs = kvs.MustAddKV(NewKV(`f3`, 3.14))
		assert.False(t, sort.IsSorted(kvs))

		t.Logf("kvs:\n%s", kvs.Pretty())

		sort.Sort(kvs)
		assert.True(t, sort.IsSorted(kvs))
		assert.Len(t, kvs, 4)
	})

	t.Run(`test-del3`, func(t *T.T) {
		pp := NewPointPoolLevel3()

		SetPointPool(pp)
		defer ClearPointPool()

		var kvs KVs

		defer func() {
			for _, kv := range kvs {
				pp.PutKV(kv)
			}
		}()

		kvs = kvs.Add(`f1`, false, false, false)
		kvs = kvs.Add(`f2`, 123, false, false)
		kvs = kvs.Add(`f3`, 123, false, false)

		t.Logf("kvs:\n%s", kvs.Pretty())
		kvs = kvs.Del(`f1`)
		t.Logf("kvs:\n%s", kvs.Pretty())

		t.Logf("pt pool: %s", pp)
	})

	t.Run(`test-del-on-pt-pool`, func(t *T.T) {
		pp := NewPointPoolLevel3()

		SetPointPool(pp)
		defer ClearPointPool()

		var kvs KVs

		defer func() {
			for _, kv := range kvs {
				pp.PutKV(kv)
			}
		}()

		kvs = kvs.Add(`f1`, false, false, false)
		kvs = kvs.Add(`f2`, 123, false, false)
		kvs = kvs.Add(`f3`, 123, false, false)

		t.Logf("kvs:\n%s", kvs.Pretty())
		kvs = kvs.Del(`f1`)
		t.Logf("kvs:\n%s", kvs.Pretty())

		t.Logf("pt pool: %s", pp)
	})

	t.Run(`test-update-on-kvs`, func(t *T.T) {
		pt := NewPointV2("ptname", nil)

		pt.pt.Fields = KVs(pt.pt.Fields).Add("f1", 1.23, false, false)

		t.Logf("point: %s", pt.Pretty())

		assert.NotNil(t, pt.Get("f1"))
	})
}

func TestKVsDel(t *T.T) {
	t.Run("del", func(t *T.T) {
		var kvs KVs

		kvs = kvs.Add(`f1`, false, false, false)
		kvs = kvs.Add(`f2`, 123, false, false)
		kvs = kvs.Add(`f3`, 123, false, false)

		kvs = kvs.Del(`f1`)
		assert.Len(t, kvs, 2)
		kvs = kvs.Del(`f3`)
		assert.Len(t, kvs, 1)
		assert.NotNil(t, kvs.Get(`f2`))
	})

	t.Run(`del-on-point-pool`, func(t *T.T) {
		var kvs KVs

		pp := &fullPointPool{}
		SetPointPool(pp)
		defer func() {
			ClearPointPool()
		}()

		kvs = kvs.Add(`f1`, false, false, false)
		kvs = kvs.Add(`f2`, 123, false, false)
		kvs = kvs.Add(`f3`, 123, false, false)

		kvs = kvs.Del(`f1`)
		assert.Len(t, kvs, 2)
		kvs = kvs.Del(`f3`)

		assert.Len(t, kvs, 1)
		assert.NotNil(t, kvs.Get(`f2`))

		_ = kvs.Add(`f-x`, 123, false, false)

		assert.True(t, pp.kvReused.Load() > 0) // key f-x reused
		assert.True(t, pp.kvCreated.Load() > 0)

		t.Logf("point pool: %s", pp.String())
	})
}

func BenchmarkKVsDel(b *T.B) {
	addTestKVs := func(kvs KVs) KVs {
		kvs = kvs.Add(`f1`, false, false, false)
		kvs = kvs.Add(`f2`, 123, false, false)
		kvs = kvs.Add(`f3`, "some string", false, false)
		kvs = kvs.Add(`f4`, []byte("hello world"), false, false)
		kvs = kvs.Add(`f5`, 3.14, false, false)
		kvs = kvs.Add(`f6`, uint(8), false, false)

		return kvs
	}

	b.Run("del-on-slice-Delete", func(b *T.B) {
		for i := 0; i < b.N; i++ {
			var kvs KVs
			kvs = addTestKVs(kvs)
			_ = kvs.Del(`f1`)
		}
	})

	b.Run("del-on-slice-Delete-with-point-pool", func(b *T.B) {
		pp := NewPointPoolLevel3()
		SetPointPool(pp)
		defer func() {
			b.Logf("point pool: %s", pp)
			ClearPointPool()
		}()

		for i := 0; i < b.N; i++ {
			var kvs KVs
			kvs = addTestKVs(kvs)
			_ = kvs.Del(`f1`)
		}
	})

	b.Run("del-on-new-slice", func(b *T.B) {
		del := func(kvs KVs, k string) KVs {
			var keep KVs // new slice
			for _, f := range kvs {
				if f.Key != k {
					keep = append(keep, f)
				}
			}
			return keep
		}

		for i := 0; i < b.N; i++ {
			var kvs KVs
			kvs = addTestKVs(kvs)
			_ = del(kvs, `f1`)
		}
	})
}
