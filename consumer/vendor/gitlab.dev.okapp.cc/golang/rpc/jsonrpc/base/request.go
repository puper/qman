package base

import (
	"context"
)

type Request struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  interface{}     `json:"params"`
	Context context.Context `json:"-"`
}

func (this *Request) DecodeParams(v interface{}) error {
	err := Decode(this.Params, v)
	return err
}
