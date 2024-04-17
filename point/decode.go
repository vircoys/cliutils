// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package point

import (
	"encoding/json"
	sync "sync"
	"time"

	"google.golang.org/protobuf/proto"
)

var decPool sync.Pool

// DecodeFn used to iterate on []*Point payload, if error returned, the iterate terminated.
type DecodeFn func([]*Point) error

type DecoderOption func(e *Decoder)

func WithDecEncoding(enc Encoding) DecoderOption {
	return func(d *Decoder) { d.enc = enc }
}

func WithDecFn(fn DecodeFn) DecoderOption {
	return func(d *Decoder) { d.fn = fn }
}

type Decoder struct {
	enc Encoding
	fn  DecodeFn

	// For line-protocol parsing, keep original error.
	detailedError error
}

func GetDecoder(opts ...DecoderOption) *Decoder {
	v := decPool.Get()
	if v == nil {
		v = newDecoder()
	}

	x := v.(*Decoder)

	for _, opt := range opts {
		if opt != nil {
			opt(x)
		}
	}

	return x
}

func PutDecoder(d *Decoder) {
	d.reset()
	decPool.Put(d)
}

func newDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) reset() {
	d.enc = 0
	d.fn = nil
	d.detailedError = nil
}

func (d *Decoder) Decode(data []byte, opts ...Option) ([]*Point, error) {
	var (
		pts []*Point
		err error
	)

	// point options
	cfg := GetCfg(opts...)
	defer PutCfg(cfg)

	switch d.enc {
	case JSON:
		var (
			ts  = time.Now()
			arr []*Point
		)

		if err := json.Unmarshal(data, &arr); err != nil {
			return nil, err
		}

		for _, x := range arr {
			ptts := x.time.UnixNano()
			if ptts > 0 { // check if precision attached
				switch cfg.precision {
				case US:
					ptts *= int64(time.Microsecond)
				case MS:
					ptts *= int64(time.Millisecond)
				case S:
					ptts *= int64(time.Second)
				case M:
					ptts *= int64(time.Minute)
				case H:
					ptts *= int64(time.Hour)

				case NS:
					// pass

				case W, D: // not used

				default:
					// pass
				}

				x.time = time.Unix(0, ptts)
			} else {
				x.time = ts
			}

			pts = append(pts, NewPointV2(x.name, x.kvs, append(opts, WithTime(x.time))...))
		}

	case Protobuf:
		var pbpts PBPoints
		if err = proto.Unmarshal(data, &pbpts); err != nil {
			return nil, err
		}

		for _, pbpt := range pbpts.Arr {
			pt := &Point{
				name:   pbpt.Name,
				kvs:    pbpt.Fields,
				time:   time.Unix(0, pbpt.Time),
				warns:  pbpt.Warns,
				debugs: pbpt.Debugs,
			}
			pt.SetFlag(Ppb)

			pts = append(pts, pt)
		}

	case LineProtocol:

		pts, err = parseLPPoints(data, cfg)
		if err != nil {
			d.detailedError = err
			return nil, simplifyLPError(err)
		}
	}

	// check point and apply callbak on each point
	if cfg.precheck || cfg.callback != nil {
		var (
			chk = &checker{cfg: cfg}
			arr []*Point
		)

		for idx, _ := range pts {
			if cfg.precheck {
				pts[idx] = chk.check(pts[idx])
				chk.reset()
			}

			if cfg.callback != nil {
				newPoint, err := cfg.callback(pts[idx])
				if err != nil {
					return nil, err
				}

				if newPoint != nil {
					arr = append(arr, newPoint)
				}
			}
		}

		if cfg.callback != nil {
			pts = arr
		}
	}

	if d.fn != nil {
		return pts, d.fn(pts)
	}

	return pts, nil
}

func (d *Decoder) DetailedError() error {
	return d.detailedError
}
