package bilibili

// 分区视频排行相关响应模型。

type ZoneVideoRankList struct {
	Note string      `json:"note"` // “根据稿件内容质量、近期的数据综合展示，动态更新”
	List []VideoInfo `json:"list"` // 视频列表
}

type ZoneVideoListInfo struct {
	Archives []VideoInfo   `json:"archives"` // 视频列表
	Page     ZoneVideoPage `json:"page"`     // 页面信息
}

// ZoneVideoPage 分区视频分页信息。
//
// 字段名为 num/size，与空间投稿的 pn/ps、合集的 page_num/page_size 不同，
// 因此保留独立类型，不做跨接口合并。
type ZoneVideoPage struct {
	Count int `json:"count"` // 总计视频数
	Num   int `json:"num"`   // 当前页码
	Size  int `json:"size"`  // 每页项数
}

type ZoneVideoRankInfo struct {
	ExpList        *string         `json:"exp_list"`         // 作用尚不明确
	ShowModuleList []string        `json:"show_module_list"` // 显示模块列表?
	Result         []RankVideoInfo `json:"result"`           // 结果本体。失败时为null
	ShowColumn     int             `json:"show_column"`      // 0。作用尚不明确
	RqtType        string          `json:"rqt_type"`         // search。作用尚不明确
	Numpages       int             `json:"numPages"`         // 页码。失败时为0
	Numresults     int             `json:"numResults"`       // 视频数。失败时为0
	CrrQuery       *string         `json:"crr_query"`        // 空。作用尚不明确
	Pagesize       int             `json:"pagesize"`         // 视频数
	SuggestKeyword *string         `json:"suggest_keyword"`  // 空。作用尚不明确
	EggInfo        *string         `json:"egg_info"`         // 作用尚不明确
	Cache          int             `json:"cache"`            // 0。作用尚不明确
	ExpBits        int             `json:"exp_bits"`         // 1。作用尚不明确
	ExpStr         *string         `json:"exp_str"`          // 空。作用尚不明确
	Seid           string          `json:"seid"`             // 一串字符串中的数字。作用尚不明确
	Msg            string          `json:"msg"`              // 结果信息。成功时为success, 反之为as error.
	EggHit         int             `json:"egg_hit"`          // 0。作用尚不明确
	Page           int             `json:"page"`             // 页码
}

type RankVideoInfo struct {
	Pubdate      string `json:"pubdate"`        // 发布时间。格式为 yyyy-MM-dd HH:mm:ss
	Pic          string `json:"pic"`            // 封面图
	Tag          string `json:"tag"`            // 标签。用 , 分隔
	Duration     int    `json:"duration"`       // 时长。单位为秒
	Id           int    `json:"id"`             // aid
	RankScore    int    `json:"rank_score"`     // 排序分数?
	Badgepay     bool   `json:"badgepay"`       // 是否有角标?
	Senddate     int    `json:"senddate"`       // 发送时间?。UNIX 秒级时间戳
	Author       string `json:"author"`         // UP主名
	Review       int    `json:"review"`         // 评论数
	Mid          int    `json:"mid"`            // UP主mid
	IsUnionVideo int    `json:"is_union_video"` // 是否为联合投稿
	RankIndex    int    `json:"rank_index"`     // 排序索引号
	Type         string `json:"type"`           // 类型。video: 视频
	Arcrank      string `json:"arcrank"`        // 0。作用尚不明确
	Play         string `json:"play"`           // 播放数
	RankOffset   int    `json:"rank_offset"`    // 排序偏移?。与 rank_index 相同
	Description  string `json:"description"`    // 简介
	VideoReview  int    `json:"video_review"`   // 弹幕数?
	IsPay        int    `json:"is_pay"`         // 是否付费?。0: 免费。1: 付费
	Favorites    int    `json:"favorites"`      // 收藏数
	Arcurl       string `json:"arcurl"`         // 视频播放页URL
	Bvid         string `json:"bvid"`           // bvid
	Title        string `json:"title"`          // 标题
	Vt           int    `json:"vt"`             // 0。作用尚不明确
	EnableVt     int    `json:"enable_vt"`      // 0。作用尚不明确
	VtDisplay    string `json:"vt_display"`     // 空。作用尚不明确
}
