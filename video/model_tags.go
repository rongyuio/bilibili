package video

// GetVideoTags 接口的响应模型。

// GetVideoTagsItem 是「获取视频TAG」data 数组中的对象。
type GetVideoTagsItem struct {
	TagID        int                   `json:"tag_id"`        // tag_id
	TagName      string                `json:"tag_name"`      // TAG名称
	Cover        string                `json:"cover"`         // TAG图片url
	HeadCover    string                `json:"head_cover"`    // TAG页面头图url
	Content      string                `json:"content"`       // TAG介绍
	ShortContent string                `json:"short_content"` // TAG简介
	Type         int                   `json:"type"`          // ？？？
	State        int                   `json:"state"`         // 0
	Ctime        int                   `json:"ctime"`         // 创建时间。时间戳
	Count        GetVideoTagsItemCount `json:"count"`         // 状态数
	IsAtten      int                   `json:"is_atten"`      // 是否关注。需要登录(Cookie)。未登录为0
	Likes        int                   `json:"likes"`         // 0。作用尚不明确
	Hates        int                   `json:"hates"`         // 0。作用尚不明确
	Attribute    int                   `json:"attribute"`     // 0。作用尚不明确
	Liked        int                   `json:"liked"`         // 是否已经点赞。需要登录(Cookie)。未登录为0
	Hated        int                   `json:"hated"`         // 是否已经点踩。需要登录(Cookie)。未登录为0
	ExtraAttr    int                   `json:"extra_attr"`    // ? ? ?
}

// GetVideoTagsItemCount 是「获取视频TAG」count 对象。
type GetVideoTagsItemCount struct {
	View  int `json:"view"`  // 0。作用尚不明确
	Use   int `json:"use"`   // 视频添加TAG数
	Atten int `json:"atten"` // TAG关注
}
