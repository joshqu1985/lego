package rpc

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/golang/glog"
)

type Monitor struct {
	callbacks []func()
	sync.Mutex
}

func NewMonitor() *Monitor {
	m := &Monitor{
		callbacks: make([]func(), 0),
	}
	go func() {
		m.listenSignals()
	}()

	return m
}

func (this *Monitor) AddShutdownCallback(fn func()) {
	this.Lock()
	this.callbacks = append(this.callbacks, fn)
	this.Unlock()
}

func (this *Monitor) listenSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	for {
		sig := <-signals
		switch sig {
		case syscall.SIGTERM:
			this.processCallbacks(signals, syscall.SIGTERM)
		case syscall.SIGINT:
			this.processCallbacks(signals, syscall.SIGINT)
		default:
			glog.Error("unregistered signal:", sig)
		}
	}
}

func (this *Monitor) processCallbacks(signals chan os.Signal, _ syscall.Signal) {
	signal.Stop(signals)

	this.Lock()
	callbacks := this.callbacks
	this.Unlock()

	for _, callback := range callbacks {
		callback()
	}
}
