package configor

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"github.com/joshqu1985/lego/utils/routine"
)

func NewLocal(file string, opts options) (Configor, error) {
	c := &localConfig{
		file: file,
		opts: opts,
	}

	if _, err := c.read(file); err != nil {
		return nil, err
	}

	if err := c.watch(); err != nil {
		return nil, err
	}
	return c, nil
}

type localConfig struct {
	file    string
	opts    options
	watcher *fsnotify.Watcher

	sync.RWMutex
	data ChangeSet
}

func (this *localConfig) Load(v any) error {
	this.RLock()
	defer this.RUnlock()
	return this.opts.Encoding.Unmarshal(this.data.Value, v)
}

func (this *localConfig) read(file string) (ChangeSet, error) {
	fp, err := os.Open(file)
	if err != nil {
		return ChangeSet{}, err
	}
	defer fp.Close()

	value, err := io.ReadAll(fp)
	if err != nil {
		return ChangeSet{}, err
	}

	data := ChangeSet{Timestamp: time.Now(), Value: value}

	this.Lock()
	this.data = data
	this.Unlock()
	return data, nil
}

func (this *localConfig) watch() error {
	if this.opts.WatchChange == nil {
		return nil
	}

	var err error
	this.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	if err := this.watcher.Add(this.file); err != nil {
		return err
	}

	routine.Go(func() { this.run() })
	return nil
}

func (this *localConfig) run() {
	for {
		select {
		case event, ok := <-this.watcher.Events:
			if !ok {
				return
			}

			if event.Has(fsnotify.Write) {
				data, err := this.read(event.Name)
				if err != nil {
					glog.Errorf("local config read err:%v", err)
					continue
				}
				this.opts.WatchChange(data)
			}
		case err, ok := <-this.watcher.Errors:
			if !ok {
				return
			}
			glog.Errorf("local config err:%v", err)
		}
	}
}
