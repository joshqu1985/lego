package rest

import (
	"time"

	"github.com/joshqu1985/lego/transport/naming"
)

type (
	RouterRegister func(*Router)

	Option func(o *options)

	options struct {
		Naming         naming.Naming
		RouterRegister func(*Router)
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

// WithRouters set router handlers.
func WithRouters(f RouterRegister) Option {
	return func(o *options) {
		o.RouterRegister = f
	}
}

// WithNaming set naming.
func WithNaming(n naming.Naming) Option {
	return func(o *options) {
		o.Naming = n
	}
}

// WithMetrics set metrics open.
func WithMetrics() Option {
	return func(o *options) {
		o.Metrics = true
	}
}
