package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户粉丝明细与关注分组相关接口。响应模型见 relation_model.go 与 relation_tag_model.go。

// GetUserFansParam 是查询用户粉丝明细（新）的请求参数。
type GetUserFansParam struct {
	Vmid         int    `json:"vmid"`                                               // 目标用户 mid
	Ps           int    `json:"ps,omitempty" request:"query,omitempty"`             // 每页项数。默认为 50
	Pn           int    `json:"pn,omitempty" request:"query,omitempty"`             // 页码。默认为 1
	Offset       string `json:"offset,omitempty" request:"query,omitempty"`         // 偏移量。从上次响应的 data.offset 获取；有效且不为 rcmd 时忽略 pn
	LastAccessTs int64  `json:"last_access_ts,omitempty" request:"query,omitempty"` // 上次访问粉丝列表的时间。秒级时间戳，从粉丝未读状态接口获取
	From         string `json:"from,omitempty" request:"query,omitempty"`           // 请求来源。为 main 且目标用户为自己时按智能推荐算法排序
}

// GetUserFans 查询用户粉丝明细（新）。
//
// 仅当登录、referer 为 bilibili.com 子域名且 UA 不含 python 时返回列表。
func (c *Client) GetUserFans(ctx context.Context, param GetUserFansParam) (*GetUserFansResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/fans"
	)
	return bilibili.Execute[*GetUserFansResult](ctx, c.Client, method, url, param)
}

// GetUserFollowers3Param 是查询用户粉丝明细（biligame 游戏中心 H5）的请求参数。
type GetUserFollowers3Param struct {
	Vmid int `json:"vmid"`                                   // 目标用户 mid
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 20
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
}

// GetUserFollowers3 查询用户粉丝明细（biligame 游戏中心 H5）。
//
// 隐私设置拒绝、未携带 Cookie、目标用户不存在三种情况均返回 code 0 与空列表，调用方无法区分。
func (c *Client) GetUserFollowers3(ctx context.Context, param GetUserFollowers3Param) (*GetUserFollowers3Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://line3-h5-mobile-api.biligame.com/game/center/h5/user/relationship/follower_list"
	)
	return bilibili.Execute[*GetUserFollowers3Result](ctx, c.Client, method, url, param)
}

// GetUserFollowersUnreadCount 获取自己粉丝列表的未读状态。
//
// 调用此接口后会重置未读状态。
func (c *Client) GetUserFollowersUnreadCount(ctx context.Context) (*GetUserFollowersUnreadCountResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followers/unread/count"
	)
	return bilibili.Execute[*GetUserFollowersUnreadCountResult](ctx, c.Client, method, url, nil)
}

// GetRelationTagUsersParam 是查询关注分组明细的请求参数。
type GetRelationTagUsersParam struct {
	Tagid     int    `json:"tagid"`                                          // 分组 id。0：默认分组。-10：特别关注。-20：所有
	OrderType string `json:"order_type,omitempty" request:"query,omitempty"` // 排序方式。按照关注顺序排列：留空。按照最常访问排列：attention
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为 20
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为 1
}

// GetRelationTagUsers 查询关注分组明细。只可查询属于自己的分组。
func (c *Client) GetRelationTagUsers(ctx context.Context, param GetRelationTagUsersParam) ([]GetRelationTagUsersItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/tag"
	)
	return bilibili.Execute[[]GetRelationTagUsersItem](ctx, c.Client, method, url, param)
}

// GetRelationTagUserParam 是查询目标用户所在分组的请求参数。
type GetRelationTagUserParam struct {
	Fid int `json:"fid"` // 目标用户 mid
}

// GetRelationTagUser 查询目标用户所在的分组。键为分组 id，值为分组名称。
func (c *Client) GetRelationTagUser(ctx context.Context, param GetRelationTagUserParam) (GetRelationTagUserResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/tag/user"
	)
	return bilibili.Execute[GetRelationTagUserResult](ctx, c.Client, method, url, param)
}

// GetRelationTagSpecial 查询所有特别关注 mid。
func (c *Client) GetRelationTagSpecial(ctx context.Context) (GetRelationTagSpecialResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/tag/special"
	)
	return bilibili.Execute[GetRelationTagSpecialResult](ctx, c.Client, method, url, nil)
}

// CreateRelationTagParam 是创建分组的请求参数。
type CreateRelationTagParam struct {
	Tag string `json:"tag"` // 分组名。最长 16 字符
}

// CreateRelationTag 创建分组。请求体为 x-www-form-urlencoded，分组名与 CSRF 均放在表单中。
func (c *Client) CreateRelationTag(ctx context.Context, param CreateRelationTagParam) (*CreateRelationTagResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/tag/create"
	)
	return bilibili.Execute[*CreateRelationTagResult](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("tag"), bilibili.FillFormCsrf(c.Client))
}

// UpdateRelationTagParam 是重命名分组的请求参数。
type UpdateRelationTagParam struct {
	Tagid int    `json:"tagid"` // 分组 id
	Name  string `json:"name"`  // 新名称。最长 16 字符
}

// UpdateRelationTag 重命名分组。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) UpdateRelationTag(ctx context.Context, param UpdateRelationTagParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/tag/update"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("tagid", "name"), bilibili.FillFormCsrf(c.Client))
	return err
}

// DeleteRelationTagParam 是删除分组的请求参数。
type DeleteRelationTagParam struct {
	Tagid int `json:"tagid"` // 分组 id
}

// DeleteRelationTag 删除分组。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) DeleteRelationTag(ctx context.Context, param DeleteRelationTagParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/tag/del"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("tagid"), bilibili.FillFormCsrf(c.Client))
	return err
}

// AddRelationTagUsersParam 是修改分组成员的请求参数。
type AddRelationTagUsersParam struct {
	Fids   []int `json:"fids"`   // 目标用户 mid 列表。每个之间用 , 间隔
	Tagids []int `json:"tagids"` // 分组 id 列表。每个之间用 , 间隔
}

// AddRelationTagUsers 修改分组成员。
//
// 请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
// 如需删除分组中的成员，请将 tagids 设为 0（移动至默认分组），而不是取关。
func (c *Client) AddRelationTagUsers(ctx context.Context, param AddRelationTagUsersParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/tags/addUsers"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("fids", "tagids"), bilibili.FillFormCsrf(c.Client))
	return err
}

// CopyRelationTagUsersParam 是复制关注到分组的请求参数。
type CopyRelationTagUsersParam struct {
	Fids   []int `json:"fids"`   // 待复制的用户 mid 列表。每个之间用 , 间隔
	Tagids []int `json:"tagids"` // 目标分组 id 列表。每个之间用 , 间隔
}

// CopyRelationTagUsers 复制关注到分组。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) CopyRelationTagUsers(ctx context.Context, param CopyRelationTagUsersParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/tags/copyUsers"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("fids", "tagids"), bilibili.FillFormCsrf(c.Client))
	return err
}

// MoveRelationTagUsersParam 是移动关注到分组的请求参数。
type MoveRelationTagUsersParam struct {
	BeforeTagids []int `json:"beforeTagids"` // 原分组 id 列表。每个之间用 , 间隔
	AfterTagids  []int `json:"afterTagids"`  // 新分组 id 列表。每个之间用 , 间隔
	Fids         []int `json:"fids"`         // 待移动的用户 mid 列表。每个之间用 , 间隔
}

// MoveRelationTagUsers 移动关注到分组。请求体为 x-www-form-urlencoded，参数与 CSRF 均放在表单中。
func (c *Client) MoveRelationTagUsers(ctx context.Context, param MoveRelationTagUsersParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/tags/moveUsers"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param,
		bilibili.MoveFormParams("beforeTagids", "afterTagids", "fids"), bilibili.FillFormCsrf(c.Client))
	return err
}
