package bilibili

import (
	"context"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// 直播相关接口。响应模型见 live_model.go。

// ReportLiveLikeParam 指定直播点赞上报参数，所有字段进入 URL 编码的表单。
type ReportLiveLikeParam struct {
	ClickTime int   `json:"click_time"` // 上报数量，由调用方指定
	RoomID    int   `json:"room_id"`    // 直播间号
	AnchorID  int64 `json:"anchor_id"`  // 主播 UID
	UID       int   `json:"uid"`        // 当前账号 UID
}

// ReportLiveLike 上报直播点赞，CSRF 自动从请求 Cookie 快照填入表单，不自动重试。
func (c *Client) ReportLiveLike(ctx context.Context, param ReportLiveLikeParam) error {
	_, err := execute[any](ctx, c, resty.MethodPost,
		"https://api.live.bilibili.com/xlive/app-ucenter/v1/like_info_v3/like/likeReportV3", param,
		func(r *resty.Request) error {
			csrf, err := csrfValue(r)
			if err != nil {
				return err
			}
			// encodeParams 已完成字段编码；仅将本接口字段移入表单，保留客户端默认 query。
			for _, key := range []string{"click_time", "room_id", "anchor_id", "uid"} {
				r.FormData[key] = append([]string(nil), r.QueryParam[key]...)
				r.QueryParam.Del(key)
			}
			r.SetFormData(map[string]string{"csrf": csrf})
			r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			return nil
		})
	return err
}

// GetLiveMedalWallParam 指定需要查询勋章墙的用户。
type GetLiveMedalWallParam struct {
	TargetID int `json:"target_id"` // 用户 UID
}

// GetLiveMedalWall 获取用户的直播勋章墙。
func (c *Client) GetLiveMedalWall(ctx context.Context, param GetLiveMedalWallParam) (*GetLiveMedalWallResult, error) {
	return execute[*GetLiveMedalWallResult](ctx, c, resty.MethodGet,
		"https://api.live.bilibili.com/xlive/web-ucenter/user/MedalWall", param)
}

// GetLiveFansMedalPanelParam 指定勋章面板的页码与每页数量。
type GetLiveFansMedalPanelParam struct {
	Page     int `json:"page"`      // 页码，从 1 开始
	PageSize int `json:"page_size"` // 每页数量
}

// GetLiveFansMedalPanel 获取当前账号的一页直播勋章面板，不自动翻页。
func (c *Client) GetLiveFansMedalPanel(ctx context.Context, param GetLiveFansMedalPanelParam) (*GetLiveFansMedalPanelResult, error) {
	return execute[*GetLiveFansMedalPanelResult](ctx, c, resty.MethodGet,
		"https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/panel", param)
}

// GetLiveActivatedMedalInfoParam 指定直播间及主播。
type GetLiveActivatedMedalInfoParam struct {
	Platform    string `json:"platform"`     // 平台，例如 pc
	RoomID      int    `json:"room_id"`      // 直播间号
	TargetID    int64  `json:"target_id"`    // 主播 UID
	WebLocation string `json:"web_location"` // 页面位置，例如 0.0
}

// GetLiveActivatedMedalInfo 获取已激活勋章及任务信息，CSRF 自动从请求 Cookie 快照填入 query。
func (c *Client) GetLiveActivatedMedalInfo(ctx context.Context, param GetLiveActivatedMedalInfoParam) (*GetLiveActivatedMedalInfoResult, error) {
	return execute[*GetLiveActivatedMedalInfoResult](ctx, c, resty.MethodGet,
		"https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/GetActivatedMedalInfo", param,
		func(r *resty.Request) error {
			csrf, err := csrfValue(r)
			if err != nil {
				return err
			}
			r.SetQueryParam("csrf", csrf)
			return nil
		})
}

type GetLiveRoomInfoParam struct {
	RoomID int `json:"room_id"` // 直播间号。可以为短号
}

// GetLiveRoomInfo 获取直播间信息
func (c *Client) GetLiveRoomInfo(ctx context.Context, param GetLiveRoomInfoParam) (*LiveRoomInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Room/get_info"
	)
	return execute[*LiveRoomInfo](ctx, c, method, url, param)
}

type UpdateLiveRoomTitleParam struct {
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识
	VisitID  string `json:"visit_id,omitempty" request:"query,omitempty"` // (?)。某种标识？
	RoomID   int    `json:"room_id"`                                      // 直播间id。必须为自己的直播间id
	Title    string `json:"title,omitempty" request:"query,omitempty"`    // 直播间标题。上限40个字符
	AreaID   int    `json:"area_id,omitempty" request:"query,omitempty"`  // 直播分区id（子分区id）。详见[直播分区](live_area.md)
	AddTag   string `json:"add_tag,omitempty" request:"query,omitempty"`  // 要添加的标签。开播设置界面上限10个字符
	DelTag   string `json:"del_tag,omitempty" request:"query,omitempty"`  // 要删除的标签。若存在add_tag时不起作用
}

// UpdateLiveRoomTitle 更新直播间信息
func (c *Client) UpdateLiveRoomTitle(ctx context.Context, param UpdateLiveRoomTitleParam) (*UpdateLiveRoomTitleResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/room/v1/Room/update"
	)
	return execute[*UpdateLiveRoomTitleResult](ctx, c, method, url, param, fillCsrf(c))
}

type StartLiveParam struct {
	RoomID   int    `json:"room_id"`  // 直播间id。必须为自己的直播间id
	AreaV2   int    `json:"area_v2"`  // 直播分区id（子分区id）。详见[直播分区]
	Platform string `json:"platform"` // 直播平台。直播姬（pc）：pc_link。web在线直播：web_link（已下线）。bililink：android_link。

	// 下面四个参数详见：https://github.com/SocialSisterYi/bilibili-API-collect/pull/1351/files 。
	// 可以调用 GetHomePageLiveVersion 方法获取 Version 和 Build 参数。
	Version string `json:"version"`          // 直播姬版本号，2025.7.20后对于某些用户必填
	Build   int    `json:"build"`            // 直播姬构建号，2025.7.20后对于某些用户必填
	Appkey  string `json:"appkey,omitempty"` // APP密钥，不填会自动计算
	Sign    string `json:"sign,omitempty"`   // APP API签名得到的sign，不填会自动计算

	Ts int `json:"ts,omitempty" request:"query,omitempty"` // 10位时间戳
}

// StartLive 开始直播
//
// 注意：为了方便使用，这个函数的很多参数做了自动填入，使用例子可以参考：https://github.com/CuteReimu/bilibili/issues/121
func (c *Client) StartLive(ctx context.Context, param StartLiveParam) (*StartLiveResult, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	r := c.newRequest(ctx)
	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/room/v1/Room/startLive"
	)

	// 未提供签名时按直播姬协议自动计算，签名依赖 Cookie 快照中的 CSRF。
	if param.Sign == "" && param.Appkey == "" {
		param = signStartLiveParam(param, cookieValue(r.Cookies, "bili_jct"))
	}

	return executeRequest[*StartLiveResult](c, r, method, url, param, fillCsrf(c))
}

// signStartLiveParam fills the appkey/sign pair required by the startLive
// signature protocol, using public live-client credentials.
func signStartLiveParam(param StartLiveParam, csrf string) StartLiveParam {
	// 已知的 Bilibili 直播姬的密钥和对应的秘钥。
	// 这些是公开的常量，用于计算 API 签名。
	const (
		appKey    = "aae92bc66f3edfab"
		appSecret = "af125a0d5279fd576c1b4418a3e8276d" //nolint:gosec
	)
	if param.Ts == 0 {
		param.Ts = int(time.Now().Unix())
	}
	param.Appkey = appKey
	param.Sign = calculateAppSign(map[string]string{
		"appkey":     param.Appkey,
		"build":      strconv.Itoa(param.Build),
		"platform":   param.Platform,
		"room_id":    strconv.Itoa(param.RoomID),
		"area_v2":    strconv.Itoa(param.AreaV2),
		"ts":         strconv.Itoa(param.Ts),
		"version":    param.Version,
		"csrf":       csrf,
		"csrf_token": csrf,
	}, appSecret)
	return param
}

type StopLiveParam struct {
	Platform string `json:"platform"` // 直播平台。直播姬（pc）：pc_link。web在线直播：web_link（已下线）。bililink：android_link。
	RoomID   int    `json:"room_id"`  // 直播间id。必须为自己的直播间id
}

// StopLive 关闭直播
func (c *Client) StopLive(ctx context.Context, param StopLiveParam) (*StopLiveResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/room/v1/Room/stopLive"
	)
	return execute[*StopLiveResult](ctx, c, method, url, param, fillCsrf(c))
}

// GetLiveAreaList 获取全部直播间分区列表
func (c *Client) GetLiveAreaList(ctx context.Context) ([]LiveAreaList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Area/getList"
	)
	return execute[[]LiveAreaList](ctx, c, method, url, nil)
}

type GetHomePageLiveVersionParam struct {
	SystemVersion int `json:"system_version"`                         // 暂不清楚。可以直接写2
	Ts            int `json:"ts,omitempty" request:"query,omitempty"` // 10位时间戳
}

// GetHomePageLiveVersion PC直播姬版本号获取
func (c *Client) GetHomePageLiveVersion(ctx context.Context, param GetHomePageLiveVersionParam) (*HomePageLiveVersion, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/liveVersionInfo/getHomePageLiveVersion"
	)
	return execute[*HomePageLiveVersion](ctx, c, method, url, param)
}
