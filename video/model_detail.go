package video

// GetVideoDetailInfo 接口的响应模型。
//
// 本文件的类型只服务于 GetVideoDetailInfo，不与其它接口共用。
// View 与 Related 同属一个接口，因此共用本文件内的嵌套类型；Related 没有分P信息。

// GetVideoDetailInfoResult 是「获取视频超详细信息」的 data 对象。
type GetVideoDetailInfoResult struct {
	View                   GetVideoDetailInfoView      `json:"View"`                       // 视频基本信息
	Card                   GetVideoDetailInfoCard      `json:"Card"`                       // 视频UP主信息
	Tags                   []GetVideoDetailInfoTag     `json:"Tags"`                       // 视频TAG信息
	Reply                  GetVideoDetailInfoReply     `json:"Reply"`                      // 视频热评信息
	Related                []GetVideoDetailInfoRelated `json:"Related"`                    // 推荐视频信息
	Spec                   any                         `json:"Spec"`                       // ？。作用尚不明确
	IsHitLabourDayActivity bool                        `json:"is_hit_labour_day_activity"` // （？）。作用尚不明确
	HotShare               any                         `json:"hot_share"`                  // ？。作用尚不明确
	Elec                   any                         `json:"elec"`                       // ？。作用尚不明确
	Recommend              any                         `json:"recommend"`                  // ？。作用尚不明确
	ViewAddit              any                         `json:"view_addit"`                 // ？。作用尚不明确
	Emergency              GetVideoDetailInfoEmergency `json:"emergency"`                  // 视频操作按钮信息
	Participle             []string                    `json:"participle"`                 // 分词信息。用于推荐
	ReplaceRecommend       bool                        `json:"replace_recommend"`          // ？。作用尚不明确
	Guide                  any                         `json:"guide"`                      // ？。作用尚不明确
	QueryTags              any                         `json:"query_tags"`                 // ？。作用尚不明确
	ModuleCtrl             any                         `json:"module_ctrl"`                // ？。作用尚不明确
}

// GetVideoDetailInfoView 是「获取视频超详细信息」data.View 对象。
//
//nolint:dupl // 按接口隔离返回体，与其它接口的 data 对象结构重复是设计使然
type GetVideoDetailInfoView struct {
	Bvid                    string                       `json:"bvid"`         // 稿件bvid
	Aid                     int                          `json:"aid"`          // 稿件avid
	Videos                  int                          `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                          `json:"tid"`          // 分区tid
	Tname                   string                       `json:"tname"`        // 子分区名称
	Copyright               int                          `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                       `json:"pic"`          // 稿件封面图片url
	Title                   string                       `json:"title"`        // 稿件标题
	PubDate                 int                          `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                          `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                       `json:"desc"`         // 视频简介
	DescV2                  []GetVideoDetailInfoDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                          `json:"state"`        // 视频状态
	Duration                int                          `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                          `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                          `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                       `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段
	Rights                  GetVideoDetailInfoRights     `json:"rights"`       // 视频属性标志
	Owner                   GetVideoDetailInfoOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetVideoDetailInfoStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                       `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                          `json:"cid"`          // 视频1P cid
	Dimension               GetVideoDetailInfoDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                          `json:"season_id"`    // 合集id
	Premiere                any                          `json:"premiere"`     // null
	TeenageMode             int                          `json:"teenage_mode"`
	IsChargeableSeason      bool                         `json:"is_chargeable_season"`
	IsStory                 bool                         `json:"is_story"`
	NoCache                 bool                         `json:"no_cache"` // 作用尚不明确
	Pages                   []GetVideoDetailInfoPage     `json:"pages"`    // 视频分P列表
	Subtitle                GetVideoDetailInfoSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetVideoDetailInfoStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                         `json:"is_season_display"`
	UserGarb                GetVideoDetailInfoUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetVideoDetailInfoHonorReply `json:"honor_reply"`
	LikeIcon                string                       `json:"like_icon"`
	ArgueInfo               GetVideoDetailInfoArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                          `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                       `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                          `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                       `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                          `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                       `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                          `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                          `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                         `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                          `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetVideoDetailInfoAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                          `json:"tid_v2"`                      // 分区tid (v2)
	TnameV2                 string                       `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                         `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                         `json:"is_upower_play"`              //
	IsUpowerPreview         bool                         `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                         `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                         `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetVideoDetailInfoUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                         `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                         `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                         `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                         `json:"is_view_self"`                // 是否尽自己可见
}

// GetVideoDetailInfoRelated 是「获取视频超详细信息」data.Related 数组中的对象。
//
// 文档注明其基本同 View，但没有分P信息。
//
//nolint:dupl // 按接口隔离返回体，与 View 结构重复是设计使然
type GetVideoDetailInfoRelated struct {
	Bvid                    string                       `json:"bvid"`         // 稿件bvid
	Aid                     int                          `json:"aid"`          // 稿件avid
	Videos                  int                          `json:"videos"`       // 稿件分P总数。默认为1
	Tid                     int                          `json:"tid"`          // 分区tid
	Tname                   string                       `json:"tname"`        // 子分区名称
	Copyright               int                          `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                     string                       `json:"pic"`          // 稿件封面图片url
	Title                   string                       `json:"title"`        // 稿件标题
	PubDate                 int                          `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime                   int                          `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc                    string                       `json:"desc"`         // 视频简介
	DescV2                  []GetVideoDetailInfoDesc     `json:"desc_v2"`      // 新版视频简介
	State                   int                          `json:"state"`        // 视频状态
	Duration                int                          `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward                 int                          `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionID               int                          `json:"mission_id"`   // 稿件参与的活动id
	RedirectURL             string                       `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段
	Rights                  GetVideoDetailInfoRights     `json:"rights"`       // 视频属性标志
	Owner                   GetVideoDetailInfoOwner      `json:"owner"`        // 视频UP主信息
	Stat                    GetVideoDetailInfoStat       `json:"stat"`         // 视频状态数
	Dynamic                 string                       `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                     int                          `json:"cid"`          // 视频1P cid
	Dimension               GetVideoDetailInfoDimension  `json:"dimension"`    // 视频1P分辨率
	SeasonID                int                          `json:"season_id"`    // 合集id
	Premiere                any                          `json:"premiere"`     // null
	TeenageMode             int                          `json:"teenage_mode"`
	IsChargeableSeason      bool                         `json:"is_chargeable_season"`
	IsStory                 bool                         `json:"is_story"`
	NoCache                 bool                         `json:"no_cache"` // 作用尚不明确
	Subtitle                GetVideoDetailInfoSubtitle   `json:"subtitle"` // 视频CC字幕信息
	Staff                   []GetVideoDetailInfoStaff    `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay         bool                         `json:"is_season_display"`
	UserGarb                GetVideoDetailInfoUserGarb   `json:"user_garb"` // 用户装扮信息
	HonorReply              GetVideoDetailInfoHonorReply `json:"honor_reply"`
	LikeIcon                string                       `json:"like_icon"`
	ArgueInfo               GetVideoDetailInfoArgueInfo  `json:"argue_info"`                  // 争议/警告信息
	UpFromV2                int                          `json:"up_from_v2"`                  // （？）。作用尚不明确
	PubLocation             string                       `json:"pub_location"`                // （？）。作用尚不明确
	Tidv2                   int                          `json:"tidv2"`                       // （？）。作用尚不明确
	Tnamev2                 string                       `json:"tnamev2"`                     // （？）。作用尚不明确
	PidV2                   int                          `json:"pid_v2"`                      // （？）。作用尚不明确
	PidNameV2               string                       `json:"pid_name_v2"`                 // （？）。作用尚不明确
	CurrentState            int                          `json:"current_state"`               // （？）。作用尚不明确
	GlobalState             int                          `json:"global_state"`                // （？）。作用尚不明确
	IsOgv                   bool                         `json:"is_ogv"`                      // （？）。作用尚不明确
	AttributeV3             int                          `json:"attribute_v3"`                // （？）。作用尚不明确
	AiRcmd                  GetVideoDetailInfoAiRcmd     `json:"ai_rcmd"`                     // （？）。作用尚不明确
	TidV2                   int                          `json:"tid_v2"`                      // 分区tid (v2)
	TnameV2                 string                       `json:"tname_v2"`                    // 子分区名称 (v2)
	IsUpowerExclusive       bool                         `json:"is_upower_exclusive"`         // 是否为充电专属视频
	IsUpowerPlay            bool                         `json:"is_upower_play"`              //
	IsUpowerPreview         bool                         `json:"is_upower_preview"`           // 充电专属视频是否支持试看
	IsUpowerExclusiveWithQa bool                         `json:"is_upower_exclusive_with_qa"` // （？）。作用尚不明确
	IsHuaSheng              bool                         `json:"is_hua_sheng"`                // （？）。作用尚不明确
	UgcSeason               GetVideoDetailInfoUgcSeason  `json:"ugc_season"`                  // 视频合集信息。不在合集中的视频无此项
	NeedJumpBv              bool                         `json:"need_jump_bv"`                // 需要跳转到BV号?
	DisableShowUpInfo       bool                         `json:"disable_show_up_info"`        // 禁止展示UP主信息?
	IsStoryPlay             bool                         `json:"is_story_play"`               // 。作用未知，可能与动态视频有关
	IsViewSelf              bool                         `json:"is_view_self"`                // 是否尽自己可见
}

// GetVideoDetailInfoDesc 是「获取视频超详细信息」desc_v2 数组中的对象。
type GetVideoDetailInfoDesc struct {
	RawText string `json:"raw_text"` // 简介内容
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizID   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
}

// GetVideoDetailInfoRights 是「获取视频超详细信息」的 rights 对象。
type GetVideoDetailInfoRights struct {
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

// GetVideoDetailInfoOwner 是「获取视频超详细信息」的 owner 对象。
type GetVideoDetailInfoOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetVideoDetailInfoStat 是「获取视频超详细信息」的 stat 对象。
type GetVideoDetailInfoStat struct {
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

// GetVideoDetailInfoDimension 是「获取视频超详细信息」的 dimension 对象。
type GetVideoDetailInfoDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetVideoDetailInfoPage 是「获取视频超详细信息」pages 数组中的对象。
type GetVideoDetailInfoPage struct {
	Cid        int                         `json:"cid"`         // 当前分P cid
	Page       int                         `json:"page"`        // 当前分P
	From       string                      `json:"from"`        // 视频来源
	Part       string                      `json:"part"`        // 当前分P标题
	Duration   int                         `json:"duration"`    // 当前分P持续时间。单位为秒
	Vid        string                      `json:"vid"`         // 站外视频vid
	Weblink    string                      `json:"weblink"`     // 站外视频跳转url
	Dimension  GetVideoDetailInfoDimension `json:"dimension"`   // 当前分P分辨率
	FirstFrame string                      `json:"first_frame"` // 分P封面
}

// GetVideoDetailInfoSubtitle 是「获取视频超详细信息」的 subtitle 对象。
type GetVideoDetailInfoSubtitle struct {
	AllowSubmit bool                             `json:"allow_submit"` // 是否允许提交字幕
	List        []GetVideoDetailInfoSubtitleItem `json:"list"`         // 字幕列表。未登录为空
}

// GetVideoDetailInfoSubtitleItem 是「获取视频超详细信息」subtitle.list 数组中的对象。
type GetVideoDetailInfoSubtitleItem struct {
	ID             int                              `json:"id"`              // 字幕id
	Lan            string                           `json:"lan"`             // 字幕语言
	LanDoc         string                           `json:"lan_doc"`         // 字幕语言名称
	IsLock         bool                             `json:"is_lock"`         // 是否锁定
	AuthorMid      int                              `json:"author_mid"`      // 字幕上传者mid
	SubtitleURL    string                           `json:"subtitle_url"`    // json格式字幕文件url
	Author         GetVideoDetailInfoSubtitleAuthor `json:"author"`          // 字幕上传者信息
	SubtitleHeight any                              `json:"subtitle_height"` // （？）。作用尚不明确
}

// GetVideoDetailInfoSubtitleAuthor 是「获取视频超详细信息」subtitle.list[].author 对象。
type GetVideoDetailInfoSubtitleAuthor struct {
	Mid           int    `json:"mid"`             // 字幕上传者mid
	Name          string `json:"name"`            // 字幕上传者昵称
	Sex           string `json:"sex"`             // 字幕上传者性别
	Face          string `json:"face"`            // 字幕上传者头像url
	Sign          string `json:"sign"`            // 字幕上传者签名
	Rank          int    `json:"rank"`            // 10000。作用尚不明确
	Birthday      int    `json:"birthday"`        // 0。作用尚不明确
	IsFakeAccount int    `json:"is_fake_account"` // 0。作用尚不明确
	IsDeleted     int    `json:"is_deleted"`      // 0。作用尚不明确
}

// GetVideoDetailInfoStaff 是「获取视频超详细信息」staff 数组中的对象。
type GetVideoDetailInfoStaff struct {
	Mid        int                             `json:"mid"`      // 成员mid
	Title      string                          `json:"title"`    // 成员名称
	Name       string                          `json:"name"`     // 成员昵称
	Face       string                          `json:"face"`     // 成员头像url
	Vip        GetVideoDetailInfoStaffVip      `json:"vip"`      // 成员大会员状态
	Official   GetVideoDetailInfoStaffOfficial `json:"official"` // 成员认证信息
	Follower   int                             `json:"follower"` // 成员粉丝数
	LabelStyle int                             `json:"label_style"`
}

// GetVideoDetailInfoStaffVip 是「获取视频超详细信息」staff[].vip 对象。
type GetVideoDetailInfoStaffVip struct {
	Type       int                           `json:"type"`         // 成员会员类型
	Status     int                           `json:"status"`       // 会员状态
	DueDate    int                           `json:"due_date"`     // 到期时间。UNIX 毫秒时间戳
	VipPayType int                           `json:"vip_pay_type"` // 支付类型
	ThemeType  int                           `json:"theme_type"`   // 0
	Label      GetVideoDetailInfoLabel       `json:"label"`        // 成员会员标签
	OttInfo    GetVideoDetailInfoVipOttInfo  `json:"ott_info"`     // （？）。作用尚不明确
	SuperVip   GetVideoDetailInfoVipSuperVip `json:"super_vip"`    // （？）。作用尚不明确
}

// GetVideoDetailInfoLabel 是「获取视频超详细信息」staff[].vip.label 对象。
type GetVideoDetailInfoLabel struct {
	Path                  string                      `json:"path"`                      // 空。作用尚不明确
	Text                  string                      `json:"text"`                      // 会员类型文案
	LabelTheme            string                      `json:"label_theme"`               // 会员标签
	TextColor             string                      `json:"text_color"`                // 会员标签文字颜色
	BgStyle               int                         `json:"bg_style"`                  // 1
	BgColor               string                      `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string                      `json:"border_color"`              // 会员标签边框颜色
	UseImgLabel           bool                        `json:"use_img_label"`             // true
	ImgLabelURIHans       string                      `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                      `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                      `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                      `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                         `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                         `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetVideoDetailInfoLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                      `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                      `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetVideoDetailInfoLabelGoto 是「获取视频超详细信息」staff[].vip.label.label_goto 对象。
type GetVideoDetailInfoLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetVideoDetailInfoVipOttInfo 是「获取视频超详细信息」staff[].vip.ott_info 对象。
type GetVideoDetailInfoVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetVideoDetailInfoVipSuperVip 是「获取视频超详细信息」staff[].vip.super_vip 对象。
type GetVideoDetailInfoVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetVideoDetailInfoStaffOfficial 是「获取视频超详细信息」staff[].official 对象。
type GetVideoDetailInfoStaffOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetVideoDetailInfoUserGarb 是「获取视频超详细信息」的 user_garb 对象。
type GetVideoDetailInfoUserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"` // 某url？
}

// GetVideoDetailInfoHonorReply 是「获取视频超详细信息」的 honor_reply 对象。
type GetVideoDetailInfoHonorReply struct {
	Honor []GetVideoDetailInfoHonor `json:"honor"`
}

// GetVideoDetailInfoHonor 是「获取视频超详细信息」honor_reply.honor 数组中的对象。
type GetVideoDetailInfoHonor struct {
	Aid                int    `json:"aid"`  // 当前稿件aid
	Type               int    `json:"type"` // 1：入站必刷收录。2：第?期每周必看。3：全站排行榜最高第?名。4：热门
	Desc               string `json:"desc"` // 描述
	WeeklyRecommendNum int    `json:"weekly_recommend_num"`
}

// GetVideoDetailInfoArgueInfo 是「获取视频超详细信息」的 argue_info 对象。
type GetVideoDetailInfoArgueInfo struct {
	ArgueLink string `json:"argue_link"` // 作用尚不明确
	ArgueMsg  string `json:"argue_msg"`  // 警告/争议提示信息
	ArgueType int    `json:"argue_type"` // 作用尚不明确
}

// GetVideoDetailInfoAiRcmd 是「获取视频超详细信息」的 ai_rcmd 对象。
type GetVideoDetailInfoAiRcmd struct {
	Id      int    `json:"id"`      // （？）。作用尚不明确
	Goto    string `json:"goto"`    // （？）。作用尚不明确
	Trackid string `json:"trackid"` // （？）。作用尚不明确
	UniqId  string `json:"uniq_id"` // （？）。作用尚不明确
}

// GetVideoDetailInfoUgcSeason 是「获取视频超详细信息」的 ugc_season 对象。
type GetVideoDetailInfoUgcSeason struct {
	Id          int                                  `json:"id"`            // 视频合集id
	Title       string                               `json:"title"`         // 视频合集标题
	Cover       string                               `json:"cover"`         // 视频合集封面url
	Mid         int                                  `json:"mid"`           // 视频合集作者id
	Intro       string                               `json:"intro"`         // 视频合集介绍
	SignState   int                                  `json:"sign_state"`    // （？）。作用尚不明确
	Attribute   int                                  `json:"attribute"`     // 稿件属性位
	Sections    []GetVideoDetailInfoUgcSeasonSection `json:"sections"`      // 视频合集中分部列表
	Stat        GetVideoDetailInfoUgcSeasonStat      `json:"stat"`          // 视频合集状态数
	EpCount     int                                  `json:"ep_count"`      // 视频合集中视频数量
	SeasonType  int                                  `json:"season_type"`   // （？）。作用尚不明确
	IsPaySeason bool                                 `json:"is_pay_season"` // 是否为付费合集
	EnableVt    int                                  `json:"enable_vt"`     // （？）。作用尚不明确
}

// GetVideoDetailInfoUgcSeasonSection 是「获取视频超详细信息」ugc_season.sections 数组中的对象。
type GetVideoDetailInfoUgcSeasonSection struct {
	SeasonID  int                                  `json:"season_id"`  // 视频合集中分部所属视频合集id
	ID        int                                  `json:"id"`         // 分部id
	SectionID int                                  `json:"section_id"` // 分部id
	Title     string                               `json:"title"`      // 视频合集中分部标题
	Type      int                                  `json:"type"`       // （？）。作用尚不明确
	Episodes  []GetVideoDetailInfoUgcSeasonEpisode `json:"episodes"`   // 视频合集中分部的视频列表
}

// GetVideoDetailInfoUgcSeasonEpisode 是「获取视频超详细信息」ugc_season.sections[].episodes[] 中的对象。
type GetVideoDetailInfoUgcSeasonEpisode struct {
	SeasonID  int                                   `json:"season_id"`  // 分部中视频所属视频合集id
	SectionID int                                   `json:"section_id"` // 分部中视频所属视频合集分部id
	Id        int                                   `json:"id"`         // 分部中视频id
	Aid       int                                   `json:"aid"`        // 视频aid
	Cid       int                                   `json:"cid"`        // 视频cid
	Title     string                                `json:"title"`      // 视频标题
	Attribute int                                   `json:"attribute"`  // 稿件属性位。文档标为已弃用，示例中仍在返回
	Arc       GetVideoDetailInfoUgcSeasonEpisodeArc `json:"arc"`        // 视频详细信息
	Page      GetVideoDetailInfoPage                `json:"page"`       // 视频分P信息
	Bvid      string                                `json:"bvid"`       // 视频bvid
	Pages     []GetVideoDetailInfoPage              `json:"pages"`      // 视频分P列表
}

// GetVideoDetailInfoUgcSeasonEpisodeArc 是「获取视频超详细信息」ugc_season.sections[].episodes[].arc 对象。
type GetVideoDetailInfoUgcSeasonEpisodeArc struct {
	Aid                int                                    `json:"aid"`                  // 稿件avid
	Videos             int                                    `json:"videos"`               // 稿件分P总数。默认为1
	TypeID             int                                    `json:"type_id"`              // 分区tid
	TypeName           string                                 `json:"type_name"`            // 子分区名称
	Copyright          int                                    `json:"copyright"`            // 视频类型。1：原创。2：转载
	Pic                string                                 `json:"pic"`                  // 稿件封面图片url
	Title              string                                 `json:"title"`                // 稿件标题
	PubDate            int                                    `json:"pubdate"`              // 稿件发布时间。秒级时间戳
	Ctime              int                                    `json:"ctime"`                // 用户投稿时间。秒级时间戳
	Desc               string                                 `json:"desc"`                 // 视频简介
	State              int                                    `json:"state"`                // 视频状态
	Duration           int                                    `json:"duration"`             // 稿件总时长。单位为秒
	Rights             GetVideoDetailInfoRights               `json:"rights"`               // 视频属性标志
	Author             GetVideoDetailInfoOwner                `json:"author"`               // UP主信息
	Stat               GetVideoDetailInfoUgcSeasonEpisodeStat `json:"stat"`                 // 视频状态数
	Dynamic            string                                 `json:"dynamic"`              // 同步发布的动态的文字内容
	Dimension          GetVideoDetailInfoDimension            `json:"dimension"`            // 视频1P分辨率
	DescV2             any                                    `json:"desc_v2"`              // 新版视频简介。合集示例中为 null
	IsChargeableSeason bool                                   `json:"is_chargeable_season"` // 是否为付费合集
	IsBlooper          bool                                   `json:"is_blooper"`           // 是否为花絮
	EnableVt           int                                    `json:"enable_vt"`            // （？）。作用尚不明确
	VtDisplay          string                                 `json:"vt_display"`           // （？）。作用尚不明确
}

// GetVideoDetailInfoUgcSeasonEpisodeStat 是「获取视频超详细信息」ugc_season.sections[].episodes[].arc.stat 对象。
type GetVideoDetailInfoUgcSeasonEpisodeStat struct {
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

// GetVideoDetailInfoUgcSeasonStat 是「获取视频超详细信息」ugc_season.stat 对象。
type GetVideoDetailInfoUgcSeasonStat struct {
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

// GetVideoDetailInfoCard 是「获取视频超详细信息」的 Card 对象。
type GetVideoDetailInfoCard struct {
	Card         GetVideoDetailInfoCardUser  `json:"card"`          // UP主名片信息
	Space        GetVideoDetailInfoCardSpace `json:"space"`         // 主页头图
	Following    bool                        `json:"following"`     // 是否关注此用户。需要登录(Cookie)。未登录为false
	ArchiveCount int                         `json:"archive_count"` // 用户稿件数
	ArticleCount int                         `json:"article_count"` // 用户专栏数
	Follower     int                         `json:"follower"`      // 粉丝数
	LikeNum      int                         `json:"like_num"`      // UP主获赞次数
}

// GetVideoDetailInfoCardUser 是「获取视频超详细信息」Card.card 对象。
type GetVideoDetailInfoCardUser struct {
	Mid            string                           `json:"mid"`              // 用户mid
	Name           string                           `json:"name"`             // 用户昵称
	Approve        bool                             `json:"approve"`          // false。作用尚不明确
	Sex            string                           `json:"sex"`              // 用户性别。男 女 保密
	Rank           string                           `json:"rank"`             // 10000。作用尚不明确
	Face           string                           `json:"face"`             // 用户头像链接
	FaceNft        int                              `json:"face_nft"`         // 是否为 nft 头像。0不是nft头像。1是 nft 头像
	DisplayRank    string                           `json:"DisplayRank"`      // 0。作用尚不明确
	Regtime        int                              `json:"regtime"`          // 0。作用尚不明确
	Spacesta       int                              `json:"spacesta"`         // 0。作用尚不明确
	Birthday       string                           `json:"birthday"`         // 空。作用尚不明确
	Place          string                           `json:"place"`            // 空。作用尚不明确
	Description    string                           `json:"description"`      // 空。作用尚不明确
	Article        int                              `json:"article"`          // 0。作用尚不明确
	Attentions     []any                            `json:"attentions"`       // 空。作用尚不明确
	Fans           int                              `json:"fans"`             // 粉丝数
	Friend         int                              `json:"friend"`           // 关注数
	Attention      int                              `json:"attention"`        // 关注数
	Sign           string                           `json:"sign"`             // 签名
	LevelInfo      GetVideoDetailInfoLevelInfo      `json:"level_info"`       // 等级
	Pendant        GetVideoDetailInfoPendant        `json:"pendant"`          // 挂件
	Nameplate      GetVideoDetailInfoNameplate      `json:"nameplate"`        // 勋章
	Official       GetVideoDetailInfoOfficial       `json:"Official"`         // 认证信息
	OfficialVerify GetVideoDetailInfoOfficialVerify `json:"official_verify"`  // 认证信息2
	Vip            GetVideoDetailInfoCardVip        `json:"vip"`              // 大会员状态
	IsSeniorMember int                              `json:"is_senior_member"` // 是否为硬核会员。0：否。1：是
}

// GetVideoDetailInfoCardSpace 是「获取视频超详细信息」Card.space 对象。
type GetVideoDetailInfoCardSpace struct {
	SImg string `json:"s_img"` // 主页头图url 小图
	LImg string `json:"l_img"` // 主页头图url 正常
}

// GetVideoDetailInfoLevelInfo 是「获取视频超详细信息」Card.card.level_info 对象。
type GetVideoDetailInfoLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetVideoDetailInfoPendant 是「获取视频超详细信息」Card.card.pendant 对象。
type GetVideoDetailInfoPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetVideoDetailInfoNameplate 是「获取视频超详细信息」Card.card.nameplate 对象。
type GetVideoDetailInfoNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetVideoDetailInfoOfficial 是「获取视频超详细信息」Card.card.Official 对象。
type GetVideoDetailInfoOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetVideoDetailInfoOfficialVerify 是「获取视频超详细信息」Card.card.official_verify 对象。
type GetVideoDetailInfoOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetVideoDetailInfoCardVip 是「获取视频超详细信息」Card.card.vip 对象。
type GetVideoDetailInfoCardVip struct {
	Type               int                           `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                           `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                           `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                           `json:"vip_pay_type"`         // 支付类型
	ThemeType          int                           `json:"theme_type"`           // 0。作用尚不明确
	Label              GetVideoDetailInfoLabel       `json:"label"`                // 会员标签
	AvatarSubscript    int                           `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                        `json:"nickname_color"`       // 会员昵称颜色
	Role               int                           `json:"role"`                 // 大角色类型
	AvatarSubscriptURL string                        `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                           `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                           `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetVideoDetailInfoVipOttInfo  `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetVideoDetailInfoVipSuperVip `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                           `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetVideoDetailInfoAvatarIcon  `json:"avatar_icon"`          // 大会员角标信息
}

// GetVideoDetailInfoAvatarIcon 是「获取视频超详细信息」Card.card.vip.avatar_icon 对象。
type GetVideoDetailInfoAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetVideoDetailInfoTag 是「获取视频超详细信息」Tags 数组中的对象。
type GetVideoDetailInfoTag struct {
	TagID        int                        `json:"tag_id"`        // tag_id
	TagName      string                     `json:"tag_name"`      // TAG名称
	Cover        string                     `json:"cover"`         // TAG图片url
	HeadCover    string                     `json:"head_cover"`    // TAG页面头图url
	Content      string                     `json:"content"`       // TAG介绍
	ShortContent string                     `json:"short_content"` // TAG简介
	Type         int                        `json:"type"`          // ？？？
	State        int                        `json:"state"`         // 0
	Ctime        int                        `json:"ctime"`         // 创建时间。时间戳
	Count        GetVideoDetailInfoTagCount `json:"count"`         // 状态数
	IsAtten      int                        `json:"is_atten"`      // 是否关注。需要登录(Cookie)。未登录为0
	Likes        int                        `json:"likes"`         // 0。作用尚不明确
	Hates        int                        `json:"hates"`         // 0。作用尚不明确
	Attribute    int                        `json:"attribute"`     // 0。作用尚不明确
	Liked        int                        `json:"liked"`         // 是否已经点赞。需要登录(Cookie)。未登录为0
	Hated        int                        `json:"hated"`         // 是否已经点踩。需要登录(Cookie)。未登录为0
	ExtraAttr    int                        `json:"extra_attr"`    // ? ? ?
}

// GetVideoDetailInfoTagCount 是「获取视频超详细信息」Tags[].count 对象。
type GetVideoDetailInfoTagCount struct {
	View  int `json:"view"`  // 0。作用尚不明确
	Use   int `json:"use"`   // 视频添加TAG数
	Atten int `json:"atten"` // TAG关注
}

// GetVideoDetailInfoEmergency 是「获取视频超详细信息」的 emergency 对象。
type GetVideoDetailInfoEmergency struct {
	NoLike  bool `json:"no_like"`  // 是否不显示点赞按钮
	NoCoin  bool `json:"no_coin"`  // 是否不显示投币按钮
	NoFav   bool `json:"no_fav"`   // 是否不显示收藏按钮
	NoShare bool `json:"no_share"` // 是否不显示分享按钮
}

// GetVideoDetailInfoReply 是「获取视频超详细信息」的 Reply 对象（评论区热评）。
type GetVideoDetailInfoReply struct {
	Page    GetVideoDetailInfoReplyPage   `json:"page"`    // 页面信息
	Replies []GetVideoDetailInfoReplyItem `json:"replies"` // 热评列表
}

// GetVideoDetailInfoReplyPage 是「获取视频超详细信息」Reply.page 对象。
type GetVideoDetailInfoReplyPage struct {
	Num    int `json:"num"`    // 当前页码
	Size   int `json:"size"`   // 每页项数
	Count  int `json:"count"`  // 根评论条数
	Acount int `json:"acount"` // 总计评论条数
}

// GetVideoDetailInfoReplyItem 是「获取视频超详细信息」Reply.replies 数组中的对象。
type GetVideoDetailInfoReplyItem struct {
	Rpid         int                                `json:"rpid"`          // 评论 rpid
	Oid          int                                `json:"oid"`           // 评论区对象 id
	Type         int                                `json:"type"`          // 评论区类型代码
	Mid          int                                `json:"mid"`           // 发送者 mid
	Root         int                                `json:"root"`          // 根评论 rpid
	Parent       int                                `json:"parent"`        // 回复父评论 rpid
	Dialog       int                                `json:"dialog"`        // 回复对方 rpid
	Count        int                                `json:"count"`         // 二级评论条数
	Rcount       int                                `json:"rcount"`        // 回复评论条数
	Floor        int                                `json:"floor"`         // 评论楼层号
	State        int                                `json:"state"`         // (?)
	FansGrade    int                                `json:"fansgrade"`     // 是否具有粉丝标签。0：无。1：有
	Attr         int                                `json:"attr"`          // 某属性位？
	Ctime        int                                `json:"ctime"`         // 评论发送时间。时间戳
	RpidStr      string                             `json:"rpid_str"`      // 评论rpid。字串格式
	RootStr      string                             `json:"root_str"`      // 根评论rpid。字串格式
	ParentStr    string                             `json:"parent_str"`    // 回复父评论rpid。字串格式
	Like         int                                `json:"like"`          // 评论获赞数
	Action       int                                `json:"action"`        // 当前用户操作状态
	Member       GetVideoDetailInfoReplyMember      `json:"member"`        // 评论发送者信息
	Content      GetVideoDetailInfoReplyContent     `json:"content"`       // 评论信息
	Replies      []*GetVideoDetailInfoReplyItem     `json:"replies"`       // 评论回复条目预览。仅嵌套一层。否则为 null
	Assist       int                                `json:"assist"`        // (?)
	Folder       GetVideoDetailInfoReplyFolder      `json:"folder"`        // 折叠信息
	UpAction     GetVideoDetailInfoReplyUpAction    `json:"up_action"`     // 评论 UP 主操作信息
	ShowFollow   bool                               `json:"show_follow"`   // (?)
	Invisible    bool                               `json:"invisible"`     // 评论是否被隐藏
	CardLabel    []GetVideoDetailInfoReplyCardLabel `json:"card_label"`    // 右上角卡片标签信息
	ReplyControl GetVideoDetailInfoReplyControl     `json:"reply_control"` // 评论提示文案信息
}

// GetVideoDetailInfoReplyMember 是「获取视频超详细信息」Reply.replies[].member 对象。
type GetVideoDetailInfoReplyMember struct {
	Mid            string                                `json:"mid"`             // 发送者 mid
	Uname          string                                `json:"uname"`           // 发送者昵称
	Sex            string                                `json:"sex"`             // 发送者性别。男 女 保密
	Sign           string                                `json:"sign"`            // 发送者签名
	Avatar         string                                `json:"avatar"`          // 发送者头像 url
	Rank           string                                `json:"rank"`            // (?)
	DisplayRank    string                                `json:"DisplayRank"`     // (?)
	LevelInfo      GetVideoDetailInfoReplyLevelInfo      `json:"level_info"`      // 发送者等级
	Pendant        GetVideoDetailInfoReplyPendant        `json:"pendant"`         // 发送者头像框信息
	Nameplate      GetVideoDetailInfoReplyNameplate      `json:"nameplate"`       // 发送者勋章信息
	OfficialVerify GetVideoDetailInfoReplyOfficialVerify `json:"official_verify"` // 发送者认证信息
	Vip            GetVideoDetailInfoReplyVip            `json:"vip"`             // 发送者会员信息
	FansDetail     *GetVideoDetailInfoReplyFansDetail    `json:"fans_detail"`     // 发送者粉丝标签
	Following      int                                   `json:"following"`       // 是否关注该用户。需要登录。否则恒为0
	IsFollowed     int                                   `json:"is_followed"`     // 是否被该用户关注。需要登录。否则恒为0
	UserSailing    GetVideoDetailInfoReplyUserSailing    `json:"user_sailing"`    // 发送者评论条目装扮信息
	IsContractor   bool                                  `json:"is_contractor"`   // 是否为合作用户？
	ContractDesc   string                                `json:"contract_desc"`   // 合作用户说明？
}

// GetVideoDetailInfoReplyLevelInfo 是「获取视频超详细信息」Reply.replies[].member.level_info 对象。
type GetVideoDetailInfoReplyLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetVideoDetailInfoReplyPendant 是「获取视频超详细信息」Reply.replies[].member.pendant 对象。
type GetVideoDetailInfoReplyPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetVideoDetailInfoReplyNameplate 是「获取视频超详细信息」Reply.replies[].member.nameplate 对象。
type GetVideoDetailInfoReplyNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetVideoDetailInfoReplyOfficialVerify 是「获取视频超详细信息」Reply.replies[].member.official_verify 对象。
type GetVideoDetailInfoReplyOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetVideoDetailInfoReplyVip 是「获取视频超详细信息」Reply.replies[].member.vip 对象。
type GetVideoDetailInfoReplyVip struct {
	VipType            int                             `json:"vipType"`              // 大会员类型。0：无。1：月会员。2：年以上会员
	VipDueDate         int                             `json:"vipDueDate"`           // 大会员到期时间。毫秒 时间戳
	DueRemark          string                          `json:"dueRemark"`            // (?)
	AccessStatus       int                             `json:"accessStatus"`         // (?)
	VipStatus          int                             `json:"vipStatus"`            // 大会员状态。0：无。1：有
	VipStatusWarn      string                          `json:"vipStatusWarn"`        // (?)
	ThemeType          int                             `json:"theme_type"`           // 会员样式 id
	Label              GetVideoDetailInfoReplyVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                             `json:"avatar_subscript"`     // (?)
	AvatarSubscriptURL string                          `json:"avatar_subscript_url"` // (?)
	NicknameColor      string                          `json:"nickname_color"`       // 昵称颜色
}

// GetVideoDetailInfoReplyVipLabel 是「获取视频超详细信息」Reply.replies[].member.vip.label 对象。
type GetVideoDetailInfoReplyVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetVideoDetailInfoReplyFansDetail 是「获取视频超详细信息」Reply.replies[].member.fans_detail 对象。
type GetVideoDetailInfoReplyFansDetail struct {
	UID          int    `json:"uid"`           // 用户 mid
	MedalID      int    `json:"medal_id"`      // 粉丝标签 id
	MedalName    string `json:"medal_name"`    // 粉丝标签名
	Score        int    `json:"score"`         // (?)
	Level        int    `json:"level"`         // 当前标签等级
	Intimacy     int    `json:"intimacy"`      // (?)
	MasterStatus int    `json:"master_status"` // (?)
	IsReceive    int    `json:"is_receive"`    // (?)
}

// GetVideoDetailInfoReplyUserSailing 是「获取视频超详细信息」Reply.replies[].member.user_sailing 对象。
type GetVideoDetailInfoReplyUserSailing struct {
	Pendant         *GetVideoDetailInfoReplyPendant `json:"pendant"`           // 头像框信息
	Cardbg          *GetVideoDetailInfoReplyCardbg  `json:"cardbg"`            // 评论卡片装扮
	CardbgWithFocus any                             `json:"cardbg_with_focus"` // (?)
}

// GetVideoDetailInfoReplyCardbg 是「获取视频超详细信息」Reply.replies[].member.user_sailing.cardbg 对象。
type GetVideoDetailInfoReplyCardbg struct {
	ID      int                        `json:"id"`       // 评论条目装扮 id
	Name    string                     `json:"name"`     // 评论条目装扮名称
	Image   string                     `json:"image"`    // 评论条目装扮图片 url
	JumpURL string                     `json:"jump_url"` // 评论条目装扮商城页面 url
	Fan     GetVideoDetailInfoReplyFan `json:"fan"`      // 粉丝专属信息
	Type    string                     `json:"type"`     // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
}

// GetVideoDetailInfoReplyFan 是「获取视频超详细信息」Reply.replies[].member.user_sailing.cardbg.fan 对象。
type GetVideoDetailInfoReplyFan struct {
	IsFan   int    `json:"is_fan"`   // 是否为粉丝专属装扮。0：否。1：是
	Number  int    `json:"number"`   // 粉丝专属编号
	Color   string `json:"color"`    // 数字颜色。颜色码
	Name    string `json:"name"`     // 装扮名称
	NumDesc string `json:"num_desc"` // 粉丝专属编号。字串格式
}

// GetVideoDetailInfoReplyContent 是「获取视频超详细信息」Reply.replies[].content 对象。
type GetVideoDetailInfoReplyContent struct {
	Message  string                           `json:"message"`  // 评论内容
	Plat     int                              `json:"plat"`     // 评论发送端
	Device   string                           `json:"device"`   // 评论发送平台设备
	Members  []GetVideoDetailInfoReplyMember  `json:"members"`  // at 到的用户信息
	Emote    any                              `json:"emote"`    // 需要渲染的表情转义
	JumpURL  any                              `json:"jump_url"` // 需要高亮的超链转义
	MaxLine  int                              `json:"max_line"` // 6。收起最大行数
	Pictures []GetVideoDetailInfoReplyPicture `json:"pictures"` // 评论图片数组
}

// GetVideoDetailInfoReplyPicture 是「获取视频超详细信息」Reply.replies[].content.pictures 数组中的对象。
type GetVideoDetailInfoReplyPicture struct {
	ImgSrc    string  `json:"img_src"`    // 图片地址
	ImgWidth  int     `json:"img_width"`  // 图片宽度
	ImgHeight int     `json:"img_height"` // 图片高度
	ImgSize   float64 `json:"img_size"`   // 图片大小。单位KB
}

// GetVideoDetailInfoReplyFolder 是「获取视频超详细信息」Reply.replies[].folder 对象。
type GetVideoDetailInfoReplyFolder struct {
	HasFolded bool   `json:"has_folded"` // 是否有被折叠的二级评论
	IsFolded  bool   `json:"is_folded"`  // 评论是否被折叠
	Rule      string `json:"rule"`       // 相关规则页面 url
}

// GetVideoDetailInfoReplyUpAction 是「获取视频超详细信息」Reply.replies[].up_action 对象。
type GetVideoDetailInfoReplyUpAction struct {
	Like  bool `json:"like"`  // 是否UP主觉得很赞
	Reply bool `json:"reply"` // 是否被UP主回复
}

// GetVideoDetailInfoReplyCardLabel 是「获取视频超详细信息」Reply.replies[].card_label 数组中的对象。
type GetVideoDetailInfoReplyCardLabel struct {
	Rpid             int    `json:"rpid"`              // 评论 rpid
	TextContent      string `json:"text_content"`      // 标签文本。已知有妙评
	TextColorDay     string `json:"text_color_day"`    // 日间文本颜色
	TextColorNight   string `json:"text_color_night"`  // 夜间文本颜色
	LabelColorDay    string `json:"label_color_day"`   // 日间标签颜色
	LabelColorNight  string `json:"label_color_night"` // 夜间标签颜色
	Image            string `json:"image"`             // 作用不明
	Type             int    `json:"type"`              // 1。作用不明
	Background       string `json:"background"`        // 背景图片 url
	BackgroundWidth  int    `json:"background_width"`  // 背景图片宽度
	BackgroundHeight int    `json:"background_height"` // 背景图片高度
	JumpURL          string `json:"jump_url"`          // 跳转链接
	Effect           int    `json:"effect"`            // 0。作用不明
	EffectStartTime  int    `json:"effect_start_time"` // 0
}

// GetVideoDetailInfoReplyControl 是「获取视频超详细信息」Reply.replies[].reply_control 对象。
type GetVideoDetailInfoReplyControl struct {
	SubReplyEntryText string `json:"sub_reply_entry_text"` // 回复提示。共 xx 条回复
	SubReplyTitleText string `json:"sub_reply_title_text"` // 回复提示。相关回复共有 xx 条
	TimeDesc          string `json:"time_desc"`            // 时间提示
	Location          string `json:"location"`             // IP属地
}
