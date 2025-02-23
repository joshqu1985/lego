package configor

import (
	"github.com/joshqu1985/lego/encoding"
)

type ChangeNotify func(ChangeSet)

type Option func(o *options)

type options struct {
	// Encoding
	Encoding encoding.Encoding

	// WatchChange
	WatchChange ChangeNotify
}

// WithToml sets encoding toml.
func WithToml() Option {
	return func(o *options) {
		o.Encoding = encoding.New("toml")
	}
}

// WithYaml sets encoding yaml.
func WithYaml() Option {
	return func(o *options) {
		o.Encoding = encoding.New("yaml")
	}
}

// WithJson sets encoding json.
func WithJson() Option {
	return func(o *options) {
		o.Encoding = encoding.New("json")
	}
}

func WithWatch(f ChangeNotify) Option {
	return func(o *options) {
		o.WatchChange = f
	}
}
