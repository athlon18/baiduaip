package face

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/antlinker/baiduaip/baidu/aip/client"
)

func TestMain(m *testing.M) {
	b, err := ioutil.ReadFile("../testdata/key.json")
	if err != nil {
		log.Fatalln(err)
	}
	var key struct {
		AppID, APIKey, SecretKey string
	}
	if err = json.Unmarshal(b, &key); err != nil {
		log.Fatalln(err)
	}
	if key.AppID == "" || key.APIKey == "" || key.SecretKey == "" {
		log.Fatalln("AppID|APIKey|SecretKey is empty")
	}
	client.Init(&client.Option{AppID: key.AppID, APIKey: key.APIKey, SecretKey: key.SecretKey, RefreshTime: 2591995})
	m.Run()
}

func TestGetGroupList(t *testing.T) {
	type args struct {
		req *GetGroupListRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				req: &GetGroupListRequest{Start: 0, Length: 50},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetGroupList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}

func TestDeleteGroup(t *testing.T) {
	type args struct {
		req *DeleteGroupRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &DeleteGroupRequest{
				GroupID: "12345678",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteGroup(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DeleteGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddGroup(t *testing.T) {
	type args struct {
		req *AddGroupRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				req: &AddGroupRequest{
					GroupID: "12345678",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddGroup(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("AddGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		req *DeleteUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &DeleteUserRequest{
				GroupID: "aa5533ff4a9d4f261e571a61c7342df5241e945e",
				UserID:  "d5a03fe796ca4547643c4ae5f1d79c64d144da22",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyUser(t *testing.T) {
	type args struct {
		req *CopyUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &CopyUserRequest{
				UserID:     "d5a03fe796ca4547643c4ae5f1d79c64d144da22",
				SrcGroupID: "aa5533ff4a9d4f261e571a61c7342df5241e945e",
				DstGroupID: "123456",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyUser(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CopyUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserList(t *testing.T) {
	type args struct {
		req *GetUserListRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &GetUserListRequest{
				GroupID: "123456",
				Start:   0,
				Length:  10,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetUserList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}

func TestGetFaceList(t *testing.T) {
	type args struct {
		req *GetUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &GetUserRequest{
				GroupID: "123456",
				UserID:  "d5a03fe796ca4547643c4ae5f1d79c64d144da22",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetFaceList(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFaceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		req *GetUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &GetUserRequest{
				GroupID: "@ALL",
				UserID:  "d5a03fe796ca4547643c4ae5f1d79c64d144da22",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetUser(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}

func TestDeleteFace(t *testing.T) {
	type args struct {
		req *DeleteFaceRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				req: &DeleteFaceRequest{
					GroupID:   "123456",
					UserID:    "d5a03fe796ca4547643c4ae5f1d79c64d144da22",
					FaceToken: "2f6df124fed06035cb2994923372061f",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFace(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/cp.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		req *AddUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &UpdateUserRequest{
				Image:     base64.StdEncoding.EncodeToString(b),
				ImageType: "BASE64",
				GroupID:   "123456",
				UserID:    "tom",
				UserInfo:  "tom_face",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := UpdateUser(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/man4.jpg")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		req *AddUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{req: &AddUserRequest{
				Image:     base64.StdEncoding.EncodeToString(b),
				ImageType: "BASE64",
				GroupID:   "123456",
				UserID:    "man4",
				UserInfo:  "图片中的第四个男人",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := AddUser(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != nil {
				b, _ := json.Marshal(gotRes)
				t.Logf("gotRes: %s", b)
			}
		})
	}
}
