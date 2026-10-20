package audio

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 音频状态数相关接口。响应模型见 audio_status_number_model.go。

// GetAudioStatusNumberParam 是查询歌曲状态数的请求参数。
type GetAudioStatusNumberParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioStatusNumber 查询歌曲状态数。
func (c *Client) GetAudioStatusNumber(ctx context.Context, param GetAudioStatusNumberParam) (*GetAudioStatusNumberResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/stat/song"
	)
	return bilibili.Execute[*GetAudioStatusNumberResult](ctx, c.Client, method, url, param)
}
