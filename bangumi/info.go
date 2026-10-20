package bangumi

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 剧集基本信息相关接口。响应模型见 model_info.go。
//
// 这三个接口把结果放在根对象的 result 字段而非标准的 data 字段，
// 因此通过 ExecuteRaw 解码 result 信封；请求链路仍沿用根包统一的
// Cookie 快照、参数编码与错误语义。

// GetBangumiInfoParam 获取剧集基本信息的参数。
type GetBangumiInfoParam struct {
	MediaID string `json:"media_id" request:"query"` // 剧集 mdid
}

// GetBangumiInfo 获取剧集基本信息（mdid 方式）。
func (c *Client) GetBangumiInfo(ctx context.Context, param GetBangumiInfoParam) (*GetBangumiInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pgc/review/user"
	)
	return executeResult[GetBangumiInfoResult](ctx, c, method, url, param)
}

// GetBangumiDetailParam 获取剧集明细的参数。season_id 与 ep_id 任选其一。
type GetBangumiDetailParam struct {
	SeasonID int `json:"season_id,omitempty" request:"query,omitempty"` // 番剧 ssid。season_id 与 ep_id 任选其一
	EpID     int `json:"ep_id,omitempty" request:"query,omitempty"`     // 剧集 epid。season_id 与 ep_id 任选其一
}

// GetBangumiDetail 获取剧集明细（web 端，ssid/epid 方式）。
func (c *Client) GetBangumiDetail(ctx context.Context, param GetBangumiDetailParam) (*GetBangumiDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pgc/view/web/season"
	)
	return executeResult[GetBangumiDetailResult](ctx, c, method, url, param)
}

// GetBangumiSectionParam 获取剧集分集信息的参数。
type GetBangumiSectionParam struct {
	SeasonID string `json:"season_id" request:"query"` // 剧集 ssid
}

// GetBangumiSection 获取剧集分集信息。
func (c *Client) GetBangumiSection(ctx context.Context, param GetBangumiSectionParam) (*GetBangumiSectionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pgc/web/season/section"
	)
	return executeResult[GetBangumiSectionResult](ctx, c, method, url, param)
}
