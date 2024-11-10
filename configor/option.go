package configor

import (
	"lego/encoding"
)

type ChangeNotify func(ChangeSet)

type Option func(o *options)

type options struct {
	// Encoding
	Encoding encoding.Encoding

	// Source
	Source int

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

// WithEtcd sets source etcd.
func WithEtcd() Option {
	return func(o *options) {
		o.Source = ETCD
	}
}

// WithApollo sets source apollo.
func WithApollo() Option {
	return func(o *options) {
		o.Source = APOLLO
	}
}

// WithNacos sets source nacos.
func WithNacos() Option {
	return func(o *options) {
		o.Source = NACOS
	}
}

func WithWatch(f ChangeNotify) Option {
	return func(o *options) {
		o.WatchChange = f
	}
}
