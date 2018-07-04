package core

import (
	"container/list"
	"sync"
)

type SubscriptionConfig struct {
	Name          string         `json:"name"`
	Topic         string         `json:"topic"`
	Tag           string         `json:"tag"`
	HandlerConfig *HandlerConfig `json:"handler_config"`
}

type Subscription struct {
	Config           *SubscriptionConfig
	mutex            sync.Mutex
	messages         *list.List
	handler          Handler
	newMessageSignal chan struct{}
	tickets          chan struct{}
}

func (this *Subscription) Start() error {
	return nil
}

func (this *Subscription) Stop() error {
	return nil
}

func (this *Subscription) Process(msg *Message) {
	this.tickets <- struct{}{}
	this.process(msg)
}

func (this *Subscription) process(msg *Message) {
	for {
		this.handler.Process(msg)
	}
}
