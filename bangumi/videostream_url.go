package bangumi

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 番剧视频流 URL 相关接口。响应模型见 model_videostream_url.go。

// GetBangumiVideoStreamParam 指定获取番剧视频流 URL 的参数。
// ep_id 与 cid 任选一个。
type GetBangumiVideoStreamParam struct {
	Avid        int    `json:"avid,omitempty" request:"query,omitempty"`          // 稿件 avid
	Bvid        string `json:"bvid,omitempty" request:"query,omitempty"`          // 稿件 bvid
	EpID        int    `json:"ep_id,omitempty" request:"query,omitempty"`         // 稿件 epid。ep_id 与 cid 任选一个
	Cid         int    `json:"cid,omitempty" request:"query,omitempty"`           // 视频 cid。ep_id 与 cid 任选一个
	Qn          int    `json:"qn,omitempty" request:"query,omitempty"`            // 视频清晰度选择。未登录默认 32（480P），登录默认 64（720P）。DASH 方式无效
	Fnval       int    `json:"fnval,omitempty" request:"query,omitempty"`         // 视频获取方式选择。0/2：flv；1：低清 mp4（仅 240P 与 360P）；16/80：dash
	Fnver       int    `json:"fnver,omitempty" request:"query,omitempty"`         // 固定为 0
	Fourk       int    `json:"fourk,omitempty" request:"query,omitempty"`         // 是否允许 4K 视频。0：画质最高 1080P；1：画质最高 4K
	Session     string `json:"session,omitempty" request:"query,omitempty"`       // 从视频播放页的网页源码中获取
	FromClient  string `json:"from_client,omitempty" request:"query,omitempty"`   // BROWSER
	DrmTechType int    `json:"drm_tech_type,omitempty" request:"query,omitempty"` // 2
}

// GetBangumiVideoStream 获取番剧视频流 URL。
//
// 该接口的响应信封使用 result 而非 data，故通过 ExecuteRaw 手动解码。
// 获取 480P 及以上清晰度需要登录 Cookie，1080P60 / 1080P+ 等清晰度需要大会员；
// 默认请求头中的 Referer 为 https://www.bilibili.com，满足取流要求。
// 返回的 URL 有效期为 120 分钟；FLV / MP4 方式可能返回多个分段。
func (c *Client) GetBangumiVideoStream(ctx context.Context, param GetBangumiVideoStreamParam) (*GetBangumiVideoStreamResult, error) {
	const (
		method   = resty.MethodGet
		endpoint = "https://api.bilibili.com/pgc/player/web/playurl"
	)
	return executeResult[GetBangumiVideoStreamResult](ctx, c, method, endpoint, param)
}
