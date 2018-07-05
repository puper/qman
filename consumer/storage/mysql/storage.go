package mysql

import (
	"fmt"
	"time"

	"code.int.thoseyears.com/golang/ppgo/helpers"
	"github.com/puper/qman/consumer/core"
)

func init() {
	core.RegisterStorage("mysql", NewStorage)
}

type Storage struct {
	tk            *time.Ticker
	stopSignal    chan struct{}
	watchCallback core.SubscriptionChangeCallback
	Config        *Config
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
		Config:     config,
		stopSignal: make(chan struct{}),
	}, nil
}

func (this *Storage) SetWatchCallback(cb core.SubscriptionChangeCallback) {
	this.watchCallback = cb
}

func (this *Storage) Start() error {
	if this.watchCallback == nil {
		return fmt.Errorf("no subscription change callback set")
	}
	this.tk = time.NewTicker(time.Second)
	for {
		select {
		case <-this.tk.C:
			this.watchCallback(&core.Event{
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
		case <-this.stopSignal:
			this.tk.Stop()
			return nil
		}
	}
	return nil
}

func (this *Storage) Stop() error {
	select {
	case this.stopSignal <- struct{}{}:
	default:
	}
	return nil
}
