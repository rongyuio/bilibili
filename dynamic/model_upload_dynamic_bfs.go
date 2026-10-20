package dynamic

// UploadDynamicBfs 的响应模型。

// UploadDynamicBfsResult 是图片动态上传接口的 data 对象。
type UploadDynamicBfsResult struct {
	ImageURL    string `json:"image_url"`    // 图片 url
	ImageWidth  int    `json:"image_width"`  // 图片宽度
	ImageHeight int    `json:"image_height"` // 图片高度
}
