package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户空间视频列表（系列 / 合集）与视频列表管理接口。文档见 video/collection.md。
//
// 响应模型见 model_series.go。注意与 collection.go 的
// GetVideoCollectionInfo / GetVideoSeriesInfo 区分：这里前三个接口查询用户空间的
// 系列列表，后五个接口管理个人空间里的「视频列表」。

// userSpaceSeasonsSeriesWebLocation 是系列和合集列表接口默认的页面位置标识。
const userSpaceSeasonsSeriesWebLocation = "333.999"

// GetUserSpaceSeriesListParam 指定要查询的用户空间。只获取系列视频。
type GetUserSpaceSeriesListParam struct {
	Mid        int    `json:"mid"`                                             // 用户 mid
	PageNum    int    `json:"page_num"`                                        // 页码索引
	PageSize   int    `json:"page_size"`                                       // 单页内容数量
	GaiaVtoken string `json:"gaia_vtoken,omitempty" request:"query,omitempty"` // 风控验证 token。被风控时（如 User-Agent 不正常）必填
}

// GetUserSpaceSeriesList 只获取系列视频
//
// 返回 data.items_lists 中的系列列表；seasons_list 为空。
func (c *Client) GetUserSpaceSeriesList(ctx context.Context, param GetUserSpaceSeriesListParam) (*GetUserSpaceSeriesListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-space/home/seasons_series"
	)
	return bilibili.Execute[*GetUserSpaceSeriesListResult](ctx, c.Client, method, url, param)
}

// GetUserSpaceSeasonsSeriesListParam 指定要查询的用户空间。获取系列和合集视频。
type GetUserSpaceSeasonsSeriesListParam struct {
	Mid         int    `json:"mid"`                                              // 用户 mid
	PageNum     int    `json:"page_num"`                                         // 页码。默认为 1
	PageSize    int    `json:"page_size"`                                        // 每页数量。默认为 20
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面位置标识。留空时由库填入 333.999
}

// GetUserSpaceSeasonsSeriesList 获取系列和合集视频
//
// 与 GetUserSpaceSeriesList 基本一致，但 data.items_lists.seasons_list 不为空。
func (c *Client) GetUserSpaceSeasonsSeriesList(ctx context.Context, param GetUserSpaceSeasonsSeriesListParam) (*GetUserSpaceSeasonsSeriesListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-space/seasons_series_list"
	)
	if param.WebLocation == "" {
		param.WebLocation = userSpaceSeasonsSeriesWebLocation
	}
	return bilibili.Execute[*GetUserSpaceSeasonsSeriesListResult](ctx, c.Client, method, url, param)
}

// GetSeriesInfoParam 指定要查询的系列。
type GetSeriesInfoParam struct {
	SeriesID int `json:"series_id"` // 系列 ID
}

// GetSeriesInfo 查询指定系列
func (c *Client) GetSeriesInfo(ctx context.Context, param GetSeriesInfoParam) (*GetSeriesInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/series/series"
	)
	return bilibili.Execute[*GetSeriesInfoResult](ctx, c.Client, method, url, param)
}

// CreateVideoSeriesParam 指定要创建的视频列表。
type CreateVideoSeriesParam struct {
	Mid         int    `json:"mid" request:"form-data"`                             // 用户 mid
	Name        string `json:"name" request:"form-data"`                            // 标题
	Keywords    string `json:"keywords,omitempty" request:"form-data,omitempty"`    // 关键词
	Description string `json:"description,omitempty" request:"form-data,omitempty"` // 简介
	Aids        string `json:"aids,omitempty" request:"form-data,omitempty"`        // 视频 aid 列表，以 , 分隔
}

// CreateVideoSeries 创建视频列表
//
// 请求体为 multipart/form-data，CSRF 放在 query 中。
func (c *Client) CreateVideoSeries(ctx context.Context, param CreateVideoSeriesParam) (*CreateVideoSeriesResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/series/series/createAndAddArchives"
	)
	return bilibili.Execute[*CreateVideoSeriesResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// DeleteVideoSeriesParam 指定要删除的视频列表。
type DeleteVideoSeriesParam struct {
	Mid      int    `json:"mid"`                                      // 用户 mid
	SeriesID int    `json:"series_id"`                                // 视频列表 ID
	Aids     string `json:"aids,omitempty" request:"query,omitempty"` // 视频 aid 列表。文档标注为「空」
}

// DeleteVideoSeries 删除视频列表
func (c *Client) DeleteVideoSeries(ctx context.Context, param DeleteVideoSeriesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/series/series/delete"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// SeriesArchivesParam 是添加 / 删除视频列表稿件共用的参数。
type SeriesArchivesParam struct {
	Mid      int    `json:"mid"`       // 用户 mid
	SeriesID int    `json:"series_id"` // 视频列表 ID
	Aids     string `json:"aids"`      // 视频 aid 列表，以 , 分隔
}

// DeleteVideoSeriesArchives 删除视频列表中的稿件
//
// 请求体为 x-www-form-urlencoded，CSRF 放在 query 中。
func (c *Client) DeleteVideoSeriesArchives(ctx context.Context, param SeriesArchivesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/series/series/delArchives"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("mid", "series_id", "aids"),
		bilibili.FillCsrf(c.Client))
	return err
}

// AddVideoSeriesArchives 添加稿件至视频列表
//
// 请求体为 x-www-form-urlencoded，CSRF 放在 query 中。
func (c *Client) AddVideoSeriesArchives(ctx context.Context, param SeriesArchivesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/series/series/addArchives"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("mid", "series_id", "aids"),
		bilibili.FillCsrf(c.Client))
	return err
}

// UpdateVideoSeriesParam 指定视频列表的修改内容。
type UpdateVideoSeriesParam struct {
	Mid         int    `json:"mid"`                                             // 用户 mid
	SeriesID    int    `json:"series_id"`                                       // 视频列表 ID
	Name        string `json:"name"`                                            // 标题
	Keywords    string `json:"keywords,omitempty" request:"query,omitempty"`    // 关键词
	Description string `json:"description,omitempty" request:"query,omitempty"` // 简介
	AddAids     string `json:"add_aids,omitempty" request:"query,omitempty"`    // 要添加的视频 aid 列表，以 , 分隔
	DelAids     string `json:"del_aids,omitempty" request:"query,omitempty"`    // 要删除的视频 aid 列表，以 , 分隔
	Aids        string `json:"aids,omitempty" request:"query,omitempty"`        // 文档标注为「空」
}

// UpdateVideoSeries 编辑视频列表信息
//
// 请求体为 x-www-form-urlencoded，CSRF 放在 query 中。
func (c *Client) UpdateVideoSeries(ctx context.Context, param UpdateVideoSeriesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/series/series/update"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("mid", "series_id", "name", "keywords", "description",
			"add_aids", "del_aids", "aids"),
		bilibili.FillCsrf(c.Client))
	return err
}
