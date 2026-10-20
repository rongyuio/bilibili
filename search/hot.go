package search

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 默认搜索与热搜相关接口。响应模型见 model_hot.go。

// GetSearchDefault 获取默认搜索内容（web端）。接口需要 WBI 签名。
func (c *Client) GetSearchDefault(ctx context.Context) (*GetSearchDefaultResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/search/default"
	)
	return bilibili.Execute[*GetSearchDefaultResult](ctx, c.Client, method, url, nil, c.FillWbi())
}

// GetSearchHotListParam 是获取热搜列表的参数。
type GetSearchHotListParam struct {
	Limit    int    `json:"limit" request:"query"`                        // 结果限制。范围 [1, 50]
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。web：web 端
}

// GetSearchHotList 获取热搜列表。接口需要 WBI 签名。
func (c *Client) GetSearchHotList(ctx context.Context, param GetSearchHotListParam) (*GetSearchHotListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/search/square"
	)
	return bilibili.Execute[*GetSearchHotListResult](ctx, c.Client, method, url, param, c.FillWbi())
}

// GetSearchHotListWeb 获取热搜列表（web端）。
//
// 该接口的响应把业务字段直接放在根对象上，没有标准的 data 信封，
// 因此不能复用 Execute：这里用 ExecuteRaw 取原始响应体后手动解码。
// context 检查、参数编码、非 200 状态码与业务码的错误语义与 Execute 一致。
func (c *Client) GetSearchHotListWeb(ctx context.Context) (*GetSearchHotListWebResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://s.search.bilibili.com/main/hotword"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, url, nil)
	if err != nil {
		return nil, err
	}
	var envelope searchHotListWebEnvelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, bilibili.NewDecodeError(method, url, body, reflect.TypeFor[searchHotListWebEnvelope](), "$", 0, err)
	}
	if envelope.Code != 0 {
		return nil, fmt.Errorf("%s %s: %w", method, url, bilibili.Error{Code: envelope.Code, Message: envelope.Message})
	}
	return &envelope.GetSearchHotListWebResult, nil
}

// searchHotListWebEnvelope 是热搜列表（web端）的根对象，业务字段直接位于根上。
type searchHotListWebEnvelope struct {
	GetSearchHotListWebResult

	Code    int    `json:"code"`
	Message string `json:"message"`
}
