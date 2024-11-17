package configor

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

const (
	LOCAL  = 0 // local file
	ETCD   = 1 // etcd
	APOLLO = 2 // apollo
	NACOS  = 3 // nacos
)

// ErrUnknowType 暂时不支持的类型错误
var ErrUnknowType = errors.New("unknow config source type")

type Configor interface {
	Load(v any) error
}

func New(file string, opts ...Option) Configor {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if option.Encoding == nil {
		option.Encoding = encoding.New("toml")
	}

	conf, err := ReadSourceConfig(file, option.Encoding)
	if err != nil {
		panic(err)
	}

	var configor Configor
	switch option.Source {
	case LOCAL:
		configor, err = NewLocal(file, option)
	case ETCD:
		configor, err = NewEtcd(conf, option)
	case APOLLO:
		configor, err = NewApollo(conf, option)
	case NACOS:
		configor, err = NewNacos(conf, option)
	default:
		configor, err = nil, ErrUnknowType
	}

	if err != nil {
		panic(err)
	}
	return configor
}

type SourceConfig struct {
	Endpoints []string `json:"endpoints" yaml:"endpoints" toml:"endpoints"`    // apollo ip        | nacos addr        | etcd endpoints
	Cluster   string   `json:"cluster" yaml:"cluster" toml:"cluster"`          // apollo cluster   | nacos namespaceid | etcd key前缀
	AppId     string   `json:"app_id" yaml:"app_id" toml:"app_id"`             // apollo appId     | nacos DataId      | etcd key前缀
	Namespace string   `json:"namespace" yaml:"namespace" toml:"namespace"`    // apollo namespace |  \                |  \
	AccessKey string   `json:"access_key" yaml:"access_key" toml:"access_key"` //  \               | nacos access_key  | etcd username
	SecretKey string   `json:"secret_key" yaml:"secret_key" toml:"secret_key"` // apollo secret    | nacos secret_key  | etcd password
}

func ReadSourceConfig(file string, encoding encoding.Encoding) (*SourceConfig, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	c := &SourceConfig{}
	return c, encoding.Unmarshal(data, c)
}

type ChangeSet struct {
	Timestamp time.Time
	Value     []byte
}

func (c *ChangeSet) Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
