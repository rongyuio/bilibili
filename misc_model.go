package bilibili

// 杂项接口（分区数据、日报）响应模型。

type ZoneLocation struct {
	Addr        string `json:"addr"`         // 公网IP地址
	Country     string `json:"country"`      // 国家/地区名
	Province    string `json:"province"`     // 省/州。非必须存在项
	City        string `json:"city"`         // 城市。非必须存在项
	Isp         string `json:"isp"`          // 运营商名
	Latitude    int    `json:"latitude"`     // 纬度
	Longitude   int    `json:"longitude"`    // 经度
	ZoneID      int    `json:"zone_id"`      // ip数据库id
	CountryCode int    `json:"country_code"` // 国家/地区代码
}

type RegionDailyCount struct {
	RegionCount map[int]int `json:"region_count"` // 分区当日投稿稿件数信息
}
