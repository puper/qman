package base

import (
	"encoding/json"
)

type ErrorInter interface {
	GetCode() int
	GetMessage() string
	Error() string
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (this *Error) GetCode() int {
	return this.Code
}

func (this *Error) GetMessage() string {
	return this.Message
}

func (this *Error) Error() string {
	bs, _ := json.Marshal(this)
	return string(bs)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	ParseError     = NewError(-32700, "Parse error")
	InvalidRequest = NewError(-32600, "Invalid Request")
	MethodNotFound = NewError(-32601, "Method not found")
	InvalidParams  = NewError(-32602, "Invalid params")
	InternalError  = NewError(-32603, "Internal error")
	AuthFailed     = NewError(-33000, "Auth failed")
)
