package mysql

import (
	"time"

	"code.int.thoseyears.com/golang/ppgo/helpers"
	"github.com/puper/qman/consumer/core"
)

func init() {
	core.RegisterStorage("mysql", NewStorage)
}

type Storage struct {
	Config *Config
}

type Config struct {
	ConnectionName string `json:"connection_name"`
}

func NewStorage(in interface{}) (core.Storage, error) {
	config := new(Config)
	err := helpers.StructDecode(in, config, "json")
	if err != nil {
		return nil, err
	}
	return &Storage{
		Config: config,
	}, nil
}

func (this *Storage) WatchSubscriptionChange(cb core.SubscriptionChangeCallback) {
	tk := time.NewTicker(time.Second)
	for range tk.C {
		cb(&core.Event{
			Type: core.EVENT_CREATE,
			Data: core.SubscriptionConfig{
				Name:  "testName",
				Topic: "testTopic",
				Tag:   "testTag",
				HandlerConfig: core.HandlerConfig{
					Name: "jsonrpc",
				},
			},
		})
	}
}
