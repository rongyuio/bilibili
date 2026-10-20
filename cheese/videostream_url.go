package cheese

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 课程视频流 URL 相关接口。响应模型见 model_videostream_url.go。

// GetCheeseVideoStreamParam 指定获取课程视频流 URL 的参数。
// avid、ep_id、cid 三者均为必要参数。
type GetCheeseVideoStreamParam struct {
	Avid  int `json:"avid"`                                      // 课程 avid
	EpID  int `json:"ep_id"`                                     // 课程 epid
	Cid   int `json:"cid"`                                       // 视频 cid
	Qn    int `json:"qn,omitempty" request:"query,omitempty"`    // 视频清晰度选择。含义见 [qn视频清晰度标识]
	Fnver int `json:"fnver,omitempty" request:"query,omitempty"` // 视频流版本。含义见 [fnver视频流版本标识]
	Fnval int `json:"fnval,omitempty" request:"query,omitempty"` // 视频流类型。含义见 [fnval视频流格式标识]
	Fourk int `json:"fourk,omitempty" request:"query,omitempty"` // 是否允许 4K 视频。0：画质最高 1080P（默认）；1：画质最高 4K
}

// GetCheeseVideoStream 获取课程视频流 URL（web 端）。
//
// 本接口为课程视频专用，与普通视频不互通；课程未购买时返回 -403。
// 需要登录 Cookie（SESSDATA），返回的 URL 有效期为 120 分钟。
func (c *Client) GetCheeseVideoStream(ctx context.Context, param GetCheeseVideoStreamParam) (*GetCheeseVideoStreamResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pugv/player/web/playurl"
	)
	return bilibili.Execute[*GetCheeseVideoStreamResult](ctx, c.Client, method, url, param)
}
