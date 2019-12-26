package face

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestDetect(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/EF2B8C2B8BCD43DA931B218759D59C22.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	req := NewDetectRequest(base64.StdEncoding.EncodeToString(b), "BASE64")
	req.MaxFaceNum = 10
	req.FaceField = "age,face_type,beauty,expression,face_shape,gender,glasses,eye_status,emotion,landmark,landmark72,landmark150,quality"
	type args struct {
		req *DetectRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{req: req},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Detect(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Detect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}
