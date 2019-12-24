package face

// 人脸识别

const (
	// 人脸检测
	detectURL = `https://aip.baidubce.com/rest/2.0/face/v3/detect`
	// 人脸对比
	matchURL = `https://aip.baidubce.com/rest/2.0/face/v3/match`
	// 人脸搜索
	searchURL = `https://aip.baidubce.com/rest/2.0/face/v3/search`
	// 人脸库管理-人脸注册
	userAddURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add`
	// 人脸库管理-人脸更新
	userUpdateURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/update`
	// 人脸库管理-删除用户
	userDeleteURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/delete`
	// 人脸库管理-用户信息查询
	userGetURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/get`
	// 人脸库管理-获取组列表
	groupGetlistURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getlist`
	// 人脸库管理-获取用户人脸列表
	faceGetlistURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/getlist`
	// 人脸库管理-获取用户列表
	groupGetusersURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getusers`
	// 人脸库管理-复制用户
	userCopyURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/copy`
	// 人脸库管理-创建用户组
	groupAddURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/add`
	// 人脸库管理-删除用户组
	groupDeleteURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/delete`
	// ⼈脸库管理-删除人脸
	faceDeleteURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/delete`
	// 在线活体检测
	onlineVerfiyURL = `https://aip.baidubce.com/rest/2.0/face/v3/faceverify`
	// 人脸搜索-M:N识别
	multiSearchURL = `https://aip.baidubce.com/rest/2.0/face/v3/multi-search`
)

const (
	// jsonContentType 请求json的内容类型
	jsonContentTypeOfRequest = `application/json; charset=utf-8`
	// jsonContentTypeResponse 响应的json内容类型
	jsonContentTypeResponse = `json`
)
