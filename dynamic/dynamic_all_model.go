package dynamic

// 关注动态流（dynamic/all）相关响应模型。

// CheckDynamicUpdateResult 是检测新动态接口的 data 本体。
type CheckDynamicUpdateResult struct {
	UpdateNum int `json:"update_num"` // 有新动态的数量
}
