package face

// SearchRequest 人脸搜索1:N的请求参数
type SearchRequest struct {
	// 图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断。 两张图片通过json格式上传
	Image string `json:"image"`
	// 图片类型
	// BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M
	// URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
	// FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
	ImageType string `json:"image_type"`
	// 从指定的group中进行查找 用逗号分隔，上限10个
	GroupIDList string `json:"group_id_list"`
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
	// 当需要对特定用户进行比对时，指定user_id进行比对。即人脸认证功能。1:1时使用
	UserID string `json:"user_id,omitempty"`
	// 查找后返回的用户数量。返回相似度最高的几个用户，默认为1，最多返回50个。
	MaxUserNum uint32 `json:"max_user_num,omitempty"`
}

// NewSearchRequest 新建人脸搜索M:N的请求
func NewSearchRequest(image, imageType, groupIDList string) *SearchRequest {
	return &SearchRequest{
		Image:       image,
		ImageType:   imageType,
		GroupIDList: groupIDList,
	}
}

// SearchResult 人脸搜索响应参数
type SearchResult struct {
	// 人脸标志
	FaceToken string `json:"face_token"`
	// 匹配的用户信息列表
	UserList []*SearchItem `json:"user_list"`
}

// SearchItem 人脸搜索的响应项
type SearchItem struct {
	// 用户所属的group_id
	GroupID string `json:"group_id"`
	// 用户的user_id
	UserID string `json:"user_id"`
	// 注册用户时携带的user_info
	UserInfo string `json:"user_info"`
	// 用户的匹配得分，推荐阈值80分
	Score float32 `json:"score"`
}

// Search 人脸搜索1:N识别
func Search(req *SearchRequest) (res *SearchResult, err error) {
	err = postJSON(searchURL, req, &res)
	return
}

// MultiSearchRequest 人脸搜索M:N的请求参数
type MultiSearchRequest struct {
	// 图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断。 两张图片通过json格式上传
	Image string `json:"image"`
	// 图片类型
	// BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M
	// URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
	// FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
	ImageType string `json:"image_type"`
	// 从指定的group中进行查找 用逗号分隔，上限10个
	GroupIDList string `json:"group_id_list"`
	// 最多处理人脸的数目, M:N时使用
	// 默认值为1(仅检测图片中面积最大的那个人脸) 最大值10
	MaxFaceNum int `json:"max_face_num,omitempty"`
	// 匹配阈值（设置阈值后，score低于此阈值的用户信息将不会返回） 最大100 最小0 默认80
	// 此阈值设置得越高，检索速度将会越快，推荐使用默认阈值80, M:N时使用
	MatchThreshold int `json:"match_threshold,omitempty"`
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
	// 查找后返回的用户数量。返回相似度最高的几个用户，默认为1，最多返回50个。
	MaxUserNum uint32 `json:"max_user_num,omitempty"`
}

// NewMultiSearchRequest 新建人脸搜索M:N的请求
func NewMultiSearchRequest(image, imageType, groupIDList string) *MultiSearchRequest {
	return &MultiSearchRequest{
		Image:       image,
		ImageType:   imageType,
		GroupIDList: groupIDList,
	}
}

// MultiSearchResult 人脸搜索M:N的响应
type MultiSearchResult struct {
	// 图片中的人脸数量
	FaceNum int `json:"face_num"`
	// 人脸信息列表
	FaceList []*MultiSearchItem `json:"face_list"`
}

// MultiSearchItem 人脸搜索M:N的响应列表项
type MultiSearchItem struct {
	// 人脸标志
	FaceToken string `json:"face_token"`
	// Location 人脸在图片中的位置
	Location *Location `json:"location"`
	// 匹配的用户信息列表
	UserList []*SearchItem `json:"user_list"`
}

// MultiSearch 人脸搜索M:N识别
func MultiSearch(req *MultiSearchRequest) (res *MultiSearchResult, err error) {
	err = postJSON(multiSearchURL, req, &res)
	return
}
