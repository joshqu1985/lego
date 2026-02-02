package broker

import (
	"context"
	"sync"

	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/utils/routine"
)

type memoryConsumer struct {
	consumers map[string]ConsumeCallback
	broker    *memoryBroker
	wg        *sync.WaitGroup
	stopWork  chan struct{}
	groupId   string
	lock      sync.RWMutex
}

func NewMemoryConsumer(conf *Config) (Consumer, error) {
	return &memoryConsumer{
		consumers: make(map[string]ConsumeCallback),
		groupId:   conf.GroupId,
		broker:    getMemoryBroker(conf.Topics),
		wg:        &sync.WaitGroup{},
		stopWork:  make(chan struct{}, len(conf.Topics)),
	}, nil
}

func (mc *memoryConsumer) Register(topic string, f ConsumeCallback) error {
	mc.lock.Lock()
	mc.consumers[topic] = f
	mc.lock.Unlock()

	mc.broker.addSubscriber(topic, mc.groupId)

	return nil
}

func (mc *memoryConsumer) Start() error {
	mc.lock.Lock()
	if len(mc.consumers) == 0 {
		mc.lock.Unlock()
		return ErrSubscriberNil
	}

	consumersCopy := make(map[string]ConsumeCallback)
	for topic, f := range mc.consumers {
		consumersCopy[topic] = f
	}
	mc.lock.Unlock()

	for topic, f := range consumersCopy {
		topicCopy, callbackCopy := topic, f
		mc.wg.Add(1)
		routine.Go(func() { mc.work(topicCopy, callbackCopy) })
	}

	return nil
}

func (mc *memoryConsumer) Close() error {
	mc.lock.Lock()
	callbacks := make(map[string]ConsumeCallback)
	for topic, f := range mc.consumers {
		callbacks[topic] = f
	}
	mc.lock.Unlock()

	total := len(callbacks)
	for i := 0; i < total; i++ {
		mc.stopWork <- struct{}{}
	}
	mc.wg.Wait()

	for topic := range callbacks {
		topicCopy := topic
		mc.broker.delSubscriber(topicCopy, mc.groupId)
	}

	return nil
}

func (mc *memoryConsumer) work(topic string, f ConsumeCallback) {
	defer mc.wg.Done()

	recv, err := mc.broker.getSubChannel(topic, mc.groupId)
	if err != nil {
		return
	}

	for {
		select {
		case msg := <-recv:
			serr := routine.Safe(func() {
				if xerr := f(context.Background(), &msg); xerr != nil {
					logs.Errorf("memory queue process err:%v", xerr)
				}
			})
			if serr != nil {
				logs.Errorf("memory queue process err:%v", serr)
			}
		case <-mc.stopWork:
			return
		}
	}
}
