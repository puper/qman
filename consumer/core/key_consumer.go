package core

import (
	"container/list"
	"sync"
)

type KeyConsumer struct {
	manager *Manager

	messagesMutex sync.Mutex
	messages      *list.List
}

func (this *KeyConsumer) Start() error {

}

func (this *KeyConsumer) Stop() {

}

func (this *KeyConsumer) Put(msg *Message) {
	this.messagesMutex.Lock()
	this.messages.PushFront(msg)
	this.messagesMutex.Unlock()
}

func (this *KeyConsumer) loop() {
	for msg := range this.messages {
		for _, sub := range this.manager.GetSubscriptions(msg.Topic, msg.Tag) {
			sub.Put(msg)
		}
	}
}
