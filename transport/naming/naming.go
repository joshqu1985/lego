package naming

import (
	"errors"
	"fmt"
)

const (
	SOURCE_NACOS = "nacos"
	SOURCE_ETCD  = "etcd"
)

type (
	Naming interface {
		Endpoints() []string
		Name() string
		Register(key, val string) error
		Deregister(key string) error
		Service(key string) RegService
	}

	RegService interface {
		Name() string
		Addrs() ([]string, error)
		AddListener(func())
	}

	Config struct {
		Source    string   `json:"source"     toml:"source"     yaml:"source"`
		Cluster   string   `json:"cluster"    toml:"cluster"    yaml:"cluster"`
		AccessKey string   `json:"access_key" toml:"access_key" yaml:"access_key"`
		SecretKey string   `json:"secret_key" toml:"secret_key" yaml:"secret_key"`
		Endpoints []string `json:"endpoints"  toml:"endpoints"  yaml:"endpoints"`
	}
)

var ErrClientNil = errors.New("client is nil")

func New(conf *Config) (Naming, error) {
	switch conf.Source {
	case SOURCE_NACOS:
		return NewNacos(conf)
	case SOURCE_ETCD:
		return NewEtcd(conf)
	default:
		return nil, fmt.Errorf("naming: invalid source %s, supported: %s, %s",
			conf.Source, SOURCE_NACOS, SOURCE_ETCD)
	}
}
