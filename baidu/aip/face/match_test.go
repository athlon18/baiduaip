package face

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestMatch(t *testing.T) {
	b1, err := ioutil.ReadFile("../testdata/tom.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadFile("../testdata/cp.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		req []*MatchRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				req: []*MatchRequest{
					NewMatchRequest(base64.StdEncoding.EncodeToString(b1), "BASE64"),
					NewMatchRequest(base64.StdEncoding.EncodeToString(b2), "BASE64"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Match(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}
