package lottery

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 活动抽奖相关接口。各接口的返回体独立定义，互不复用。

// DoActivityLotteryParam 指定活动抽奖操作，所有字段进入 URL 编码表单。
type DoActivityLotteryParam struct {
	GaiaVtoken string `json:"gaia_vtoken"` // 风控验证 token；空字符串仍发送
	Num        int    `json:"num"`         // 本次抽奖次数
	PageID     string `json:"page_id"`     // 活动页面 ID
	Sid        string `json:"sid"`         // 活动抽奖配置 ID
}

// formContentType 显式声明 URL 编码表单请求体。
func formContentType(r *resty.Request) error {
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return nil
}

// DoActivityLottery 执行活动抽奖，自动填入 CSRF，不自动重试或解析中奖明细。
func (c *Client) DoActivityLottery(ctx context.Context, param DoActivityLotteryParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/lottery/x/do"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("gaia_vtoken", "num", "page_id", "sid"),
		bilibili.FillFormCsrf(c.Client),
		formContentType,
	)
	return err
}

// GetActivityLotteryTimesParam 指定活动抽奖配置。
type GetActivityLotteryTimesParam struct {
	Sid string `json:"sid"` // 活动抽奖配置 ID
}

// GetActivityLotteryTimesResult 是活动剩余抽奖次数的 data 对象。
type GetActivityLotteryTimesResult struct {
	Times         int `json:"times"`           // 剩余抽奖次数
	LotteryType   int `json:"lottery_type"`    // 抽奖类型
	Points        int `json:"points"`          // 积分
	PointsPerTime int `json:"points_per_time"` // 每次抽奖消耗积分
	Intergral     any `json:"intergral"`       // 作用尚不明确
	Stime         int `json:"stime"`           // 开始时间。UNIX 秒级时间戳
	Etime         int `json:"etime"`           // 结束时间。UNIX 秒级时间戳
}

// GetActivityLotteryTimes 获取当前账号的活动剩余抽奖次数，CSRF 自动填入 query。
func (c *Client) GetActivityLotteryTimes(ctx context.Context, param GetActivityLotteryTimesParam) (*GetActivityLotteryTimesResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/lottery/x/mytimes"
	)
	return bilibili.Execute[*GetActivityLotteryTimesResult](ctx, c.Client, method, url, param,
		bilibili.FillCsrf(c.Client))
}

// dynamicLotteryWebLocation 是动态抽奖接口默认的页面位置标识。
const dynamicLotteryWebLocation = "333.1330"

// GetDynamicLotteryInfoParam 指定用户抽奖动态及请求来源。
type GetDynamicLotteryInfoParam struct {
	BusinessID   string `json:"business_id"`                            // 抽奖动态 ID
	BusinessType int    `json:"business_type"`                          // 业务类型，动态使用 1
	WebLocation  string `json:"web_location" request:"query,omitempty"` // 页面位置；留空时由库填入 dynamicLotteryWebLocation
}

// GetDynamicLotteryInfoResult 是动态抽奖信息的 data 对象。
type GetDynamicLotteryInfoResult struct {
	LotteryID         int                                 `json:"lottery_id"`          // 抽奖 id
	SenderUID         int                                 `json:"sender_uid"`          // 发送者 mid
	BusinessType      int                                 `json:"business_type"`       // 业务类型。1：动态抽奖
	BusinessID        int64                               `json:"business_id"`         // 业务 id。可能与请求参数不同
	Status            int                                 `json:"status"`              // 抽奖状态。0：未开奖，2：已开奖
	LotteryTime       int                                 `json:"lottery_time"`        // 开奖时间。UNIX 秒级时间戳
	LotteryAtNum      int                                 `json:"lottery_at_num"`      // at 人数?
	LotteryFeedLimit  int                                 `json:"lottery_feed_limit"`  // 转发限制?
	NeedPost          int                                 `json:"need_post"`           // 是否需要发布?
	FirstPrize        int                                 `json:"first_prize"`         // 一等奖数量
	SecondPrize       int                                 `json:"second_prize"`        // 二等奖数量
	ThirdPrize        int                                 `json:"third_prize"`         // 三等奖数量
	Ts                int                                 `json:"ts"`                  // 发布时间。UNIX 秒级时间戳
	Participants      int                                 `json:"participants"`        // 参与人数
	HasChargeRight    bool                                `json:"has_charge_right"`    // 是否有充值权?
	Participated      bool                                `json:"participated"`        // 是否参与过?
	Followed          bool                                `json:"followed"`            // 是否关注过
	Reposted          bool                                `json:"reposted"`            // 是否转发过
	LotteryDetailURL  string                              `json:"lottery_detail_url"`  // 抽奖详情页 URL?
	FirstPrizeCmt     string                              `json:"first_prize_cmt"`     // 一等奖奖品名称
	ThirdPrizeCmt     string                              `json:"third_prize_cmt"`     // 三等奖奖品名称
	FirstPrizePic     string                              `json:"first_prize_pic"`     // 一等奖奖品图片 URL
	SecondPrizePic    string                              `json:"second_prize_pic"`    // 二等奖奖品图片 URL
	ThirdPrizePic     string                              `json:"third_prize_pic"`     // 三等奖奖品图片 URL
	VipBatchSign      string                              `json:"vip_batch_sign"`      // 作用尚不明确
	VipRedirectURL    string                              `json:"vip_redirect_url"`    // 作用尚不明确
	UpowerRedirectURL string                              `json:"upower_redirect_url"` // 作用尚不明确
	PrizeTypeFirst    GetDynamicLotteryInfoPrizeType      `json:"prize_type_first"`    // 一等奖奖品类型
	LotteryResult     *GetDynamicLotteryInfoLotteryResult `json:"lottery_result"`      // 抽奖结果。无则不存在
}

// GetDynamicLotteryInfoPrizeType 是一等奖奖品类型。
type GetDynamicLotteryInfoPrizeType struct {
	Type  int                             `json:"type"`  // 奖品类型?
	Value GetDynamicLotteryInfoPrizeValue `json:"value"` // 奖品价值?
}

// GetDynamicLotteryInfoPrizeValue 是奖品价值。
type GetDynamicLotteryInfoPrizeValue struct {
	Stype int `json:"stype"` // 子奖品类型?
	Count int `json:"count"` // 数量
}

// GetDynamicLotteryInfoLotteryResult 是抽奖结果。
type GetDynamicLotteryInfoLotteryResult struct {
	FirstPrizeResult []GetDynamicLotteryInfoWinner `json:"first_prize_result"` // 一等奖结果
}

// GetDynamicLotteryInfoWinner 是一名中奖者。
type GetDynamicLotteryInfoWinner struct {
	UID          int    `json:"uid"`           // 中奖者 mid
	Name         string `json:"name"`          // 中奖者名称
	Face         string `json:"face"`          // 中奖者头像 URL
	HongbaoMoney int    `json:"hongbao_money"` // 奖品金额?
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

// webLocationOrDefault 在调用方未指定页面标识时返回接口默认值。
func webLocationOrDefault(webLocation, fallback string) string {
	if webLocation == "" {
		return fallback
	}
	return webLocation
}

// dynamicLotteryDeviceHandler 注入按 WebLocation 生成的设备信息 JSON。
func dynamicLotteryDeviceHandler(webLocation string) bilibili.Handler {
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
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/lottery_svr/v1/lottery_svr/lottery_notice"
	)
	param.WebLocation = webLocationOrDefault(param.WebLocation, dynamicLotteryWebLocation)
	return bilibili.Execute[*GetDynamicLotteryInfoResult](ctx, c.Client, method, url, param,
		dynamicLotteryDeviceHandler(param.WebLocation))
}
