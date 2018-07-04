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
	var (
		err error
	)
	m := &Manager{
		Config:         config,
		topicConsumers: make(map[string]*TopicComsumer),
	}
	m.Consumer, err = sarama.NewConsumer(m.Config.Brokers, nil)
	return m, err
}

func (this *Manager) Start() error {
	this.storage.WatchSubscriptionChange(this.onSubscriptionChange)
	return nil
}

func (this *Manager) Stop() error {
	return this.Consumer.Close()
}

func (this *Manager) onSubscriptionChange(ev *Event) {
	if ev.Type == EVENT_DELETE {
	} else {

	}
}
