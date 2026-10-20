package videoranking

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 分区视频排行相关接口。响应模型见 model_ranking.go。

// GetZoneVideoRankListParam 是获取分区视频排行榜列表的请求参数。
type GetZoneVideoRankListParam struct {
	Tid  int    `json:"tid,omitempty" request:"query,omitempty"`  // 目标分区tid，可不填。可调用 GetAllZoneInfos 获取，或者直接自行查阅 video_zone.csv
	Type string `json:"type,omitempty" request:"query,omitempty"` // 未知。默认为：all，且为目前唯一已知值。怀疑为稿件类型，但没有找到其他值佐证。
}

// GetZoneVideoRankList 获取分区视频排行榜列表
func (c *Client) GetZoneVideoRankList(ctx context.Context, param GetZoneVideoRankListParam) (*GetZoneVideoRankListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/ranking/v2"
	)
	return bilibili.Execute[*GetZoneVideoRankListResult](ctx, c.Client, method, url, param)
}

// GetZoneVideoListNewParam 是获取分区最新视频列表的请求参数。
type GetZoneVideoListNewParam struct {
	Pn  int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
	Ps  int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为14, 留空为5
	Rid int `json:"rid"`                                    // 目标分区tid
}

// GetZoneVideoListNew 获取分区最新视频列表
func (c *Client) GetZoneVideoListNew(ctx context.Context, param GetZoneVideoListNewParam) (*GetZoneVideoListNewResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/dynamic/region"
	)
	return bilibili.Execute[*GetZoneVideoListNewResult](ctx, c.Client, method, url, param)
}

// GetZoneVideoListWithTagParam 是获取分区标签近期互动列表的请求参数。
type GetZoneVideoListWithTagParam struct {
	Ps    int `json:"ps,omitempty" request:"query,omitempty"` // 视频数。默认为14, 留空为5
	Pn    int `json:"pn,omitempty" request:"query,omitempty"` // 列数。留空为1
	Rid   int `json:"rid"`                                    // 目标分区id。参见仓库根目录的 video_zone.csv
	TagID int `json:"tag_id"`                                 // 目标标签id
}

// GetZoneVideoListWithTag 获取分区标签近期互动列表
func (c *Client) GetZoneVideoListWithTag(ctx context.Context, param GetZoneVideoListWithTagParam) (*GetZoneVideoListWithTagResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/dynamic/tag"
	)
	return bilibili.Execute[*GetZoneVideoListWithTagResult](ctx, c.Client, method, url, param)
}

// GetZoneVideoListRecentParam 是获取分区近期投稿列表的请求参数。
type GetZoneVideoListRecentParam struct {
	Ps   int `json:"ps,omitempty" request:"query,omitempty"`   // 视频数。默认为14, 留空为5
	Pn   int `json:"pn,omitempty" request:"query,omitempty"`   // 页码。默认为1
	Rid  int `json:"rid,omitempty" request:"query,omitempty"`  // 目标分区id。参见仓库根目录的 video_zone.csv
	Type int `json:"type,omitempty" request:"query,omitempty"` // 类型?。默认为0
}

// GetZoneVideoListRecent 获取分区近期投稿列表
func (c *Client) GetZoneVideoListRecent(ctx context.Context, param GetZoneVideoListRecentParam) (*GetZoneVideoListRecentResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/newlist"
	)
	return bilibili.Execute[*GetZoneVideoListRecentResult](ctx, c.Client, method, url, param)
}

// GetZoneVideoListByOrderParam 是获取分区近期投稿列表（带排序）的请求参数。
type GetZoneVideoListByOrderParam struct {
	MainVer    string `json:"main_ver,omitempty" request:"query,omitempty"`    // 主页版本。默认为 v3
	SearchType string `json:"search_type"`                                     // 搜索类型。默认为 video
	ViewType   string `json:"view_type"`                                       // 查看类型?。默认为 hot_rank
	CopyRight  int    `json:"copy_right,omitempty" request:"query,omitempty"`  // 版权?。默认为 -1
	NewWebTag  int    `json:"new_web_tag,omitempty" request:"query,omitempty"` // 标签?。默认为 1
	Order      string `json:"order,omitempty" request:"query,omitempty"`       // 排序方式。click: 按播放排序(默认)。scores: 按评论数排序。stow: 按收藏排序。coin: 按硬币数排序。dm: 按弹幕数排序
	CateID     int    `json:"cate_id"`                                         // 分区id。留空会导致响应中data中result为null, 参见仓库根目录的 video_zone.csv
	Page       int    `json:"page,omitempty" request:"query,omitempty"`        // 页码。默认以 1 开始
	PageSize   int    `json:"pagesize"`                                        // 视频数。默认为 30, 留空会导致 -500
	TimeFrom   int    `json:"time_from"`                                       // 起始时间。yyyyMMdd, 默认为 time_to - 7
	TimeTo     int    `json:"time_to"`                                         // 结束时间。yyyyMMdd, 默认为当前时间(大于起始时间)
}

// GetZoneVideoListByOrder 获取分区近期投稿列表 (带排序)
func (c *Client) GetZoneVideoListByOrder(ctx context.Context, param GetZoneVideoListByOrderParam) (*GetZoneVideoListByOrderResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/newlist_rank"
	)
	return bilibili.Execute[*GetZoneVideoListByOrderResult](ctx, c.Client, method, url, param)
}
