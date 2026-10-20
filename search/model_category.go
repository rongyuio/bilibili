package search

import "encoding/json"

// 分类搜索相关响应模型。

// CategorySearchResult 是分类搜索的 data 对象。
type CategorySearchResult struct {
	Seid           string                  `json:"seid"`            // 搜索seid
	Page           int                     `json:"page"`            // 当前页码
	PageSize       int                     `json:"pagesize"`        // 每页条数，固定20
	NumResults     int                     `json:"numResults"`      // 总条数，最大值为1000
	NumPages       int                     `json:"numPages"`        // 总计分页数，最大值为50
	SuggestKeyword string                  `json:"suggest_keyword"` // 空。作用尚不明确
	RqtType        string                  `json:"rqt_type"`        // search。作用尚不明确
	CostTime       CategorySearchCostTime  `json:"cost_time"`       // 详细搜索用时
	ExpList        map[string]bool         `json:"exp_list"`        // 实验命中情况，键名为实验id，随账号与实验分组变化
	EggHit         int                     `json:"egg_hit"`         // 0。作用尚不明确
	PageInfo       *CategorySearchPageInfo `json:"pageinfo"`        // 副分页信息，只在搜索类型为直播间及主播时有效
	Result         json.RawMessage         `json:"result"`          // 结果列表。搜索类型为直播间及主播时为对象（见 CategorySearchLiveResult），其他类型时为数组（元素见各 *Item 类型）
	ShowColumn     int                     `json:"show_column"`     // 0。作用尚不明确
	IsHitWebInf    bool                    `json:"is_hit_web_inf"`  // 是否命中web接口信息，实测为 false
	InBlackKey     int                     `json:"in_black_key"`    // 是否命中黑名单关键词，实测为 0
	InWhiteKey     int                     `json:"in_white_key"`    // 是否命中白名单关键词，实测为 0
}

// CategorySearchCostTime 是分类搜索的详细搜索用时。
type CategorySearchCostTime struct {
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

// CategorySearchPageInfo 是分类搜索的副分页信息。
type CategorySearchPageInfo struct {
	LiveRoom *CategorySearchPageInfoItem `json:"live_room"` // 直播间信息
	LiveUser *CategorySearchPageInfoItem `json:"live_user"` // 主播信息
}

// CategorySearchPageInfoItem 是分类搜索副分页信息中的一项。
type CategorySearchPageInfoItem struct {
	NumPages   int `json:"numPages"`   // 总计页数
	NumResults int `json:"numResults"` // 总计项数
	Total      int `json:"total"`      // 总计项数
	Pages      int `json:"pages"`      // 总计页数
}

// CategorySearchLiveResult 搜索类型为直播间及主播时的 result 对象。
type CategorySearchLiveResult struct {
	LiveRoom []*CategorySearchLiveRoomItem `json:"live_room"` // 直播间搜索结果
	LiveUser []*CategorySearchLiveUserItem `json:"live_user"` // 主播搜索结果
}

// CategorySearchRecTag 视频条目的推荐标签。
type CategorySearchRecTag struct {
	TagName  string `json:"tag_name"`  // 标签名称
	TagStyle int    `json:"tag_style"` // 标签样式
}

// CategorySearchVideoItem 搜索类型为 video 时的结果条目。
type CategorySearchVideoItem struct {
	Type             string                 `json:"type"`                // 结果类型，固定为 video
	ID               int                    `json:"id"`                  // 稿件avid
	Author           string                 `json:"author"`              // UP主昵称
	Mid              int                    `json:"mid"`                 // UP主mid
	TypeID           string                 `json:"typeid"`              // 视频分区tid
	TypeName         string                 `json:"typename"`            // 视频子分区名
	Arcurl           string                 `json:"arcurl"`              // 视频重定向url
	Aid              int                    `json:"aid"`                 // 稿件avid
	Bvid             string                 `json:"bvid"`                // 稿件bvid
	Title            string                 `json:"title"`               // 视频标题，关键字用 xml 标签标注
	Description      string                 `json:"description"`         // 视频简介
	Arcrank          string                 `json:"arcrank"`             // 0。作用尚不明确
	Pic              string                 `json:"pic"`                 // 视频封面url
	Play             int                    `json:"play"`                // 视频播放量
	VideoReview      int                    `json:"video_review"`        // 视频弹幕量
	Favorites        int                    `json:"favorites"`           // 视频收藏数
	Tag              string                 `json:"tag"`                 // 视频TAG，每项用 , 分隔
	Review           int                    `json:"review"`              // 视频评论数
	Pubdate          int64                  `json:"pubdate"`             // 视频投稿时间，时间戳
	Senddate         int64                  `json:"senddate"`            // 视频发布时间，时间戳
	Duration         string                 `json:"duration"`            // 视频时长，HH:MM
	BadgePay         bool                   `json:"badgepay"`            // false。作用尚不明确
	HitColumns       []string               `json:"hit_columns"`         // 关键字匹配类型：title、description、author、tag
	ViewType         string                 `json:"view_type"`           // 空。作用尚不明确
	IsPay            int                    `json:"is_pay"`              // 0。作用尚不明确
	IsUnionVideo     int                    `json:"is_union_video"`      // 是否为合作视频，0：否，1：是
	RecTags          []any                  `json:"rec_tags"`            // 作用尚不明确，实测为 null
	NewRecTags       []CategorySearchRecTag `json:"new_rec_tags"`        // 推荐标签
	RankScore        int                    `json:"rank_score"`          // 结果排序量化值
	BizData          any                    `json:"biz_data"`            // 业务数据，实测为 null
	IsChargeVideo    int                    `json:"is_charge_video"`     // 是否为充电专属视频，0：否，1：是
	EpisodeCountText string                 `json:"episode_count_text"`  // 分集数文本，实测为空
	IsIntervene      int                    `json:"is_intervene"`        // 是否被干预，实测为 0
	IsLiveRoomInline int                    `json:"is_live_room_inline"` // 是否内嵌直播间，实测为 0
	SpreadID         int                    `json:"spread_id"`           // 推广id，实测为 0
	CnyFlag          int                    `json:"cny_flag"`            // 是否春节活动，实测为 0
	CardStatus       int                    `json:"card_status"`         // 卡片状态，实测为 0
}

// CategorySearchMediaItem 搜索类型为 media_bangumi 或 media_ft 时的结果条目。
type CategorySearchMediaItem struct {
	Type           string                       `json:"type"`             // 结果类型，media_bangumi：番剧，media_ft：影视
	MediaID        int                          `json:"media_id"`         // 剧集mdid
	SeasonID       int                          `json:"season_id"`        // 剧集ssid
	Title          string                       `json:"title"`            // 剧集标题，关键字用 xml 标签标注
	OrgTitle       string                       `json:"org_title"`        // 剧集原名，关键字用 xml 标签标注，可为空
	Cover          string                       `json:"cover"`            // 剧集封面url
	MediaType      int                          `json:"media_type"`       // 剧集类型：1番剧，2电影，3纪录片，4国创，5电视剧，7综艺
	Areas          string                       `json:"areas"`            // 地区
	Styles         string                       `json:"styles"`           // 风格
	Cv             string                       `json:"cv"`               // 声优
	Staff          string                       `json:"staff"`            // 制作组
	PlayState      int                          `json:"play_state"`       // 0。作用尚不明确
	GotoURL        string                       `json:"goto_url"`         // 剧集重定向url
	Desc           string                       `json:"desc"`             // 简介
	Corner         int                          `json:"corner"`           // 角标有无：2无，13有
	Pubtime        int64                        `json:"pubtime"`          // 开播时间，时间戳
	MediaMode      int                          `json:"media_mode"`       // 2。作用尚不明确
	IsAvid         bool                         `json:"is_avid"`          // false。作用尚不明确
	FixPubtimeStr  string                       `json:"fix_pubtime_str"`  // 开播时间重写信息，优先级高于 pubtime，可为空
	MediaScore     *CategorySearchMediaScore    `json:"media_score"`      // 评分信息，无效时为 null
	HitColumns     []string                     `json:"hit_columns"`      // 关键字匹配类型：title、org_title，无效时为 null
	AllNetName     string                       `json:"all_net_name"`     // 空。作用尚不明确
	AllNetIcon     string                       `json:"all_net_icon"`     // 空。作用尚不明确
	AllNetURL      string                       `json:"all_net_url"`      // 空。作用尚不明确
	AngleTitle     string                       `json:"angle_title"`      // 角标内容
	AngleColor     int                          `json:"angle_color"`      // 角标颜色：0红色，2橙色
	DisplayInfo    []CategorySearchBadge        `json:"display_info"`     // 剧集标志信息
	HitEpids       string                       `json:"hit_epids"`        // 关键字匹配分集标题的分集epid，多个用 , 分隔
	PgcSeasonID    int                          `json:"pgc_season_id"`    // 剧集ssid
	SeasonType     int                          `json:"season_type"`      // 剧集类型：1番剧，2电影，3纪录片，4国创，5电视剧，7综艺
	SeasonTypeName string                       `json:"season_type_name"` // 剧集类型文字
	SelectionStyle string                       `json:"selection_style"`  // 分集选择按钮风格：horizontal横排式，grid按钮式
	EpSize         int                          `json:"ep_size"`          // 结果匹配的分集数
	URL            string                       `json:"url"`              // 剧集重定向url
	ButtonText     string                       `json:"button_text"`      // 观看按钮文字
	IsFollow       int                          `json:"is_follow"`        // 是否追番，0：否，1：是。需要登录，未登录恒为 0
	IsSelection    int                          `json:"is_selection"`     // 1。作用尚不明确
	Eps            []CategorySearchMediaEpisode `json:"eps"`              // 结果匹配的分集信息
	Badges         []CategorySearchBadge        `json:"badges"`           // 剧集标志信息
}

// CategorySearchMediaScore 是番剧条目的评分信息。
type CategorySearchMediaScore struct {
	UserCount int     `json:"user_count"` // 总计评分人数
	Score     float64 `json:"score"`      // 评分
}

// CategorySearchBadge 剧集标志信息。
type CategorySearchBadge struct {
	BgColorNight     string `json:"bg_color_night"`     // 夜间背景颜色，颜色码
	Text             string `json:"text"`               // 剧集标志
	BorderColor      string `json:"border_color"`       // 背景颜色，颜色码
	BgStyle          int    `json:"bg_style"`           // 1
	TextColor        string `json:"text_color"`         // 文字颜色，颜色码
	BgColor          string `json:"bg_color"`           // 背景颜色，颜色码
	TextColorNight   string `json:"text_color_night"`   // 夜间文字颜色，颜色码
	BorderColorNight string `json:"border_color_night"` // 夜间背景颜色，颜色码
}

// CategorySearchMediaEpisode 番剧条目的分集信息。
type CategorySearchMediaEpisode struct {
	ID          int                   `json:"id"`           // 分集epid
	Cover       string                `json:"cover"`        // 分集封面url
	Title       string                `json:"title"`        // 完整标题
	URL         string                `json:"url"`          // 分集重定向url
	ReleaseDate string                `json:"release_date"` // 空
	Badges      []CategorySearchBadge `json:"badges"`       // 分集标志
	IndexTitle  string                `json:"index_title"`  // 短标题
	LongTitle   string                `json:"long_title"`   // 单集标题
}

// CategorySearchLiveRoomItem 搜索类型为 live_room 时的结果条目。
type CategorySearchLiveRoomItem struct {
	Type       string   `json:"type"`        // 结果类型，固定为 live_room
	RankOffset int      `json:"rank_offset"` // 搜索结果排名值
	UID        int      `json:"uid"`         // 主播mid
	Tags       string   `json:"tags"`        // 直播间TAG，多个用 , 分隔
	HitColumns []string `json:"hit_columns"` // 关键字匹配类型：title、cate_name、uname
	LiveTime   string   `json:"live_time"`   // 开播时间，YYYY-MM-DD HH:MM:SS
	CateName   string   `json:"cate_name"`   // 子分区名
	LiveStatus int      `json:"live_status"` // 1
	Uname      string   `json:"uname"`       // 主播昵称
	Uface      string   `json:"uface"`       // 主播头像url
	UserCover  string   `json:"user_cover"`  // 直播间封面url
	ShortID    int      `json:"short_id"`    // 0。作用尚不明确
	Area       int      `json:"area"`        // 1。作用尚不明确
	Title      string   `json:"title"`       // 直播间标题，关键字用 xml 标签标注
	Cover      string   `json:"cover"`       // 关键帧截图url
	Online     int      `json:"online"`      // 在线人数
	RankIndex  int      `json:"rank_index"`  // 0。作用尚不明确
	RankScore  int      `json:"rank_score"`  // 结果排序量化值
	RoomID     int      `json:"roomid"`      // 直播间id
	Attentions int      `json:"attentions"`  // 主播粉丝数
}

// CategorySearchLiveUserItem 搜索类型为 live_user 时的结果条目。
type CategorySearchLiveUserItem struct {
	Type       string   `json:"type"`        // 结果类型，固定为 live_user
	RankOffset int      `json:"rank_offset"` // 搜索结果排名值
	UID        int      `json:"uid"`         // 主播mid
	Tags       string   `json:"tags"`        // 直播间TAG，多个用 , 分隔
	LiveTime   string   `json:"live_time"`   // 开播时间，YYYY-MM-DD HH:MM:SS。未开播为 0000-00-00 00:00:00
	HitColumns []string `json:"hit_columns"` // 关键字匹配类型：uname
	LiveStatus int      `json:"live_status"` // 是否开播，0：未开播，1：已开播
	Area       int      `json:"area"`        // 1。作用尚不明确
	IsLive     bool     `json:"is_live"`     // 是否开播
	Uname      string   `json:"uname"`       // 主播昵称，关键字用 xml 标签标注
	Uface      string   `json:"uface"`       // 主播头像url
	RankIndex  int      `json:"rank_index"`  // 0。作用尚不明确
	RankScore  int      `json:"rank_score"`  // 结果排序量化值
	RoomID     int      `json:"roomid"`      // 直播间id
	Attentions int      `json:"attentions"`  // 主播粉丝数
}

// CategorySearchArticleItem 搜索类型为 article 时的结果条目。
type CategorySearchArticleItem struct {
	Type         string   `json:"type"`          // 结果类型，固定为 article
	RankOffset   int      `json:"rank_offset"`   // 搜索结果排名值
	PubTime      int64    `json:"pub_time"`      // 投稿时间，时间戳
	Like         int      `json:"like"`          // 获赞数
	Title        string   `json:"title"`         // 标题
	Mid          int      `json:"mid"`           // UP主mid
	ImageURLs    []string `json:"image_urls"`    // 封面图组
	TemplateID   int      `json:"template_id"`   // 作用尚不明确
	CategoryID   int      `json:"category_id"`   // 专栏分区
	View         int      `json:"view"`          // 阅读数
	Reply        int      `json:"reply"`         // 评论数
	RankIndex    int      `json:"rank_index"`    // 0。作用尚不明确
	Desc         string   `json:"desc"`          // 文章预览
	RankScore    int      `json:"rank_score"`    // 结果排序量化值
	ID           int      `json:"id"`            // 专栏cvid
	CategoryName string   `json:"category_name"` // 子分区名
}

// CategorySearchTopicItem 搜索类型为 topic 时的结果条目。
type CategorySearchTopicItem struct {
	Type        string   `json:"type"`        // 结果类型，固定为 topic
	Description string   `json:"description"` // 简介
	Pubdate     int64    `json:"pubdate"`     // 发布时间，时间戳
	Title       string   `json:"title"`       // 标题
	Favourite   int      `json:"favourite"`   // 0。作用尚不明确
	HitColumns  []string `json:"hit_columns"` // 关键字匹配类型：title
	Review      int      `json:"review"`      // 0。作用尚不明确
	RankOffset  int      `json:"rank_offset"` // 搜索结果排名值
	Cover       string   `json:"cover"`       // 话题封面url
	Update      int64    `json:"update"`      // 上传时间，时间戳
	Mid         int      `json:"mid"`         // 0。作用尚不明确
	Click       int      `json:"click"`       // 作用尚不明确
	TpType      int      `json:"tp_type"`     // 作用尚不明确
	Keyword     string   `json:"keyword"`     // 空。作用尚不明确
	TpID        int      `json:"tp_id"`       // 话题tp
	RankIndex   int      `json:"rank_index"`  // 0。作用尚不明确
	Author      string   `json:"author"`      // UP主昵称
	Arcurl      string   `json:"arcurl"`      // 话题页面重定向url
	RankScore   int      `json:"rank_score"`  // 结果排序量化值
}

// CategorySearchUserItem 搜索类型为 bili_user 时的结果条目。
type CategorySearchUserItem struct {
	Type           string                            `json:"type"`            // 结果类型，固定为 bili_user
	Mid            int                               `json:"mid"`             // 用户mid
	Uname          string                            `json:"uname"`           // 用户昵称
	Usign          string                            `json:"usign"`           // 用户签名
	Fans           int                               `json:"fans"`            // 用户粉丝数
	Videos         int                               `json:"videos"`          // 用户稿件数
	Upic           string                            `json:"upic"`            // 用户头像url
	VerifyInfo     string                            `json:"verify_info"`     // 空。作用尚不明确
	Level          int                               `json:"level"`           // 用户等级
	Gender         int                               `json:"gender"`          // 用户性别：1男，2女，3私密
	IsUpuser       int                               `json:"is_upuser"`       // 是否为UP主，0：否，1：是
	IsLive         int                               `json:"is_live"`         // 是否正在直播，0：否，1：是
	RoomID         int                               `json:"room_id"`         // 用户直播间id
	Res            []CategorySearchUserVideo         `json:"res"`             // 用户投稿内容
	OfficialVerify *CategorySearchUserOfficialVerify `json:"official_verify"` // 用户认证信息
	HitColumns     []string                          `json:"hit_columns"`     // 关键字匹配类型：uname
}

// CategorySearchUserVideo 用户条目中的投稿视频。
type CategorySearchUserVideo struct {
	Aid          int    `json:"aid"`            // 稿件avid
	Bvid         string `json:"bvid"`           // 稿件bvid
	Title        string `json:"title"`          // 视频标题
	Pubdate      int64  `json:"pubdate"`        // 视频投稿时间，时间戳
	Arcurl       string `json:"arcurl"`         // 视频页面重定向url
	Pic          string `json:"pic"`            // 视频封面图片url
	Play         string `json:"play"`           // 播放量
	Dm           int    `json:"dm"`             // 弹幕量
	Coin         int    `json:"coin"`           // 投币数
	Fav          int    `json:"fav"`            // 收藏数
	Desc         string `json:"desc"`           // 视频简介
	Duration     string `json:"duration"`       // 视频时长，MM:SS
	IsPay        int    `json:"is_pay"`         // 0。作用尚不明确
	IsUnionVideo int    `json:"is_union_video"` // 是否为合作视频，0：否，1：是
}

// CategorySearchUserOfficialVerify 用户认证信息。
type CategorySearchUserOfficialVerify struct {
	Type int    `json:"type"` // 是否认证：127无，0个人认证，1组织认证
	Desc string `json:"desc"` // 认证名称
}

// CategorySearchPhotoItem 搜索类型为 photo 时的结果条目。
type CategorySearchPhotoItem struct {
	Type       string   `json:"type"`        // 结果类型，固定为 photo
	Count      int      `json:"count"`       // 图片数
	Like       int      `json:"like"`        // 收藏数
	Title      string   `json:"title"`       // 相簿标题
	HitColumns []string `json:"hit_columns"` // 关键字匹配类型：title、description、source_tag
	RankOffset int      `json:"rank_offset"` // 搜索结果排名值
	Cover      string   `json:"cover"`       // 相簿封面url
	Mid        int      `json:"mid"`         // UP主mid
	Uname      string   `json:"uname"`       // UP主昵称
	RankIndex  int      `json:"rank_index"`  // 0。作用尚不明确
	View       int      `json:"view"`        // 观看次数
	ID         int      `json:"id"`          // 相簿id
	RankScore  int      `json:"rank_score"`  // 结果排序量化值
}
