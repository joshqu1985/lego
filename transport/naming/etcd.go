package naming

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"

	clientv3 "go.etcd.io/etcd/client/v3" //nolint:gomodguard

	"github.com/joshqu1985/lego/logs"
)

type (
	etcd struct {
		client   *clientv3.Client
		quit     chan struct{}
		config   *Config
		services map[string]*etcdService
		prefix   string
		key      string
		val      string
		lease    clientv3.LeaseID
		sync.RWMutex
	}

	etcdService struct {
		client    *clientv3.Client
		values    map[string]string
		prefix    string
		key       string
		listeners []func()
		revision  int64
		sync.RWMutex
	}
)

func NewEtcd(conf *Config) (Naming, error) {
	c := &etcd{
		services: make(map[string]*etcdService),
		quit:     make(chan struct{}),
		config:   conf,
	}

	// etcd前缀 naming:${conf.Cluster}:
	c.prefix = "naming" + ":" + conf.Cluster

	err := c.init(conf)

	return c, err
}

func (e *etcd) Endpoints() []string {
	return e.config.Endpoints
}

func (e *etcd) Name() string {
	return SOURCE_ETCD
}

func (e *etcd) Register(key, val string) error {
	e.key, e.val = key, val

	if e.client == nil {
		return ErrClientNil
	}

	if err := e.register(); err != nil {
		return err
	}

	return e.keepAlive()
}

func (e *etcd) Deregister(_ string) error {
	key := fmt.Sprintf("%s:%s:%d", e.prefix, e.key, e.lease)
	_, err := e.client.Delete(context.Background(), key)

	return err
}

func (e *etcd) Service(key string) RegService {
	e.Lock()
	defer e.Unlock()

	service, ok := e.services[key]
	if ok {
		return service
	}

	service = &etcdService{
		prefix: e.prefix,
		key:    key,
		client: e.client,
	}
	go func() { _ = service.Watch() }()

	e.services[key] = service

	return service
}

func (e *etcd) Close() {
	if e.client != nil {
		e.client.Close()
	}
	close(e.quit)
}

func (e *etcd) init(conf *Config) error {
	config := clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.AccessKey,
		Password:    conf.SecretKey,
		DialTimeout: 3 * time.Second,
	}

	var err error
	e.client, err = clientv3.New(config)

	return err
}

func (e *etcd) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()

	resp, err := e.client.Grant(ctx, 10)
	if err != nil {
		return err
	}
	e.lease = resp.ID

	key := fmt.Sprintf("%s:%s:%d", e.prefix, e.key, e.lease)
	_, err = e.client.Put(ctx, key, e.val,
		clientv3.WithLease(e.lease))

	return err
}

func (e *etcd) keepAlive() error {
	ch, err := e.client.KeepAlive(context.Background(), e.lease)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					_ = e.revoke()

					return
				}
			case <-e.quit:
				_ = e.revoke()

				return
			}
		}
	}()

	return nil
}

func (e *etcd) revoke() error {
	_, err := e.client.Revoke(context.Background(), e.lease)

	return err
}

// ----------- etcdService -------------

func (es *etcdService) Name() string {
	return es.key
}

func (es *etcdService) Addrs() ([]string, error) {
	if es.client == nil {
		return nil, errors.New("etcd client is nil")
	}

	es.RLock()
	revision, addrs := es.revision, es.values
	es.RUnlock()

	if revision != 0 && len(addrs) != 0 {
		return lo.Values(addrs), nil
	}

	addrs, err := es.load()
	if err != nil {
		return nil, err
	}

	return lo.Values(addrs), nil
}

func (es *etcdService) AddListener(f func()) {
	es.Lock()
	es.listeners = append(es.listeners, f)
	es.Unlock()
}

func (es *etcdService) load() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()

	key := fmt.Sprintf("%s:%s:", es.prefix, es.key)
	resp, err := es.client.Get(ctx, key, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		return nil, err
	}

	addrs := make(map[string]string)
	for _, kv := range resp.Kvs {
		addrs[string(kv.Key)] = string(kv.Value)
	}

	es.Lock()
	es.revision, es.values = resp.Header.Revision, addrs
	es.Unlock()

	return addrs, nil
}

func (es *etcdService) Watch() error {
	key := fmt.Sprintf("%s:%s:", es.prefix, es.key)
	watchCh := es.client.Watch(context.Background(), key,
		clientv3.WithPrefix(), clientv3.WithRev(es.revision+1))

	for {
		select {
		case resp, ok := <-watchCh:
			if !ok {
				return errors.New("etcd watch chan has been closed")
			}
			if resp.Canceled || resp.Err() != nil {
				return resp.Err()
			}
			es.processEvents(resp.Events)
		}
	}
}

func (es *etcdService) processEvents(events []*clientv3.Event) {
	es.Lock()
	listeners := es.listeners
	es.Unlock()

	for _, ev := range events {
		switch ev.Type {
		case clientv3.EventTypePut:
			es.Lock()
			es.values[string(ev.Kv.Key)] = string(ev.Kv.Value)
			es.Unlock()
		case clientv3.EventTypeDelete:
			es.Lock()
			delete(es.values, string(ev.Kv.Key))
			es.Unlock()
		default:
			logs.Errorf("unknown etcd watch event type:%v", ev.Type)
		}
	}

	for _, listener := range listeners {
		listener()
	}
}
