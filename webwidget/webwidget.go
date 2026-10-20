package webwidget

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetRegionBannerParam 指定要查询轮播图的分区。
type GetRegionBannerParam struct {
	RegionID int `json:"region_id" request:"query"` // 目标分区 id
}

// GetRegionBannerItem 是一条分区轮播图。
type GetRegionBannerItem struct {
	Image    string `json:"image"`     // 封面资源路径
	Title    string `json:"title"`     // 封面标题
	SubTitle string `json:"sub_title"` // 封面子标题
	URL      string `json:"url"`       // 点击封面后指向的链接。可能为视频地址，也可能是活动地址
	Rid      int    `json:"rid"`       // 分区 id
}

// GetRegionBannerResult 是分区轮播图的 data 对象。
type GetRegionBannerResult struct {
	RegionBannerList []GetRegionBannerItem `json:"region_banner_list"` // 轮播图列表
}

// GetRegionBanner 获取指定分区的轮播图（web 端）。
func (c *Client) GetRegionBanner(ctx context.Context, param GetRegionBannerParam) (*GetRegionBannerResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-show/region/banner"
	)
	return bilibili.Execute[*GetRegionBannerResult](ctx, c.Client, method, url, param)
}

// GetWebHeaderParam 指定首页头图资源。
type GetWebHeaderParam struct {
	ResourceID int `json:"resource_id" request:"query"` // 资源 ID。默认为 142，实测可为任意有效整数
}

// GetWebHeaderResult 是首页头图的 data 对象。
type GetWebHeaderResult struct {
	Name         string `json:"name"`           // 空
	Pic          string `json:"pic"`            // 静态头图 URL
	Litpic       string `json:"litpic"`         // bilibili logo URL
	URL          string `json:"url"`            // 空
	IsSplitLayer int    `json:"is_split_layer"` // 是否分层。1：是
	SplitLayer   string `json:"split_layer"`    // 分层信息。一个套在字符串里的 JSON 对象
	RequestID    string `json:"request_id"`     // 请求 id
}

// GetWebHeader 获取首页头图。
func (c *Client) GetWebHeader(ctx context.Context, param GetWebHeaderParam) (*GetWebHeaderResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-show/page/header"
	)
	return bilibili.Execute[*GetWebHeaderResult](ctx, c.Client, method, url, param)
}
