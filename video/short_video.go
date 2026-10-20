package video

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频推荐相关接口。响应模型见 recommend_model.go。

// GetStoryVideoListParam 是获取短视频播放列表的参数。
type GetStoryVideoListParam struct {
	Aid          int64  `json:"aid,omitempty" request:"query,omitempty"`           // 点击视频的aid。会影响后续视频内容
	DisplayID    int    `json:"display_id,omitempty" request:"query,omitempty"`    // 视频列表页数。从1开始，第1页会得到比其他页多aid处所填视频
	AccessKey    string `json:"access_key,omitempty" request:"query,omitempty"`    // APP登录Token。APP方式必要
	AdExtra      string `json:"ad_extra,omitempty" request:"query,omitempty"`      // 额外广告？
	Appkey       string `json:"appkey,omitempty" request:"query,omitempty"`        // APP密钥。APP方式必要
	AutoPlay     int    `json:"auto_play,omitempty" request:"query,omitempty"`     // 自动播放。可为0
	Build        int    `json:"build,omitempty" request:"query,omitempty"`         // 版本。可为8130300
	Bvid         string `json:"bvid,omitempty" request:"query,omitempty"`          // 视频的bv号。可为空，但没有会导致返回不正常
	CLocale      string `json:"c_locale,omitempty" request:"query,omitempty"`      // 语言。如zh_CN
	Channel      int    `json:"channel,omitempty" request:"query,omitempty"`       // 频道
	Cid          int64  `json:"cid,omitempty" request:"query,omitempty"`           // 目标频道id
	Contain      bool   `json:"contain,omitempty" request:"query,omitempty"`       // 未知。可为false
	CreativeID   int64  `json:"creative_id,omitempty" request:"query,omitempty"`   // 未知。默认为0
	DeviceName   string `json:"device_name,omitempty" request:"query,omitempty"`   // 设备名称。随意字符串都行
	DisableRcmd  int    `json:"disable_rcmd,omitempty" request:"query,omitempty"`  // 未知。默认为1
	Epid         int64  `json:"epid,omitempty" request:"query,omitempty"`          // 未知。默认为0
	FeedStatus   int    `json:"feed_status,omitempty" request:"query,omitempty"`   // 未知。默认为0
	Fnval        int    `json:"fnval,omitempty" request:"query,omitempty"`         // 视频流类型
	Fnver        int    `json:"fnver,omitempty" request:"query,omitempty"`         // 请求时提供的fnver。可为0
	ForceHost    int    `json:"force_host,omitempty" request:"query,omitempty"`    // 源url类型。0：无限制。1：使用http。2：使用https
	Fourk        int    `json:"fourk,omitempty" request:"query,omitempty"`         // 是否允许4K视频。0：最高1080P（默认）。1：最高4K
	From         int    `json:"from,omitempty" request:"query,omitempty"`          // 未知。可为7
	FromSpmid    string `json:"from_spmid,omitempty" request:"query,omitempty"`    // 未知。如tm.recommend.0.0
	Goto         string `json:"goto,omitempty" request:"query,omitempty"`          // 未知
	MobiApp      string `json:"mobi_app,omitempty" request:"query,omitempty"`      // 平台标识。如web、android
	Network      string `json:"network,omitempty" request:"query,omitempty"`       // 网络。如wifi
	OgvStyle     int    `json:"ogv_style,omitempty" request:"query,omitempty"`     // 未知。默认为0
	Platform     string `json:"platform,omitempty" request:"query,omitempty"`      // 平台。如web或android
	PlayerNet    int    `json:"player_net,omitempty" request:"query,omitempty"`    // 未知。默认为1
	Pull         int    `json:"pull,omitempty" request:"query,omitempty"`          // 未知。默认为1
	Pn           int    `json:"pn,omitempty" request:"query,omitempty"`            // 似乎不是页码。默认为32
	RequestFrom  int    `json:"request_from,omitempty" request:"query,omitempty"`  // 未知。默认为0
	SLocale      string `json:"s_locale,omitempty" request:"query,omitempty"`      // 语言。如zh_CN
	Spmid        string `json:"spmid,omitempty" request:"query,omitempty"`         // 未知。如main.ugc-video-detail-vertical.0.0
	Statistics   string `json:"statistics,omitempty" request:"query,omitempty"`    // 位置。如{"appId":1,"platform":3,"version":"8.13.0","abtest":""}
	StoryParam   string `json:"story_param,omitempty" request:"query,omitempty"`   // 未知
	Trackid      string `json:"trackid,omitempty" request:"query,omitempty"`       // 路径id？未知
	Ts           int64  `json:"ts,omitempty" request:"query,omitempty"`            // 秒级时间戳
	VideoMode    int    `json:"video_mode,omitempty" request:"query,omitempty"`    // 视频模式？未知。可为2
	VoiceBalance int    `json:"voice_balance,omitempty" request:"query,omitempty"` // 未知。默认为1
	Sign         string `json:"sign,omitempty" request:"query,omitempty"`          // APP签名。APP方式必要
}

// GetStoryVideoList 根据点击视频获取的短视频播放列表
//
// 在APP端点击主页视频后发出的请求。该接口走 APP 域名，sign 由调用方按 APP 签名规则生成，
// 库不计算签名；接口文档标注为未验证。
func (c *Client) GetStoryVideoList(ctx context.Context, param GetStoryVideoListParam) (*GetStoryVideoListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/feed/index/story"
	)
	return bilibili.Execute[*GetStoryVideoListResult](ctx, c.Client, method, url, param)
}

// GetShortVideoListParam 是获取短视频模式视频列表的参数。
type GetShortVideoListParam struct {
	Fnval        int    `json:"fnval,omitempty" request:"query,omitempty"`         // 视频流格式标识。默认为272
	Fnver        int    `json:"fnver,omitempty" request:"query,omitempty"`         // 视频流版本标识。恒为1
	ForceHost    int    `json:"force_host,omitempty" request:"query,omitempty"`    // 源url类型。0：无限制。1：使用http。2：使用https
	Fourk        int    `json:"fourk,omitempty" request:"query,omitempty"`         // 是否允许4K视频。0：最高1080P（默认）。1：最高4K
	Guidance     int    `json:"guidance,omitempty" request:"query,omitempty"`      // 0
	HTTPSURLReq  int    `json:"https_url_req,omitempty" request:"query,omitempty"` // 0
	InlineDanmu  int    `json:"inline_danmu,omitempty" request:"query,omitempty"`  // 2
	InlineSound  int    `json:"inline_sound,omitempty" request:"query,omitempty"`  // 1
	InterestID   int    `json:"interest_id,omitempty" request:"query,omitempty"`   // 0
	LoginEvent   int    `json:"login_event,omitempty" request:"query,omitempty"`   // 登录状态。0：登录。1：未登录
	MobiApp      string `json:"mobi_app,omitempty" request:"query,omitempty"`      // 设备类型。如android
	Network      string `json:"network,omitempty" request:"query,omitempty"`       // 网络类型。如wifi
	OpenEvent    int    `json:"open_event,omitempty" request:"query,omitempty"`    // 未知
	Platform     string `json:"platform,omitempty" request:"query,omitempty"`      // 设备类型。如android
	Pull         bool   `json:"pull,omitempty" request:"query,omitempty"`          // 未知
	Qn           int    `json:"qn,omitempty" request:"query,omitempty"`            // 似乎是画质。如32
	RecsysMode   int    `json:"recsys_mode,omitempty" request:"query,omitempty"`   // 0
	SLocale      string `json:"s_locale,omitempty" request:"query,omitempty"`      // 语言。如zh_CN
	VideoMode    int    `json:"video_mode,omitempty" request:"query,omitempty"`    // 未知。如1
	AccessKey    string `json:"accessKey,omitempty" request:"query,omitempty"`     // 登录成功后返回的accessToken。添加此参数会返回个性化内容和横幅
	VoiceBalance int    `json:"voice_balance,omitempty" request:"query,omitempty"` // 音量均衡？默认为1
}

// GetShortVideoList 获取短视频模式视频列表
//
// 该接口走 APP 域名，参数按 APP 端填写；接口文档标注为未验证。
func (c *Client) GetShortVideoList(ctx context.Context, param GetShortVideoListParam) (*GetShortVideoListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/feed/index"
	)
	return bilibili.Execute[*GetShortVideoListResult](ctx, c.Client, method, url, param)
}
