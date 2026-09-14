package bilibili

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// 历史记录与稍后再看接口。响应模型见 history_model.go。

type GetHistoryParam struct {
	Max      int    `json:"max,omitempty" request:"query,omitempty"`      // 历史记录截止目标 id。默认为 0。稿件：稿件 avid。剧集（番剧 / 影视）：剧集 ssid。直播：直播间 id。文集：文集 rlid。文章：文章 cvid
	Business string `json:"business,omitempty" request:"query,omitempty"` // 历史记录截止目标业务类型。默认为空。archive：稿件。pgc：剧集（番剧 / 影视）。live：直播。article-list：文集。article：文章
	ViewAt   int    `json:"view_at,omitempty" request:"query,omitempty"`  // 历史记录截止时间。时间戳。默认为 0。0 为当前 时间
	Type     string `json:"type,omitempty" request:"query,omitempty"`     // 历史记录分类筛选。all：全部类型（默认）。archive：稿件。live：直播。article：文章
	Ps       int    `json:"ps,omitempty" request:"query,omitempty"`       // 每页项数。默认为 20，最大 30
}

// GetHistory 获取历史记录列表
func (c *Client) GetHistory(ctx context.Context, param GetHistoryParam) (*HistoryInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/history/cursor"
	)
	return execute[*HistoryInfo](ctx, c, method, url, param)
}

type DeleteHistoryParam struct {
	Kid string `json:"kid"` // 删除的目标记录，格式为{业务类型}_{目标id}详见备注。视频：archive_{稿件avid}。直播：live_{直播间id}。专栏：article_{专栏cvid}。剧集：pgc_{剧集ssid}。文集：article-list_{文集rlid}
}

// DeleteHistory 删除历史记录
func (c *Client) DeleteHistory(ctx context.Context, param DeleteHistoryParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/delete"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// ClearHistory 清空历史记录
func (c *Client) ClearHistory(ctx context.Context) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/clear"
	)
	_, err := execute[any](ctx, c, method, url, nil, fillCsrf(c))
	return err
}

type SetHistoryDisableParam struct {
	Switch bool `json:"switch,omitempty" request:"query,omitempty"` // 停用开关。true：停用。false：正常。默认为false
}

// SetHistoryDisable 停用历史记录
func (c *Client) SetHistoryDisable(ctx context.Context, param SetHistoryDisableParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/shadow/set"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// GetHistoryDisableState 查询历史记录停用状态 true：停用 false：正常
func (c *Client) GetHistoryDisableState(ctx context.Context) (bool, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/history/shadow"
	)
	return execute[bool](ctx, c, method, url, nil)
}

// AddToView 视频添加稍后再看
func (c *Client) AddToView(ctx context.Context, param VideoParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/toview/add"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type AddChannelAllToViewParam struct {
	Cid int `json:"cid"` // 目标频道id
	Mid int `json:"mid"` // 目标频道所属的用户mid
}

// AddChannelkAllToView 添加频道中所有视频到稍后再看
func (c *Client) AddChannelkAllToView(ctx context.Context, param AddChannelAllToViewParam) error {
	const (
		method = resty.MethodPost
		url    = "https://space.bilibili.com/ajax/channel/addAllToView"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// GetToViewList 获取稍后再看视频列表
func (c *Client) GetToViewList(ctx context.Context) (*ToViewInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/history/toview"
	)
	return execute[*ToViewInfo](ctx, c, method, url, nil)
}

type DeleteToViewParam struct {
	Viewed bool `json:"viewed,omitempty" request:"query,omitempty"` // 是否删除所有已观看的视频。true：删除已观看视 频。false：不删除已观看视频。默认为false
	Aid    int  `json:"aid,omitempty" request:"query,omitempty"`    // 删除的目标记录的avid
}

// DeleteToView 删除稍后再看视频
func (c *Client) DeleteToView(ctx context.Context, param DeleteToViewParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/toview/del"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// ClearToView 清空稍后再看视频列表
func (c *Client) ClearToView(ctx context.Context) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v2/history/toview/clear"
	)
	_, err := execute[any](ctx, c, method, url, nil, fillCsrf(c))
	return err
}
