package container

import (
	"sync"
	"time"

	"github.com/joshqu1985/lego/utils/routine"
)

type (
	Batcher struct {
		queue     chan []any
		ticker    *time.Ticker
		done      chan struct{}
		entries   []any
		interval  time.Duration
		bulkSize  int64
		queueSize int64
		sync.RWMutex
	}

	BatcherOption func(o *Batcher)
)

func NewBatcher(opts ...BatcherOption) *Batcher {
	var batcher Batcher
	for _, opt := range opts {
		opt(&batcher)
	}
	if batcher.interval == 0 {
		batcher.interval = time.Second
	}
	if batcher.bulkSize == 0 {
		batcher.bulkSize = 500
	}
	if batcher.queueSize == 0 {
		batcher.queueSize = 100
	}
	batcher.entries = make([]any, 0, batcher.bulkSize)
	batcher.queue = make(chan []any, batcher.queueSize)

	routine.Go(func() { batcher.intervalFlush() })

	return &batcher
}

func (b *Batcher) Put(data any) {
	b.Lock()
	defer b.Unlock()

	b.entries = append(b.entries, data)
	if len(b.entries) >= int(b.bulkSize) {
		b.flush()
	}
}

func (b *Batcher) Queue() chan []any {
	return b.queue
}

func (b *Batcher) Close() error {
	close(b.done)
	b.ticker.Stop()

	b.Lock()
	defer b.Unlock()

	if len(b.entries) > 0 {
		b.flush()
	}
	return nil
}

func (b *Batcher) intervalFlush() {
	b.ticker = time.NewTicker(b.interval)
	defer b.ticker.Stop()

	for {
		select {
		case <-b.ticker.C:
			b.Lock()
			b.flush()
			b.Unlock()
		case <-b.done:
			return
		}
	}
}

func (b *Batcher) flush() {
	if len(b.entries) == 0 {
		return
	}
	b.queue <- b.entries
	b.entries = make([]any, 0, b.bulkSize)
}

// WithInterval 设置批量的最大等待时间.
func WithInterval(interval time.Duration) BatcherOption {
	return func(b *Batcher) {
		b.interval = interval
	}
}

// WithBulkSize 设置批量的最大条数.
func WithBulkSize(bulkSize int64) BatcherOption {
	return func(b *Batcher) {
		b.bulkSize = bulkSize
	}
}

// WithQueueSize 设置队列的最大长度.
func WithQueueSize(queueSize int64) BatcherOption {
	return func(b *Batcher) {
		b.queueSize = queueSize
	}
}
