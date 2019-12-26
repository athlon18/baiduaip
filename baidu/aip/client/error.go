package client

import (
	"encoding/json"
	"fmt"
)

// RequestError 请求错误
type RequestError interface {
	Code() int
	Message() string
}

// Response API响应数据
type Response struct {
	// 错误码
	ErrorCode int `json:"error_code"`
	// 错误描述信息
	ErrorMsg string `json:"error_msg"`
	// 请求ID
	LogID uint64 `json:"log_id"`
	// 时间戳
	Timestamp int64 `json:"timestamp"`
	// cached
	Cached int `json:"cached,omitempty"`
	// 响应结果
	Reult json.RawMessage `json:"result,omitempty"`
}

// Code 实现RequestError
func (res *Response) Code() int {
	return res.ErrorCode
}

// Message 实现RequestError
func (res *Response) Message() string {
	return res.ErrorMsg
}

// Error 实现error接口
func (res *Response) Error() string {
	return fmt.Sprintf("error_code: %d, error_msg: %s", res.ErrorCode, res.ErrorMsg)
}
