package core

import (
	"sync"
)

type SubscriptionConfig struct {
	Name          string        `json:"name"`
	Topic         string        `json:"topic"`
	Tag           string        `json:"tag"`
	HandlerConfig HandlerConfig `json:"handler_config"`
}

type Subscription struct {
	Config       *SubscriptionConfig
	mutex        sync.Mutex
	handlerMutex sync.RWMutex
	handler      Handler
}

func NewSubscription(config *SubscriptionConfig) *Subscription {
	return &Subscription{
		Config: config,
	}
}

func (this *Subscription) UpdateConfig(config *SubscriptionConfig) {
	h, err := NewHandler(&config.HandlerConfig)
	if err == nil {
		this.Config = config
		this.handlerMutex.Lock()
		this.handler = h
		this.handlerMutex.Unlock()
	}
}

func (this *Subscription) Start() error {
	return nil
}

func (this *Subscription) Stop() error {
	return nil
}

func (this *Subscription) Process(msg *MessageWithResult) {
	this.handlerMutex.RLock()
	handler := this.handler
	this.handlerMutex.RUnlock()
	handler.Process(msg)
}
