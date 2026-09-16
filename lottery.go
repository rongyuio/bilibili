package bilibili

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// 活动抽奖相关接口。响应模型见 lottery_model.go。

// DoActivityLotteryParam 指定活动抽奖操作，所有字段进入 URL 编码表单。
type DoActivityLotteryParam struct {
	GaiaVtoken string `json:"gaia_vtoken"` // 风控验证 token；空字符串仍发送
	Num        int    `json:"num"`         // 本次抽奖次数
	PageID     string `json:"page_id"`     // 活动页面 ID
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

// dynamicLotteryWebLocation 是动态抽奖接口默认的页面位置标识。
const dynamicLotteryWebLocation = "333.1330"

// GetDynamicLotteryInfoParam 指定用户抽奖动态及请求来源。
type GetDynamicLotteryInfoParam struct {
	BusinessID   string `json:"business_id"`                            // 抽奖动态 ID
	BusinessType int    `json:"business_type"`                          // 业务类型，动态使用 1
	WebLocation  string `json:"web_location" request:"query,omitempty"` // 页面位置；留空时由库填入 dynamicLotteryWebLocation
}

// deviceReqJSON 按 webLocation 生成 x-bili-device-req-json，platform 与 device 固定为 web/pc。
// webLocation 经 JSON 转义后再拼接，避免引号或反斜杠破坏结构。
func deviceReqJSON(webLocation string) (string, error) {
	spmid, err := json.Marshal(webLocation)
	if err != nil {
		return "", err
	}
	return `{"platform":"web","device":"pc","spmid":` + string(spmid) + `}`, nil
}

// dynamicLotteryDeviceHandler 注入按 WebLocation 生成的设备信息 JSON。
func dynamicLotteryDeviceHandler(webLocation string) paramHandler {
	return func(r *resty.Request) error {
		body, err := deviceReqJSON(webLocation)
		if err != nil {
			return err
		}
		r.SetQueryParam("x-bili-device-req-json", body)
		return nil
	}
}

// GetDynamicLotteryInfo 查询用户抽奖动态的抽奖信息。
// x-bili-device-req-json 由库按 WebLocation 生成，调用方无需手写 JSON。
func (c *Client) GetDynamicLotteryInfo(ctx context.Context, param GetDynamicLotteryInfoParam) (*GetDynamicLotteryInfoResult, error) {
	param.WebLocation = webLocationOrDefault(param.WebLocation, dynamicLotteryWebLocation)
	return execute[*GetDynamicLotteryInfoResult](ctx, c, resty.MethodGet,
		"https://api.vc.bilibili.com/lottery_svr/v1/lottery_svr/lottery_notice", param,
		dynamicLotteryDeviceHandler(param.WebLocation))
}
