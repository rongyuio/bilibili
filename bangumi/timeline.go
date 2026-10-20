package bangumi

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 番剧或影视时间线相关接口。响应模型见 model_timeline.go。

// GetBangumiTimelineParam 获取番剧或影视时间线的参数。
type GetBangumiTimelineParam struct {
	Types  string `json:"types" request:"query"`  // 类别。1：番剧。3：电影。4：国创
	Before int    `json:"before" request:"query"` // 开始于前几日。取值范围为 [0, 7]
	After  int    `json:"after" request:"query"`  // 结束于后几日。取值范围为 [0, 7]
}

// GetBangumiTimeline 获取番剧或影视时间线。
//
// 该接口把结果放在 result 而不是 data 中，因此通过 ExecuteRaw 解码
// result 信封，但仍沿用根包统一的 Cookie 快照、参数编码与错误处理语义。
func (c *Client) GetBangumiTimeline(ctx context.Context, param GetBangumiTimelineParam) ([]GetBangumiTimelineDay, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pgc/web/timeline"
	)
	days, err := executeResult[[]GetBangumiTimelineDay](ctx, c, method, url, param)
	if err != nil || days == nil {
		return nil, err
	}
	return *days, nil
}
