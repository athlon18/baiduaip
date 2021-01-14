package ocr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/athlon18/baiduaip/baidu/aip/client"
)

func postJSON(uri string, req, v interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}
	return client.DefaultClient.Do(http.MethodPost, uri, jsonContentTypeOfRequest, jsonContentTypeResponse, bytes.NewReader(b), v)
}

func post(uri string, req url.Values, v interface{}) error {
	return client.DefaultClient.Do(http.MethodPost, uri, formDataTypeOfRequest, jsonContentTypeResponse, strings.NewReader(req.Encode()), v)
}
