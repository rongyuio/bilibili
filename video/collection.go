package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频合集、视频列表与关键词检索接口。响应模型见 model_collection.go。

// GetVideoCollectionInfoParam 指定要查询的视频合集。
type GetVideoCollectionInfoParam struct {
	Mid         int  `json:"mid"`                                              // UP 主 ID
	SeasonID    int  `json:"season_id"`                                        // 视频合集 ID
	SortReverse bool `json:"sort_reverse,omitempty" request:"query,omitempty"` // 未知
	PageNum     int  `json:"page_num,omitempty" request:"query,omitempty"`     // 页码索引
	PageSize    int  `json:"page_size,omitempty" request:"query,omitempty"`    // 单页内容数量
}

// GetVideoCollectionInfo 获取视频合集信息
func (c *Client) GetVideoCollectionInfo(ctx context.Context, param GetVideoCollectionInfoParam) (*GetVideoCollectionInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-space/seasons_archives_list"
	)
	return bilibili.Execute[*GetVideoCollectionInfoResult](ctx, c.Client, method, url, param)
}

// GetVideoByKeywordsParam 指定关键词检索条件。
type GetVideoByKeywordsParam struct {
	Mid      int    `json:"mid"`                                           // 用户 mid
	Keywords string `json:"keywords"`                                      // 关键词。可为空, 即获取所有视频
	Ps       int    `json:"ps,omitempty" request:"query,omitempty"`        // 每页视频数。默认为 0, 留空为 20
	Pn       int    `json:"pn,omitempty" request:"query,omitempty"`        // 页码。留空为 1
	Orderby  string `json:"orderby,omitempty" request:"query,omitempty"`   // 排序方式。最新发布: pubdate(默认)。最多播放: views。senddate: 最新发布
	SeriesID int    `json:"series_id,omitempty" request:"query,omitempty"` // 系列 ID。用于过滤结果, 即若某一视频包含在系列内则不返回该视频
}

// GetVideoByKeywords 根据关键词查找视频
func (c *Client) GetVideoByKeywords(ctx context.Context, param GetVideoByKeywordsParam) (*GetVideoByKeywordsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/series/recArchivesByKeywords"
	)
	return bilibili.Execute[*GetVideoByKeywordsResult](ctx, c.Client, method, url, param)
}

// GetVideoSeriesInfoParam 指定要查询的视频列表（系列）。
type GetVideoSeriesInfoParam struct {
	Mid        int    `json:"mid"`                                             // UP 主 ID
	SeriesID   int    `json:"series_id"`                                       // 视频合集 ID
	Sort       string `json:"sort,omitempty" request:"query,omitempty"`        // 未知
	Pn         int    `json:"pn,omitempty" request:"query,omitempty"`          // 页码索引
	Ps         int    `json:"ps,omitempty" request:"query,omitempty"`          // 单页内容数量
	CurrentMid int    `json:"current_mid,omitempty" request:"query,omitempty"` // 单页内容数量
}

// GetVideoSeriesInfo 获取视频列表信息（在个人空间里创建的叫做视频列表，在创作中心里创建的叫合集，注意区分）
func (c *Client) GetVideoSeriesInfo(ctx context.Context, param GetVideoSeriesInfoParam) (*GetVideoSeriesInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/series/archives"
	)
	return bilibili.Execute[*GetVideoSeriesInfoResult](ctx, c.Client, method, url, param)
}
