package electric

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 微信/支付宝方式充电相关接口。各接口的返回体独立定义，互不复用。
// 文档见 electric/WeChat&Alipay.md。

// CheckElecPayOrderStatusParam 是检查扫码支付结果的请求参数。
type CheckElecPayOrderStatusParam struct {
	QRToken string `json:"qr_token" request:"query"` // 扫码秘钥
}

// CheckElecPayOrderStatusResult 是检查扫码支付结果的 data 对象。
type CheckElecPayOrderStatusResult struct {
	QRToken string  `json:"qr_token"` // 扫码秘钥
	OrderNo *string `json:"order_no"` // 留言token。未成功则无此项，用于添加充电留言
	Mid     int     `json:"mid"`      // 当前用户mid
	Status  *int    `json:"status"`   // 状态值。若秘钥错误则无此项。1：已支付。2：未扫描。3：未确认
}

// CheckElecPayOrderStatus 检查扫码支付结果。
func (c *Client) CheckElecPayOrderStatus(ctx context.Context, param CheckElecPayOrderStatusParam) (*CheckElecPayOrderStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/ugcpay/trade/elec/pay/order/status"
	)
	return bilibili.Execute[*CheckElecPayOrderStatusResult](ctx, c.Client, method, url, param)
}
