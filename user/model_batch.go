package user

// 用户批量查询相关响应模型。

// NameToUIDResult 用户名查 mid 结果。
type NameToUIDResult struct {
	UIDList []NameToUIDItem `json:"uid_list"` // UID 列表
}

// NameToUIDItem 用户名与对应的 mid。
type NameToUIDItem struct {
	Name string `json:"name"` // 用户名
	UID  string `json:"uid"`  // mid（UID）
}
