package face

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestSearch(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/tom.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		req *SearchRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "1:1",
			args:    args{req: NewSearchRequest(base64.StdEncoding.EncodeToString(b), "BASE64", "123456")},
			wantErr: false,
		},
		{
			name:    "1:N",
			args:    args{req: NewSearchRequest(base64.StdEncoding.EncodeToString(b), "BASE64", "123456")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "1:N" {
				tt.args.req.MaxUserNum = 10
			}
			gotRes, err := Search(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}

func TestMultiSearch(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/EF2B8C2B8BCD43DA931B218759D59C22.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		req *MultiSearchRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "M:N",
			args:    args{req: NewMultiSearchRequest(base64.RawStdEncoding.EncodeToString(b), "BASE64", "123456")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.req.MaxUserNum = 10
			tt.args.req.MaxFaceNum = 10
			tt.args.req.MatchThreshold = 0
			gotRes, err := MultiSearch(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}
