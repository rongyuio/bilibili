package article

// GetArticleInfo 接口的响应模型。
//
// 本文件的类型只服务于 GetArticleInfo，不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// GetArticleInfoResult 是「获取专栏文章基本信息」的 data 对象。
type GetArticleInfoResult struct {
	Like            int                          `json:"like"`              // 是否点赞。0：未点赞。1：已点赞。需要登录(Cookie) 。未登录为0
	Attention       bool                         `json:"attention"`         // 是否关注文章作者。false：未关注。true：已关注。需要登录(Cookie) 。未登录为false
	Favorite        bool                         `json:"favorite"`          // 是否收藏。false：未收藏。true：已收藏。需要登录(Cookie) 。未登录为false
	Coin            int                          `json:"coin"`              // 为文章投币数
	Stats           GetArticleInfoStats          `json:"stats"`             // 状态数信息
	Title           string                       `json:"title"`             // 文章标题
	BannerURL       string                       `json:"banner_url"`        // 文章头图url
	Mid             int                          `json:"mid"`               // 文章作者mid
	AuthorName      string                       `json:"author_name"`       // 文章作者昵称
	IsAuthor        bool                         `json:"is_author"`         // true。作用尚不明确
	ImageURLs       []string                     `json:"image_urls"`        // 动态封面
	OriginImageURLs []string                     `json:"origin_image_urls"` // 封面图片
	Shareable       bool                         `json:"shareable"`         // true。作用尚不明确
	ShowLaterWatch  bool                         `json:"show_later_watch"`  // true。作用尚不明确
	ShowSmallWindow bool                         `json:"show_small_window"` // true。作用尚不明确
	InList          bool                         `json:"in_list"`           // 是否收于文集。false：否。true：是
	Pre             int                          `json:"pre"`               // 上一篇文章cvid。无为0
	Next            int                          `json:"next"`              // 下一篇文章cvid。无为0
	ShareChannels   []GetArticleInfoShareChannel `json:"share_channels"`    // 分享方式列表
	Type            int                          `json:"type"`              // 文章类别。0：文章。2：笔记
	DisableShare    bool                         `json:"disable_share"`     // 是否禁止分享。实测为 false
}

// GetArticleInfoStats 是文章的状态数信息。
type GetArticleInfoStats struct {
	View     int `json:"view"`     // 阅读数
	Favorite int `json:"favorite"` // 收藏数
	Like     int `json:"like"`     // 点赞数
	Dislike  int `json:"dislike"`  // 点踩数
	Reply    int `json:"reply"`    // 评论数
	Share    int `json:"share"`    // 分享数
	Coin     int `json:"coin"`     // 投币数
	Dynamic  int `json:"dynamic"`  // 动态转发数
}

// GetArticleInfoShareChannel 是文章的一种分享方式。
type GetArticleInfoShareChannel struct {
	Name         string `json:"name"`          // 分享名称
	Picture      string `json:"picture"`       // 分享图片url
	ShareChannel string `json:"share_channel"` // 分享代号
}
