package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户状态数相关接口。响应模型见 model_status.go。

// RelationStatParam 是查询关系状态数的请求参数。
type RelationStatParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
	Vmid      int    `json:"vmid" request:"query"`                           // 目标用户mid
}

// GetRelationStat 查询关系状态数。
func (c *Client) GetRelationStat(ctx context.Context, param RelationStatParam) (*GetRelationStatResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/stat"
	)
	return bilibili.Execute[*GetRelationStatResult](ctx, c.Client, method, url, param)
}

// UpStatParam 是查询UP主状态数的请求参数。
type UpStatParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token。APP方式必要
	Mid       int    `json:"mid" request:"query"`                            // 目标用户mid
}

// GetUpStat 查询UP主状态数。需要任意用户登录，否则不返回任何数据。
func (c *Client) GetUpStat(ctx context.Context, param UpStatParam) (*GetUpStatResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/upstat"
	)
	return bilibili.Execute[*GetUpStatResult](ctx, c.Client, method, url, param)
}

// NavNumParam 是查询用户导航栏状态数的请求参数。
type NavNumParam struct {
	Mid         int    `json:"mid" request:"query"`                              // 目标用户mid
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面标识，默认 333.999
}

// GetNavNum 查询用户导航栏状态数。
func (c *Client) GetNavNum(ctx context.Context, param NavNumParam) (*GetNavNumResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/navnum"
	)
	return bilibili.Execute[*GetNavNumResult](ctx, c.Client, method, url, param)
}

// AlbumUploadCountParam 是查询相簿投稿数的请求参数。
type AlbumUploadCountParam struct {
	Uid int `json:"uid" request:"query"` // 目标用户mid
}

// GetAlbumUploadCount 查询相簿投稿数。
func (c *Client) GetAlbumUploadCount(ctx context.Context, param AlbumUploadCountParam) (*GetAlbumUploadCountResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_draw/v1/doc/upload_count"
	)
	return bilibili.Execute[*GetAlbumUploadCountResult](ctx, c.Client, method, url, param)
}
