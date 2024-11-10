package naming

import (
	"sync"
)

type Naming interface {
	Endpoints() []string
	Name() string
	Register(key, val string) error
	Service(key string) RegService
}

type RegService interface {
	Name() string
	Addrs() ([]string, error)
	AddListener(func())
}

type Config struct {
	Endpoints []string `json:"endpoints" yaml:"endpoints" toml:"endpoints"`    // nacos addr        | etcd endpoints
	Cluster   string   `json:"cluster" yaml:"cluster" toml:"cluster"`          // nacos namespaceid | etcd key前缀
	AppId     string   `json:"app_id" yaml:"app_id" toml:"app_id"`             // nacos DataId      | -
	AccessKey string   `json:"access_key" yaml:"access_key" toml:"access_key"` // nacos access_key  | etcd username
	SecretKey string   `json:"secret_key" yaml:"secret_key" toml:"secret_key"` // nacos secret_key  | etcd password
}

var (
	globalNaming Naming
	once         sync.Once
)

func New(conf Config) Naming {
	once.Do(func() {
		globalNaming = NewEtcd(&conf)
	})
	return globalNaming
}
