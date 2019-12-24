package client

// RequestError 请求错误
type RequestError interface {
	Code() int
	Message() string
}
