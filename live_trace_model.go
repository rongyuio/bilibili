package bilibili

// Web 端直播数据上报相关响应模型（live-trace 域）。

// LiveHeartBeatResult 是进房（EnterLiveRoom）与心跳（SendLiveHeartBeat）接口共用的响应。
// 调用方应把 SecretKey、SecretRule、Timestamp 原样带入下一包心跳。
type LiveHeartBeatResult struct {
	HeartbeatInterval int    `json:"heartbeat_interval"` // 下一次心跳的间隔（秒）
	SecretKey         string `json:"secret_key"`         // 下一包心跳的签名密钥
	SecretRule        []int  `json:"secret_rule"`        // 下一包心跳的签名哈希规则
	Timestamp         int64  `json:"timestamp"`          // 服务器时间戳（毫秒），作为下一包的 Ets
}
