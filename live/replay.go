package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播回放相关接口。响应模型见 replay_model.go。

// GetLiveReplayListParam 是获取自己的直播回放列表的请求参数。
type GetLiveReplayListParam struct {
	Page     int `json:"page,omitempty" request:"query,omitempty"`      // 页码。默认为第 1 页
	PageSize int `json:"page_size,omitempty" request:"query,omitempty"` // 每页内容数量。默认为 30，最大 30
}

// GetLiveReplayList 获取自己的直播回放列表，只能获取最近 14 天的回放。
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetLiveReplayList(ctx context.Context, param GetLiveReplayListParam) (*GetLiveReplayListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/AnchorGetReplayList"
	)
	return bilibili.Execute[*GetLiveReplayListResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetOtherLiveReplayListParam 是获取某位主播的回放列表的请求参数。
type GetOtherLiveReplayListParam struct {
	LiveUID     int64  `json:"live_uid"`                                         // 目标主播的 uid
	TimeRange   int    `json:"time_range,omitempty" request:"query,omitempty"`   // 回放列表的时间范围。默认为近 14 天。1：近 3 天。2：近 7 天。3：近 14 天
	Page        int    `json:"page,omitempty" request:"query,omitempty"`         // 页码。默认为第 1 页
	PageSize    int    `json:"page_size,omitempty" request:"query,omitempty"`    // 每页内容数量。默认为 30，最大 30
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetOtherLiveReplayList 获取某位主播的回放列表。
//
// 需要获得那位主播的授权才能获取数据；需要 Cookie（SESSDATA）。
func (c *Client) GetOtherLiveReplayList(ctx context.Context, param GetOtherLiveReplayListParam) (*GetOtherLiveReplayListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/videoService/GetOtherSliceList"
	)
	return bilibili.Execute[*GetOtherLiveReplayListResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetLivePublishedSliceListParam 是获取自己已发布片段信息的请求参数。
type GetLivePublishedSliceListParam struct {
	Page     int `json:"page,omitempty" request:"query,omitempty"`      // 页码。默认为第 1 页
	PageSize int `json:"page_size,omitempty" request:"query,omitempty"` // 每页内容数量。默认为 20，最大 20
}

// GetLivePublishedSliceList 获取自己已发布片段的信息，只能获取主播自己的已发布片段。
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetLivePublishedSliceList(ctx context.Context, param GetLivePublishedSliceListParam) (*GetLivePublishedSliceListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/AnchorGetVideoSliceList"
	)
	return bilibili.Execute[*GetLivePublishedSliceListResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetOtherPublishedSliceListParam 是获取你为某位主播剪辑的已发布片段的请求参数。
type GetOtherPublishedSliceListParam struct {
	LiveUID     int64  `json:"live_uid"`                                         // 目标主播的 uid
	Page        int    `json:"page,omitempty" request:"query,omitempty"`         // 页码。默认为第 1 页
	PageSize    int    `json:"page_size,omitempty" request:"query,omitempty"`    // 每页内容数量。默认为 20，最大 20
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetOtherPublishedSliceList 获取你为某位主播剪辑的已发布片段。
//
// 需要获得那位主播的授权才能获取数据；需要 Cookie（SESSDATA）。
func (c *Client) GetOtherPublishedSliceList(ctx context.Context, param GetOtherPublishedSliceListParam) (*GetOtherPublishedSliceListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/videoService/GetPublishedList"
	)
	return bilibili.Execute[*GetOtherPublishedSliceListResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetLiveReplayDraftListParam 是获取自己的回放剪辑草稿列表的请求参数。
type GetLiveReplayDraftListParam struct {
	Page     int `json:"page,omitempty" request:"query,omitempty"`      // 页码。默认为第 1 页
	PageSize int `json:"page_size,omitempty" request:"query,omitempty"` // 每页内容数量。默认为 30，最大 30
}

// GetLiveReplayDraftList 获取自己的回放剪辑草稿列表，只能获取主播自己的草稿。
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetLiveReplayDraftList(ctx context.Context, param GetLiveReplayDraftListParam) (*GetLiveReplayDraftListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/GetDraftList"
	)
	return bilibili.Execute[*GetLiveReplayDraftListResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetOtherReplayDraftListParam 是获取你为某位主播剪辑的回放草稿的请求参数。
type GetOtherReplayDraftListParam struct {
	LiveUID     int64  `json:"live_uid"`                                         // 目标主播的 uid
	Page        int    `json:"page,omitempty" request:"query,omitempty"`         // 页码。默认为第 1 页
	PageSize    int    `json:"page_size,omitempty" request:"query,omitempty"`    // 每页内容数量。默认为 30，最大 30
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetOtherReplayDraftList 获取你为某位主播剪辑的回放草稿。
//
// 需要获得那位主播的授权才能获取数据；需要 Cookie（SESSDATA）。
func (c *Client) GetOtherReplayDraftList(ctx context.Context, param GetOtherReplayDraftListParam) (*GetOtherReplayDraftListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/videoService/GetDraftList"
	)
	return bilibili.Execute[*GetOtherReplayDraftListResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetLiveReplayInfoParam 是获取指定回放信息的请求参数。
type GetLiveReplayInfoParam struct {
	RecordID int64 `json:"record_id"` // 直播回放 id
}

// GetLiveReplayInfo 获取指定回放的信息。
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetLiveReplayInfo(ctx context.Context, param GetLiveReplayInfoParam) (*GetLiveReplayInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/GetAnchorVideoUidRecordsSubsect"
	)
	return bilibili.Execute[*GetLiveReplayInfoResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetLiveSliceStreamParam 是获取自己某场直播切片视频流的请求参数。
type GetLiveSliceStreamParam struct {
	LiveKey     string `json:"live_key"`                                         // 标记直播场次的 key
	StartTime   int64  `json:"start_time"`                                       // 直播开始时间戳
	EndTime     int64  `json:"end_time"`                                         // 直播结束时间戳
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetLiveSliceStream 获取自己某场直播的切片视频流。
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetLiveSliceStream(ctx context.Context, param GetLiveSliceStreamParam) (*GetLiveSliceStreamResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/GetSliceStream"
	)
	return bilibili.Execute[*GetLiveSliceStreamResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetOtherLiveSliceStreamParam 是获取某位主播某场直播切片视频流的请求参数。
type GetOtherLiveSliceStreamParam struct {
	LiveKey     string `json:"live_key"`                                         // 标记直播场次的 key
	StartTime   int64  `json:"start_time"`                                       // 直播开始时间戳
	EndTime     int64  `json:"end_time"`                                         // 直播结束时间戳
	LiveUID     int64  `json:"live_uid"`                                         // 目标主播的 uid
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetOtherLiveSliceStream 获取某位主播某场直播的切片视频流。
//
// 需要获得那位主播的授权才能获取数据；需要 Cookie（SESSDATA）。
func (c *Client) GetOtherLiveSliceStream(ctx context.Context, param GetOtherLiveSliceStreamParam) (*GetOtherLiveSliceStreamResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/videoService/GetUserSliceStream"
	)
	return bilibili.Execute[*GetOtherLiveSliceStreamResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetLiveReplaySessionDataParam 是获取自己某场直播会话数据的请求参数。
type GetLiveReplaySessionDataParam struct {
	LiveKey     string `json:"live_key"`                                         // 标记直播场次的 key
	StartTm     string `json:"start_tm"`                                         // 开始时间，格式为 yyyy-mm-dd HH:MM:SS（UTC+08:00）；取值对实际无影响
	EndTm       string `json:"end_tm"`                                           // 结束时间，格式为 yyyy-mm-dd HH:MM:SS（UTC+08:00）；取值对实际无影响
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetLiveReplaySessionData 获取自己某场直播的直播会话数据。
//
// 需要 Cookie（SESSDATA）。
func (c *Client) GetLiveReplaySessionData(ctx context.Context, param GetLiveReplaySessionDataParam) (*GetLiveReplaySessionDataResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/GetLiveSessionData"
	)
	return bilibili.Execute[*GetLiveReplaySessionDataResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetOtherLiveReplaySessionDataParam 是获取某位主播某场直播会话数据的请求参数。
type GetOtherLiveReplaySessionDataParam struct {
	LiveKey     string `json:"live_key"`                                         // 标记直播场次的 key
	StartTm     string `json:"start_tm"`                                         // 开始时间，格式为 yyyy-mm-dd HH:MM:SS（UTC+08:00）；取值对实际无影响
	EndTm       string `json:"end_tm"`                                           // 结束时间，格式为 yyyy-mm-dd HH:MM:SS（UTC+08:00）；取值对实际无影响
	LiveUID     int64  `json:"live_uid"`                                         // 目标主播的 uid
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)
}

// GetOtherLiveReplaySessionData 获取某位主播某场直播的直播会话数据。
//
// 若已提供 live_key 与 live_uid，则必须正确提供 start_tm 与 end_tm，
// 否则服务器可能无法成功响应。需要获得那位主播的授权才能获取数据；需要 Cookie（SESSDATA）。
func (c *Client) GetOtherLiveReplaySessionData(ctx context.Context, param GetOtherLiveReplaySessionDataParam) (*GetOtherLiveReplaySessionDataResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-room/v1/videoService/GetUserLiveSessionData"
	)
	return bilibili.Execute[*GetOtherLiveReplaySessionDataResult](ctx, c.Client, method, url, param, liveReferer(0))
}

// GetLiveManualHighlightListParam 是获取某场直播用户高光列表的请求参数。
type GetLiveManualHighlightListParam struct {
	LiveKey     string `json:"live_key"`                                         // 标记直播场次的 key
	ManualType  int    `json:"manual_type"`                                      // 查询类型。1：回溯录制。2：已保存片段
	LiveUID     int64  `json:"live_uid,omitempty" request:"query,omitempty"`     // 目标主播的 uid。通过粉丝使用的接口时必要；填写后使用粉丝接口
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // (?)。作用尚不明确
}

// GetLiveManualHighlightList 获取某场直播的用户高光列表。
//
// 不填 LiveUID 时使用主播接口（只能查询自己的场次）；填写 LiveUID 时使用粉丝接口，
// 需要获得那位主播的授权。需要 Cookie（SESSDATA）。
func (c *Client) GetLiveManualHighlightList(ctx context.Context, param GetLiveManualHighlightListParam) (*GetLiveManualHighlightListResult, error) {
	const (
		method = resty.MethodGet
	)
	url := "https://api.live.bilibili.com/xlive/app-blink/v1/anchorVideo/UserGetManualHighlightList"
	if param.LiveUID != 0 {
		url = "https://api.live.bilibili.com/xlive/web-room/v1/videoService/UserGetManualHighlightList"
	}
	return bilibili.Execute[*GetLiveManualHighlightListResult](ctx, c.Client, method, url, param, liveReferer(0))
}
