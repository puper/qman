package jsonrpc

import (
	"time"

	"code.int.thoseyears.com/golang/ppgo/helpers"
	"github.com/puper/qman/consumer/core"
	"gitlab.dev.okapp.cc/golang/rpc/jsonrpc/client"
)

func init() {
	core.RegisterHandler("jsonrpc", NewHandler)
}

type Handler struct {
	config *Config
	c      *client.Client
}

type Config struct {
	ServerAddr string        `json:"server_addr"`
	Method     string        `json:"method"`
	Timeout    time.Duration `json:"timeout"`
	RetryCount int           `json:"retry_count"`
}

func NewHandler(in interface{}) (core.Handler, error) {
	config := new(Config)
	err := helpers.StructDecode(in, config, "json")
	if err != nil {
		return nil, err
	}
	return &Handler{
		config: config,
		c:      client.New(config.ServerAddr, config.Timeout),
	}, nil
}

func (this *Handler) Process(msg *core.MessageWithResult) {
	arg := map[string]interface{}{
		"message": msg,
	}
	i := 0
	for {
		resp := this.c.CallTimeout(nil, this.config.Method, arg, this.config.Timeout)
		if resp.Error == nil {
			return
		}
		if this.config.RetryCount == i {
			return
		}
		i++
	}
	msg.Done(nil)
}
