package user

// 用户粉丝明细与关注分组的响应模型（续）。
//
// relation_model.go 已定义 GetUserFans*、GetUserFollowers3*、GetUserFollowersUnreadCount*、
// GetRelationTagUsersItem 与 CreateRelationTagResult；本文件只补齐其余接口的返回体。
// 各接口的返回体独立定义，即使字段相同也不与其它接口共用类型。

// GetRelationTagUserResult 是查询目标用户所在分组的 data 对象。
// 键为分组 id（字符串形式），值为分组名称；目标用户位于默认分组时不会出现在此对象中。
type GetRelationTagUserResult map[string]string

// GetRelationTagSpecialResult 是查询所有特别关注 mid 的 data 数组。
type GetRelationTagSpecialResult []int
