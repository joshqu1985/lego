package broker

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/rs/xid"

	"github.com/joshqu1985/lego/container"
	"github.com/joshqu1985/lego/utils/routine"
	"github.com/joshqu1985/lego/utils/utime"
)

type (
	memoryProducer struct {
		broker *memoryBroker
	}

	memoryBroker struct {
		sublist    map[string][]Subscriber
		stopDelay  chan struct{}
		stopWorker chan struct{}
		queues     sync.Map
		sync.RWMutex
	}

	Subscriber struct {
		output  chan Message
		groupId string
	}

	TopicQueue struct {
		linked   *container.Linked
		priority *container.Priority
	}
)

var (
	broker *memoryBroker
	once   sync.Once
	lock   sync.RWMutex
)

func NewMemoryProducer(conf *Config) (Producer, error) {
	mbroker := getMemoryBroker(conf.Topics)
	if mbroker == nil {
		return nil, errors.New("memory broker not initialized")
	}

	return &memoryProducer{broker: mbroker}, nil
}

func (mp *memoryProducer) Send(ctx context.Context, topic string, msg *Message) error {
	return mp.SendDelay(ctx, topic, msg, 0)
}

func (mp *memoryProducer) SendDelay(ctx context.Context, topic string, msg *Message, stamp int64) error {
	if msg == nil {
		return ErrMessageIsNil
	}
	msg.SetTopic(topic)
	msg.SetMsgId(xid.New().String())

	data, ok := mp.broker.queues.Load(topic)
	if !ok {
		return ErrTopicNotFound
	}

	queue, _ := data.(*TopicQueue)
	if int64(queue.priority.Size()+queue.linked.Size()) > MaxMessageCount {
		return ErrQueueIsFull
	}

	if stamp > time.Now().Unix() {
		queue.priority.Put(msg, stamp)
	} else {
		queue.linked.Put(msg)
	}

	return nil
}

func (mp *memoryProducer) Close() error {
	if mp.broker == nil {
		return nil
	}

	return mp.broker.Close()
}

func getMemoryBroker(topics map[string]string) *memoryBroker {
	once.Do(func() {
		keys := make([]string, 0)
		for key := range topics {
			keys = append(keys, key)
		}
		lock.RLock()
		broker = newMemoryBroker(keys)
		lock.RUnlock()
	})

	lock.RLock()
	defer lock.RUnlock()

	return broker
}

func newMemoryBroker(topics []string) *memoryBroker {
	mbroker := &memoryBroker{
		sublist:    make(map[string][]Subscriber),
		stopDelay:  make(chan struct{}, 1),
		stopWorker: make(chan struct{}, 1),
	}
	for _, topic := range topics {
		mbroker.queues.LoadOrStore(topic, &TopicQueue{
			linked:   container.NewLinked(),
			priority: container.NewPriority(),
		})
	}
	mbroker.run()

	return mbroker
}

func (mp *memoryBroker) Close() error {
	mp.stopDelay <- struct{}{}
	mp.stopWorker <- struct{}{}

	return nil
}

func (mp *memoryBroker) run() {
	queuesCopy := make(map[string]*TopicQueue)
	mp.queues.Range(func(key, val any) bool {
		topic, _ := key.(string)
		queue, _ := val.(*TopicQueue)
		queuesCopy[topic] = queue

		return true
	})

	for topic, queue := range queuesCopy {
		topicCopy, queueCopy := topic, queue
		routine.Go(func() { mp.fetchWorker(topicCopy, queueCopy) })
		routine.Go(func() { mp.delayWorker(queueCopy) })
	}
}

func (mp *memoryBroker) fetchWorker(topic string, queue *TopicQueue) {
	attempt := 0
	for {
		select {
		case <-mp.stopWorker:
			return
		default:
			utime.Sleep(backoff(attempt, 100*time.Millisecond, 1*time.Second))
			sublist, err := mp.getSubscribers(topic)
			if err != nil || len(sublist) == 0 {
				attempt++

				continue
			}

			data, err := queue.linked.Get()
			if err != nil {
				attempt++

				continue
			}
			msg, _ := data.(*Message)

			for _, sub := range sublist {
				sub.output <- *msg
			}
			attempt = 0
		}
	}
}

func (mp *memoryBroker) delayWorker(queue *TopicQueue) {
	attempt := 0
	for {
		select {
		case <-mp.stopDelay:
			return
		default:
			utime.Sleep(backoff(attempt, 100*time.Millisecond, 1*time.Second))
			msg, score, err := queue.priority.Top()
			if err != nil || score > time.Now().Unix() {
				attempt++

				continue
			}

			if _, xerr := queue.priority.Get(); xerr != nil {
				attempt++

				continue
			}

			attempt = 0
			queue.linked.Put(msg)
		}
	}
}

func (mp *memoryBroker) getSubChannel(topic, groupId string) (<-chan Message, error) {
	mp.Lock()
	defer mp.Unlock()

	subs, ok := mp.sublist[topic]
	if !ok {
		return nil, ErrSubscriberNil
	}

	for _, sub := range subs {
		if sub.groupId == groupId {
			return sub.output, nil
		}
	}
	return nil, ErrSubscriberNil
}

func (mp *memoryBroker) addSubscriber(topic, groupId string) {
	mp.Lock()
	defer mp.Unlock()

	subs, ok := mp.sublist[topic]
	if !ok {
		mp.sublist[topic] = []Subscriber{{
			groupId: groupId,
			output:  make(chan Message, 100),
		}}

		return
	}

	for _, sub := range subs { // exist
		if sub.groupId == groupId {
			return
		}
	}

	subs = append(subs, Subscriber{
		groupId: groupId,
		output:  make(chan Message, 100),
	})
	mp.sublist[topic] = subs
}

func (mp *memoryBroker) delSubscriber(topic, groupId string) {
	mp.Lock()
	defer mp.Unlock()

	subs, ok := mp.sublist[topic]
	if !ok {
		return
	}

	items := make([]Subscriber, 0)
	for _, sub := range subs {
		if sub.groupId != groupId {
			items = append(items, sub)
		}
	}
	mp.sublist[topic] = items
}

func (mp *memoryBroker) getSubscribers(topic string) ([]Subscriber, error) {
	mp.RLock()
	defer mp.RUnlock()

	subs, ok := mp.sublist[topic]
	if !ok {
		return nil, ErrSubscriberNil
	}

	result := make([]Subscriber, len(subs))
	copy(result, subs)

	return subs, nil
}

func backoff(attempt int, minVal, maxVal time.Duration) time.Duration {
	d := time.Duration(attempt*attempt) * minVal
	if d > maxVal {
		d = maxVal
	}

	return d
}
