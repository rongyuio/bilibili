package vip

// 大会员中心与大积分相关响应模型。
// 每个接口的返回体独立定义，不与其他接口复用类型。

// GetVipCenterCombineResult 是大会员中心信息接口的 `data` 对象。
type GetVipCenterCombineResult struct {
	User              GetVipCenterCombineUser              `json:"user"`               // 用户信息
	Wallet            GetVipCenterCombineWallet            `json:"wallet"`             // 钱包信息
	UnionVip          []GetVipCenterCombineUnionVip        `json:"union_vip"`          // 联合会员信息列表。web 端为 null，APP 端为数组
	OtherOpenInfo     []GetVipCenterCombineOtherOpenInfo   `json:"other_open_info"`    // 其他开通方式信息列表。web 端为 null，APP 端为数组
	Privileges        []GetVipCenterCombinePrivilege       `json:"privileges"`         // 会员特权信息列表
	Banners           []GetVipCenterCombineBanner          `json:"banners"`            // banner 卡片列表。web 端为空
	Welfare           GetVipCenterCombineWelfare           `json:"welfare"`            // 福利信息
	RecommendPendants GetVipCenterCombineRecommendPendants `json:"recommend_pendants"` // 推荐头像框信息
	RecommendCards    GetVipCenterCombineRecommendCards    `json:"recommend_cards"`    // 推荐装扮信息
	Sort              []GetVipCenterCombineSort            `json:"sort"`               // 扩展 row 排序
	InReview          bool                                 `json:"in_review"`          // (?)
	BigPoint          GetVipCenterCombineBigPoint          `json:"big_point"`          // 大积分信息
	HotList           GetVipCenterCombineHotList           `json:"hot_list"`           // 热门榜单类型信息
	IntegrationTask   bool                                 `json:"integration_task"`   // 是否整合任务？实测为 false，作用尚不明确
	FreeWelfare       []GetVipCenterCombineFreeWelfare     `json:"free_welfare"`       // 免费福利列表
	ExtraParamas      GetVipCenterCombineExtraParamas      `json:"extra_paramas"`      // 额外参数
	HitAb             bool                                 `json:"hit_ab"`             // 是否命中 AB 测试？实测为 false，作用尚不明确
}

// GetVipCenterCombineUser 是 `data` 中的 `user` 对象。
type GetVipCenterCombineUser struct {
	Account              GetVipCenterCombineAccount       `json:"account"`                // 账号基本信息
	Vip                  GetVipCenterCombineVip           `json:"vip"`                    // 账号会员信息
	Tv                   GetVipCenterCombineTv            `json:"tv"`                     // 电视会员信息
	BackgroundImageSmall string                           `json:"background_image_small"` // 空
	BackgroundImageBig   string                           `json:"background_image_big"`   // 空
	PanelTitle           string                           `json:"panel_title"`            // 用户昵称
	PanelSubTitle        string                           `json:"panel_sub_title"`        // 面板副标题。如 大会员已过期
	AvatarPendant        GetVipCenterCombineAvatarPendant `json:"avatar_pendant"`         // 用户头像框信息
	VipOverdueExplain    string                           `json:"vip_overdue_explain"`    // 大会员提示文案。有效期 / 到期
	TvOverdueExplain     string                           `json:"tv_overdue_explain"`     // 电视大会员提示文案。有效期 / 到期
	AccountExceptionText string                           `json:"account_exception_text"` // 空
	IsAutoRenew          bool                             `json:"is_auto_renew"`          // 是否自动续费
	IsTvAutoRenew        bool                             `json:"is_tv_auto_renew"`       // 是否自动续费电视大会员
	SurplusSeconds       int64                            `json:"surplus_seconds"`        // 大会员到期剩余时间。单位为秒
	VipKeepTime          int64                            `json:"vip_keep_time"`          // 持续开通大会员时间。单位为秒
	Renew                GetVipCenterCombineRenew         `json:"renew"`                  // (?)
	Notice               GetVipCenterCombineNotice        `json:"notice"`                 // (?)
}

// GetVipCenterCombineAccount 是 `user` 中的 `account` 对象。
type GetVipCenterCombineAccount struct {
	Mid            int    `json:"mid"`              // 用户 mid
	Name           string `json:"name"`             // 昵称
	Sex            string `json:"sex"`              // 性别。男 / 女 / 保密
	Face           string `json:"face"`             // 头像 url
	Sign           string `json:"sign"`             // 签名
	Rank           int    `json:"rank"`             // 等级
	Birthday       int64  `json:"birthday"`         // 生日。秒时间戳
	IsFakeAccount  int    `json:"is_fake_account"`  // (?)
	IsDeleted      int    `json:"is_deleted"`       // 是否注销。0：正常。1：注销
	InRegAudit     int    `json:"in_reg_audit"`     // 是否注册审核。0：正常。1：审核
	IsSeniorMember int    `json:"is_senior_member"` // 是否转正。0：未转正。1：正式会员
}

// GetVipCenterCombineVip 是 `user` 中的 `vip` 对象。
type GetVipCenterCombineVip struct {
	Mid                  int                             `json:"mid"`                     // 用户 mid
	VipType              int                             `json:"vip_type"`                // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	VipStatus            int                             `json:"vip_status"`              // 会员状态。0：无。1：有
	VipDueDate           int64                           `json:"vip_due_date"`            // 会员过期时间。毫秒时间戳
	VipPayType           int                             `json:"vip_pay_type"`            // 支付类型。0：未支付（常见于官方账号）。1：已支付
	ThemeType            int                             `json:"theme_type"`              // (?)
	Label                GetVipCenterCombineLabel        `json:"label"`                   // 会员标签
	AvatarSubscript      int                             `json:"avatar_subscript"`        // 是否显示会员图标。0：不显示。1：显示
	AvatarSubscriptURL   string                          `json:"avatar_subscript_url"`    // 会员图标 url
	NicknameColor        string                          `json:"nickname_color"`          // 会员昵称颜色。颜色码，一般为#FB7299
	IsNewUser            bool                            `json:"is_new_user"`             // (?)
	TipMaterial          *GetVipCenterCombineTipMaterial `json:"tip_material"`            // (?)。无效时为 null
	VipIsValid           bool                            `json:"vip_is_valid"`            // 大会员是否有效。实测为 false
	VipIsOverdue         bool                            `json:"vip_is_overdue"`          // 大会员是否已过期。实测为 true
	VipExpireDays        int                             `json:"vip_expire_days"`         // 大会员过期天数。实测为 -38
	IsTvVip              bool                            `json:"is_tv_vip"`               // 是否为电视大会员
	VipMembershipIsVip   bool                            `json:"vip_membership_is_vip"`   // 是否为大会员？作用尚不明确
	SuperMembershipIsVip bool                            `json:"super_membership_is_vip"` // 是否为超级大会员？作用尚不明确
}

// GetVipCenterCombineTipMaterial 是 `vip` 中的 `tip_material` 对象。
type GetVipCenterCombineTipMaterial struct {
	WhiteTip         string `json:"white_tip"`          // 白色提示文字。实测为空
	Link             string `json:"link"`               // 跳转链接。实测为空
	Subscript        string `json:"subscript"`          // 下标文字。实测为空
	OrderTrackParams any    `json:"order_track_params"` // 订单追踪参数。无效时为 null
}

// GetVipCenterCombineLabel 是 `vip` 中的 `label` 对象。
type GetVipCenterCombineLabel struct {
	Text                  string `json:"text"`                      // 会员类型文案。如 大会员 / 年度大会员
	LabelTheme            string `json:"label_theme"`               // 会员标签。如 annual_vip
	TextColor             string `json:"text_color"`                // 会员标签文本颜色
	BgStyle               int    `json:"bg_style"`                  // 1
	BgColor               string `json:"bg_color"`                  // 会员标签背景颜色
	BorderColor           string `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool   `json:"use_img_label"`             // true
	ImgLabelUriHans       string `json:"img_label_uri_hans"`        // 空
	ImgLabelUriHant       string `json:"img_label_uri_hant"`        // 空
	ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
}

// GetVipCenterCombineTv 是 `user` 中的 `tv` 对象。
type GetVipCenterCombineTv struct {
	Type       int   `json:"type"`         // 电视大会员类型。0：无。1：月大会员。2：年度及以上大会员
	VipPayType int   `json:"vip_pay_type"` // 电视大会员支付类型。0：未支付。1：已支付
	Status     int   `json:"status"`       // 电视大会员状态。0：无。1：有
	DueDate    int64 `json:"due_date"`     // 电视大会员过期时间。毫秒时间戳
}

// GetVipCenterCombineAvatarPendant 是 `user` 中的 `avatar_pendant` 对象。
type GetVipCenterCombineAvatarPendant struct {
	Image             string `json:"image"`               // 头像框 url
	ImageEnhance      string `json:"image_enhance"`       // 头像框 url。动态图
	ImageEnhanceFrame string `json:"image_enhance_frame"` // 动态头像框帧波普版 url
}

// GetVipCenterCombineRenew 是 `user` 中的 `renew` 对象。
type GetVipCenterCombineRenew struct {
	Text string `json:"text"` // (?)
	Link string `json:"link"` // (?)
}

// GetVipCenterCombineNotice 是 `user` 中的 `notice` 对象。
type GetVipCenterCombineNotice struct {
	Text             string `json:"text"`               // (?)
	TvText           string `json:"tv_text"`            // (?)
	Type             int    `json:"type"`               // (?)
	CanClose         bool   `json:"can_close"`          // (?)
	SurplusSeconds   int64  `json:"surplus_seconds"`    // (?)
	TvSurplusSeconds int64  `json:"tv_surplus_seconds"` // (?)
}

// GetVipCenterCombineWallet 是 `data` 中的 `wallet` 对象。
type GetVipCenterCombineWallet struct {
	Coupon            int  `json:"coupon"`             // 当前 B 币券
	Point             int  `json:"point"`              // (?)
	PrivilegeReceived bool `json:"privilege_received"` // (?)
}

// GetVipCenterCombineUnionVip 是 `union_vip` 数组中的对象。仅 APP 端存在该信息。
type GetVipCenterCombineUnionVip struct {
	Image2x string `json:"image_2x"` // 联合会员图片小 url
	Image3x string `json:"image_3x"` // 联合会员图片大 url
	Title   string `json:"title"`    // 联合会员名称
	Link    string `json:"link"`     // 联合会员开通页 url
	Sort    int    `json:"sort"`     // 排列顺序
}

// GetVipCenterCombineOtherOpenInfo 是 `other_open_info` 数组中的对象。仅 APP 端存在该信息。
type GetVipCenterCombineOtherOpenInfo struct {
	Title   string `json:"title"`    // 其他开通方式名称
	URL     string `json:"url"`      // 其他开通方式开通页 url
	IconURL string `json:"icon_url"` // 其他开通方式图标 url
	Desc    string `json:"desc"`     // 提示文案
	Sort    int    `json:"sort"`     // 排列顺序
}

// GetVipCenterCombinePrivilege 是 `privileges` 数组中的对象（特权父类）。
type GetVipCenterCombinePrivilege struct {
	ID              int                                 `json:"id"`               // 特权父类 id
	Name            string                              `json:"name"`             // 类型名称
	ChildPrivileges []GetVipCenterCombineChildPrivilege `json:"child_privileges"` // 特权子类列表
}

// GetVipCenterCombineChildPrivilege 是 `child_privileges` 数组中的对象（特权子类）。
type GetVipCenterCombineChildPrivilege struct {
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
	Type               int    `json:"type"`                 // 类型？目前为0
	HotType            int    `json:"hot_type"`             // 是否热门特权。0：普通特权。1：热门特权
	NewType            int    `json:"new_type"`             // 是否新特权。0：普通特权。1：新特权
	ID                 int    `json:"id"`                   // 特权子类 id
}

// GetVipCenterCombineBanner 是 `banners` 数组中的对象。仅 APP 端存在该信息。
type GetVipCenterCombineBanner struct {
	ID          int    `json:"id"`           // banner 卡片 id
	Index       int    `json:"index"`        // banner 卡片排序
	Image       string `json:"image"`        // banner 卡片图片 url
	Title       string `json:"title"`        // banner 卡片标题
	URI         string `json:"uri"`          // banner 卡片跳转页 url
	TrackParams any    `json:"track_params"` // 上报参数
}

// GetVipCenterCombineWelfare 是 `data` 中的 `welfare` 对象。
type GetVipCenterCombineWelfare struct {
	Count int                              `json:"count"` // 福利数
	List  []GetVipCenterCombineWelfareItem `json:"list"`  // 福利项目列表
}

// GetVipCenterCombineWelfareItem 是 `welfare.list` 数组中的对象。
type GetVipCenterCombineWelfareItem struct {
	ID          int    `json:"id"`           // 福利 id
	Name        string `json:"name"`         // 福利名称
	HomepageURI string `json:"homepage_uri"` // 福利图片 url
	BackdropURI string `json:"backdrop_uri"` // 福利图片 banner url
	Tid         int    `json:"tid"`          // (?)。目前为0
	Rank        int    `json:"rank"`         // 排列顺序
	ReceiveURI  string `json:"receive_uri"`  // 福利跳转页 url
}

// GetVipCenterCombineRecommendPendant 是推荐头像框条目。
type GetVipCenterCombineRecommendPendant struct {
	ID      int    `json:"id"`       // 推荐位条目 id
	Name    string `json:"name"`     // 推荐位条目名称
	Image   string `json:"image"`    // 推荐位条目图片 url
	JumpURL string `json:"jump_url"` // 推荐位条目页面 url
}

// GetVipCenterCombineRecommendPendants 是 `data` 中的 `recommend_pendants` 对象。
type GetVipCenterCombineRecommendPendants struct {
	JumpURL string                                `json:"jump_url"` // 头像框商城页面跳转 url
	List    []GetVipCenterCombineRecommendPendant `json:"list"`     // 推荐头像框列表
}

// GetVipCenterCombineRecommendCard 是推荐个性装扮条目。
type GetVipCenterCombineRecommendCard struct {
	ID      int    `json:"id"`       // 推荐位条目 id
	Name    string `json:"name"`     // 推荐位条目名称
	Image   string `json:"image"`    // 推荐位条目图片 url
	JumpURL string `json:"jump_url"` // 推荐位条目页面 url
}

// GetVipCenterCombineRecommendCards 是 `data` 中的 `recommend_cards` 对象。
type GetVipCenterCombineRecommendCards struct {
	JumpURL string                             `json:"jump_url"` // 推荐个性装扮商城页面跳转 url
	List    []GetVipCenterCombineRecommendCard `json:"list"`     // 推荐个性装扮列表
}

// GetVipCenterCombineSort 是 `sort` 数组中的对象。
type GetVipCenterCombineSort struct {
	Key  string `json:"key"`  // 扩展 row 字段名
	Sort int    `json:"sort"` // 排列顺序
}

// GetVipCenterCombineBigPoint 是 `data` 中的 `big_point` 对象。
type GetVipCenterCombineBigPoint struct {
	PointInfo      GetVipCenterCombinePointInfo `json:"point_info"`       // 点数信息
	SignInfo       GetVipCenterCombineSignInfo  `json:"sign_info"`        // 签到信息
	SkuInfo        GetVipCenterCombineSkuInfo   `json:"sku_info"`         // 大积分商品预览
	Tips           any                          `json:"tips"`             // 文档表格标为 bool，实测为 [{"content": ...}] 数组，故用 any
	PointSwitchOff any                          `json:"point_switch_off"` // 文档表格标为 obj，实测为 false，故用 any
	SkuPriceHidden bool                         `json:"sku_price_hidden"` // 是否隐藏商品价格？作用尚不明确
}

// GetVipCenterCombinePointInfo 是大积分点数信息。
type GetVipCenterCombinePointInfo struct {
	Point       int `json:"point"`        // 当前拥有大积分数量
	ExpirePoint int `json:"expire_point"` // 失效积分？目前为0
	ExpireTime  int `json:"expire_time"`  // 失效时间？目前为0
	ExpireDays  int `json:"expire_days"`  // 失效天数？目前为0
}

// GetVipCenterCombineSignInfo 是大积分签到信息。
type GetVipCenterCombineSignInfo struct {
	SignRemind   bool `json:"sign_remind"`   // (?)
	Benefit      int  `json:"benefit"`       // 签到收益。单位为积分
	BonusBenefit int  `json:"bonus_benefit"` // (?)
	NormalRemind bool `json:"normal_remind"` // (?)
	MuggleTask   bool `json:"muggle_task"`   // (?)
	ExpValue     int  `json:"exp_value"`     // 签到经验值？实测为 3，作用尚不明确
}

// GetVipCenterCombineSkuInfo 是 `big_point` 中的 `sku_info` 对象。
type GetVipCenterCombineSkuInfo struct {
	Skus []GetVipCenterCombineSku `json:"skus"` // 大积分商品列表
}

// GetVipCenterCombineSku 是大积分商品。
type GetVipCenterCombineSku struct {
	Base GetVipCenterCombineSkuBase `json:"base"` // 商品信息
}

// GetVipCenterCombineSkuBase 是大积分商品详情。
type GetVipCenterCombineSkuBase struct {
	Token             string                       `json:"token"`               // 商品 token。用于购买 & 明细页面
	Title             string                       `json:"title"`               // 商品名称
	Picture           string                       `json:"picture"`             // 商品图片 url
	RotationPictures  []string                     `json:"rotation_pictures"`   // 商品图片组
	Price             GetVipCenterCombinePrice     `json:"price"`               // 价格信息
	Inventory         GetVipCenterCombineInventory `json:"inventory"`           // 库存信息
	UserType          int                          `json:"user_type"`           // 2
	ExchangeLimitType int                          `json:"exchange_limit_type"` // 2 / 3 / 4
	ExchangeLimitNum  int                          `json:"exchange_limit_num"`  // 限购数量
	StartTime         int64                        `json:"start_time"`          // 起售时间。秒时间戳
	EndTime           int64                        `json:"end_time"`            // 止售时间。秒时间戳
	State             int                          `json:"state"`               // 状态？2
}

// GetVipCenterCombinePrice 是 `base` 中的 `price` 对象。
type GetVipCenterCombinePrice struct {
	Origin    int                           `json:"origin"`    // 商品原价。单位为积分
	Sale      int                           `json:"sale"`      // 售价。单位为积分
	Promotion *GetVipCenterCombinePromotion `json:"promotion"` // 折扣信息。无折扣时为 null
}

// GetVipCenterCombinePromotion 是 `price` 中的 `promotion` 对象。
type GetVipCenterCombinePromotion struct {
	Price    int    `json:"price"`    // 折后价格。单位为积分
	Type     int    `json:"type"`     // 折扣类型。1：普通折扣。2：秒杀
	Discount int    `json:"discount"` // 折扣力度
	Label    string `json:"label"`    // 标签文案。如 5折 / 秒杀
}

// GetVipCenterCombineInventory 是 `base` 中的 `inventory` 对象。
type GetVipCenterCombineInventory struct {
	AvailableNum int   `json:"available_num"`  // 库存总量
	UsedNum      int   `json:"used_num"`       // 已售数量
	SurplusNum   int   `json:"surplus_num"`    // 剩余数量
	IsSoldOut    bool  `json:"is_sold_out"`    // 是否售罄。实测为 false
	NextSoldTime int64 `json:"next_sold_time"` // 下次开售时间？秒时间戳
}

// GetVipCenterCombineHotList 是 `data` 中的 `hot_list` 对象。
type GetVipCenterCombineHotList struct {
	Taps []GetVipCenterCombineHotTap `json:"taps"` // 热门类型列表
}

// GetVipCenterCombineHotTap 是 `hot_list.taps` 数组中的对象。
type GetVipCenterCombineHotTap struct {
	Oid       string `json:"oid"`        // 分类数据。类似 JSON
	RankID    int    `json:"rank_id"`    // 分类 id
	RankTitle string `json:"rank_title"` // 分类名称
}

// GetVipCenterCombineFreeWelfare 是 `free_welfare` 数组中的对象。
type GetVipCenterCombineFreeWelfare struct {
	ID             int    `json:"id"`             // 福利 id
	Icon           string `json:"icon"`           // 福利图标 url
	Title          string `json:"title"`          // 福利标题
	TitleColor     string `json:"titleColor"`     // 标题颜色。空字符串表示使用默认颜色
	Subtitle       string `json:"subtitle"`       // 副标题
	SubtitleColor  string `json:"subtitleColor"`  // 副标题颜色。空字符串表示使用默认颜色
	Subtitle2      string `json:"subtitle2"`      // 副标题 2
	Subtitle2Color string `json:"subtitle2Color"` // 副标题 2 颜色。空字符串表示使用默认颜色
	TaskState      int    `json:"task_state"`     // 任务状态。实测为 1
	Link           string `json:"link"`           // 跳转链接
	Channel        string `json:"channel"`        // 渠道标识。如 outer_h5
}

// GetVipCenterCombineExtraParamas 是 `data` 中的 `extra_paramas` 对象。过半是 AB 测试分组。
type GetVipCenterCombineExtraParamas struct {
	SwitchOn              bool   `json:"switch_on"`                // 开关是否开启。实测为 false
	BirthdaySkuSwitchOn   bool   `json:"birthday_sku_switch_on"`   // 生日商品开关是否开启。实测为 false
	IsBuyBirthdaySku      bool   `json:"is_buy_birthday_sku"`      // 是否购买生日商品。实测为 false
	IsBirthdayOff         bool   `json:"is_birthday_off"`          // 生日是否已关闭。实测为 false
	PhoneBindState        int    `json:"phone_bind_state"`         // 手机号绑定状态。实测为 0
	AppTimes              int    `json:"app_times"`                // APP 次数？作用尚不明确
	NaAb                  int    `json:"na_ab"`                    // na AB 测试分组？作用尚不明确
	NaAbGroupID           string `json:"na_ab_group_id"`           // na AB 测试分组 id？作用尚不明确
	IsSameToLastSession   bool   `json:"is_same_to_last_session"`  // 是否与上次会话相同。实测为 false
	IsWhiteList           bool   `json:"is_white_list"`            // 是否在白名单中。实测为 false
	CloudAb               int    `json:"cloud_ab"`                 // cloud AB 测试分组？作用尚不明确
	CloudAbGroupID        string `json:"cloud_ab_group_id"`        // cloud AB 测试分组 id？作用尚不明确
	BannerAb              int    `json:"banner_ab"`                // banner AB 测试分组？作用尚不明确
	BannerAbGroupID       string `json:"banner_ab_group_id"`       // banner AB 测试分组 id？作用尚不明确
	NowTime               int    `json:"now_time"`                 // 当前时间？作用尚不明确
	FreeWelfareAb         int    `json:"free_welfare_ab"`          // 免费福利 AB 测试分组？作用尚不明确
	DeviceLimitAb         int    `json:"device_limit_ab"`          // 设备限制 AB 测试分组？作用尚不明确
	OfflineAb             int    `json:"offline_ab"`               // 线下 AB 测试分组？作用尚不明确
	WelfareModuleHeightAb int    `json:"welfare_module_height_ab"` // 福利模块高度 AB 测试分组？作用尚不明确
	MyPrivilegeUpgradeAb  int    `json:"my_privilege_upgrade_ab"`  // 我的特权升级 AB 测试分组？作用尚不明确
	HitDailyMustGet       bool   `json:"hit_daily_must_get"`       // 是否命中每日必领。实测为 false
}

// GetVipPointHomepageCombineResult 是大积分中心信息接口的 `data` 对象。
type GetVipPointHomepageCombineResult struct {
	VipInfo         GetVipPointHomepageCombineVipInfo         `json:"vip_info"`         // 大会员信息
	Content         string                                    `json:"content"`          // 空
	PointInfo       GetVipPointHomepageCombinePointInfo       `json:"point_info"`       // 点数信息
	Task            GetVipPointHomepageCombineTask            `json:"task"`             // 任务信息
	Banner          []GetVipPointHomepageCombineBanner        `json:"banner"`           // banner 卡片列表
	GoodsCategory   []GetVipPointHomepageCombineGoodsCategory `json:"goods_category"`   // 商品类型列表
	GoodsSkus       []GetVipPointHomepageCombineSku           `json:"goods_skus"`       // 商品明细列表。包含所有类型的商品
	CurrentTS       int64                                     `json:"current_ts"`       // 本次请求时间。秒时间戳
	IntegrationTask bool                                      `json:"integration_task"` // 是否整合任务？作用尚不明确
}

// GetVipPointHomepageCombineVipInfo 是 `data` 中的 `vip_info` 对象。
type GetVipPointHomepageCombineVipInfo struct {
	Type            int                                   `json:"type"`             // 大会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status          int                                   `json:"status"`           // 大会员状态。0：无。1：有
	DueDate         int64                                 `json:"due_date"`         // 大会员过期时间。毫秒时间戳
	VipPayType      int                                   `json:"vip_pay_type"`     // 支付类型。0：未支付。1：已支付
	StartTime       int64                                 `json:"start_time"`       // 首次大会员开通时间。秒时间戳
	PaidType        int                                   `json:"paid_type"`        // 0
	Mid             int                                   `json:"mid"`              // 当前用户 mid
	Role            int                                   `json:"role"`             // 大会员类型。1：月度。3：年度。7：十年。15：百年
	TvVipStatus     int                                   `json:"tv_vip_status"`    // 电视大会员状态。0：无。1：有
	TvVipPayType    int                                   `json:"tv_vip_pay_type"`  // 电视大会员支付类型。0：未支付。1：已支付
	TvDueDate       int64                                 `json:"tv_due_date"`      // 电视大会员过期时间。秒时间戳
	VipMembership   *GetVipPointHomepageCombineMembership `json:"vip_membership"`   // 大会员身份信息。无效时为 null
	OttMembership   any                                   `json:"ott_membership"`   // 电视大会员身份信息。实测为 null
	SuperMembership any                                   `json:"super_membership"` // 超级大会员身份信息。实测为 null
	VipRecentTime   int64                                 `json:"vip_recent_time"`  // 最近一次大会员开通时间。秒时间戳
}

// GetVipPointHomepageCombineMembership 是 `vip_info` 中的 `vip_membership` 对象。
type GetVipPointHomepageCombineMembership struct {
	IsVip  bool `json:"is_vip"` // 是否为大会员。实测为 false
	Status int  `json:"status"` // 大会员状态。实测为 0
}

// GetVipPointHomepageCombinePointInfo 是大积分点数信息。
type GetVipPointHomepageCombinePointInfo struct {
	Point       int `json:"point"`        // 当前拥有大积分数量
	ExpirePoint int `json:"expire_point"` // 失效积分？目前为0
	ExpireTime  int `json:"expire_time"`  // 失效时间？目前为0
	ExpireDays  int `json:"expire_days"`  // 失效天数？目前为0
}

// GetVipPointHomepageCombineTask 是 `data` 中的 `task` 对象。
type GetVipPointHomepageCombineTask struct {
	TaskItem  []GetVipPointHomepageCombineTaskItem `json:"task_item"`  // 当前页面显示的任务
	TaskCount int                                  `json:"task_count"` // 任务总数
	Signed    bool                                 `json:"signed"`     // false
	Score     int                                  `json:"score"`      // 5
	ExpValue  int                                  `json:"exp_value"`  // 完成任务可得经验值？作用尚不明确
}

// GetVipPointHomepageCombineTaskItem 是 `task.task_item` 数组中的对象。
type GetVipPointHomepageCombineTaskItem struct {
	TaskCode      string `json:"task_code"`      // 任务 id。详见大积分任务列表
	State         int    `json:"state"`          // 任务状态。0：未领取。1：未完成。3：已完成/已领取
	Title         string `json:"title"`          // 任务名称
	Icon          string `json:"icon"`           // 任务图标 url
	Subtitle      string `json:"subtitle"`       // 任务副标题
	Explain       string `json:"explain"`        // 任务详情说明
	Link          string `json:"link"`           // 任务跳转 url。web 为目标页面 url，APP 为 APP 内 uri
	VipLimit      int    `json:"vip_limit"`      // 是否为大会员专享任务。0：否。1：是
	CompleteTimes int    `json:"complete_times"` // 已完成次数
	MaxTimes      int    `json:"max_times"`      // 最大完成次数
	RecallNum     int    `json:"recall_num"`     // 重试次数？目前为0
}

// GetVipPointHomepageCombineBanner 是 `banner` 数组中的对象。
type GetVipPointHomepageCombineBanner struct {
	Link     string `json:"link"`      // 跳转 url
	ImageURL string `json:"image_url"` // banner 卡片 url
}

// GetVipPointHomepageCombineGoodsCategory 是 `goods_category` 数组中的对象。
type GetVipPointHomepageCombineGoodsCategory struct {
	ID    int    `json:"id"`    // 商品类型 id。详见大积分商品类型
	Name  string `json:"name"`  // 商品类型名称
	State int    `json:"state"` // 状态。目前为2
}

// GetVipPointHomepageCombineSku 是大积分商品。
type GetVipPointHomepageCombineSku struct {
	Base GetVipPointHomepageCombineSkuBase `json:"base"` // 商品信息
}

// GetVipPointHomepageCombineSkuBase 是大积分商品详情。
type GetVipPointHomepageCombineSkuBase struct {
	Token             string                              `json:"token"`               // 商品 token。用于购买 & 明细页面
	Title             string                              `json:"title"`               // 商品名称
	Picture           string                              `json:"picture"`             // 商品图片 url
	RotationPictures  []string                            `json:"rotation_pictures"`   // 商品图片组
	Price             GetVipPointHomepageCombinePrice     `json:"price"`               // 价格信息
	Inventory         GetVipPointHomepageCombineInventory `json:"inventory"`           // 库存信息
	UserType          int                                 `json:"user_type"`           // 2
	ExchangeLimitType int                                 `json:"exchange_limit_type"` // 2 / 3 / 4
	ExchangeLimitNum  int                                 `json:"exchange_limit_num"`  // 限购数量
	StartTime         int64                               `json:"start_time"`          // 起售时间。秒时间戳
	EndTime           int64                               `json:"end_time"`            // 止售时间。秒时间戳
	State             int                                 `json:"state"`               // 状态？2
}

// GetVipPointHomepageCombinePrice 是 `base` 中的 `price` 对象。
type GetVipPointHomepageCombinePrice struct {
	Origin    int                                  `json:"origin"`    // 商品原价。单位为积分
	Sale      int                                  `json:"sale"`      // 售价。单位为积分
	Promotion *GetVipPointHomepageCombinePromotion `json:"promotion"` // 折扣信息。无折扣时为 null
}

// GetVipPointHomepageCombinePromotion 是 `price` 中的 `promotion` 对象。
type GetVipPointHomepageCombinePromotion struct {
	Price    int    `json:"price"`    // 折后价格。单位为积分
	Type     int    `json:"type"`     // 折扣类型。1：普通折扣。2：秒杀
	Discount int    `json:"discount"` // 折扣力度
	Label    string `json:"label"`    // 标签文案。如 5折 / 秒杀
}

// GetVipPointHomepageCombineInventory 是 `base` 中的 `inventory` 对象。
type GetVipPointHomepageCombineInventory struct {
	AvailableNum int   `json:"available_num"`  // 库存总量
	UsedNum      int   `json:"used_num"`       // 已售数量
	SurplusNum   int   `json:"surplus_num"`    // 剩余数量
	IsSoldOut    bool  `json:"is_sold_out"`    // 是否售罄。实测为 false
	NextSoldTime int64 `json:"next_sold_time"` // 下次开售时间？秒时间戳
}

// GetVipPointListResult 是大积分改变记录接口的 `data` 对象。
type GetVipPointListResult struct {
	Total        int                   `json:"total"`          // 总记录数
	BigPointList []GetVipPointListItem `json:"big_point_list"` // 积分记录数组。不存在时为 null
}

// GetVipPointListItem 是 `big_point_list` 数组中的对象。
type GetVipPointListItem struct {
	Point      int    `json:"point"`       // 改变的积分数。消耗为负数
	ChangeTime int64  `json:"change_time"` // 改变时间戳。秒
	Remark     string `json:"remark"`      // 批注，积分改变的原因
	OrderNo    string `json:"order_no"`    // 编号
	ImageURL   string `json:"image_url"`   // 商品图片。消耗时存在，获取时为空
}
