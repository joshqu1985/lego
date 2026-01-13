package configor

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"

	clientv3 "go.etcd.io/etcd/client/v3" //nolint:gomodguard

	"github.com/joshqu1985/lego/utils/routine"
)

type etcdConfig struct {
	opts   options
	conf   *SourceConfig
	client *clientv3.Client
	prefix string
	data   ChangeSet
	sync.RWMutex
}

func NewEtcd(conf *SourceConfig, opts options) (Configor, error) {
	c := &etcdConfig{
		conf: conf,
		opts: opts,
	}

	// etcd配置前缀config:${conf.Cluster}:${conf.AppId}:
	prefix := strings.Join([]string{"config", conf.Cluster, conf.AppId}, ":")
	c.prefix = prefix + ":"

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

func (ec *etcdConfig) Load(v any) error {
	ec.RLock()
	defer ec.RUnlock()

	return ec.opts.Encoding.Unmarshal(ec.data.Value, v)
}

func (ec *etcdConfig) watch() error {
	if ec.opts.WatchChange == nil {
		return nil
	}

	ch := ec.client.Watch(context.Background(), ec.prefix, clientv3.WithPrefix())
	routine.Go(func() { ec.run(ch) })

	return nil
}

func (ec *etcdConfig) read() (ChangeSet, error) {
	resp, err := ec.client.Get(context.Background(), ec.prefix, clientv3.WithPrefix())
	if err != nil {
		return ChangeSet{}, err
	}

	if resp == nil || len(resp.Kvs) == 0 {
		return ChangeSet{}, fmt.Errorf("data not found: %s", ec.prefix)
	}

	var buffer bytes.Buffer
	for _, kv := range resp.Kvs {
		_, _ = buffer.WriteString(string(kv.Value) + "\n")
	}
	data := ChangeSet{Timestamp: time.Now(), Value: buffer.Bytes()}

	ec.Lock()
	ec.data = data
	ec.Unlock()

	return data, nil
}

func (ec *etcdConfig) run(ch clientv3.WatchChan) {
	for {
		select {
		case resp, ok := <-ch:
			if !ok {
				return
			}
			if len(resp.Events) == 0 {
				continue
			}
			data, err := ec.read()
			if err != nil {
				glog.Errorf("etcd config read err:%v", err)

				continue
			}
			_ = routine.Safe(func() {
				ec.opts.WatchChange(data)
			})
		}
	}
}

func (ec *etcdConfig) init(conf *SourceConfig) error {
	config := clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.AccessKey,
		Password:    conf.SecretKey,
		DialTimeout: 3 * time.Second,
	}

	var err error
	ec.client, err = clientv3.New(config)

	return err
}
