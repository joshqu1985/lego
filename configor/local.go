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

type localConfig struct {
	opts    options
	watcher *fsnotify.Watcher
	file    string
	data    ChangeSet
	sync.RWMutex
}

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

func (lc *localConfig) Load(v any) error {
	lc.RLock()
	defer lc.RUnlock()

	return lc.opts.Encoding.Unmarshal(lc.data.Value, v)
}

func (lc *localConfig) read(file string) (ChangeSet, error) {
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

	lc.Lock()
	lc.data = data
	lc.Unlock()

	return data, nil
}

func (lc *localConfig) watch() error {
	if lc.opts.WatchChange == nil {
		return nil
	}

	var err error
	lc.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	if xerr := lc.watcher.Add(lc.file); xerr != nil {
		return xerr
	}

	routine.Go(func() { lc.run() })

	return nil
}

func (lc *localConfig) run() {
	for {
		select {
		case event, ok := <-lc.watcher.Events:
			if !ok {
				return
			}

			if event.Has(fsnotify.Write) {
				data, err := lc.read(event.Name)
				if err != nil {
					glog.Errorf("local config read err:%v", err)

					continue
				}
				lc.opts.WatchChange(data)
			}
		case err, ok := <-lc.watcher.Errors:
			if !ok {
				return
			}
			glog.Errorf("local config err:%v", err)
		}
	}
}
