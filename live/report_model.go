package live

// Web 端直播心跳上报的响应模型（live-report 域）。

// ReportLiveWebHeartBeatResult 是 Web 端直播心跳上报接口的响应 data。
type ReportLiveWebHeartBeatResult struct {
	NextInterval int `json:"next_interval"` // 下次心跳间隔（秒）
}
