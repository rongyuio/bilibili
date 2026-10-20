package live

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播间用户相关接口。响应模型见 live_user_model.go，文档见 live/user.md。

// GetMyLiveMedalsParam 指定当前账号粉丝勋章的翻页条件。
type GetMyLiveMedalsParam struct {
	PageSize int `json:"page_size" request:"query"` // 每页数量。最大为 10，超出返回 1002002 参数异常
	Page     int `json:"page" request:"query"`      // 返回结果页数。从 1 开始；只填 Page 不填 PageSize 会 500
}

// GetMyLiveMedals 获取当前账号持有的粉丝勋章信息。
//
// 认证方式：Cookie（SESSDATA）或 APP。旧接口
// fans_medal/v5/live_fans_medal/iApiMedal 已不可用。
func (c *Client) GetMyLiveMedals(ctx context.Context, param GetMyLiveMedalsParam) (*GetMyLiveMedalsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-ucenter/v1/user/GetMyMedals"
	)
	return bilibili.Execute[*GetMyLiveMedalsResult](ctx, c.Client, method, url, param,
		liveReferer(0))
}

// GetLiveSignInfo 获取当前账号本月的直播签到信息。
//
// 认证方式：Cookie（SESSDATA）或 APP。
func (c *Client) GetLiveSignInfo(ctx context.Context) (*GetLiveSignInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/WebGetSignInfo"
	)
	return bilibili.Execute[*GetLiveSignInfoResult](ctx, c.Client, method, url, nil,
		liveReferer(0))
}

// GetLastMonthSignDays 获取当前账号上月的直播签到信息。
//
// 认证方式：Cookie（SESSDATA）或 APP。
func (c *Client) GetLastMonthSignDays(ctx context.Context) (*GetLastMonthSignDaysResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/sign/getLastMonthSignDays"
	)
	return bilibili.Execute[*GetLastMonthSignDaysResult](ctx, c.Client, method, url, nil,
		liveReferer(0))
}

// GetOnlineGoldRankParam 指定要查询的直播间贡献榜。
type GetOnlineGoldRankParam struct {
	RoomID   int64 `json:"roomId" request:"query"`   // 房间号
	Page     int   `json:"page" request:"query"`     // 返回结果页数
	PageSize int   `json:"pageSize" request:"query"` // 返回结果页大小。最大 50
	RUID     int64 `json:"ruid" request:"query"`     // 主播 uid
}

// GetOnlineGoldRank 查询直播间的贡献榜（在线榜）。
//
// 认证方式：无。文档中的 URL 带双斜杠（api.live.bilibili.com//xlive/...），
// 这里归一为单斜杠。
func (c *Client) GetOnlineGoldRank(ctx context.Context, param GetOnlineGoldRankParam) (*GetOnlineGoldRankResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/general-interface/v1/rank/getOnlineGoldRank"
	)
	return bilibili.Execute[*GetOnlineGoldRankResult](ctx, c.Client, method, url, param,
		liveReferer(param.RoomID))
}

// GetLiveGuardActiveParam 指定要查询观看时长的主播。
type GetLiveGuardActiveParam struct {
	Platform string `json:"platform" request:"query"` // 平台，固定为 android
	RUID     int64  `json:"ruid" request:"query"`     // 主播 uid
}

// GetLiveGuardActive 查询当前账号在某直播间的观看时长与大航海信息。
//
// 认证方式：Cookie（SESSDATA）或 APP。
func (c *Client) GetLiveGuardActive(ctx context.Context, param GetLiveGuardActiveParam) (*GetLiveGuardActiveResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/general-interface/v1/guard/GuardActive"
	)
	return bilibili.Execute[*GetLiveGuardActiveResult](ctx, c.Client, method, url, param,
		liveReferer(0))
}

// GetLiveUserCardParam 指定要查询的用户及其所在直播间的主播。
type GetLiveUserCardParam struct {
	UID  int64 `json:"uid" request:"query"`  // 目标用户 uid
	RUID int64 `json:"ruid" request:"query"` // 主播 uid
}

// GetLiveUserCard 查询某用户在直播间内的信息（直播间用户卡片）。
//
// 认证方式：无。目标用户隐藏粉丝牌时 FansMedal 为 nil。
func (c *Client) GetLiveUserCard(ctx context.Context, param GetLiveUserCardParam) (*GetLiveUserCardResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-ucenter/v2/card/user"
	)
	return bilibili.Execute[*GetLiveUserCardResult](ctx, c.Client, method, url, param,
		liveReferer(0))
}
