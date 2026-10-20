package live

// GetDanmuInfoResult 是获取信息流认证秘钥的响应。
type GetDanmuInfoResult struct {
	Group            string             `json:"group"`              // 分组，固定为 live
	BusinessID       int                `json:"business_id"`        // 业务 id，示例为 0
	RefreshRowFactor float64            `json:"refresh_row_factor"` // 刷新行因子，示例为 0.125
	RefreshRate      int                `json:"refresh_rate"`       // 刷新频率，示例为 100
	MaxDelay         int                `json:"max_delay"`          // 最大延迟，单位毫秒，示例为 5000
	Token            string             `json:"token"`              // 认证秘钥
	HostList         []GetDanmuInfoHost `json:"host_list"`          // 信息流服务器节点列表
}

// GetDanmuInfoHost 是信息流服务器节点。
type GetDanmuInfoHost struct {
	Host    string `json:"host"`     // 服务器域名
	Port    int    `json:"port"`     // TCP 端口
	WSSPort int    `json:"wss_port"` // WSS 端口
	WSPort  int    `json:"ws_port"`  // WS 端口
}
