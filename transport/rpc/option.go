package rpc

import (
	"google.golang.org/grpc"

	"lego/transport/naming"
)

type (
	RegisterFunc func(*grpc.Server)

	Option func(o *options)
)

type options struct {
	// Target
	Target string

	// Timeout
	Timeout int64

	// Naming
	Naming naming.Naming

	// RegisterFunc
	RegisterFunc RegisterFunc
}

// WithTarget sets target.
func WithTarget(target string) Option {
	return func(o *options) {
		o.Target = target
	}
}

// WithTimeout set timeout.
func WithTimeout(timeout int64) Option {
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

func WithRegisterFunc(f RegisterFunc) Option {
	return func(o *options) {
		o.RegisterFunc = f
	}
}
