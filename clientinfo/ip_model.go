package clientinfo

import "github.com/rongyuio/bilibili"

// IP 归属地相关响应模型。各接口的返回体独立定义，互不复用。

// GetResourceIPResult 是 APP 端根据请求 IP 确定属地的 data 对象。
type GetResourceIPResult struct {
	Addr        string                  `json:"addr"`         // 公网 IP 地址
	Country     string                  `json:"country"`      // 国家/地区名
	Province    string                  `json:"province"`     // 省/州。非必须存在项
	City        string                  `json:"city"`         // 城市。非必须存在项
	ISP         string                  `json:"isp"`          // 运营商名
	Latitude    bilibili.NumberOrString `json:"latitude"`     // 纬度
	Longitude   bilibili.NumberOrString `json:"longitude"`    // 经度
	ZoneID      int64                   `json:"zone_id"`      // IP 数据库 id
	CountryCode int                     `json:"country_code"` // 国家/地区代码
}

// GetIPInfoNewResult 是查询任意 IP 地址归属地的 data 对象。
type GetIPInfoNewResult struct {
	Addr      string                  `json:"addr"`      // 公网 IP 地址
	Country   string                  `json:"country"`   // 国家/地区名
	Province  string                  `json:"province"`  // 省/州。非必须存在项
	City      string                  `json:"city"`      // 城市。非必须存在项
	ISP       string                  `json:"isp"`       // 运营商名
	Latitude  bilibili.NumberOrString `json:"latitude"`  // 纬度
	Longitude bilibili.NumberOrString `json:"longitude"` // 经度
}

// GetWebRoomIPInfoResult 是直播端根据请求 IP 确定属地的 data 对象。
type GetWebRoomIPInfoResult struct {
	Addr        string                  `json:"addr"`         // 公网 IP 地址
	Country     string                  `json:"country"`      // 国家/地区名
	Province    string                  `json:"province"`     // 省/州。非必须存在项
	City        string                  `json:"city"`         // 城市。非必须存在项
	ISP         string                  `json:"isp"`          // 运营商名
	Latitude    bilibili.NumberOrString `json:"latitude"`     // 纬度
	Longitude   bilibili.NumberOrString `json:"longitude"`    // 经度
	ZoneID      int64                   `json:"zone_id"`      // IP 数据库 id
	CountryCode int                     `json:"country_code"` // 国家/地区代码
}
