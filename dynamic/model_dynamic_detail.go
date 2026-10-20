package dynamic

// GetDynamicDetail 的响应模型。

// GetDynamicDetailResult 是特定动态卡片信息的 data 对象。
type GetDynamicDetailResult struct {
	Card   *GetDynamicDetailCard `json:"card"` // 动态卡片内容
	Result int                   `json:"result"`
	Gt     int                   `json:"_gt_"`
}

// GetDynamicDetailCard 动态卡片内容。因为 ActivityInfos 、 Desc 、 Display 等字段会随着此动态类型不同发生一定的变化，无法统一，因此都转换成了 map[string]any ，请自行解析
type GetDynamicDetailCard struct {
	ActivityInfos map[string]any `json:"activity_infos"` // 该条动态参与的活动
	Card          string         `json:"card"`           // 动态详细信息
	Desc          map[string]any `json:"desc"`           // 动态相关信息
	Display       map[string]any `json:"display"`        // 动态部分的可操作项
	ExtendJSON    string         `json:"extend_json"`    // 动态扩展项
}
