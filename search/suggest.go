package search

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 搜索建议接口。响应模型见 model_suggest.go。

// GetSearchSuggestParam 指定获取搜索建议关键词的参数。
type GetSearchSuggestParam struct {
	Term          string  `json:"term" request:"query"`                          // 需要获得建议的输入内容
	MainVer       string  `json:"main_ver" request:"query,default=v1"`           // 版本。默认为 v1
	Highlight     string  `json:"highlight,omitempty" request:"query,omitempty"` // 任意值，无明显作用。默认为空
	Func          string  `json:"func" request:"query,default=suggest"`          // 函数。默认为 suggest
	SuggestType   string  `json:"suggest_type" request:"query,default=accurate"` // 建议类型。默认为 accurate
	SubType       string  `json:"sub_type" request:"query,default=tag"`          // 子类型。默认为 tag
	UserID        int64   `json:"userid,omitempty" request:"query,omitempty"`    // 本用户 mid。可能用于个性化推荐
	BangumiAccNum int     `json:"bangumi_acc_num" request:"query,default=1"`     // 番剧累积数。默认为 1
	SpecialAccNum int     `json:"special_acc_num" request:"query,default=1"`     // 特殊累积数。默认为 1
	TopicAccNum   int     `json:"topic_acc_num" request:"query,default=1"`       // 话题累积数。默认为 1
	UpuserAccNum  int     `json:"upuser_acc_num" request:"query,default=1"`      // UP主累积数。默认为 1
	TagNum        int     `json:"tag_num" request:"query,default=10"`            // Tag 数。默认为 10
	SpecialNum    int     `json:"special_num" request:"query,default=10"`        // 特殊推荐数。默认为 10
	BangumiNum    int     `json:"bangumi_num" request:"query,default=10"`        // 番剧推荐数。默认为 10
	UpuserNum     int     `json:"upuser_num" request:"query,default=3"`          // UP主推荐数。默认为 3
	Rnd           float64 `json:"rnd,omitempty" request:"query,omitempty"`       // 一个随机浮点数，由 Math.random() 生成
	Buvid         string  `json:"buvid,omitempty" request:"query,omitempty"`     // 同 Cookie 中的 buvid3
	Spmid         string  `json:"spmid" request:"query,default=333.1007"`        // 页面标识。默认为 333.1007
}

// searchSuggestResponse 是搜索建议接口的根对象，结果位于 result。
// 该接口不使用标准的 code/message/data 信封，因此不能复用 Execute。
type searchSuggestResponse struct {
	Code   int                     `json:"code"`   // 返回值。0：成功
	Result *GetSearchSuggestResult `json:"result"` // 搜索建议结果
}

// searchSuggestNoResultCode 是接口在没有任何候选关键词时返回的业务码。
// 文档只标注了 0（成功）；实测无结果时返回 3 且 result.tag 为空数组，因此按空结果处理。
const searchSuggestNoResultCode = 3

// GetSearchSuggest 获取搜索建议关键词（web端）。
//
// 搜索建议最多提供 10 个候选关键词，内容可为任意语言，中文拼音支持全拼联想。
// 该接口的响应把结果放在根对象的 result 字段而非标准的 data 字段，
// 故直接走 ExecuteRaw 取原始响应体并手动解码。
func (c *Client) GetSearchSuggest(ctx context.Context, param GetSearchSuggestParam) (*GetSearchSuggestResult, error) {
	const (
		method   = resty.MethodGet
		endpoint = "https://s.search.bilibili.com/main/suggest"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, endpoint, param)
	if err != nil {
		return nil, err
	}
	var envelope searchSuggestResponse
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, bilibili.NewDecodeError(method, endpoint, body, reflect.TypeFor[searchSuggestResponse](), "$", 0, err)
	}
	if envelope.Code != 0 && envelope.Code != searchSuggestNoResultCode {
		return nil, fmt.Errorf("%s %s: %w", method, endpoint, bilibili.Error{Code: envelope.Code})
	}
	if envelope.Result == nil {
		return &GetSearchSuggestResult{}, nil
	}
	return envelope.Result, nil
}
