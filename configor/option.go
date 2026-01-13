package configor

import (
	"github.com/joshqu1985/lego/encoding"
)

type (
	ChangeNotify func(ChangeSet)

	Option func(o *options)

	options struct {
		// Encoding
		Encoding encoding.Encoding

		// WatchChange
		WatchChange ChangeNotify
	}
)

// WithToml sets encoding toml.
func WithToml() Option {
	return func(o *options) {
		o.Encoding = encoding.New(ENCODING_TOML)
	}
}

// WithYaml sets encoding yaml.
func WithYaml() Option {
	return func(o *options) {
		o.Encoding = encoding.New(ENCODING_YAML)
	}
}

// WithJson sets encoding json.
func WithJson() Option {
	return func(o *options) {
		o.Encoding = encoding.New(ENCODING_JSON)
	}
}

func WithWatch(f ChangeNotify) Option {
	return func(o *options) {
		o.WatchChange = f
	}
}
