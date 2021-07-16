package tracer

import (
	"fmt"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

type Option func(t *Tracer)

func WithEnv(env string) Option {
	return func(t *Tracer) {
		t.env = env
	}
}

func WithAgentAddr(host string, port int) Option {
	addr := fmt.Sprintf("%s:%d", host, port)
	return func(t *Tracer) {
		t.Host = host
		t.Port = port
		t.addr = addr
	}
}

func WithService(name, version string) Option {
	return func(t *Tracer) {
		t.Service = name
		t.Version = version
	}
}

func WithDebug(debug bool) Option {
	return func(t *Tracer) {
		t.Debug = debug
	}
}

func WithLogger(logger ddtrace.Logger) Option {
	return func(t *Tracer) {
		t.logger = logger
	}
}

// func WithFinishTime(t time.Time) tracer.FinishOption {
// 	return tracer.FinishTime(t)
// }

// func WithError(err error) tracer.FinishOption {
// 	return tracer.WithError(err)
// }
