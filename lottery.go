package bilibili

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"
)

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
			csrf := cookieValue(r.Cookies, "bili_jct")
			if csrf == "" {
				return errors.New("B站登录过期")
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
			csrf := cookieValue(r.Cookies, "bili_jct")
			if csrf == "" {
				return errors.New("B站登录过期")
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

// GetActivityLotteryTimesResult contains the corresponding lottery response fields.
type GetActivityLotteryTimesResult struct {
	Times         int         `json:"times"`
	LotteryType   int         `json:"lottery_type"`
	Points        int         `json:"points"`
	PointsPerTime int         `json:"points_per_time"`
	Intergral     interface{} `json:"intergral"`
	Stime         int         `json:"stime"`
	Etime         int         `json:"etime"`
}

// GetDynamicLotteryInfoResult contains the corresponding lottery response fields.
type GetDynamicLotteryInfoResult struct {
	LotteryID         int                          `json:"lottery_id"`
	SenderUID         int                          `json:"sender_uid"`
	BusinessType      int                          `json:"business_type"`
	BusinessID        int64                        `json:"business_id"`
	Status            int                          `json:"status"` // 2=已开奖
	LotteryTime       int                          `json:"lottery_time"`
	LotteryAtNum      int                          `json:"lottery_at_num"`
	LotteryFeedLimit  int                          `json:"lottery_feed_limit"`
	NeedPost          int                          `json:"need_post"`
	FirstPrize        int                          `json:"first_prize"`
	SecondPrize       int                          `json:"second_prize"`
	ThirdPrize        int                          `json:"third_prize"`
	Ts                int                          `json:"ts"`
	Participants      int                          `json:"participants"`
	HasChargeRight    bool                         `json:"has_charge_right"`
	Participated      bool                         `json:"participated"`
	Followed          bool                         `json:"followed"`
	Reposted          bool                         `json:"reposted"`
	LotteryDetailURL  string                       `json:"lottery_detail_url"`
	FirstPrizeCmt     string                       `json:"first_prize_cmt"`
	ThirdPrizeCmt     string                       `json:"third_prize_cmt"`
	FirstPrizePic     string                       `json:"first_prize_pic"`
	SecondPrizePic    string                       `json:"second_prize_pic"`
	ThirdPrizePic     string                       `json:"third_prize_pic"`
	VipBatchSign      string                       `json:"vip_batch_sign"`
	VipRedirectURL    string                       `json:"vip_redirect_url"`
	UpowerRedirectURL string                       `json:"upower_redirect_url"`
	PrizeTypeFirst    DynamicLotteryFirstPrizeType `json:"prize_type_first"`
	LotteryResult     DynamicLotteryResult         `json:"lottery_result"`
}

// DynamicLotteryFirstPrizeType contains the corresponding lottery response fields.
type DynamicLotteryFirstPrizeType struct {
	Type  int                      `json:"type"`
	Value DynamicLotteryPrizeValue `json:"value"`
}

// DynamicLotteryPrizeValue contains the corresponding lottery response fields.
type DynamicLotteryPrizeValue struct {
	Stype int `json:"stype"`
	Count int `json:"count"`
}

// DynamicLotteryResult contains the corresponding lottery response fields.
type DynamicLotteryResult struct {
	FirstPrizeResult []DynamicLotteryWinner `json:"first_prize_result"`
}

// DynamicLotteryWinner contains the corresponding lottery response fields.
type DynamicLotteryWinner struct {
	UID          int    `json:"uid"`
	Name         string `json:"name"`
	Face         string `json:"face"`
	HongbaoMoney int    `json:"hongbao_money"`
}
