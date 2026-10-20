package audio

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 音频流URL相关接口。响应模型见 audio_musicstream_url_model.go。

// GetAudioStreamURLParam 是获取音频流URL（web 端）的请求参数。
type GetAudioStreamURLParam struct {
	Sid       int `json:"sid" request:"query"`                           // 音频auid
	Quality   int `json:"quality,omitempty" request:"query,omitempty"`   // 音质代码。0：128K。1：192K。2：320K。3：无损FLAC
	Privilege int `json:"privilege,omitempty" request:"query,omitempty"` // 固定为2。不必要
}

// GetAudioStreamURL 获取音频流URL（web端）
//
// 本接口仅能获取192K音质的音频；付费歌曲只能获取30秒试听片段。
func (c *Client) GetAudioStreamURL(ctx context.Context, param GetAudioStreamURLParam) (*GetAudioStreamURLResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/url"
	)
	return bilibili.Execute[*GetAudioStreamURLResult](ctx, c.Client, method, url, param)
}

// GetPaidAudioStreamURLParam 是获取音频流URL（可获取付费音频）的请求参数。
type GetPaidAudioStreamURLParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
	SongID    int    `json:"songid" request:"query"`                         // 音频auid
	Quality   int    `json:"quality" request:"query"`                        // 音质代码。0：128K。1：192K。2：320K。3：无损FLAC
	Privilege int    `json:"privilege" request:"query"`                      // 必须为2
	Mid       int    `json:"mid" request:"query"`                            // 当前用户mid。可为任意值
	Platform  string `json:"platform" request:"query"`                       // 平台标识。可为任意值
}

// GetPaidAudioStreamURL 获取音频流URL（可获取付费音频）
//
// 认证方式为APP或Cookie（SESSDATA）；付费音乐需要大会员或音乐包账号，
// 无损音质需要登录用户为会员，否则返回试听片段。
func (c *Client) GetPaidAudioStreamURL(ctx context.Context, param GetPaidAudioStreamURLParam) (*GetPaidAudioStreamURLResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/audio/music-service-c/url"
	)
	return bilibili.Execute[*GetPaidAudioStreamURLResult](ctx, c.Client, method, url, param)
}
