package core

type HandlerConfig struct {
	Name   string
	Config interface{}
}

type Handler interface {
	Process(*Message)
}
