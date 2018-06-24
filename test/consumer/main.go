package main

import (
	"github.com/Shopify/sarama"
)

type QMan struct {
}

type Message struct {
	sarama.ConsumerMessage
}

type Subscription struct {
}

func (this *QMan) Put(msg *Message) {
}

func (this *Subscription) Put(msg *Message) {
	//每个订阅一个offset
}

type Storage struct {
}
