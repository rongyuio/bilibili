package dynamic

// 导航栏动态响应模型。
//
// 本包不复用根包或其他子包的响应体：结果与所有嵌套对象都以接口方法名
// 为前缀独立定义，便于后续单个接口的字段变更不影响其他调用。

// GetDynamicNavListResult 是导航栏动态列表。
type GetDynamicNavListResult struct {
	HasMore        bool                    `json:"has_more"`        // 是否有更多数据
	Items          []GetDynamicNavListItem `json:"items"`           // 数据数组
	Offset         string                  `json:"offset"`          // 偏移量。等于 items 中最后一条记录的 id，获取下一页时使用
	UpdateBaseline string                  `json:"update_baseline"` // 更新基线。等于 items 中第一条记录的 id_str
	UpdateNum      int                     `json:"update_num"`      // 本次获取到了多少条新动态（更新基线以上的动态条数）
}

// GetDynamicNavListItem 是导航栏中的一条动态。
type GetDynamicNavListItem struct {
	Author    GetDynamicNavListAuthor `json:"author"`     // UP主
	Cover     string                  `json:"cover"`      // 封面URL
	IDStr     string                  `json:"id_str"`     // 动态id
	JumpURL   string                  `json:"jump_url"`   // 跳转地址
	PubTime   string                  `json:"pub_time"`   // 发布时间。文字表述的相对时间
	Rid       int64                   `json:"rid"`        // 关联id。视频即 aid
	Title     string                  `json:"title"`      // 标题
	Type      int                     `json:"type"`       // 动态类型。8：视频
	Visible   bool                    `json:"visible"`    // 是否可见
	IconBadge any                     `json:"icon_badge"` // 实测为 null
	Fold      *GetDynamicNavListFold  `json:"fold"`       // 折叠信息。实测多数为 null
}

// GetDynamicNavListAuthor 是导航栏动态的 UP 主。
type GetDynamicNavListAuthor struct {
	Face     string                    `json:"face"`     // 头像URL
	JumpURL  string                    `json:"jump_url"` // 空间动态页地址
	Mid      int64                     `json:"mid"`      // UP主 mid
	Name     string                    `json:"name"`     // UP主昵称
	Official GetDynamicNavListOfficial `json:"official"` // 认证信息
	Vip      GetDynamicNavListVip      `json:"vip"`      // 大会员信息
}

// GetDynamicNavListOfficial 是 UP 主的认证信息。
type GetDynamicNavListOfficial struct {
	Desc  string `json:"desc"`  // 认证备注。无为空
	Role  int    `json:"role"`  // 认证级别
	Title string `json:"title"` // 认证名称
	Type  int    `json:"type"`  // 认证类型
}

// GetDynamicNavListVip 是 UP 主的大会员信息。
type GetDynamicNavListVip struct {
	AvatarIcon         GetDynamicNavListAvatarIcon `json:"avatar_icon"`          // 头像挂件
	AvatarSubscript    int                         `json:"avatar_subscript"`     // 是否显示会员铭牌
	AvatarSubscriptURL string                      `json:"avatar_subscript_url"` // 会员铭牌图片URL
	DueDate            int64                       `json:"due_date"`             // 大会员到期时间。毫秒时间戳
	Label              GetDynamicNavListVipLabel   `json:"label"`                // 会员铭牌样式
	NicknameColor      string                      `json:"nickname_color"`       // 昵称颜色
	Role               int                         `json:"role"`                 // 会员角色
	Status             int                         `json:"status"`               // 会员状态
	ThemeType          int                         `json:"theme_type"`           // 会员样式 id
	TvDueDate          int64                       `json:"tv_due_date"`          // 电视大会员到期时间。毫秒时间戳
	TvVipPayType       int                         `json:"tv_vip_pay_type"`      // 电视大会员付费类型
	TvVipStatus        int                         `json:"tv_vip_status"`        // 电视大会员状态
	Type               int                         `json:"type"`                 // 会员类型
	VipPayType         int                         `json:"vip_pay_type"`         // 大会员付费类型
}

// GetDynamicNavListAvatarIcon 是会员头像挂件。
type GetDynamicNavListAvatarIcon struct {
	IconResource any `json:"icon_resource"` // 图标资源。示例为空对象，结构未公开
	IconType     int `json:"icon_type"`     // 图标类型
}

// GetDynamicNavListVipLabel 是会员铭牌样式（12 字段，含图片铭牌）。
type GetDynamicNavListVipLabel struct {
	BgColor               string `json:"bg_color"`                  // 背景颜色
	BgStyle               int    `json:"bg_style"`                  // 背景样式
	BorderColor           string `json:"border_color"`              // 描边颜色
	ImgLabelURIHans       string `json:"img_label_uri_hans"`        // 简体图片铭牌URL
	ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"` // 简体图片铭牌静态URL
	ImgLabelURIHant       string `json:"img_label_uri_hant"`        // 繁体图片铭牌URL
	ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"` // 繁体图片铭牌静态URL
	LabelTheme            string `json:"label_theme"`               // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员
	Path                  string `json:"path"`                      // (?)
	Text                  string `json:"text"`                      // 会员类型文案
	TextColor             string `json:"text_color"`                // 文字颜色
	UseImgLabel           bool   `json:"use_img_label"`             // 是否使用图片铭牌
}

// GetDynamicNavListFold 是多条同类动态被折叠时的说明。
type GetDynamicNavListFold struct {
	IDs       []string `json:"ids"`       // 被折叠的动态id
	Statement string   `json:"statement"` // 折叠说明
	Type      int      `json:"type"`      // 折叠类型
	Users     []any    `json:"users"`     // 折叠用户。示例为空数组
}
