package bangumi

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 番剧追番相关接口。响应模型见本文件。
//
// 这两个接口把结果放在根对象的 result 字段而非标准的 data 字段，
// 因此通过 ExecuteRaw 解码 result 信封；请求链路仍沿用根包统一的
// Cookie 快照、参数编码、CSRF 与错误语义。

// FollowBangumiParam 指定要追的番剧。
type FollowBangumiParam struct {
	SeasonID string `json:"season_id" request:"query"` // 剧集 ssid
}

// FollowBangumiResult 是追番接口 result 字段的数据本体。
type FollowBangumiResult struct {
	Fmid     int    `json:"fmid"`     // 实测为 0
	Relation bool   `json:"relation"` // 实测为 false
	Status   int    `json:"status"`   // 实测为 2
	Toast    string `json:"toast"`    // 提示文案。实测为「自己追的番就要好好看完哟^o^」
}

// FollowBangumi 追番，CSRF 从请求 Cookie 快照读取后作为 URL 参数提交。
func (c *Client) FollowBangumi(ctx context.Context, param FollowBangumiParam) (*FollowBangumiResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/pgc/web/follow/add"
	)
	return executeResult[FollowBangumiResult](ctx, c, method, url, param, bilibili.FillCsrf(c.Client))
}

// UnfollowBangumiParam 指定要取消追番的番剧。
type UnfollowBangumiParam struct {
	SeasonID string `json:"season_id" request:"query"` // 剧集 ssid
}

// UnfollowBangumiResult 是取消追番接口 result 字段的数据本体。
type UnfollowBangumiResult struct {
	Fmid     int    `json:"fmid"`     // 实测为 0
	Relation bool   `json:"relation"` // 文档表格为 true，但响应示例为 false
	Status   int    `json:"status"`   // 实测为 0
	Toast    string `json:"toast"`    // 提示文案。实测为「已取消追番」
}

// UnfollowBangumi 取消追番，CSRF 从请求 Cookie 快照读取后作为 URL 参数提交。
func (c *Client) UnfollowBangumi(ctx context.Context, param UnfollowBangumiParam) (*UnfollowBangumiResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/pgc/web/follow/del"
	)
	return executeResult[UnfollowBangumiResult](ctx, c, method, url, param, bilibili.FillCsrf(c.Client))
}
