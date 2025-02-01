package naming

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcd(conf *Config) (Naming, error) {
	c := &etcd{
		services: make(map[string]*etcdService),
		quit:     make(chan struct{}),
		config:   conf,
	}

	// etcd前缀 discovery:${conf.Cluster}:
	c.prefix = strings.Join([]string{"naming", conf.Cluster}, ":")
	return c, c.init(conf)
}

type etcd struct {
	prefix string
	key    string
	val    string
	lease  clientv3.LeaseID
	client *clientv3.Client
	quit   chan struct{}
	config *Config

	services map[string]*etcdService
	sync.RWMutex
}

func (this *etcd) Endpoints() []string {
	return this.config.Endpoints
}

func (this *etcd) Name() string {
	return "etcd"
}

func (this *etcd) Register(key, val string) error {
	this.key, this.val = key, val

	if this.client == nil {
		return fmt.Errorf("etcd client is nil")
	}

	if err := this.register(); err != nil {
		return err
	}

	return this.keepAlive()
}

func (this *etcd) Deregister(_ string) error {
	key := fmt.Sprintf("%s:%s:%d", this.prefix, this.key, this.lease)
	_, err := this.client.Delete(context.Background(), key)
	return err
}

func (this *etcd) Service(key string) RegService {
	this.Lock()
	defer this.Unlock()

	service, ok := this.services[key]
	if ok {
		return service
	}

	service = &etcdService{
		prefix: this.prefix,
		key:    key,
		client: this.client,
	}
	go func() { _ = service.Watch() }()

	this.services[key] = service
	return service
}

func (this *etcd) Close() {
	if this.client != nil {
		this.client.Close()
	}
	close(this.quit)
}

func (this *etcd) init(conf *Config) (err error) {
	config := clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.AccessKey,
		Password:    conf.SecretKey,
		DialTimeout: 3 * time.Second,
	}

	this.client, err = clientv3.New(config)
	return err
}

func (this *etcd) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()

	resp, err := this.client.Grant(ctx, 10)
	if err != nil {
		return err
	}
	this.lease = resp.ID

	key := fmt.Sprintf("%s:%s:%d", this.prefix, this.key, this.lease)
	_, err = this.client.Put(ctx, key, this.val,
		clientv3.WithLease(this.lease))
	return err
}

func (this *etcd) keepAlive() error {
	ch, err := this.client.KeepAlive(context.Background(), this.lease)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					_ = this.revoke()
					return
				}
			case <-this.quit:
				_ = this.revoke()
				return
			}
		}
	}()
	return nil
}

func (this *etcd) revoke() error {
	_, err := this.client.Revoke(context.Background(), this.lease)
	return err
}

type etcdService struct {
	prefix string
	key    string
	client *clientv3.Client

	revision int64
	values   map[string]string
	sync.RWMutex
	listeners []func()
}

func (this *etcdService) Name() string {
	return this.key
}

func (this *etcdService) Addrs() ([]string, error) {
	if this.client == nil {
		return nil, fmt.Errorf("etcd client is nil")
	}

	this.RLock()
	revision, addrs := this.revision, this.values
	this.RUnlock()

	if revision != 0 && len(addrs) != 0 {
		return lo.Values(addrs), nil
	}

	addrs, err := this.load()
	if err != nil {
		return nil, err
	}
	return lo.Values(addrs), nil
}

func (this *etcdService) AddListener(f func()) {
	this.Lock()
	this.listeners = append(this.listeners, f)
	this.Unlock()
}

func (this *etcdService) load() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()

	key := fmt.Sprintf("%s:%s:", this.prefix, this.key)
	resp, err := this.client.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	addrs := map[string]string{}
	for _, kv := range resp.Kvs {
		addrs[string(kv.Key)] = string(kv.Value)
	}

	this.Lock()
	this.revision, this.values = resp.Header.Revision, addrs
	this.Unlock()
	return addrs, nil
}

func (this *etcdService) Watch() error {
	key := fmt.Sprintf("%s:%s:", this.prefix, this.key)
	watchCh := this.client.Watcher.Watch(context.Background(), key,
		clientv3.WithPrefix(), clientv3.WithRev(this.revision+1))

	for {
		select {
		case resp, ok := <-watchCh:
			if !ok {
				return fmt.Errorf("etcd watch chan has been closed")
			}
			if resp.Canceled || resp.Err() != nil {
				return resp.Err()
			}
			this.processEvents(resp.Events)
		}
	}
}

func (this *etcdService) processEvents(events []*clientv3.Event) {
	this.Lock()
	listeners := this.listeners
	this.Unlock()

	for _, ev := range events {
		switch ev.Type {
		case clientv3.EventTypePut:
			this.Lock()
			this.values[string(ev.Kv.Key)] = string(ev.Kv.Value)
			this.Unlock()
		case clientv3.EventTypeDelete:
			this.Lock()
			delete(this.values, string(ev.Kv.Key))
			this.Unlock()
		default:
			glog.Errorf("unknown etcd watch event type:%v", ev.Type)
		}
	}

	for _, listener := range listeners {
		listener()
	}
}
