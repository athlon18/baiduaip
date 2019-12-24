package face

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antlinker/baiduaip/baidu/aip/client"
)

func postJSON(uri string, req, v interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}
	return client.DefaultClient.Do(http.MethodPost, uri, jsonContentTypeOfRequest, jsonContentTypeResponse, bytes.NewReader(b), v)
}
