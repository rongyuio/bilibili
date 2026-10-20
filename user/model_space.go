package user

// 用户空间、名片与投稿视频相关响应模型。
//
// 每个接口的返回体独立定义，嵌套对象也带方法前缀，避免接口间共用类型。

// GetUserVideosResult 是查询用户投稿视频明细的 data 对象。
type GetUserVideosResult struct {
	List           GetUserVideosList           `json:"list"`            // 列表信息
	Page           GetUserVideosPage           `json:"page"`            // 页面信息
	EpisodicButton GetUserVideosEpisodicButton `json:"episodic_button"` // “播放全部“按钮
	IsRisk         bool                        `json:"is_risk"`
	GaiaResType    int                         `json:"gaia_res_type"`
	GaiaData       any                         `json:"gaia_data"`
}

// GetUserVideosList 是投稿视频列表信息。
type GetUserVideosList struct {
	Tlist map[int]GetUserVideosArea `json:"tlist"` // 投稿视频分区索引
	Vlist []GetUserVideosItem       `json:"vlist"` // 投稿视频列表
}

// GetUserVideosArea 是投稿视频分区索引中的一项。
type GetUserVideosArea struct {
	Count int    `json:"count"` // 投稿至该分区的视频数
	Name  string `json:"name"`  // 该分区名称
	Tid   int    `json:"tid"`   // 该分区tid
}

// GetUserVideosItem 是投稿视频列表中的一项。
type GetUserVideosItem struct {
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

// GetUserVideosPage 是空间投稿视频分页信息。
type GetUserVideosPage struct {
	Count int `json:"count"` // 总计稿件数
	Pn    int `json:"pn"`    // 当前页码
	Ps    int `json:"ps"`    // 每页项数
}

// GetUserVideosEpisodicButton 是“播放全部“按钮信息。
type GetUserVideosEpisodicButton struct {
	Text string `json:"text"` // 按钮文字
	URI  string `json:"uri"`  // 全部播放页url
}

// GetUserSpaceDetailResult 是获取用户空间详细信息的 data 对象。
type GetUserSpaceDetailResult struct {
	Mid             int                              `json:"mid"`              // mid
	Name            string                           `json:"name"`             // 昵称
	Sex             string                           `json:"sex"`              // 性别。男/女/保密
	Face            string                           `json:"face"`             // 头像链接
	FaceNft         int                              `json:"face_nft"`         // 是否为 NFT 头像。0：不是 NFT 头像。1：是 NFT 头像
	FaceNftType     int                              `json:"face_nft_type"`    // NFT 头像类型？
	Sign            string                           `json:"sign"`             // 签名
	Rank            int                              `json:"rank"`             // 用户权限等级。目前应该无任何作用。5000：0级未答题。10000：普通会员。20000：字幕君。25000：VIP。30000：真·职人。32000：管理员
	Level           int                              `json:"level"`            // 当前等级。0-6 级
	Jointime        int                              `json:"jointime"`         // 注册时间。此接口返回恒为0
	Moral           int                              `json:"moral"`            // 节操值。此接口返回恒为0
	Silence         int                              `json:"silence"`          // 封禁状态。0：正常。1：被封
	Coins           int                              `json:"coins"`            // 硬币数。需要登录（Cookie） 。只能查看自己的。默认为0
	FansBadge       bool                             `json:"fans_badge"`       // 是否具有粉丝勋章。false：无。true：有
	FansMedal       GetUserSpaceDetailFansMedal      `json:"fans_medal"`       // 粉丝勋章信息
	Official        GetUserSpaceDetailOfficial       `json:"official"`         // 认证信息
	Vip             GetUserSpaceDetailVip            `json:"vip"`              // 会员信息
	Pendant         GetUserSpaceDetailPendant        `json:"pendant"`          // 头像框信息
	Nameplate       GetUserSpaceDetailNameplate      `json:"nameplate"`        // 勋章信息
	UserHonourInfo  GetUserSpaceDetailUserHonourInfo `json:"user_honour_info"` // （？）
	IsFollowed      bool                             `json:"is_followed"`      // 是否关注此用户。true：已关注。false：未关注。需要登录（Cookie） 。未登录恒为false
	TopPhoto        string                           `json:"top_photo"`        // 主页头图链接
	Theme           any                              `json:"theme"`            // （？）
	SysNotice       GetUserSpaceDetailSysNotice      `json:"sys_notice"`       // 系统通知。无内容则为空对象。主要用于展示如用户争议、纪念账号等等的小黄条
	LiveRoom        GetUserSpaceDetailLiveRoom       `json:"live_room"`        // 直播间信息
	Birthday        string                           `json:"birthday"`         // 生日。MM-DD。如设置隐私为空
	School          GetUserSpaceDetailSchool         `json:"school"`           // 学校
	Profession      GetUserSpaceDetailProfession     `json:"profession"`       // 专业资质信息
	Tags            any                              `json:"tags"`             // 个人标签
	Series          GetUserSpaceDetailSeries         `json:"series"`
	IsSeniorMember  int                              `json:"is_senior_member"` // 是否为硬核会员。0：否。1：是
	McnInfo         any                              `json:"mcn_info"`         // （？）
	GaiaResType     int                              `json:"gaia_res_type"`    // （？）
	GaiaData        any                              `json:"gaia_data"`        // （？）
	IsRisk          bool                             `json:"is_risk"`          // （？）
	Elec            GetUserSpaceDetailElec           `json:"elec"`             // 充电信息
	Contract        GetUserSpaceDetailContract       `json:"contract"`         // 是否显示老粉计划
	CertificateShow bool                             `json:"certificate_show"` // （？）
	NameRender      GetUserSpaceDetailNameRender     `json:"name_render"`      // 昵称渲染信息
}

// GetUserSpaceDetailMedal 是粉丝勋章信息。
type GetUserSpaceDetailMedal struct {
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

// GetUserSpaceDetailFansMedal 是粉丝勋章信息的外层对象。
type GetUserSpaceDetailFansMedal struct {
	Show  bool                    `json:"show"`
	Wear  bool                    `json:"wear"`  // 是否佩戴了粉丝勋章
	Medal GetUserSpaceDetailMedal `json:"medal"` // 粉丝勋章信息
}

// GetUserSpaceDetailOfficial 是用户空间中的认证信息。
type GetUserSpaceDetailOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetUserSpaceDetailVip 是用户空间中的大会员信息。
type GetUserSpaceDetailVip struct {
	Type               int                           `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                           `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                           `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                           `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int                           `json:"theme_type"`           // 0。作用尚不明确
	Label              GetUserSpaceDetailLabel       `json:"label"`                // 会员标签
	AvatarSubscript    int                           `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                        `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int                           `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptURL string                        `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                           `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                           `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetUserSpaceDetailVipOttInfo  `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetUserSpaceDetailVipSuperVip `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                           `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetUserSpaceDetailAvatarIcon  `json:"avatar_icon"`          // 大会员角标信息
}

// GetUserSpaceDetailLabel 是会员标签。
type GetUserSpaceDetailLabel struct {
	Path                  string                      `json:"path"`                      // 空。作用尚不明确
	Text                  string                      `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                      `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                      `json:"text_color"`                // 会员标签
	BgStyle               int                         `json:"bg_style"`                  // 1
	BgColor               string                      `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                      `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                        `json:"use_img_label"`             // true
	ImgLabelURIHans       string                      `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                      `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                      `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                      `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                         `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                         `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetUserSpaceDetailLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                      `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                      `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetUserSpaceDetailLabelGoto 是会员标签中的 label_goto 对象。
type GetUserSpaceDetailLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetUserSpaceDetailVipOttInfo 是大会员信息中的 ott_info 对象。
type GetUserSpaceDetailVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetUserSpaceDetailVipSuperVip 是大会员信息中的 super_vip 对象。
type GetUserSpaceDetailVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetUserSpaceDetailAvatarIcon 是大会员角标信息。
type GetUserSpaceDetailAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetUserSpaceDetailPendant 是头像框信息。
type GetUserSpaceDetailPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetUserSpaceDetailNameplate 是勋章信息。
type GetUserSpaceDetailNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetUserSpaceDetailUserHonourInfo 是用户荣誉信息。
type GetUserSpaceDetailUserHonourInfo struct {
	Mid    int    `json:"mid"`    // 0
	Colour string `json:"colour"` // null
	Tags   any    `json:"tags"`   // null
}

// GetUserSpaceDetailSysNotice 是系统通知。
type GetUserSpaceDetailSysNotice struct {
	ID         int    `json:"id"`          // id
	Content    string `json:"content"`     // 显示文案
	URL        string `json:"url"`         // 跳转地址
	NoticeType int    `json:"notice_type"` // 提示类型。1,2
	Icon       string `json:"icon"`        // 前缀图标
	TextColor  string `json:"text_color"`  // 文字颜色
	BgColor    string `json:"bg_color"`    // 背景颜色
}

// GetUserSpaceDetailLiveRoom 是直播间信息。
type GetUserSpaceDetailLiveRoom struct {
	RoomStatus    int                           `json:"roomStatus"` // 直播间状态。0：无房间。1：有房间
	LiveStatus    int                           `json:"liveStatus"` // 直播状态。0：未开播。1：直播中
	URL           string                        `json:"url"`        // 直播间网页 url
	Title         string                        `json:"title"`      // 直播间标题
	Cover         string                        `json:"cover"`      // 直播间封面 url
	WatchedShow   GetUserSpaceDetailWatchedShow `json:"watched_show"`
	RoomID        int                           `json:"roomid"`         // 直播间 id(短号)
	RoundStatus   int                           `json:"roundStatus"`    // 轮播状态。0：未轮播。1：轮播
	BroadcastType int                           `json:"broadcast_type"` // 0
}

// GetUserSpaceDetailWatchedShow 是直播间观看信息。
type GetUserSpaceDetailWatchedShow struct {
	Switch       bool   `json:"switch"` // ?
	Num          int    `json:"num"`    // total watched users
	TextSmall    string `json:"text_small"`
	TextLarge    string `json:"text_large"`
	Icon         string `json:"icon"`          // watched icon url
	IconLocation string `json:"icon_location"` // ?
	IconWeb      string `json:"icon_web"`      // watched icon url
}

// GetUserSpaceDetailSchool 是学校信息。
type GetUserSpaceDetailSchool struct {
	Name string `json:"name"` // 就读大学名称。没有则为空
}

// GetUserSpaceDetailProfession 是专业资质信息。
type GetUserSpaceDetailProfession struct {
	Name          string `json:"name"`           // 资质名称
	Department    string `json:"department"`     // 职位
	Title         string `json:"title"`          // 所属机构
	IsShow        int    `json:"is_show"`        // 是否显示。0：不显示。1：显示
	CategoryOne   string `json:"category_one"`   // 资质分类。实测为空
	CertificateNo string `json:"certificate_no"` // 资质证书编号。实测为空
}

// GetUserSpaceDetailSeries 是用户成长系列信息。
type GetUserSpaceDetailSeries struct {
	UserUpgradeStatus int  `json:"user_upgrade_status"` // (?)
	ShowUpgradeWindow bool `json:"show_upgrade_window"` // (?)
}

// GetUserSpaceDetailElec 是充电信息。
type GetUserSpaceDetailElec struct {
	ShowInfo GetUserSpaceDetailElecShowInfo `json:"show_info"`
}

// GetUserSpaceDetailElecShowInfo 是充电信息中的 show_info 对象。
type GetUserSpaceDetailElecShowInfo struct {
	Show    bool   `json:"show"`     // 是否开通了充电
	State   int    `json:"state"`    // 状态。-1：未开通。1：已开通
	Title   string `json:"title"`    // 空串
	Icon    string `json:"icon"`     // 空串
	JumpURL string `json:"jump_url"` // 空串
}

// GetUserSpaceDetailContract 是是否显示老粉计划。
type GetUserSpaceDetailContract struct {
	IsDisplay       bool `json:"is_display"`        // true/false。在页面中未使用此字段
	IsFollowDisplay bool `json:"is_follow_display"` // 是否在显示老粉计划。true：显示。false：不显示
}

// GetUserSpaceDetailNameRender 是昵称渲染信息。
type GetUserSpaceDetailNameRender struct {
	ColorsInfo   GetUserSpaceDetailNameRenderColors `json:"colors_info"`   // 昵称颜色信息（？）。作用尚不明确
	RenderScheme string                             `json:"render_scheme"` // "Default" 或 "Colorful"。（？）。作用尚不明确
}

// GetUserSpaceDetailNameRenderColors 是昵称颜色信息。
type GetUserSpaceDetailNameRenderColors struct {
	Color    []GetUserSpaceDetailNameRenderColor `json:"color"`     // 昵称颜色（？）。作用尚不明确
	ColorIDs []string                            `json:"color_ids"` // 昵称颜色 id（？）。作用尚不明确。实测为 ["6"]
}

// GetUserSpaceDetailNameRenderColor 是昵称颜色。
type GetUserSpaceDetailNameRenderColor struct {
	ColorDay   string `json:"color_day"`   // 浅色模式昵称颜色。HEX颜色代码
	ColorNight string `json:"color_night"` // 深色模式昵称颜色。HEX颜色代码
}

// GetUserCardResult 是获取用户用户名片的 data 对象。
type GetUserCardResult struct {
	Card         GetUserCardInfo `json:"card"`          // 卡片信息
	Following    bool            `json:"following"`     // 是否关注此用户。true：已关注。false：未关注。需要登录(Cookie)。未登录为false
	ArchiveCount int             `json:"archive_count"` // 用户稿件数
	ArticleCount int             `json:"article_count"` // 0。**作用尚不明确**
	Follower     int             `json:"follower"`      // 粉丝数
	LikeNum      int             `json:"like_num"`      // 点赞数
}

// GetUserCardInfo 是用户名片中的卡片信息。
type GetUserCardInfo struct {
	Mid            string                    `json:"mid"`             // 用户mid
	Approve        bool                      `json:"approve"`         // false。**作用尚不明确**
	Name           string                    `json:"name"`            // 用户昵称
	Sex            string                    `json:"sex"`             // 用户性别。男 女 保密
	Face           string                    `json:"face"`            // 用户头像链接
	DisplayRank    string                    `json:"DisplayRank"`     // 0。**作用尚不明确**
	Regtime        int                       `json:"regtime"`         // 0。**作用尚不明确**
	Spacesta       int                       `json:"spacesta"`        // 用户状态。0：正常。-2：被封禁
	Birthday       string                    `json:"birthday"`        // 空。**作用尚不明确**
	Place          string                    `json:"place"`           // 空。**作用尚不明确**
	Description    string                    `json:"description"`     // 空。**作用尚不明确**
	Article        int                       `json:"article"`         // 0。**作用尚不明确**
	Attentions     any                       `json:"attentions"`      // 空。**作用尚不明确**
	Fans           int                       `json:"fans"`            // 粉丝数
	Friend         int                       `json:"friend"`          // 关注数
	Attention      int                       `json:"attention"`       // 关注数
	Sign           string                    `json:"sign"`            // 签名
	LevelInfo      GetUserCardLevelInfo      `json:"level_info"`      // 等级
	Pendant        GetUserCardPendant        `json:"pendant"`         // 挂件
	Nameplate      GetUserCardNameplate      `json:"nameplate"`       // 勋章
	Official       GetUserCardOfficial       `json:"Official"`        // 认证信息
	OfficialVerify GetUserCardOfficialVerify `json:"official_verify"` // 认证信息2
	Vip            GetUserCardVip            `json:"vip"`             // 大会员状态
	Space          GetUserCardSpace          `json:"space"`           // 主页头图
}

// GetUserCardLevelInfo 是用户名片中的等级信息。
type GetUserCardLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 用户等级
	CurrentMin   int `json:"current_min"`   // 0
	CurrentExp   int `json:"current_exp"`   // 0
	NextExp      int `json:"next_exp"`      // 0
}

// GetUserCardPendant 是用户名片中的挂件信息。
type GetUserCardPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetUserCardNameplate 是用户名片中的勋章信息。
type GetUserCardNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetUserCardOfficial 是用户名片中的认证信息。
type GetUserCardOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetUserCardOfficialVerify 是用户名片中的认证信息2。
type GetUserCardOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetUserCardVip 是用户名片中的大会员状态。
type GetUserCardVip struct {
	VipType       int    `json:"vipType"`       // 大会员类型。0：无。1：月度大会员。2：年度及以上大会员
	DueRemark     string `json:"dueRemark"`     // 空。**作用尚不明确**
	AccessStatus  int    `json:"accessStatus"`  // 0。**作用尚不明确**
	VipStatus     int    `json:"vipStatus"`     // 大会员状态。0：无。1：有
	VipStatusWarn string `json:"vipStatusWarn"` // 空。**作用尚不明确**
	ThemeType     int    `json:"theme_type"`    // 0。**作用尚不明确**
}

// GetUserCardSpace 是用户名片中的主页头图。
type GetUserCardSpace struct {
	SImg string `json:"s_img"` // 主页头图url 小图
	LImg string `json:"l_img"` // 主页头图url 正常
}

// GetMyUserSpaceDetailResult 是获取登录用户空间详细信息的 data 对象。
type GetMyUserSpaceDetailResult struct {
	Mid            int                             `json:"mid"`             // mid
	Name           string                          `json:"name"`            // 昵称
	Sex            string                          `json:"sex"`             // 性别。男 女 保密
	Face           string                          `json:"face"`            // 头像图片url
	Sign           string                          `json:"sign"`            // 签名
	Rank           int                             `json:"rank"`            // 10000。**作用尚不明确**
	Level          int                             `json:"level"`           // 当前等级。0-6级
	Jointime       int                             `json:"jointime"`        // 0。**作用尚不明确**
	Moral          int                             `json:"moral"`           // 节操。默认70
	Silence        int                             `json:"silence"`         // 封禁状态。0：正常。1：被封
	EmailStatus    int                             `json:"email_status"`    // 已验证邮箱。0：未验证。1：已验证
	TelStatus      int                             `json:"tel_status"`      // 已验证手机号。0：未验证。1：已验证
	Identification int                             `json:"identification"`  // 1。**作用尚不明确**
	Vip            GetMyUserSpaceDetailVip         `json:"vip"`             // 大会员状态
	Pendant        GetMyUserSpaceDetailPendant     `json:"pendant"`         // 头像框信息
	Nameplate      GetMyUserSpaceDetailNameplate   `json:"nameplate"`       // 勋章信息
	Official       GetMyUserSpaceDetailOfficial    `json:"official"`        // 认证信息
	Birthday       int                             `json:"birthday"`        // 生日。时间戳
	IsTourist      int                             `json:"is_tourist"`      // 0。**作用尚不明确**
	IsFakeAccount  int                             `json:"is_fake_account"` // 0。**作用尚不明确**
	PinPrompting   int                             `json:"pin_prompting"`   // 0。**作用尚不明确**
	IsDeleted      int                             `json:"is_deleted"`      // 0。**作用尚不明确**
	InRegAudit     int                             `json:"in_reg_audit"`
	IsRipUser      bool                            `json:"is_rip_user"`
	Profession     GetMyUserSpaceDetailProfession  `json:"profession"`   // 专业资质
	Coins          float64                         `json:"coins"`        // 硬币数
	Following      int                             `json:"following"`    // 粉丝数
	Follower       int                             `json:"follower"`     // 粉丝数
	Honours        GetMyUserSpaceDetailHonours     `json:"honours"`      // 荣誉信息
	DigitalId      string                          `json:"digital_id"`   // 数字id。实测为空
	DigitalType    int                             `json:"digital_type"` // 数字id类型。实测为 -2
	Attestation    GetMyUserSpaceDetailAttestation `json:"attestation"`  // 认证信息
	LevelExp       GetMyUserSpaceDetailLevelExp    `json:"level_exp"`    // 等级经验信息
}

// GetMyUserSpaceDetailVip 是登录用户空间中的大会员状态。
type GetMyUserSpaceDetailVip struct {
	Type            int                       `json:"type"`             // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status          int                       `json:"status"`           // 会员状态。0：无。1：有
	DueDate         int                       `json:"due_date"`         // 会员过期时间。Unix时间戳(毫秒)
	ThemeType       int                       `json:"theme_type"`       // 0。作用尚不明确
	Label           GetMyUserSpaceDetailLabel `json:"label"`            // 会员标签
	AvatarSubscript int                       `json:"avatar_subscript"` // 是否显示会员图标。0：不显示。1：显示
	NicknameColor   string                    `json:"nickname_color"`   // 会员昵称颜色。颜色码
}

// GetMyUserSpaceDetailLabel 是登录用户空间中的会员标签。
type GetMyUserSpaceDetailLabel struct {
	Path                  string                        `json:"path"`                      // 空。作用尚不明确
	Text                  string                        `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                        `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                        `json:"text_color"`                // 会员标签
	BgStyle               int                           `json:"bg_style"`                  // 1
	BgColor               string                        `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                        `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                          `json:"use_img_label"`             // true
	ImgLabelURIHans       string                        `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                        `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                        `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                        `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                           `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                           `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetMyUserSpaceDetailLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                        `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                        `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetMyUserSpaceDetailLabelGoto 是会员标签中的 label_goto 对象。
type GetMyUserSpaceDetailLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetMyUserSpaceDetailPendant 是登录用户空间中的头像框信息。
type GetMyUserSpaceDetailPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetMyUserSpaceDetailNameplate 是登录用户空间中的勋章信息。
type GetMyUserSpaceDetailNameplate struct {
	Nid        int    `json:"nid"`         // 勋章id
	Name       string `json:"name"`        // 勋章名称
	Image      string `json:"image"`       // 勋章图标
	ImageSmall string `json:"image_small"` // 勋章图标（小）
	Level      string `json:"level"`       // 勋章等级
	Condition  string `json:"condition"`   // 获取条件
}

// GetMyUserSpaceDetailOfficial 是登录用户空间中的认证信息。
type GetMyUserSpaceDetailOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetMyUserSpaceDetailProfession 是专业资质。
type GetMyUserSpaceDetailProfession struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ShowName string `json:"show_name"`
}

// GetMyUserSpaceDetailHonours 是荣誉信息。
type GetMyUserSpaceDetailHonours struct {
	Mid               int                              `json:"mid"`                 // 用户 mid
	Colour            GetMyUserSpaceDetailHonourColour `json:"colour"`              // 荣誉颜色
	Tags              []any                            `json:"tags"`                // 荣誉标签。实测为 null
	IsLatest100Honour int                              `json:"is_latest_100honour"` // 是否为最近 100 个荣誉？。0：否。1：是。实测为 0
}

// GetMyUserSpaceDetailHonourColour 是荣誉颜色。
type GetMyUserSpaceDetailHonourColour struct {
	Dark   string `json:"dark"`   // 深色模式下的颜色。实测为 #CE8620
	Normal string `json:"normal"` // 浅色模式下的颜色。实测为 #F0900B
}

// GetMyUserSpaceDetailAttestation 是认证信息。
type GetMyUserSpaceDetailAttestation struct {
	Type       int                                       `json:"type"`        // 认证类型。见 docs/user/official_role.md
	CommonInfo GetMyUserSpaceDetailAttestationCommonInfo `json:"common_info"` // 认证信息
	SpliceInfo GetMyUserSpaceDetailAttestationSpliceInfo `json:"splice_info"` // 拼接认证信息
	Icon       string                                    `json:"icon"`        // 认证图标url。实测为空
	Desc       string                                    `json:"desc"`        // 认证备注。实测为空
}

// GetMyUserSpaceDetailAttestationCommonInfo 是认证信息中的 common_info 对象。
type GetMyUserSpaceDetailAttestationCommonInfo struct {
	Title       string `json:"title"`        // 认证信息。实测为空
	Prefix      string `json:"prefix"`       // 认证前缀。如 bilibili UP主认证
	PrefixTitle string `json:"prefix_title"` // 带前缀的认证信息
}

// GetMyUserSpaceDetailAttestationSpliceInfo 是认证信息中的 splice_info 对象。
type GetMyUserSpaceDetailAttestationSpliceInfo struct {
	Title string `json:"title"` // 认证信息。实测为空
}

// GetMyUserSpaceDetailLevelExp 是等级经验信息。
type GetMyUserSpaceDetailLevelExp struct {
	CurrentLevel int `json:"current_level"` // 当前等级。0-6级
	CurrentMin   int `json:"current_min"`   // 指当前等级从多少经验值开始
	CurrentExp   int `json:"current_exp"`   // 当前账户的经验值
	NextExp      int `json:"next_exp"`      // 下一个等级所需的经验值（不是还需要多少）
	LevelUp      int `json:"level_up"`      // 升级时间。Unix秒级时间戳
}

// JoinOldFansResult 是加入老粉计划的 data 对象。
type JoinOldFansResult struct {
	AllowMessage bool   `json:"allow_message"` // true
	InputText    string `json:"input_text"`    // UP主加油！看好你噢
	InputTitle   string `json:"input_title"`   // 感谢你对UP主的特别支持，“老粉”可期！私信留言鼓励下TA吧
}

// FansSendMessageResult 是老粉计划发送留言的 data 对象。
type FansSendMessageResult struct {
	SuccessToast string `json:"success_toast"` // "提交成功，UP主已收到留言~"
}

// BatchGetUserCardsResult 是获取多用户详细信息 data 数组中的用户对象。
type BatchGetUserCardsResult struct {
	Mid     int    `json:"mid"`     // mid
	Name    string `json:"name"`    // 昵称
	Face    string `json:"face"`    // 头像链接
	Sign    string `json:"sign"`    // 签名
	Rank    int    `json:"rank"`    // 用户权限等级
	Level   int    `json:"level"`   // 当前等级。0-6 级
	Silence int    `json:"silence"` // 封禁状态。0：正常。1：被封
}
