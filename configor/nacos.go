package configor

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NewNacos(conf *SourceConfig, opts options) (Configor, error) {
	c := &nacosConfig{
		opts: opts,
		conf: conf,
	}

	if err := c.init(conf); err != nil {
		return nil, err
	}

	if _, err := c.read(); err != nil {
		return nil, err
	}

	if err := c.watch(); err != nil {
		return nil, err
	}
	return c, nil
}

type nacosConfig struct {
	conf   *SourceConfig
	opts   options
	client config_client.IConfigClient

	sync.RWMutex
	data ChangeSet
}

func (this *nacosConfig) Load(v any) error {
	this.RLock()
	defer this.RUnlock()
	return this.opts.Encoding.Unmarshal(this.data.Value, v)
}

var defaultGroup = "DEFAULT_GROUP"

func (this *nacosConfig) read() (ChangeSet, error) {
	value, err := this.client.GetConfig(vo.ConfigParam{
		DataId: this.conf.AppId,
		Group:  defaultGroup,
	})
	if err != nil {
		return ChangeSet{}, err
	}

	data := ChangeSet{Timestamp: time.Now(), Value: []byte(value)}

	this.Lock()
	this.data = data
	this.Unlock()
	return data, nil
}

func (this *nacosConfig) watch() error {
	if this.opts.WatchChange == nil {
		return nil
	}

	err := this.client.ListenConfig(vo.ConfigParam{
		DataId:   this.conf.AppId,
		Group:    defaultGroup,
		OnChange: this.callback,
	})

	return err
}

func (this *nacosConfig) callback(_, _, _, value string) {
	data := ChangeSet{
		Timestamp: time.Now(),
		Value:     []byte(value),
	}

	this.Lock()
	this.data = data
	this.Unlock()

	this.opts.WatchChange(data)
}

func (this *nacosConfig) init(conf *SourceConfig) (err error) {
	clientConfig := constant.NewClientConfig(
		constant.WithNamespaceId(conf.Cluster),
		constant.WithNotLoadCacheAtStart(true),
	)
	serverConfigs := make([]constant.ServerConfig, 0)
	for _, endpoint := range conf.Endpoints {
		host, port, err := net.SplitHostPort(endpoint)
		if err != nil {
			return err
		}

		p, err := strconv.ParseUint(port, 10, 64)
		if err != nil {
			return err
		}

		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: host,
			Port:   p,
		})
	}
	this.client, err = clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  clientConfig,
		ServerConfigs: serverConfigs,
	})
	return err
}
