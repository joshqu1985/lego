package configor

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

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
	switch conf.Source {
	case "etcd":
		configor, err = NewEtcd(conf, option)
	case "apollo":
		configor, err = NewApollo(conf, option)
	case "nacos":
		configor, err = NewNacos(conf, option)
	default:
		configor, err = NewLocal(file, option)
	}

	if err != nil {
		panic(err)
	}
	return configor
}

/*
 * apollo 配置项
 *   cluster: 应用所属的集群 默认default
 *   app_id: 用来标识应用身份(服务名 *唯一*)
 *   namespace: 配置项的集合(应用的下一层)
 * nacos 配置项
 *   cluster: nacos的namespaceid
 *   app_id: nacos的dataId
 * etcd 配置项
 *   cluster: etcd的key前缀
 *   app_id:  etcd的key前缀
 */
type SourceConfig struct {
	Source    string   `json:"source" yaml:"source" toml:"source"`             // apollo      | nacos        | etcd
	Endpoints []string `json:"endpoints" yaml:"endpoints" toml:"endpoints"`    //  ip         |  addr        |  endpoints
	Cluster   string   `json:"cluster" yaml:"cluster" toml:"cluster"`          //  cluster    |  namespaceid |  key前缀
	AppId     string   `json:"app_id" yaml:"app_id" toml:"app_id"`             //  appId      |  DataId      |  key前缀
	Namespace string   `json:"namespace" yaml:"namespace" toml:"namespace"`    //  namespace  |  \           |  \
	AccessKey string   `json:"access_key" yaml:"access_key" toml:"access_key"` //  \          |  access_key  |  username
	SecretKey string   `json:"secret_key" yaml:"secret_key" toml:"secret_key"` //  secret     |  secret_key  |  password
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
