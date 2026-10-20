package video

// GetVideoInfo 接口的响应模型。
//
// 本文件的类型只服务于 GetVideoInfo，不与其它接口共用；即使字段相同，
// 其它接口也在各自 model_*.go 中单独定义。

// GetVideoInfoResult 是「获取视频详细信息」的 data 对象。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的 data 对象结构重复是设计使然
type GetVideoInfoResult struct {
	Bvid                    string                 `json:"bvid"`         // 稿件bvid
	Aid                     int                    `json:"aid"`          // 稿件avid
	Videos                  int                    `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                    `json:"tid"`          // 分区tid
	Tname                   string                 `json:"tname"`        // 子分区名称
	Copyright               int                    `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                 `json:"pic"`          // 稿件封面图片url
	Title                   string                 `json:"title"`        // 稿件标题
	PubDate                 int                    `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                    `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                 `json:"desc"`         // 视频简介
	DescV2                  []GetVideoInfoDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                    `json:"state"`        // 视频状态
	Duration                int                    `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                    `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                    `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                 `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights                  GetVideoInfoRights     `json:"rights"`       // 视频属性标志
	Owner                   GetVideoInfoOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetVideoInfoStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                 `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                    `json:"cid"`          // 视频1P cid
	Dimension               GetVideoInfoDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                    `json:"season_id"`    // 合集id
	Premiere                any                    `json:"premiere"`     // null
	TeenageMode             int                    `json:"teenage_mode"`
	IsChargeableSeason      bool                   `json:"is_chargeable_season"`
	IsStory                 bool                   `json:"is_story"`
	NoCache                 bool                   `json:"no_cache"` // 作用尚不明确
	Pages                   []GetVideoInfoPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetVideoInfoSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetVideoInfoStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                   `json:"is_season_display"`
	UserGarb                GetVideoInfoUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetVideoInfoHonorReply `json:"honor_reply"`
	LikeIcon                string                 `json:"like_icon"`
	ArgueInfo               GetVideoInfoArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                    `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                 `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                    `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                 `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                    `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                 `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                    `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                    `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                   `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                    `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetVideoInfoAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                    `json:"tid_v2"`                      // 分区tid (v2)
	TnameV2                 string                 `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                   `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                   `json:"is_upower_play"`              //
	IsUpowerPreview         bool                   `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                   `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                   `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetVideoInfoUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                   `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                   `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                   `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                   `json:"is_view_self"`                // 是否尽自己可见
}

// GetVideoInfoDesc 是「获取视频详细信息」desc_v2 数组中的对象。
type GetVideoInfoDesc struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetVideoInfoRights 是「获取视频详细信息」的 rights 对象。
type GetVideoInfoRights struct {
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
	PayFreeWatch  int `json:"pay_free_watch"`  // （？）。作用尚不明确
}

// GetVideoInfoOwner 是「获取视频详细信息」的 owner 对象。
type GetVideoInfoOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetVideoInfoStat 是「获取视频详细信息」的 stat 对象。
type GetVideoInfoStat struct {
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
	FavG       int    `json:"fav_g"`      // （？）。作用尚不明确
	LikeG      int    `json:"like_g"`     // （？）。作用尚不明确
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetVideoInfoDimension 是「获取视频详细信息」的 dimension 对象。
type GetVideoInfoDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetVideoInfoPage 是「获取视频详细信息」pages 数组中的对象。
type GetVideoInfoPage struct {
	Cid        int                   `json:"cid"`         // 当前分P cid
	Page       int                   `json:"page"`        // 当前分P
	From       string                `json:"from"`        // 视频来源。vupload：普通上传（B站）。hunan：芒果TV。qq：腾讯
	Part       string                `json:"part"`        // 当前分P标题
	Duration   int                   `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                `json:"vid"`         // 站外视频vid
	Weblink    string                `json:"weblink"`     // 站外视频跳转url
	Dimension  GetVideoInfoDimension `json:"dimension"`   // 当前分P分辨率。有部分视频无法获取分辨率
	FirstFrame string                `json:"first_frame"` // 分P封面
}

// GetVideoInfoSubtitle 是「获取视频详细信息」的 subtitle 对象。
//
// 文档中 subtitle 是 {allow_submit, list[]} 容器，字幕条目见 GetVideoInfoSubtitleItem。
type GetVideoInfoSubtitle struct {
	AllowSubmit bool                       `json:"allow_submit"` // 是否允许提交字幕
	List        []GetVideoInfoSubtitleItem `json:"list"`         // 字幕列表。未登录为空
}

// GetVideoInfoSubtitleItem 是「获取视频详细信息」subtitle.list 数组中的对象。
type GetVideoInfoSubtitleItem struct {
	ID             int                        `json:"id"`              // 字幕id
	Lan            string                     `json:"lan"`             // 字幕语言
	LanDoc         string                     `json:"lan_doc"`         // 字幕语言名称
	IsLock         bool                       `json:"is_lock"`         // 是否锁定
	AuthorMid      int                        `json:"author_mid"`      // 字幕上传者mid
	SubtitleURL    string                     `json:"subtitle_url"`    // json格式字幕文件url
	Author         GetVideoInfoSubtitleAuthor `json:"author"`          // 字幕上传者信息
	SubtitleHeight any                        `json:"subtitle_height"` // （？）。作用尚不明确
}

// GetVideoInfoSubtitleAuthor 是「获取视频详细信息」subtitle.list[].author 对象。
type GetVideoInfoSubtitleAuthor struct {
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

// GetVideoInfoStaff 是「获取视频详细信息」staff 数组中的对象。
type GetVideoInfoStaff struct {
	Mid        int                       `json:"mid"`      // 成员mid
	Title      string                    `json:"title"`    // 成员名称
	Name       string                    `json:"name"`     // 成员昵称
	Face       string                    `json:"face"`     // 成员头像url
	Vip        GetVideoInfoStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetVideoInfoStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                       `json:"follower"` // 成员粉丝数
	LabelStyle int                       `json:"label_style"`
}

// GetVideoInfoStaffVip 是「获取视频详细信息」staff[].vip 对象。
type GetVideoInfoStaffVip struct {
	Type       int                     `json:"type"`         // 成员会员类型。0：无。1：月会员。2：年会员
	Status     int                     `json:"status"`       // 会员状态。0：无。1：有
	DueDate    int                     `json:"due_date"`     // 到期时间。UNIX 毫秒时间戳
	VipPayType int                     `json:"vip_pay_type"` // 支付类型
	ThemeType  int                     `json:"theme_type"`   // 0
	Label      GetVideoInfoLabel       `json:"label"`        // 成员会员标签
	OttInfo    GetVideoInfoVipOttInfo  `json:"ott_info"`     // （？）。作用尚不明确
	SuperVip   GetVideoInfoVipSuperVip `json:"super_vip"`    // （？）。作用尚不明确
}

// GetVideoInfoLabel 是「获取视频详细信息」staff[].vip.label 对象。
type GetVideoInfoLabel struct {
	Path                  string                `json:"path"`                      // 空。作用尚不明确
	Text                  string                `json:"text"`                      // 会员类型文案
	LabelTheme            string                `json:"label_theme"`               // 会员标签
	TextColor             string                `json:"text_color"`                // 会员标签文字颜色
	BgStyle               int                   `json:"bg_style"`                  // 1
	BgColor               string                `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                `json:"border_color"`              // 会员标签边框颜色
	UseImgLabel           bool                  `json:"use_img_label"`             // true
	ImgLabelURIHans       string                `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                   `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                   `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetVideoInfoLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetVideoInfoLabelGoto 是「获取视频详细信息」staff[].vip.label.label_goto 对象。
type GetVideoInfoLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetVideoInfoVipOttInfo 是「获取视频详细信息」staff[].vip.ott_info 对象。
type GetVideoInfoVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetVideoInfoVipSuperVip 是「获取视频详细信息」staff[].vip.super_vip 对象。
type GetVideoInfoVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetVideoInfoStaffOfficial 是「获取视频详细信息」staff[].official 对象。
type GetVideoInfoStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetVideoInfoUserGarb 是「获取视频详细信息」的 user_garb 对象。
type GetVideoInfoUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetVideoInfoHonorReply 是「获取视频详细信息」的 honor_reply 对象。
type GetVideoInfoHonorReply struct {
	Honor []GetVideoInfoHonor `json:"honor"`
}

// GetVideoInfoHonor 是「获取视频详细信息」honor_reply.honor 数组中的对象。
type GetVideoInfoHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetVideoInfoArgueInfo 是「获取视频详细信息」的 argue_info 对象。
type GetVideoInfoArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetVideoInfoAiRcmd 是「获取视频详细信息」的 ai_rcmd 对象。
type GetVideoInfoAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetVideoInfoUgcSeason 是「获取视频详细信息」的 ugc_season 对象。
type GetVideoInfoUgcSeason struct {
	Id          int                            `json:"id"`            // 视频合集id
	Title       string                         `json:"title"`         // 视频合集标题
	Cover       string                         `json:"cover"`         // 视频合集封面url。文档字段表未列，取自本节 JSON 示例
	Mid         int                            `json:"mid"`           // 视频合集作者id。文档字段表标为 str，示例中是数字
	Intro       string                         `json:"intro"`         // 视频合集介绍
	SignState   int                            `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                            `json:"attribute"`     // 稿件属性位
	Sections    []GetVideoInfoUgcSeasonSection `json:"sections"`      // 视频合集中分部列表，名称可由up主自定义，默认为正片
	Stat        GetVideoInfoUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                            `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                            `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                           `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                            `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetVideoInfoUgcSeasonSection 是「获取视频详细信息」ugc_season.sections 数组中的对象。
type GetVideoInfoUgcSeasonSection struct {
	SeasonID int `json:"season_id"` // 视频合集中分部所属视频合集id
	// 分部的 id。文档字段表写作 section_id，但本节 JSON 示例里是 id，
	// 且示例中不存在 section_id。两个都留着 —— 缺失的键不会填值，不会丢数据
	ID        int                            `json:"id"`
	SectionID int                            `json:"section_id"`
	Title     string                         `json:"title"`    // 视频合集中分部标题
	Type      int                            `json:"type"`     // （？）。作用尚不明确
	Episodes  []GetVideoInfoUgcSeasonEpisode `json:"episodes"` // 视频合集中分部的视频列表
}

// GetVideoInfoUgcSeasonEpisode 是「获取视频详细信息」ugc_season.sections[].episodes[] 中的对象。
type GetVideoInfoUgcSeasonEpisode struct {
	SeasonID  int                             `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                             `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                             `json:"id"`         // 分部中视频id
	Aid       int                             `json:"aid"`        // 视频aid
	Cid       int                             `json:"cid"`        // 视频cid
	Title     string                          `json:"title"`      // 视频标题
	Attribute int                             `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetVideoInfoUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetVideoInfoPage                `json:"page"`       // 视频分P信息
	Bvid      string                          `json:"bvid"`       // 视频bvid
	Pages     []GetVideoInfoPage              `json:"pages"`      // 视频分P列表
}

// GetVideoInfoUgcSeasonEpisodeArc 是「获取视频详细信息」ugc_season.sections[].episodes[].arc 对象。
//
// 文档称「基本同获取视频详细信息的 data 对象」，但实测键名有差异 ——
// author 而非 owner、type_id/type_name 而非 tid/tname、stat.fav 而非
// stat.favorite。复用 GetVideoInfoResult 会把这些字段静默丢掉，所以单独建型。
type GetVideoInfoUgcSeasonEpisodeArc struct {
	Aid                int                              `json:"aid"`                  // 稿件avid
	Videos             int                              `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                              `json:"type_id"`              // 分区tid
	TypeName           string                           `json:"type_name"`            // 子分区名称
	Copyright          int                              `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string                           `json:"pic"`                  // 稿件封面图片url
	Title              string                           `json:"title"`                // 稿件标题
	PubDate            int                              `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                              `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string                           `json:"desc"`                 // 视频简介
	State              int                              `json:"state"`                // 视频状态
	Duration           int                              `json:"duration"`             // 稿件总时长。单位为秒
	Rights             GetVideoInfoRights               `json:"rights"`               // 视频属性标志
	Author             GetVideoInfoOwner                `json:"author"`               // UP主信息
	Stat               GetVideoInfoUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                           `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetVideoInfoDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                              `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                             `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                             `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                              `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                           `json:"vt_display"`           // （？）。作用尚不明确
}

// GetVideoInfoUgcSeasonEpisodeStat 是「获取视频详细信息」ugc_season.sections[].episodes[].arc.stat 对象。
//
// 有 fav、vv、argue_msg，无 favorite、no_reprint，与 GetVideoInfoStat 不同型。
type GetVideoInfoUgcSeasonEpisodeStat struct {
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

// GetVideoInfoUgcSeasonStat 是「获取视频详细信息」ugc_season.stat 对象。
//
// 与 GetVideoInfoStat 不同：收藏数键名是 fav 而非 favorite，且多了 vv。
type GetVideoInfoUgcSeasonStat struct {
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
