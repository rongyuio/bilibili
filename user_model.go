package bilibili

// 用户空间详情（GetUserSpaceDetail）响应模型。
//
// SpaceVip / UserCardVip / MyVip 等会员变体各自保留独立类型，
// 仅在字段完全一致时才考虑合并。

type GetUserSpaceDetailParam struct {
	Mid int `json:"mid"` // 目标用户mid
}

type SpaceVip struct {
	Type               int    `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int    `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int    `json:"due_date"`             // 会员过期时间。毫秒时间戳
	VipPayType         int    `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int    `json:"theme_type"`           // 0。作用尚不明确
	Label              Label  `json:"label"`                // 会员标签
	AvatarSubscript    int    `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int    `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptUrl string `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int    `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int    `json:"tv_vip_pay_type"`      // 电视大会员支付类型
}

type Medal struct {
	Uid              int    `json:"uid"`                // 此用户mid
	TargetId         int    `json:"target_id"`          // 粉丝勋章所属UP的mid
	MedalId          int    `json:"medal_id"`           // 粉丝勋章id
	Level            int    `json:"level"`              // 粉丝勋章等级
	MedalName        string `json:"medal_name"`         // 粉丝勋章名称
	MedalColor       int    `json:"medal_color"`        // 颜色
	Intimacy         int    `json:"intimacy"`           // 当前亲密度
	NextIntimacy     int    `json:"next_intimacy"`      // 下一等级所需亲密度
	DayLimit         int    `json:"day_limit"`          // 每日亲密度获取上限
	TodayFeed        int    `json:"today_feed"`         // 今日已获得亲密度
	MedalColorStart  int    `json:"medal_color_start"`  // 粉丝勋章颜色。十进制数，可转为十六进制颜色代码
	MedalColorEnd    int    `json:"medal_color_end"`    // 粉丝勋章颜色。十进制数，可转为十六进制颜色代码
	MedalColorBorder int    `json:"medal_color_border"` // 粉丝勋章边框颜色。十进制数，可转为十六进制颜色代码
	IsLighted        int    `json:"is_lighted"`
	LightStatus      int    `json:"light_status"`
	WearingStatus    int    `json:"wearing_status"` // 当前是否佩戴。0：未佩戴。1：已佩戴
	Score            int    `json:"score"`
}

type FansMedal struct {
	Show  bool  `json:"show"`
	Wear  bool  `json:"wear"`  // 是否佩戴了粉丝勋章
	Medal Medal `json:"medal"` // 粉丝勋章信息
}

type SysNotice struct {
	Id         int    `json:"id"`          // id
	Content    string `json:"content"`     // 显示文案
	Url        string `json:"url"`         // 跳转地址
	NoticeType int    `json:"notice_type"` // 提示类型。1,2
	Icon       string `json:"icon"`        // 前缀图标
	TextColor  string `json:"text_color"`  // 文字颜色
	BgColor    string `json:"bg_color"`    // 背景颜色
}

type WatchedShow struct {
	Switch       bool   `json:"switch"` // ?
	Num          int    `json:"num"`    // total watched users
	TextSmall    string `json:"text_small"`
	TextLarge    string `json:"text_large"`
	Icon         string `json:"icon"`          // watched icon url
	IconLocation string `json:"icon_location"` // ?
	IconWeb      string `json:"icon_web"`      // watched icon url
}

type LiveRoom struct {
	Roomstatus    int         `json:"roomStatus"` // 直播间状态。0：无房间。1：有房间
	Livestatus    int         `json:"liveStatus"` // 直播状态。0：未开播。1：直播中
	Url           string      `json:"url"`        // 直播间网页 url
	Title         string      `json:"title"`      // 直播间标题
	Cover         string      `json:"cover"`      // 直播间封面 url
	WatchedShow   WatchedShow `json:"watched_show"`
	Roomid        int         `json:"roomid"`         // 直播间 id(短号)
	Roundstatus   int         `json:"roundStatus"`    // 轮播状态。0：未轮播。1：轮播
	BroadcastType int         `json:"broadcast_type"` // 0
}

type School struct {
	Name string `json:"name"` // 就读大学名称。没有则为空
}

type Profession struct {
	Name       string `json:"name"`       // 资质名称
	Department string `json:"department"` // 职位
	Title      string `json:"title"`      // 所属机构
	IsShow     int    `json:"is_show"`    // 是否显示。0：不显示。1：显示
}

type Elec struct {
	ShowInfo struct {
		Show    bool   `json:"show"`     // 是否开通了充电
		State   int    `json:"state"`    // 状态。-1：未开通。1：已开通
		Title   string `json:"title"`    // 空串
		Icon    string `json:"icon"`     // 空串
		JumpUrl string `json:"jump_url"` // 空串
	} `json:"show_info"`
}

type Contract struct {
	IsDisplay       bool `json:"is_display"`        // true/false。在页面中未使用此字段
	IsFollowDisplay bool `json:"is_follow_display"` // 是否在显示老粉计划。true：显示。false：不显示
}

type UserSpaceDetail struct {
	Mid            int       `json:"mid"`           // mid
	Name           string    `json:"name"`          // 昵称
	Sex            string    `json:"sex"`           // 性别。男/女/保密
	Face           string    `json:"face"`          // 头像链接
	FaceNft        int       `json:"face_nft"`      // 是否为 NFT 头像。0：不是 NFT 头像。1：是 NFT 头像
	FaceNftType    int       `json:"face_nft_type"` // NFT 头像类型？
	Sign           string    `json:"sign"`          // 签名
	Rank           int       `json:"rank"`          // 用户权限等级。目前应该无任何作用。5000：0级未答题。10000：普通会员。20000：字幕君。25000：VIP。30000：真·职人。32000：管理员
	Level          int       `json:"level"`         // 当前等级。0-6 级
	Jointime       int       `json:"jointime"`      // 注册时间。此接口返回恒为0
	Moral          int       `json:"moral"`         // 节操值。此接口返回恒为0
	Silence        int       `json:"silence"`       // 封禁状态。0：正常。1：被封
	Coins          int       `json:"coins"`         // 硬币数。需要登录（Cookie） 。只能查看自己的。默认为0
	FansBadge      bool      `json:"fans_badge"`    // 是否具有粉丝勋章。false：无。true：有
	FansMedal      FansMedal `json:"fans_medal"`    // 粉丝勋章信息
	Official       Official  `json:"official"`      // 认证信息
	Vip            SpaceVip  `json:"vip"`           // 会员信息
	Pendant        Pendant   `json:"pendant"`       // 头像框信息
	Nameplate      Nameplate `json:"nameplate"`     // 勋章信息
	UserHonourInfo struct {
		Mid    int    `json:"mid"`    // 0
		Colour string `json:"colour"` // null
		Tags   any    `json:"tags"`   // null
	} `json:"user_honour_info"` // （？）
	IsFollowed bool       `json:"is_followed"` // 是否关注此用户。true：已关注。false：未关注。需要登录（Cookie） 。未登录恒为false
	TopPhoto   string     `json:"top_photo"`   // 主页头图链接
	Theme      any        `json:"theme"`       // （？）
	SysNotice  SysNotice  `json:"sys_notice"`  // 系统通知。无内容则为空对象。主要用于展示如用户争议、纪念账号等等的小黄条
	LiveRoom   LiveRoom   `json:"live_room"`   // 直播间信息
	Birthday   string     `json:"birthday"`    // 生日。MM-DD。如设置隐私为空
	School     School     `json:"school"`      // 学校
	Profession Profession `json:"profession"`  // 专业资质信息
	Tags       any        `json:"tags"`        // 个人标签
	Series     struct {
		UserUpgradeStatus int  `json:"user_upgrade_status"` // (?)
		ShowUpgradeWindow bool `json:"show_upgrade_window"` // (?)
	} `json:"series"`
	IsSeniorMember int      `json:"is_senior_member"` // 是否为硬核会员。0：否。1：是
	McnInfo        any      `json:"mcn_info"`         // （？）
	GaiaResType    int      `json:"gaia_res_type"`    // （？）
	GaiaData       any      `json:"gaia_data"`        // （？）
	IsRisk         bool     `json:"is_risk"`          // （？）
	Elec           Elec     `json:"elec"`             // 充电信息
	Contract       Contract `json:"contract"`         // 是否显示老粉计划
}
