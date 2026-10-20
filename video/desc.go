package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetVideoDescParam 指定要查询简介的稿件。avid 与 bvid 任选一个。
type GetVideoDescParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件 avid
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件 bvid
}

// GetVideoDesc 获取视频简介。
func (c *Client) GetVideoDesc(ctx context.Context, param GetVideoDescParam) (string, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/desc"
	)
	return bilibili.Execute[string](ctx, c.Client, method, url, param)
}
