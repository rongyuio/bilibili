package dynamic

// GetDynamicLikeList 的响应模型。

// GetDynamicLikeListResult 是动态点赞列表的 data 对象。
type GetDynamicLikeListResult struct {
	ItemLikes  []GetDynamicLikeListItem `json:"item_likes"`  // 点赞信息列表主体
	HasMore    int                      `json:"has_more"`    // 是否还有下一页
	TotalCount int                      `json:"total_count"` // 总计点赞数
	Gt         int                      `json:"_gt_"`        // 固定值0
}

// GetDynamicLikeListItem 是点赞列表中的一条点赞记录。
type GetDynamicLikeListItem struct {
	UID      int                            `json:"uid"`
	Time     int                            `json:"time"`
	FaceURL  string                         `json:"face_url"`
	Uname    string                         `json:"uname"`
	UserInfo GetDynamicLikeListItemUserInfo `json:"user_info"`
	Attend   int                            `json:"attend"`
}

// GetDynamicLikeListItemUserInfo 是点赞用户的详细信息。
type GetDynamicLikeListItemUserInfo struct {
	UID            int                                  `json:"uid"`
	Uname          string                               `json:"uname"`
	Face           string                               `json:"face"`
	Rank           string                               `json:"rank"`
	OfficialVerify GetDynamicLikeListItemOfficialVerify `json:"official_verify"`
	Vip            GetDynamicLikeListItemVip            `json:"vip"`
	Pendant        GetDynamicLikeListItemPendant        `json:"pendant"`
	Sign           string                               `json:"sign"`
	LevelInfo      GetDynamicLikeListItemLevelInfo      `json:"level_info"`
}

// GetDynamicLikeListItemOfficialVerify 是点赞用户的认证标识。
type GetDynamicLikeListItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetDynamicLikeListItemVip 是点赞用户的会员信息（camelCase 变体）。
type GetDynamicLikeListItemVip struct {
	VipType            int                            `json:"vipType"`
	VipDueDate         int64                          `json:"vipDueDate"`
	VipStatus          int                            `json:"vipStatus"`
	ThemeType          int                            `json:"themeType"`
	Label              GetDynamicLikeListItemVipLabel `json:"label"`
	AvatarSubscript    int                            `json:"avatar_subscript"`
	NicknameColor      string                         `json:"nickname_color"`
	Role               int                            `json:"role"`
	AvatarSubscriptURL string                         `json:"avatar_subscript_url"`
}

// GetDynamicLikeListItemVipLabel 是点赞用户的会员铭牌。
type GetDynamicLikeListItemVipLabel struct {
	Path        string `json:"path"`         // (?)
	Text        string `json:"text"`         // 会员类型文案
	LabelTheme  string `json:"label_theme"`  // 会员类型
	TextColor   string `json:"text_color"`   // 文字颜色?
	BgStyle     int    `json:"bg_style"`     // (?)
	BgColor     string `json:"bg_color"`     // 背景颜色?
	BorderColor string `json:"border_color"` // 描边颜色?
}

// GetDynamicLikeListItemPendant 是点赞用户的挂件信息。
type GetDynamicLikeListItemPendant struct {
	Pid               int    `json:"pid"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	Expire            int    `json:"expire"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
}

// GetDynamicLikeListItemLevelInfo 是点赞用户的等级信息。
type GetDynamicLikeListItemLevelInfo struct {
	CurrentLevel int `json:"current_level"`
}
