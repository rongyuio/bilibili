package bilibili

// 大会员相关响应模型。

type VipPrivilegeInfo struct {
	Type            int         `json:"type"`              // 卡券类型。详见 list 数组表格中的 type 项
	State           int         `json:"state"`             // 兑换状态。0：未兑换。1：已兑换。2：未完成（若需要完成）
	ExpireTime      int         `json:"expire_time"`       // 本轮卡券过期时间戳。当月月底/当日24点
	VipType         int         `json:"vip_type"`          // 当前用户的大会员状态。2：年度大会员
	NextReceiveDays int         `json:"next_receive_days"` // 距下一轮兑换剩余天数。无权限时，每月任务固定为 0，每日固定为 1
	PeriodEndUnix   int         `json:"period_end_unix"`   // 下一轮兑换开始时间戳。秒级时间戳
	IsCount         bool        `json:"is_count"`          // 是否计入？。实测为 true 或 false
	CouponCode      string      `json:"coupon_code"`       // 卡券代码。实测为空
	AppDescribe     string      `json:"app_describe"`      // app 内描述。实测为空
	ReciveState     int         `json:"recive_state"`      // 领取状态。0：未领取。1：已领取。2：未完成
	SalaryType      int         `json:"salary_type"`       // 发放类型？。实测为 0 或 1
	ExpParams       ExpParams   `json:"exp_params"`        // 实验参数。非必要。无效时为 null
	ExtraParams     ExtraParams `json:"extra_params"`      // 额外参数。非必要。无效时为 null
}

// ExpParams `list[]` 中 `exp_params` 对象。实验参数，非必要（无效时为 null）
type ExpParams struct {
	ExpGroupTag string `json:"exp_group_tag"` // 实验分组标签。实测为空或 45476
	HitValue    int    `json:"hit_value"`     // 实验命中值。实测为 0 或 2
}

// ExtraParams `list[]` 中 `extra_params` 对象。额外参数，非必要（无效时为 null）
type ExtraParams struct {
	IsAlloweReceive         string `json:"is_allowe_receive"`           // 是否允许领取。"true" / "false"，注意是字符串
	IsShow                  string `json:"is_show"`                     // 是否展示。"true" / "false"，注意是字符串	LastSalaryTime           string `json:"last_salary_time"`            // 上次发放时间。字符串形式的秒级时间戳，实测为 "0"
	Now                     string `json:"now"`                         // 当前时间。字符串形式的秒级时间戳
	ComicShowCouponPropInfo string `json:"comic_show_coupon_prop_info"` // 漫展优惠券道具信息。JSON 字符串，含 prop_id / prop_name 等字段
}

type VipPrivilege struct {
	List            []VipPrivilegeInfo `json:"list"`              // 卡券信息列表
	IsShortVip      bool               `json:"is_short_vip"`      // (?)
	IsFreightOpen   bool               `json:"is_freight_open"`   // (?)
	Level           int                `json:"level"`             // 当前等级
	CurExp          int                `json:"cur_exp"`           // 当前拥有经验值
	NextExp         int                `json:"next_exp"`          // 升级所需经验值。满级时为 -1
	IsVip           bool               `json:"is_vip"`            // 是否为大会员
	IsSeniorMember  int                `json:"is_senior_member"`  // (?)
	Format060102    int                `json:"format060102"`      // (?)
	KeeptimeStart   int                `json:"keeptime_start"`    // 大会员当前阶段开始时间。秒级时间戳
	ComicShowCoupon ComicShowCoupon    `json:"comic_show_coupon"` // 漫展优惠券信息
}

// ComicShowCoupon `data` 中的 `comic_show_coupon` 对象
type ComicShowCoupon struct {
	Hide           bool         `json:"hide"`            // 是否隐藏。实测为 false
	State          int          `json:"state"`           // 状态。实测为 0
	Count          int          `json:"count"`           // 优惠券数量。实测为 2
	Title          string       `json:"title"`           // 标题。实测为 大会员专属票务优惠券
	SubTitle       string       `json:"sub_title"`       // 副标题。实测为 立减2元 | 满120减5
	Coupons        []CouponInfo `json:"coupons"`         // 优惠券列表
	PromotionRules []string     `json:"promotion_rules"` // 优惠规则列表。如 立减2元
	HideV2         bool         `json:"hide_v2"`         // 是否隐藏（v2）。实测为 false
	ExpireTip      string       `json:"expire_tip"`      // 过期提示。实测为空
	ExpireTime     int          `json:"expire_time"`     // 过期时间。实测为 0
}

// CouponInfo `comic_show_coupon.coupons[]` 中的对象
type CouponInfo struct {
	Amount           int    `json:"amount"`            // 优惠金额。单位为分
	PaymentThreshold string `json:"payment_threshold"` // 使用门槛。如 立减、满120可用
	Title            string `json:"title"`             // 标题。实测为 大会员专享票务优惠券
	SubTitle         string `json:"sub_title"`         // 副标题。实测为 领取后15天内可用
}

type VipUserAccount struct {
	Mid            int    `json:"mid"`              // 用户 mid
	Name           string `json:"name"`             // 昵称
	Sex            string `json:"sex"`              // 性别。男 / 女 / 保密
	Face           string `json:"face"`             // 头像 url
	Sign           string `json:"sign"`             // 签名
	Rank           int    `json:"rank"`             // 等级
	Birthday       int    `json:"birthday"`         // 生日。秒时间戳
	IsFakeAccount  int    `json:"is_fake_account"`  // (?)
	IsDeleted      int    `json:"is_deleted"`       // 是否注销。0：正常。1：注销
	InRegAudit     int    `json:"in_reg_audit"`     // 是否注册审核。0：正常。1：审核
	IsSeniorMember int    `json:"is_senior_member"` // 是否转正。0：未转正。1：正式会员
}

type VipUserVip struct {
	Mid                  int    `json:"mid"`                     // 用户 mid
	VipType              int    `json:"vip_type"`                // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	VipStatus            int    `json:"vip_status"`              // 会员状态。0：无。1：有
	VipDueDate           int    `json:"vip_due_date"`            // 会员过期时间。毫秒时间戳
	VipPayType           int    `json:"vip_pay_type"`            // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType            int    `json:"theme_type"`              // (?)
	Label                Label  `json:"label"`                   // 会员标签
	AvatarSubscript      int    `json:"avatar_subscript"`        // 是否显示会员图标。0：不显示。1：显示
	NicknameColor        string `json:"nickname_color"`          // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	IsNewUser            bool   `json:"is_new_user"`             // (?)
	TipMaterial          any    `json:"tip_material"`            // (?)
	VipIsValid           bool   `json:"vip_is_valid"`            // 大会员是否有效。实测为 false
	VipIsOverdue         bool   `json:"vip_is_overdue"`          // 大会员是否已过期。实测为 true
	VipExpireDays        int    `json:"vip_expire_days"`         // 大会员过期天数。实测为 -38，即已过期 38 天
	IsTvVip              bool   `json:"is_tv_vip"`               // 是否为电视大会员。实测为 false
	VipMembershipIsVip   bool   `json:"vip_membership_is_vip"`   // 是否为大会员？。实测为 false，作用尚不明确
	SuperMembershipIsVip bool   `json:"super_membership_is_vip"` // 是否为超级大会员？。实测为 false，作用尚不明确
}

type VipUserTv struct {
	Type       int `json:"type"`         // 电视大会员类型。0：无。1：月大会员。2：年度及以上大会员
	VipPayType int `json:"vip_pay_type"` // 电视大支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	Status     int `json:"status"`       // 电视大会员状态。0：无。1：有
	DueDate    int `json:"due_date"`     // 电视大会员过期时间。毫秒时间戳
}

type AvatarPendant struct {
	Image             string `json:"image"`               // 头像框 url
	ImageEnhance      string `json:"image_enhance"`       // 头像框 url。动态图
	ImageEnhanceFrame string `json:"image_enhance_frame"` // 动态头像框帧波普版 url
}

type VipUser struct {
	Account              VipUserAccount `json:"account"`                // 账号基本信息
	Vip                  VipUserVip     `json:"vip"`                    // 账号会员信息
	Tv                   VipUserTv      `json:"tv"`                     // 电视会员信息
	BackgroundImageSmall string         `json:"background_image_small"` // 空
	BackgroundImageBig   string         `json:"background_image_big"`   // 空
	PanelTitle           string         `json:"panel_title"`            // 用户昵称
	PanelSubTitle        string         `json:"panel_sub_title"`        // 面板副标题。如 大会员已过期
	AvatarPendant        AvatarPendant  `json:"avatar_pendant"`         // 用户头像框信息
	VipOverdueExplain    string         `json:"vip_overdue_explain"`    // 大会员提示文案。有效期 / 到期
	TvOverdueExplain     string         `json:"tv_overdue_explain"`     // 电视大会员提示文案。有效期 / 到期
	AccountExceptionText string         `json:"account_exception_text"` // 空
	IsAutoRenew          bool           `json:"is_auto_renew"`          // 是否自动续费。true：是。false：否
	IsTvAutoRenew        bool           `json:"is_tv_auto_renew"`       // 是否自动续费电视大会员。true：是。false：否
	SurplusSeconds       int            `json:"surplus_seconds"`        // 大会员到期剩余时间。单位为秒
	VipKeepTime          int            `json:"vip_keep_time"`          // 持续开通大会员时间。单位为秒
	Renew                any            `json:"renew"`                  // (?)
	Notice               any            `json:"notice"`                 // (?)
}

type VipWallet struct {
	Coupon            int  `json:"coupon"`             // 当前 B 币券
	Point             int  `json:"point"`              // (?)
	PrivilegeReceived bool `json:"privilege_received"` // (?)
}

type ChildPrivilege struct {
	FirstID            int    `json:"first_id"`             // 特权父类 id
	ReportID           string `json:"report_id"`            // 上报 id。该特权的代号？
	Name               string `json:"name"`                 // 特权名称
	Desc               string `json:"desc"`                 // 特权简介文案
	Explain            string `json:"explain"`              // 特权介绍正文
	IconURL            string `json:"icon_url"`             // 特权图标 url
	IconGrayURL        string `json:"icon_gray_url"`        // 特权图标灰色主题 url。某些项目无此字段
	BackgroundImageURL string `json:"background_image_url"` // 背景图片 url
	Link               string `json:"link"`                 // 特权介绍页 url
	ImageURL           string `json:"image_url"`            // 特权示例图 url
	Type               int    `json:"type"`                 // 类型？。目前为0
	HotType            int    `json:"hot_type"`             // 是否热门特权。0：普通特权。1：热门特权
	NewType            int    `json:"new_type"`             // 是否新特权。0：普通特权。1：新特权
	ID                 int    `json:"id"`                   // 特权子类 id
}

type Privilege struct {
	ID              int              `json:"id"`               // 特权父类 id
	Name            string           `json:"name"`             // 类型名称
	ChildPrivileges []ChildPrivilege `json:"child_privileges"` // 特权子类列表
}

type Banner struct {
	ID          int    `json:"id"`           // banner 卡片 id
	Index       int    `json:"index"`        // banner 卡片排序
	Image       string `json:"image"`        // banner 卡片图片 url
	Title       string `json:"title"`        // banner 卡片标题
	URI         string `json:"uri"`          // banner 卡片跳转页 url
	TrackParams any    `json:"track_params"` // 上报参数
}

type WelfareItem struct {
	ID          int    `json:"id"`           // 福利 id
	Name        string `json:"name"`         // 福利名称
	HomepageURI string `json:"homepage_uri"` // 福利图片 url
	BackdropURI string `json:"backdrop_uri"` // 福利图片 banner url
	Tid         int    `json:"tid"`          // (?)。目前为0
	Rank        int    `json:"rank"`         // 排列顺序
	ReceiveURI  string `json:"receive_uri"`  // 福利跳转页 url
	Image       string `json:"image"`        // 头像框图片 url
	JumpUrl     string `json:"jump_url"`     // 头像框页面 url
}

type Welfare struct {
	Count int           `json:"count"` // 福利数
	List  []WelfareItem `json:"list"`  // 福利项目列表
}

// RecommendItem 是推荐位条目，推荐头像框与推荐个性装扮共用同一返回结构。
type RecommendItem struct {
	ID      int    `json:"id"`       // 推荐位条目 id
	Name    string `json:"name"`     // 推荐位条目名称
	Image   string `json:"image"`    // 推荐位条目图片 url
	JumpURL string `json:"jump_url"` // 推荐位条目页面 url
}

// RecommendPendant 是 RecommendItem 的别名（推荐头像框）。
type RecommendPendant = RecommendItem

// RecommendCard 是 RecommendItem 的别名（推荐个性装扮）。
type RecommendCard = RecommendItem

type RecommendPendants struct {
	JumpURL string             `json:"jump_url"` // 头像框商城页面跳转 url
	List    []RecommendPendant `json:"list"`     // 推荐头像框列表
}

type RecommendCards struct {
	JumpURL string          `json:"jump_url"` // 推荐个性装扮商城页面跳转 url
	List    []RecommendCard `json:"list"`     // 推荐个性装扮列表
}

type Sort struct {
	Key  string `json:"key"`  // 扩展 row 字段名
	Sort int    `json:"sort"` // 排列顺序
}

type PointInfo struct {
	Point       int `json:"point"`        // 当前拥有大积分数量
	ExpirePoint int `json:"expire_point"` // 失效积分？。目前为0
	ExpireTime  int `json:"expire_time"`  // 失效时间？。目前为0
	ExpireDays  int `json:"expire_days"`  // 失效天数？。目前为0
}

type SignInfo struct {
	SignRemind   bool `json:"sign_remind"`   // (?)
	Benefit      int  `json:"benefit"`       // 签到收益。单位为积分
	BonusBenefit int  `json:"bonus_benefit"` // (?)
	NormalRemind bool `json:"normal_remind"` // (?)
	MuggleTask   bool `json:"muggle_task"`   // (?)
	ExpValue     int  `json:"exp_value"`     // 签到经验值？。实测为 3，作用尚不明确
}

type BigPoint struct {
	PointInfo      PointInfo `json:"point_info"` // 点数信息
	SignInfo       SignInfo  `json:"sign_info"`  // 签到信息
	SkuInfo        any       `json:"sku_info"`   // 大积分商品预览
	Tips           bool      `json:"tips"`
	PointSwitchOff any       `json:"point_switch_off"`
	SkuPriceHidden bool      `json:"sku_price_hidden"` // 是否隐藏商品价格？。实测为 false，作用尚不明确
}

type VipCenterInfo struct {
	User              VipUser           `json:"user"`               // 用户信息
	Wallet            VipWallet         `json:"wallet"`             // 钱包信息
	UnionVip          any               `json:"union_vip"`          // 联合会员信息列表，web 端：null。APP 端：array
	OtherOpenInfo     any               `json:"other_open_info"`    // 其他开通方式信息列表，web 端：null。APP 端：array
	Privileges        []Privilege       `json:"privileges"`         // 会员特权信息列表
	Banners           []Banner          `json:"banners"`            // banner 卡片列表。web 端为空
	Welfare           Welfare           `json:"welfare"`            // 福利信息
	RecommendPendants RecommendPendants `json:"recommend_pendants"` // 推荐头像框信息
	RecommendCards    RecommendCards    `json:"recommend_cards"`    // 推荐装扮信息
	Sort              []Sort            `json:"sort"`
	InReview          bool              `json:"in_review"`
	BigPoint          BigPoint          `json:"big_point"`        // 大积分信息
	HotList           any               `json:"hot_list"`         // 热门榜单类型信息
	IntegrationTask   bool              `json:"integration_task"` // 是否整合任务？。实测为 false，作用尚不明确
	FreeWelfare       []FreeWelfare     `json:"free_welfare"`     // 免费福利列表
	ExtraParamas      ExtraParamas      `json:"extra_paramas"`    // 额外参数
	HitAb             bool              `json:"hit_ab"`           // 是否命中 AB 测试？。实测为 false，作用尚不明确
}

// FreeWelfare `data` 中 `free_welfare` 数组中的对象
type FreeWelfare struct {
	Id             int    `json:"id"`             // 福利 id
	Icon           string `json:"icon"`           // 福利图标 url
	Title          string `json:"title"`          // 福利标题
	TitleColor     string `json:"titleColor"`     // 标题颜色。颜色码，空字符串表示使用默认颜色
	Subtitle       string `json:"subtitle"`       // 副标题
	SubtitleColor  string `json:"subtitleColor"`  // 副标题颜色。颜色码，空字符串表示使用默认颜色
	Subtitle2      string `json:"subtitle2"`      // 副标题 2
	Subtitle2Color string `json:"subtitle2Color"` // 副标题 2 颜色。颜色码，空字符串表示使用默认颜色
	TaskState      int    `json:"task_state"`     // 任务状态。实测为 1
	Link           string `json:"link"`           // 跳转链接
	Channel        string `json:"channel"`        // 渠道标识。如 outer_h5
}

// ExtraParamas `data` 中 `extra_paramas` 对象。实测 21 个字段，过半是 AB 测试分组
type ExtraParamas struct {
	SwitchOn              bool   `json:"switch_on"`                // 开关是否开启。实测为 false
	BirthdaySkuSwitchOn   bool   `json:"birthday_sku_switch_on"`   // 生日商品开关是否开启。实测为 false
	IsBuyBirthdaySku      bool   `json:"is_buy_birthday_sku"`      // 是否购买生日商品。实测为 false
	IsBirthdayOff         bool   `json:"is_birthday_off"`          // 生日是否已关闭。实测为 false
	PhoneBindState        int    `json:"phone_bind_state"`         // 手机号绑定状态。实测为 0
	AppTimes              int    `json:"app_times"`                // APP 次数？。实测为 0，作用尚不明确
	NaAb                  int    `json:"na_ab"`                    // na AB 测试分组？。实测为 0，作用尚不明确
	NaAbGroupId           string `json:"na_ab_group_id"`           // na AB 测试分组 id？。实测为空，作用尚不明确
	IsSameToLastSession   bool   `json:"is_same_to_last_session"`  // 是否与上次会话相同。实测为 false
	IsWhiteList           bool   `json:"is_white_list"`            // 是否在白名单中。实测为 false
	CloudAb               int    `json:"cloud_ab"`                 // cloud AB 测试分组？。实测为 0，作用尚不明确
	CloudAbGroupId        string `json:"cloud_ab_group_id"`        // cloud AB 测试分组 id？。实测为空，作用尚不明确
	BannerAb              int    `json:"banner_ab"`                // banner AB 测试分组？。实测为 0，作用尚不明确
	BannerAbGroupId       string `json:"banner_ab_group_id"`       // banner AB 测试分组 id？。实测为空，作用尚不明确
	NowTime               int    `json:"now_time"`                 // 当前时间？。实测为 0，作用尚不明确
	FreeWelfareAb         int    `json:"free_welfare_ab"`          // 免费福利 AB 测试分组？。实测为 0，作用尚不明确
	DeviceLimitAb         int    `json:"device_limit_ab"`          // 设备限制 AB 测试分组？。实测为 0，作用尚不明确
	OfflineAb             int    `json:"offline_ab"`               // 线下 AB 测试分组？。实测为 0，作用尚不明确
	WelfareModuleHeightAb int    `json:"welfare_module_height_ab"` // 福利模块高度 AB 测试分组？。实测为 0，作用尚不明确
	MyPrivilegeUpgradeAb  int    `json:"my_privilege_upgrade_ab"`  // 我的特权升级 AB 测试分组？。实测为 0，作用尚不明确
	HitDailyMustGet       bool   `json:"hit_daily_must_get"`       // 是否命中每日必领。实测为 false
}
