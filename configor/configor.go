package configor

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

const (
	ENCODING_JSON = "json"
	ENCODING_YAML = "yaml"
	ENCODING_TOML = "toml"

	SOURCE_ETCD   = "etcd"
	SOURCE_APOLLO = "apollo"
	SOURCE_NACOS  = "nacos"
)

type (
	Configor interface {
		Load(v any) error
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
	 *   app_id:  etcd的key前缀.
	 */
	SourceConfig struct {
		Source    string   `json:"source"     toml:"source"     yaml:"source"`
		Cluster   string   `json:"cluster"    toml:"cluster"    yaml:"cluster"`
		AppId     string   `json:"app_id"     toml:"app_id"     yaml:"app_id"`
		Namespace string   `json:"namespace"  toml:"namespace"  yaml:"namespace"`
		AccessKey string   `json:"access_key" toml:"access_key" yaml:"access_key"`
		SecretKey string   `json:"secret_key" toml:"secret_key" yaml:"secret_key"`
		Endpoints []string `json:"endpoints"  toml:"endpoints"  yaml:"endpoints"`
	}

	ChangeSet struct {
		Timestamp time.Time
		Value     []byte
	}
)

func New(file string, opts ...Option) (Configor, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if option.Encoding == nil {
		option.Encoding = encoding.New(ENCODING_TOML)
	}

	conf, err := ReadSourceConfig(file, option.Encoding)
	if err != nil {
		return nil, err
	}

	switch conf.Source {
	case SOURCE_ETCD:
		return NewEtcd(conf, option)
	case SOURCE_APOLLO:
		return NewApollo(conf, option)
	case SOURCE_NACOS:
		return NewNacos(conf, option)
	default:
		return NewLocal(file, option)
	}
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

func (c *ChangeSet) Sum(data []byte) string {
	h := sha256.New()
	_, _ = h.Write(data)

	return hex.EncodeToString(h.Sum(nil))
}
