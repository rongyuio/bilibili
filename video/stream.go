package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频流地址接口。响应模型见 model_stream.go。

// GetVideoStreamParam 指定要获取播放地址的分P与清晰度。
type GetVideoStreamParam struct {
	Avid        int    `json:"avid,omitempty" request:"query,omitempty"`         // 稿件 avid。avid 与 bvid 任选一个
	Bvid        string `json:"bvid,omitempty" request:"query,omitempty"`         // 稿件 bvid。avid 与 bvid 任选一个
	Cid         int    `json:"cid"`                                              // 视频 cid
	Qn          int    `json:"qn,omitempty" request:"query,omitempty"`           // 视频清晰度选择。未登录默认 32（480P），登录后默认 64（720P）。含义见 [上表](#qn视频清晰度标识)。DASH 格式时无效
	Fnval       int    `json:"fnval,omitempty" request:"query,omitempty"`        // 视频流格式标识。默认值为1（MP4 格式）。含义见 [ 上表](#fnval视频流格式标识)
	Fnver       int    `json:"fnver,omitempty" request:"query,omitempty"`        // 0
	Fourk       int    `json:"fourk,omitempty" request:"query,omitempty"`        // 是否允许 4K 视频。画质最高 1080P：0（默认）。画 质最高 4K：1
	Session     string `json:"session,omitempty" request:"query,omitempty"`      // 从视频播放页的 HTML 中获取
	Otype       string `json:"otype,omitempty" request:"query,omitempty"`        // 固定为json
	Type        string `json:"type,omitempty" request:"query,omitempty"`         // 目前为空
	Platform    string `json:"platform,omitempty" request:"query,omitempty"`     // pc：web播放（默认值，视频流存在 referer鉴权）。html5：移动端 HTML5 播放（仅支持 MP4 格式，无 referer 鉴权可以直接使用video标签播放）
	HighQuality int    `json:"high_quality,omitempty" request:"query,omitempty"` // 是否高画质。platform=html5时，此值 为1可使画质为1080p
}

// GetVideoStream 获取视频流地址_web端
func (c *Client) GetVideoStream(ctx context.Context, param GetVideoStreamParam) (*GetVideoStreamResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/wbi/playurl"
	)
	return bilibili.Execute[*GetVideoStreamResult](ctx, c.Client, method, url, param, c.FillWbi())
}
