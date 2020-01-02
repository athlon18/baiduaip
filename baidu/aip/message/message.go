// Package message 百度智能云的响应信息
package message

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
	Result json.RawMessage `json:"result,omitempty"`
}

// Code 返回响应的错误码, 如果不是0, 则请求失败
func (res *Response) Code() int {
	return res.ErrorCode
}

// Message 返回错误信息描述
func (res *Response) Message() string {
	return res.ErrorMsg
}

// Error Response 的错误信息格式为错误码和错误消息
func (res *Response) Error() string {
	return fmt.Sprintf("error_code: %d, error_msg: %s", res.ErrorCode, res.ErrorMsg)
}
