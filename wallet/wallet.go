package wallet

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetUserWalletParam 是获取用户钱包的请求参数。
type GetUserWalletParam struct {
	PanelType    int    `json:"panelType,omitempty" request:"json,omitempty"` // 面板类型。默认为 3
	PlatformType int    `json:"platformType" request:"json"`                  // 平台类型。任意有效数字，默认为 3
	Timestamp    int64  `json:"timestamp,omitempty" request:"json,omitempty"` // 当前时间。UNIX 毫秒时间戳
	TraceID      int64  `json:"traceId,omitempty" request:"json,omitempty"`   // 追踪 ID。与 timestamp 值相同
	Version      string `json:"version,omitempty" request:"json,omitempty"`   // 版本。默认 1.0
}

// GetUserWalletResult 是获取用户钱包接口的 data 对象。
type GetUserWalletResult struct {
	Mid                  int64  `json:"mid"`                  // 用户 mid
	TotalBp              int64  `json:"totalBp"`              // 总计 B 币
	DefaultBp            int64  `json:"defaultBp"`            // 默认 B 币
	IosBp                int64  `json:"iosBp"`                // iOS B 币
	CouponBalance        int64  `json:"couponBalance"`        // 优惠券余额
	AvailableBp          int64  `json:"availableBp"`          // 可用 B 币
	UnavailableBp        int64  `json:"unavailableBp"`        // 不可用 B 币
	UnavailableReason    string `json:"unavailableReason"`    // 不可用原因
	Tip                  string `json:"tip"`                  // 提示信息
	NeedShowClassBalance int    `json:"needShowClassBalance"` // 需要显示类余额
}

// GetUserWallet 获取用户钱包信息。
func (c *Client) GetUserWallet(ctx context.Context, param GetUserWalletParam) (*GetUserWalletResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://pay.bilibili.com/paywallet/wallet/getUserWallet"
	)
	return bilibili.Execute[*GetUserWalletResult](ctx, c.Client, method, url, param)
}
