package core

import (
	"container/list"
	"sync"
)

type Subscription struct {
	mutex            sync.Mutex
	messages         *list.List
	handler          Handler
	newMessageSignal chan struct{}
	tickets          chan struct{}

	subscriptions map[string]map[string]*Subscription

	MaxConcurrentCount    int
	MaxFlightMessageCount int
}

func (this *Subscription) Put(msg *Message) {
	this.mutex.Lock()
	this.messages.PushFront(msg)
	this.mutex.Unlock()
	select {
	case this.newMessageSignal <- struct{}{}:
	default:
	}
}

func (this *Subscription) loop() {
	for {
		this.mutex.Lock()
		msg := this.messages.Back()
		this.mutex.Unlock()
		if msg == nil {
			<-this.newMessageSignal
		} else {
			this.process(msg)
		}
	}
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
