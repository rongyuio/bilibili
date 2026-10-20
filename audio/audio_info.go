package audio

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 音频基本信息相关接口。响应模型见 audio_info_model.go。

// GetAudioInfoParam 是查询音频基本信息的请求参数。
type GetAudioInfoParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioInfo 查询音频基本信息。
func (c *Client) GetAudioInfo(ctx context.Context, param GetAudioInfoParam) (*GetAudioInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/song/info"
	)
	return bilibili.Execute[*GetAudioInfoResult](ctx, c.Client, method, url, param)
}

// GetAudioTagsParam 是查询音频 TAG 列表的请求参数。
type GetAudioTagsParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioTags 查询音频 TAG 列表。
func (c *Client) GetAudioTags(ctx context.Context, param GetAudioTagsParam) ([]GetAudioTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/tag/song"
	)
	return bilibili.Execute[[]GetAudioTagsItem](ctx, c.Client, method, url, param)
}

// GetAudioMembersParam 是查询音频创作成员列表的请求参数。
type GetAudioMembersParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioMembers 查询音频创作成员列表。
func (c *Client) GetAudioMembers(ctx context.Context, param GetAudioMembersParam) ([]GetAudioMembersGroup, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/member/song"
	)
	return bilibili.Execute[[]GetAudioMembersGroup](ctx, c.Client, method, url, param)
}

// GetAudioLyricParam 是获取音频歌词的请求参数。
type GetAudioLyricParam struct {
	Sid int `json:"sid" request:"query"` // 音频 auid
}

// GetAudioLyric 获取音频歌词，返回 lrc 格式文本。
func (c *Client) GetAudioLyric(ctx context.Context, param GetAudioLyricParam) (string, error) {
	const (
		method = resty.MethodGet
		url    = "https://www.bilibili.com/audio/music-service-c/web/song/lyric"
	)
	return bilibili.Execute[string](ctx, c.Client, method, url, param)
}
