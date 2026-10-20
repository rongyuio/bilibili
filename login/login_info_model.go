package login

// 登录基本信息相关响应模型。各接口的返回体独立定义，互不复用。

// GetNavInfoLevelInfo 是导航栏用户信息中的等级信息。
//
// next_exp 在 6 级时为字符串 "--"，因此用 any 承接 num 与 str 两种形态。
type GetNavInfoLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 当前等级
	CurrentMin   int `json:"current_min"`   // 当前等级经验最低值
	CurrentExp   int `json:"current_exp"`   // 当前经验
	NextExp      any `json:"next_exp"`      // 升级下一等级需达到的经验。小于 6 级时为 num，6 级时为 str "--"
}

// GetNavInfoOfficial 是导航栏用户信息中的认证信息。
type GetNavInfoOfficial struct {
	Role  int    `json:"role"`  // 认证类型
	Title string `json:"title"` // 认证信息。无为空
	Desc  string `json:"desc"`  // 认证备注。无为空
	Type  int    `json:"type"`  // 是否认证。-1：无。0：认证
}

// GetNavInfoOfficialVerify 是导航栏用户信息中的认证信息 2。
type GetNavInfoOfficialVerify struct {
	Type int    `json:"type"` // 是否认证。-1：无。0：认证
	Desc string `json:"desc"` // 认证信息。无为空
}

// GetNavInfoPendant 是导航栏用户信息中的头像框信息。
type GetNavInfoPendant struct {
	Pid               int    `json:"pid"`                 // 挂件id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片url
	JumpURL           string `json:"jump_url"`            // 挂件跳转url
	Type              string `json:"type"`                // 装扮类型。suit：一般装扮。vip_suit：vip 装扮
	Expire            int    `json:"expire"`              // 固定值0，作用尚不明确
	ImageEnhance      string `json:"image_enhance"`       // 头像框图片url
	ImageEnhanceFrame string `json:"image_enhance_frame"` // (?)
}

// GetNavInfoLabelGoto 是会员标签中的跳转信息。
type GetNavInfoLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetNavInfoLabel 是导航栏用户信息中的会员标签。
type GetNavInfoLabel struct {
	Path                  string              `json:"path"`                      // 空。作用尚不明确
	Text                  string              `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string              `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string              `json:"text_color"`                // 会员标签
	BgStyle               int                 `json:"bg_style"`                  // 1
	BgColor               string              `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string              `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                `json:"use_img_label"`             // true
	ImgLabelURIHans       string              `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string              `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string              `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string              `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                 `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                 `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetNavInfoLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string              `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string              `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetNavInfoVipOttInfo 是导航栏用户信息 vip 中的 ott_info 对象。
type GetNavInfoVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetNavInfoVipSuperVip 是导航栏用户信息 vip 中的 super_vip 对象。
type GetNavInfoVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetNavInfoAvatarIcon 是导航栏用户信息 vip 中的大会员角标信息。
type GetNavInfoAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetNavInfoVip 是导航栏用户信息中的会员信息。
type GetNavInfoVip struct {
	Type               int                   `json:"type"`                 // （？）。作用尚不明确
	Status             int                   `json:"status"`               // （？）。作用尚不明确
	DueDate            int                   `json:"due_date"`             // （？）。作用尚不明确，毫秒时间戳
	VipPayType         int                   `json:"vip_pay_type"`         // （？）。作用尚不明确
	ThemeType          int                   `json:"theme_type"`           // （？）。作用尚不明确
	Label              GetNavInfoLabel       `json:"label"`                // 会员标签。结构与 vip_label 相同
	AvatarSubscript    int                   `json:"avatar_subscript"`     // （？）。作用尚不明确
	NicknameColor      string                `json:"nickname_color"`       // （？）。作用尚不明确
	Role               int                   `json:"role"`                 // （？）。作用尚不明确
	AvatarSubscriptURL string                `json:"avatar_subscript_url"` // （？）。作用尚不明确
	TvVipStatus        int                   `json:"tv_vip_status"`        // （？）。作用尚不明确
	TvVipPayType       int                   `json:"tv_vip_pay_type"`      // （？）。作用尚不明确
	OttInfo            GetNavInfoVipOttInfo  `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetNavInfoVipSuperVip `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                   `json:"tv_due_date"`          // （？）。作用尚不明确，秒级时间戳
	AvatarIcon         GetNavInfoAvatarIcon  `json:"avatar_icon"`          // （？）。作用尚不明确，含 icon_resource 子对象
}

// GetNavInfoWallet 是导航栏用户信息中的 B币钱包信息。
type GetNavInfoWallet struct {
	Mid           int64 `json:"mid"`             // 登录用户 mid
	BcoinBalance  int   `json:"bcoin_balance"`   // 拥有 B币数
	CouponBalance int   `json:"coupon_balance"`  // 每月奖励 B币数
	CouponDueTime int64 `json:"coupon_due_time"` // (?)。作用尚不明确
}

// GetNavInfoWbiImg 是导航栏用户信息中的 Wbi 签名实时口令。
type GetNavInfoWbiImg struct {
	ImgURL string `json:"img_url"` // Wbi 签名参数 imgKey 的伪装 url
	SubURL string `json:"sub_url"` // Wbi 签名参数 subKey 的伪装 url
}

// GetNavInfoResult 是导航栏用户信息接口的 data 对象。
type GetNavInfoResult struct {
	IsLogin            bool                     `json:"isLogin"`              // 是否已登录。false：未登录。true：已登录
	EmailVerified      int                      `json:"email_verified"`       // 是否验证邮箱地址。0：未验证。1：已验证
	Face               string                   `json:"face"`                 // 用户头像 url
	LevelInfo          GetNavInfoLevelInfo      `json:"level_info"`           // 等级信息
	Mid                int64                    `json:"mid"`                  // 用户 mid
	MobileVerified     int                      `json:"mobile_verified"`      // 是否验证手机号。0：未验证。1：已验证
	Money              float64                  `json:"money"`                // 拥有硬币数
	Moral              int                      `json:"moral"`                // 当前节操值。上限为 70
	Official           GetNavInfoOfficial       `json:"official"`             // 认证信息
	OfficialVerify     GetNavInfoOfficialVerify `json:"officialVerify"`       // 认证信息 2
	Pendant            GetNavInfoPendant        `json:"pendant"`              // 头像框信息
	Scores             int                      `json:"scores"`               // (?)。作用尚不明确
	Uname              string                   `json:"uname"`                // 用户昵称
	VipDueDate         int64                    `json:"vipDueDate"`           // 会员到期时间。毫秒 时间戳
	VipStatus          int                      `json:"vipStatus"`            // 会员开通状态。0：无。1：有
	VipType            int                      `json:"vipType"`              // 会员类型。0：无。1：月度大会员。2：年度及以上大会员
	VipPayType         int                      `json:"vip_pay_type"`         // 会员开通状态。0：无。1：有
	VipThemeType       int                      `json:"vip_theme_type"`       // (?)。作用尚不明确
	VipLabel           GetNavInfoLabel          `json:"vip_label"`            // 会员标签
	VipAvatarSubscript int                      `json:"vip_avatar_subscript"` // 是否显示会员图标。0：不显示。1：显示
	VipNicknameColor   string                   `json:"vip_nickname_color"`   // 会员昵称颜色。颜色码
	Vip                GetNavInfoVip            `json:"vip"`                  // 会员信息
	LegalRegion        string                   `json:"legal_region"`         // (?)。作用尚不明确
	IpRegion           string                   `json:"ip_region"`            // (?)。作用尚不明确
	Wallet             GetNavInfoWallet         `json:"wallet"`               // B币钱包信息
	HasShop            bool                     `json:"has_shop"`             // 是否拥有推广商品。false：无。true：有
	ShopURL            string                   `json:"shop_url"`             // 商品推广页面 url
	AllowanceCount     int                      `json:"allowance_count"`      // (?)。作用尚不明确
	AnswerStatus       int                      `json:"answer_status"`        // (?)。作用尚不明确
	IsSeniorMember     int                      `json:"is_senior_member"`     // 是否硬核会员。0：非硬核会员。1：硬核会员
	WbiImg             GetNavInfoWbiImg         `json:"wbi_img"`              // Wbi 签名实时口令。该字段即使用户未登录也存在
	IsJury             bool                     `json:"is_jury"`              // 是否风纪委员。true：风纪委员。false：非风纪委员
}

// GetHomeUserInfoLevelInfo 是已弃用的登录用户信息接口中的等级信息。
type GetHomeUserInfoLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 当前等级
	CurrentMin   int `json:"current_min"`   // 当前等级经验最低值
	CurrentExp   int `json:"current_exp"`   // 当前经验
	NextExp      int `json:"next_exp"`      // 升级下一等级需达到的经验
}

// GetHomeUserInfoResult 是已弃用的登录用户信息接口的 data 对象。
type GetHomeUserInfoResult struct {
	LevelInfo        GetHomeUserInfoLevelInfo `json:"level_info"`        // 等级信息
	BCoins           int                      `json:"bCoins"`            // 拥有 B币数
	Coins            float64                  `json:"coins"`             // 拥有硬币数
	Face             string                   `json:"face"`              // 登录用户头像 url
	NameplateCurrent any                      `json:"nameplate_current"` // 登录用户勋章 url。可能为 null
	PendantCurrent   string                   `json:"pendant_current"`   // 登录用户头像框 url
	Uname            string                   `json:"uname"`             // 登录用户昵称
	UserStatus       string                   `json:"userStatus"`        // 登录用户状态
	VipType          int                      `json:"vipType"`           // 大会员类型。0：无。1：月度。2：年度
	VipStatus        int                      `json:"vipStatus"`         // 会员开通状态。0：无。1：有
	OfficialVerify   int                      `json:"official_verify"`   // 是否认证。-1：无。0：认证
	PointBalance     int                      `json:"pointBalance"`      // 0。作用尚不明确
}

// GetAppAccountMyInfoVipLabel 是 APP 端登录用户信息中的会员标签。
type GetAppAccountMyInfoVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetAppAccountMyInfoVip 是 APP 端登录用户信息中的会员信息。
type GetAppAccountMyInfoVip struct {
	Type            int                         `json:"type"`             // 大会员类型。0：无。1：月度。2：年度
	Status          int                         `json:"status"`           // 会员开通状态。0：无。1：有
	DueDate         int64                       `json:"due_date"`         // 大会员到期时间。毫秒 时间戳
	VipPayType      int                         `json:"vip_pay_type"`     // 会员开通状态。0：无。1：有
	ThemeType       int                         `json:"theme_type"`       // 会员开通状态。0：无。1：有
	Label           GetAppAccountMyInfoVipLabel `json:"label"`            // 大会员信息
	AvatarSubscript int                         `json:"avatar_subscript"` // 是否显示会员图标。0：不显示。1：显示
	NicknameColor   string                      `json:"nickname_color"`   // 会员昵称颜色。颜色码
}

// GetAppAccountMyInfoOfficial 是 APP 端登录用户信息中的认证信息。
type GetAppAccountMyInfoOfficial struct {
	Role  int    `json:"role"`  // 认证类型
	Title string `json:"title"` // 认证信息。无为空
	Desc  string `json:"desc"`  // 认证备注。无为空
	Type  int    `json:"type"`  // 认证备注。无为空
}

// GetAppAccountMyInfoInvite 是 APP 端登录用户信息中的邀请信息。
type GetAppAccountMyInfoInvite struct {
	InviteRemind int  `json:"invite_remind"` // 1。作用尚不明确
	Display      bool `json:"display"`       // true。作用尚不明确
}

// GetAppAccountMyInfoResult 是 APP 端登录用户信息接口的 data 对象。
type GetAppAccountMyInfoResult struct {
	Mid            int64                       `json:"mid"`            // 用户 mid
	Name           string                      `json:"name"`           // 用户昵称
	Sign           string                      `json:"sign"`           // 用户签名
	Coins          float64                     `json:"coins"`          // 拥有硬币数
	Birthday       string                      `json:"birthday"`       // 用户生日。YYYY-MM-DD
	Face           string                      `json:"face"`           // 用户头像 url
	Sex            int                         `json:"sex"`            // 用户性别。0：私密。1：男。2：女
	Level          int                         `json:"level"`          // 用户等级。0-6
	Rank           int                         `json:"rank"`           // 1000。作用尚不明确
	Silence        int                         `json:"silence"`        // 用户是否被封禁。0：正常。1：封禁
	Vip            GetAppAccountMyInfoVip      `json:"vip"`            // 会员信息
	EmailStatus    int                         `json:"email_status"`   // 是否验证邮箱地址。0：未验证。1：已验证
	TelStatus      int                         `json:"tel_status"`     // 是否验证手机号。0：未验证。1：已验证
	Official       GetAppAccountMyInfoOfficial `json:"official"`       // 认证信息
	Identification int                         `json:"identification"` // 1。作用尚不明确
	Invite         GetAppAccountMyInfoInvite   `json:"invite"`         // 邀请信息
	IsTourist      int                         `json:"is_tourist"`     // 0。作用尚不明确
	PinPrompting   int                         `json:"pin_prompting"`  // 0。作用尚不明确
}

// GetNavStatResult 是登录用户状态数接口的 data 对象。
type GetNavStatResult struct {
	Following    int `json:"following"`     // 关注数
	Follower     int `json:"follower"`      // 粉丝数
	DynamicCount int `json:"dynamic_count"` // 发布动态数
	FansEffect   any `json:"fans_effect"`   // (?)。实测为 null
}

// GetCoinResult 是获取硬币数接口的 data 对象。
type GetCoinResult struct {
	Money *float64 `json:"money"` // 当前硬币数。硬币为 0 时为 null
}
