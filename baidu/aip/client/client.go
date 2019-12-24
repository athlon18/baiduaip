package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/antlinker/baiduaip/baidu/aip/store"
)

const (
	// 认证地址
	authURL = `https://aip.baidubce.com/oauth/2.0/token`
	// 授权类型 key
	grantType = `grant_type`
	// 授权类型 value
	clientCredentials = `client_credentials`
	// 客户端ID key
	clientID = `client_id`
	// 客户端密钥 key
	clientSecret = `client_secret`

	// 请求类型
	contentType = "Content-Type"
	// formContentType 类型
	formContentType = "application/x-www-form-urlencoded"
)

var (
	// DefaultClient 初始默认客户端
	DefaultClient *Client
)

// Client 百度ai的客户端
type Client struct {
	sync.Mutex
	option           *Option
	client           *http.Client
	accessTokenStore store.AccessTokenStore
}

// NewClient 新建客户端
func NewClient(opts ...*Option) *Client {
	opt := mergeOptions(opts...)
	if opt.APIKey == "" || opt.SecretKey == "" {
		panic("appID or appSecret is empty")
	}
	client := &Client{
		option: opt,
	}
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   opt.ConnectTimeout,
			KeepAlive: opt.KeepAlive,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          10,
		IdleConnTimeout:       opt.IdleConnTimeout,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client.client = &http.Client{
		Transport: tr,
	}
	return client
}

// InitClient 初始化客户端
func InitClient(opts ...*Option) *Client {
	return NewClient(opts...)
}

// SetAccessTokenStore 设置访问令牌存储
func (c *Client) SetAccessTokenStore(s store.AccessTokenStore) {
	c.Lock()
	defer c.Unlock()
	if s == nil {
		c.accessTokenStore = store.DefaultAccessTokenStore()
		return
	}
	c.accessTokenStore = s
}

// Option 选项
type Option struct {
	// AppID
	AppID string
	// apiKey
	APIKey string
	// SecretKey
	SecretKey string
	// IsCloudUser
	IsCloudUser bool
	// AccessToken刷新时间,单位秒
	RefreshTime int64
	// 连接超时时间
	ConnectTimeout time.Duration
	// 长连接时间
	KeepAlive time.Duration
	// 空闲连接超时时间
	IdleConnTimeout time.Duration
}

// DefaultOptions 默认选项
func DefaultOptions() *Option {
	return &Option{
		RefreshTime:     3600, // 默认提前1小时刷新令牌
		ConnectTimeout:  30 * time.Second,
		KeepAlive:       30 * time.Second,
		IdleConnTimeout: 90 * time.Second,
	}
}

// mergeOptions 合并选项
func mergeOptions(opts ...*Option) *Option {
	o := DefaultOptions()

	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.AppID != "" {
			o.AppID = opt.AppID
		}
		if opt.APIKey != "" {
			o.APIKey = opt.APIKey
		}
		if opt.SecretKey != "" {
			o.SecretKey = opt.SecretKey
		}
		if opt.RefreshTime > 0 {
			o.RefreshTime = opt.RefreshTime
		}
		if opt.KeepAlive > 0 {
			o.KeepAlive = opt.KeepAlive
		}
		if opt.IdleConnTimeout > 0 {
			o.IdleConnTimeout = opt.KeepAlive
		}
		if opt.ConnectTimeout > 0 {
			o.ConnectTimeout = opt.ConnectTimeout
		}
		if opt.IsCloudUser {
			o.IsCloudUser = opt.IsCloudUser
		}
	}
	return o
}

// AccessToken 访问令牌
type AccessToken struct {
	// AccessToken 访问令牌
	AccessToken string `json:"access_token"`
	// ExpiresIn 有效期, 默认1个月
	ExpiresIn int64 `json:"expires_in"`
	// Error 错误信息
	Error string `json:"error"`
	// ErrorDescription 错误描述
	ErrorDescription string `json:"error_description"`
}

// auth 客户端获取访问令牌
func (c *Client) getAccessToken() (token string, err error) {
	// 如果存储为空
	if c.accessTokenStore == nil {
		t, err := c.auth()
		if err != nil {
			return "", err
		}
		return t.AccessToken, nil
	}
	// 从存储获取令牌
	t, err := c.accessTokenStore.Get(c.option.AppID)
	if err != nil && err != store.ErrNotFound {
		return "", fmt.Errorf("访问令牌存储类型错误:%w", err)
	}
	// 检查是否有效和过期
	if t != nil && !t.Expired(c.option.RefreshTime) {
		return t.AccessToken, nil
	}
	c.Lock()
	defer c.Unlock()
	// 获取令牌
	accessToken, err := c.auth()
	if err != nil {
		return "", err
	}
	// 保存令牌
	t = &store.AccessToken{
		AccessToken: accessToken.AccessToken,
		ExpiredAt:   time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second),
	}
	c.accessTokenStore.Set(c.option.AppID, t)
	return t.AccessToken, nil
}

// auth 客户端鉴权认证
func (c *Client) auth() (token *AccessToken, err error) {
	// 设置参数
	values := url.Values{}
	values.Set(grantType, clientCredentials)
	values.Set(clientID, c.option.APIKey)
	values.Set(clientSecret, c.option.SecretKey)
	req, err := http.NewRequest(http.MethodPost, authURL, strings.NewReader(values.Encode()))
	if err != nil {
		err = fmt.Errorf("http.NewRequest: %w", err)
		return
	}
	// 填充请求头
	req.Header.Set(contentType, formContentType)
	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	if c.parseReponse(resp, "json", &token); err != nil {
		err = fmt.Errorf("http request: %w", err)
		return
	}
	if token == nil {
		err = fmt.Errorf("get access_token error: 未知原因")
		return
	}
	if token.Error != "" {
		err = fmt.Errorf("get access_token error: %s, description: %s", token.Error, token.ErrorDescription)
		return
	}
	return
}

// newRequestWithAccessToken 新建http带认证访问令牌的请求
func (c *Client) newRequestWithAccessToken(method, uri, ctype string, body io.Reader) (*http.Request, error) {
	accessToken, err := c.getAccessToken()
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("ur.Parse: %w", err)
	}
	params := u.Query()
	params.Set("access_token", accessToken)
	u.RawQuery = params.Encode()
	r, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	r.Header.Set(contentType, ctype)
	return r, nil
}

// Request 执行http请求
func (c *Client) doRequestWithAccessToken(method, uri, contentType, typ string, body io.Reader, data interface{}) error {
	req, err := c.newRequestWithAccessToken(method, uri, contentType, body)
	if err != nil {
		return err
	}
	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}
	return c.parseReponse(resp, typ, data)
}

// Do 执行请求
func (c *Client) Do(method, uri, contentType, typ string, body io.Reader, data interface{}) (err error) {
	err = c.doRequestWithAccessToken(method, uri, contentType, typ, body, &data)
	if data == nil {
		return
	}
	v, ok := data.(RequestError)
	if !ok {
		return
	}
	// 令牌无效或者过期, 重试一次
	if errCode := v.Code(); errCode == 110 || errCode == 111 {
		err = c.doRequestWithAccessToken(method, uri, contentType, typ, body, &data)
	}
	return
}

// doRequest 执行请求
func (c *Client) doRequest(req *http.Request) (res *http.Response, err error) {
	res, err = c.client.Do(req)
	if err != nil {
		err = fmt.Errorf("http.Client.Do: %w", err)
		return
	}
	return res, nil
}

// parseReponse 解析响应, type可以是json,xml,v是的值
func (c *Client) parseReponse(resp *http.Response, typ string, v interface{}) (err error) {
	if resp == nil {
		err = errors.New("响应数据为空")
		return
	}
	defer resp.Body.Close()

	switch typ {
	case "raw": // 直接读取body
		w, ok := v.(io.Writer)
		if !ok {
			err = fmt.Errorf("parse %s 必须实现io.Writer", typ)
			return
		}
		if _, err := io.Copy(w, resp.Body); err != nil {
			return fmt.Errorf("read %s data: %w", typ, err)
		}
	case "json": // 解析为json
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body: %w", err)
		}
		err = json.Unmarshal(b, &v)
		if err != nil {
			err = fmt.Errorf("parse %s data: %w", typ, err)
		}
	default:
		err = errors.New("不支持的类型")
		return
	}
	return nil
}
