package live

// 直播间视频流相关响应模型。

// GetLiveStreamURLResult 是直播间视频流接口的 data 对象。
type GetLiveStreamURLResult struct {
	CurrentQuality     int                                  `json:"current_quality"`     // 当前画质代码 qn
	AcceptQuality      []string                             `json:"accept_quality"`      // 可选画质参数（qn 代码列表）
	CurrentQn          int                                  `json:"current_qn"`          // 当前画质代码 quality
	QualityDescription []GetLiveStreamURLQualityDescription `json:"quality_description"` // 可选画质参数（quality 列表）
	Durl               []GetLiveStreamURLDurl               `json:"durl"`                // 直播流 url 组，第一项为主线服务器
}

// GetLiveStreamURLQualityDescription 描述一个 quality 画质代码及其名称。
type GetLiveStreamURLQualityDescription struct {
	Qn   int    `json:"qn"`   // 画质代码
	Desc string `json:"desc"` // 该代码对应的画质名称
}

// GetLiveStreamURLDurl 是直播流 url 组中的一条线路。
type GetLiveStreamURLDurl struct {
	URL        string `json:"url"`         // 直播流 url，flv 或 m3u8 格式。注意带有转义
	Length     int    `json:"length"`      // 0。作用尚不明确
	Order      int    `json:"order"`       // 服务器线路序号
	StreamType int    `json:"stream_type"` // 0。作用尚不明确
	P2pType    int    `json:"p2p_type"`    // 0。作用尚不明确
}
