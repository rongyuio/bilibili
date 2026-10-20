package audio

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 音频投币与收藏相关接口。
//
// 这两个接口的 data 层均为标量：查询音频收藏状态为 bool，查询音频投币数为
// num，因此直接返回标量，不定义结果结构体。

// GetAudioCollectionStatusParam 是查询音频收藏状态的请求参数。
type GetAudioCollectionStatusParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioCollectionStatus 查询音频收藏状态。
func (c *Client) GetAudioCollectionStatus(ctx context.Context, param GetAudioCollectionStatusParam) (bool, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/collections/songs-coll"
	)
	return bilibili.Execute[bool](ctx, c.Client, method, url, param)
}

// GetAudioCoinCountParam 是查询音频投币数的请求参数。
type GetAudioCoinCountParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioCoinCount 查询音频投币数。
func (c *Client) GetAudioCoinCount(ctx context.Context, param GetAudioCoinCountParam) (int, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/coin/audio"
	)
	return bilibili.Execute[int](ctx, c.Client, method, url, param)
}
