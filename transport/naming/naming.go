package naming

import (
	"fmt"
	"sync"
)

type Naming interface {
	Endpoints() []string
	Name() string
	Register(key, val string) error
	Deregister(key string) error
	Service(key string) RegService
}

type RegService interface {
	Name() string
	Addrs() ([]string, error)
	AddListener(func())
}

type Config struct {
	Source    string   `json:"source" yaml:"source" toml:"source"`             // nacos | etcd
	Endpoints []string `json:"endpoints" yaml:"endpoints" toml:"endpoints"`    // nacos addr        | etcd endpoints
	Cluster   string   `json:"cluster" yaml:"cluster" toml:"cluster"`          // nacos namespaceid | etcd key前缀
	AccessKey string   `json:"access_key" yaml:"access_key" toml:"access_key"` // nacos access_key  | etcd username
	SecretKey string   `json:"secret_key" yaml:"secret_key" toml:"secret_key"` // nacos secret_key  | etcd password
}

var (
	globalNaming Naming
	once         sync.Once
)

func Get() Naming {
	return globalNaming
}

func Init(conf Config) (Naming, error) {
	var err error
	once.Do(func() {
		globalNaming, err = newNaming(&conf)
	})
	return globalNaming, err
}

func newNaming(conf *Config) (Naming, error) {
	switch conf.Source {
	case "nacos":
		return NewNacos(conf)
	case "etcd":
		return NewEtcd(conf)
	default:
		return nil, fmt.Errorf("invalid source: %s", conf.Source)
	}
}
