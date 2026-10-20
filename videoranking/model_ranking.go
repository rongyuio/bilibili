package videoranking

// 分区视频排行接口的响应模型。
//
// 每个接口的返回体独立定义，嵌套对象也不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// GetZoneVideoRankListResult 是「获取分区视频排行榜列表」的 data 对象。
type GetZoneVideoRankListResult struct {
	Note string                     `json:"note"` // “根据稿件内容质量、近期的数据综合展示，动态更新”
	List []GetZoneVideoRankListItem `json:"list"` // 视频列表
}

// GetZoneVideoRankListItem 是「获取分区视频排行榜列表」列表中的视频对象。
//
// 字段集同根包的 VideoInfo，但按接口独立定义，不与其它接口共用。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的视频对象结构重复是设计使然
type GetZoneVideoRankListItem struct {
	Bvid                    string                             `json:"bvid"`         // 稿件bvid
	Aid                     int                                `json:"aid"`          // 稿件avid
	Videos                  int                                `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                                `json:"tid"`          // 分区tid
	Tname                   string                             `json:"tname"`        // 子分区名称
	Copyright               int                                `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                             `json:"pic"`          // 稿件封面图片url
	Title                   string                             `json:"title"`        // 稿件标题
	PubDate                 int                                `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                                `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                             `json:"desc"`         // 视频简介
	DescV2                  []GetZoneVideoRankListItemDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                                `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int                                `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                                `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                                `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                             `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetZoneVideoRankListItemRights     `json:"rights"`       // 视频属性标志
	Owner                   GetZoneVideoRankListItemOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetZoneVideoRankListItemStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                             `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                                `json:"cid"`          // 视频1P cid
	Dimension               GetZoneVideoRankListItemDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                                `json:"season_id"`    // 合集id
	Premiere                any                                `json:"premiere"`     // null
	TeenageMode             int                                `json:"teenage_mode"`
	IsChargeableSeason      bool                               `json:"is_chargeable_season"`
	IsStory                 bool                               `json:"is_story"`
	NoCache                 bool                               `json:"no_cache"` // 作用尚不明确
	Pages                   []GetZoneVideoRankListItemPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetZoneVideoRankListItemSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetZoneVideoRankListItemStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                               `json:"is_season_display"`
	UserGarb                GetZoneVideoRankListItemUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetZoneVideoRankListItemHonorReply `json:"honor_reply"`
	LikeIcon                string                             `json:"like_icon"`
	ArgueInfo               GetZoneVideoRankListItemArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                                `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                             `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                                `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                             `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                                `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                             `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                                `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                                `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                               `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                                `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetZoneVideoRankListItemAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                                `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string                             `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                               `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                               `json:"is_upower_play"`              //
	IsUpowerPreview         bool                               `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                               `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                               `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetZoneVideoRankListItemUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                               `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                               `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                               `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                               `json:"is_view_self"`                // 是否尽自己可见
}

// GetZoneVideoRankListItemDesc 是「获取分区视频排行榜列表」视频对象中的 desc_v2 数组元素。
type GetZoneVideoRankListItemDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetZoneVideoRankListItemRights 是「获取分区视频排行榜列表」视频对象中的 rights。
type GetZoneVideoRankListItemRights struct {
	Bp            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	Hd5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UgcPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UgcPayPreview int `json:"ugc_pay_preview"` // 0。作用尚不明确
	NoBackground  int `json:"no_background"`   // 0。作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0。作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0。作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0。作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0。作用尚不明确
}

// GetZoneVideoRankListItemOwner 是「获取分区视频排行榜列表」视频对象中的 owner。
type GetZoneVideoRankListItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetZoneVideoRankListItemStat 是「获取分区视频排行榜列表」视频对象中的 stat。
type GetZoneVideoRankListItemStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetZoneVideoRankListItemDimension 是「获取分区视频排行榜列表」视频对象中的 dimension。
type GetZoneVideoRankListItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetZoneVideoRankListItemPage 是「获取分区视频排行榜列表」视频对象中的 pages 数组元素。
type GetZoneVideoRankListItemPage struct {
	Cid        int                               `json:"cid"`         // 当前分P cid
	Page       int                               `json:"page"`        // 当前分P
	From       string                            `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                            `json:"part"`        // 当前分P标题
	Duration   int                               `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                            `json:"vid"`         // 站外视频vid
	Weblink    string                            `json:"weblink"`     // 站外视频跳转url
	Dimension  GetZoneVideoRankListItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                            `json:"first_frame"` // 分P封面
}

// GetZoneVideoRankListItemSubtitle 是「获取分区视频排行榜列表」视频对象中的 subtitle。
type GetZoneVideoRankListItemSubtitle struct {
	ID          int                                    `json:"id"`           // 字幕id
	Lan         string                                 `json:"lan"`          // 字幕语言
	LanDoc      string                                 `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                                   `json:"is_lock"`      // 是否锁定
	AuthorMid   int                                    `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string                                 `json:"subtitle_url"` // json格式字幕文件url
	Author      GetZoneVideoRankListItemSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// GetZoneVideoRankListItemSubtitleAuthor 是「获取分区视频排行榜列表」视频对象中的 subtitle.author。
type GetZoneVideoRankListItemSubtitleAuthor struct {
	Mid           int    `json:"mid"`             // 字幕上传者mid
	Name          string `json:"name"`            // 字幕上传者昵称
	Sex           string `json:"sex"`             // 字幕上传者性别。男 女 保密
	Face          string `json:"face"`            // 字幕上传者头像url
	Sign          string `json:"sign"`            // 字幕上传者签名
	Rank          int    `json:"rank"`            // 10000。作用尚不明确
	Birthday      int    `json:"birthday"`        // 0。作用尚不明确
	IsFakeAccount int    `json:"is_fake_account"` // 0。作用尚不明确
	IsDeleted     int    `json:"is_deleted"`      // 0。作用尚不明确
}

// GetZoneVideoRankListItemStaff 是「获取分区视频排行榜列表」视频对象中的 staff 数组元素。
type GetZoneVideoRankListItemStaff struct {
	Mid        int                                   `json:"mid"`      // 成员mid
	Title      string                                `json:"title"`    // 成员名称
	Name       string                                `json:"name"`     // 成员昵称
	Face       string                                `json:"face"`     // 成员头像url
	Vip        GetZoneVideoRankListItemStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetZoneVideoRankListItemStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                                   `json:"follower"` // 成员粉丝数
	LabelStyle int                                   `json:"label_style"`
}

// GetZoneVideoRankListItemStaffVip 是「获取分区视频排行榜列表」视频对象中的 staff[].vip。
type GetZoneVideoRankListItemStaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

// GetZoneVideoRankListItemStaffOfficial 是「获取分区视频排行榜列表」视频对象中的 staff[].official。
type GetZoneVideoRankListItemStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetZoneVideoRankListItemUserGarb 是「获取分区视频排行榜列表」视频对象中的 user_garb。
type GetZoneVideoRankListItemUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetZoneVideoRankListItemHonorReply 是「获取分区视频排行榜列表」视频对象中的 honor_reply。
type GetZoneVideoRankListItemHonorReply struct {
	Honor []GetZoneVideoRankListItemHonor `json:"honor"`
}

// GetZoneVideoRankListItemHonor 是「获取分区视频排行榜列表」视频对象中的 honor_reply.honor 数组元素。
type GetZoneVideoRankListItemHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetZoneVideoRankListItemArgueInfo 是「获取分区视频排行榜列表」视频对象中的 argue_info。
type GetZoneVideoRankListItemArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetZoneVideoRankListItemAiRcmd 是「获取分区视频排行榜列表」视频对象中的 ai_rcmd。
type GetZoneVideoRankListItemAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetZoneVideoRankListItemUgcSeason 是「获取分区视频排行榜列表」视频对象中的 ugc_season。
type GetZoneVideoRankListItemUgcSeason struct {
	Id          int                                        `json:"id"`            // 视频合集id
	Title       string                                     `json:"title"`         // 视频合集标题
	Cover       string                                     `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                                        `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                                     `json:"intro"`         // 视频合集介绍
	SignState   int                                        `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                        `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []GetZoneVideoRankListItemUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetZoneVideoRankListItemUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                        `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                        `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                       `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                        `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetZoneVideoRankListItemUgcSeasonStat 是「获取分区视频排行榜列表」视频对象中的 ugc_season.stat。
type GetZoneVideoRankListItemUgcSeasonStat struct {
	SeasonID int `json:"season_id"` // 视频合集id
	View     int `json:"view"`      // 视频合集总浏览量
	Danmaku  int `json:"danmaku"`   // 视频合集总弹幕量
	Reply    int `json:"reply"`     // 视频合集总评论量
	Fav      int `json:"fav"`       // 视频合集总收藏数
	Coin     int `json:"coin"`      // 视频合集总投币数
	Share    int `json:"share"`     // 视频合集总分享数
	NowRank  int `json:"now_rank"`  // 视频合集当前排名
	HisRank  int `json:"his_rank"`  // 视频合集历史排名
	Like     int `json:"like"`      // 视频合集总获赞数
	Vt       int `json:"vt"`        // （？）。作用尚不明确
	Vv       int `json:"vv"`        // （？）。作用尚不明确
}

// GetZoneVideoRankListItemUgcSeasonSection 是「获取分区视频排行榜列表」视频对象中的 ugc_season.sections 数组元素。
type GetZoneVideoRankListItemUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                                        `json:"id"`
	SectionID int                                        `json:"section_id"`
	Title     string                                     `json:"title"`    // 视频合集中分部标题
	Type      int                                        `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetZoneVideoRankListItemUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetZoneVideoRankListItemUgcSeasonEpisode 是「获取分区视频排行榜列表」视频对象中的 ugc_season.sections[].episodes 数组元素。
type GetZoneVideoRankListItemUgcSeasonEpisode struct {
	SeasonID  int                                         `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                         `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                         `json:"id"`         // 分部中视频id
	Aid       int                                         `json:"aid"`        // 视频aid
	Cid       int                                         `json:"cid"`        // 视频cid
	Title     string                                      `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                                         `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetZoneVideoRankListItemUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetZoneVideoRankListItemPage                `json:"page"`       // 视频分P信息
	Bvid      string                                      `json:"bvid"`       // 视频bvid
	Pages     []GetZoneVideoRankListItemPage              `json:"pages"`      // 视频分P列表
}

// GetZoneVideoRankListItemUgcSeasonEpisodeArc 是「获取分区视频排行榜列表」视频对象中的 ugc_season.sections[].episodes[].arc。
type GetZoneVideoRankListItemUgcSeasonEpisodeArc struct {
	Aid                int                                          `json:"aid"`                  // 稿件avid
	Videos             int                                          `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                                          `json:"type_id"`              // 分区tid
	TypeName           string                                       `json:"type_name"`            // 子分区名称
	Copyright          int                                          `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string                                       `json:"pic"`                  // 稿件封面图片url
	Title              string                                       `json:"title"`                // 稿件标题
	PubDate            int                                          `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                                          `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string                                       `json:"desc"`                 // 视频简介
	State              int                                          `json:"state"`                // 视频状态
	Duration           int                                          `json:"duration"`             // 稿件总时长。单位为秒
	Rights             GetZoneVideoRankListItemRights               `json:"rights"`               // 视频属性标志
	Author             GetZoneVideoRankListItemOwner                `json:"author"`               // UP主信息
	Stat               GetZoneVideoRankListItemUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                       `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetZoneVideoRankListItemDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                          `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                         `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                         `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                          `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                       `json:"vt_display"`           // （？）。作用尚不明确
}

// GetZoneVideoRankListItemUgcSeasonEpisodeStat 是「获取分区视频排行榜列表」视频对象中的 ugc_season.sections[].episodes[].arc.stat。
type GetZoneVideoRankListItemUgcSeasonEpisodeStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Fav        int    `json:"fav"`        // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	ArgueMsg   string `json:"argue_msg"`  // 警告信息
	Vt         int    `json:"vt"`         // （？）。作用尚不明确
	Vv         int    `json:"vv"`         // （？）。作用尚不明确
}

// GetZoneVideoListNewResult 是「获取分区最新视频列表」的 data 对象。
type GetZoneVideoListNewResult struct {
	Archives []GetZoneVideoListNewItem `json:"archives"` // 视频列表
	Page     GetZoneVideoListNewPage   `json:"page"`     // 页面信息
}

// GetZoneVideoListNewPage 是「获取分区最新视频列表」的分页信息。
//
// 字段名为 num/size，与空间投稿的 pn/ps、合集的 page_num/page_size 不同，
// 因此保留独立类型，不做跨接口合并。
type GetZoneVideoListNewPage struct {
	Count int `json:"count"` // 总计视频数
	Num   int `json:"num"`   // 当前页码
	Size  int `json:"size"`  // 每页项数
}

// GetZoneVideoListNewItem 是「获取分区最新视频列表」列表中的视频对象。
//
// 字段集同根包的 VideoInfo，但按接口独立定义，不与其它接口共用。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的视频对象结构重复是设计使然
type GetZoneVideoListNewItem struct {
	Bvid                    string                            `json:"bvid"`         // 稿件bvid
	Aid                     int                               `json:"aid"`          // 稿件avid
	Videos                  int                               `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                               `json:"tid"`          // 分区tid
	Tname                   string                            `json:"tname"`        // 子分区名称
	Copyright               int                               `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                            `json:"pic"`          // 稿件封面图片url
	Title                   string                            `json:"title"`        // 稿件标题
	PubDate                 int                               `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                               `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                            `json:"desc"`         // 视频简介
	DescV2                  []GetZoneVideoListNewItemDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                               `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int                               `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                               `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                               `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                            `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetZoneVideoListNewItemRights     `json:"rights"`       // 视频属性标志
	Owner                   GetZoneVideoListNewItemOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetZoneVideoListNewItemStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                            `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                               `json:"cid"`          // 视频1P cid
	Dimension               GetZoneVideoListNewItemDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                               `json:"season_id"`    // 合集id
	Premiere                any                               `json:"premiere"`     // null
	TeenageMode             int                               `json:"teenage_mode"`
	IsChargeableSeason      bool                              `json:"is_chargeable_season"`
	IsStory                 bool                              `json:"is_story"`
	NoCache                 bool                              `json:"no_cache"` // 作用尚不明确
	Pages                   []GetZoneVideoListNewItemPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetZoneVideoListNewItemSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetZoneVideoListNewItemStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                              `json:"is_season_display"`
	UserGarb                GetZoneVideoListNewItemUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetZoneVideoListNewItemHonorReply `json:"honor_reply"`
	LikeIcon                string                            `json:"like_icon"`
	ArgueInfo               GetZoneVideoListNewItemArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                               `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                            `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                               `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                            `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                               `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                            `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                               `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                               `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                              `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                               `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetZoneVideoListNewItemAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                               `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string                            `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                              `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                              `json:"is_upower_play"`              //
	IsUpowerPreview         bool                              `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                              `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                              `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetZoneVideoListNewItemUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                              `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                              `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                              `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                              `json:"is_view_self"`                // 是否尽自己可见
}

// GetZoneVideoListNewItemDesc 是「获取分区最新视频列表」视频对象中的 desc_v2 数组元素。
type GetZoneVideoListNewItemDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetZoneVideoListNewItemRights 是「获取分区最新视频列表」视频对象中的 rights。
type GetZoneVideoListNewItemRights struct {
	Bp            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	Hd5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UgcPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UgcPayPreview int `json:"ugc_pay_preview"` // 0。作用尚不明确
	NoBackground  int `json:"no_background"`   // 0。作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0。作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0。作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0。作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0。作用尚不明确
}

// GetZoneVideoListNewItemOwner 是「获取分区最新视频列表」视频对象中的 owner。
type GetZoneVideoListNewItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetZoneVideoListNewItemStat 是「获取分区最新视频列表」视频对象中的 stat。
type GetZoneVideoListNewItemStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetZoneVideoListNewItemDimension 是「获取分区最新视频列表」视频对象中的 dimension。
type GetZoneVideoListNewItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetZoneVideoListNewItemPage 是「获取分区最新视频列表」视频对象中的 pages 数组元素。
type GetZoneVideoListNewItemPage struct {
	Cid        int                              `json:"cid"`         // 当前分P cid
	Page       int                              `json:"page"`        // 当前分P
	From       string                           `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                           `json:"part"`        // 当前分P标题
	Duration   int                              `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                           `json:"vid"`         // 站外视频vid
	Weblink    string                           `json:"weblink"`     // 站外视频跳转url
	Dimension  GetZoneVideoListNewItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                           `json:"first_frame"` // 分P封面
}

// GetZoneVideoListNewItemSubtitle 是「获取分区最新视频列表」视频对象中的 subtitle。
type GetZoneVideoListNewItemSubtitle struct {
	ID          int                                   `json:"id"`           // 字幕id
	Lan         string                                `json:"lan"`          // 字幕语言
	LanDoc      string                                `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                                  `json:"is_lock"`      // 是否锁定
	AuthorMid   int                                   `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string                                `json:"subtitle_url"` // json格式字幕文件url
	Author      GetZoneVideoListNewItemSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// GetZoneVideoListNewItemSubtitleAuthor 是「获取分区最新视频列表」视频对象中的 subtitle.author。
type GetZoneVideoListNewItemSubtitleAuthor struct {
	Mid           int    `json:"mid"`             // 字幕上传者mid
	Name          string `json:"name"`            // 字幕上传者昵称
	Sex           string `json:"sex"`             // 字幕上传者性别。男 女 保密
	Face          string `json:"face"`            // 字幕上传者头像url
	Sign          string `json:"sign"`            // 字幕上传者签名
	Rank          int    `json:"rank"`            // 10000。作用尚不明确
	Birthday      int    `json:"birthday"`        // 0。作用尚不明确
	IsFakeAccount int    `json:"is_fake_account"` // 0。作用尚不明确
	IsDeleted     int    `json:"is_deleted"`      // 0。作用尚不明确
}

// GetZoneVideoListNewItemStaff 是「获取分区最新视频列表」视频对象中的 staff 数组元素。
type GetZoneVideoListNewItemStaff struct {
	Mid        int                                  `json:"mid"`      // 成员mid
	Title      string                               `json:"title"`    // 成员名称
	Name       string                               `json:"name"`     // 成员昵称
	Face       string                               `json:"face"`     // 成员头像url
	Vip        GetZoneVideoListNewItemStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetZoneVideoListNewItemStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                                  `json:"follower"` // 成员粉丝数
	LabelStyle int                                  `json:"label_style"`
}

// GetZoneVideoListNewItemStaffVip 是「获取分区最新视频列表」视频对象中的 staff[].vip。
type GetZoneVideoListNewItemStaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

// GetZoneVideoListNewItemStaffOfficial 是「获取分区最新视频列表」视频对象中的 staff[].official。
type GetZoneVideoListNewItemStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetZoneVideoListNewItemUserGarb 是「获取分区最新视频列表」视频对象中的 user_garb。
type GetZoneVideoListNewItemUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetZoneVideoListNewItemHonorReply 是「获取分区最新视频列表」视频对象中的 honor_reply。
type GetZoneVideoListNewItemHonorReply struct {
	Honor []GetZoneVideoListNewItemHonor `json:"honor"`
}

// GetZoneVideoListNewItemHonor 是「获取分区最新视频列表」视频对象中的 honor_reply.honor 数组元素。
type GetZoneVideoListNewItemHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetZoneVideoListNewItemArgueInfo 是「获取分区最新视频列表」视频对象中的 argue_info。
type GetZoneVideoListNewItemArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetZoneVideoListNewItemAiRcmd 是「获取分区最新视频列表」视频对象中的 ai_rcmd。
type GetZoneVideoListNewItemAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetZoneVideoListNewItemUgcSeason 是「获取分区最新视频列表」视频对象中的 ugc_season。
type GetZoneVideoListNewItemUgcSeason struct {
	Id          int                                       `json:"id"`            // 视频合集id
	Title       string                                    `json:"title"`         // 视频合集标题
	Cover       string                                    `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                                       `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                                    `json:"intro"`         // 视频合集介绍
	SignState   int                                       `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                       `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []GetZoneVideoListNewItemUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetZoneVideoListNewItemUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                       `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                       `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                      `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                       `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetZoneVideoListNewItemUgcSeasonStat 是「获取分区最新视频列表」视频对象中的 ugc_season.stat。
type GetZoneVideoListNewItemUgcSeasonStat struct {
	SeasonID int `json:"season_id"` // 视频合集id
	View     int `json:"view"`      // 视频合集总浏览量
	Danmaku  int `json:"danmaku"`   // 视频合集总弹幕量
	Reply    int `json:"reply"`     // 视频合集总评论量
	Fav      int `json:"fav"`       // 视频合集总收藏数
	Coin     int `json:"coin"`      // 视频合集总投币数
	Share    int `json:"share"`     // 视频合集总分享数
	NowRank  int `json:"now_rank"`  // 视频合集当前排名
	HisRank  int `json:"his_rank"`  // 视频合集历史排名
	Like     int `json:"like"`      // 视频合集总获赞数
	Vt       int `json:"vt"`        // （？）。作用尚不明确
	Vv       int `json:"vv"`        // （？）。作用尚不明确
}

// GetZoneVideoListNewItemUgcSeasonSection 是「获取分区最新视频列表」视频对象中的 ugc_season.sections 数组元素。
type GetZoneVideoListNewItemUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                                       `json:"id"`
	SectionID int                                       `json:"section_id"`
	Title     string                                    `json:"title"`    // 视频合集中分部标题
	Type      int                                       `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetZoneVideoListNewItemUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetZoneVideoListNewItemUgcSeasonEpisode 是「获取分区最新视频列表」视频对象中的 ugc_season.sections[].episodes 数组元素。
type GetZoneVideoListNewItemUgcSeasonEpisode struct {
	SeasonID  int                                        `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                        `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                        `json:"id"`         // 分部中视频id
	Aid       int                                        `json:"aid"`        // 视频aid
	Cid       int                                        `json:"cid"`        // 视频cid
	Title     string                                     `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                                        `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetZoneVideoListNewItemUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetZoneVideoListNewItemPage                `json:"page"`       // 视频分P信息
	Bvid      string                                     `json:"bvid"`       // 视频bvid
	Pages     []GetZoneVideoListNewItemPage              `json:"pages"`      // 视频分P列表
}

// GetZoneVideoListNewItemUgcSeasonEpisodeArc 是「获取分区最新视频列表」视频对象中的 ugc_season.sections[].episodes[].arc。
type GetZoneVideoListNewItemUgcSeasonEpisodeArc struct {
	Aid                int                                         `json:"aid"`                  // 稿件avid
	Videos             int                                         `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                                         `json:"type_id"`              // 分区tid
	TypeName           string                                      `json:"type_name"`            // 子分区名称
	Copyright          int                                         `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string                                      `json:"pic"`                  // 稿件封面图片url
	Title              string                                      `json:"title"`                // 稿件标题
	PubDate            int                                         `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                                         `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string                                      `json:"desc"`                 // 视频简介
	State              int                                         `json:"state"`                // 视频状态
	Duration           int                                         `json:"duration"`             // 稿件总时长。单位为秒
	Rights             GetZoneVideoListNewItemRights               `json:"rights"`               // 视频属性标志
	Author             GetZoneVideoListNewItemOwner                `json:"author"`               // UP主信息
	Stat               GetZoneVideoListNewItemUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                      `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetZoneVideoListNewItemDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                         `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                        `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                        `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                         `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                      `json:"vt_display"`           // （？）。作用尚不明确
}

// GetZoneVideoListNewItemUgcSeasonEpisodeStat 是「获取分区最新视频列表」视频对象中的 ugc_season.sections[].episodes[].arc.stat。
type GetZoneVideoListNewItemUgcSeasonEpisodeStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Fav        int    `json:"fav"`        // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	ArgueMsg   string `json:"argue_msg"`  // 警告信息
	Vt         int    `json:"vt"`         // （？）。作用尚不明确
	Vv         int    `json:"vv"`         // （？）。作用尚不明确
}

// GetZoneVideoListWithTagResult 是「获取分区标签近期互动列表」的 data 对象。
type GetZoneVideoListWithTagResult struct {
	Archives []GetZoneVideoListWithTagItem `json:"archives"` // 视频列表
	Page     GetZoneVideoListWithTagPage   `json:"page"`     // 页面信息
}

// GetZoneVideoListWithTagPage 是「获取分区标签近期互动列表」的分页信息。
//
// 字段名为 num/size，与空间投稿的 pn/ps、合集的 page_num/page_size 不同，
// 因此保留独立类型，不做跨接口合并。
type GetZoneVideoListWithTagPage struct {
	Count int `json:"count"` // 总计视频数
	Num   int `json:"num"`   // 当前页码
	Size  int `json:"size"`  // 每页项数
}

// GetZoneVideoListWithTagItem 是「获取分区标签近期互动列表」列表中的视频对象。
//
// 字段集同根包的 VideoInfo，但按接口独立定义，不与其它接口共用。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的视频对象结构重复是设计使然
type GetZoneVideoListWithTagItem struct {
	Bvid                    string                                `json:"bvid"`         // 稿件bvid
	Aid                     int                                   `json:"aid"`          // 稿件avid
	Videos                  int                                   `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                                   `json:"tid"`          // 分区tid
	Tname                   string                                `json:"tname"`        // 子分区名称
	Copyright               int                                   `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                                `json:"pic"`          // 稿件封面图片url
	Title                   string                                `json:"title"`        // 稿件标题
	PubDate                 int                                   `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                                   `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                                `json:"desc"`         // 视频简介
	DescV2                  []GetZoneVideoListWithTagItemDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                                   `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int                                   `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                                   `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                                   `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                                `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetZoneVideoListWithTagItemRights     `json:"rights"`       // 视频属性标志
	Owner                   GetZoneVideoListWithTagItemOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetZoneVideoListWithTagItemStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                                `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                                   `json:"cid"`          // 视频1P cid
	Dimension               GetZoneVideoListWithTagItemDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                                   `json:"season_id"`    // 合集id
	Premiere                any                                   `json:"premiere"`     // null
	TeenageMode             int                                   `json:"teenage_mode"`
	IsChargeableSeason      bool                                  `json:"is_chargeable_season"`
	IsStory                 bool                                  `json:"is_story"`
	NoCache                 bool                                  `json:"no_cache"` // 作用尚不明确
	Pages                   []GetZoneVideoListWithTagItemPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetZoneVideoListWithTagItemSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetZoneVideoListWithTagItemStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                                  `json:"is_season_display"`
	UserGarb                GetZoneVideoListWithTagItemUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetZoneVideoListWithTagItemHonorReply `json:"honor_reply"`
	LikeIcon                string                                `json:"like_icon"`
	ArgueInfo               GetZoneVideoListWithTagItemArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                                   `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                                `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                                   `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                                `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                                   `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                                `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                                   `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                                   `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                                  `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                                   `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetZoneVideoListWithTagItemAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                                   `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string                                `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                                  `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                                  `json:"is_upower_play"`              //
	IsUpowerPreview         bool                                  `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                                  `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                                  `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetZoneVideoListWithTagItemUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                                  `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                                  `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                                  `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                                  `json:"is_view_self"`                // 是否尽自己可见
}

// GetZoneVideoListWithTagItemDesc 是「获取分区标签近期互动列表」视频对象中的 desc_v2 数组元素。
type GetZoneVideoListWithTagItemDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetZoneVideoListWithTagItemRights 是「获取分区标签近期互动列表」视频对象中的 rights。
type GetZoneVideoListWithTagItemRights struct {
	Bp            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	Hd5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UgcPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UgcPayPreview int `json:"ugc_pay_preview"` // 0。作用尚不明确
	NoBackground  int `json:"no_background"`   // 0。作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0。作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0。作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0。作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0。作用尚不明确
}

// GetZoneVideoListWithTagItemOwner 是「获取分区标签近期互动列表」视频对象中的 owner。
type GetZoneVideoListWithTagItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetZoneVideoListWithTagItemStat 是「获取分区标签近期互动列表」视频对象中的 stat。
type GetZoneVideoListWithTagItemStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetZoneVideoListWithTagItemDimension 是「获取分区标签近期互动列表」视频对象中的 dimension。
type GetZoneVideoListWithTagItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetZoneVideoListWithTagItemPage 是「获取分区标签近期互动列表」视频对象中的 pages 数组元素。
type GetZoneVideoListWithTagItemPage struct {
	Cid        int                                  `json:"cid"`         // 当前分P cid
	Page       int                                  `json:"page"`        // 当前分P
	From       string                               `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                               `json:"part"`        // 当前分P标题
	Duration   int                                  `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                               `json:"vid"`         // 站外视频vid
	Weblink    string                               `json:"weblink"`     // 站外视频跳转url
	Dimension  GetZoneVideoListWithTagItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                               `json:"first_frame"` // 分P封面
}

// GetZoneVideoListWithTagItemSubtitle 是「获取分区标签近期互动列表」视频对象中的 subtitle。
type GetZoneVideoListWithTagItemSubtitle struct {
	ID          int                                       `json:"id"`           // 字幕id
	Lan         string                                    `json:"lan"`          // 字幕语言
	LanDoc      string                                    `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                                      `json:"is_lock"`      // 是否锁定
	AuthorMid   int                                       `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string                                    `json:"subtitle_url"` // json格式字幕文件url
	Author      GetZoneVideoListWithTagItemSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// GetZoneVideoListWithTagItemSubtitleAuthor 是「获取分区标签近期互动列表」视频对象中的 subtitle.author。
type GetZoneVideoListWithTagItemSubtitleAuthor struct {
	Mid           int    `json:"mid"`             // 字幕上传者mid
	Name          string `json:"name"`            // 字幕上传者昵称
	Sex           string `json:"sex"`             // 字幕上传者性别。男 女 保密
	Face          string `json:"face"`            // 字幕上传者头像url
	Sign          string `json:"sign"`            // 字幕上传者签名
	Rank          int    `json:"rank"`            // 10000。作用尚不明确
	Birthday      int    `json:"birthday"`        // 0。作用尚不明确
	IsFakeAccount int    `json:"is_fake_account"` // 0。作用尚不明确
	IsDeleted     int    `json:"is_deleted"`      // 0。作用尚不明确
}

// GetZoneVideoListWithTagItemStaff 是「获取分区标签近期互动列表」视频对象中的 staff 数组元素。
type GetZoneVideoListWithTagItemStaff struct {
	Mid        int                                      `json:"mid"`      // 成员mid
	Title      string                                   `json:"title"`    // 成员名称
	Name       string                                   `json:"name"`     // 成员昵称
	Face       string                                   `json:"face"`     // 成员头像url
	Vip        GetZoneVideoListWithTagItemStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetZoneVideoListWithTagItemStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                                      `json:"follower"` // 成员粉丝数
	LabelStyle int                                      `json:"label_style"`
}

// GetZoneVideoListWithTagItemStaffVip 是「获取分区标签近期互动列表」视频对象中的 staff[].vip。
type GetZoneVideoListWithTagItemStaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

// GetZoneVideoListWithTagItemStaffOfficial 是「获取分区标签近期互动列表」视频对象中的 staff[].official。
type GetZoneVideoListWithTagItemStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetZoneVideoListWithTagItemUserGarb 是「获取分区标签近期互动列表」视频对象中的 user_garb。
type GetZoneVideoListWithTagItemUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetZoneVideoListWithTagItemHonorReply 是「获取分区标签近期互动列表」视频对象中的 honor_reply。
type GetZoneVideoListWithTagItemHonorReply struct {
	Honor []GetZoneVideoListWithTagItemHonor `json:"honor"`
}

// GetZoneVideoListWithTagItemHonor 是「获取分区标签近期互动列表」视频对象中的 honor_reply.honor 数组元素。
type GetZoneVideoListWithTagItemHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetZoneVideoListWithTagItemArgueInfo 是「获取分区标签近期互动列表」视频对象中的 argue_info。
type GetZoneVideoListWithTagItemArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetZoneVideoListWithTagItemAiRcmd 是「获取分区标签近期互动列表」视频对象中的 ai_rcmd。
type GetZoneVideoListWithTagItemAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetZoneVideoListWithTagItemUgcSeason 是「获取分区标签近期互动列表」视频对象中的 ugc_season。
type GetZoneVideoListWithTagItemUgcSeason struct {
	Id          int                                           `json:"id"`            // 视频合集id
	Title       string                                        `json:"title"`         // 视频合集标题
	Cover       string                                        `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                                           `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                                        `json:"intro"`         // 视频合集介绍
	SignState   int                                           `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                           `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []GetZoneVideoListWithTagItemUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetZoneVideoListWithTagItemUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                           `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                           `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                          `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                           `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetZoneVideoListWithTagItemUgcSeasonStat 是「获取分区标签近期互动列表」视频对象中的 ugc_season.stat。
type GetZoneVideoListWithTagItemUgcSeasonStat struct {
	SeasonID int `json:"season_id"` // 视频合集id
	View     int `json:"view"`      // 视频合集总浏览量
	Danmaku  int `json:"danmaku"`   // 视频合集总弹幕量
	Reply    int `json:"reply"`     // 视频合集总评论量
	Fav      int `json:"fav"`       // 视频合集总收藏数
	Coin     int `json:"coin"`      // 视频合集总投币数
	Share    int `json:"share"`     // 视频合集总分享数
	NowRank  int `json:"now_rank"`  // 视频合集当前排名
	HisRank  int `json:"his_rank"`  // 视频合集历史排名
	Like     int `json:"like"`      // 视频合集总获赞数
	Vt       int `json:"vt"`        // （？）。作用尚不明确
	Vv       int `json:"vv"`        // （？）。作用尚不明确
}

// GetZoneVideoListWithTagItemUgcSeasonSection 是「获取分区标签近期互动列表」视频对象中的 ugc_season.sections 数组元素。
type GetZoneVideoListWithTagItemUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                                           `json:"id"`
	SectionID int                                           `json:"section_id"`
	Title     string                                        `json:"title"`    // 视频合集中分部标题
	Type      int                                           `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetZoneVideoListWithTagItemUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetZoneVideoListWithTagItemUgcSeasonEpisode 是「获取分区标签近期互动列表」视频对象中的 ugc_season.sections[].episodes 数组元素。
type GetZoneVideoListWithTagItemUgcSeasonEpisode struct {
	SeasonID  int                                            `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                            `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                            `json:"id"`         // 分部中视频id
	Aid       int                                            `json:"aid"`        // 视频aid
	Cid       int                                            `json:"cid"`        // 视频cid
	Title     string                                         `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                                            `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetZoneVideoListWithTagItemUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetZoneVideoListWithTagItemPage                `json:"page"`       // 视频分P信息
	Bvid      string                                         `json:"bvid"`       // 视频bvid
	Pages     []GetZoneVideoListWithTagItemPage              `json:"pages"`      // 视频分P列表
}

// GetZoneVideoListWithTagItemUgcSeasonEpisodeArc 是「获取分区标签近期互动列表」视频对象中的 ugc_season.sections[].episodes[].arc。
type GetZoneVideoListWithTagItemUgcSeasonEpisodeArc struct {
	Aid                int                                             `json:"aid"`                  // 稿件avid
	Videos             int                                             `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                                             `json:"type_id"`              // 分区tid
	TypeName           string                                          `json:"type_name"`            // 子分区名称
	Copyright          int                                             `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string                                          `json:"pic"`                  // 稿件封面图片url
	Title              string                                          `json:"title"`                // 稿件标题
	PubDate            int                                             `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                                             `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string                                          `json:"desc"`                 // 视频简介
	State              int                                             `json:"state"`                // 视频状态
	Duration           int                                             `json:"duration"`             // 稿件总时长。单位为秒
	Rights             GetZoneVideoListWithTagItemRights               `json:"rights"`               // 视频属性标志
	Author             GetZoneVideoListWithTagItemOwner                `json:"author"`               // UP主信息
	Stat               GetZoneVideoListWithTagItemUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                          `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetZoneVideoListWithTagItemDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                             `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                            `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                            `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                             `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                          `json:"vt_display"`           // （？）。作用尚不明确
}

// GetZoneVideoListWithTagItemUgcSeasonEpisodeStat 是「获取分区标签近期互动列表」视频对象中的 ugc_season.sections[].episodes[].arc.stat。
type GetZoneVideoListWithTagItemUgcSeasonEpisodeStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Fav        int    `json:"fav"`        // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	ArgueMsg   string `json:"argue_msg"`  // 警告信息
	Vt         int    `json:"vt"`         // （？）。作用尚不明确
	Vv         int    `json:"vv"`         // （？）。作用尚不明确
}

// GetZoneVideoListRecentResult 是「获取分区近期投稿列表」的 data 对象。
type GetZoneVideoListRecentResult struct {
	Archives []GetZoneVideoListRecentItem `json:"archives"` // 视频列表
	Page     GetZoneVideoListRecentPage   `json:"page"`     // 页面信息
}

// GetZoneVideoListRecentPage 是「获取分区近期投稿列表」的分页信息。
//
// 字段名为 num/size，与空间投稿的 pn/ps、合集的 page_num/page_size 不同，
// 因此保留独立类型，不做跨接口合并。
type GetZoneVideoListRecentPage struct {
	Count int `json:"count"` // 总计视频数
	Num   int `json:"num"`   // 当前页码
	Size  int `json:"size"`  // 每页项数
}

// GetZoneVideoListRecentItem 是「获取分区近期投稿列表」列表中的视频对象。
//
// 字段集同根包的 VideoInfo，但按接口独立定义，不与其它接口共用。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的视频对象结构重复是设计使然
type GetZoneVideoListRecentItem struct {
	Bvid                    string                               `json:"bvid"`         // 稿件bvid
	Aid                     int                                  `json:"aid"`          // 稿件avid
	Videos                  int                                  `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                                  `json:"tid"`          // 分区tid
	Tname                   string                               `json:"tname"`        // 子分区名称
	Copyright               int                                  `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                               `json:"pic"`          // 稿件封面图片url
	Title                   string                               `json:"title"`        // 稿件标题
	PubDate                 int                                  `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                                  `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                               `json:"desc"`         // 视频简介
	DescV2                  []GetZoneVideoListRecentItemDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                                  `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int                                  `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                                  `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                                  `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                               `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetZoneVideoListRecentItemRights     `json:"rights"`       // 视频属性标志
	Owner                   GetZoneVideoListRecentItemOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetZoneVideoListRecentItemStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                               `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                                  `json:"cid"`          // 视频1P cid
	Dimension               GetZoneVideoListRecentItemDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                                  `json:"season_id"`    // 合集id
	Premiere                any                                  `json:"premiere"`     // null
	TeenageMode             int                                  `json:"teenage_mode"`
	IsChargeableSeason      bool                                 `json:"is_chargeable_season"`
	IsStory                 bool                                 `json:"is_story"`
	NoCache                 bool                                 `json:"no_cache"` // 作用尚不明确
	Pages                   []GetZoneVideoListRecentItemPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetZoneVideoListRecentItemSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetZoneVideoListRecentItemStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                                 `json:"is_season_display"`
	UserGarb                GetZoneVideoListRecentItemUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetZoneVideoListRecentItemHonorReply `json:"honor_reply"`
	LikeIcon                string                               `json:"like_icon"`
	ArgueInfo               GetZoneVideoListRecentItemArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                                  `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                               `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                                  `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                               `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                                  `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                               `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                                  `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                                  `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                                 `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                                  `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetZoneVideoListRecentItemAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                                  `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string                               `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                                 `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                                 `json:"is_upower_play"`              //
	IsUpowerPreview         bool                                 `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                                 `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                                 `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetZoneVideoListRecentItemUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                                 `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                                 `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                                 `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                                 `json:"is_view_self"`                // 是否尽自己可见
}

// GetZoneVideoListRecentItemDesc 是「获取分区近期投稿列表」视频对象中的 desc_v2 数组元素。
type GetZoneVideoListRecentItemDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetZoneVideoListRecentItemRights 是「获取分区近期投稿列表」视频对象中的 rights。
type GetZoneVideoListRecentItemRights struct {
	Bp            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	Hd5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UgcPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UgcPayPreview int `json:"ugc_pay_preview"` // 0。作用尚不明确
	NoBackground  int `json:"no_background"`   // 0。作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0。作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0。作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0。作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0。作用尚不明确
}

// GetZoneVideoListRecentItemOwner 是「获取分区近期投稿列表」视频对象中的 owner。
type GetZoneVideoListRecentItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetZoneVideoListRecentItemStat 是「获取分区近期投稿列表」视频对象中的 stat。
type GetZoneVideoListRecentItemStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetZoneVideoListRecentItemDimension 是「获取分区近期投稿列表」视频对象中的 dimension。
type GetZoneVideoListRecentItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetZoneVideoListRecentItemPage 是「获取分区近期投稿列表」视频对象中的 pages 数组元素。
type GetZoneVideoListRecentItemPage struct {
	Cid        int                                 `json:"cid"`         // 当前分P cid
	Page       int                                 `json:"page"`        // 当前分P
	From       string                              `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                              `json:"part"`        // 当前分P标题
	Duration   int                                 `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                              `json:"vid"`         // 站外视频vid
	Weblink    string                              `json:"weblink"`     // 站外视频跳转url
	Dimension  GetZoneVideoListRecentItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                              `json:"first_frame"` // 分P封面
}

// GetZoneVideoListRecentItemSubtitle 是「获取分区近期投稿列表」视频对象中的 subtitle。
type GetZoneVideoListRecentItemSubtitle struct {
	ID          int                                      `json:"id"`           // 字幕id
	Lan         string                                   `json:"lan"`          // 字幕语言
	LanDoc      string                                   `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                                     `json:"is_lock"`      // 是否锁定
	AuthorMid   int                                      `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string                                   `json:"subtitle_url"` // json格式字幕文件url
	Author      GetZoneVideoListRecentItemSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// GetZoneVideoListRecentItemSubtitleAuthor 是「获取分区近期投稿列表」视频对象中的 subtitle.author。
type GetZoneVideoListRecentItemSubtitleAuthor struct {
	Mid           int    `json:"mid"`             // 字幕上传者mid
	Name          string `json:"name"`            // 字幕上传者昵称
	Sex           string `json:"sex"`             // 字幕上传者性别。男 女 保密
	Face          string `json:"face"`            // 字幕上传者头像url
	Sign          string `json:"sign"`            // 字幕上传者签名
	Rank          int    `json:"rank"`            // 10000。作用尚不明确
	Birthday      int    `json:"birthday"`        // 0。作用尚不明确
	IsFakeAccount int    `json:"is_fake_account"` // 0。作用尚不明确
	IsDeleted     int    `json:"is_deleted"`      // 0。作用尚不明确
}

// GetZoneVideoListRecentItemStaff 是「获取分区近期投稿列表」视频对象中的 staff 数组元素。
type GetZoneVideoListRecentItemStaff struct {
	Mid        int                                     `json:"mid"`      // 成员mid
	Title      string                                  `json:"title"`    // 成员名称
	Name       string                                  `json:"name"`     // 成员昵称
	Face       string                                  `json:"face"`     // 成员头像url
	Vip        GetZoneVideoListRecentItemStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetZoneVideoListRecentItemStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                                     `json:"follower"` // 成员粉丝数
	LabelStyle int                                     `json:"label_style"`
}

// GetZoneVideoListRecentItemStaffVip 是「获取分区近期投稿列表」视频对象中的 staff[].vip。
type GetZoneVideoListRecentItemStaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

// GetZoneVideoListRecentItemStaffOfficial 是「获取分区近期投稿列表」视频对象中的 staff[].official。
type GetZoneVideoListRecentItemStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetZoneVideoListRecentItemUserGarb 是「获取分区近期投稿列表」视频对象中的 user_garb。
type GetZoneVideoListRecentItemUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetZoneVideoListRecentItemHonorReply 是「获取分区近期投稿列表」视频对象中的 honor_reply。
type GetZoneVideoListRecentItemHonorReply struct {
	Honor []GetZoneVideoListRecentItemHonor `json:"honor"`
}

// GetZoneVideoListRecentItemHonor 是「获取分区近期投稿列表」视频对象中的 honor_reply.honor 数组元素。
type GetZoneVideoListRecentItemHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetZoneVideoListRecentItemArgueInfo 是「获取分区近期投稿列表」视频对象中的 argue_info。
type GetZoneVideoListRecentItemArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetZoneVideoListRecentItemAiRcmd 是「获取分区近期投稿列表」视频对象中的 ai_rcmd。
type GetZoneVideoListRecentItemAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetZoneVideoListRecentItemUgcSeason 是「获取分区近期投稿列表」视频对象中的 ugc_season。
type GetZoneVideoListRecentItemUgcSeason struct {
	Id          int                                          `json:"id"`            // 视频合集id
	Title       string                                       `json:"title"`         // 视频合集标题
	Cover       string                                       `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                                          `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                                       `json:"intro"`         // 视频合集介绍
	SignState   int                                          `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                          `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []GetZoneVideoListRecentItemUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetZoneVideoListRecentItemUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                          `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                          `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                         `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                          `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetZoneVideoListRecentItemUgcSeasonStat 是「获取分区近期投稿列表」视频对象中的 ugc_season.stat。
type GetZoneVideoListRecentItemUgcSeasonStat struct {
	SeasonID int `json:"season_id"` // 视频合集id
	View     int `json:"view"`      // 视频合集总浏览量
	Danmaku  int `json:"danmaku"`   // 视频合集总弹幕量
	Reply    int `json:"reply"`     // 视频合集总评论量
	Fav      int `json:"fav"`       // 视频合集总收藏数
	Coin     int `json:"coin"`      // 视频合集总投币数
	Share    int `json:"share"`     // 视频合集总分享数
	NowRank  int `json:"now_rank"`  // 视频合集当前排名
	HisRank  int `json:"his_rank"`  // 视频合集历史排名
	Like     int `json:"like"`      // 视频合集总获赞数
	Vt       int `json:"vt"`        // （？）。作用尚不明确
	Vv       int `json:"vv"`        // （？）。作用尚不明确
}

// GetZoneVideoListRecentItemUgcSeasonSection 是「获取分区近期投稿列表」视频对象中的 ugc_season.sections 数组元素。
type GetZoneVideoListRecentItemUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                                          `json:"id"`
	SectionID int                                          `json:"section_id"`
	Title     string                                       `json:"title"`    // 视频合集中分部标题
	Type      int                                          `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetZoneVideoListRecentItemUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetZoneVideoListRecentItemUgcSeasonEpisode 是「获取分区近期投稿列表」视频对象中的 ugc_season.sections[].episodes 数组元素。
type GetZoneVideoListRecentItemUgcSeasonEpisode struct {
	SeasonID  int                                           `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                           `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                           `json:"id"`         // 分部中视频id
	Aid       int                                           `json:"aid"`        // 视频aid
	Cid       int                                           `json:"cid"`        // 视频cid
	Title     string                                        `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                                           `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetZoneVideoListRecentItemUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetZoneVideoListRecentItemPage                `json:"page"`       // 视频分P信息
	Bvid      string                                        `json:"bvid"`       // 视频bvid
	Pages     []GetZoneVideoListRecentItemPage              `json:"pages"`      // 视频分P列表
}

// GetZoneVideoListRecentItemUgcSeasonEpisodeArc 是「获取分区近期投稿列表」视频对象中的 ugc_season.sections[].episodes[].arc。
type GetZoneVideoListRecentItemUgcSeasonEpisodeArc struct {
	Aid                int                                            `json:"aid"`                  // 稿件avid
	Videos             int                                            `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                                            `json:"type_id"`              // 分区tid
	TypeName           string                                         `json:"type_name"`            // 子分区名称
	Copyright          int                                            `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string                                         `json:"pic"`                  // 稿件封面图片url
	Title              string                                         `json:"title"`                // 稿件标题
	PubDate            int                                            `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                                            `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string                                         `json:"desc"`                 // 视频简介
	State              int                                            `json:"state"`                // 视频状态
	Duration           int                                            `json:"duration"`             // 稿件总时长。单位为秒
	Rights             GetZoneVideoListRecentItemRights               `json:"rights"`               // 视频属性标志
	Author             GetZoneVideoListRecentItemOwner                `json:"author"`               // UP主信息
	Stat               GetZoneVideoListRecentItemUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                         `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetZoneVideoListRecentItemDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                            `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                           `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                           `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                            `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                         `json:"vt_display"`           // （？）。作用尚不明确
}

// GetZoneVideoListRecentItemUgcSeasonEpisodeStat 是「获取分区近期投稿列表」视频对象中的 ugc_season.sections[].episodes[].arc.stat。
type GetZoneVideoListRecentItemUgcSeasonEpisodeStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Fav        int    `json:"fav"`        // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	ArgueMsg   string `json:"argue_msg"`  // 警告信息
	Vt         int    `json:"vt"`         // （？）。作用尚不明确
	Vv         int    `json:"vv"`         // （？）。作用尚不明确
}

// GetZoneVideoListByOrderResult 是「获取分区近期投稿列表（带排序）」的 data 对象。
type GetZoneVideoListByOrderResult struct {
	ExpList        *string                       `json:"exp_list"`         // 作用尚不明确
	ShowModuleList []string                      `json:"show_module_list"` // 显示模块列表?
	Result         []GetZoneVideoListByOrderItem `json:"result"`           // 结果本体。失败时为null
	ShowColumn     int                           `json:"show_column"`      // 0。作用尚不明确
	RqtType        string                        `json:"rqt_type"`         // search。作用尚不明确
	NumPages       int                           `json:"numPages"`         // 页码。失败时为0
	NumResults     int                           `json:"numResults"`       // 视频数。失败时为0
	CrrQuery       *string                       `json:"crr_query"`        // 空。作用尚不明确
	PageSize       int                           `json:"pagesize"`         // 视频数
	SuggestKeyword *string                       `json:"suggest_keyword"`  // 空。作用尚不明确
	EggInfo        *string                       `json:"egg_info"`         // 作用尚不明确
	Cache          int                           `json:"cache"`            // 0。作用尚不明确
	ExpBits        int                           `json:"exp_bits"`         // 1。作用尚不明确
	ExpStr         *string                       `json:"exp_str"`          // 空。作用尚不明确
	Seid           string                        `json:"seid"`             // 一串字符串中的数字。作用尚不明确
	Msg            string                        `json:"msg"`              // 结果信息。成功时为success, 反之为as error.
	EggHit         int                           `json:"egg_hit"`          // 0。作用尚不明确
	Page           int                           `json:"page"`             // 页码
}

// GetZoneVideoListByOrderItem 是「获取分区近期投稿列表（带排序）」result 数组中的对象。
type GetZoneVideoListByOrderItem struct {
	PubDate      string `json:"pubdate"`        // 发布时间。格式为 yyyy-MM-dd HH:mm:ss
	Pic          string `json:"pic"`            // 封面图
	Tag          string `json:"tag"`            // 标签。用 , 分隔
	Duration     int    `json:"duration"`       // 时长。单位为秒
	ID           int    `json:"id"`             // aid
	RankScore    int    `json:"rank_score"`     // 排序分数?
	BadgePay     bool   `json:"badgepay"`       // 是否有角标?
	SendDate     int    `json:"senddate"`       // 发送时间?。UNIX 秒级时间戳
	Author       string `json:"author"`         // UP主名
	Review       int    `json:"review"`         // 评论数
	Mid          int    `json:"mid"`            // UP主mid
	IsUnionVideo int    `json:"is_union_video"` // 是否为联合投稿
	RankIndex    int    `json:"rank_index"`     // 排序索引号
	Type         string `json:"type"`           // 类型。video: 视频
	ArcRank      string `json:"arcrank"`        // 0。作用尚不明确
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
