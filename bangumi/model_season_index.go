package bangumi

// GetBangumiSeasonIndex 接口的响应模型。
//
// 本文件的类型只服务于 GetBangumiSeasonIndex，不与其它接口共用。

// 剧集索引相关响应模型。

// GetBangumiSeasonIndexResult 获取剧集索引的结果。
type GetBangumiSeasonIndexResult struct {
	HasNext int                         `json:"has_next"` // 是否有下一页。1：有，0：无
	List    []GetBangumiSeasonIndexItem `json:"list"`     // 剧集列表
	Num     int                         `json:"num"`      // 当前页码
	Size    int                         `json:"size"`     // 每页数量
	Total   int                         `json:"total"`    // 总剧集数
}

// GetBangumiSeasonIndexItem 剧集列表中的剧集信息。
type GetBangumiSeasonIndexItem struct {
	Badge        string                         `json:"badge"`         // 角标文本。例如「大会员」「独家」
	BadgeInfo    GetBangumiSeasonIndexBadgeInfo `json:"badge_info"`    // 角标信息
	BadgeType    int                            `json:"badge_type"`    // 角标类型
	Cover        string                         `json:"cover"`         // 封面图片 URL
	FirstEp      GetBangumiSeasonIndexFirstEp   `json:"first_ep"`      // 第一话信息
	IndexShow    string                         `json:"index_show"`    // 剧集进度显示。例如「全48话」
	IsFinish     int                            `json:"is_finish"`     // 是否完结。1：完结，0：未完结
	Link         string                         `json:"link"`          // 剧集播放页链接
	MediaID      int                            `json:"media_id"`      // 剧集 md id
	Order        string                         `json:"order"`         // 排序显示文本。例如「9.9分」
	OrderType    string                         `json:"order_type"`    // 排序类型
	Score        string                         `json:"score"`         // 评分
	SeasonID     int                            `json:"season_id"`     // 剧集 ss id
	SeasonStatus int                            `json:"season_status"` // 剧集状态。普通剧集为 2，大会员或独家等有角标的剧集为 13
	SeasonType   int                            `json:"season_type"`   // 剧集类别。值同参数中的 season_type
	SubTitle     string                         `json:"subTitle"`      // 副标题
	Title        string                         `json:"title"`         // 标题
	TitleIcon    string                         `json:"title_icon"`    // 标题图标 URL
}

// GetBangumiSeasonIndexBadgeInfo 剧集列表项中的角标信息。
type GetBangumiSeasonIndexBadgeInfo struct {
	BgColor      string `json:"bg_color"`       // 角标背景色。日间模式
	BgColorNight string `json:"bg_color_night"` // 角标背景色。夜间模式
	Text         string `json:"text"`           // 角标文本
}

// GetBangumiSeasonIndexFirstEp 剧集列表项中的第一话信息。
type GetBangumiSeasonIndexFirstEp struct {
	Cover string `json:"cover"` // 第一话封面 URL
	EpID  int    `json:"ep_id"` // 第一话 ep id
}
