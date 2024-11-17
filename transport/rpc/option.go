package rpc

import (
	"google.golang.org/grpc"

	"github.com/joshqu1985/lego/transport/naming"
)

type (
	HandlersRegister func(*grpc.Server)

	Option func(o *options)
)

type options struct {
	// Timeout
	Timeout int64

	// Naming
	Naming naming.Naming

	// HandlersRegister
	HandlersRegister HandlersRegister
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

func WithHandlers(f HandlersRegister) Option {
	return func(o *options) {
		o.HandlersRegister = f
	}
}
