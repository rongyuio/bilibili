package dynamic

// GetDynamicLiveUserList 的响应模型。

// GetDynamicLiveUserListResult 是正在直播的已关注者的 data 对象。
type GetDynamicLiveUserListResult struct {
	Count int                          `json:"count"` // 直播者数量
	Group string                       `json:"group"` // 固定值"default"，作用尚不明确
	Items []GetDynamicLiveUserListItem `json:"items"` // 直播者列表
	Gt    int                          `json:"_gt_"`  // 固定值0，作用尚不明确
}

// GetDynamicLiveUserListItem 是一个正在直播的已关注者。
type GetDynamicLiveUserListItem struct {
	UID   int    `json:"uid"`   // 直播者id
	Uname string `json:"uname"` // 直播者昵称
	Face  string `json:"face"`  // 直播者头像
	Link  string `json:"link"`  // 直播链接
	Title string `json:"title"` // 直播标题
}
