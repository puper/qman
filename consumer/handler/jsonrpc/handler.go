package jsonrpc

import (
	"time"

	"github.com/puper/qman/consumer/core"
	"gitlab.dev.okapp.cc/golang/rpc/jsonrpc/client"
)

type Handler struct {
	ServiceAddr string
	Method      string
	Timeout     time.Duration
	RetryCount  int
	c           *client.Client
}

func (this *Handler) Process(msg *core.Message) {
	arg := map[string]interface{}{
		"message": msg,
	}
	i := 0
	for {
		resp := this.c.CallTimeout(nil, this.Method, arg, this.Timeout)
		if resp.Error == nil {
			return
		}
		if this.RetryCount == i {
			return
		}
		i++
	}
}
