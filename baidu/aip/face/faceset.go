package face

// AddUserRequest 人脸注册请求参数
type AddUserRequest struct {
	// 图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断。 两张图片通过json格式上传
	Image string `json:"image"`
	// 图片类型
	// BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M
	// URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
	// FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
	ImageType string `json:"image_type"`
	// 用户组id，标识一组用户（由数字、字母、下划线组成），长度限制48B。产品建议：根据您的业务需求，可以将需要注册的用户，按照业务划分，分配到不同的group下，例如按照会员手机尾号作为groupid，用于刷脸支付、会员计费消费等，这样可以尽可能控制每个group下的用户数与人脸数，提升检索的准确率
	GroupID string `json:"group_id"`
	// 用户id（由数字、字母、下划线组成），长度限制128B
	UserID string `json:"user_id"`
	// 用户资料，长度限制256B 默认空
	UserInfo string `json:"user_info,omitempty"`
	// 图片质量控制
	// NONE: 不进行控制
	// LOW:较低的质量要求
	// NORMAL: 一般的质量要求
	// HIGH: 较高的质量要求
	// 默认 NONE
	// 若图片质量不满足要求，则返回结果中会提示质量检测失败
	QualityControl string `json:"quality_control,omitempty"`
	// 活体检测控制
	// NONE: 不进行控制
	// LOW:较低的活体要求(高通过率 低攻击拒绝率)
	// NORMAL: 一般的活体要求(平衡的攻击拒绝率, 通过率)
	// HIGH: 较高的活体要求(高攻击拒绝率 低通过率)
	// 默认 NONE
	// 若活体检测结果不满足要求，则返回结果中会提示活体检测失败
	LivenessControl string `json:"liveness_control,omitempty"`
	// 操作方式
	// 如果请求接口: https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add
	// APPEND: 当user_id在库中已经存在时，对此user_id重复注册时，新注册的图片默认会追加到该user_id下
	// REPLACE : 当对此user_id重复注册时,则会用新图替换库中该user_id下所有图片, 注册或更新时使用
	// 默认使用APPEND
	// 如果请求接口: https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/update
	// 	UPDATE: 会使用新图替换库中该user_id下所有图片, 若user_id不存在则会报错
	// REPLACE : 当user_id不存在时, 则会注册这个user_id的用户
	// 默认使用UPDATE
	ActionType string `json:"action_type,omitempty"`
}

// NewAddUserRequest 新建人脸注册请求
func NewAddUserRequest(image, imageType, groupID, userID string) *AddUserRequest {
	return &AddUserRequest{
		Image:     image,
		ImageType: imageType,
		GroupID:   groupID,
		UserID:    userID,
	}
}

// AddUserResult 添加用户的响应结果
type AddUserResult struct {
	// 人脸图片的唯一标识
	FaceToken string `json:"face_token"`
	// 人脸在图片中的位置
	Location *Location `json:"location"`
}

// AddUser 人脸注册
func AddUser(req *AddUserRequest) (res *AddUserResult, err error) {
	err = postJSON(userAddURL, req, &res)
	return
}

// UpdateUserRequest 别名
type UpdateUserRequest = AddUserRequest

// UpdateUserResult 别名
type UpdateUserResult = AddUserResult

// UpdateUser 人脸更新
func UpdateUser(req *UpdateUserRequest) (res *UpdateUserResult, err error) {
	err = postJSON(userUpdateURL, req, &res)
	return
}

// DeleteFaceRequest 人脸删除的请求
type DeleteFaceRequest struct {
	// 用户组id（由数字、字母、下划线组成） 长度限制48B，删除指定group_id中的user_id信息
	GroupID string `json:"group_id"`
	// 用户id（由数字、字母、下划线组成），长度限制128B
	UserID string `json:"user_id"`
	// 需要删除的人脸图片token，（由数字、字母、下划线组成）长度限制64B
	FaceToken string `json:"face_token"`
}

// DeleteFace 人脸删除
func DeleteFace(req *DeleteFaceRequest) (err error) {
	return postJSON(faceDeleteURL, req, nil)
}

// GetUserRequest 用户信息查询
type GetUserRequest struct {
	// 用户组id(由数字、字母、下划线组成，长度限制48B)，如传入“@ALL”则从所有组中查询用户信息。
	// 注：处于不同组，但uid相同的用户，我们认为是同一个用户。
	GroupID string `json:"group_id"`
	// 用户id（由数字、字母、下划线组成），长度限制48B
	UserID string `json:"user_id"`
}

// NewGetUserRequest 新建用户查询请求
func NewGetUserRequest(userID, groupID string) *GetUserRequest {
	return &GetUserRequest{
		UserID:  userID,
		GroupID: groupID,
	}
}

// GetUserResult 查询用户信息的结果
type GetUserResult struct {
	// 查询到的用户列表
	UserList []*UserItem `json:"user_list"`
}

// UserItem 用户查询的列表项
type UserItem struct {
	// 用户组id，被查询用户的所在组
	GroupID string `json:"group_id"`
	// 用户资料，被查询用户的资料
	UserInfo string `json:"user_info"`
}

// GetUser 用户信息查询
func GetUser(req *GetUserRequest) (res *GetUserResult, err error) {
	err = postJSON(userGetURL, req, &res)
	return
}

// GetFaceListRequest 别名
// 注: 作为GetFaceList的参数时, "@ALL"无效
type GetFaceListRequest = GetUserRequest

// GetFaceListResult 获取用户人脸列表的结果
type GetFaceListResult struct {
	// 人脸列表
	FaceList []*ListItem `json:"face_list"`
}

// ListItem 人脸列表项
type ListItem struct {
	// 人脸图片的唯一标识
	FaceToken string `json:"face_token"`
	// 人脸创建时间
	Ctime string `json:"ctime"`
}

// GetFaceList 获取用户人脸列表
func GetFaceList(req *GetFaceListRequest) (res *GetFaceListResult, err error) {
	err = postJSON(faceGetlistURL, req, &res)
	return
}

// GetUserListRequest 获取用户列表的请求参数
type GetUserListRequest struct {
	// 用户组id，长度限制48B
	GroupID string `json:"group_id"`
	// 默认值0，起始序号
	Start uint32 `json:"start"`
	// 返回数量，默认值100，最大值1000
	Length uint32 `json:"length"`
}

// GetUserListResult 查询用户列表的结果
type GetUserListResult struct {
	// 用户ID列表
	UserIDList []string `json:"user_id_list"`
}

// GetUserList 获取用户列表
func GetUserList(req *GetUserListRequest) (res *GetUserListResult, err error) {
	err = postJSON(groupGetusersURL, req, &res)
	return
}

// CopyUserRequest 复制用户的请求
type CopyUserRequest struct {
	UserID     string `json:"user_id"`      // 用户id，长度限制48B
	SrcGroupID string `json:"src_group_id"` // 从指定组里复制信息
	DstGroupID string `json:"dst_group_id"` // 需要添加用户的组id
}

// CopyUser 复制用户
func CopyUser(req *CopyUserRequest) error {
	return postJSON(userCopyURL, req, nil)
}

// DeleteUserRequest 删除用户的请求
type DeleteUserRequest struct {
	// 用户组id（由数字、字母、下划线组成） 长度限制48B，删除指定group_id中的user_id信息
	GroupID string `json:"group_id"`
	// 用户id（由数字、字母、下划线组成），长度限制128B
	UserID string `json:"user_id"`
}

// DeleteUser 删除用户
func DeleteUser(req *DeleteUserRequest) (err error) {
	return postJSON(userDeleteURL, req, nil)
}

// AddGroupRequest 创建用户组的请求
type AddGroupRequest struct {
	// 用户组id，标识一组用户（由数字、字母、下划线组成），长度限制48B。
	GroupID string `json:"group_id"`
}

// AddGroup 创建用户组
func AddGroup(req *AddGroupRequest) error {
	return postJSON(groupAddURL, req, nil)
}

// DeleteGroupRequest 删除用户组
type DeleteGroupRequest struct {
	// 用户组id，标识一组用户（由数字、字母、下划线组成），长度限制48B。
	GroupID string `json:"group_id"`
}

// DeleteGroup 删除用户组
func DeleteGroup(req *DeleteGroupRequest) error {
	return postJSON(groupDeleteURL, req, nil)
}

// GetGroupListRequest 组列表查询的请求参数
type GetGroupListRequest struct {
	// 默认值0，起始序号
	Start uint32 `json:"start"`
	// 返回数量，默认值100，最大值1000
	Length uint32 `json:"length"`
}

// GetGroupListResult 查询组列表的响应结果
type GetGroupListResult struct {
	// 组id列表
	GroupIDList []string `json:"group_id_list"`
}

// GetGroupList 组列表查询
func GetGroupList(req *GetGroupListRequest) (res *GetGroupListResult, err error) {
	err = postJSON(groupGetlistURL, req, &res)
	return
}
