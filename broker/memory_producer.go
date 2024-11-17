package broker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/xid"

	"github.com/joshqu1985/lego/container"
	"github.com/joshqu1985/lego/utils"
)

func NewMemoryProducer(conf Config) (Producer, error) {
	return &memoryProducer{
		broker: getMemoryBroker(conf.Topics),
	}, nil
}

type memoryProducer struct {
	broker *memoryBroker
}

func (this *memoryProducer) Send(ctx context.Context, topic string, msg *Message, args ...int64) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}
	msg.Topic = topic
	msg.MessageId = xid.New().String()

	data, ok := this.broker.queues.Load(topic)
	if !ok {
		return fmt.Errorf("topic not found")
	}

	queue, _ := data.(*TopicQueue)
	if len(args) != 0 && args[0] > time.Now().Unix() {
		queue.priority.Put(msg, args[0])
	} else {
		queue.linked.Put(msg)
	}
	return nil
}

func (this *memoryProducer) Close() error {
	return nil
}

var (
	broker *memoryBroker
	once   sync.Once
)

func getMemoryBroker(topics map[string]string) *memoryBroker {
	once.Do(func() {
		keys := []string{}
		for key := range topics {
			keys = append(keys, key)
		}
		broker = newMemoryBroker(keys)
	})
	return broker
}

func newMemoryBroker(topics []string) *memoryBroker {
	broker := &memoryBroker{
		sublist: map[string][]Subscriber{},
	}
	for _, topic := range topics {
		broker.queues.LoadOrStore(topic, &TopicQueue{
			linked:   container.NewLinked(),
			priority: container.NewPriority(),
		})
	}
	broker.run()
	return broker
}

type memoryBroker struct {
	queues sync.Map

	sync.RWMutex
	sublist map[string][]Subscriber
}

func (this *memoryBroker) run() {
	this.queues.Range(func(key, val any) bool {
		queue := val.(*TopicQueue)
		go this.fetchWorker(key.(string), queue)
		go this.delayWorker(queue)
		return true
	})
}

func (this *memoryBroker) fetchWorker(topic string, queue *TopicQueue) {
	attempt := 0
	for {
		utils.Sleep(utils.Backoff(attempt, 100*time.Millisecond, 1*time.Second))

		sublist, err := this.getSubscribers(topic)
		if err != nil {
			attempt++
			continue
		}

		if len(sublist) == 0 {
			attempt++
			continue
		}

		data, err := queue.linked.Get()
		if err != nil {
			attempt++
			continue
		}
		msg := data.(*Message)

		for _, sub := range sublist {
			sub.output <- *msg
		}
		attempt = 0
	}
}

func (this *memoryBroker) delayWorker(queue *TopicQueue) {
	attempt := 0
	for {
		utils.Sleep(utils.Backoff(attempt, 100*time.Millisecond, 1*time.Second))

		msg, score, err := queue.priority.Top()
		if err != nil {
			attempt++
			continue
		}

		if score > time.Now().Unix() {
			attempt++
			continue
		}

		if _, err := queue.priority.Get(); err != nil {
			attempt++
			continue
		}

		attempt = 0
		queue.linked.Put(msg)
	}
}

func (this *memoryBroker) getSubChannel(topic, groupId string) (<-chan Message, error) {
	this.Lock()
	defer this.Unlock()

	subs, ok := this.sublist[topic]
	if !ok {
		return nil, fmt.Errorf("subscriber not found")
	}

	for _, sub := range subs {
		if sub.groupId == groupId {
			return sub.output, nil
		}
	}
	return nil, fmt.Errorf("subscriber not found")
}

func (this *memoryBroker) addSubscriber(topic, groupId string) {
	this.Lock()
	defer this.Unlock()

	subs, ok := this.sublist[topic]
	if !ok {
		this.sublist[topic] = []Subscriber{{
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
	this.sublist[topic] = subs
}

func (this *memoryBroker) delSubscriber(topic, groupId string) {
	this.Lock()
	defer this.Unlock()

	subs, ok := this.sublist[topic]
	if !ok {
		return
	}

	items := []Subscriber{}
	for _, sub := range subs {
		if sub.groupId != groupId {
			items = append(items, sub)
		}
	}
	this.sublist[topic] = items
}

func (this *memoryBroker) getSubscribers(topic string) ([]Subscriber, error) {
	this.Lock()
	defer this.Unlock()

	subs, ok := this.sublist[topic]
	if !ok {
		return nil, fmt.Errorf("subscriber not found")
	}
	return subs, nil
}

type Subscriber struct {
	groupId string
	output  chan Message
}

type TopicQueue struct {
	linked   *container.Linked
	priority *container.Priority
}
