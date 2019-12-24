package face

import "fmt"

// MatchRequest 人脸对比请求参数
type MatchRequest struct {
	// 图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断。 两张图片通过json格式上传
	Image string `json:"image"`
	// 图片类型
	// BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M
	// URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
	// FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
	ImageType string `json:"image_type"`
	//  人脸的类型
	// LIVE：表示生活照：通常为手机、相机拍摄的人像图片、或从网络获取的人像图片等，
	// IDCARD：表示身份证芯片照：二代身份证内置芯片中的人像照片，
	// WATERMARK：表示带水印证件照：一般为带水印的小图，如公安网小图
	// CERT：表示证件照片：如拍摄的身份证、工卡、护照、学生证等证件图片
	// 默认LIVE
	FaceType string `json:"face_type,omitempty"`
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
}

// NewMatchRequest 新建对比请求
func NewMatchRequest(image, imageType string) *MatchRequest {
	return &MatchRequest{
		Image:     image,
		ImageType: imageType,
	}
}

// MatchResponse 人脸对比响应数据
type MatchResponse struct {
	// 错误码
	ErrorCode int `json:"error_code"`
	// 错误描述信息
	ErrorMsg string `json:"error_msg"`
	// 人脸相似度得分，推荐阈值80分
	Score float32 `json:"score"`
	// 人脸信息列表
	FaceList []*MatchItem `json:"face_list"`
}

// Code 返回错误码
func (a *MatchResponse) Code() int {
	return a.ErrorCode
}

// Message 返回错误信息
func (a *MatchResponse) Message() string {
	return a.ErrorMsg
}

// Error 实现error接口
func (a *MatchResponse) Error() string {
	return fmt.Sprintf("error_code: %d, error_msg: %s", a.ErrorCode, a.ErrorMsg)
}

// MatchItem 人脸对比的项目
type MatchItem struct {
	// 人脸的唯一标志
	FaceToken string `json:"face_token"`
}

// Match 人脸对比
func Match(req []*MatchRequest) (res *MatchResponse, err error) {
	err = postJSON(matchURL, req, res)
	return
}
