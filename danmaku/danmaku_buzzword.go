package danmaku

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 名词解释相关接口。各接口的返回体独立定义，互不复用。

// GetBuzzwordListParam 指定要拉取的名词列表。
type GetBuzzwordListParam struct {
	TypeID int `json:"type_id" request:"query"`                // 名词类型。如：4代表“哔哩哔哩热词图鉴”，5代表“语言类”
	Pn     int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
	Ps     int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为20
}

// GetBuzzwordList 拉取名词解释列表
func (c *Client) GetBuzzwordList(ctx context.Context, param GetBuzzwordListParam) (*GetBuzzwordListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/buzzword/list"
	)
	return bilibili.Execute[*GetBuzzwordListResult](ctx, c.Client, method, url, param)
}

// GetBuzzwordDetailParam 指定要查询的名词。
type GetBuzzwordDetailParam struct {
	ID int `json:"id" request:"query"` // 名词id。如：1代表“kksk”，2代表“打脸”
}

// GetBuzzwordDetail 查询名词解释详情
func (c *Client) GetBuzzwordDetail(ctx context.Context, param GetBuzzwordDetailParam) (*GetBuzzwordDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/buzzword/detail"
	)
	return bilibili.Execute[*GetBuzzwordDetailResult](ctx, c.Client, method, url, param)
}
