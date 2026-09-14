package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 活动抽奖相关接口。响应模型见 lottery_model.go。

// DoActivityLotteryParam 指定活动抽奖操作，所有字段进入 URL 编码表单。
type DoActivityLotteryParam struct {
	GaiaVtoken string `json:"gaia_vtoken"` // 风控验证 token；空字符串仍发送
	Num        int    `json:"num"`         // 本次抽奖次数
	PageId     string `json:"page_id"`     // 活动页面 ID
	Sid        string `json:"sid"`         // 活动抽奖配置 ID
}

// DoActivityLottery 执行活动抽奖，自动填入 CSRF，不自动重试或解析中奖明细。
func (c *Client) DoActivityLottery(ctx context.Context, param DoActivityLotteryParam) error {
	_, err := execute[any](ctx, c, resty.MethodPost,
		"https://api.bilibili.com/x/lottery/x/do", param, func(r *resty.Request) error {
			csrf, err := csrfValue(r)
			if err != nil {
				return err
			}
			// 复用统一参数编码，再将本接口字段移入表单。
			for _, key := range []string{"gaia_vtoken", "num", "page_id", "sid"} {
				r.FormData[key] = append([]string(nil), r.QueryParam[key]...)
				r.QueryParam.Del(key)
			}
			r.SetFormData(map[string]string{"csrf": csrf})
			r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			return nil
		})
	return err
}

// GetActivityLotteryTimesParam 指定活动抽奖配置。
type GetActivityLotteryTimesParam struct {
	Sid string `json:"sid"` // 活动抽奖配置 ID
}

// GetActivityLotteryTimes 获取当前账号的活动剩余抽奖次数，CSRF 自动填入 query。
func (c *Client) GetActivityLotteryTimes(ctx context.Context, param GetActivityLotteryTimesParam) (*GetActivityLotteryTimesResult, error) {
	return execute[*GetActivityLotteryTimesResult](ctx, c, resty.MethodGet,
		"https://api.bilibili.com/x/lottery/x/mytimes", param, func(r *resty.Request) error {
			csrf, err := csrfValue(r)
			if err != nil {
				return err
			}
			r.SetQueryParam("csrf", csrf)
			return nil
		})
}

// GetDynamicLotteryInfoParam 指定用户抽奖动态及请求来源。
type GetDynamicLotteryInfoParam struct {
	BusinessId    string `json:"business_id"`            // 抽奖动态 ID
	BusinessType  int    `json:"business_type"`          // 业务类型，动态使用 1
	WebLocation   string `json:"web_location"`           // 页面位置，例如 333.1330
	DeviceReqJSON string `json:"x-bili-device-req-json"` // 设备请求信息 JSON 字符串，进入 query
}

// GetDynamicLotteryInfo 查询用户抽奖动态的抽奖信息
func (c *Client) GetDynamicLotteryInfo(ctx context.Context, param GetDynamicLotteryInfoParam) (*GetDynamicLotteryInfoResult, error) {
	return execute[*GetDynamicLotteryInfoResult](ctx, c, resty.MethodGet,
		"https://api.vc.bilibili.com/lottery_svr/v1/lottery_svr/lottery_notice", param)
}
