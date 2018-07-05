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

func (this *Message) WithResult() *MessageWithResult {
	return &MessageWithResult{
		Message: this,
		done:    make(chan struct{}, 1),
	}
}

type MessageWithResult struct {
	Message *Message
	result  *MessageResult
	done    chan struct{}
}

type MessageResult struct {
	Success  bool
	Response string
}

func (this *MessageWithResult) Done(result *MessageResult) {
	this.result = result
	close(this.done)
}

func (this *MessageWithResult) Result() interface{} {
	<-this.done
	return this.result
}
