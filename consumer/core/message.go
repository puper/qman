package core

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

type Message struct {
	sarama.ConsumerMessage `json:"-"`
	Tag                    string `json:"tag"`
	Key                    string `json:"-"`
	Value                  []byte `json:"value"`
}

func NewMessage(cm *sarama.ConsumerMessage) (*Message, error) {
	msg := new(Message)
	err := json.Unmarshal(cm.Value, msg)
	if err != nil {
		return nil, err
	}
	msg.ConsumerMessage = *cm
	msg.Key = string(cm.Key)
	return msg, nil
}

func (this *Message) WithContext() {
	return
}

type MessageWithContext struct {
	Message *Message
}
