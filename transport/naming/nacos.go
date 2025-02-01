package naming

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/golang/glog"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/samber/lo"
)

func NewNacos(conf *Config) (Naming, error) {
	c := &nacos{
		services: make(map[string]*nacosService),
		config:   conf,
	}

	return c, c.init(conf)
}

type nacos struct {
	key    string
	val    string
	client naming_client.INamingClient
	config *Config

	services map[string]*nacosService
	sync.RWMutex
}

func (this *nacos) Endpoints() []string {
	return this.config.Endpoints
}

func (this *nacos) Name() string {
	return "nacos"
}

func (this *nacos) Register(key, val string) error {
	this.key, this.val = key, val

	if this.client == nil {
		return fmt.Errorf("etcd client is nil")
	}

	return this.register()
}

func (this *nacos) Deregister(_ string) error {
	host, port, err := net.SplitHostPort(this.val)
	if err != nil {
		return err
	}
	iport, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return err
	}

	request := vo.DeregisterInstanceParam{
		ServiceName: this.key,
		Ip:          host,
		Port:        iport,
		Ephemeral:   true,
	}
	_, err = this.client.DeregisterInstance(request)
	return err
}

func (this *nacos) Service(key string) RegService {
	this.Lock()
	defer this.Unlock()

	service, ok := this.services[key]
	if ok {
		return service
	}

	service = &nacosService{
		key:    key,
		client: this.client,
	}
	go func() { _ = service.Watch() }()

	this.services[key] = service

	return service
}

func (this *nacos) Close() {
	if this.client != nil {
		this.client.CloseClient()
	}
}

func (this *nacos) init(conf *Config) (err error) {
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
	this.client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	return err
}

func (this *nacos) register() error {
	host, port, err := net.SplitHostPort(this.val)
	if err != nil {
		return err
	}
	iport, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return err
	}

	request := vo.RegisterInstanceParam{
		ServiceName: this.key,
		Ip:          host,
		Port:        iport,
		Enable:      true,
		Healthy:     true,
		Weight:      1.0,
		Ephemeral:   true,
	}
	success, err := this.client.RegisterInstance(request)
	if err != nil || !success {
		glog.Errorf("register nacos failed %v name:%s ip:%s port:%d", err, this.key, host, port)
	}
	return err
}

type nacosService struct {
	key    string
	client naming_client.INamingClient

	values map[string]string
	sync.RWMutex
	listeners []func()
}

func (this *nacosService) Name() string {
	return this.key
}

func (this *nacosService) Addrs() ([]string, error) {
	if this.client == nil {
		return nil, fmt.Errorf("nacos client is nil")
	}

	this.RLock()
	addrs := this.values
	this.RUnlock()

	if len(addrs) != 0 {
		return lo.Values(addrs), nil
	}

	addrs, err := this.load()
	if err != nil {
		return nil, err
	}

	return lo.Values(addrs), nil
}

func (this *nacosService) AddListener(f func()) {
	this.Lock()
	this.listeners = append(this.listeners, f)
	this.Unlock()
}

func (this *nacosService) load() (map[string]string, error) {
	service, err := this.client.GetService(vo.GetServiceParam{
		ServiceName: this.key,
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

	this.Lock()
	this.values = addrs
	this.Unlock()
	return addrs, nil
}

func (this *nacosService) Watch() error {
	param := &vo.SubscribeParam{
		ServiceName:       this.key,
		SubscribeCallback: this.processEvents,
	}
	return this.client.Subscribe(param)
}

func (this *nacosService) processEvents(items []model.Instance, err error) {
	if err != nil {
		glog.Errorf("nacos subscribe callback err:%v", err)
		return
	}

	addrs := make(map[string]string)
	for _, item := range items {
		if !item.Enable {
			continue
		}
		addrs[item.InstanceId] = fmt.Sprintf("%s:%d", item.Ip, item.Port)
	}
	this.Lock()
	this.values = addrs
	listeners := this.listeners
	this.Unlock()

	for _, listener := range listeners {
		listener()
	}
}
