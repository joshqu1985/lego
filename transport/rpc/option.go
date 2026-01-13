package rpc

import (
	"time"

	"google.golang.org/grpc"

	"github.com/joshqu1985/lego/transport/naming"
)

type (
	RouterRegister func(*grpc.Server)

	Option func(o *options)

	options struct {
		Naming         naming.Naming
		RouterRegister RouterRegister
		Timeout        time.Duration
		Metrics        bool
	}
)

// WithTimeout set timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.Timeout = timeout
	}
}

// WithNaming sets naming.
func WithNaming(n naming.Naming) Option {
	return func(o *options) {
		o.Naming = n
	}
}

// WithRouters register grpc handler.
func WithRouters(f RouterRegister) Option {
	return func(o *options) {
		o.RouterRegister = f
	}
}

// WithMetrics set metrics open.
func WithMetrics() Option {
	return func(o *options) {
		o.Metrics = true
	}
}
