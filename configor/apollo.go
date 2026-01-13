package configor

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/golang/glog"
)

type (
	apolloConfig struct {
		opts   options
		client agollo.Client
		conf   *SourceConfig
		data   ChangeSet
		sync.RWMutex
	}

	CustomChangeListener struct {
		client       agollo.Client
		apolloConfig *apolloConfig
		notify       ChangeNotify
	}
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

func (ac *apolloConfig) Load(v any) error {
	ac.RLock()
	defer ac.RUnlock()

	return ac.opts.Encoding.Unmarshal(ac.data.Value, v)
}

func (ac *apolloConfig) init(conf *SourceConfig) error {
	if len(conf.Endpoints) == 0 {
		return errors.New("endpoints is empty")
	}

	meta := &config.AppConfig{
		IP:             conf.Endpoints[0], //
		Cluster:        conf.Cluster,      // ex: default / dev / test
		AppID:          conf.AppId,        //
		Secret:         conf.SecretKey,    //
		NamespaceName:  conf.Namespace,    // ex: application
		IsBackupConfig: true,
	}

	var err error
	ac.client, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return meta, nil
	})

	return err
}

func (ac *apolloConfig) read() (ChangeSet, error) {
	cache := ac.client.GetConfigCache(ac.conf.Namespace)
	if cache == nil {
		return ChangeSet{}, errors.New("apollo config cache is nil")
	}

	var buffer bytes.Buffer
	cache.Range(func(key, value any) bool {
		_, _ = buffer.WriteString(fmt.Sprintf("%s\n", value))

		return true
	})
	data := ChangeSet{Timestamp: time.Now(), Value: buffer.Bytes()}

	ac.Lock()
	ac.data = data
	ac.Unlock()

	return data, nil
}

func (ac *apolloConfig) watch() error {
	if ac.opts.WatchChange == nil {
		return nil
	}

	listener := &CustomChangeListener{
		client:       ac.client,
		apolloConfig: ac,
		notify:       ac.opts.WatchChange,
	}
	ac.client.AddChangeListener(listener)

	return nil
}

func (c *CustomChangeListener) OnChange(_ *storage.ChangeEvent) {
	val, err := c.apolloConfig.read()
	if err != nil {
		glog.Errorf("apollo config read err:%v", err)

		return
	}

	c.apolloConfig.Lock()
	c.apolloConfig.data = val
	data := c.apolloConfig.data
	c.apolloConfig.Unlock()

	if c.notify != nil {
		c.notify(data)
	}
}

func (c *CustomChangeListener) OnNewestChange(_ *storage.FullChangeEvent) {
}
