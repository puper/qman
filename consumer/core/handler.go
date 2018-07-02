package core

type Handler interface {
	Process(*Message)
}
