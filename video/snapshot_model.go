package video

// 视频快照相关响应模型。

// GetVideoSnapshotResult 是视频快照（web 端）的响应。
type GetVideoSnapshotResult struct {
	Pvdata     string         `json:"pvdata"`      // bin格式截取时间表url
	ImgXLen    int            `json:"img_x_len"`   // 每行图片数。一般为10
	ImgYLen    int            `json:"img_y_len"`   // 每列图片数。一般为10
	ImgXSize   int            `json:"img_x_size"`  // 每张图片长。一般为160
	ImgYSize   int            `json:"img_y_size"`  // 每张图片宽。一般为90
	Image      []string       `json:"image"`       // 图片拼版url列表。第一张拼版占满时延续第二张
	Index      []int          `json:"index"`       // json数组格式截取时间表。首项恒为0，其后依次为各图片的截取时间（秒）。请求时 index=1 才有值
	VideoShots map[string]any `json:"video_shots"` // 作用尚不明确。实测为空对象
	Indexs     map[string]any `json:"indexs"`      // 作用尚不明确。实测为空对象
}
