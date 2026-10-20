package dynamic

// 动态 @ 用户搜索（dynamic/atlist）相关响应模型。
//
// 本包不复用根包或其他子包的响应体：结果与所有嵌套对象都以接口方法名
// 为前缀独立定义，便于后续单个接口的字段变更不影响其他调用。

// SearchDynamicMentionUserResult 是 @ 用户搜索接口的 data 本体。
type SearchDynamicMentionUserResult struct {
	Groups []SearchDynamicMentionUserGroup `json:"groups"` // 用户分组，按是否关注划分；未找到用户时为空数组
}

// SearchDynamicMentionUserGroup 是搜索结果的用户分组。
type SearchDynamicMentionUserGroup struct {
	GroupName string                         `json:"group_name"` // 分组名称：我的关注 或 其他
	GroupType int                            `json:"group_type"` // 分组类型：2：我的关注；4：其他
	Items     []SearchDynamicMentionUserItem `json:"items"`      // 搜索到的用户，按认证状态与粉丝数排序
}

// SearchDynamicMentionUserItem 是搜索到的用户信息。
type SearchDynamicMentionUserItem struct {
	Face               string `json:"face"`                 // 用户头像 url
	Fans               int64  `json:"fans"`                 // 用户粉丝数
	Name               string `json:"name"`                 // 用户昵称
	OfficialVerifyType int    `json:"official_verify_type"` // 用户认证状态：-1：无；0：个人认证；1：机构认证
	UID                string `json:"uid"`                  // 用户 mid
}
