package core

import (
	"container/list"
	"sync"
)

const (
	STATE_NOT_STARTED = 0
	STATE_STARTED     = 1
	STATE_STOPPED     = 2
)

type KeyConsumer struct {
	tc *TopicComsumer

	messagesMutex sync.Mutex
	messages      *list.List
	state         int
}

func NewKeyConsumer(tc *TopicComsumer) *KeyConsumer {
	kc := &KeyConsumer{
		tc:       tc,
		messages: list.New(),
	}
	return kc
}

func (this *KeyConsumer) Start() error {
	return nil
}

func (this *KeyConsumer) Stop() error {
	return nil
}

func (this *KeyConsumer) Put(msg *Message) bool {
	this.messagesMutex.Lock()
	defer this.messagesMutex.Unlock()
	if this.state == STATE_NOT_STARTED {
		this.messages.PushFront(msg)
		go this.loop()
		return true
	} else if this.state == STATE_STARTED {
		this.messages.PushFront(msg)
		return true
	}
	return false
}

func (this *KeyConsumer) loop() {
	for {
		this.messagesMutex.Lock()
		e := this.messages.Back()
		if e == nil {
			break
		}
		this.messagesMutex.Unlock()
		msg := e.Value.(*Message)
		this.tc.subscriptionsMutex.RLock()
		for _, sub := range this.tc.subscriptions[msg.Tag] {
			sub.Process(msg)
		}
		this.tc.subscriptionsMutex.RUnlock()
	}
}
