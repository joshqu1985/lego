package configor

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/joshqu1985/lego/utils/routine"
	clientv3 "go.etcd.io/etcd/client/v3"
)

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

type etcdConfig struct {
	conf   *SourceConfig
	opts   options
	client *clientv3.Client
	prefix string

	sync.RWMutex
	data ChangeSet
}

func (this *etcdConfig) Load(v any) error {
	this.RLock()
	defer this.RUnlock()
	return this.opts.Encoding.Unmarshal(this.data.Value, v)
}

func (this *etcdConfig) watch() error {
	if this.opts.WatchChange == nil {
		return nil
	}

	ch := this.client.Watcher.Watch(context.Background(), this.prefix, clientv3.WithPrefix())
	routine.Go(func() { this.run(ch) })

	return nil
}

func (this *etcdConfig) read() (ChangeSet, error) {
	resp, err := this.client.Get(context.Background(), this.prefix, clientv3.WithPrefix())
	if err != nil {
		return ChangeSet{}, err
	}

	if resp == nil || len(resp.Kvs) == 0 {
		return ChangeSet{}, fmt.Errorf("data not found: %s", this.prefix)
	}

	var buffer bytes.Buffer
	for _, kv := range resp.Kvs {
		buffer.WriteString(fmt.Sprintf("%s\n", string(kv.Value)))
	}
	data := ChangeSet{Timestamp: time.Now(), Value: buffer.Bytes()}

	this.Lock()
	this.data = data
	this.Unlock()
	return data, nil
}

func (this *etcdConfig) run(ch clientv3.WatchChan) {
	for {
		select {
		case resp, ok := <-ch:
			if !ok {
				return
			}
			if len(resp.Events) == 0 {
				continue
			}
			data, err := this.read()
			if err != nil {
				glog.Errorf("etcd config read err:%v", err)
				continue
			}
			routine.Safe(func() {
				this.opts.WatchChange(data)
			})
		}
	}
}

func (this *etcdConfig) init(conf *SourceConfig) (err error) {
	config := clientv3.Config{
		Endpoints:   conf.Endpoints,
		Username:    conf.AccessKey,
		Password:    conf.SecretKey,
		DialTimeout: 3 * time.Second,
	}
	this.client, err = clientv3.New(config)
	return err
}
