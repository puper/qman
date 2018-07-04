package core

import (
	"fmt"
)

var (
	handlers = make(map[string]NewHandlerFunc)
)

func RegisterHandler(name string, f NewHandlerFunc) {
	handlers[name] = f
}

type NewHandlerFunc func(interface{}) (Handler, error)

type HandlerConfig struct {
	Name   string
	Config interface{}
}

type Handler interface {
	Process(*Message)
}

func NewHandler(hc *HandlerConfig) (Handler, error) {
	creator, ok := handlers[hc.Name]
	if ok {
		return creator(hc.Config)
	}
	return nil, fmt.Errorf("handler `%v` not registered", hc.Name)
}
