package opus

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 空间图文接口。响应模型见 opus_space_model.go。

// GetUserSpaceOpusParam 指定要查询的用户空间图文。
type GetUserSpaceOpusParam struct {
	HostMid     int64  `json:"host_mid" request:"query"`                         // 被获取用户的 mid (UID)，必填
	Page        int    `json:"page,omitempty" request:"query,omitempty"`         // 假装自己是分页信息。以 1 开始
	Offset      string `json:"offset,omitempty" request:"query,omitempty"`       // 偏移信息。取上一次返回的 offset（同时也是最后一条的 opus id），默认为空
	Type        string `json:"type,omitempty" request:"query,omitempty"`         // 类型。all：全部（默认）。article：专栏。dynamic：动态
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 页面位置标识。网页端为 333.1387
}

// GetUserSpaceOpus 获取用户空间图文。
//
// 该接口只能获取图文信息，无法获取转发动态一类的内容。
// w_rid / wts 由 WBI 签名生成，文档标注为非必要，因此此处未启用签名。
func (c *Client) GetUserSpaceOpus(ctx context.Context, param GetUserSpaceOpusParam) (*GetUserSpaceOpusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/opus/feed/space"
	)
	return bilibili.Execute[*GetUserSpaceOpusResult](ctx, c.Client, method, url, param)
}
