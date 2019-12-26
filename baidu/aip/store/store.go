package store

import (
	"errors"
	"sync"
	"time"
)

// AccessToken 访问令牌
type AccessToken struct {
	// AccessToken 访问令牌
	AccessToken string
	// ExpiredAt 过期时间
	ExpiredAt time.Time
	// ExpiresIn 有效期, 默认1个月
	// ExpiresIn int64
	// RefreshTime 最后一次更新时间
	// RefreshTime int64
}

// Expired 令牌过期
func (a *AccessToken) Expired(n int64) bool {
	ct := time.Now()
	// 令牌过期
	if ct.After(a.ExpiredAt) {
		return true
	}
	// 最少提前60秒刷新
	if n < 60 {
		n = 60
	}
	// 是否需要刷新令牌
	if d := a.ExpiredAt.Sub(ct); d <= time.Duration(n)*time.Second {
		return true
	}
	return false
}

var (
	// ErrNotFound 不存在
	ErrNotFound = errors.New("NotFound")
	// ErrTypeInvalid 值的类型不正确
	ErrTypeInvalid = errors.New("TypeInvalid")
	// ErrExpired 过期
	ErrExpired = errors.New("Expired")
)

// AccessTokenStore 访问令牌存储
type AccessTokenStore interface {
	Set(key string, token *AccessToken) error
	Get(key string) (*AccessToken, error)
}

// Memory 内存存储
type Memory struct {
	sync.Map
}

// DefaultAccessTokenStore 默认的访问令牌存储
func DefaultAccessTokenStore() AccessTokenStore {
	return &Memory{}
}

// Set 存储访问令牌
func (a *Memory) Set(key string, token *AccessToken) error {
	a.Store(key, token)
	return nil
}

// Get 查询访问令牌
func (a *Memory) Get(key string) (*AccessToken, error) {
	v, ok := a.Load(key)
	if !ok {
		return nil, ErrNotFound
	}
	token, ok := v.(*AccessToken)
	if !ok {
		return nil, ErrTypeInvalid
	}
	return token, nil
}
