package dynamic

// 动态草稿相关响应模型。
//
// 本包不复用根包或其他子包的响应体：结果与所有嵌套对象都以接口方法名
// 为前缀独立定义，便于后续单个接口的字段变更不影响其他调用。

// GetDynamicDraftListResult 是当前账号的动态草稿列表。
type GetDynamicDraftListResult struct {
	Drafts []GetDynamicDraftListItem `json:"drafts"` // 草稿列表
}

// GetDynamicDraftListItem 是一条动态草稿。
type GetDynamicDraftListItem struct {
	DraftID     string                         `json:"draft_id"`     // 草稿id
	PublishTime int64                          `json:"publish_time"` // 定时发送的秒级时间戳
	Type        int                            `json:"type"`         // 动态类型，参考获取动态详情
	UID         int64                          `json:"uid"`          // 自己的mid
	UserProfile GetDynamicDraftListUserProfile `json:"user_profile"` // 自己的用户信息
	Request     string                         `json:"request"`      // 动态内容，该项为 JSON 转成的字符串，内容参考发表纯文本动态
}

// GetDynamicDraftListUserProfile 是草稿所属用户的用户信息。
type GetDynamicDraftListUserProfile struct {
	Info      GetDynamicDraftListUserInfo      `json:"info"`       // 基础信息
	Card      GetDynamicDraftListCard          `json:"card"`       // 卡片信息
	Vip       GetDynamicDraftListUserVip       `json:"vip"`        // 大会员信息
	Pendant   GetDynamicDraftListUserPendant   `json:"pendant"`    // 挂件信息
	Rank      string                           `json:"rank"`       // 排行
	Sign      string                           `json:"sign"`       // 签名
	LevelInfo GetDynamicDraftListUserLevelInfo `json:"level_info"` // 等级信息
}

// GetDynamicDraftListUserInfo 是草稿所属用户的基础信息。
type GetDynamicDraftListUserInfo struct {
	UID     int    `json:"uid"`      // 用户 mid
	Uname   string `json:"uname"`    // 用户名
	Face    string `json:"face"`     // 头像 URL
	FaceNft int    `json:"face_nft"` // 是否为 NFT 头像
}

// GetDynamicDraftListCard 是草稿所属用户的卡片信息。
type GetDynamicDraftListCard struct {
	OfficialVerify GetDynamicDraftListOfficialVerify `json:"official_verify"` // 认证信息
}

// GetDynamicDraftListOfficialVerify 是草稿所属用户的认证信息。
type GetDynamicDraftListOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetDynamicDraftListUserVip 是草稿所属用户的大会员信息。
type GetDynamicDraftListUserVip struct {
	VipType            int                         `json:"vipType"`              // 大会员类型
	VipDueDate         int64                       `json:"vipDueDate"`           // 大会员到期时间，毫秒时间戳
	VipStatus          int                         `json:"vipStatus"`            // 大会员状态
	ThemeType          int                         `json:"themeType"`            // 会员样式 id
	Label              GetDynamicDraftListVipLabel `json:"label"`                // 会员铭牌样式
	AvatarSubscript    int                         `json:"avatar_subscript"`     // 是否显示会员图标
	NicknameColor      string                      `json:"nickname_color"`       // 会员昵称颜色
	Role               int                         `json:"role"`                 // 大角色类型
	AvatarSubscriptURL string                      `json:"avatar_subscript_url"` // 大会员角标地址
}

// GetDynamicDraftListVipLabel 是草稿所属用户的会员铭牌（7 字段）。
type GetDynamicDraftListVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型。vip：大会员。annual_vip：年度大会员。ten_annual_vip：十年大会员。hundred_annual_vip：百年大会员
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetDynamicDraftListUserPendant 是草稿所属用户的挂件信息。
type GetDynamicDraftListUserPendant struct {
	Pid               int    `json:"pid"`                 // 挂件 id
	Name              string `json:"name"`                // 挂件名称
	Image             string `json:"image"`               // 挂件图片 URL
	Expire            int    `json:"expire"`              // 挂件到期时间
	ImageEnhance      string `json:"image_enhance"`       // 增强挂件图片 URL
	ImageEnhanceFrame string `json:"image_enhance_frame"` // 增强挂件边框 URL
}

// GetDynamicDraftListUserLevelInfo 是草稿所属用户的等级信息。
type GetDynamicDraftListUserLevelInfo struct {
	CurrentLevel int `json:"current_level"` // 当前等级
}
