package base

type Response struct {
	ID      interface{} `json:"id,omitempty"`
	JsonRPC string      `json:"jsonrpc"`
	Error   *Error      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func NewResponse(id interface{}, err *Error, result interface{}) *Response {
	return &Response{
		ID:      id,
		JsonRPC: "2.0",
		Error:   err,
		Result:  result,
	}
}
