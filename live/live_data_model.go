package live

// 直播数据相关响应模型。

// GetLiveSessionDataResult 是单场直播的数据（data 部分）。
type GetLiveSessionDataResult struct {
	LiveTime     int64 `json:"LiveTime"`     // 该场次的直播时长，单位秒；直播结束后仍会增长
	AddFans      int   `json:"AddFans"`      // 该场次新增粉丝数
	HamsterRmb   int64 `json:"HamsterRmb"`   // 该场次收益
	NewFansClub  int   `json:"NewFansClub"`  // 该场次新获得的粉丝勋章数量
	DanmuNum     int   `json:"DanmuNum"`     // 该场次弹幕条数
	MaxOnline    int   `json:"MaxOnline"`    // 该场次最大在线人数（待验证）
	WatchedCount int   `json:"WatchedCount"` // 该场次看过人数
}

// GetLivePerformanceResult 是直播表现（data 部分），用于渲染雷达图。
type GetLivePerformanceResult struct {
	Graph   []GetLivePerformanceGraph `json:"graph"`   // 图表数据，每项是一个数据类别
	Propose any                       `json:"propose"` // (?)，通常为 null
}

// GetLivePerformanceGraph 是雷达图中的一个数据类别。
type GetLivePerformanceGraph struct {
	Name  string  `json:"name"`  // 显示的名称
	Index string  `json:"index"` // 类别标识，例如 income
	Me    float64 `json:"me"`    // 自己在该类别的数值
	Max   float64 `json:"max"`   // 该类别的最大值（雷达图最外侧）
	Aver  float64 `json:"aver"`  // 同水平主播在该类别的数值
}
