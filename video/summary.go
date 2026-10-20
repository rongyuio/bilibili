package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频AI总结相关接口。响应模型见 summary_model.go。

// GetVideoAISummaryParam 是获取视频AI总结的参数。
type GetVideoAISummaryParam struct {
	Aid   int    `json:"aid,omitempty" request:"query,omitempty"`    // 稿件 avid。avid 与 bvid 任选一个
	Bvid  string `json:"bvid,omitempty" request:"query,omitempty"`   // 稿件 bvid。avid 与 bvid 任选一个
	Cid   int    `json:"cid"`                                        // 视频 cid
	UpMid int    `json:"up_mid,omitempty" request:"query,omitempty"` // UP主 mid
}

// GetVideoAISummary 获取视频AI总结内容
//
// 需要登录态（Cookie SESSDATA），并使用 WBI 签名。
// data.code 表示摘要状态：-1 不支持AI摘要，0 有摘要，1 无摘要（未识别到语音）。
func (c *Client) GetVideoAISummary(ctx context.Context, param GetVideoAISummaryParam) (*GetVideoAISummaryResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/view/conclusion/get"
	)
	return bilibili.Execute[*GetVideoAISummaryResult](ctx, c.Client, method, url, param, c.FillWbi())
}
