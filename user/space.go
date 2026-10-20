package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户空间、名片与投稿视频相关接口。响应模型见 model_space.go。

// GetUserVideosParam 是查询用户投稿视频明细的请求参数。
type GetUserVideosParam struct {
	Mid     int    `json:"mid"`                                         // 目标用户mid
	Order   string `json:"order,omitempty" request:"query,omitempty"`   // 排序方式。默认为pubdate。最新发布：pubdate。最多播放：click。最多收藏：stow
	Tid     int    `json:"tid,omitempty" request:"query,omitempty"`     // 筛选目标分区。默认为0。0：不进行分区筛选。分区tid为所筛选的分区
	Keyword string `json:"keyword,omitempty" request:"query,omitempty"` // 关键词筛选。用于使用关键词搜索该UP主视频稿件
	Pn      int    `json:"pn,omitempty" request:"query,omitempty"`      // 页码。默认为 1
	Ps      int    `json:"ps,omitempty" request:"query,omitempty"`      // 每页项数。默认为 30
}

// GetUserVideos 查询用户投稿视频明细。
func (c *Client) GetUserVideos(ctx context.Context, param GetUserVideosParam) (*GetUserVideosResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/arc/search"
	)
	return bilibili.Execute[*GetUserVideosResult](ctx, c.Client, method, url, param, c.FillWbi())
}

// GetUserSpaceDetailParam 是获取用户空间详细信息的请求参数。
type GetUserSpaceDetailParam struct {
	Mid int `json:"mid"` // 目标用户mid
}

// GetUserSpaceDetail 获取用户空间详细信息。
func (c *Client) GetUserSpaceDetail(ctx context.Context, param GetUserSpaceDetailParam) (*GetUserSpaceDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/acc/info"
	)
	return bilibili.Execute[*GetUserSpaceDetailResult](ctx, c.Client, method, url, param, c.FillWbi())
}

// GetUserCardParam 是获取用户名片的请求参数。
type GetUserCardParam struct {
	Mid   int  `json:"mid"`                                       // 目标用户mid
	Photo bool `json:"photo,omitempty" request:"query,omitempty"` // 是否请求用户主页头图。true：是。false：否
}

// GetUserCard 获取用户用户名片 免登录。
func (c *Client) GetUserCard(ctx context.Context, param GetUserCardParam) (*GetUserCardResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/card"
	)
	return bilibili.Execute[*GetUserCardResult](ctx, c.Client, method, url, param)
}

// GetMyUserSpaceDetail 获取登录用户空间详细信息。
func (c *Client) GetMyUserSpaceDetail(ctx context.Context) (*GetMyUserSpaceDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/myinfo"
	)
	return bilibili.Execute[*GetMyUserSpaceDetailResult](ctx, c.Client, method, url, nil)
}

// CheckNickNameParam 是检查昵称是否可注册的请求参数。
type CheckNickNameParam struct {
	NickName string `json:"nickName"` // 目标昵称。最长为16字符
}

// CheckNickName 检查昵称是否可注册。
func (c *Client) CheckNickName(ctx context.Context, param CheckNickNameParam) error {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/web/generic/check/nickname"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param)
	return err
}

// JoinOldFansParam 是加入老粉计划的请求参数。
type JoinOldFansParam struct {
	Aid      string `json:"aid,omitempty" request:"query,omitempty"`      // 空串
	UpMid    string `json:"up_mid"`                                       // UP主UID
	Source   string `json:"source,omitempty" request:"query,omitempty"`   // "4"
	Scene    string `json:"scene,omitempty" request:"query,omitempty"`    // "105"
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // "web"
	MobiApp  string `json:"mobi_app,omitempty" request:"query,omitempty"` // "pc"
}

// JoinOldFans 加入老粉计划。
func (c *Client) JoinOldFans(ctx context.Context, param JoinOldFansParam) (*JoinOldFansResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v1/contract/add_contract"
	)
	return bilibili.Execute[*JoinOldFansResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// FansSendMessageParam 是老粉计划发送留言的请求参数。
type FansSendMessageParam struct {
	Aid     string `json:"aid,omitempty" request:"query,omitempty"`    // 空串
	UpMid   string `json:"up_mid"`                                     // UP主UID
	Source  string `json:"source,omitempty" request:"query,omitempty"` // "4"
	Scene   string `json:"scene,omitempty" request:"query,omitempty"`  // "105"
	Content string `json:"content"`                                    // 留言内容
}

// FansSendMessage 老粉计划发送留言。
func (c *Client) FansSendMessage(ctx context.Context, param FansSendMessageParam) (*FansSendMessageResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v1/contract/add_message"
	)
	return bilibili.Execute[*FansSendMessageResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// BatchGetUserCardsParam 是批量获取多用户详细信息的请求参数。
type BatchGetUserCardsParam struct {
	UIDs []int `json:"uids"` // 目标用户的UID列表
}

// BatchGetUserCards 获取多用户详细信息。
func (c *Client) BatchGetUserCards(ctx context.Context, param BatchGetUserCardsParam) ([]*BatchGetUserCardsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/account/v1/user/cards"
	)
	return bilibili.Execute[[]*BatchGetUserCardsResult](ctx, c.Client, method, url, param)
}
