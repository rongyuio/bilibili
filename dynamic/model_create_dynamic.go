package dynamic

// CreateDynamic 的响应模型。

// CreateDynamicResult 是发表纯文本动态接口的 data 对象。
type CreateDynamicResult struct {
	Result       int    `json:"result"`         // 0
	Errmsg       string `json:"errmsg"`         // 像是服务器日志一样的东西
	DynamicID    int    `json:"dynamic_id"`     // 动态 id
	CreateResult int    `json:"create_result"`  // 1
	DynamicIDStr string `json:"dynamic_id_str"` // 动态 id。字符串格式
	Gt           int    `json:"_gt_"`           // 0
}
