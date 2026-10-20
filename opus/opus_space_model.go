package opus

// 空间图文接口的响应模型。
//
// 与详情接口一样，本文件独立定义全部对象与数组元素，不复用其他接口的响应体。

// GetUserSpaceOpusResult 是获取用户空间图文接口的 data。
type GetUserSpaceOpusResult struct {
	HasMore   bool                   `json:"has_more"`   // 是否还有更多
	Items     []GetUserSpaceOpusItem `json:"items"`      // 信息本体
	Offset    string                 `json:"offset"`     // 偏移信息。下次翻页请求时传入的 offset（同时也是最后一条的 opus id）
	UpdateNum int                    `json:"update_num"` // 更新数
}

// GetUserSpaceOpusItem 是空间图文中的一条内容。
type GetUserSpaceOpusItem struct {
	Content string                 `json:"content"`  // 文本内容
	Cover   *GetUserSpaceOpusCover `json:"cover"`    // 封面信息
	JumpURL string                 `json:"jump_url"` // 跳转 URL
	OpusID  string                 `json:"opus_id"`  // opus id
	Stat    GetUserSpaceOpusStat   `json:"stat"`     // 统计信息
}

// GetUserSpaceOpusCover 是图文封面的图片信息。
type GetUserSpaceOpusCover struct {
	Height int    `json:"height"` // 高度
	URL    string `json:"url"`    // 图片 URL
	Width  int    `json:"width"`  // 宽度
}

// GetUserSpaceOpusStat 是图文的统计信息。
type GetUserSpaceOpusStat struct {
	Like string `json:"like"` // 点赞数，接口返回字符串
	View string `json:"view"` // 浏览数，仅获取登录状态下自己的图文时存在，同样是字符串
}
