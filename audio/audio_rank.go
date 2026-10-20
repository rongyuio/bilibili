package audio

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 音频榜单相关接口。响应模型见 audio_rank_model.go。
//
// 这些接口的 csrf 参数均为非必要，因此不作为调用方参数，也不使用 CSRF 处理函数。

// GetAudioRankPeriodListParam 是获取音频榜单每期列表的请求参数。
type GetAudioRankPeriodListParam struct {
	ListType int `json:"list_type" request:"query"` // 榜单类型。1：热榜。2：原创榜
}

// GetAudioRankPeriodList 获取音频榜单每期列表。
func (c *Client) GetAudioRankPeriodList(ctx context.Context, param GetAudioRankPeriodListParam) (*GetAudioRankPeriodListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/copyright-music-publicity/toplist/all_period"
	)
	return bilibili.Execute[*GetAudioRankPeriodListResult](ctx, c.Client, method, url, param)
}

// GetAudioRankDetailParam 是查询音频榜单单期信息的请求参数。
type GetAudioRankDetailParam struct {
	ListID int `json:"list_id" request:"query"` // 榜单id。见 GetAudioRankPeriodList
}

// GetAudioRankDetail 查询音频榜单单期信息。
func (c *Client) GetAudioRankDetail(ctx context.Context, param GetAudioRankDetailParam) (*GetAudioRankDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/copyright-music-publicity/toplist/detail"
	)
	return bilibili.Execute[*GetAudioRankDetailResult](ctx, c.Client, method, url, param)
}

// GetAudioRankMusicListParam 是获取音频榜单单期内容的请求参数。
type GetAudioRankMusicListParam struct {
	ListID int `json:"list_id" request:"query"` // 榜单id。见 GetAudioRankPeriodList
}

// GetAudioRankMusicList 获取音频榜单单期内容。
func (c *Client) GetAudioRankMusicList(ctx context.Context, param GetAudioRankMusicListParam) (*GetAudioRankMusicListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/copyright-music-publicity/toplist/music_list"
	)
	return bilibili.Execute[*GetAudioRankMusicListResult](ctx, c.Client, method, url, param)
}
