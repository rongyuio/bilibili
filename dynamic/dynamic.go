package dynamic

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 动态相关接口。各接口的返回体独立定义，互不复用。

// errMissingDedeUserID 表示 cookie 里没有 DedeUserID：网页端登录态不完整。
//
// 做成包级静态错误，调用方可以拿 errors.Is 判断「是没登录，还是别的失败」。
// 错误串以小写 b 开头，避免 staticcheck ST1005 把开头的拉丁大写字母判为违规。
var errMissingDedeUserID = errors.New("b站登录过期：缺少 DedeUserID")

// SearchDynamicAtParam 指定 at 用户时的搜索条件。
type SearchDynamicAtParam struct {
	UID     int    `json:"uid"`     // 自己的uid
	Keyword string `json:"keyword"` // 搜索关键字
}

// SearchDynamicAt 根据关键字搜索用户(at别人时的填充列表)
func (c *Client) SearchDynamicAt(ctx context.Context, param SearchDynamicAtParam) (*SearchDynamicAtResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_mix/v1/dynamic_mix/at_search"
	)
	return bilibili.Execute[*SearchDynamicAtResult](ctx, c.Client, method, url, param)
}

// GetDynamicRepostDetailParam 指定要查询转发列表的动态。
type GetDynamicRepostDetailParam struct {
	DynamicID int `json:"dynamic_id"`                                 // 动态id
	Offset    int `json:"offset,omitempty" request:"query,omitempty"` // 偏移量
}

// GetDynamicRepostDetail 获取动态转发列表
func (c *Client) GetDynamicRepostDetail(ctx context.Context, param GetDynamicRepostDetailParam) (*GetDynamicRepostDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/repost_detail"
	)
	return bilibili.Execute[*GetDynamicRepostDetailResult](ctx, c.Client, method, url, param)
}

// GetDynamicLikeListParam 指定要查询点赞列表的动态。
type GetDynamicLikeListParam struct {
	DynamicID int64 `json:"dynamic_id"`                             // 动态id
	Pn        int64 `json:"pn,omitempty" request:"query,omitempty"` // 页码
	Ps        int64 `json:"ps,omitempty" request:"query,omitempty"` // 每页数量。该值不得大于20
}

// GetDynamicLikeList 获取动态点赞列表
func (c *Client) GetDynamicLikeList(ctx context.Context, param GetDynamicLikeListParam) (*GetDynamicLikeListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_like/v1/dynamic_like/spec_item_likes"
	)
	return bilibili.Execute[*GetDynamicLikeListResult](ctx, c.Client, method, url, param)
}

// GetDynamicLiveUserListParam 指定直播中关注者的每页数量。
type GetDynamicLiveUserListParam struct {
	Size int `json:"size,omitempty" request:"query,omitempty"` // 每页显示数。默认为10
}

// GetDynamicLiveUserList 获取正在直播的已关注者
func (c *Client) GetDynamicLiveUserList(ctx context.Context, param GetDynamicLiveUserListParam) (*GetDynamicLiveUserListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/w_live_users"
	)
	return bilibili.Execute[*GetDynamicLiveUserListResult](ctx, c.Client, method, url, param)
}

// GetDynamicUpListParam 指定更新 UP 主列表的青少年模式。
type GetDynamicUpListParam struct {
	TeenagersMode int `json:"teenagers_mode,omitempty" request:"query,omitempty"` // 是否开启青少年模式。否：0。是：1
}

// GetDynamicUpList 获取发布新动态的已关注者
func (c *Client) GetDynamicUpList(ctx context.Context, param GetDynamicUpListParam) (*GetDynamicUpListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/w_dyn_uplist"
	)
	return bilibili.Execute[*GetDynamicUpListResult](ctx, c.Client, method, url, param)
}

// RemoveDynamicParam 指定要删除的动态。
type RemoveDynamicParam struct {
	DynamicID int `json:"dynamic_id"` // 动态id
}

// RemoveDynamic 删除动态
func (c *Client) RemoveDynamic(ctx context.Context, param RemoveDynamicParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/rm_dynamic"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// GetDynamicDetailParam 是 RemoveDynamicParam 的别名。
type GetDynamicDetailParam = RemoveDynamicParam

// GetDynamicDetail 获取特定动态卡片信息
func (c *Client) GetDynamicDetail(ctx context.Context, param GetDynamicDetailParam) (*GetDynamicDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail"
	)
	return bilibili.Execute[*GetDynamicDetailResult](ctx, c.Client, method, url, param)
}

// GetDynamicPortal 获取最近更新UP主列表（其实就是获取自己的动态门户）
func (c *Client) GetDynamicPortal(ctx context.Context) (*GetDynamicPortalResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/portal"
	)
	return bilibili.Execute[*GetDynamicPortalResult](ctx, c.Client, method, url, nil)
}

// UploadDynamicBfsParam 指定图片动态上传参数。文件本身通过 UploadDynamicBfs 的 file 参数传入。
type UploadDynamicBfsParam struct {
	Category string `json:"category" request:"query"` // 图片分类
}

// UploadDynamicBfs 为图片动态上传图片。
//
// 文件字段交由 Resty 生成 multipart boundary；CSRF 从请求 Cookie 快照填入 query。
func (c *Client) UploadDynamicBfs(ctx context.Context, param UploadDynamicBfsParam, fileName string, file io.Reader) (url string, size bilibili.Size, err error) {
	const (
		method   = resty.MethodPost
		endpoint = "https://api.bilibili.com/x/dynamic/feed/draw/upload_bfs"
	)
	data, err := bilibili.Execute[*UploadDynamicBfsResult](ctx, c.Client, method, endpoint, param,
		bilibili.FillCsrf(c.Client), func(r *resty.Request) error {
			r.SetFileReader("file_up", fileName, file)
			return nil
		})
	if err != nil {
		return "", bilibili.Size{}, err
	}
	return data.ImageURL, bilibili.Size{Width: data.ImageWidth, Height: data.ImageHeight}, nil
}

// CreateDynamicParam 指定要发表的纯文本动态。
type CreateDynamicParam struct {
	DynamicID       int                   `json:"dynamic_id"`                                            // 0
	Type            int                   `json:"type"`                                                  // 4
	Rid             int                   `json:"rid"`                                                   // 0
	Content         string                `json:"content"`                                               // 动态内容
	UpChooseComment int                   `json:"up_choose_comment,omitempty" request:"query,omitempty"` // 0
	UpCloseComment  int                   `json:"up_close_comment,omitempty" request:"query,omitempty"`  // 0
	Extension       string                `json:"extension,omitempty" request:"query,omitempty"`         // 位置信息
	AtUIDs          string                `json:"at_uids,omitempty" request:"query,omitempty"`           // 动态中 at 到的用户的 uid。使用逗号,分隔
	Ctrl            []bilibili.FormatCtrl `json:"ctrl,omitempty" request:"query,omitempty"`              // 特殊格式控制 (如 at 别人时的蓝字体和链接)
}

// CreateDynamic 发表纯文本动态
func (c *Client) CreateDynamic(ctx context.Context, param CreateDynamicParam) (*CreateDynamicResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/create"
	)
	return bilibili.Execute[*CreateDynamicResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// GetUserSpaceDynamicParam 指定用户空间动态的查询条件。
type GetUserSpaceDynamicParam struct {
	Offset         string `json:"offset"`          // 分页偏移量
	HostMid        string `json:"host_mid"`        // 用户UID
	TimezoneOffset int    `json:"timezone_offset"` // -480
	Features       string `json:"features"`        // itemOpusStyle
}

// GetUserSpaceDynamic 获取用户空间动态，mid就是用户UID，无需登录。
//
// 返回结构较为繁琐
func (c *Client) GetUserSpaceDynamic(ctx context.Context, param GetUserSpaceDynamicParam) (*GetUserSpaceDynamicResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space"
	)
	return bilibili.Execute[*GetUserSpaceDynamicResult](ctx, c.Client, method, url, param)
}

// dynamicFeedAllFeatures 是网页版动态流默认请求的功能开关串。
//
// 它随前端版本漂移，但少几个开关只会让卡片少几项可选内容、不影响返回结构，
// 所以留成常量而不是暴露给调用方。
const dynamicFeedAllFeatures = "itemOpusStyle,listOnlyfans,opusBigCover,onlyfansVote,decorationCard,onlyfansAssetsV2,forwardListHidden,ugcDelete,onlyfansQaCard,commentsNewVersion,avatarAutoTheme,sunflowerStyle,cardsEnhance,eva3CardOpus,eva3CardVideo,eva3CardComment,eva3CardVote,eva3CardUser"

// dynamicWebLocation 是动态相关接口默认的页面位置标识。
const dynamicWebLocation = "333.1365"

// GetDynamicFeedAllParam 指定关注动态流的分类与分页。
type GetDynamicFeedAllParam struct {
	FeedType       string `json:"type" request:"query,omitempty"`            // 分类：all(默认) / video / pgc / article
	Offset         string `json:"offset" request:"query,omitempty"`          // 分页偏移量，取上一页返回的 Offset
	TimezoneOffset int    `json:"timezone_offset" request:"query,omitempty"` // 时区偏移，网页端为 -480；传 0 时由库填入
	Features       string `json:"features" request:"query,omitempty"`        // 功能开关，留空由库填入网页端默认值
	WebLocation    string `json:"web_location" request:"query,omitempty"`    // 页面位置标识，留空由库填入默认值
}

// GetDynamicFeedAll 获取关注动态流（动态首页）。
func (c *Client) GetDynamicFeedAll(ctx context.Context, param GetDynamicFeedAllParam) (*GetDynamicFeedAllResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/all"
	)
	return bilibili.Execute[*GetDynamicFeedAllResult](ctx, c.Client, method, url, normalizeDynamicFeedAllParam(param), bilibili.FillParam("platform", "web"))
}

// normalizeDynamicFeedAllParam 补齐网页端固定带、调用方通常不关心的那几个参数。
func normalizeDynamicFeedAllParam(param GetDynamicFeedAllParam) GetDynamicFeedAllParam {
	if param.FeedType == "" {
		param.FeedType = "all"
	}
	if param.TimezoneOffset == 0 {
		param.TimezoneOffset = -480
	}
	if param.Features == "" {
		param.Features = dynamicFeedAllFeatures
	}
	if param.WebLocation == "" {
		param.WebLocation = dynamicWebLocation
	}
	return param
}

// dynamicRepostScene 是发表动态接口里「转发」的 scene 值（图文是 2）。
const dynamicRepostScene = 4

// RepostDynamicParam 指定被转发的原动态与转发文案。
type RepostDynamicParam struct {
	DynamicID int64  // 原动态 ID，必填
	Content   string // 转发时附带的文字，可为空
	UploadID  string // 客户端上传标识；留空由库按「mid_秒级时间戳_四位随机」生成
}

// repostDynContentNode 是动态正文里的一个富文本节点。
// type：1=文本、2=@用户、3=抽奖、4=投票；后三者的 biz_id 放对应的 id。
type repostDynContentNode struct {
	RawText string `json:"raw_text"`
	Type    int    `json:"type"`
	BizID   string `json:"biz_id"`
}

// repostDynamicBody 是发表动态接口的请求体。
//
// ⚠️ 被转发的原动态 ID **不在** dyn_req.repost_src 里：网页版构造完 dyn_req 之后会把
// repost_src 删掉，改放到外层的 web_repost_src.dyn_id_str（scene=5 的视频转发同理）。
// 照着文档里 scene=2 那份示例去拼 repost_src 是发不出去的。
type repostDynamicBody struct {
	DynReq       repostDynReq `json:"dyn_req"`
	WebRepostSrc repostWebSrc `json:"web_repost_src"`
}

// repostDynReq 是动态本体。
type repostDynReq struct {
	Content  repostDynContent `json:"content"`
	Scene    int              `json:"scene"`
	UploadID string           `json:"upload_id"`
	Meta     repostDynMeta    `json:"meta"`
	Option   repostDynOption  `json:"option"`
}

// repostDynContent 是动态正文。
type repostDynContent struct {
	Contents []repostDynContentNode `json:"contents"`
}

// repostDynMeta 是动态的来源信息。
type repostDynMeta struct {
	AppMeta repostDynAppMeta `json:"app_meta"`
}

// repostDynAppMeta 是动态来源的平台标识。
type repostDynAppMeta struct {
	From    string `json:"from"`
	MobiApp string `json:"mobi_app"`
}

// repostDynOption 是动态的互动设置。
type repostDynOption struct {
	Aigc int `json:"aigc"` // 1=AI 生成、2=非 AI
}

// repostWebSrc 指出被转发的原动态。
type repostWebSrc struct {
	DynIDStr string `json:"dyn_id_str"`
}

// RepostDynamic 转发一条动态（网页版「发表动态」接口，scene=4）。
//
// 需要登录态：mid 取自 Cookie DedeUserID，CSRF 取自 Cookie bili_jct，任一缺失都返回错误。
// 该接口不需要 WBI 签名，CSRF 走 URL 参数。
func (c *Client) RepostDynamic(ctx context.Context, param RepostDynamicParam) (*RepostDynamicResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/dynamic/feed/create/dyn"
	)
	if param.DynamicID <= 0 {
		// 根包未导出参数错误构造器，这里直接填充其导出字段；
		// 具体原因文本由根包的 reason 字段承载，子包无法设置。
		return nil, &bilibili.ParamError{
			RootType:  "RepostDynamicParam",
			GoField:   "DynamicID",
			Parameter: "dyn_id_str",
			Location:  "json",
		}
	}
	return bilibili.Execute[*RepostDynamicResult](ctx, c.Client, method, url, nil,
		repostDynamicHandler(param), bilibili.FillCsrf(c.Client))
}

// repostDynamicHandler 构造转发请求体，并补上网页端固定带的 platform 参数。
//
// 必须在 fillCsrf 之前注册：这里设置 Content-Type 与 body，fillCsrf 只往 query 里补 CSRF。
func repostDynamicHandler(param RepostDynamicParam) bilibili.Handler {
	return func(r *resty.Request) error {
		mid := dynamicCookieValue(r.Cookies, "DedeUserID")
		if mid == "" {
			return errMissingDedeUserID
		}

		uploadID := param.UploadID
		if uploadID == "" {
			uploadID = newDynamicUploadID(mid)
		}

		// 没有文案时给空数组：网页版不带文字转发时同样没有文本节点。
		contents := make([]repostDynContentNode, 0, 1)
		if param.Content != "" {
			contents = append(contents, repostDynContentNode{RawText: param.Content, Type: 1, BizID: ""})
		}

		body, err := json.Marshal(repostDynamicBody{
			DynReq: repostDynReq{
				Content:  repostDynContent{Contents: contents},
				Scene:    dynamicRepostScene,
				UploadID: uploadID,
				Meta:     repostDynMeta{AppMeta: repostDynAppMeta{From: "create.dynamic.web", MobiApp: "web"}},
				Option:   repostDynOption{Aigc: 2},
			},
			WebRepostSrc: repostWebSrc{DynIDStr: strconv.FormatInt(param.DynamicID, 10)},
		})
		if err != nil {
			return fmt.Errorf("编码转发动态请求体: %w", err)
		}

		r.SetQueryParam("platform", "web")
		r.SetHeader("Content-Type", "application/json")
		r.SetBody(body)
		return nil
	}
}

// dynamicCookieValue 从请求 Cookie 快照中取指定名称的值。
func dynamicCookieValue(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie != nil && cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

// newDynamicUploadID 生成网页版风格的 upload_id：「发送人 mid_秒级时间戳_四位随机整数」。
// 取不到随机数时退化为 0 —— 这个字段本身不是必要条件。
func newDynamicUploadID(mid string) string {
	var b [2]byte
	_, _ = rand.Read(b[:])
	n := int(b[0])<<8 | int(b[1])
	return fmt.Sprintf("%s_%d_%d", mid, time.Now().Unix(), n%10000)
}
