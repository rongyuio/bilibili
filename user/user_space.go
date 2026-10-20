package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户空间相关接口：主页（置顶、代表作、个人TAG、公告、空间设置、头图）、
// 最近访问内容、投稿、课程与订阅。响应模型见 user_space_model.go。

type GetUserSpaceTopVideoParam struct {
	Vmid int `json:"vmid" request:"query"` // 目标用户mid
}

// GetUserSpaceTopVideo 查询用户置顶视频
func (c *Client) GetUserSpaceTopVideo(ctx context.Context, param GetUserSpaceTopVideoParam) (*GetUserSpaceTopVideoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/top/arc"
	)
	return bilibili.Execute[*GetUserSpaceTopVideoResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceMasterpieceListParam struct {
	Vmid int `json:"vmid" request:"query"` // 目标用户mid
}

// GetUserSpaceMasterpieceList 查询用户代表作视频列表。最多3个，无则为空
func (c *Client) GetUserSpaceMasterpieceList(ctx context.Context, param GetUserSpaceMasterpieceListParam) ([]GetUserSpaceMasterpieceListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/masterpiece"
	)
	return bilibili.Execute[[]GetUserSpaceMasterpieceListItem](ctx, c.Client, method, url, param)
}

type GetUserSpaceTagsParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid
}

// GetUserSpaceTags 查看用户个人TAG。上限5条，内容由用户自定义
func (c *Client) GetUserSpaceTags(ctx context.Context, param GetUserSpaceTagsParam) ([]GetUserSpaceTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/acc/tags"
	)
	return bilibili.Execute[[]GetUserSpaceTagsItem](ctx, c.Client, method, url, param)
}

type SetUserSpaceTagsParam struct {
	Tags []string `json:"tags,omitempty" request:"query,omitempty"` // 要设置的TAG内容。最多5个，各TAG之间用逗号分隔，单个TAG小于10字符。留空或省略为删除
}

// SetUserSpaceTags 修改个人TAG。请求体为 x-www-form-urlencoded，TAG 与 CSRF 均放在表单中
func (c *Client) SetUserSpaceTags(ctx context.Context, param SetUserSpaceTagsParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/space/acc/tags/set"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("tags"), bilibili.FillFormCsrf(c.Client))
	return err
}

type GetUserSpaceNoticeParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid
}

// GetUserSpaceNotice 查看用户空间公告。无公告时返回空串
func (c *Client) GetUserSpaceNotice(ctx context.Context, param GetUserSpaceNoticeParam) (string, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/notice"
	)
	return bilibili.Execute[string](ctx, c.Client, method, url, param)
}

type SetUserSpaceNoticeParam struct {
	Notice string `json:"notice,omitempty" request:"query,omitempty"` // 要设置的公告内容。少于150字符。留空或省略为删除公告
}

// SetUserSpaceNotice 修改空间公告。请求体为 x-www-form-urlencoded，公告与 CSRF 均放在表单中
func (c *Client) SetUserSpaceNotice(ctx context.Context, param SetUserSpaceNoticeParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/space/notice/set"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("notice"), bilibili.FillFormCsrf(c.Client))
	return err
}

type GetUserSpaceSettingsParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid
}

// GetUserSpaceSettings 查询空间设置（隐私权限、板块布局、头图）
//
// 该接口根对象使用 status 而不是 code，库按 code 缺失为 0 处理，data 仍可正常解码。
func (c *Client) GetUserSpaceSettings(ctx context.Context, param GetUserSpaceSettingsParam) (*GetUserSpaceSettingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "http://space.bilibili.com/ajax/settings/getSettings" //nolint:revive // 接口地址原样保留
	)
	return bilibili.Execute[*GetUserSpaceSettingsResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceTopPhotoListParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid。非负数，即使该用户并不存在
}

// GetUserSpaceTopPhotoList 查询可用头图列表（Web端）
//
// 该接口根对象使用 status 而不是 code，库按 code 缺失为 0 处理，data 仍可正常解码。
func (c *Client) GetUserSpaceTopPhotoList(ctx context.Context, param GetUserSpaceTopPhotoListParam) ([]GetUserSpaceTopPhotoListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://space.bilibili.com/ajax/topphoto/getlist"
	)
	return bilibili.Execute[[]GetUserSpaceTopPhotoListItem](ctx, c.Client, method, url, param)
}

type GetUserSpaceLastPlayGameParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid
}

// GetUserSpaceLastPlayGame 查询用户最近玩过的游戏V2。隐私未公开时 data 为 null
func (c *Client) GetUserSpaceLastPlayGame(ctx context.Context, param GetUserSpaceLastPlayGameParam) (*GetUserSpaceLastPlayGameResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/lastplaygame/v2"
	)
	return bilibili.Execute[*GetUserSpaceLastPlayGameResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceCoinVideoListParam struct {
	Vmid int `json:"vmid" request:"query"` // 目标用户mid
}

// GetUserSpaceCoinVideoList 查询用户最近投币视频（Web）。隐私未公开时 data 为 null
func (c *Client) GetUserSpaceCoinVideoList(ctx context.Context, param GetUserSpaceCoinVideoListParam) ([]GetUserSpaceCoinVideoListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/coin/video"
	)
	return bilibili.Execute[[]GetUserSpaceCoinVideoListItem](ctx, c.Client, method, url, param)
}

type GetUserSpaceCoinVideoListAppParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token
	Appkey    string `json:"appkey" request:"query"`                         // APP密钥
	Vmid      int    `json:"vmid" request:"query"`                           // 目标用户mid
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为1
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为20
}

// GetUserSpaceCoinVideoListApp 查询用户最近投币视频（APP）
func (c *Client) GetUserSpaceCoinVideoListApp(ctx context.Context, param GetUserSpaceCoinVideoListAppParam) (*GetUserSpaceCoinVideoListAppResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/space/coinarc"
	)
	return bilibili.Execute[*GetUserSpaceCoinVideoListAppResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceLikeVideoListParam struct {
	Vmid int `json:"vmid" request:"query"` // 目标用户mid
}

// GetUserSpaceLikeVideoList 查询用户最近点赞视频（Web）。隐私未公开时 data 为 null
//
// 文档字段表把 data 写成数组，但响应示例实际为 {"list": [...]}，这里按示例建模。
func (c *Client) GetUserSpaceLikeVideoList(ctx context.Context, param GetUserSpaceLikeVideoListParam) (*GetUserSpaceLikeVideoListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/like/video"
	)
	return bilibili.Execute[*GetUserSpaceLikeVideoListResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceLikeVideoListAppParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP登录Token
	Appkey    string `json:"appkey" request:"query"`                         // APP密钥
	Vmid      int    `json:"vmid" request:"query"`                           // 目标用户mid
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为1
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为20
}

// GetUserSpaceLikeVideoListApp 查询用户最近点赞视频（APP）
func (c *Client) GetUserSpaceLikeVideoListApp(ctx context.Context, param GetUserSpaceLikeVideoListAppParam) (*GetUserSpaceLikeVideoListAppResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.bilibili.com/x/v2/space/likearc"
	)
	return bilibili.Execute[*GetUserSpaceLikeVideoListAppResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceArchiveCursorParam struct {
	Vmid          int    `json:"vmid" request:"query"`                               // 目标用户mid
	Aid           int    `json:"aid,omitempty" request:"query,omitempty"`            // 请求返回起始视频。首次请求不需要，填写上次请求返回最后视频的aid
	Order         string `json:"order,omitempty" request:"query,omitempty"`          // 排序方式。click：最多播放。pubdate：最新发布
	Sort          string `json:"sort,omitempty" request:"query,omitempty"`           // 创建时间：asc。修改时间：desc
	AccessKey     string `json:"access_key,omitempty" request:"query,omitempty"`     // APP登录Token。实测全程未使用
	Appkey        string `json:"appkey" request:"query"`                             // APP密钥。取自 APPKey 公开表
	Build         int    `json:"build,omitempty" request:"query,omitempty"`          // 版本。实测值任意
	CLocale       string `json:"c_locale,omitempty" request:"query,omitempty"`       // 语言。如 zh_CN
	Channel       int    `json:"channel,omitempty" request:"query,omitempty"`        // 频道
	Fnval         int    `json:"fnval,omitempty" request:"query,omitempty"`          // 视频流类型
	Fnver         int    `json:"fnver,omitempty" request:"query,omitempty"`          // 请求时提供的fnver。可为0
	ForceHost     int    `json:"force_host,omitempty" request:"query,omitempty"`     // 源url类型。0：无限制。1：使用http。2：使用https
	Fourk         int    `json:"fourk,omitempty" request:"query,omitempty"`          // 是否允许4K视频。0：最高1080P。1：最高4K
	IncludeCursor bool   `json:"include_cursor,omitempty" request:"query,omitempty"` // 未知。可为false
	MobiApp       string `json:"mobi_app,omitempty" request:"query,omitempty"`       // 平台标识。可为 web、android 等
	Platform      string `json:"platform,omitempty" request:"query,omitempty"`       // 平台。可为 web
	Ps            int    `json:"ps,omitempty" request:"query,omitempty"`             // 每页条数。默认为20
	Pn            int    `json:"pn,omitempty" request:"query,omitempty"`             // 似乎不是页码。默认为32
	PnPolicy      int    `json:"pn_policy,omitempty" request:"query,omitempty"`      // 未知。可为1
	SLocale       string `json:"s_locale,omitempty" request:"query,omitempty"`       // 语言。如 zh_CN
	Statistics    string `json:"statistics,omitempty" request:"query,omitempty"`     // 位置。如 {"appId":1,"platform":3,"version":"8.13.0","abtest":""}
	Ts            int64  `json:"ts" request:"query"`                                 // 秒级时间戳。必须存在且非空，值本身不被校验
	VoiceBalance  int    `json:"voice_balance,omitempty" request:"query,omitempty"`  // 未知。默认为1
	Sign          string `json:"sign" request:"query"`                               // APP签名。算法见 APP API 签名与鉴权文档
}

// GetUserSpaceArchiveCursor 查询用户投稿明细（APP、无需wbi鉴权）
//
// 不能裸打：必须按 APP 签名规则计算 sign（md5，小写十六进制），本方法不代算签名。
func (c *Client) GetUserSpaceArchiveCursor(ctx context.Context, param GetUserSpaceArchiveCursorParam) (*GetUserSpaceArchiveCursorResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.biliapi.com/x/v2/space/archive/cursor"
	)
	return bilibili.Execute[*GetUserSpaceArchiveCursorResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceArticleListParam struct {
	Mid  int    `json:"mid" request:"query"`                      // 用户uid
	Pn   int    `json:"pn,omitempty" request:"query,omitempty"`   // 页码。默认为1
	Ps   int    `json:"ps,omitempty" request:"query,omitempty"`   // 每页项数。默认为30。范围：[1,30]
	Sort string `json:"sort,omitempty" request:"query,omitempty"` // 排序方式。publish_time：最新发布。view：最多阅读。fav：最多收藏。默认：publish_time
}

// GetUserSpaceArticleList 查询用户投稿专栏明细（WBI 签名）
//
// 与 GetUserArticleList（旧接口 x/space/article）不是同一接口；data 结构在响应层面一致，
// 但本包按接口隔离返回体，单独定义 GetUserSpaceArticleListResult。
func (c *Client) GetUserSpaceArticleList(ctx context.Context, param GetUserSpaceArticleListParam) (*GetUserSpaceArticleListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/article"
	)
	return bilibili.Execute[*GetUserSpaceArticleListResult](ctx, c.Client, method, url, param, c.FillWbi())
}

type GetUserSpaceAudioListParam struct {
	Mid   int    `json:"mid" request:"query"`                       // 用户mid
	Pn    int    `json:"pn,omitempty" request:"query,omitempty"`    // 页码。默认为1
	Ps    int    `json:"ps,omitempty" request:"query,omitempty"`    // 每页项数。默认为30
	Order string `json:"order,omitempty" request:"query,omitempty"` // 排序方式。1：最新发布。2：最多播放。3：最多收藏
}

// GetUserSpaceAudioList 查询用户投稿音频明细
//
// 文档参数表写 mid，但请求示例用的是 uid，这里按参数表保留 mid。
func (c *Client) GetUserSpaceAudioList(ctx context.Context, param GetUserSpaceAudioListParam) (*GetUserSpaceAudioListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/audio/music-service/web/song/upper"
	)
	return bilibili.Execute[*GetUserSpaceAudioListResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceCourseListParam struct {
	Mid int `json:"mid" request:"query"`                    // 目标用户mid
	Pn  int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
	Ps  int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为20
}

// GetUserSpaceCourseList 查询用户发布的课程列表
func (c *Client) GetUserSpaceCourseList(ctx context.Context, param GetUserSpaceCourseListParam) (*GetUserSpaceCourseListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pugv/app/web/season/page"
	)
	return bilibili.Execute[*GetUserSpaceCourseListResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceBangumiPreviewListParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid
}

// GetUserSpaceBangumiPreviewList 查询用户追番预览列表。最多15条
//
// 该接口根对象使用 status 而不是 code，库按 code 缺失为 0 处理，data 仍可正常解码。
func (c *Client) GetUserSpaceBangumiPreviewList(ctx context.Context, param GetUserSpaceBangumiPreviewListParam) (*GetUserSpaceBangumiPreviewListResult, error) {
	const (
		method = resty.MethodGet
		url    = "http://space.bilibili.com/ajax/Bangumi/getList" //nolint:revive // 接口地址原样保留
	)
	return bilibili.Execute[*GetUserSpaceBangumiPreviewListResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceBangumiFollowListParam struct {
	Vmid int `json:"vmid" request:"query"`                   // 目标用户mid
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为15。定义域1-30
	Type int `json:"type" request:"query"`                   // 查询类型。1：追番。2：追剧
}

// GetUserSpaceBangumiFollowList 查询用户追番（追剧）明细。隐私未公开时返回业务错误 53013
func (c *Client) GetUserSpaceBangumiFollowList(ctx context.Context, param GetUserSpaceBangumiFollowListParam) (*GetUserSpaceBangumiFollowListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/bangumi/follow/list"
	)
	return bilibili.Execute[*GetUserSpaceBangumiFollowListResult](ctx, c.Client, method, url, param)
}

type GetUserSpaceSubscribedTagsParam struct {
	Mid int `json:"mid" request:"query"` // 目标用户mid
}

// GetUserSpaceSubscribedTags 查询用户关注的TAG（话题）。只显示前100个
//
// 该接口根对象使用 status 而不是 code，库按 code 缺失为 0 处理，data 仍可正常解码。
func (c *Client) GetUserSpaceSubscribedTags(ctx context.Context, param GetUserSpaceSubscribedTagsParam) (*GetUserSpaceSubscribedTagsResult, error) {
	const (
		method = resty.MethodGet
		url    = "http://space.bilibili.com/ajax/tags/getSubList" //nolint:revive // 接口地址原样保留
	)
	return bilibili.Execute[*GetUserSpaceSubscribedTagsResult](ctx, c.Client, method, url, param)
}
