package face

import "fmt"

// DetectResponse 人脸检查响应参数
type DetectResponse struct {
	// 错误码
	ErrorCode int `json:"error_code"`
	// 错误描述信息
	ErrorMsg string `json:"error_msg"`
	// FaceNum 检测到的人脸数量
	FaceNum int `json:"face_num"`
	// FaceList 人脸信息列表
	FaceList []*DetectItem `json:"face_list"`
}

// Code 返回错误码
func (a *DetectResponse) Code() int {
	return a.ErrorCode
}

// Message 返回错误信息
func (a *DetectResponse) Message() string {
	return a.ErrorMsg
}

// Error 实现error接口
func (a *DetectResponse) Error() string {
	return fmt.Sprintf("error_code: %d, error_msg: %s", a.ErrorCode, a.ErrorMsg)
}

// DetectItem 人脸信息
type DetectItem struct {
	// 人脸图片的唯一标识
	FaceToken string `json:"face_token"`
	// 人脸在图片中的位置
	Location *Location `json:"location"`
	// 人脸置信度，范围[0-1]，代表一张人脸的概率，0最小，1最大
	FaceProbability float64 `json:"face_probability"`
	// 人脸旋转角度参数
	Angle *DetectAngle `json:"angle"`
	// 年龄 ，当face_field包含age时返回
	Age float64 `json:"age"`
	// 美丑打分，范围0-100，越大表示越美。当face_fields包含beauty时返回
	Beauty int64 `json:"beauty"`
	// 表情，当 face_field包含expression时返回
	Expression *DetectExtension `json:"expression"`
	// 脸型，当face_field包含face_shape时返回
	FaceShape *DetectExtension `json:"face_shape"`
	// 性别，face_field包含gender时返回
	Gender *DetectExtension `json:"gender"`
	// 是否带眼镜，face_field包含glasses时返回
	Glasses *DetectExtension `json:"glasses"`
	// 双眼状态（睁开/闭合） face_field包含eye_status时返回
	EyeStatus *DetectEyeStatus `json:"eyestatus"`
	// 情绪 face_field包含emotion时返回
	Emotion *DetectExtension `json:"emotion"`
	// 真实人脸/卡通人脸 face_field包含face_type时返回
	FaceType *DetectExtension `json:"FaceType"`
	// 4个关键点位置，左眼中心、右眼中心、鼻尖、嘴中心。face_field包含landmark时返回
	LandMark []*DetectLandMark `json:"landmark"`
	// 72个特征点位置 face_field包含landmark150时返回
	LandMark72 []*DetectLandMark `json:"landmark72"`
	// 150个特征点位置 face_field包含landmark150时返回
	LandMark150 []*DetectLandMark `json:"landmark150"`
	// 人脸质量信息。face_field包含quality时返回
	Quality *DetectQuality `json:"quality"`
}

// Location 人脸在图片中的位置
type Location struct {
	// 人脸距离左边界的距离
	Left float64 `json:"left"`
	// 人脸距离上边界的距离
	Top float64 `json:"top"`
	// 人脸区域的宽度
	Width float64 `json:"width"`
	// 人脸区域的高度
	Height float64 `json:"height"`
	// 人脸框相对于竖直方向的顺时针旋转角,[-180,180]
	Rotation int64 `json:"rotation"`
}

// DetectAngle 人脸旋转角度参数
type DetectAngle struct {
	// 三维旋转之左右旋转角[-90(左), 90(右)]
	Yaw float64 `json:"yaw"`
	// 三维旋转之俯仰角度[-90(上), 90(下)]
	Pitch float64 `json:"pitch"`
	// 平面内旋转角[-180(逆时针), 180(顺时针)]
	Roll float64 `json:"roll"`
}

// DetectLandMark 特征点
type DetectLandMark struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// DetectExtension 表情
type DetectExtension struct {
	// 分类
	Type string `json:"type"`
	// 值
	Probability float64 `json:"probability"`
}

// DetectEyeStatus 双眼状态
type DetectEyeStatus struct {
	// 左眼状态 [0,1]取值，越接近0闭合的可能性越大
	LeftEye float64 `json:"left_eye"`
	// 右眼状态 [0,1]取值，越接近0闭合的可能性越大
	RightEye float64 `json:"right_eye"`
}

// DetectQuality 人脸质量信息
type DetectQuality struct {
	// 人脸各部分遮挡的概率，范围[0~1]，0表示完整，1表示不完整
	Occlusion *DetectOcclusion `json:"occlusion"`
	// 人脸模糊程度，范围[0~1]，0表示清晰，1表示模糊
	Blur float64 `json:"blur"`
	// 取值范围在[0~255], 表示脸部区域的光照程度 越大表示光照越好
	Illumination float64 `json:"illumination"`
	// 人脸完整度，0或1, 0为人脸溢出图像边界，1为人脸都在图像边界内
	Completeness int64 `json:"completeness"`
}

// DetectOcclusion 人脸各部分遮挡的概率
type DetectOcclusion struct {
	// 左眼遮挡比例，[0-1] ，1表示完全遮挡
	LeftEye float64 `json:"left_eye"`
	// 右眼遮挡比例，[0-1] ， 1表示完全遮挡
	RightEye float64 `json:"right_eye"`
	// 鼻子遮挡比例，[0-1] ， 1表示完全遮挡
	Nose float64 `json:"nose"`
	// 嘴巴遮挡比例，[0-1] ， 1表示完全遮挡
	Mouth float64 `json:"mouth"`
	// 左脸颊遮挡比例，[0-1] ， 1表示完全遮挡
	LeftCheek float64 `json:"left_cheek"`
	// 右脸颊遮挡比例，[0-1] ， 1表示完全遮挡
	RightCheek float64 `json:"right_cheek"`
	// 下巴遮挡比例，，[0-1] ， 1表示完全遮挡
	Chin float64 `json:"chin"`
}

// DetectRequest 人脸检查请求参数
type DetectRequest struct {
	// 图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断
	Image string `json:"image"`
	// 图片类型
	// BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M
	// URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
	// FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
	ImageType string `json:"image_type"`
	// 包括age,beauty,expression,face_shape,gender,glasses,landmark,landmark150,race,quality,eye_status,emotion,face_type信息,逗号分隔. 默认只返回face_token、人脸框、概率和旋转角度
	FaceField string `json:"face_field,omitempty"`
	// 最多处理人脸的数目，默认值为1，仅检测图片中面积最大的那个人脸；最大值10，检测图片中面积最大的几张人脸。
	MaxFaceNum uint32 `json:"max_face_num,omitempty"`
	// 人脸的类型
	// LIVE表示生活照：通常为手机、相机拍摄的人像图片、或从网络获取的人像图片等
	// IDCARD表示身份证芯片照：二代身份证内置芯片中的人像照片
	// WATERMARK表示带水印证件照：一般为带水印的小图，如公安网小图
	// CERT表示证件照片：如拍摄的身份证、工卡、护照、学生证等证件图片
	// 默认LIVE
	FaceType string `json:"face_type,omitempty"`
	// 活体控制 检测结果中不符合要求的人脸会被过滤
	// NONE: 不进行控制
	// LOW:较低的活体要求(高通过率 低攻击拒绝率)
	// NORMAL: 一般的活体要求(平衡的攻击拒绝率, 通过率)
	// HIGH: 较高的活体要求(高攻击拒绝率 低通过率)
	// 默认NONE
	LivenessControl string `json:"liveness_control,omitempty"`
}

// NewDetectRequest 新建人脸检测请求
func NewDetectRequest(image, imageType string) *DetectRequest {
	return &DetectRequest{
		Image:     image,
		ImageType: imageType,
	}
}

// Detect 人脸检测
func Detect(req *DetectRequest) (res *DetectResponse, err error) {
	err = postJSON(detectURL, req, &res)
	return
}
