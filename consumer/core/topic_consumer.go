package core

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type TopicComsumer struct {
	Topic string

	manager  *Manager
	messages chan *Message

	keyConsumersMutex sync.RWMutex
	keyConsumers      map[string]*KeyConsumer

	subscriptionsMutex sync.RWMutex
	subscriptions      map[string]*Subscription
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
				this.messages <- msg
			}
		}(pc)
	}
	return nil
}

func (this *TopicComsumer) Stop() {

}

func (this *TopicComsumer) GetKeyConsumer(key string) *KeyConsumer {
	this.keyConsumersMutex.RLock()
	defer this.keyConsumersMutex.RUnlock()
	kc, ok := this.keyConsumers[key]
	if !ok {
	}
	return kc
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

func (this *TopicComsumer) loop() {
	for msg := range this.messages {
		if msg.Key != "" {
			this.GetKeyConsumer(msg.Key).Put(msg)
		} else {
			for _, sub := range this.GetSubscriptions(msg.Topic, msg.Tag) {
				sub.Put(msg)
			}
		}
	}
}
