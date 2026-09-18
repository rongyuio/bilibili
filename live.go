package bilibili

import (
	"context"
	"crypto/rand"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
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

// liveMedalWebLocation 是已激活勋章接口默认的页面位置标识。
const liveMedalWebLocation = "0.0"

// GetLiveActivatedMedalInfoParam 指定直播间及主播。
type GetLiveActivatedMedalInfoParam struct {
	Platform    string `json:"platform"`                               // 平台，例如 pc
	RoomID      int    `json:"room_id"`                                // 直播间号
	TargetID    int64  `json:"target_id"`                              // 主播 UID
	WebLocation string `json:"web_location" request:"query,omitempty"` // 页面位置；留空时由库填入 liveMedalWebLocation
}

// GetLiveActivatedMedalInfo 获取已激活勋章及任务信息，CSRF 自动从请求 Cookie 快照填入 query。
func (c *Client) GetLiveActivatedMedalInfo(ctx context.Context, param GetLiveActivatedMedalInfoParam) (*GetLiveActivatedMedalInfoResult, error) {
	param.WebLocation = webLocationOrDefault(param.WebLocation, liveMedalWebLocation)
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

// GetLiveWebAreaList 获取 Web 端直播间分区列表（与 GetLiveAreaList 不同源）。
func (c *Client) GetLiveWebAreaList(ctx context.Context) (*LiveWebAreaList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-interface/v1/index/getWebAreaList"
	)
	return execute[*LiveWebAreaList](ctx, c, method, url, nil, fillParam("source_id", "2"))
}

// GetLiveAreaRoomListParam 指定获取分区房间列表的条件。
type GetLiveAreaRoomListParam struct {
	Platform     string `json:"platform" request:"default=web"` // 平台，一般为 web。留空时自动填 web
	ParentAreaID int64  `json:"parent_area_id"`                 // 父分区 id，来自 GetLiveWebAreaList
	AreaID       int64  `json:"area_id"`                        // 子分区 id。0 表示全部
	SortType     string `json:"sort_type"`                      // 排序方式；可为空，空值也会发送
	Page         int    `json:"page" request:"default=1"`       // 页码，从 1 开始。留空时自动填 1
}

// GetLiveAreaRoomList 获取直播二级分区的房间列表，WBI 签名。
// 注意：WBI 签名后按库约定会清空 Referer，但本接口实测在不携带直播域 Referer 时
// 会被风控拦截（-352），因此签名后重新补回直播域 Referer/Origin。
func (c *Client) GetLiveAreaRoomList(ctx context.Context, param GetLiveAreaRoomListParam) (*LiveAreaRoomList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-interface/v1/second/getList"
	)
	return execute[*LiveAreaRoomList](ctx, c, method, url, param, c.fillWbi(),
		func(r *resty.Request) error {
			r.SetHeader("Referer", "https://live.bilibili.com/")
			r.SetHeader("Origin", "https://live.bilibili.com")
			return nil
		})
}

// CheckLiveAnchorLotteryParam 指定要查询天选时刻的直播间。
type CheckLiveAnchorLotteryParam struct {
	RoomID int64 `json:"room_id" request:"field=roomid"` // 直播间号。接口参数名是 roomid（无下划线）
}

// CheckLiveAnchorLottery 查询直播间当前的天选时刻。
func (c *Client) CheckLiveAnchorLottery(ctx context.Context, param CheckLiveAnchorLotteryParam) (*CheckLiveAnchorLotteryResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Check"
	)
	return execute[*CheckLiveAnchorLotteryResult](ctx, c, method, url, param,
		func(r *resty.Request) error {
			r.SetHeader("Referer", "https://live.bilibili.com/")
			r.SetHeader("Origin", "https://live.bilibili.com")
			return nil
		})
}

// JoinLiveAnchorLotteryParam 指定参与天选时刻的抽奖信息，字段来自 CheckLiveAnchorLottery 的结果。
type JoinLiveAnchorLotteryParam struct {
	ID      int64  // 天选抽奖 id
	GiftID  int64  // 礼物 id
	GiftNum int    // 礼物数量
	VisitID string // 访问标识；留空时由库自动生成
}

// liveAnchorJoinForm 是参与天选时刻的表单字段。
type liveAnchorJoinForm struct {
	ID       int64  `json:"id"`
	GiftID   int64  `json:"gift_id"`
	GiftNum  int    `json:"gift_num"`
	VisitID  string `json:"visit_id"`
	Platform string `json:"platform"` // 固定为 pc
}

// JoinLiveAnchorLottery 参与天选时刻抽奖，CSRF 自动填入表单。
// 注意：要求赠礼的天选会消耗瓜子，参与前请自行检查 CheckLiveAnchorLotteryResult.GiftPrice。
func (c *Client) JoinLiveAnchorLottery(ctx context.Context, param JoinLiveAnchorLotteryParam) (*JoinLiveAnchorLotteryResult, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	if param.VisitID == "" {
		visitID, err := randomVisitID()
		if err != nil {
			return nil, err
		}
		param.VisitID = visitID
	}
	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/xlive/lottery-interface/v1/Anchor/Join"
	)
	// withParams 的默认分支已将 Content-Type 设为 application/x-www-form-urlencoded。
	return execute[*JoinLiveAnchorLotteryResult](ctx, c, method, url, liveAnchorJoinForm{
		ID:       param.ID,
		GiftID:   param.GiftID,
		GiftNum:  param.GiftNum,
		VisitID:  param.VisitID,
		Platform: "pc",
	}, moveFormParams("id", "gift_id", "gift_num", "visit_id", "platform"), fillFormCsrf(c))
}

// GetLiveNotice 获取直播公告。B 站会通过本接口的 Set-Cookie 下发 LIVE_BUVID，
// 响应 Cookie 已由请求链路自动合并进客户端，因此调用本方法即可完成补全。
func (c *Client) GetLiveNotice(ctx context.Context) error {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/news/v1/notice/recom"
	)
	_, err := execute[any](ctx, c, method, url, nil, fillParam("product", "live"))
	return err
}

// randomVisitID 生成 12 位随机访问标识：首位 1-9、中间 10 位随机小写字母数字、末尾固定 0。
func randomVisitID() (string, error) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 11)
	if _, err := rand.Read(b); err != nil {
		return "", errors.WithStack(err)
	}
	visitID := make([]byte, 12)
	visitID[0] = '1' + b[0]%9
	for i := range 10 {
		visitID[i+1] = charset[int(b[i+1])%len(charset)]
	}
	visitID[11] = '0'
	return string(visitID), nil
}
