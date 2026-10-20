package dynamic

// GetDynamicPortal 的响应模型。

// GetDynamicPortalResult 是最近更新 UP 主列表的 data 对象。
type GetDynamicPortalResult struct {
	MyInfo GetDynamicPortalMyInfo `json:"my_info"` // 个人关注的一些信息
	UpList []GetDynamicPortalUp   `json:"up_list"` // 最近更新的up主列表
}

// GetDynamicPortalMyInfo 是动态门户中的个人关注信息。
type GetDynamicPortalMyInfo struct {
	Dyns      int                             `json:"dyns"`      // 个人动态
	Face      string                          `json:"face"`      // 头像url
	FaceNft   int                             `json:"face_nft"`  // 含义尚不明确
	Follower  int                             `json:"follower"`  // 粉丝数量
	Following int                             `json:"following"` // 我的关注
	LevelInfo GetDynamicPortalMyInfoLevelInfo `json:"level_info"`
	Mid       int                             `json:"mid"`      // 账户mid
	Name      string                          `json:"name"`     // 账户名称
	Official  GetDynamicPortalMyInfoOfficial  `json:"official"` // 认证信息
	SpaceBg   string                          `json:"space_bg"` // 账户个人中心的背景横幅url
	Vip       GetDynamicPortalMyInfoVip       `json:"vip"`      // vip信息
}

// GetDynamicPortalMyInfoLevelInfo 是本人等级信息。
type GetDynamicPortalMyInfoLevelInfo struct {
	CurrentExp   int   `json:"current_exp"`
	CurrentLevel int   `json:"current_level"` // 当前等级，0-6级
	CurrentMin   int   `json:"current_min"`
	LevelUp      int64 `json:"level_up"`
	NextExp      int   `json:"next_exp"`
}

// GetDynamicPortalMyInfoOfficial 是本人认证信息。
type GetDynamicPortalMyInfoOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetDynamicPortalMyInfoVip 是本人会员信息。
type GetDynamicPortalMyInfoVip struct {
	Type               int                                 `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                                 `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                                 `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                                 `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int                                 `json:"theme_type"`           // 0。作用尚不明确
	Label              GetDynamicPortalMyInfoVipLabel      `json:"label"`                // 会员标签
	AvatarSubscript    int                                 `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                              `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int                                 `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptURL string                              `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                                 `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                                 `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetDynamicPortalMyInfoVipOttInfo    `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetDynamicPortalMyInfoVipSuperVip   `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                                 `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetDynamicPortalMyInfoVipAvatarIcon `json:"avatar_icon"`          // 大会员角标信息
}

// GetDynamicPortalMyInfoVipLabel 是本人会员标签。
type GetDynamicPortalMyInfoVipLabel struct {
	Path                  string                             `json:"path"`                      // 空。作用尚不明确
	Text                  string                             `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                             `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                             `json:"text_color"`                // 会员标签
	BgStyle               int                                `json:"bg_style"`                  // 1
	BgColor               string                             `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                             `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                               `json:"use_img_label"`             // true
	ImgLabelURIHans       string                             `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                             `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                             `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                             `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                                `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                                `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetDynamicPortalMyInfoVipLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                             `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                             `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetDynamicPortalMyInfoVipLabelGoto 是会员标签中的跳转信息。
type GetDynamicPortalMyInfoVipLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetDynamicPortalMyInfoVipOttInfo 是会员信息中的 OTT 信息。
type GetDynamicPortalMyInfoVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetDynamicPortalMyInfoVipSuperVip 是会员信息中的超级会员信息。
type GetDynamicPortalMyInfoVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetDynamicPortalMyInfoVipAvatarIcon 是大会员角标信息。
type GetDynamicPortalMyInfoVipAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetDynamicPortalUp 是一位最近更新的 UP 主。
type GetDynamicPortalUp struct {
	Face            string `json:"face"`       // UP主头像
	HasUpdate       bool   `json:"has_update"` // 最近是否有更新
	IsReserveRecall bool   `json:"is_reserve_recall"`
	Mid             int    `json:"mid"`   // UP主mid
	Uname           string `json:"uname"` // UP主昵称
}
