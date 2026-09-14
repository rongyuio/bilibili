package bilibili

// 用户空间、名片与用户视频相关响应模型。
//
// SpaceVip 与 CardVip 字段（名称、JSON 标签、类型）完全一致，合并为 CardVip
// 并保留 SpaceVip 别名；UserCardVip、MyVip 的字段名或字段集不同，各自保留。

type GetUserSpaceDetailParam struct {
	Mid int `json:"mid"` // 目标用户mid
}

// SpaceVip 是用户空间的大会员信息，与 CardVip 字段完全一致。
type SpaceVip = CardVip

type Medal struct {
	UID              int    `json:"uid"`                // 此用户mid
	TargetID         int    `json:"target_id"`          // 粉丝勋章所属UP的mid
	MedalID          int    `json:"medal_id"`           // 粉丝勋章id
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
	ID         int    `json:"id"`          // id
	Content    string `json:"content"`     // 显示文案
	URL        string `json:"url"`         // 跳转地址
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
	RoomStatus    int         `json:"roomStatus"` // 直播间状态。0：无房间。1：有房间
	LiveStatus    int         `json:"liveStatus"` // 直播状态。0：未开播。1：直播中
	URL           string      `json:"url"`        // 直播间网页 url
	Title         string      `json:"title"`      // 直播间标题
	Cover         string      `json:"cover"`      // 直播间封面 url
	WatchedShow   WatchedShow `json:"watched_show"`
	RoomID        int         `json:"roomid"`         // 直播间 id(短号)
	RoundStatus   int         `json:"roundStatus"`    // 轮播状态。0：未轮播。1：轮播
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

// ElecShowInfo 是充电信息。
type ElecShowInfo struct {
	Show    bool   `json:"show"`     // 是否开通了充电
	State   int    `json:"state"`    // 状态。-1：未开通。1：已开通
	Title   string `json:"title"`    // 空串
	Icon    string `json:"icon"`     // 空串
	JumpURL string `json:"jump_url"` // 空串
}

type Elec struct {
	ShowInfo ElecShowInfo `json:"show_info"`
}

type Contract struct {
	IsDisplay       bool `json:"is_display"`        // true/false。在页面中未使用此字段
	IsFollowDisplay bool `json:"is_follow_display"` // 是否在显示老粉计划。true：显示。false：不显示
}

// UserHonourInfo 是用户荣誉信息。
type UserHonourInfo struct {
	Mid    int    `json:"mid"`    // 0
	Colour string `json:"colour"` // null
	Tags   any    `json:"tags"`   // null
}

// UserSeries 是用户成长系列信息。
type UserSeries struct {
	UserUpgradeStatus int  `json:"user_upgrade_status"` // (?)
	ShowUpgradeWindow bool `json:"show_upgrade_window"` // (?)
}

type UserSpaceDetail struct {
	Mid            int            `json:"mid"`              // mid
	Name           string         `json:"name"`             // 昵称
	Sex            string         `json:"sex"`              // 性别。男/女/保密
	Face           string         `json:"face"`             // 头像链接
	FaceNft        int            `json:"face_nft"`         // 是否为 NFT 头像。0：不是 NFT 头像。1：是 NFT 头像
	FaceNftType    int            `json:"face_nft_type"`    // NFT 头像类型？
	Sign           string         `json:"sign"`             // 签名
	Rank           int            `json:"rank"`             // 用户权限等级。目前应该无任何作用。5000：0级未答题。10000：普通会员。20000：字幕君。25000：VIP。30000：真·职人。32000：管理员
	Level          int            `json:"level"`            // 当前等级。0-6 级
	Jointime       int            `json:"jointime"`         // 注册时间。此接口返回恒为0
	Moral          int            `json:"moral"`            // 节操值。此接口返回恒为0
	Silence        int            `json:"silence"`          // 封禁状态。0：正常。1：被封
	Coins          int            `json:"coins"`            // 硬币数。需要登录（Cookie） 。只能查看自己的。默认为0
	FansBadge      bool           `json:"fans_badge"`       // 是否具有粉丝勋章。false：无。true：有
	FansMedal      FansMedal      `json:"fans_medal"`       // 粉丝勋章信息
	Official       Official       `json:"official"`         // 认证信息
	Vip            SpaceVip       `json:"vip"`              // 会员信息
	Pendant        Pendant        `json:"pendant"`          // 头像框信息
	Nameplate      Nameplate      `json:"nameplate"`        // 勋章信息
	UserHonourInfo UserHonourInfo `json:"user_honour_info"` // （？）
	IsFollowed     bool           `json:"is_followed"`      // 是否关注此用户。true：已关注。false：未关注。需要登录（Cookie） 。未登录恒为false
	TopPhoto       string         `json:"top_photo"`        // 主页头图链接
	Theme          any            `json:"theme"`            // （？）
	SysNotice      SysNotice      `json:"sys_notice"`       // 系统通知。无内容则为空对象。主要用于展示如用户争议、纪念账号等等的小黄条
	LiveRoom       LiveRoom       `json:"live_room"`        // 直播间信息
	Birthday       string         `json:"birthday"`         // 生日。MM-DD。如设置隐私为空
	School         School         `json:"school"`           // 学校
	Profession     Profession     `json:"profession"`       // 专业资质信息
	Tags           any            `json:"tags"`             // 个人标签
	Series         UserSeries     `json:"series"`
	IsSeniorMember int            `json:"is_senior_member"` // 是否为硬核会员。0：否。1：是
	McnInfo        any            `json:"mcn_info"`         // （？）
	GaiaResType    int            `json:"gaia_res_type"`    // （？）
	GaiaData       any            `json:"gaia_data"`        // （？）
	IsRisk         bool           `json:"is_risk"`          // （？）
	Elec           Elec           `json:"elec"`             // 充电信息
	Contract       Contract       `json:"contract"`         // 是否显示老粉计划
}

type VideoArea struct {
	Count int    `json:"count"` // 投稿至该分区的视频数
	Name  string `json:"name"`  // 该分区名称
	Tid   int    `json:"tid"`   // 该分区tid
}

type UserVideo struct {
	Aid          int    `json:"aid"` // 稿件avid
	Attribute    int    `json:"attribute"`
	Author       string `json:"author"`      // 视频UP主。不一定为目标用户（合作视频）
	Bvid         string `json:"bvid"`        // 稿件bvid
	Comment      int    `json:"comment"`     // 视频评论数
	Copyright    string `json:"copyright"`   // 视频版权类型
	Created      int    `json:"created"`     // 投稿时间。时间戳
	Description  string `json:"description"` // 视频简介
	EnableVt     int    `json:"enable_vt"`
	HideClick    bool   `json:"hide_click"`     // false。作用尚不明确
	IsPay        int    `json:"is_pay"`         // 0。作用尚不明确
	IsUnionVideo int    `json:"is_union_video"` // 是否为合作视频。0：否。1：是
	Length       string `json:"length"`         // 视频长度。MM:SS
	Mid          int    `json:"mid"`            // 视频UP主mid。不一定为目标用户（合作视频）
	Meta         any    `json:"meta"`           // 无数据时为 null
	Pic          string `json:"pic"`            // 视频封面
	Play         int    `json:"play"`           // 视频播放次数
	Review       int    `json:"review"`         // 0。作用尚不明确
	Subtitle     string `json:"subtitle"`       // 空。作用尚不明确
	Title        string `json:"title"`          // 视频标题
	Typeid       int    `json:"typeid"`         // 视频分区tid
	VideoReview  int    `json:"video_review"`   // 视频弹幕数
}

type UserVideosList struct {
	Tlist map[int]VideoArea `json:"tlist"` // 投稿视频分区索引
	Vlist []UserVideo       `json:"vlist"` // 投稿视频列表
}

// UserVideoPage 空间投稿视频分页信息。
//
// 各接口分页字段命名不统一（此处为 pn/ps，分区列表为 num/size，
// 合集为 page_num/page_size），因此保留独立类型，不做跨接口合并。
type UserVideoPage struct {
	Count int `json:"count"` // 总计稿件数
	Pn    int `json:"pn"`    // 当前页码
	Ps    int `json:"ps"`    // 每页项数
}

type EpisodicButton struct {
	Text string `json:"text"` // 按钮文字
	URI  string `json:"uri"`  // 全部播放页url
}

type UserVideos struct {
	List           UserVideosList `json:"list"`            // 列表信息
	Page           UserVideoPage  `json:"page"`            // 页面信息
	EpisodicButton EpisodicButton `json:"episodic_button"` // “播放全部“按钮
	IsRisk         bool           `json:"is_risk"`
	GaiaResType    int            `json:"gaia_res_type"`
	GaiaData       any            `json:"gaia_data"`
}

type UserCardVip struct {
	VipType       int    `json:"vipType"`       // 大会员类型。0：无。1：月度大会员。2：年度及以上大会员
	DueRemark     string `json:"dueRemark"`     // 空。**作用尚不明确**
	AccessStatus  int    `json:"accessStatus"`  // 0。**作用尚不明确**
	VipStatus     int    `json:"vipStatus"`     // 大会员状态。0：无。1：有
	VipStatusWarn string `json:"vipStatusWarn"` // 空。**作用尚不明确**
	ThemeType     int    `json:"theme_type"`    // 0。**作用尚不明确**
}

type UserCardInfo struct {
	Mid            string         `json:"mid"`             // 用户mid
	Approve        bool           `json:"approve"`         // false。**作用尚不明确**
	Name           string         `json:"name"`            // 用户昵称
	Sex            string         `json:"sex"`             // 用户性别。男 女 保密
	Face           string         `json:"face"`            // 用户头像链接
	DisplayRank    string         `json:"DisplayRank"`     // 0。**作用尚不明确**
	Regtime        int            `json:"regtime"`         // 0。**作用尚不明确**
	Spacesta       int            `json:"spacesta"`        // 用户状态。0：正常。-2：被封禁
	Birthday       string         `json:"birthday"`        // 空。**作用尚不明确**
	Place          string         `json:"place"`           // 空。**作用尚不明确**
	Description    string         `json:"description"`     // 空。**作用尚不明确**
	Article        int            `json:"article"`         // 0。**作用尚不明确**
	Attentions     any            `json:"attentions"`      // 空。**作用尚不明确**
	Fans           int            `json:"fans"`            // 粉丝数
	Friend         int            `json:"friend"`          // 关注数
	Attention      int            `json:"attention"`       // 关注数
	Sign           string         `json:"sign"`            // 签名
	LevelInfo      LevelInfo      `json:"level_info"`      // 等级
	Pendant        Pendant        `json:"pendant"`         // 挂件
	Nameplate      Nameplate      `json:"nameplate"`       // 勋章
	Official       Official       `json:"Official"`        // 认证信息
	OfficialVerify OfficialVerify `json:"official_verify"` // 认证信息2
	Vip            UserCardVip    `json:"vip"`             // 大会员状态
	Space          CardSpace      `json:"space"`           // 主页头图
}

type UserCard struct {
	Card         UserCardInfo `json:"card"`          // 卡片信息
	Following    bool         `json:"following"`     // 是否关注此用户。true：已关注。false：未关注。需要登录(Cookie)。未登录为false
	ArchiveCount int          `json:"archive_count"` // 用户稿件数
	ArticleCount int          `json:"article_count"` // 0。**作用尚不明确**
	Follower     int          `json:"follower"`      // 粉丝数
	LikeNum      int          `json:"like_num"`      // 点赞数
}

type MyVip struct {
	Type            int    `json:"type"`             // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status          int    `json:"status"`           // 会员状态。0：无。1：有
	DueDate         int    `json:"due_date"`         // 会员过期时间。Unix时间戳(毫秒)
	ThemeType       int    `json:"theme_type"`       // 0。作用尚不明确
	Label           Label  `json:"label"`            // 会员标签
	AvatarSubscript int    `json:"avatar_subscript"` // 是否显示会员图标。0：不显示。1：显示
	NicknameColor   string `json:"nickname_color"`   // 会员昵称颜色。颜色码
}

type MyProfession struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ShowName string `json:"show_name"`
}

type MyUserSpaceDetail struct {
	Mid            int          `json:"mid"`             // mid
	Name           string       `json:"name"`            // 昵称
	Sex            string       `json:"sex"`             // 性别。男 女 保密
	Face           string       `json:"face"`            // 头像图片url
	Sign           string       `json:"sign"`            // 签名
	Rank           int          `json:"rank"`            // 10000。**作用尚不明确**
	Level          int          `json:"level"`           // 当前等级。0-6级
	Jointime       int          `json:"jointime"`        // 0。**作用尚不明确**
	Moral          int          `json:"moral"`           // 节操。默认70
	Silence        int          `json:"silence"`         // 封禁状态。0：正常。1：被封
	EmailStatus    int          `json:"email_status"`    // 已验证邮箱。0：未验证。1：已验证
	TelStatus      int          `json:"tel_status"`      // 已验证手机号。0：未验证。1：已验证
	Identification int          `json:"identification"`  // 1。**作用尚不明确**
	Vip            MyVip        `json:"vip"`             // 大会员状态
	Pendant        Pendant      `json:"pendant"`         // 头像框信息
	Nameplate      Nameplate    `json:"nameplate"`       // 勋章信息
	Official       Official     `json:"official"`        // 认证信息
	Birthday       int          `json:"birthday"`        // 生日。时间戳
	IsTourist      int          `json:"is_tourist"`      // 0。**作用尚不明确**
	IsFakeAccount  int          `json:"is_fake_account"` // 0。**作用尚不明确**
	PinPrompting   int          `json:"pin_prompting"`   // 0。**作用尚不明确**
	IsDeleted      int          `json:"is_deleted"`      // 0。**作用尚不明确**
	InRegAudit     int          `json:"in_reg_audit"`
	IsRipUser      bool         `json:"is_rip_user"`
	Profession     MyProfession `json:"profession"` // 专业资质
	Coins          float64      `json:"coins"`      // 硬币数
	Following      int          `json:"following"`  // 粉丝数
	Follower       int          `json:"follower"`   // 粉丝数
}

type JoinOldFansResult struct {
	AllowMessage bool   `json:"allow_message"` // true
	InputText    string `json:"input_text"`    // UP主加油！看好你噢
	InputTitle   string `json:"input_title"`   // 感谢你对UP主的特别支持，“老粉”可期！私信留言鼓励下TA吧
}

type FansSendMessageResult struct {
	SuccessToast string `json:"success_toast"` // "提交成功，UP主已收到留言~"
}

type BatchGetUserCardsResult struct {
	Mid     int    `json:"mid"`     // mid
	Name    string `json:"name"`    // 昵称
	Face    string `json:"face"`    // 头像链接
	Sign    string `json:"sign"`    // 签名
	Rank    int    `json:"rank"`    // 用户权限等级
	Level   int    `json:"level"`   // 当前等级。0-6 级
	Silence int    `json:"silence"` // 封禁状态。0：正常。1：被封
}
