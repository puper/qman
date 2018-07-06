package producer

import (
	"encoding/json"
)

type Message struct {
	Topic string `json:"-"`
	Key   string `json:"-"`

	BusinessID string `json:"business_id"`
	Tag        string `json:"tag"`
	Value      string `json:"value"`
}

func (this *Message) Encode() []byte {
	bs, _ := json.Marshal(this)
	return bs
}
