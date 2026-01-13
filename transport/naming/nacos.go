package naming

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/samber/lo"

	"github.com/joshqu1985/lego/logs"
)

type (
	nacos struct {
		client   naming_client.INamingClient
		config   *Config
		services map[string]*nacosService
		key      string
		val      string
		sync.RWMutex
	}

	nacosService struct {
		client    naming_client.INamingClient
		values    map[string]string
		key       string
		listeners []func()
		sync.RWMutex
	}
)

func NewNacos(conf *Config) (Naming, error) {
	c := &nacos{
		services: make(map[string]*nacosService),
		config:   conf,
	}

	err := c.init(conf)

	return c, err
}

func (n *nacos) Endpoints() []string {
	return n.config.Endpoints
}

func (n *nacos) Name() string {
	return SOURCE_NACOS
}

func (n *nacos) Register(key, val string) error {
	n.key, n.val = key, val

	if n.client == nil {
		return ErrClientNil
	}

	return n.register()
}

func (n *nacos) Deregister(_ string) error {
	host, port, err := net.SplitHostPort(n.val)
	if err != nil {
		return err
	}
	iport, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return err
	}

	request := vo.DeregisterInstanceParam{
		ServiceName: n.key,
		Ip:          host,
		Port:        iport,
		Ephemeral:   true,
	}
	_, err = n.client.DeregisterInstance(request)

	return err
}

func (n *nacos) Service(key string) RegService {
	n.Lock()

	service, ok := n.services[key]
	if ok {
		n.Unlock()

		return service
	}

	service = &nacosService{
		key:    key,
		client: n.client,
	}
	n.services[key] = service
	n.Unlock()

	go func(s *nacosService) {
		if err := service.Watch(); err != nil {
			logs.Errorf("nacos watch err:%v", err)
		}
	}(service)

	return service
}

func (n *nacos) Close() {
	if n.client != nil {
		n.client.CloseClient()
	}
}

func (n *nacos) init(conf *Config) error {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(conf.Cluster),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
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
	n.client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	return err
}

func (n *nacos) register() error {
	host, port, err := net.SplitHostPort(n.val)
	if err != nil {
		return err
	}
	iport, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return err
	}

	request := vo.RegisterInstanceParam{
		ServiceName: n.key,
		Ip:          host,
		Port:        iport,
		Enable:      true,
		Healthy:     true,
		Weight:      1.0,
		Ephemeral:   true,
	}
	success, err := n.client.RegisterInstance(request)
	if err != nil || !success {
		logs.Errorf("register nacos failed %v name:%s ip:%s port:%d", err, n.key, host, port)
	}

	return err
}

// ----------------- nacosService -----------------

func (ns *nacosService) Name() string {
	return ns.key
}

func (ns *nacosService) Addrs() ([]string, error) {
	if ns.client == nil {
		return nil, errors.New("nacos client is nil")
	}

	ns.RLock()
	addrs := ns.values
	ns.RUnlock()

	if len(addrs) != 0 {
		return lo.Values(addrs), nil
	}

	addrs, err := ns.load()
	if err != nil {
		return nil, err
	}

	return lo.Values(addrs), nil
}

func (ns *nacosService) AddListener(f func()) {
	ns.Lock()
	ns.listeners = append(ns.listeners, f)
	ns.Unlock()
}

func (ns *nacosService) load() (map[string]string, error) {
	service, err := ns.client.GetService(vo.GetServiceParam{
		ServiceName: ns.key,
	})
	if err != nil {
		return nil, err
	}

	addrs := make(map[string]string)
	for _, item := range service.Hosts {
		if !item.Enable || !item.Healthy || item.Weight <= 0 {
			continue
		}
		addrs[item.InstanceId] = fmt.Sprintf("%s:%d", item.Ip, item.Port)
	}

	ns.Lock()
	ns.values = addrs
	ns.Unlock()

	return addrs, nil
}

func (ns *nacosService) Watch() error {
	param := &vo.SubscribeParam{
		ServiceName:       ns.key,
		SubscribeCallback: ns.processEvents,
	}

	return ns.client.Subscribe(param)
}

func (ns *nacosService) processEvents(items []model.Instance, err error) {
	if err != nil {
		logs.Errorf("nacos subscribe callback err:%v", err)

		return
	}

	addrs := make(map[string]string)
	for _, item := range items {
		if !item.Enable {
			continue
		}
		addrs[item.InstanceId] = fmt.Sprintf("%s:%d", item.Ip, item.Port)
	}
	ns.Lock()
	ns.values = addrs
	listeners := ns.listeners
	ns.Unlock()

	for _, listener := range listeners {
		listener()
	}
}
