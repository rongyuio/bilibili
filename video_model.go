package bilibili

// 视频相关响应模型。
//
// Owner 是视频与收藏夹共用的作者简要信息（mid/name/face），FavourUpper 为其别名；
// Author、DynamicUpUserInfo 等其余作者类型字段集不同，保留独立定义。

type DescV2 struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

type VideoRights struct {
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

// Owner 是视频稿件与收藏夹内容共用的作者简要信息。
//
// FavourUpper 为本类型的别名；其余业务的作者类型字段集不同，保留独立定义。
type Owner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

type VideoStat struct {
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

type VideoSubtitleAuthor struct {
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

type VideoSubtitle struct {
	ID          int                 `json:"id"`           // 字幕id
	Lan         string              `json:"lan"`          // 字幕语言
	LanDoc      string              `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                `json:"is_lock"`      // 是否锁定
	AuthorMid   int                 `json:"author_mid"`   // 字幕上传者mid
	SubtitleURL string              `json:"subtitle_url"` // json格式字幕文件url
	Author      VideoSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

type StaffVip struct {
	Type      int `json:"type"`       // 成员会员类型。0：无。1：月会员。2：年会员
	Status    int `json:"status"`     // 会员状态。0：无。1：有
	ThemeType int `json:"theme_type"` // 0
}

type Staff struct {
	Mid        int      `json:"mid"`      // 成员mid
	Title      string   `json:"title"`    // 成员名称
	Name       string   `json:"name"`     // 成员昵称
	Face       string   `json:"face"`     // 成员头像url
	Vip        StaffVip `json:"vip"`      // 成员大会员状态
	Official   Official `json:"official"` // 成员认证信息
	Follower   int      `json:"follower"` // 成员粉丝数
	LabelStyle int      `json:"label_style"`
}

type UserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

type Honor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

type HonorReply struct {
	Honor []Honor `json:"honor"`
}

type ArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}
type TopRecommendVideoList struct {
	BusinessCard          any                     `json:"business_card"`            // 无意义
	FloorInfo             any                     `json:"floor_info"`               // 无意义
	Item                  []TopRecommendVideoItem `json:"item"`                     // 推荐列表
	Mid                   int                     `json:"mid"`                      // 用户mid,未登录为0
	PreloadExposePct      float64                 `json:"preload_expose_pct"`       // 用于预加载?
	PreloadFloorExposePct float64                 `json:"preload_floor_expose_pct"` // 用于预加载?
	SideBarColumn         []any                   `json:"side_bar_column"`          // 边栏列表?	可参考字段 item 及对应功能文档
	UserFeature           any                     `json:"user_feature"`             // 无意义
}

// RcmdReason 推荐理由
type RcmdReason struct {
	ReasonType int    `json:"reason_type"` // 原因类型
	Content    string `json:"content"`     // 原因描述（仅当 reason_type 为 3 时存在）
}

type TopRecommendVideoItem struct {
	AvFeature       any        `json:"av_feature"`        // 暂无参考意义
	BusinessInfo    any        `json:"business_info"`     // 商业推广信息，通常为 null
	Bvid            string     `json:"bvid"`              // 视频 bvid
	Cid             int        `json:"cid"`               // 视频 cid
	DislikeSwitch   int        `json:"dislike_switch"`    // 不感兴趣开关
	DislikeSwitchPC int        `json:"dislike_switch_pc"` // PC端不感兴趣开关
	Duration        int        `json:"duration"`          // 视频时长
	EnableVt        int        `json:"enable_vt"`         // 未知作用
	Goto            string     `json:"goto"`              // 目标类型 (av, ogv, live)
	ID              int        `json:"id"`                // 视频 avid / 直播间 id
	IsFollowed      int        `json:"is_followed"`       // 是否已关注
	IsStock         int        `json:"is_stock"`          // 未知作用
	OgvInfo         any        `json:"ogv_info"`          // 通常为 null
	Owner           Owner      `json:"owner"`             // 视频 UP 主信息
	Pic             string     `json:"pic"`               // 视频封面
	Pic43           string     `json:"pic_4_3"`           // 4:3 比例封面
	Pos             int        `json:"pos"`               // 位置
	PubDate         int        `json:"pubdate"`           // 发布时间（秒级时间戳）
	RcmdReason      RcmdReason `json:"rcmd_reason"`       // 推荐理由
	RoomInfo        any        `json:"room_info"`         // 通常为 null
	ShowInfo        int        `json:"show_info"`         // 展示信息（1: 普通视频, 0: 直播）
	Stat            any        `json:"stat"`              // 视频状态信息
	Title           string     `json:"title"`             // 视频标题
	TrackID         string     `json:"track_id"`          // 跟踪标识
	URI             string     `json:"uri"`               // 目标页 URI
	VtDisplay       string     `json:"vt_display"`        // 未知作用
}
type VideoInfo struct {
	Bvid                    string        `json:"bvid"`         // 稿件bvid
	Aid                     int           `json:"aid"`          // 稿件avid
	Videos                  int           `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int           `json:"tid"`          // 分区tid
	Tname                   string        `json:"tname"`        // 子分区名称
	Copyright               int           `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string        `json:"pic"`          // 稿件封面图片url
	Title                   string        `json:"title"`        // 稿件标题
	PubDate                 int           `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int           `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string        `json:"desc"`         // 视频简介
	DescV2                  []DescV2      `json:"desc_v2"`      // 新版视频简介
	State                   int           `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration                int           `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int           `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int           `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string        `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  VideoRights   `json:"rights"`       // 视频属性标志
	Owner                   Owner         `json:"owner"`        // 视频UP主信息
	Stat                    VideoStat     `json:"stat"`         // 视频状态数
	Dynamic                 string        `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int           `json:"cid"`          // 视频1P cid
	Dimension               Dimension     `json:"dimension"`    // 视频1P分辨率
	SeasonID                int           `json:"season_id"`    // 合集id
	Premiere                any           `json:"premiere"`     // null
	TeenageMode             int           `json:"teenage_mode"`
	IsChargeableSeason      bool          `json:"is_chargeable_season"`
	IsStory                 bool          `json:"is_story"`
	NoCache                 bool          `json:"no_cache"` // 作用尚不明确
	Pages                   []VideoPage   `json:"pages"`    // 视频分P列表
	Subtitle                VideoSubtitle `json:"subtitle"` // 视频CC字幕信息
	Staff                   []Staff       `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool          `json:"is_season_display"`
	UserGarb                UserGarb      `json:"user_garb"` // 用户装扮信息
	HonorReply              HonorReply    `json:"honor_reply"`
	LikeIcon                string        `json:"like_icon"`
	ArgueInfo               ArgueInfo     `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int           `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string        `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int           `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string        `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int           `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string        `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int           `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int           `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool          `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int           `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  AiRcmd        `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int           `json:"tid_v2"`                      // 分区tid (v2)。详情见[视频分区一览 (v2)](video_zone_v2.md)
	TnameV2                 string        `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool          `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool          `json:"is_upower_play"`              //
	IsUpowerPreview         bool          `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool          `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool          `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               UgcSeason     `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool          `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool          `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool          `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool          `json:"is_view_self"`                // 是否尽自己可见
}

// AiRcmd AI 推荐信息。`Related[]` 与 `View` 上是同一个结构
type AiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// CardVip 是视频卡片与用户空间共用的大会员信息。
//
// SpaceVip 是本类型的别名；UserCardVip、MyVip、VipUserVip 等字段集或命名
// 不同，保留独立类型。
type CardVip struct {
	Type               int         `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int         `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int         `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int         `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int         `json:"theme_type"`           // 0。作用尚不明确
	Label              Label       `json:"label"`                // 会员标签
	AvatarSubscript    int         `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string      `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int         `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptURL string      `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int         `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int         `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            VipOttInfo  `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           VipSuperVip `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int         `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         AvatarIcon  `json:"avatar_icon"`          // 大会员角标信息
}

// VipOttInfo `vip` 中的 `ott_info` 对象
type VipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// VipSuperVip `vip` 中的 `super_vip` 对象
type VipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// AvatarIcon 大会员角标信息。`vip` 中的 `avatar_icon` 对象。
// 视频卡片与用户信息里的 `vip` 都带这一项
type AvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

type VideoCard struct {
	Mid            string         `json:"mid"`              // 用户mid
	Name           string         `json:"name"`             // 用户昵称
	Approve        bool           `json:"approve"`          // false。作用尚不明确
	Sex            string         `json:"sex"`              // 用户性别。男 女 保密
	Rank           string         `json:"rank"`             // 10000。作用尚不明确
	Face           string         `json:"face"`             // 用户头像链接
	FaceNft        int            `json:"face_nft"`         // 是否为 nft 头像。0不是nft头像。1是 nft 头像
	DisplayRank    string         `json:"DisplayRank"`      // 0。作用尚不明确
	Regtime        int            `json:"regtime"`          // 0。作用尚不明确
	Spacesta       int            `json:"spacesta"`         // 0。作用尚不明确
	Birthday       string         `json:"birthday"`         // 空。作用尚不明确
	Place          string         `json:"place"`            // 空。作用尚不明确
	Description    string         `json:"description"`      // 空。作用尚不明确
	Article        int            `json:"article"`          // 0。作用尚不明确
	Attentions     []any          `json:"attentions"`       // 空。作用尚不明确
	Fans           int            `json:"fans"`             // 粉丝数
	Friend         int            `json:"friend"`           // 关注数
	Attention      int            `json:"attention"`        // 关注数
	Sign           string         `json:"sign"`             // 签名
	LevelInfo      LevelInfo      `json:"level_info"`       // 等级
	Pendant        Pendant        `json:"pendant"`          // 挂件
	Nameplate      Nameplate      `json:"nameplate"`        // 勋章
	Official       Official       `json:"Official"`         // 认证信息
	OfficialVerify OfficialVerify `json:"official_verify"`  // 认证信息2
	Vip            CardVip        `json:"vip"`              // 大会员状态
	IsSeniorMember int            `json:"is_senior_member"` // 是否为硬核会员。0：否。1：是
}

type VideoDetailInfoCard struct {
	Card         VideoCard `json:"card"`          // UP主名片信息
	Space        CardSpace `json:"space"`         // 主页头图
	Following    bool      `json:"following"`     // 是否关注此用户。true：已关注。false：未关注。需要登录(Cookie) 。未登录为false
	ArchiveCount int       `json:"archive_count"` // 用户稿件数
	ArticleCount int       `json:"article_count"` // 用户专栏数
	Follower     int       `json:"follower"`      // 粉丝数
	LikeNum      int       `json:"like_num"`      // UP主获赞次数
}

type VideoDetailInfo struct {
	View                   VideoInfo           `json:"View"`                       // 视频基本信息
	Card                   VideoDetailInfoCard `json:"Card"`                       // 视频UP主信息
	Tags                   []VideoTag          `json:"Tags"`                       // 视频TAG信息
	Reply                  CommentsHotReply    `json:"Reply"`                      // 视频热评信息
	Related                []VideoInfo         `json:"Related"`                    // 推荐视频信息
	Spec                   any                 `json:"Spec"`                       // ？。作用尚不明确
	IsHitLabourDayActivity bool                `json:"is_hit_labour_day_activity"` // （？）。作用尚不明确
	HotShare               any                 `json:"hot_share"`                  // ？。作用尚不明确
	Elec                   any                 `json:"elec"`                       // ？。作用尚不明确
	Recommend              any                 `json:"recommend"`                  // ？。作用尚不明确
	ViewAddit              any                 `json:"view_addit"`                 // ？。作用尚不明确
	Emergency              Emergency           `json:"emergency"`                  // 视频操作按钮信息
	Participle             []string            `json:"participle"`                 // 分词信息。用于推荐
	ReplaceRecommend       bool                `json:"replace_recommend"`          // ？。作用尚不明确
	Guide                  any                 `json:"guide"`                      // ？。作用尚不明确
	QueryTags              any                 `json:"query_tags"`                 // ？。作用尚不明确
	ModuleCtrl             any                 `json:"module_ctrl"`                // ？。作用尚不明确
}

// Emergency 视频操作按钮信息。控制详情页各操作按钮是否展示
type Emergency struct {
	NoLike  bool `json:"no_like"`  // 是否不显示点赞按钮
	NoCoin  bool `json:"no_coin"`  // 是否不显示投币按钮
	NoFav   bool `json:"no_fav"`   // 是否不显示收藏按钮
	NoShare bool `json:"no_share"` // 是否不显示分享按钮
}

// UgcSeason 视频合集信息。不在合集中的视频无此项
type UgcSeason struct {
	Id          int                `json:"id"`            // 视频合集id
	Title       string             `json:"title"`         // 视频合集标题
	Cover       string             `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string             `json:"intro"`         // 视频合集介绍
	SignState   int                `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                `json:"attribute"`     // 稿件属性位。详情见[属性数据文档](attribute_data.md#attribute字段值(稿件属性位))
	Sections    []UgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        UgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool               `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                `json:"enable_vt"`     // （？）。作用尚不明确
}

// UgcSeasonStat `ugc_season` 中的 `stat` 对象。
// 与 `VideoStat` 不同：收藏数键名是 `fav` 而非 `favorite`，且多了 `vv`
type UgcSeasonStat struct {
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

// UgcSeasonSection `ugc_season` 中的 `sections` 数组中的对象
type UgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 `section_id`，但本节 JSON 示例里是 `id`，
	// 且示例中不存在 `section_id`。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                `json:"id"`
	SectionID int                `json:"section_id"`
	Title     string             `json:"title"`    // 视频合集中分部标题
	Type      int                `json:"type"`     // （？）。作用尚不明确
	Episodes  []UgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// UgcSeasonEpisode `ugc_season.sections[].episodes[]` 中的对象
type UgcSeasonEpisode struct {
	SeasonID  int                 `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                 `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                 `json:"id"`         // 分部中视频id
	Aid       int                 `json:"aid"`        // 视频aid
	Cid       int                 `json:"cid"`        // 视频cid
	Title     string              `json:"title"`      // 视频标题。合集列表中展示的标题，默认视频真实标题
	Attribute int                 `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       UgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      VideoPage           `json:"page"`       // 视频分P信息
	Bvid      string              `json:"bvid"`       // 视频bvid
	Pages     []VideoPage         `json:"pages"`      // 视频分P列表
}

// UgcSeasonEpisodeArc `ugc_season.sections[].episodes[].arc` 对象，视频详细信息。
//
// 文档称「基本同『获取视频详细信息(web端)』中的 data 对象」，但实测键名有差异 ——
// `author` 而非 `owner`、`type_id`/`type_name` 而非 `tid`/`tname`、`stat.fav` 而非
// `stat.favorite`。复用 VideoInfo 会把这些字段静默丢掉，所以单独建型。
type UgcSeasonEpisodeArc struct {
	Aid                int                  `json:"aid"`                  // 稿件avid
	Videos             int                  `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                  `json:"type_id"`              // 分区tid
	TypeName           string               `json:"type_name"`            // 子分区名称
	Copyright          int                  `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string               `json:"pic"`                  // 稿件封面图片url
	Title              string               `json:"title"`                // 稿件标题
	PubDate            int                  `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                  `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string               `json:"desc"`                 // 视频简介
	State              int                  `json:"state"`                // 视频状态
	Duration           int                  `json:"duration"`             // 稿件总时长。单位为秒
	Rights             VideoRights          `json:"rights"`               // 视频属性标志
	Author             Owner                `json:"author"`               // UP主信息
	Stat               UgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string               `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          Dimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                  `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                 `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                 `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                  `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string               `json:"vt_display"`           // （？）。作用尚不明确
}

// UgcSeasonEpisodeStat `ugc_season.sections[].episodes[].arc.stat` 对象。
// 有 `fav`、`vv`、`argue_msg`，无 `favorite`、`no_reprint`，与 `VideoStat` 不同型
type UgcSeasonEpisodeStat struct {
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

type Dimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

type VideoPage struct {
	Cid        int       `json:"cid"`         // 当前分P cid
	Page       int       `json:"page"`        // 当前分P
	From       string    `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string    `json:"part"`        // 当前分P标题
	Duration   int       `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string    `json:"vid"`         // 站外视频vid
	Weblink    string    `json:"weblink"`     // 站外视频跳转url
	Dimension  Dimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string    `json:"first_frame"` // 分P封面
}

type StatusCount struct {
	View  int `json:"view"`  // 0。作用尚不明确
	Use   int `json:"use"`   // 视频添加TAG数
	Atten int `json:"atten"` // TAG关注
}

type VideoTag struct {
	TagID        int         `json:"tag_id"`        // tag_id
	TagName      string      `json:"tag_name"`      // TAG名称
	Cover        string      `json:"cover"`         // TAG图片url
	HeadCover    string      `json:"head_cover"`    // TAG页面头图url
	Content      string      `json:"content"`       // TAG介绍
	ShortContent string      `json:"short_content"` // TAG简介
	Type         int         `json:"type"`          // ？？？
	State        int         `json:"state"`         // 0
	Ctime        int         `json:"ctime"`         // 创建时间。时间戳
	Count        StatusCount `json:"count"`         // 状态数
	IsAtten      int         `json:"is_atten"`      // 是否关注。0：未关注。1：已关注。需要登录(Cookie) 。未登录为0
	Likes        int         `json:"likes"`         // 0。作用尚不明确
	Hates        int         `json:"hates"`         // 0。作用尚不明确
	Attribute    int         `json:"attribute"`     // 0。作用尚不明确
	Liked        int         `json:"liked"`         // 是否已经点赞。0：未点赞。1：已点赞。需要登录(Cookie) 。未登录为0
	Hated        int         `json:"hated"`         // 是否已经点踩。0：未点踩。1：已点踩。需要登录(Cookie) 。未登录为0
	ExtraAttr    int         `json:"extra_attr"`    // ? ? ?
}

type CoinVideoResult struct {
	Like bool `json:"like"` // 是否点赞成功。true：成功。false：失败。已赞过则附加点赞失败
}

type FavourVideoResult struct {
	Prompt bool `json:"prompt"` // 是否为未关注用户收藏。false：否。true：是
}

type LikeCoinFavourResult struct {
	Like     bool `json:"like"`     // 是否点赞成功。true：成功。false：失败
	Coin     bool `json:"coin"`     // 是否投币成功。true：成功。false：失败
	Fav      bool `json:"fav"`      // 是否收藏成功。true：成功。false：失败
	Multiply int  `json:"multiply"` // 投币枚数。默认为2
}

type ShowSwitch struct {
	Total bool `json:"total"` // 展示所有终端总计人数
	Count bool `json:"count"` // 展示web端实时在线人数
}

type VideoOnlineInfo struct {
	Total      string     `json:"total"`       // 所有终端总计人数。例如10万+
	Count      string     `json:"count"`       // web端实时在线人数
	ShowSwitch ShowSwitch `json:"show_switch"` // 数据显示控制
	Abtest     Abtest     `json:"abtest"`      // AB测试分组。实测为 {"group":"b"}
}

// Abtest AB 测试分组
type Abtest struct {
	Group string `json:"group"` // 分组标识。实测为 b
}

type VideoStatusNumber struct {
	Aid        int            `json:"aid"`        // 稿件avid
	Bvid       string         `json:"bvid"`       // 稿件bvid
	View       NumberOrString `json:"view"`       // 正常：播放次数(num)。屏蔽："--"(str)
	Danmaku    int            `json:"danmaku"`    // 弹幕条数
	Reply      int            `json:"reply"`      // 评论条数
	Favorite   int            `json:"favorite"`   // 收藏人数
	Coin       int            `json:"coin"`       // 投币枚数
	Share      int            `json:"share"`      // 分享次数
	NowRank    int            `json:"now_rank"`   // 0。作用尚不明确
	HisRank    int            `json:"his_rank"`   // 历史最高排行
	Like       int            `json:"like"`       // 获赞次数
	Dislike    int            `json:"dislike"`    // 0。作用尚不明确
	NoReprint  int            `json:"no_reprint"` // 禁止转载标志。0：无。1：禁止
	Copyright  int            `json:"copyright"`  // 版权标志。1：自制。2：转载
	ArgueMsg   string         `json:"argue_msg"`  // 警告信息。默认为空
	Evaluation string         `json:"evaluation"` // 视频评分。默认为空
}
