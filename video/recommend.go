package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 首页视频推荐接口。响应模型见 model_recommend.go。

// GetTopRecommendVideoParam 指定首页推荐的翻页与相关性参数。
type GetTopRecommendVideoParam struct {
	FreshType  int `json:"fresh_type,omitempty" request:"query,omitempty"`   // 相关性。默认为3 。 值越大推荐内容越相关
	Version    int `json:"version,omitempty" request:"query,omitempty"`      // web端新旧版本:0为旧版本1为新版本。默认为 0 。1,0分别为新旧web端
	Ps         int `json:"ps,omitempty" request:"query,omitempty"`           // pagesize 单页返回的记录条数默认为10或8。默认为10 。当version为1时默认为8
	FreshIdx   int `json:"fresh_idx,omitempty" request:"query,omitempty"`    // 翻页相关。默认为1 。 与翻页相关
	FreshIdx1H int `json:"fresh_idx_1h,omitempty" request:"query,omitempty"` // 翻页相关。默认为1 。 与翻页相关
}

// GetTopRecommendVideo 获取首页视频推荐列表
func (c *Client) GetTopRecommendVideo(ctx context.Context, param GetTopRecommendVideoParam) (*GetTopRecommendVideoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/index/top/feed/rcmd"
	)
	return bilibili.Execute[*GetTopRecommendVideoResult](ctx, c.Client, method, url, param)
}
