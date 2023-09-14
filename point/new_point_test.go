// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package point

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
	T "testing"
	"time"

	"github.com/GuanceCloud/cliutils"
	"github.com/stretchr/testify/assert"
)

func getNFields(n int) map[string]interface{} {
	i := 0
	fields := map[string]interface{}{}
	for {
		var v interface{}
		v = i // int

		if i%2 == 0 { // string
			v = fmt.Sprintf("fieldv-%d", i)
		}

		if i%3 == 0 { // float
			v = rand.NormFloat64()
		}

		if i%4 == 0 { // bool
			if i/2%2 == 0 {
				v = false
			} else {
				v = true
			}
		}

		fields[fmt.Sprintf("field-%d", i)] = v
		if i > n {
			return fields
		}

		i++
	}
}

func TestV2NewPoint(t *T.T) {
	t.Run("valid-fields", func(t *T.T) {
		pt := NewPointV2("abc", NewKVs(
			map[string]interface{}{
				"[]byte":  []byte("abc"),
				"[]uint8": []uint8("abc"),

				"b-false":   false,
				"b-true":    true,
				"float":     1.0,
				"float32":   float32(1.0),
				"float64":   float64(1.0),
				"float64-2": float64(1.1),
				"i":         int(1),
				"i16":       int16(1),
				"i32":       int32(1),
				"i64":       int64(1),
				"i8":        int8(1),
				"u":         uint(1),
				"u16":       uint16(1),
				"u32":       uint32(1),
				"u64-large": uint64(math.MaxInt64 + 1), // skipped in expect string
				"u64":       uint64(1),
				"u8":        uint8(1),
			}), WithTime(time.Unix(0, 123)))

		assert.Equal(t, `abc []byte="abc",[]uint8="abc",b-false=false,b-true=true,float=1,float32=1,float64=1,float64-2=1.1,i=1i,i16=1i,i32=1i,i64=1i,i8=1i,u=1i,u16=1i,u32=1i,u64=1i,u8=1i 123`,
			pt.LineProto())
	})

	t.Run("valid-fields-under-pb", func(t *T.T) {
		kvs := map[string]interface{}{
			"[]byte":    []byte("abc"),
			"[]uint8":   []uint8("abc"),
			"b-false":   false,
			"b-true":    true,
			"float":     1.0,
			"float32":   float32(1.0),
			"float64":   float64(1.0),
			"float64-2": float64(1.1),
			"i":         int(1),
			"i16":       int16(1),
			"i32":       int32(1),
			"i64":       int64(1),
			"i8":        int8(1),
			"u":         uint(1),
			"u16":       uint16(1),
			"u32":       uint32(1),
			"u64":       uint64(1),
			"u64-large": uint64(math.MaxInt64 + 1), // skipped in expect string
			"u8":        uint8(1),
		}
		pt := NewPointV2(`abc`, NewKVs(kvs), WithTime(time.Unix(0, 123)), WithEncoding(Protobuf))
		expect := `abc []byte="abc",[]uint8="abc",b-false=false,b-true=true,float=1,float32=1,float64=1,float64-2=1.1,i=1i,i16=1i,i32=1i,i64=1i,i8=1i,u=1i,u16=1i,u32=1i,u64=1i,u8=1i 123`
		assert.Equal(t, expect, pt.LineProto())
	})

	t.Run("basic", func(t *T.T) {
		kvs := NewKVs(map[string]interface{}{"f1": 12}).MustAddTag(`t1`, `tval1`)
		pt := NewPointV2(`abc`, kvs, WithTime(time.Unix(0, 123)))

		assert.Equal(t, "abc,t1=tval1 f1=12i 123", pt.LineProto())
	})
}

func TestNewPoint(t *T.T) {
	cases := []struct {
		opts []Option

		tname, name, expect string

		t map[string]string
		f map[string]interface{}

		warns int
	}{
		{
			tname: "valid-fields",
			opts:  []Option{WithTime(time.Unix(0, 123))},
			name:  "valid-fields",
			f: map[string]interface{}{
				"[]byte":  []byte("abc"),
				"[]uint8": []uint8("abc"),

				"b-false":   false,
				"b-true":    true,
				"float":     1.0,
				"float32":   float32(1.0),
				"float64":   float64(1.0),
				"float64-2": float64(1.1),
				"i":         int(1),
				"i16":       int16(1),
				"i32":       int32(1),
				"i64":       int64(1),
				"i8":        int8(1),
				"u":         uint(1),
				"u16":       uint16(1),
				"u32":       uint32(1),
				"u64":       uint64(1),
				"u64-large": uint64(math.MaxInt64 + 1), // skipped in expect string
				"u8":        uint8(1),
			},
			expect: `valid-fields []byte="abc",[]uint8="abc",b-false=false,b-true=true,float=1,float32=1,float64=1,float64-2=1.1,i=1i,i16=1i,i32=1i,i64=1i,i8=1i,u=1i,u16=1i,u32=1i,u64=1i,u8=1i 123`,
		},

		{
			tname: "valid-fields-under-pb",
			opts:  []Option{WithTime(time.Unix(0, 123)), WithEncoding(Protobuf)},
			name:  "valid-fields",
			f: map[string]interface{}{
				"[]byte":    []byte("abc"),
				"[]uint8":   []uint8("abc"),
				"b-false":   false,
				"b-true":    true,
				"float":     1.0,
				"float32":   float32(1.0),
				"float64":   float64(1.0),
				"float64-2": float64(1.1),
				"i":         int(1),
				"i16":       int16(1),
				"i32":       int32(1),
				"i64":       int64(1),
				"i8":        int8(1),
				"u":         uint(1),
				"u16":       uint16(1),
				"u32":       uint32(1),
				"u64":       uint64(1),
				"u64-large": uint64(math.MaxInt64 + 1), // skipped in expect string
				"u8":        uint8(1),
			},
			expect: `valid-fields []byte="abc",[]uint8="abc",b-false=false,b-true=true,float=1,float32=1,float64=1,float64-2=1.1,i=1i,i16=1i,i32=1i,i64=1i,i8=1i,u=1i,u16=1i,u32=1i,u64=1i,u8=1i 123`,
		},

		{
			tname: "exceed-measurement-len",
			opts:  []Option{WithTime(time.Unix(0, 123)), WithMaxMeasurementLen(10)},

			name:   "name-exceed-len",
			f:      map[string]interface{}{"f1": 123},
			expect: `name-excee f1=123i 123`,
			warns:  1,
		},

		{
			tname:  "empty-measurement",
			opts:   []Option{WithTime(time.Unix(0, 123))},
			name:   "",
			f:      map[string]interface{}{"f1": 123},
			expect: fmt.Sprintf(`%s f1=123i 123`, DefaultMeasurementName),
			warns:  1,
		},

		//{
		//	tname:  "exceed-tag-kv-compose",
		//	opts:   []Option{WithTime(time.Unix(0, 123)), WithMaxKVComposeLen(10)},
		//	name:   "abc",
		//	t:      map[string]string{"t1": "12345", "t2": "ssclh"},
		//	f:      map[string]interface{}{"f1": 123},
		//	expect: `abc,t1=12345 f1=123i 123`,
		//	warns:  1,
		// },

		{
			tname:  "basic",
			opts:   []Option{WithTime(time.Unix(0, 123))},
			name:   "abc",
			t:      map[string]string{"t1": "tval1"},
			f:      map[string]interface{}{"f1": 12},
			expect: "abc,t1=tval1 f1=12i 123",
		},
		{
			tname:  "metric-with-dot-in-field-key",
			name:   "abc",
			opts:   append(DefaultMetricOptions(), WithTime(time.Unix(0, 123))),
			t:      map[string]string{"t1": "tval1"},
			f:      map[string]interface{}{"f.1": 12},
			expect: "abc,t1=tval1 f.1=12i 123",
		},
		{
			tname:  "metric-with-dot-in-tag-key",
			name:   "abc",
			opts:   append(DefaultMetricOptions(), WithTime(time.Unix(0, 123))),
			t:      map[string]string{"t.1": "tval1"},
			f:      map[string]interface{}{"f1": 12},
			expect: "abc,t.1=tval1 f1=12i 123",
		},
		{
			tname: "with-dot-in-t-f-key-on-non-metric-type",
			name:  "abc",
			opts:  append(DefaultObjectOptions(), WithTime(time.Unix(0, 123))),

			t:      map[string]string{"t1": "tval1"},
			f:      map[string]interface{}{"f.1": 12},
			expect: fmt.Sprintf(`abc,t1=tval1 f_1=12i,name="%s" 123`, defaultObjectName),
			warns:  2,
		},

		{
			tname:  "with-dot-in-tag-field-key",
			name:   "abc",
			opts:   append(DefaultObjectOptions(), WithTime(time.Unix(0, 123))),
			t:      map[string]string{"t1": "abc", "t.2": "xyz"},
			f:      map[string]interface{}{"f1": 123, "f.2": "def"},
			expect: fmt.Sprintf(`abc,t1=abc,t_2=xyz f1=123i,f_2="def",name="%s" 123`, defaultObjectName),
			warns:  3,
		},

		{
			tname: "both-exceed-max-field-tag-count",
			name:  "abc",
			t: map[string]string{
				"t1": "abc",
				"t2": "xyz",
				"t3": "abc",
				"t4": "xyz",
				"t5": "abc",
				"t6": "abc",
				"t7": "abc",
				"t8": "abc",
				"t9": "abc",
			},
			f: map[string]interface{}{
				"f1": 123,
				"f2": "def",
				"f3": "def",
				"f4": "def",
				"f5": "def",
				"f6": "def",
				"f7": "def",
				"f8": "def",
				"f9": "def",
			},
			opts: []Option{
				WithTime(time.Unix(0, 123)),
				WithMaxTags(1),
				WithMaxFields(1),
				WithKeySorted(true)},
			expect: `abc,t1=abc f1=123i 123`,
			warns:  2,
		},

		{
			tname: "exceed-max-field-count",
			name:  "abc",
			opts:  []Option{WithTime(time.Unix(0, 123)), WithMaxFields(1), WithKeySorted(true)},
			t: map[string]string{
				"t1": "abc",
				"t2": "xyz",
			},
			f: map[string]interface{}{
				"f1": 123,
				"f2": "def",
				"f3": "def",
				"f4": "def",
				"f5": "def",
				"f6": "def",
				"f7": "def",
				"f8": "def",
				"f9": "def",
			},
			expect: `abc,t1=abc,t2=xyz f1=123i 123`,
			warns:  1,
		},

		{
			tname: "exceed-max-tag-count",
			opts:  []Option{WithTime(time.Unix(0, 123)), WithMaxTags(1), WithKeySorted(true)},
			name:  "abc",
			t: map[string]string{
				"t1": "abc",
				"t2": "xyz",
				"t3": "abc",
				"t4": "xyz",
				"t5": "abc",
				"t6": "abc",
				"t7": "abc",
				"t8": "abc",
				"t9": "abc",
			},
			f: map[string]interface{}{
				"f1": 123,
			},
			expect: `abc,t1=abc f1=123i 123`,
			warns:  1,
		},

		{
			tname:  "exceed-max-tag-key-len",
			opts:   []Option{WithTime(time.Unix(0, 123)), WithMaxTagKeyLen(1)},
			name:   "abc",
			t:      map[string]string{"t1": "x"},
			f:      map[string]interface{}{"f1": 123},
			expect: `abc,t=x f1=123i 123`,
			warns:  1,
		},

		{
			tname:  "exceed-max-tag-value-len",
			opts:   []Option{WithTime(time.Unix(0, 123)), WithMaxTagValLen(3)},
			name:   "abc",
			t:      map[string]string{"t": "1234"},
			f:      map[string]interface{}{"f1": 123},
			expect: `abc,t=123 f1=123i 123`,
			warns:  1,
		},

		{
			tname:  "exceed-max-field-key-len",
			name:   "abc",
			opts:   []Option{WithTime(time.Unix(0, 123)), WithMaxFieldKeyLen(3)},
			f:      map[string]interface{}{"f123": 123},
			expect: `abc f12=123i 123`,
			warns:  1,
		},

		{
			tname:  "exceed-max-field-val-len",
			name:   "abc",
			opts:   []Option{WithTime(time.Unix(0, 123)), WithMaxFieldValLen(3)},
			f:      map[string]interface{}{"f1": "hello"},
			expect: `abc f1="hel" 123`,
			warns:  1,
		},

		{
			tname: "with-disabled-tag-key-source",
			name:  "abc",
			opts:  append(DefaultLoggingOptions(), WithTime(time.Unix(0, 123))),

			t:      map[string]string{"source": "s1"},
			f:      map[string]interface{}{"f1": 123},
			expect: fmt.Sprintf(`abc f1=123i,status="%s" 123`, defaultLoggingStatus),
			warns:  2,
		},
		{
			tname:  "with-disabled-field-key",
			name:   "abc",
			opts:   append(DefaultObjectOptions(), WithTime(time.Unix(0, 123))),
			t:      map[string]string{"class": "xyz"},
			f:      map[string]interface{}{"class": 123, "f1": 1},
			expect: fmt.Sprintf(`abc f1=1i,name="%s" 123`, defaultObjectName),

			// NOTE: tag key `class` override field `class`, then the tag disabled
			warns: 2,
		},
		{
			tname:  "normal",
			opts:   []Option{WithTime(time.Unix(0, 123))},
			name:   "abc",
			t:      map[string]string{},
			f:      map[string]interface{}{"f1": 123},
			expect: "abc f1=123i 123",
		},

		{
			tname:  "invalid-category",
			opts:   []Option{WithTime(time.Unix(0, 123))},
			name:   "abc",
			f:      map[string]interface{}{"f1": 123},
			expect: `abc f1=123i 123`,
		},

		{
			tname: "nil-opiton",
			name:  "abc",
			t:     map[string]string{},
			f:     map[string]interface{}{"f1": 123},
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *T.T) {
			pt, err := NewPoint(tc.name, tc.t, tc.f, tc.opts...)

			assert.NoError(t, err)

			x := pt.LineProto()

			if tc.expect != "" {
				assert.Equal(t, tc.expect, x, "pt: %s, kvs: %s", pt.Pretty(), pt.kvs.Pretty())
			} else {
				assert.NotEqual(t, x, "", "got %s", pt.Pretty())
				t.Logf("got %s", x)
			}

			assert.Equal(t, tc.warns, len(pt.warns), "pt: %s", pt.Pretty())
		})
	}
}

func TestPointKeySorted(t *testing.T) {
	t.Run("sorted", func(t *testing.T) {
		pt, err := NewPoint("basic",
			map[string]string{
				"t1": "v1",
				"t2": "v2",
				"t3": "v1",
				"t4": "v2",
				"t5": "v1",
				"t6": "v2",
				"t7": "v1",
				"t8": "v2",
			},
			map[string]any{
				"f1": 1,
				"f2": 2,
				"f3": 3,
				"f4": 1,
				"f5": 2,
				"f6": 3,
				"f7": 1,
				"f8": 2,
				"f9": 3,
			},
			WithKeySorted(true),
		)

		assert.NoError(t, err)

		assert.True(t, sort.IsSorted(pt.kvs))

		t.Logf("pt: %s", pt.Pretty())
	})

	t.Run("not-sorted", func(t *testing.T) {
		pt, err := NewPoint("basic",
			map[string]string{
				"t1": "v1",
				"t2": "v2",
				"t3": "v1",
				"t4": "v2",
				"t5": "v1",
				"t6": "v2",
				"t7": "v1",
				"t8": "v2",
			},
			map[string]any{
				"f1": 1,
				"f2": 2,
				"f3": 3,
				"f4": 1,
				"f5": 2,
				"f6": 3,
				"f7": 1,
				"f8": 2,
				"f9": 3,
			},
			WithKeySorted(false),
		)

		assert.NoError(t, err)

		assert.False(t, sort.IsSorted(pt.kvs))

		t.Logf("pt: %s", pt.Pretty())
	})
}

var (
	__largeKeyForBench = cliutils.CreateRandomString(128)
	__largeValForBench = cliutils.CreateRandomString(1024)
)

var benchCases = []struct {
	name string
	t    map[string]string
	f    map[string]interface{}
}{
	{
		name: "3-tag-10-field",
		t: map[string]string{
			"t1": "val1",
			"t2": "val2",
			"t3": "val3",
		},
		f: map[string]interface{}{
			"f1":  123,
			"f2":  "abc",
			"f3":  45.6,
			"f4":  123,
			"f5":  "abc",
			"f6":  45.6,
			"f7":  123,
			"f8":  "abc",
			"f9":  45.6,
			"f10": false,
		},
	},
	{
		name: "3-tag-10-long-key-field",
		t: map[string]string{
			"t1": "val1",
			"t2": "val2",
			"t3": "val3",
		},
		f: map[string]interface{}{
			"f1" + __largeKeyForBench: 123,
			"f2" + __largeKeyForBench: "abc",
			"f3" + __largeKeyForBench: 45.6,
			"f4" + __largeKeyForBench: 123,
			"f5" + __largeKeyForBench: "abc",
			"f6" + __largeKeyForBench: 45.6,
			"f7" + __largeKeyForBench: 123,
			"f8" + __largeKeyForBench: "abc",
			"f9" + __largeKeyForBench: 45.6,
			"f0" + __largeKeyForBench: 3.14,
		},
	},

	{
		name: "3-tag-10-long-key-tag",
		t: map[string]string{
			"t1" + __largeKeyForBench: "val1",
			"t2" + __largeKeyForBench: "val2",
			"t3" + __largeKeyForBench: "val3",
			"t4" + __largeKeyForBench: "val4",
			"t5" + __largeKeyForBench: "val5",
			"t6" + __largeKeyForBench: "val6",
			"t7" + __largeKeyForBench: "val7",
			"t8" + __largeKeyForBench: "val8",
			"t9" + __largeKeyForBench: "val9",
			"t0" + __largeKeyForBench: "val0",
		},
		f: map[string]interface{}{
			"f1":  123,
			"f2":  "abc",
			"f3":  45.6,
			"f4":  123,
			"f5":  "abc",
			"f6":  45.6,
			"f7":  123,
			"f8":  "abc",
			"f9":  45.6,
			"f10": false,
		},
	},
}

func BenchmarkNewPoint(b *T.B) {
	b.Run(`with-pool-cfg`, func(b *T.B) {
		ptName := `abc`
		kvs := NewKVs(map[string]any{"f1": 123, "f2": 3.14})
		for i := 0; i < b.N; i++ {
			NewPointV2(ptName, kvs)
		}
	})

	b.Run(`without-pool-cfg`, func(b *T.B) {
		ptName := `abc`
		kvs := NewKVs(map[string]any{"f1": 123, "f2": 3.14})
		for i := 0; i < b.N; i++ {
			doNewPoint(ptName, kvs, newCfg()) // slower ~17% than pooled
		}
	})

	b.Run(`with-key-sorted`, func(b *T.B) {
		ptName := `abc`
		kvs := NewKVs(map[string]any{
			"f1": 123,
			"f2": 3.14,
			"f3": "str",

			"_f1": 123,
			"_f2": 3.14,
			"_f3": "str",

			"_f1_": 123,
			"_f2_": 3.14,
			"_f3_": "str",
		})

		for i := 0; i < b.N; i++ {
			NewPointV2(ptName, kvs, WithKeySorted(true))
		}
	})

	b.Run(`without-key-sorted`, func(b *T.B) {
		ptName := `abc`
		kvs := NewKVs(map[string]any{
			"f1": 123,
			"f2": 3.14,
			"f3": "str",

			"_f1": 123,
			"_f2": 3.14,
			"_f3": "str",

			"_f1_": 123,
			"_f2_": 3.14,
			"_f3_": "str",
		})

		for i := 0; i < b.N; i++ {
			NewPointV2(ptName, kvs)
		}
	})
}

func FuzzPLPBEquality(f *testing.F) {
	cases := []struct {
		measurement string
		tagk        string
		tagv        string

		i64  int64
		u64  uint64
		str  string
		b    bool
		f    float64
		d    []byte
		time int64
	}{
		{
			measurement: "",
			tagk:        "tag",
			tagv:        "tval",

			i64:  int64(1),
			u64:  uint64(123),
			str:  "hello",
			b:    false,
			f:    3.14,
			d:    []byte("world"),
			time: 123,
		},
	}

	for _, tc := range cases {
		f.Add(tc.measurement, tc.tagk, tc.tagv, tc.i64, tc.u64, tc.str, tc.b, tc.f, tc.d, tc.time)
	}

	f.Fuzz(func(t *testing.T,
		measurement, tagk, tagv string,
		i64 int64, u64 uint64, str string, b bool, f float64, d []byte, ts int64,
	) {
		lppt, err := NewPoint(measurement,
			map[string]string{tagk: tagv},
			map[string]interface{}{
				"i64": i64,
				"u64": u64,
				"str": str,
				"b":   b,
				"f":   f,
				"d":   d, // dropped under line-protocol
			},
			WithTime(time.Unix(0, ts)),
			WithDotInKey(true) /* random string may contains '.' */)

		assert.NoError(t, err)

		pbpt, err := NewPoint(measurement,
			map[string]string{tagk: tagv},
			map[string]interface{}{
				"i64": i64,
				"u64": u64,
				"str": str,
				"b":   b,
				"f":   f,
				"d":   d,
			},
			WithTime(time.Unix(0, ts)),
			WithDotInKey(true), // random string may contains '.'
			WithEncoding(Protobuf))

		assert.NoError(t, err)

		_ = pbpt
		_ = lppt
	})
}
