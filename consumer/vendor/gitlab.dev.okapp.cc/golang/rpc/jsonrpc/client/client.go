package client

import (
	"encoding/base64"
	"strconv"
	"sync/atomic"

	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	jsonrpc "gitlab.dev.okapp.cc/golang/rpc/jsonrpc/base"
)

type Client struct {
	Url        string
	id         int64
	c          *http.Client
	timeout    time.Duration
	AuthKey    string
	AuthSecret string
}

func New(url string, timeout time.Duration) *Client {
	client := new(Client)
	client.Url = url
	client.timeout = timeout
	return client
}

func (c *Client) Call(ctx *jsonrpc.Context, method string, params interface{}) *jsonrpc.Response {
	return c.callTimeout(ctx, method, params, false, c.timeout)
}

func (c *Client) CallTimeout(ctx *jsonrpc.Context, method string, params interface{}, timeout time.Duration) *jsonrpc.Response {
	return c.callTimeout(ctx, method, params, false, timeout)
}

func (c *Client) Notify(ctx *jsonrpc.Context, method string, params interface{}) *jsonrpc.Response {
	return c.callTimeout(ctx, method, params, true, c.timeout)
}

func (c *Client) NotifyTimeout(ctx *jsonrpc.Context, method string, params interface{}, timeout time.Duration) *jsonrpc.Response {
	return c.callTimeout(ctx, method, params, true, timeout)
}

func (c *Client) call(ctx *jsonrpc.Context, method string, params interface{}, notify bool) *jsonrpc.Response {
	return c.callTimeout(ctx, method, params, notify, 0)
}

func (c *Client) callTimeout(ctx *jsonrpc.Context, method string, params interface{}, notify bool, timeout time.Duration) *jsonrpc.Response {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
	}
	var id interface{}
	if !notify {
		id = atomic.AddInt64(&c.id, 1)
		payload["id"] = id
	} else {
		id = nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return jsonrpc.NewResponse(id, jsonrpc.InvalidParams, nil)
	}
	buf := bytes.NewBuffer(data)
	hc := &http.Client{}
	hc.Timeout = timeout
	request, err := http.NewRequest("POST", c.Url, buf)
	request.Header.Set("Content-type", "application/json")
	if c.AuthKey != "" {
		auth := &jsonrpc.AuthParams{
			AppKey:       c.AuthKey,
			Time:         time.Now().Unix(),
			RandomString: randomString(8),
		}
		auth.Sign = jsonrpc.SignData(auth.AppKey, strconv.Itoa(int(auth.Time)), auth.RandomString, c.AuthSecret, string(data))
		bs, _ := json.Marshal(auth)
		request.Header.Set("X-JSONRPC-AUTH", base64.StdEncoding.EncodeToString(bs))
	}
	if ctx == nil {
		ctx = jsonrpc.NewContext(map[string]string{})
	}
	request.Header.Set("X-JSONRPC-CONTEXT", ctx.String())
	resp, err := hc.Do(request)

	if err != nil {
		return jsonrpc.NewResponse(id, jsonrpc.NewError(-32604, err.Error()), nil)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	respPlayload := &jsonrpc.Response{}
	err = decoder.Decode(respPlayload)
	if notify {
		return jsonrpc.NewResponse(nil, nil, nil)
	}
	if err != nil {
		return jsonrpc.NewResponse(id, jsonrpc.NewError(-32605, "response not valid json"), nil)
	}

	return respPlayload
}

func (this *Client) Decode(in, out interface{}) error {
	return jsonrpc.Decode(in, out)
}

func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
