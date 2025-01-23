package errorx

import (
	"encoding/json"
)

type ResponseError struct {
	Data any
	Code int
}

func (e *ResponseError) Error() string {
	jsonStr, err := json.Marshal(e.Data)
	if err != nil {
		return "error"
	}
	return string(jsonStr)
}

func WithCodeResponse(code int, resp any) *ResponseError {
	return &ResponseError{
		Data: resp,
		Code: code,
	}
}
