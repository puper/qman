package core

import (
	"sync"

	"github.com/Shopify/sarama"
)

type Manager struct {
	storage Storage

	topicConsumersMutex sync.RWMutex
	topicConsumers      map[string]*TopicComsumer

	Config   *Config
	Consumer sarama.Consumer
}

func New(config *Config) (*Manager, error) {
	m := &Manager{
		Config:         config,
		topicConsumers: make(map[string]*TopicComsumer),
	}
	return m, nil
}

func (this *Manager) Start() error {
	var err error
	this.Consumer, err = sarama.NewConsumer(this.Config.Brokers, nil)
	if err != nil {
		return err
	}
	return this.storage.WatchSubscriptionChange(this.onSubscriptionChange)
}

func (this *Manager) Stop() error {
	return nil
}

func (this *Manager) onSubscriptionChange(ev *Event) {
	this.subscriptionsMutex.Lock()
	defer this.subscriptionsMutex.Unlock()
	if ev.Type == EVENT_DELETE {
		tc, ok := this.topicConsumers[ev.Type]
		if ok {

		}
	} else {

	}
}

func (this *Manager) GetSubscriptions(topic, tag string) []*Subscription {
	this.subscriptionsMutex.RLock()
	defer this.subscriptionsMutex.RUnlock()
	subs := make([]*Subscription, len(this.subscriptions[topic][tag]))
	for _, sub := range this.subscriptions[topic][tag] {
		subs = append(subs, sub)
	}
	return subs
}
