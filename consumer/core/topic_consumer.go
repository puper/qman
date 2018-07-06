package core

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type TopicConsumer struct {
	Topic string

	manager           *Manager
	keyConsumersMutex sync.RWMutex
	keyConsumers      map[string]*KeyConsumer

	subscriptionsMutex sync.RWMutex
	subscriptions      map[string]map[string]*Subscription
}

func NewTopicConsumer(manager *Manager, topic string) *TopicConsumer {
	return &TopicConsumer{
		manager:       manager,
		Topic:         topic,
		keyConsumers:  make(map[string]*KeyConsumer),
		subscriptions: make(map[string]map[string]*Subscription),
	}
}

func (this *TopicConsumer) Start() error {
	partitions, err := this.manager.Consumer.Partitions(this.Topic)
	if err != nil {
		return err
	}
	if len(partitions) == 0 {
		return fmt.Errorf("no partitions of %v", this.Topic)
	}
	for _, p := range partitions {
		pc, err := this.manager.Consumer.ConsumePartition(this.Topic, p, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		go func(pc sarama.PartitionConsumer) {
			for messsage := range pc.Messages() {
				msg, err := NewMessage(messsage)
				if err != nil {
					continue
				}
				this.Put(msg)
			}
		}(pc)
	}
	return nil
}

func (this *TopicConsumer) Stop() error {
	for _, kc := range this.keyConsumers {
		if err := kc.Stop(); err != nil {
			return err
		}
	}
	return nil
}

func (this *TopicConsumer) Put(msg *Message) {
	if msg.Key == "" {
		this.subscriptionsMutex.RLock()
		for _, sub := range this.subscriptions[msg.Tag] {
			sub.Process(msg.WithResult())
		}
		this.subscriptionsMutex.RUnlock()
	} else {
		this.keyConsumersMutex.RLock()
		kc, ok := this.keyConsumers[msg.Key]
		kc.Put(msg)
		this.keyConsumersMutex.RUnlock()
		if !ok {
			this.keyConsumersMutex.Lock()
			kc = NewKeyConsumer(this)
			this.keyConsumers[msg.Key] = kc
			kc.Put(msg)
			this.keyConsumersMutex.Unlock()
		}
	}
}

//统统不加锁，使用的地方加锁
func (this *TopicConsumer) DeleteSubscription(tag, name string) {
	delete(this.subscriptions[tag], name)
	if len(this.subscriptions[tag]) == 0 {
		delete(this.subscriptions, tag)
	}
}

func (this *TopicConsumer) SetSubscription(tag, name string, sub *Subscription) {
	_, ok := this.subscriptions[tag]
	if !ok {
		this.subscriptions[tag] = make(map[string]*Subscription)
	}
	this.subscriptions[tag][name] = sub
}

func (this *TopicConsumer) GetSubscription(tag, name string) (*Subscription, bool) {
	sub, ok := this.subscriptions[tag][name]
	return sub, ok
}

func (this *TopicConsumer) GetSubscriptionCount() int {
	return len(this.subscriptions)
}
