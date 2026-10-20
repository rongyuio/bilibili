package dynamic

// SearchDynamicAt 的响应模型。

// SearchDynamicAtResult 是根据关键字搜索用户的 data 对象。
type SearchDynamicAtResult struct {
	Groups []SearchDynamicAtGroup `json:"groups"` // 内容分组
	Gt     int                    `json:"_gt_"`   // 固定值0
}

// SearchDynamicAtGroup 是一个用户分组。
type SearchDynamicAtGroup struct {
	GroupType int                        `json:"group_type"` // 2:我的关注。4:其他
	GroupName string                     `json:"group_name"` // 分组名字
	Items     []SearchDynamicAtGroupItem `json:"items"`      // 用户信息
}

// SearchDynamicAtGroupItem 是分组中的一位用户。
type SearchDynamicAtGroupItem struct {
	UID                int    `json:"uid"`                  // 用户id
	Uname              string `json:"uname"`                // 用户昵称
	Face               string `json:"face"`                 // 用户头像url
	Fans               int    `json:"fans"`                 // 用户粉丝数
	OfficialVerifyType int    `json:"official_verify_type"` // 认证信息?
}
