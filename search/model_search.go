package search

// 搜索相关响应模型。

// IntegratedSearchResult 是综合搜索的 data 对象。
type IntegratedSearchResult struct {
	Seid           string                       `json:"seid"`             // 搜索id
	Page           int                          `json:"page"`             // 页数，固定为1
	PageSize       int                          `json:"page_size"`        // 每页条数，固定为20
	NumResults     int                          `json:"numResults"`       // 总条数，最大值为1000
	NumPages       int                          `json:"numPages"`         // 分页数，最大值为50
	SuggestKeyword string                       `json:"suggest_keyword"`  // 作用尚不明确
	RqtType        string                       `json:"rqt_type"`         // 作用尚不明确
	CostTime       IntegratedSearchCostTime     `json:"cost_time"`        // 详细搜索用时
	ExpList        map[string]any               `json:"exp_list"`         // 作用尚不明确
	EggHit         int                          `json:"egg_hit"`          // 作用尚不明确
	PageInfo       IntegratedSearchPageInfo     `json:"pageinfo"`         // 分类页数信息
	TopTList       IntegratedSearchTopList      `json:"top_tlist"`        // 分类结果数目信息
	ShowColumn     int                          `json:"show_column"`      // 作用尚不明确
	ShowModuleList []string                     `json:"show_module_list"` // 返回结果类型列表
	Result         []IntegratedSearchResultItem `json:"result"`           // 结果列表
}

// IntegratedSearchCostTime 是综合搜索的详细搜索用时。
type IntegratedSearchCostTime struct {
	ParamsCheck      string `json:"params_check"`         // 参数检查时间
	IllegalHandler   string `json:"illegal_handler"`      // 违规处理时间
	AsResponseFormat string `json:"as_response_format"`   // 响应格式化时间
	AsRequest        string `json:"as_request"`           // 请求时间
	SaveCache        string `json:"save_cache"`           // 缓存保存时间
	DeserializeResp  string `json:"deserialize_response"` // 反序列化时间
	AsRequestFormat  string `json:"as_request_format"`    // 请求格式化时间
	Total            string `json:"total"`                // 总耗时
	MainHandler      string `json:"main_handler"`         // 主处理时间
}

// IntegratedSearchPageInfo 是综合搜索的分类页数信息。
type IntegratedSearchPageInfo struct {
	Pgc           IntegratedSearchItemCount `json:"pgc"`            // PGC相关内容数
	LiveRoom      IntegratedSearchItemCount `json:"live_room"`      // 直播间数
	Photo         IntegratedSearchItemCount `json:"photo"`          // 相簿数
	Topic         IntegratedSearchItemCount `json:"topic"`          // 话题数
	Video         IntegratedSearchItemCount `json:"video"`          // 视频数
	User          IntegratedSearchItemCount `json:"user"`           // 用户数
	BiliUser      IntegratedSearchItemCount `json:"bili_user"`      // B站用户数
	MediaFt       IntegratedSearchItemCount `json:"media_ft"`       // 电影数
	Article       IntegratedSearchItemCount `json:"article"`        // 专栏数
	MediaBangumi  IntegratedSearchItemCount `json:"media_bangumi"`  // 番剧数
	Special       IntegratedSearchItemCount `json:"special"`        // 特别项目数
	OperationCard IntegratedSearchItemCount `json:"operation_card"` // 运营卡片数
	UpUser        IntegratedSearchItemCount `json:"upuser"`         // UP主数
	Movie         IntegratedSearchItemCount `json:"movie"`          // 电影数
	LiveAll       IntegratedSearchItemCount `json:"live_all"`       // 直播相关数
	Tv            IntegratedSearchItemCount `json:"tv"`             // 电视数
	Live          IntegratedSearchItemCount `json:"live"`           // 直播间数
	Bangumi       IntegratedSearchItemCount `json:"bangumi"`        // 番剧数
	Activity      IntegratedSearchItemCount `json:"activity"`       // 活动数
	LiveMaster    IntegratedSearchItemCount `json:"live_master"`    // 主播数
	LiveUser      IntegratedSearchItemCount `json:"live_user"`      // 直播用户数
}

// IntegratedSearchItemCount 是综合搜索某一分类的条目计数。
type IntegratedSearchItemCount struct {
	NumResults int `json:"numResults"` // 结果总数
	Total      int `json:"total"`      // 总计数量
	Pages      int `json:"pages"`      // 分页数量
}

// IntegratedSearchTopList 是综合搜索的分类结果数目信息。
type IntegratedSearchTopList struct {
	Pgc           int `json:"pgc"`            // PGC内容数
	LiveRoom      int `json:"live_room"`      // 直播数
	Photo         int `json:"photo"`          // 相簿数
	Topic         int `json:"topic"`          // 话题数
	Video         int `json:"video"`          // 视频数
	User          int `json:"user"`           // 用户数
	BiliUser      int `json:"bili_user"`      // B站用户数
	MediaFt       int `json:"media_ft"`       // 电影数
	Article       int `json:"article"`        // 专栏数
	MediaBangumi  int `json:"media_bangumi"`  // 番剧数
	Card          int `json:"card"`           // 作用不明确
	OperationCard int `json:"operation_card"` // 运营卡片数
	UpUser        int `json:"upuser"`         // UP主数
	Movie         int `json:"movie"`          // 电影数
	LiveAll       int `json:"live_all"`       // 直播相关数
	Tv            int `json:"tv"`             // 电视数
	Live          int `json:"live"`           // 直播间数
	Special       int `json:"special"`        // 特别项目数
	Bangumi       int `json:"bangumi"`        // 番剧数
	Activity      int `json:"activity"`       // 活动数
	LiveMaster    int `json:"live_master"`    // 主播数
	LiveUser      int `json:"live_user"`      // 直播用户数
}

// IntegratedSearchResultItem 是综合搜索结果列表中的一项。
type IntegratedSearchResultItem struct {
	ResultType string           `json:"result_type"` // 结果类型 与result数组对应的项相同
	Data       []map[string]any `json:"data"`        // 具体结果数据 结果为该项所对应的对象条目格式
}
