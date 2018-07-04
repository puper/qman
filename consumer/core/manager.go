package core

import (
	"sync"

	"github.com/Shopify/sarama"
)

type Manager struct {
	storage Storage

	topicConsumersMutex sync.RWMutex
	topicConsumers      map[string]*TopicConsumer

	Config   *Config
	Consumer sarama.Consumer
}

func New(config *Config) (*Manager, error) {
	var (
		err error
	)
	m := &Manager{
		Config:         config,
		topicConsumers: make(map[string]*TopicConsumer),
	}
	m.Consumer, err = sarama.NewConsumer(m.Config.Brokers, nil)
	return m, err
}

func (this *Manager) SetStorage(s Storage) {
	this.storage = s
}

func (this *Manager) Start() error {
	this.storage.WatchSubscriptionChange(this.onSubscriptionChange)
	return nil
}

func (this *Manager) Stop() error {
	for _, tc := range this.topicConsumers {
		tc.Stop()
	}
	return this.Consumer.Close()
}

func (this *Manager) onSubscriptionChange(ev *Event) {
	tc, ok := this.topicConsumers[ev.Data.Topic]
	if ev.Type == EVENT_DELETE {
		if ok {
			tc.DeleteSubscription(ev.Data.Tag, ev.Data.Name)
			if tc.GetSubscriptionCount() == 0 {
				tc.Stop()
				delete(this.topicConsumers, ev.Data.Topic)
			}
		}
	} else {
		if !ok {
			tc = NewTopicConsumer(this, ev.Data.Topic)
			this.topicConsumers[ev.Data.Topic] = tc
		}
		sub, ok := tc.GetSubscription(ev.Data.Tag, ev.Data.Name)
		if !ok {
			sub = NewSubscription(&ev.Data)
			tc.SetSubscription(ev.Data.Tag, ev.Data.Name, sub)
		} else {
			sub.UpdateConfig(&ev.Data)
		}

	}
}
