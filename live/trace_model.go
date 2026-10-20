package live

// Web 端直播数据上报相关响应模型（live-trace 域）。
// 各接口的返回体独立定义，不与其他接口复用。

// EnterLiveRoomResult 是进房上报（EnterLiveRoom）的 data 对象。
// 调用方应把 SecretKey、SecretRule、Timestamp 原样带入下一包心跳。
type EnterLiveRoomResult struct {
	HeartbeatInterval int    `json:"heartbeat_interval"` // 下一次心跳的间隔（秒）
	SecretKey         string `json:"secret_key"`         // 下一包心跳的签名密钥
	SecretRule        []int  `json:"secret_rule"`        // 下一包心跳的签名哈希规则
	Timestamp         int64  `json:"timestamp"`          // 服务器时间戳（毫秒），作为下一包的 Ets
}

// SendLiveHeartBeatResult 是心跳上报（SendLiveHeartBeat）的 data 对象。
// 调用方应把 SecretKey、SecretRule、Timestamp 原样带入下一包心跳。
type SendLiveHeartBeatResult struct {
	HeartbeatInterval int    `json:"heartbeat_interval"` // 下一次心跳的间隔（秒）
	SecretKey         string `json:"secret_key"`         // 下一包心跳的签名密钥
	SecretRule        []int  `json:"secret_rule"`        // 下一包心跳的签名哈希规则
	Timestamp         int64  `json:"timestamp"`          // 服务器时间戳（毫秒），作为下一包的 Ets
}

// ReportWebWatchEnterResult 是观看时长新链路进房（ReportWebWatchEnter）的 data 对象。
type ReportWebWatchEnterResult struct {
	STKY string `json:"stky"` // 下一次心跳使用的心跳密钥（每次轮换）
	SID  string `json:"sid"`  // 会话 id，进房后原样携带
	HBIL int    `json:"hbil"` // 心跳间隔（秒）
}

// ReportWebWatchHeartBeatResult 是观看时长新链路心跳（ReportWebWatchHeartBeat）的 data 对象。
type ReportWebWatchHeartBeatResult struct {
	STKY string `json:"stky"` // 下一次心跳使用的心跳密钥（每次轮换）
	SID  string `json:"sid"`  // 会话 id，进房后原样携带
	HBIL int    `json:"hbil"` // 心跳间隔（秒）
}
