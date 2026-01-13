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

type nacosConfig struct {
	opts   options
	client config_client.IConfigClient
	conf   *SourceConfig
	data   ChangeSet
	sync.RWMutex
}

var defaultGroup = "DEFAULT_GROUP"

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

func (nc *nacosConfig) Load(v any) error {
	nc.RLock()
	defer nc.RUnlock()

	return nc.opts.Encoding.Unmarshal(nc.data.Value, v)
}

func (nc *nacosConfig) read() (ChangeSet, error) {
	value, err := nc.client.GetConfig(vo.ConfigParam{
		DataId: nc.conf.AppId,
		Group:  defaultGroup,
	})
	if err != nil {
		return ChangeSet{}, err
	}

	data := ChangeSet{Timestamp: time.Now(), Value: []byte(value)}

	nc.Lock()
	nc.data = data
	nc.Unlock()

	return data, nil
}

func (nc *nacosConfig) watch() error {
	if nc.opts.WatchChange == nil {
		return nil
	}

	return nc.client.ListenConfig(vo.ConfigParam{
		DataId:   nc.conf.AppId,
		Group:    defaultGroup,
		OnChange: nc.callback,
	})
}

func (nc *nacosConfig) callback(_, _, _, value string) {
	data := ChangeSet{
		Timestamp: time.Now(),
		Value:     []byte(value),
	}

	nc.Lock()
	nc.data = data
	nc.Unlock()

	nc.opts.WatchChange(data)
}

func (nc *nacosConfig) init(conf *SourceConfig) error {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(conf.Cluster),
		constant.WithNotLoadCacheAtStart(true),
	)
	serverConfigs := make([]constant.ServerConfig, 0)
	for _, endpoint := range conf.Endpoints {
		host, port, err := net.SplitHostPort(endpoint)
		if err != nil {
			return err
		}
		iport, err := strconv.ParseUint(port, 10, 64)
		if err != nil {
			return err
		}

		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: host,
			Port:   iport,
		})
	}

	var err error
	nc.client, err = clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})

	return err
}
