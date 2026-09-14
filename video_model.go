package bilibili

// 视频详情（GetVideoInfo / 推荐列表）响应模型。
//
// Owner 等作者模型保留独立定义，不与其他业务的作者类型合并。

type DescV2 struct {
	RawText string `json:"raw_text"` // 简介内容。type=1时显示原文。type=2时显示'@'+raw_text+' '并链接至biz_id的主页
	Type    int    `json:"type"`     // 类型。1：普通，2：@他人
	BizId   int    `json:"biz_id"`   // 被@用户的mid。=0，当type=1
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
	Id          int                 `json:"id"`           // 字幕id
	Lan         string              `json:"lan"`          // 字幕语言
	LanDoc      string              `json:"lan_doc"`      // 字幕语言名称
	IsLock      bool                `json:"is_lock"`      // 是否锁定
	AuthorMid   int                 `json:"author_mid"`   // 字幕上传者mid
	SubtitleUrl string              `json:"subtitle_url"` // json格式字幕文件url
	Author      VideoSubtitleAuthor `json:"author"`       // 字幕上传者信息
}

type VideoSubtitles struct {
	AllowSubmit bool            `json:"allow_submit"` // 是否允许提交字幕
	List        []VideoSubtitle `json:"list"`         // 字幕列表
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
	UrlImageAniCut string `json:"url_image_ani_cut"` // 某url？
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
	Id              int        `json:"id"`                // 视频 avid / 直播间 id
	IsFollowed      int        `json:"is_followed"`       // 是否已关注
	IsStock         int        `json:"is_stock"`          // 未知作用
	OgvInfo         any        `json:"ogv_info"`          // 通常为 null
	Owner           Owner      `json:"owner"`             // 视频 UP 主信息
	Pic             string     `json:"pic"`               // 视频封面
	Pic43           string     `json:"pic_4_3"`           // 4:3 比例封面
	Pos             int        `json:"pos"`               // 位置
	Pubdate         int        `json:"pubdate"`           // 发布时间（秒级时间戳）
	RcmdReason      RcmdReason `json:"rcmd_reason"`       // 推荐理由
	RoomInfo        any        `json:"room_info"`         // 通常为 null
	ShowInfo        int        `json:"show_info"`         // 展示信息（1: 普通视频, 0: 直播）
	Stat            any        `json:"stat"`              // 视频状态信息
	Title           string     `json:"title"`             // 视频标题
	TrackId         string     `json:"track_id"`          // 跟踪标识
	Uri             string     `json:"uri"`               // 目标页 URI
	VtDisplay       string     `json:"vt_display"`        // 未知作用
}
type VideoInfo struct {
	Bvid               string        `json:"bvid"`         // 稿件bvid
	Aid                int           `json:"aid"`          // 稿件avid
	Videos             int           `json:"videos"`       // 稿件分P总数。默认为1
	Tid                int           `json:"tid"`          // 分区tid
	Tname              string        `json:"tname"`        // 子分区名称
	Copyright          int           `json:"copyright"`    // 视频类型。1：原创。2：转载
	Pic                string        `json:"pic"`          // 稿件封面图片url
	Title              string        `json:"title"`        // 稿件标题
	Pubdate            int           `json:"pubdate"`      // 稿件发布时间。秒级时间戳
	Ctime              int           `json:"ctime"`        // 用户投稿时间。秒级时间戳
	Desc               string        `json:"desc"`         // 视频简介
	DescV2             []DescV2      `json:"desc_v2"`      // 新版视频简介
	State              int           `json:"state"`        // 视频状态。详情见[属性数据文档](attribute_data.md#state字段值(稿件状态))
	Duration           int           `json:"duration"`     // 稿件总时长(所有分P)。单位为秒
	Forward            int           `json:"forward"`      // 撞车视频跳转avid。仅撞车视频存在此字段
	MissionId          int           `json:"mission_id"`   // 稿件参与的活动id
	RedirectUrl        string        `json:"redirect_url"` // 重定向url。仅番剧或影视视频存在此字段。用于番剧&影视的av/bv->ep
	Rights             VideoRights   `json:"rights"`       // 视频属性标志
	Owner              Owner         `json:"owner"`        // 视频UP主信息
	Stat               VideoStat     `json:"stat"`         // 视频状态数
	Dynamic            string        `json:"dynamic"`      // 视频同步发布的的动态的文字内容
	Cid                int           `json:"cid"`          // 视频1P cid
	Dimension          Dimension     `json:"dimension"`    // 视频1P分辨率
	SeasonId           int           `json:"season_id"`    // 合集id
	Premiere           any           `json:"premiere"`     // null
	TeenageMode        int           `json:"teenage_mode"`
	IsChargeableSeason bool          `json:"is_chargeable_season"`
	IsStory            bool          `json:"is_story"`
	NoCache            bool          `json:"no_cache"` // 作用尚不明确
	Pages              []VideoPage   `json:"pages"`    // 视频分P列表
	Subtitle           VideoSubtitle `json:"subtitle"` // 视频CC字幕信息
	Staff              []Staff       `json:"staff"`    // 合作成员列表。非合作视频无此项
	IsSeasonDisplay    bool          `json:"is_season_display"`
	UserGarb           UserGarb      `json:"user_garb"` // 用户装扮信息
	HonorReply         HonorReply    `json:"honor_reply"`
	LikeIcon           string        `json:"like_icon"`
	ArgueInfo          ArgueInfo     `json:"argue_info"` // 争议/警告信息
}
