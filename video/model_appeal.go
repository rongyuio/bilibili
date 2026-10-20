package video

// 稿件投诉相关响应模型。

// GetAppealTagsItem 是一种稿件投诉类型。
type GetAppealTagsItem struct {
	Tid      int                    `json:"tid"`      // 类型 tid
	Business int                    `json:"business"` // 意义不明
	Weight   int                    `json:"weight"`   // 权重
	Round    int                    `json:"round"`    // 意义不明
	State    int                    `json:"state"`    // 意义不明
	Name     string                 `json:"name"`     // 类型名称
	Remark   string                 `json:"remark"`   // 类型备注
	Ctime    string                 `json:"ctime"`    // 意义不明
	Mtime    string                 `json:"mtime"`    // 意义不明
	Controls []GetAppealTagsControl `json:"controls"` // 详细信息填写提示。没有时为 null
}

// GetAppealTagsControl 是投诉类型的填写提示。
type GetAppealTagsControl struct {
	Tid         int    `json:"tid"`         // 类型 tid
	Bid         int    `json:"bid"`         // 意义不明
	Name        string `json:"name"`        // 提示名称
	Title       string `json:"title"`       // 提示标题
	Component   string `json:"component"`   // 需要填入的类型
	Placeholder string `json:"placeholder"` // 文本框占位符
	Required    int    `json:"required"`    // 是否为必填
}
