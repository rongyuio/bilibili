package video

// 视频在线人数（APP 端）响应模型。

// GetVideoOnlineInfoAppResult 是视频在线人数（APP 端）的 data 对象。
type GetVideoOnlineInfoAppResult struct {
	Online     GetVideoOnlineInfoAppOnline      `json:"online"`      // 所有终端总计人数
	LikeSwitch bool                             `json:"like_switch"` // 是否显示点赞按钮。实测为 true
	ShowSwitch *GetVideoOnlineInfoAppShowSwitch `json:"show_switch"` // 数据显示控制。文档单独列出，响应示例中未出现，可能为 null
}

// GetVideoOnlineInfoAppOnline 是所有终端总计人数。
type GetVideoOnlineInfoAppOnline struct {
	TotalText string `json:"total_text"` // 所有终端总计人数文本。文档备注实测为空，响应示例为 "8.8万+人在看"
}

// GetVideoOnlineInfoAppShowSwitch 是数据显示控制。
type GetVideoOnlineInfoAppShowSwitch struct {
	Online string `json:"online"` // 所有终端总计人数。例如 "10万+人在看"
}
