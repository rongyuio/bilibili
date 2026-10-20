package live

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 直播相关接口。响应模型见 live_model.go。

// errNilContext 表示调用方传入了 nil context。
var errNilContext = errors.New("request context is nil")

// checkContext 在任何请求准备开始前拒绝 nil / 已结束的 context。
func checkContext(ctx context.Context) error {
	if ctx == nil {
		return errNilContext
	}
	return ctx.Err()
}

// ReportLiveLikeParam 指定直播点赞上报参数，所有字段进入 URL 编码的表单。
type ReportLiveLikeParam struct {
	ClickTime int   `json:"click_time"` // 上报数量，由调用方指定
	RoomID    int   `json:"room_id"`    // 直播间号
	AnchorID  int64 `json:"anchor_id"`  // 主播 UID
	UID       int   `json:"uid"`        // 当前账号 UID
}

// ReportLiveLike 上报直播点赞，CSRF 自动从请求 Cookie 快照填入表单，不自动重试。
func (c *Client) ReportLiveLike(ctx context.Context, param ReportLiveLikeParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/xlive/app-ucenter/v1/like_info_v3/like/likeReportV3"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("click_time", "room_id", "anchor_id", "uid"),
		bilibili.FillFormCsrf(c.Client),
		liveReferer(int64(param.RoomID)))
	return err
}

// GetLiveMedalWallParam 指定需要查询勋章墙的用户。
type GetLiveMedalWallParam struct {
	TargetID int `json:"target_id"` // 用户 UID
}

// GetLiveMedalWall 获取用户的直播勋章墙。
func (c *Client) GetLiveMedalWall(ctx context.Context, param GetLiveMedalWallParam) (*GetLiveMedalWallResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-ucenter/user/MedalWall"
	)
	return bilibili.Execute[*GetLiveMedalWallResult](ctx, c.Client, method, url, param)
}

// GetLiveFansMedalPanelParam 指定勋章面板的页码与每页数量。
type GetLiveFansMedalPanelParam struct {
	Page     int `json:"page"`      // 页码，从 1 开始
	PageSize int `json:"page_size"` // 每页数量
}

// GetLiveFansMedalPanel 获取当前账号的一页直播勋章面板，不自动翻页。
func (c *Client) GetLiveFansMedalPanel(ctx context.Context, param GetLiveFansMedalPanelParam) (*GetLiveFansMedalPanelResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/panel"
	)
	return bilibili.Execute[*GetLiveFansMedalPanelResult](ctx, c.Client, method, url, param)
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
	if param.WebLocation == "" {
		param.WebLocation = liveMedalWebLocation
	}
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-ucenter/v1/fansMedal/GetActivatedMedalInfo"
	)
	return bilibili.Execute[*GetLiveActivatedMedalInfoResult](ctx, c.Client, method, url, param,
		bilibili.FillCsrf(c.Client))
}

// GetLiveRoomInfoParam 指定要查询的直播间。
type GetLiveRoomInfoParam struct {
	RoomID int `json:"room_id"` // 直播间号。可以为短号
}

// GetLiveRoomInfo 获取直播间信息
func (c *Client) GetLiveRoomInfo(ctx context.Context, param GetLiveRoomInfoParam) (*GetLiveRoomInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Room/get_info"
	)
	return bilibili.Execute[*GetLiveRoomInfoResult](ctx, c.Client, method, url, param)
}

// UpdateLiveRoomTitleParam 指定直播间信息更新参数。
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
	return bilibili.Execute[*UpdateLiveRoomTitleResult](ctx, c.Client, method, url, param,
		bilibili.FillCsrf(c.Client))
}

// StartLiveParam 指定开播参数。
type StartLiveParam struct {
	RoomID   int    `json:"room_id"`  // 直播间id。必须为自己的直播间id
	AreaV2   int    `json:"area_v2"`  // 直播分区id（子分区id）。详见[直播分区]
	Platform string `json:"platform"` // 直播平台。直播姬（pc）：pc_link。web在线直播：web_link（已下线）。bililink：android_link。

	// 下面四个参数的说明见接口文档。
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

	// 未提供签名时按直播姬协议自动计算，签名依赖 Cookie 快照中的 CSRF。
	if param.Sign == "" && param.Appkey == "" {
		param = signStartLiveParam(param, cookieValue(c.GetCookies(), "bili_jct"))
	}

	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/room/v1/Room/startLive"
	)
	return bilibili.Execute[*StartLiveResult](ctx, c.Client, method, url, param,
		bilibili.FillCsrf(c.Client))
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

// calculateAppSign 计算 APP API 签名
// 按照 Bilibili APP API 签名算法：参数按 key 排序后拼接，加上秘钥后计算 MD5
func calculateAppSign(params map[string]string, appSecret string) string {
	// 收集所有非空参数
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v != "" {
			keys = append(keys, k)
		}
	}

	// 按 key 排序
	slices.Sort(keys)

	// 构建查询字符串
	query := url.Values{}
	for _, k := range keys {
		if params[k] != "" {
			query.Set(k, params[k])
		}
	}

	// 拼接参数和秘钥
	signStr := query.Encode() + appSecret

	// 计算 MD5
	hash := md5.Sum([]byte(signStr))
	return hex.EncodeToString(hash[:])
}

// cookieValue 从 Cookie 列表里按名称取值。
func cookieValue(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie != nil && cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

// StopLiveParam 指定关播参数。
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
	return bilibili.Execute[*StopLiveResult](ctx, c.Client, method, url, param,
		bilibili.FillCsrf(c.Client))
}

// GetLiveAreaList 获取全部直播间分区列表
func (c *Client) GetLiveAreaList(ctx context.Context) ([]GetLiveAreaListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/room/v1/Area/getList"
	)
	return bilibili.Execute[[]GetLiveAreaListItem](ctx, c.Client, method, url, nil)
}

// GetHomePageLiveVersionParam 指定直播姬版本查询参数。
type GetHomePageLiveVersionParam struct {
	SystemVersion int `json:"system_version"`                         // 暂不清楚。可以直接写2
	Ts            int `json:"ts,omitempty" request:"query,omitempty"` // 10位时间戳
}

// GetHomePageLiveVersion PC直播姬版本号获取
func (c *Client) GetHomePageLiveVersion(ctx context.Context, param GetHomePageLiveVersionParam) (*GetHomePageLiveVersionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/app-blink/v1/liveVersionInfo/getHomePageLiveVersion"
	)
	return bilibili.Execute[*GetHomePageLiveVersionResult](ctx, c.Client, method, url, param)
}

// GetLiveWebAreaList 获取 Web 端直播间分区列表（与 GetLiveAreaList 不同源）。
func (c *Client) GetLiveWebAreaList(ctx context.Context) (*GetLiveWebAreaListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-interface/v1/index/getWebAreaList"
	)
	return bilibili.Execute[*GetLiveWebAreaListResult](ctx, c.Client, method, url, nil,
		bilibili.FillParam("source_id", "2"))
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
	return bilibili.Execute[*CheckLiveAnchorLotteryResult](ctx, c.Client, method, url, param,
		liveReferer(param.RoomID))
}

// liveReferer 给**直播域**的接口设 Referer / Origin。
//
// ⚠️ 直播接口全都走 `api.live.bilibili.com`，而真实的调用发生在**直播间页面**上
// （`https://live.bilibili.com/<房间号>`）—— 与 `New()` 给的那对主站头
// （`https://www.bilibili.com`）**对不上**。从主站首页永远发不出这些请求，
// 这个不一致本身就是个能被认出来的特征：风控会回业务码 **-352（需要进行验证码校验）**。
//
// 这个坑在库里出现过三次，前两次都只补了一个接口：
//   - `3c63c65` GetLiveAreaRoomList（WBI 签名会清空 Referer，那之后要补回来）
//   - Anchor/Check
//
// 而 **Anchor/Join 一直漏着** —— 于是"查得到、参与不了"：2026-10-12 线上每次 Join
// 都被 -352 挡，还把这个账号送进了风控名单（连在浏览器里手动参加都要验证码）。
// 所以这次抽成一个函数，别再逐个接口手写。
//
// roomID > 0 时 Referer 带房间号（与浏览器那条请求逐字一致）；拿不到房间号的调用方
// 传 0，退成直播首页 —— 域对得上，比主站那对强得多。
func liveReferer(roomID int64) bilibili.Handler {
	return func(r *resty.Request) error {
		referer := "https://live.bilibili.com/"
		if roomID > 0 {
			referer = fmt.Sprintf("https://live.bilibili.com/%d", roomID)
		}
		r.SetHeader("Referer", referer)
		r.SetHeader("Origin", "https://live.bilibili.com")
		return nil
	}
}

// JoinLiveAnchorLotteryParam 指定参与天选时刻的抽奖信息，字段来自 CheckLiveAnchorLottery 的结果。
type JoinLiveAnchorLotteryParam struct {
	ID      int64  // 天选抽奖 id
	GiftID  int64  // 礼物 id
	GiftNum int    // 礼物数量
	VisitID string // 访问标识；留空时由库自动生成

	// RoomID **只用来拼 Referer**（见 liveReferer），不进请求表单。
	//
	// 接口本身不需要它，但浏览器那条请求的 Referer 是 `live.bilibili.com/<房间号>`，
	// 而调用方从 Check 那边本来就知道房间号。留 0 则退成直播首页。
	RoomID int64
}

// joinLiveAnchorLotteryForm 是参与天选时刻的表单字段。
type joinLiveAnchorLotteryForm struct {
	ID       int64  `json:"id"`
	GiftID   int64  `json:"gift_id"`
	GiftNum  int    `json:"gift_num"`
	VisitID  string `json:"visit_id"`
	Platform string `json:"platform"` // 固定为 pc
}

// JoinLiveAnchorLottery 参与天选时刻抽奖，CSRF 自动填入表单。
// 注意：要求赠礼的天选需赠送消耗电池的礼物（GiftPrice 为礼物单价），参与前请自行检查 CheckLiveAnchorLotteryResult.GiftPrice。
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
	return bilibili.Execute[*JoinLiveAnchorLotteryResult](ctx, c.Client, method, url, joinLiveAnchorLotteryForm{
		ID:       param.ID,
		GiftID:   param.GiftID,
		GiftNum:  param.GiftNum,
		VisitID:  param.VisitID,
		Platform: "pc",
	}, bilibili.MoveFormParams("id", "gift_id", "gift_num", "visit_id", "platform"),
		bilibili.FillFormCsrf(c.Client),
		liveReferer(param.RoomID))
}

// GetLiveNotice 获取直播公告。B 站会通过本接口的 Set-Cookie 下发 LIVE_BUVID，
// 响应 Cookie 已由请求链路自动合并进客户端，因此调用本方法即可完成补全。
func (c *Client) GetLiveNotice(ctx context.Context) error {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/news/v1/notice/recom"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, nil,
		bilibili.FillParam("product", "live"))
	return err
}

// randomVisitID 生成 12 位随机访问标识：首位 1-9、中间 10 位随机小写字母数字、末尾固定 0。
func randomVisitID() (string, error) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 11)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	visitID := make([]byte, 12)
	visitID[0] = '1' + b[0]%9
	for i := range 10 {
		visitID[i+1] = charset[int(b[i+1])%len(charset)]
	}
	visitID[11] = '0'
	return string(visitID), nil
}

// GetLiveWebRoomListParam 指定获取直播间列表的条件（Web 端新接口）。
type GetLiveWebRoomListParam struct {
	Platform     string `json:"platform" request:"default=web"`   // 平台。留空时自动填 web
	ParentAreaID int64  `json:"parent_area_id"`                   // 父分区 id
	AreaID       int64  `json:"area_id"`                          // 子分区 id。0 表示全部
	SortType     string `json:"sort_type"`                        // 排序方式；可为空
	Page         int    `json:"page" request:"default=1"`         // 页码，从 1 开始
	WebLocation  string `json:"web_location" request:"omitempty"` // 页面位置标识；留空时由库填入 444.7
}

// GetLiveWebRoomList 获取直播间列表（Web 端新接口，WBI 签名）。
// 注意：旧的 second/getList 接口已被风控拦截（-352），当前网页端使用本接口。
func (c *Client) GetLiveWebRoomList(ctx context.Context, param GetLiveWebRoomListParam) (*GetLiveWebRoomListResult, error) {
	if param.WebLocation == "" {
		param.WebLocation = liveWebLocationDefault
	}
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-interface/v1/index/getList"
	)
	return bilibili.Execute[*GetLiveWebRoomListResult](ctx, c.Client, method, url, param,
		c.FillWbi())
}

// liveWebLocationDefault 是直播间列表接口默认的页面位置标识。
const liveWebLocationDefault = "444.7"

// GetLiveHotRankList 获取首页人气榜直播间列表（Web 端）。
// 响应中的 LotStatus 标记直播间是否有进行中的抽奖类活动，可用于初筛。
func (c *Client) GetLiveHotRankList(ctx context.Context) (*GetLiveHotRankListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.live.bilibili.com/xlive/web-interface/v1/index/getHotRankList"
	)
	return bilibili.Execute[*GetLiveHotRankListResult](ctx, c.Client, method, url, nil,
		bilibili.FillParam("web_location", liveWebLocationDefault))
}

// 发送直播弹幕的出厂默认值（与网页端一致）。
const (
	liveDanmakuDefaultColor    = 16777215 // 白色
	liveDanmakuDefaultFontSize = 25       // 标准字号
	liveDanmakuDefaultMode     = 1        // 滚动弹幕
	liveDanmakuDefaultRnd      = 1        // 固定值，源实现如此
)

// SendLiveDanmakuParam 指定要发送的直播弹幕，CSRF 自动填入表单。
//
// 服务端对弹幕有频率与长度限制（普通用户单条 20 字符、发送间隔数秒），
// 本库不代为限流，调用方需自行控制节奏。
type SendLiveDanmakuParam struct {
	RoomID   int64  // 直播间号
	Msg      string // 弹幕内容
	Color    int    // 弹幕颜色，十进制 RGB 值；0 表示默认白色
	FontSize int    // 弹幕字号；0 表示默认 25
	Mode     int    // 弹幕模式；0 表示默认滚动弹幕
	Bubble   int    // 气泡 id；0 表示不使用气泡
}

// sendLiveDanmakuForm 是发送直播弹幕的表单字段。
type sendLiveDanmakuForm struct {
	Msg      string `json:"msg"`
	RoomID   int64  `json:"roomid"`
	Bubble   int    `json:"bubble"`
	Color    int    `json:"color"`
	FontSize int    `json:"fontsize"`
	Mode     int    `json:"mode"`
	Rnd      int    `json:"rnd"`
}

// SendLiveDanmaku 发送一条直播弹幕（文本，非表情）。需要登录态：CSRF 取自 Cookie
// 的 bili_jct，由库自动填入。
func (c *Client) SendLiveDanmaku(ctx context.Context, param SendLiveDanmakuParam) (*SendLiveDanmakuResult, error) {
	color := param.Color
	if color == 0 {
		color = liveDanmakuDefaultColor
	}
	fontSize := param.FontSize
	if fontSize == 0 {
		fontSize = liveDanmakuDefaultFontSize
	}
	mode := param.Mode
	if mode == 0 {
		mode = liveDanmakuDefaultMode
	}

	const (
		method = resty.MethodPost
		url    = "https://api.live.bilibili.com/msg/send"
	)
	return bilibili.Execute[*SendLiveDanmakuResult](ctx, c.Client, method, url, sendLiveDanmakuForm{
		Msg:      param.Msg,
		RoomID:   param.RoomID,
		Bubble:   param.Bubble,
		Color:    color,
		FontSize: fontSize,
		Mode:     mode,
		Rnd:      liveDanmakuDefaultRnd,
	}, bilibili.MoveFormParams("msg", "roomid", "bubble", "color", "fontsize", "mode", "rnd"),
		bilibili.FillFormCsrf(c.Client),
		liveReferer(param.RoomID))
}
