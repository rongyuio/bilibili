package video

// GetVideoPageList 接口的响应模型。

// GetVideoPageListItem 是「获取视频分P列表」data 数组中的对象。
type GetVideoPageListItem struct {
	Cid        int                           `json:"cid"`         // 当前分P cid
	Page       int                           `json:"page"`        // 当前分P
	From       string                        `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                        `json:"part"`        // 当前分P标题
	Duration   int                           `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                        `json:"vid"`         // 站外视频vid
	Weblink    string                        `json:"weblink"`     // 站外视频跳转url
	Dimension  GetVideoPageListItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                        `json:"first_frame"` // 分P封面
}

// GetVideoPageListItemDimension 是「获取视频分P列表」dimension 对象。
type GetVideoPageListItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}
