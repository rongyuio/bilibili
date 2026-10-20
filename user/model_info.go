package user

// 批量获取用户详细信息接口的响应模型。
//
// 两个接口的响应形状不同：GetUserCards 的 data 是以 mid 为键的对象，
// GetUserInfos 的 data 是用户对象数组，因此各自保留独立模型。

// GetUserCardsItem 是 GetUserCards 响应 data 中以 mid 为键的用户对象。
type GetUserCardsItem struct {
	Mid        string                  `json:"mid"`         // mid（字符串形式）
	Face       string                  `json:"face"`        // 头像链接
	Name       string                  `json:"name"`        // 昵称
	Official   GetUserCardsOfficial    `json:"official"`    // 认证信息。基本同「用户空间详细信息」中的 data.official
	Vip        GetUserCardsVip         `json:"vip"`         // 会员信息。部分数值字段在本接口中为字符串
	NameRender *GetUserCardsNameRender `json:"name_render"` // 昵称渲染信息。无效时为 null
}

// GetUserCardsOfficial 是 GetUserCards 响应中的认证信息。
type GetUserCardsOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetUserCardsVip 是 GetUserCards 响应中的会员信息。
type GetUserCardsVip struct {
	AvatarSubscript    int               `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	AvatarSubscriptURL string            `json:"avatar_subscript_url"` // 大会员角标地址
	DueDate            string            `json:"due_date"`             // 会员过期时间。毫秒时间戳，字符串形式
	NicknameColor      string            `json:"nickname_color"`       // 会员昵称颜色
	Role               string            `json:"role"`                 // 大角色类型。字符串形式
	Status             int               `json:"status"`               // 会员状态。0：无。1：有
	ThemeType          int               `json:"theme_type"`           // 0。作用尚不明确
	Type               int               `json:"type"`                 // 会员类型。0：无。1：月度大会员。2：年度及以上大会员
	Label              GetUserCardsLabel `json:"label"`                // 会员标签
}

// GetUserCardsLabel 是 GetUserCards 响应中的会员标签。
type GetUserCardsLabel struct {
	Path                  string                `json:"path"`                      // 空。作用尚不明确
	Text                  string                `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                `json:"text_color"`                // 会员标签
	BgStyle               int                   `json:"bg_style"`                  // 1
	BgColor               string                `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                  `json:"use_img_label"`             // true
	ImgLabelURIHans       string                `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                   `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                   `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetUserCardsLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetUserCardsLabelGoto 是会员标签中的 label_goto 对象。
type GetUserCardsLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetUserCardsNameRender 是 GetUserCards 响应中的昵称渲染信息。
type GetUserCardsNameRender struct {
	ColorsInfo   GetUserCardsNameRenderColors `json:"colors_info"`   // 昵称颜色信息（？）。作用尚不明确
	RenderScheme string                       `json:"render_scheme"` // "Default" 或 "Colorful"。（？）。作用尚不明确
}

// GetUserCardsNameRenderColors 是昵称颜色信息。
type GetUserCardsNameRenderColors struct {
	Color    []GetUserCardsNameRenderColor `json:"color"`     // 昵称颜色（？）。作用尚不明确
	ColorIDs []string                      `json:"color_ids"` // 昵称颜色 id（？）。作用尚不明确。实测为 ["6"]
}

// GetUserCardsNameRenderColor 是昵称颜色。
type GetUserCardsNameRenderColor struct {
	ColorDay   string `json:"color_day"`   // 浅色模式昵称颜色。HEX颜色代码
	ColorNight string `json:"color_night"` // 深色模式昵称颜色。HEX颜色代码
}

// GetUserInfosItem 是 GetUserInfos 响应 data 数组中的用户对象。
type GetUserInfosItem struct {
	Mid           int                    `json:"mid"`             // mid
	Name          string                 `json:"name"`            // 昵称
	Face          string                 `json:"face"`            // 头像链接
	Sign          string                 `json:"sign"`            // 签名
	Rank          int                    `json:"rank"`            // 用户权限等级
	Level         int                    `json:"level"`           // 当前等级。0-6 级
	Silence       int                    `json:"silence"`         // 封禁状态。0：正常。1：被封
	Vip           GetUserInfosVip        `json:"vip"`             // 大会员信息
	Official      GetUserInfosOfficial   `json:"official"`        // 认证信息。文档字段表拼写有误
	IsFakeAccount int                    `json:"is_fake_account"` // 是否为虚假账号。文档标为 bool，实际返回 0/1
	ExpertInfo    GetUserInfosExpertInfo `json:"expert_info"`     // 专业信息
	JsonAva       GetUserInfosJsonAva    `json:"json_ava"`        // 头像图层信息
}

// GetUserInfosOfficial 是 GetUserInfos 响应中的认证信息。
type GetUserInfosOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetUserInfosVip 是 GetUserInfos 响应中的大会员信息。
type GetUserInfosVip struct {
	Type               int                     `json:"type"`                 // 会员类型。0：无。1：月大会员。2：年度及以上大会员
	Status             int                     `json:"status"`               // 会员状态。0：无。1：有
	DueDate            int                     `json:"due_date"`             // 会员过期时间。Unix时间戳(毫秒)
	VipPayType         int                     `json:"vip_pay_type"`         // 支付类型。0：未支付（常见于官方账号）。1：已支付（以正常渠道获取的大会员均为此值）
	ThemeType          int                     `json:"theme_type"`           // 0。作用尚不明确
	Label              GetUserInfosLabel       `json:"label"`                // 会员标签
	AvatarSubscript    int                     `json:"avatar_subscript"`     // 是否显示会员图标。0：不显示。1：显示
	NicknameColor      string                  `json:"nickname_color"`       // 会员昵称颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	Role               int                     `json:"role"`                 // 大角色类型。1：月度大会员。3：年度大会员。7：十年大会员。15：百年大会员
	AvatarSubscriptURL string                  `json:"avatar_subscript_url"` // 大会员角标地址
	TvVipStatus        int                     `json:"tv_vip_status"`        // 电视大会员状态。0：未开通
	TvVipPayType       int                     `json:"tv_vip_pay_type"`      // 电视大会员支付类型
	OttInfo            GetUserInfosVipOttInfo  `json:"ott_info"`             // （？）。作用尚不明确
	SuperVip           GetUserInfosVipSuperVip `json:"super_vip"`            // （？）。作用尚不明确
	TvDueDate          int                     `json:"tv_due_date"`          // 电视大会员过期时间。秒级时间戳
	AvatarIcon         GetUserInfosAvatarIcon  `json:"avatar_icon"`          // 大会员角标信息
}

// GetUserInfosLabel 是 GetUserInfos 响应中的会员标签。
type GetUserInfosLabel struct {
	Path                  string                `json:"path"`                      // 空。作用尚不明确
	Text                  string                `json:"text"`                      // 会员类型文案。大会员 年度大会员 十年大会员 百年大会员 最强绿鲤鱼
	LabelTheme            string                `json:"label_theme"`               // 会员标签。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十 年大会员。hundred_annual_vip：百年大会员。fools_day_hundred_annual_vip：最强绿鲤鱼
	TextColor             string                `json:"text_color"`                // 会员标签
	BgStyle               int                   `json:"bg_style"`                  // 1
	BgColor               string                `json:"bg_color"`                  // 会员标签背景颜色。颜色码，一般为#FB7299，曾用于愚人节改变大会员配色
	BorderColor           string                `json:"border_color"`              // 会员标签边框颜色。未使用
	UseImgLabel           bool                  `json:"use_img_label"`             // true
	ImgLabelURIHans       string                `json:"img_label_uri_hans"`        // 空串
	ImgLabelURIHant       string                `json:"img_label_uri_hant"`        // 空串
	ImgLabelURIHansStatic string                `json:"img_label_uri_hans_static"` // 大会员牌子图片。简体版
	ImgLabelURIHantStatic string                `json:"img_label_uri_hant_static"` // 大会员牌子图片。繁体版
	LabelId               int                   `json:"label_id"`                  // （？）。作用尚不明确
	LabelType             int                   `json:"label_type"`                // （？）。作用尚不明确
	LabelGoto             GetUserInfosLabelGoto `json:"label_goto"`                // （？）。作用尚不明确
	ImgLabelUriI18n       string                `json:"img_label_uri_i18n"`        // （？）。作用尚不明确
	ImgLabelUriI18nStatic string                `json:"img_label_uri_i18n_static"` // （？）。作用尚不明确
}

// GetUserInfosLabelGoto 是会员标签中的 label_goto 对象。
type GetUserInfosLabelGoto struct {
	Mobile string `json:"mobile"` // （？）。作用尚不明确
	PcWeb  string `json:"pc_web"` // （？）。作用尚不明确
}

// GetUserInfosVipOttInfo 是大会员信息中的 ott_info 对象。
type GetUserInfosVipOttInfo struct {
	VipType      int    `json:"vip_type"`       // （？）。作用尚不明确
	PayType      int    `json:"pay_type"`       // （？）。作用尚不明确
	PayChannelId string `json:"pay_channel_id"` // （？）。作用尚不明确
	Status       int    `json:"status"`         // （？）。作用尚不明确
	OverdueTime  int    `json:"overdue_time"`   // （？）。作用尚不明确
}

// GetUserInfosVipSuperVip 是大会员信息中的 super_vip 对象。
type GetUserInfosVipSuperVip struct {
	IsSuperVip bool `json:"is_super_vip"` // （？）。作用尚不明确
}

// GetUserInfosAvatarIcon 是大会员角标信息。
type GetUserInfosAvatarIcon struct {
	IconType     int `json:"icon_type"`     // （？）。作用尚不明确
	IconResource any `json:"icon_resource"` // （？）。作用尚不明确
}

// GetUserInfosExpertInfo 是 GetUserInfos 响应中的专业信息。
type GetUserInfosExpertInfo struct {
	Title string `json:"title"` // 头衔
	State int    `json:"state"` // 状态
	Type  int    `json:"type"`  // 类型
	Desc  string `json:"desc"`  // 描述
}

// GetUserInfosJsonAva 是头像图层信息。
type GetUserInfosJsonAva struct {
	ContainerSize  GetUserInfosJsonAvaContainerSize  `json:"container_size"`  // 头像容器尺寸
	FallbackLayers GetUserInfosJsonAvaFallbackLayers `json:"fallback_layers"` // 头像图层信息
	Mid            string                            `json:"mid"`             // 用户 mid。实测为字符串形式
}

// GetUserInfosJsonAvaContainerSize 是头像容器尺寸。
type GetUserInfosJsonAvaContainerSize struct {
	Width  float64 `json:"width"`  // 容器宽度。实测为 1.375 或 1.35
	Height float64 `json:"height"` // 容器高度。同上
}

// GetUserInfosJsonAvaFallbackLayers 是头像降级渲染图层分组。
type GetUserInfosJsonAvaFallbackLayers struct {
	Layers          []GetUserInfosJsonAvaLayer `json:"layers"`            // 图层数组
	IsCriticalGroup bool                       `json:"is_critical_group"` // 是否为关键分组。实测为 true
}

// GetUserInfosJsonAvaLayer 是头像渲染层中的一层。
type GetUserInfosJsonAvaLayer struct {
	Visible     bool                               `json:"visible"`      // 是否可见。实测为 true
	GeneralSpec GetUserInfosAvatarLayerGeneralSpec `json:"general_spec"` // 通用规格
	LayerConfig GetUserInfosJsonAvaLayerConfig     `json:"layer_config"` // 图层配置
	Resource    GetUserInfosJsonAvaResource        `json:"resource"`     // 图层资源
}

// GetUserInfosAvatarLayerGeneralSpec 是头像渲染层的通用规格。
type GetUserInfosAvatarLayerGeneralSpec struct {
	PosSpec    GetUserInfosAvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec GetUserInfosAvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   GetUserInfosAvatarLayerSizeSpec   `json:"size_spec"`
}

// GetUserInfosAvatarLayerPosSpec 是头像渲染层的位置规格。
type GetUserInfosAvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// GetUserInfosAvatarLayerRenderSpec 是头像渲染层的渲染规格。
type GetUserInfosAvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// GetUserInfosAvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type GetUserInfosAvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// GetUserInfosJsonAvaLayerConfig 是图层配置。
type GetUserInfosJsonAvaLayerConfig struct {
	Tags       map[string]any                `json:"tags"`        // 图层标签。键如 AVATAR_LAYER、PENDENT_LAYER、ICON_LAYER
	IsCritical bool                          `json:"is_critical"` // 是否为关键图层。实测为 true
	LayerMask  *GetUserInfosJsonAvaLayerMask `json:"layer_mask"`  // 图层蒙版。仅部分图层存在
}

// GetUserInfosJsonAvaLayerMask 是图层蒙版。
type GetUserInfosJsonAvaLayerMask struct {
	GeneralSpec GetUserInfosAvatarLayerGeneralSpec `json:"general_spec"` // 通用规格。结构与 layers 中的 general_spec 相同
	MaskSrc     GetUserInfosJsonAvaMaskSrc         `json:"mask_src"`     // 蒙版来源
}

// GetUserInfosJsonAvaMaskSrc 是蒙版来源。
type GetUserInfosJsonAvaMaskSrc struct {
	SrcType int                     `json:"src_type"` // 来源类型。实测为 3
	Draw    GetUserInfosJsonAvaDraw `json:"draw"`     // 绘制配置
}

// GetUserInfosJsonAvaDraw 是绘制配置。
type GetUserInfosJsonAvaDraw struct {
	DrawType    int                            `json:"draw_type"`    // 绘制类型。实测为 1
	FillMode    int                            `json:"fill_mode"`    // 填充模式。实测为 1
	ColorConfig GetUserInfosJsonAvaColorConfig `json:"color_config"` // 颜色配置
}

// GetUserInfosJsonAvaColorConfig 是绘制颜色配置。
type GetUserInfosJsonAvaColorConfig struct {
	IsDarkModeAware bool                      `json:"is_dark_mode_aware"` // 是否区分夜间模式。实测为 true
	Day             GetUserInfosJsonAvaColor  `json:"day"`                // 日间颜色
	Night           *GetUserInfosJsonAvaColor `json:"night"`              // 夜间颜色。部分图层没有
}

// GetUserInfosJsonAvaColor 是颜色值。
type GetUserInfosJsonAvaColor struct {
	Argb string `json:"argb"` // ARGB 颜色值，如 #FF000000
}

// GetUserInfosJsonAvaResource 是图层资源。
type GetUserInfosJsonAvaResource struct {
	ResType       int                            `json:"res_type"`        // 资源类型。实测为 3 或 5
	ResImage      *GetUserInfosAvatarImage       `json:"res_image"`       // 图片资源
	ResNativeDraw *GetUserInfosJsonAvaNativeDraw `json:"res_native_draw"` // 原生绘制资源
}

// GetUserInfosAvatarImage 是头像图层的图片资源。
type GetUserInfosAvatarImage struct {
	ImageSrc GetUserInfosAvatarImageSrc `json:"image_src"` // 图片资源
}

// GetUserInfosAvatarImageSrc 是头像图层的图片来源。
type GetUserInfosAvatarImageSrc struct {
	Local       int                       `json:"local,omitempty"`       // 本地占位资源
	Placeholder int                       `json:"placeholder,omitempty"` // 占位类型
	Remote      *GetUserInfosAvatarRemote `json:"remote,omitempty"`      // 远程资源
	SrcType     int                       `json:"src_type"`              // 资源来源类型
}

// GetUserInfosAvatarRemote 是头像渲染层的远程资源。
type GetUserInfosAvatarRemote struct {
	BfsStyle string `json:"bfs_style"` // BFS 样式
	URL      string `json:"url"`       // 资源 URL
}

// GetUserInfosJsonAvaNativeDraw 是原生绘制资源。
type GetUserInfosJsonAvaNativeDraw struct {
	DrawSrc GetUserInfosJsonAvaDrawSrc `json:"draw_src"` // 绘制来源
}

// GetUserInfosJsonAvaDrawSrc 是原生绘制来源。
type GetUserInfosJsonAvaDrawSrc struct {
	SrcType int                     `json:"src_type"` // 来源类型。实测为 3
	Draw    GetUserInfosJsonAvaDraw `json:"draw"`     // 绘制配置。结构与 mask_src 中的 draw 相同
}
