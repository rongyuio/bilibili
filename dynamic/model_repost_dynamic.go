package dynamic

// RepostDynamic 的响应模型。

// RepostDynamicResult 是发表动态接口（含转发）的响应。
//
// 注意它与 CreateDynamicResult 不是同一套字段：后者属于旧的 dynamic_svr/create 接口。
type RepostDynamicResult struct {
	DynID    int64  `json:"dyn_id"`     // 新动态 id
	DynIDStr string `json:"dyn_id_str"` // 新动态 id，字符串格式
	DynType  int    `json:"dyn_type"`   // 动态类型：1=转发、2=图文
	DynRid   int64  `json:"dyn_rid"`    // 动态业务 id
	Result   int    `json:"result"`     // 0
	Message  string `json:"message"`    // 错误信息
}
