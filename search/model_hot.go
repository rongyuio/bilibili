package search

// 默认搜索与热搜相关响应模型。

// GetSearchDefaultResult 是默认搜索内容，即搜索框默认填充的内容。
type GetSearchDefaultResult struct {
	Seid      string `json:"seid"`       // 搜索 seid
	ID        int64  `json:"id"`         // 默认搜索 id
	Type      int    `json:"type"`       // 类型。实测为 0
	ShowName  string `json:"show_name"`  // 显示文字
	Name      string `json:"name"`       // 名称。实测为空
	GotoType  int    `json:"goto_type"`  // 跳转类型。1：视频
	GotoValue string `json:"goto_value"` // 搜索目标 id。视频：稿件 avid
	URL       string `json:"url"`        // 搜索目标跳转 url
}

// GetSearchHotListResult 是热搜列表（wbi/search/square）的 data 对象。
type GetSearchHotListResult struct {
	Trending GetSearchHotListTrending `json:"trending"` // 热搜榜单
}

// GetSearchHotListTrending 是热搜榜单。
type GetSearchHotListTrending struct {
	Title   string                 `json:"title"`    // 标题
	TrackID string                 `json:"trackid"`  // 跟踪 ID
	List    []GetSearchHotListItem `json:"list"`     // 热搜列表，按名次顺序
	TopList []GetSearchHotListItem `json:"top_list"` // 置顶热搜列表。实测为空
}

// GetSearchHotListItem 是热搜列表中的一项。
type GetSearchHotListItem struct {
	Keyword  string `json:"keyword"`   // 关键词
	ShowName string `json:"show_name"` // 显示文字
	Icon     string `json:"icon"`      // 图标 URL
	URI      string `json:"uri"`       // 跳转 uri。实测为空
	Goto     string `json:"goto"`      // 跳转类型。实测为空
}

// GetSearchHotListWebResult 是热搜列表（web端，s.search.bilibili.com/main/hotword）。
// 该接口没有 data 信封，业务字段直接位于根对象上。
type GetSearchHotListWebResult struct {
	ExpStr         string                    `json:"exp_str"`          // 作用尚不明确
	Cost           GetSearchHotListWebCost   `json:"cost"`             // 详细搜索用时
	Seid           string                    `json:"seid"`             // 搜索 seid
	Timestamp      int64                     `json:"timestamp"`        // 榜单统计时间。Unix 秒级时间戳
	List           []GetSearchHotListWebItem `json:"list"`             // 热搜列表，按名次顺序，共 10 项
	HotwordEggInfo string                    `json:"hotword_egg_info"` // 热搜彩蛋信息。实测为 "0"
}

// GetSearchHotListWebCost 是热搜列表（web端）的搜索耗时明细。
type GetSearchHotListWebCost struct {
	ReasRequest         string `json:"reas_request"`         // 请求耗时
	ParamsCheck         string `json:"params_check"`         // 参数检查耗时
	ReasResponseFormat  string `json:"reas_response_format"` // 响应格式化耗时
	DeserializeResponse string `json:"deserialize_response"` // 反序列化耗时
	ReasRequestFormat   string `json:"reas_request_format"`  // 请求格式化耗时
	Total               string `json:"total"`                // 总耗时
	MainHandler         string `json:"main_handler"`         // 主处理耗时
}

// GetSearchHotListWebItem 是热搜列表（web端）中的一项。
type GetSearchHotListWebItem struct {
	Status     string                       `json:"status"`      // 状态。实测为空
	HotID      int64                        `json:"hot_id"`      // 热词 id
	Keyword    string                       `json:"keyword"`     // 关键词
	ResourceID int                          `json:"resource_id"` // 资源 id
	GotoType   int                          `json:"goto_type"`   // 跳转类型。实测为 0
	Res        []any                        `json:"res"`         // 跳转资源。实测为空数组
	ShowName   string                       `json:"show_name"`   // 完整关键词
	Pos        int                          `json:"pos"`         // 名次。1-10
	WordType   int                          `json:"word_type"`   // 条目属性。4：新。5：热。6：雾/咒/小丑（愚人节）。7：直播中。8：默认（无标签）。9：梗。11：话题。12：独家
	ID         int                          `json:"id"`          // 名次。1-10
	GotoValue  string                       `json:"goto_value"`  // 跳转值。实测为空
	LiveID     []any                        `json:"live_id"`     // 直播 id 列表。实测为空数组
	NameType   string                       `json:"name_type"`   // 名称类型。实测为空
	Icon       string                       `json:"icon"`        // 图标 url
	CallReason int                          `json:"call_reason"` // 上榜原因。实测为 1 / 2 / 3
	HeatLayer  string                       `json:"heat_layer"`  // 热度层级。实测为 B / A
	SetGray    int                          `json:"set_gray"`    // 是否置灰。实测为 0
	CardValues []any                        `json:"card_values"` // 卡片数据。实测为空数组
	HeatScore  int64                        `json:"heat_score"`  // 热度分值。实测为 7482526 等
	StatDatas  GetSearchHotListWebStatDatas `json:"stat_datas"`  // 统计信息
}

// GetSearchHotListWebStatDatas 是热搜列表（web端）的统计信息。
type GetSearchHotListWebStatDatas struct {
	CnyFlag      string `json:"cny_flag"`      // 是否为春节活动。布尔型转换的字符串，实测为 "0"
	Etime        string `json:"etime"`         // 统计结束时间。Unix 秒级时间戳
	Stime        string `json:"stime"`         // 统计开始时间。Unix 秒级时间戳
	IsCommercial string `json:"is_commercial"` // 是否为商业推广。布尔型转换的字符串，实测为 "0"
}
