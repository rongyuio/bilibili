package videoranking

// 热门视频与每周必看接口的响应模型。
//
// 每个接口的返回体独立定义，嵌套对象也不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// GetPopularVideoListResult 是「获取当前热门视频列表」的 data 对象。
type GetPopularVideoListResult struct {
	List   []GetPopularVideoListItem `json:"list"`    // 视频列表。基本同「获取视频详细信息（web端）」中的 data 对象
	NoMore bool                      `json:"no_more"` // 是否有更多数据。false：下页还有数据。true：下页没有数据
}

// GetPopularVideoListItem 是「获取当前热门视频列表」列表中的视频对象。
//
// 字段集同根包的 VideoInfo，但按接口独立定义，不与其它接口共用。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的视频对象结构重复是设计使然
type GetPopularVideoListItem struct {
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
	DescV2                  []GetPopularVideoListItemDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                               `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int                               `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                               `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                               `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                            `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetPopularVideoListItemRights     `json:"rights"`       // 视频属性标志
	Owner                   GetPopularVideoListItemOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetPopularVideoListItemStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                            `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                               `json:"cid"`          // 视频1P cid
	Dimension               GetPopularVideoListItemDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                               `json:"season_id"`    // 合集id
	Premiere                any                               `json:"premiere"`     // null
	TeenageMode             int                               `json:"teenage_mode"`
	IsChargeableSeason      bool                              `json:"is_chargeable_season"`
	IsStory                 bool                              `json:"is_story"`
	NoCache                 bool                              `json:"no_cache"` // 作用尚不明确
	Pages                   []GetPopularVideoListItemPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetPopularVideoListItemSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetPopularVideoListItemStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                              `json:"is_season_display"`
	UserGarb                GetPopularVideoListItemUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetPopularVideoListItemHonorReply `json:"honor_reply"`
	LikeIcon                string                            `json:"like_icon"`
	ArgueInfo               GetPopularVideoListItemArgueInfo  `json:"argue_info"`                  // 争议/警告信息
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
	AiRcmd                  GetPopularVideoListItemAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                               `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string                            `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                              `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                              `json:"is_upower_play"`              //
	IsUpowerPreview         bool                              `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                              `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                              `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetPopularVideoListItemUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                              `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                              `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                              `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                              `json:"is_view_self"`                // 是否尽自己可见
}

// GetPopularVideoListItemDesc 是「获取当前热门视频列表」视频对象中的 desc_v2 数组元素。
type GetPopularVideoListItemDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetPopularVideoListItemRights 是「获取当前热门视频列表」视频对象中的 rights。
type GetPopularVideoListItemRights struct {
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

// GetPopularVideoListItemOwner 是「获取当前热门视频列表」视频对象中的 owner。
type GetPopularVideoListItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetPopularVideoListItemStat 是「获取当前热门视频列表」视频对象中的 stat。
type GetPopularVideoListItemStat struct {
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

// GetPopularVideoListItemDimension 是「获取当前热门视频列表」视频对象中的 dimension。
type GetPopularVideoListItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetPopularVideoListItemPage 是「获取当前热门视频列表」视频对象中的 pages 数组元素。
type GetPopularVideoListItemPage struct {
	Cid        int                              `json:"cid"`         // 当前分P cid
	Page       int                              `json:"page"`        // 当前分P
	From       string                           `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                           `json:"part"`        // 当前分P标题
	Duration   int                              `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                           `json:"vid"`         // 站外视频vid
	Weblink    string                           `json:"weblink"`     // 站外视频跳转url
	Dimension  GetPopularVideoListItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                           `json:"first_frame"` // 分P封面
}

// GetPopularVideoListItemSubtitle 是「获取当前热门视频列表」视频对象中的 subtitle。
type GetPopularVideoListItemSubtitle struct {
	ID          int                                   `json:"id"`           // 字幕id
	Lan         string                                `json:"lan"`          // 字幕语言
	LanDoc      string                                `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                                  `json:"is_lock"`      // 是否锁定
	AuthorMid   int                                   `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string                                `json:"subtitle_url"` // json格式字幕文件url
	Author      GetPopularVideoListItemSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// GetPopularVideoListItemSubtitleAuthor 是「获取当前热门视频列表」视频对象中的 subtitle.author。
type GetPopularVideoListItemSubtitleAuthor struct {
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

// GetPopularVideoListItemStaff 是「获取当前热门视频列表」视频对象中的 staff 数组元素。
type GetPopularVideoListItemStaff struct {
	Mid        int                                  `json:"mid"`      // 成员mid
	Title      string                               `json:"title"`    // 成员名称
	Name       string                               `json:"name"`     // 成员昵称
	Face       string                               `json:"face"`     // 成员头像url
	Vip        GetPopularVideoListItemStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetPopularVideoListItemStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                                  `json:"follower"` // 成员粉丝数
	LabelStyle int                                  `json:"label_style"`
}

// GetPopularVideoListItemStaffVip 是「获取当前热门视频列表」视频对象中的 staff[].vip。
type GetPopularVideoListItemStaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

// GetPopularVideoListItemStaffOfficial 是「获取当前热门视频列表」视频对象中的 staff[].official。
type GetPopularVideoListItemStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetPopularVideoListItemUserGarb 是「获取当前热门视频列表」视频对象中的 user_garb。
type GetPopularVideoListItemUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetPopularVideoListItemHonorReply 是「获取当前热门视频列表」视频对象中的 honor_reply。
type GetPopularVideoListItemHonorReply struct {
	Honor []GetPopularVideoListItemHonor `json:"honor"`
}

// GetPopularVideoListItemHonor 是「获取当前热门视频列表」视频对象中的 honor_reply.honor 数组元素。
type GetPopularVideoListItemHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetPopularVideoListItemArgueInfo 是「获取当前热门视频列表」视频对象中的 argue_info。
type GetPopularVideoListItemArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetPopularVideoListItemAiRcmd 是「获取当前热门视频列表」视频对象中的 ai_rcmd。
type GetPopularVideoListItemAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetPopularVideoListItemUgcSeason 是「获取当前热门视频列表」视频对象中的 ugc_season。
type GetPopularVideoListItemUgcSeason struct {
	Id          int                                       `json:"id"`            // 视频合集id
	Title       string                                    `json:"title"`         // 视频合集标题
	Cover       string                                    `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                                       `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                                    `json:"intro"`         // 视频合集介绍
	SignState   int                                       `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                       `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []GetPopularVideoListItemUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetPopularVideoListItemUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                       `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                       `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                      `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                       `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetPopularVideoListItemUgcSeasonStat 是「获取当前热门视频列表」视频对象中的 ugc_season.stat。
type GetPopularVideoListItemUgcSeasonStat struct {
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

// GetPopularVideoListItemUgcSeasonSection 是「获取当前热门视频列表」视频对象中的 ugc_season.sections 数组元素。
type GetPopularVideoListItemUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                                       `json:"id"`
	SectionID int                                       `json:"section_id"`
	Title     string                                    `json:"title"`    // 视频合集中分部标题
	Type      int                                       `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetPopularVideoListItemUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetPopularVideoListItemUgcSeasonEpisode 是「获取当前热门视频列表」视频对象中的 ugc_season.sections[].episodes 数组元素。
type GetPopularVideoListItemUgcSeasonEpisode struct {
	SeasonID  int                                        `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                        `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                        `json:"id"`         // 分部中视频id
	Aid       int                                        `json:"aid"`        // 视频aid
	Cid       int                                        `json:"cid"`        // 视频cid
	Title     string                                     `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                                        `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetPopularVideoListItemUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetPopularVideoListItemPage                `json:"page"`       // 视频分P信息
	Bvid      string                                     `json:"bvid"`       // 视频bvid
	Pages     []GetPopularVideoListItemPage              `json:"pages"`      // 视频分P列表
}

// GetPopularVideoListItemUgcSeasonEpisodeArc 是「获取当前热门视频列表」视频对象中的 ugc_season.sections[].episodes[].arc。
type GetPopularVideoListItemUgcSeasonEpisodeArc struct {
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
	Rights             GetPopularVideoListItemRights               `json:"rights"`               // 视频属性标志
	Author             GetPopularVideoListItemOwner                `json:"author"`               // UP主信息
	Stat               GetPopularVideoListItemUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                      `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetPopularVideoListItemDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                         `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                        `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                        `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                         `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                      `json:"vt_display"`           // （？）。作用尚不明确
}

// GetPopularVideoListItemUgcSeasonEpisodeStat 是「获取当前热门视频列表」视频对象中的 ugc_season.sections[].episodes[].arc.stat。
type GetPopularVideoListItemUgcSeasonEpisodeStat struct {
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

// GetWeeklyPopularListResult 是「每周必看全部列表」的 data 对象。
type GetWeeklyPopularListResult struct {
	List []GetWeeklyPopularListItem `json:"list"` // 全部期数列表，按期数从大到小排列
}

// GetWeeklyPopularListItem 是「每周必看全部列表」中的一期。
type GetWeeklyPopularListItem struct {
	Number  int    `json:"number"`  // 期数
	Subject string `json:"subject"` // 主题
	Status  int    `json:"status"`  // 状态。2：已结束
	Name    string `json:"name"`    // 名称。yyyy第n期 MM.dd - MM.dd
}

// GetWeeklyPopularDetailResult 是「每周必看选期详细信息」的 data 对象。
type GetWeeklyPopularDetailResult struct {
	Config   GetWeeklyPopularDetailConfig `json:"config"`   // 选期信息
	Reminder string                       `json:"reminder"` // 提醒。如「每周五晚18:00更新」
	List     []GetWeeklyPopularDetailItem `json:"list"`     // 选期视频列表。基本同「获取视频详细信息（web端）」中的 data 对象
}

// GetWeeklyPopularDetailConfig 是「每周必看选期详细信息」的 config 对象。
type GetWeeklyPopularDetailConfig struct {
	ID            int    `json:"id"`             // 选期ID
	Type          string `json:"type"`           // 选期类型。如 weekly_selected
	Number        int    `json:"number"`         // 期数
	Subject       string `json:"subject"`        // 主题
	Stime         int    `json:"stime"`          // 开始时间。秒级时间戳
	Etime         int    `json:"etime"`          // 结束时间。秒级时间戳
	Status        int    `json:"status"`         // 状态。2：已结束
	Name          string `json:"name"`           // 名称。yyyy第n期 MM.dd - MM.dd
	Label         string `json:"label"`          // 标题
	Hint          string `json:"hint"`           // 提示
	Color         int    `json:"color"`          // 颜色?。作用尚不明确
	Cover         string `json:"cover"`          // 封面url
	ShareTitle    string `json:"share_title"`    // 分享标题
	ShareSubtitle string `json:"share_subtitle"` // 分享副标题
	MediaID       int    `json:"media_id"`       // 媒体ID?。作用尚不明确
}

// GetWeeklyPopularDetailItem 是「每周必看选期详细信息」列表中的视频对象。
//
// 字段集同根包的 VideoInfo，但按接口独立定义，不与其它接口共用。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的视频对象结构重复是设计使然
type GetWeeklyPopularDetailItem struct {
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
	DescV2                  []GetWeeklyPopularDetailItemDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                                  `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int                                  `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                                  `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                                  `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                               `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetWeeklyPopularDetailItemRights     `json:"rights"`       // 视频属性标志
	Owner                   GetWeeklyPopularDetailItemOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetWeeklyPopularDetailItemStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                               `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                                  `json:"cid"`          // 视频1P cid
	Dimension               GetWeeklyPopularDetailItemDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                                  `json:"season_id"`    // 合集id
	Premiere                any                                  `json:"premiere"`     // null
	TeenageMode             int                                  `json:"teenage_mode"`
	IsChargeableSeason      bool                                 `json:"is_chargeable_season"`
	IsStory                 bool                                 `json:"is_story"`
	NoCache                 bool                                 `json:"no_cache"` // 作用尚不明确
	Pages                   []GetWeeklyPopularDetailItemPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetWeeklyPopularDetailItemSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetWeeklyPopularDetailItemStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                                 `json:"is_season_display"`
	UserGarb                GetWeeklyPopularDetailItemUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetWeeklyPopularDetailItemHonorReply `json:"honor_reply"`
	LikeIcon                string                               `json:"like_icon"`
	ArgueInfo               GetWeeklyPopularDetailItemArgueInfo  `json:"argue_info"`                  // 争议/警告信息
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
	AiRcmd                  GetWeeklyPopularDetailItemAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                                  `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string                               `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                                 `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                                 `json:"is_upower_play"`              //
	IsUpowerPreview         bool                                 `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                                 `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                                 `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetWeeklyPopularDetailItemUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                                 `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                                 `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                                 `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                                 `json:"is_view_self"`                // 是否尽自己可见
}

// GetWeeklyPopularDetailItemDesc 是「每周必看选期详细信息」视频对象中的 desc_v2 数组元素。
type GetWeeklyPopularDetailItemDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetWeeklyPopularDetailItemRights 是「每周必看选期详细信息」视频对象中的 rights。
type GetWeeklyPopularDetailItemRights struct {
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

// GetWeeklyPopularDetailItemOwner 是「每周必看选期详细信息」视频对象中的 owner。
type GetWeeklyPopularDetailItemOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetWeeklyPopularDetailItemStat 是「每周必看选期详细信息」视频对象中的 stat。
type GetWeeklyPopularDetailItemStat struct {
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

// GetWeeklyPopularDetailItemDimension 是「每周必看选期详细信息」视频对象中的 dimension。
type GetWeeklyPopularDetailItemDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetWeeklyPopularDetailItemPage 是「每周必看选期详细信息」视频对象中的 pages 数组元素。
type GetWeeklyPopularDetailItemPage struct {
	Cid        int                                 `json:"cid"`         // 当前分P cid
	Page       int                                 `json:"page"`        // 当前分P
	From       string                              `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                              `json:"part"`        // 当前分P标题
	Duration   int                                 `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                              `json:"vid"`         // 站外视频vid
	Weblink    string                              `json:"weblink"`     // 站外视频跳转url
	Dimension  GetWeeklyPopularDetailItemDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                              `json:"first_frame"` // 分P封面
}

// GetWeeklyPopularDetailItemSubtitle 是「每周必看选期详细信息」视频对象中的 subtitle。
type GetWeeklyPopularDetailItemSubtitle struct {
	ID          int                                      `json:"id"`           // 字幕id
	Lan         string                                   `json:"lan"`          // 字幕语言
	LanDoc      string                                   `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                                     `json:"is_lock"`      // 是否锁定
	AuthorMid   int                                      `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string                                   `json:"subtitle_url"` // json格式字幕文件url
	Author      GetWeeklyPopularDetailItemSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

// GetWeeklyPopularDetailItemSubtitleAuthor 是「每周必看选期详细信息」视频对象中的 subtitle.author。
type GetWeeklyPopularDetailItemSubtitleAuthor struct {
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

// GetWeeklyPopularDetailItemStaff 是「每周必看选期详细信息」视频对象中的 staff 数组元素。
type GetWeeklyPopularDetailItemStaff struct {
	Mid        int                                     `json:"mid"`      // 成员mid
	Title      string                                  `json:"title"`    // 成员名称
	Name       string                                  `json:"name"`     // 成员昵称
	Face       string                                  `json:"face"`     // 成员头像url
	Vip        GetWeeklyPopularDetailItemStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetWeeklyPopularDetailItemStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                                     `json:"follower"` // 成员粉丝数
	LabelStyle int                                     `json:"label_style"`
}

// GetWeeklyPopularDetailItemStaffVip 是「每周必看选期详细信息」视频对象中的 staff[].vip。
type GetWeeklyPopularDetailItemStaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

// GetWeeklyPopularDetailItemStaffOfficial 是「每周必看选期详细信息」视频对象中的 staff[].official。
type GetWeeklyPopularDetailItemStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetWeeklyPopularDetailItemUserGarb 是「每周必看选期详细信息」视频对象中的 user_garb。
type GetWeeklyPopularDetailItemUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetWeeklyPopularDetailItemHonorReply 是「每周必看选期详细信息」视频对象中的 honor_reply。
type GetWeeklyPopularDetailItemHonorReply struct {
	Honor []GetWeeklyPopularDetailItemHonor `json:"honor"`
}

// GetWeeklyPopularDetailItemHonor 是「每周必看选期详细信息」视频对象中的 honor_reply.honor 数组元素。
type GetWeeklyPopularDetailItemHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetWeeklyPopularDetailItemArgueInfo 是「每周必看选期详细信息」视频对象中的 argue_info。
type GetWeeklyPopularDetailItemArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetWeeklyPopularDetailItemAiRcmd 是「每周必看选期详细信息」视频对象中的 ai_rcmd。
type GetWeeklyPopularDetailItemAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetWeeklyPopularDetailItemUgcSeason 是「每周必看选期详细信息」视频对象中的 ugc_season。
type GetWeeklyPopularDetailItemUgcSeason struct {
	Id          int                                          `json:"id"`            // 视频合集id
	Title       string                                       `json:"title"`         // 视频合集标题
	Cover       string                                       `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                                          `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                                       `json:"intro"`         // 视频合集介绍
	SignState   int                                          `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                          `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []GetWeeklyPopularDetailItemUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetWeeklyPopularDetailItemUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                          `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                          `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                         `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                          `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetWeeklyPopularDetailItemUgcSeasonStat 是「每周必看选期详细信息」视频对象中的 ugc_season.stat。
type GetWeeklyPopularDetailItemUgcSeasonStat struct {
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

// GetWeeklyPopularDetailItemUgcSeasonSection 是「每周必看选期详细信息」视频对象中的 ugc_season.sections 数组元素。
type GetWeeklyPopularDetailItemUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                                          `json:"id"`
	SectionID int                                          `json:"section_id"`
	Title     string                                       `json:"title"`    // 视频合集中分部标题
	Type      int                                          `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetWeeklyPopularDetailItemUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetWeeklyPopularDetailItemUgcSeasonEpisode 是「每周必看选期详细信息」视频对象中的 ugc_season.sections[].episodes 数组元素。
type GetWeeklyPopularDetailItemUgcSeasonEpisode struct {
	SeasonID  int                                           `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                           `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                           `json:"id"`         // 分部中视频id
	Aid       int                                           `json:"aid"`        // 视频aid
	Cid       int                                           `json:"cid"`        // 视频cid
	Title     string                                        `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                                           `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetWeeklyPopularDetailItemUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetWeeklyPopularDetailItemPage                `json:"page"`       // 视频分P信息
	Bvid      string                                        `json:"bvid"`       // 视频bvid
	Pages     []GetWeeklyPopularDetailItemPage              `json:"pages"`      // 视频分P列表
}

// GetWeeklyPopularDetailItemUgcSeasonEpisodeArc 是「每周必看选期详细信息」视频对象中的 ugc_season.sections[].episodes[].arc。
type GetWeeklyPopularDetailItemUgcSeasonEpisodeArc struct {
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
	Rights             GetWeeklyPopularDetailItemRights               `json:"rights"`               // 视频属性标志
	Author             GetWeeklyPopularDetailItemOwner                `json:"author"`               // UP主信息
	Stat               GetWeeklyPopularDetailItemUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                         `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetWeeklyPopularDetailItemDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                            `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                           `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                           `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                            `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                         `json:"vt_display"`           // （？）。作用尚不明确
}

// GetWeeklyPopularDetailItemUgcSeasonEpisodeStat 是「每周必看选期详细信息」视频对象中的 ugc_season.sections[].episodes[].arc.stat。
type GetWeeklyPopularDetailItemUgcSeasonEpisodeStat struct {
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
