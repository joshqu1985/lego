package configor

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

func NewApollo(conf *SourceConfig, opts options) (Configor, error) {
	c := &apolloConfig{
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

type apolloConfig struct {
	opts   options
	conf   *SourceConfig
	client agollo.Client

	sync.RWMutex
	data ChangeSet
}

func (this *apolloConfig) Load(v any) error {
	this.RLock()
	defer this.RUnlock()
	return this.opts.Encoding.Unmarshal(this.data.Value, v)
}

func (this *apolloConfig) init(conf *SourceConfig) (err error) {
	if len(conf.Endpoints) == 0 {
		return fmt.Errorf("endpoints is empty")
	}

	meta := &config.AppConfig{
		IP:             conf.Endpoints[0], //
		Cluster:        conf.Cluster,      // ex: default / dev / test
		AppID:          conf.AppId,        //
		Secret:         conf.SecretKey,    //
		NamespaceName:  conf.Namespace,    // ex: application
		IsBackupConfig: true,
	}
	this.client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return meta, nil
	})
	return err
}

func (this *apolloConfig) read() (ChangeSet, error) {
	cache := this.client.GetConfigCache(this.conf.Namespace)
	if cache == nil {
		return ChangeSet{}, fmt.Errorf("apollo config cache is nil")
	}

	var buffer bytes.Buffer
	cache.Range(func(key, value interface{}) bool {
		buffer.WriteString(fmt.Sprintf("%s\n", value))
		return true
	})
	data := ChangeSet{Timestamp: time.Now(), Value: buffer.Bytes()}

	this.Lock()
	this.data = data
	this.Unlock()
	return data, nil
}

func (this *apolloConfig) watch() error {
	if this.opts.WatchChange == nil {
		return nil
	}

	listener := &CustomChangeListener{
		client:       this.client,
		apolloConfig: this,
		notify:       this.opts.WatchChange,
	}
	this.client.AddChangeListener(listener)
	return nil
}

type CustomChangeListener struct {
	client       agollo.Client
	apolloConfig *apolloConfig
	notify       ChangeNotify
}

func (c *CustomChangeListener) OnChange(_ *storage.ChangeEvent) {
	data, err := c.apolloConfig.read()
	if err != nil {
		log.Println("apollo config read err:", err)
		return
	}
	c.notify(data)
}

func (c *CustomChangeListener) OnNewestChange(_ *storage.FullChangeEvent) {
}
