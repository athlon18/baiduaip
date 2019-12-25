package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var (
	img    []byte
	client *Client
)

func TestMain(m *testing.M) {
	b, err := ioutil.ReadFile("./testdata/EF2B8C2B8BCD43DA931B218759D59C22.jpeg")
	if err != nil {
		log.Fatalln(err)
	}
	img = b
	img = make([]byte, len(b))
	copy(img, b)
	var key struct {
		AppID, APIKey, SecretKey string
	}
	b, err = ioutil.ReadFile("./testdata/key.json")
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(b, &key); err != nil {
		log.Fatalln(err)
	}
	if key.AppID == "" || key.APIKey == "" || key.SecretKey == "" {
		log.Fatalln("AppID|APIKey|SecretKey is empty")
	}
	client = NewClient(&Option{AppID: key.AppID, APIKey: key.APIKey, SecretKey: key.SecretKey, RefreshTime: 2591995})
	client.SetAccessTokenStore(nil)
	m.Run()
}

// 测试获取访问令牌
func TestClient_auth(t *testing.T) {
	type fields struct {
		AppID, APIKey, SecretKey string
		client                   *Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "getAccessTokenFailed",
			fields: fields{
				client: NewClient(&Option{AppID: "123456", APIKey: "123456", SecretKey: "123456"}),
			},
			wantErr: true,
		},
		{
			name: "getAccessTokenOK",
			fields: fields{
				client: client,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := tt.fields.client.auth()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != nil {
				t.Logf("accessToken: %s, expiresIn: %d", gotToken.AccessToken, gotToken.ExpiresIn)
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		uri  string
		v    url.Values
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "getAccessTokenOK",
			fields: fields{
				client: client,
			},
			args: args{
				uri: "https://aip.baidubce.com/rest/2.0/face/v3/detect",
				v: url.Values{
					"image":        []string{base64.StdEncoding.EncodeToString(img)},
					"image_type":   []string{"BASE64"},
					"face_field":   []string{"gender", "age"},
					"max_face_num": []string{"10"},
				},
				data: new(json.RawMessage),
			},
			wantErr: false,
		},
		{
			name: "getAccessTokenFailed",
			fields: fields{
				client: NewClient(&Option{AppID: "123456", APIKey: "123456", SecretKey: "123456"}),
			},
			args: args{
				uri: "https://aip.baidubce.com/rest/2.0/face/v3/detect",
				v: url.Values{
					"image":        []string{base64.StdEncoding.EncodeToString(img)},
					"image_type":   []string{"BASE64"},
					"face_field":   []string{"gender", "age"},
					"max_face_num": []string{"10"},
				},
				data: new(json.RawMessage),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.client.Do(http.MethodPost, tt.args.uri, "application/json", "json", strings.NewReader(tt.args.v.Encode()), tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// ExampleDo 发送并接收JSON请求和响应
func Example() {
	c := DefaultClient
	uri := "uri"
	var r bytes.Buffer
	data := new(json.RawMessage)
	if err := c.Do(http.MethodPost, uri, "application/json", "json", &r, data); err != nil {
		fmt.Printf("%s\n", err)
	}
}

// 测试从内存读取访问令牌
func TestClient_getAccessToken(t *testing.T) {
	type fields struct {
		client *Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "getAccessToken1",
			fields: fields{
				client: client,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				_, cached, err := tt.fields.client.GetAccessToken()
				if (err != nil) != tt.wantErr {
					t.Errorf("Client.getAccessToken() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("get access_token success, cached: %t", cached)
				// time.Sleep(time.Second)
			}
		})
	}
}
