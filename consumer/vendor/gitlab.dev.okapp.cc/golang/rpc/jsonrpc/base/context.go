package base

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Context struct {
	values map[string]string
}

func NewContext(values map[string]string) *Context {
	return &Context{
		values: values,
	}
}

func (this *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (this *Context) Done() <-chan struct{} {
	return nil
}

func (this *Context) Err() error {
	return nil
}

func (this *Context) Value(key interface{}) interface{} {
	k, ok := key.(string)
	if !ok {
		return nil
	}
	v := this.values[k]
	return v
}

func (this *Context) Values() map[string]string {
	return this.values
}

func (this *Context) String() string {
	v, err := json.Marshal(this.values)
	if err != nil {
		v = []byte("{}")
	}
	return base64.StdEncoding.EncodeToString(v)
}

func ContextFromGin(ctx *gin.Context) *Context {
	v := make(map[string]string)
	data, err := base64.StdEncoding.DecodeString(ctx.GetHeader("X-JSONRPC-CONTEXT"))
	if err != nil {
		return NewContext(v)
	}
	json.Unmarshal(data, &v)
	//add client ip
	v["remoteAddr"] = strings.Split(ctx.Request.RemoteAddr, ":")[0]
	return NewContext(v)
}
