package core

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type TopicComsumer struct {
	Topic string

	manager           *Manager
	keyConsumersMutex sync.RWMutex
	keyConsumers      map[string]*KeyConsumer

	subscriptionsMutex sync.RWMutex
	subscriptions      map[string]map[string]*Subscription
}

func (this *TopicComsumer) Start() error {
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

func (this *TopicComsumer) Stop() error {
	for _, kc := range this.keyConsumers {
		if err := kc.Stop(); err != nil {
			return err
		}
	}
	return nil
}

func (this *TopicComsumer) GetSubscriptions(tag string) []*Subscription {
	this.subscriptionsMutex.RLock()
	defer this.subscriptionsMutex.RUnlock()
	result := make([]*Subscription, len(this.subscriptions[tag]))
	for _, sub := range this.subscriptions[tag] {
		result = append(result, sub)
	}
	return result
}

func (this *TopicComsumer) Put(msg *Message) {
	if msg.Key == "" {
		this.subscriptionsMutex.RLock()
		for _, sub := range this.subscriptions[msg.Tag] {
			sub.Process(msg)
		}
		this.subscriptionsMutex.RUnlock()
	} else {
		for {

		}
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
