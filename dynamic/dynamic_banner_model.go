package dynamic

// 动态首页公告栏接口响应模型。
//
// 本包不复用根包或其他子包的响应体：结果与所有嵌套对象都以接口方法名
// 为前缀独立定义，便于后续单个接口的字段变更不影响其他调用。

// GetDynamicBannerResult 是动态首页公告栏（新版）的响应数据。
type GetDynamicBannerResult struct {
	Banners []GetDynamicBannerItem `json:"banners"` // 横幅列表。套了个娃
}

// GetDynamicBannerItem 是动态首页公告栏的一条横幅。
type GetDynamicBannerItem struct {
	BannerID  int64  `json:"banner_id"`  // 横幅 id。旧版接口无
	EndTime   int64  `json:"end_time"`   // 结束时间。UNIX 秒级时间戳
	ImgURL    string `json:"img_url"`    // 图片 URL
	Link      string `json:"link"`       // 跳转链接
	Platform  int    `json:"platform"`   // 平台。与请求参数对应
	Position  string `json:"position"`   // 位置。与请求参数对应
	StartTime int64  `json:"start_time"` // 开始时间。UNIX 秒级时间戳
	Title     string `json:"title"`      // 标题
	Weight    int    `json:"weight"`     // 权重
}

// GetDynamicBannerLegacyResult 是动态首页公告栏（旧版）的响应数据。
type GetDynamicBannerLegacyResult struct {
	Banners []GetDynamicBannerLegacyItem `json:"banners"` // 横幅列表。套了个娃
}

// GetDynamicBannerLegacyItem 是旧版公告栏的一条横幅。与新版相比没有 banner_id。
type GetDynamicBannerLegacyItem struct {
	EndTime   int64  `json:"end_time"`   // 结束时间。UNIX 秒级时间戳
	ImgURL    string `json:"img_url"`    // 图片 URL
	Link      string `json:"link"`       // 跳转链接
	Platform  int    `json:"platform"`   // 平台。与请求参数对应
	Position  string `json:"position"`   // 位置。与请求参数对应
	StartTime int64  `json:"start_time"` // 开始时间。UNIX 秒级时间戳
	Title     string `json:"title"`      // 标题
	Weight    int    `json:"weight"`     // 权重
}
