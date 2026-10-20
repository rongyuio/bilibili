package dynamic

// GetDynamicUpList 的响应模型。

// GetDynamicUpListResult 是发布新动态的已关注者的 data 对象。
type GetDynamicUpListResult struct {
	ButtonStatement string                 `json:"button_statement"` // 固定值空，作用尚不明确
	Items           []GetDynamicUpListItem `json:"items"`            // 更新者列表
	Gt              int                    `json:"_gt_"`             // 固定值0，作用尚不明确
}

// GetDynamicUpListItem 是一条发布新动态的已关注者记录。
type GetDynamicUpListItem struct {
	UserProfile GetDynamicUpListItemUserProfile `json:"user_profile"`
	HasUpdate   int                             `json:"has_update"`
}

// GetDynamicUpListItemUserProfile 是更新者的用户信息。
type GetDynamicUpListItemUserProfile struct {
	Info      GetDynamicUpListItemUserInfo  `json:"info"`
	Card      GetDynamicUpListItemCard      `json:"card"`
	Vip       GetDynamicUpListItemVip       `json:"vip"`
	Pendant   GetDynamicUpListItemPendant   `json:"pendant"`
	Rank      string                        `json:"rank"`
	Sign      string                        `json:"sign"`
	LevelInfo GetDynamicUpListItemLevelInfo `json:"level_info"`
}

// GetDynamicUpListItemUserInfo 是更新者的基础信息。
type GetDynamicUpListItemUserInfo struct {
	UID   int    `json:"uid"`
	Uname string `json:"uname"`
	Face  string `json:"face"`
}

// GetDynamicUpListItemCard 是更新者的卡片信息。
type GetDynamicUpListItemCard struct {
	OfficialVerify GetDynamicUpListItemOfficialVerify `json:"official_verify"`
}

// GetDynamicUpListItemOfficialVerify 是更新者的认证标识。
type GetDynamicUpListItemOfficialVerify struct {
	Type int    `json:"type"` // 是否认证，-1：无。0：个人认证。1：机构认证
	Desc string `json:"desc"` // 认证信息，无为空
}

// GetDynamicUpListItemVip 是更新者的会员信息，其标签仅有 path 字段。
type GetDynamicUpListItemVip struct {
	VipType       int                          `json:"vipType"`
	VipDueDate    int64                        `json:"vipDueDate"`
	DueRemark     string                       `json:"dueRemark"`
	AccessStatus  int                          `json:"accessStatus"`
	VipStatus     int                          `json:"vipStatus"`
	VipStatusWarn string                       `json:"vipStatusWarn"`
	ThemeType     int                          `json:"themeType"`
	Label         GetDynamicUpListItemVipLabel `json:"label"`
}

// GetDynamicUpListItemVipLabel 是更新者的会员标签（仅含 path）。
type GetDynamicUpListItemVipLabel struct {
	Path string `json:"path"`
}

// GetDynamicUpListItemPendant 是更新者的挂件信息。
type GetDynamicUpListItemPendant struct {
	Pid          int    `json:"pid"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Expire       int    `json:"expire"`
	ImageEnhance string `json:"image_enhance"`
}

// GetDynamicUpListItemLevelInfo 是更新者的等级信息。
type GetDynamicUpListItemLevelInfo struct {
	CurrentLevel int    `json:"current_level"`
	CurrentMin   int    `json:"current_min"`
	CurrentExp   int    `json:"current_exp"`
	NextExp      string `json:"next_exp"`
}
