package rpc

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joshqu1985/lego/logs"
)

type Monitor struct {
	callbacks []func()
	lock      sync.Mutex
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

func (m *Monitor) AddShutdownCallback(fn func()) {
	m.lock.Lock()
	m.callbacks = append(m.callbacks, fn)
	m.lock.Unlock()
}

func (m *Monitor) listenSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	for {
		sig := <-signals
		switch sig {
		case syscall.SIGTERM:
			m.processCallbacks(signals, syscall.SIGTERM)
		case syscall.SIGINT:
			m.processCallbacks(signals, syscall.SIGINT)
		default:
			logs.Error("unregistered signal:", sig)
		}
	}
}

func (m *Monitor) processCallbacks(signals chan os.Signal, _ syscall.Signal) {
	signal.Stop(signals)

	m.lock.Lock()
	callbacks := m.callbacks
	m.lock.Unlock()

	for _, callback := range callbacks {
		callback()
	}
}
