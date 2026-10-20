package audio

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 音频歌单与音频收藏夹相关接口。响应模型见 audio_music_list_model.go。

// GetAudioCollectionListParam 是查询自己创建的歌单的请求参数。
type GetAudioCollectionListParam struct {
	Pn int `json:"pn" request:"query"` // 页码
	Ps int `json:"ps" request:"query"` // 每页项数
}

// GetAudioCollectionList 查询自己创建的歌单。
func (c *Client) GetAudioCollectionList(ctx context.Context, param GetAudioCollectionListParam) (*GetAudioCollectionListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/collections/list"
	)
	return bilibili.Execute[*GetAudioCollectionListResult](ctx, c.Client, method, url, param)
}

// GetAudioCollectionInfoParam 是查询音频收藏夹（默认歌单）信息的请求参数。
type GetAudioCollectionInfoParam struct {
	Sid int `json:"sid" request:"query"` // 音频收藏夹 mlid。必须为默认收藏夹 mlid
}

// GetAudioCollectionInfo 查询音频收藏夹（默认歌单）信息。
func (c *Client) GetAudioCollectionInfo(ctx context.Context, param GetAudioCollectionInfoParam) (*GetAudioCollectionInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/collections/info"
	)
	return bilibili.Execute[*GetAudioCollectionInfoResult](ctx, c.Client, method, url, param)
}

// GetAudioHotMenuListParam 是查询热门歌单的请求参数。
type GetAudioHotMenuListParam struct {
	Pn int `json:"pn" request:"query"` // 页码
	Ps int `json:"ps" request:"query"` // 每页项数
}

// GetAudioHotMenuList 查询热门歌单。
func (c *Client) GetAudioHotMenuList(ctx context.Context, param GetAudioHotMenuListParam) (*GetAudioHotMenuListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/menu/hit"
	)
	return bilibili.Execute[*GetAudioHotMenuListResult](ctx, c.Client, method, url, param)
}

// GetAudioRankMenuListParam 是查询热门榜的请求参数。
type GetAudioRankMenuListParam struct {
	Pn int `json:"pn" request:"query"` // 页码
	Ps int `json:"ps" request:"query"` // 每页项数
}

// GetAudioRankMenuList 查询热门榜。
func (c *Client) GetAudioRankMenuList(ctx context.Context, param GetAudioRankMenuListParam) (*GetAudioRankMenuListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/menu/rank"
	)
	return bilibili.Execute[*GetAudioRankMenuListResult](ctx, c.Client, method, url, param)
}
