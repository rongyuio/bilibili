package user

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 用户关注 / 粉丝 / 黑名单等关系相关接口。响应模型见 model_relation.go。

// GetUserFollowersParam 是查询用户粉丝明细的请求参数。
type GetUserFollowersParam struct {
	Vmid int `json:"vmid"`                                   // 目标用户 mid
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 50
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1。仅可查看前 1000 名粉丝
}

// GetUserFollowers 查询用户粉丝明细（需要登录）。
func (c *Client) GetUserFollowers(ctx context.Context, param GetUserFollowersParam) (*GetUserFollowersResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followers"
	)
	return bilibili.Execute[*GetUserFollowersResult](ctx, c.Client, method, url, param)
}

// GetUserFollowingsParam 是查询用户关注明细的请求参数。
type GetUserFollowingsParam struct {
	Vmid      int    `json:"vmid"`                                           // 目标用户 mid
	OrderType string `json:"order_type,omitempty" request:"query,omitempty"` // 排序方式。当目标用户为自己时有效。按照关注顺序排列：留空。按照最常访问排列：attention
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为 50
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为 1。其他用户仅可查看前 100 个
}

// GetUserFollowings 查询用户关注明细（需要登录）。
func (c *Client) GetUserFollowings(ctx context.Context, param GetUserFollowingsParam) (*GetUserFollowingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followings"
	)
	return bilibili.Execute[*GetUserFollowingsResult](ctx, c.Client, method, url, param)
}

// GetUserFollowings2Param 是查询用户关注明细2的请求参数。
type GetUserFollowings2Param struct {
	Vmid  int    `json:"vmid"`                                      // 目标用户 mid
	Order string `json:"order,omitempty" request:"query,omitempty"` // 排序方式。按照降序排列：desc。按照升序排列：asc。默认降序排列
	Ps    int    `json:"ps,omitempty" request:"query,omitempty"`    // 每页项数。默认为 50
	Pn    int    `json:"pn,omitempty" request:"query,omitempty"`    // 页码。默认为 1。仅可查看前 5 页
}

// GetUserFollowings2 查询用户关注明细2。
//
// 仅可查看前 5 页，对于已设置可见性隐私关注列表的用户，则返回的List为nil，Total为0。
func (c *Client) GetUserFollowings2(ctx context.Context, param GetUserFollowings2Param) (*GetUserFollowings2Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.biliapi.net/x/v2/relation/followings"
	)
	return bilibili.Execute[*GetUserFollowings2Result](ctx, c.Client, method, url, param)
}

// GetUserFollowings3Param 是 GetUserFollowersParam 的别名。
type GetUserFollowings3Param = GetUserFollowersParam

// GetUserFollowings3 查询用户关注明细3。
//
// 对于设置了可见性隐私关注列表的用户会返回空列表。
func (c *Client) GetUserFollowings3(ctx context.Context, param GetUserFollowings3Param) (*GetUserFollowings3Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://line3-h5-mobile-api.biligame.com/game/center/h5/user/relationship/following_list"
	)
	return bilibili.Execute[*GetUserFollowings3Result](ctx, c.Client, method, url, param)
}

// SearchUserFollowingsParam 是搜索关注明细的请求参数。
type SearchUserFollowingsParam struct {
	Vmid string `json:"vmid"`                                     // 目标用户 mid
	Name string `json:"name,omitempty" request:"query,omitempty"` // 搜索关键词
	Ps   int    `json:"ps,omitempty" request:"query,omitempty"`   // 每页项数。默认为 50
	Pn   int    `json:"pn,omitempty" request:"query,omitempty"`   // 页码。默认为 1
}

// SearchUserFollowings 搜索关注明细。
func (c *Client) SearchUserFollowings(ctx context.Context, param SearchUserFollowingsParam) (*SearchUserFollowingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followings/search"
	)
	return bilibili.Execute[*SearchUserFollowingsResult](ctx, c.Client, method, url, param)
}

// GetSameFollowingsParam 是 GetUserFollowersParam 的别名。
type GetSameFollowingsParam = GetUserFollowersParam

// GetSameFollowings 查询共同关注明细。
func (c *Client) GetSameFollowings(ctx context.Context, param GetSameFollowingsParam) (*GetSameFollowingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/same/followings"
	)
	return bilibili.Execute[*GetSameFollowingsResult](ctx, c.Client, method, url, param)
}

// GetWhispers 查询悄悄关注明细。
func (c *Client) GetWhispers(ctx context.Context) (*GetWhispersResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/whispers"
	)
	return bilibili.Execute[*GetWhispersResult](ctx, c.Client, method, url, nil)
}

// GetFriends 查询互相关注明细。
func (c *Client) GetFriends(ctx context.Context) (*GetFriendsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/friends"
	)
	return bilibili.Execute[*GetFriendsResult](ctx, c.Client, method, url, nil)
}

// GetBlacksParam 是查询黑名单明细的请求参数。
type GetBlacksParam struct {
	Ps int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 50
	Pn int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
}

// GetBlacks 查询黑名单明细。
func (c *Client) GetBlacks(ctx context.Context, param GetBlacksParam) (*GetBlacksResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/blacks"
	)
	return bilibili.Execute[*GetBlacksResult](ctx, c.Client, method, url, param)
}

// ModifyRelationAct 是操作用户关系的操作代码。
type ModifyRelationAct int

const (
	ModifyRelationActFollow     ModifyRelationAct = iota + 1 // 关注，无法对已注销或不存在的用户进行此操作
	ModifyRelationActUnfollow                                // 取关
	ModifyRelationActWhisper                                 // 悄悄关注，现已下线，使用本操作代码请求接口会提示“请求错误”
	ModifyRelationActUnwhisper                               // 取消悄悄关注
	ModifyRelationActBlack                                   // 拉黑
	ModifyRelationActUnblack                                 // 取消拉黑
	ModifyRelationActUnfollower                              // 踢出粉丝
)

// ModifyRelationParam 是操作用户关系的请求参数。
type ModifyRelationParam struct {
	Fid   int               `json:"fid"`    // 目标用户mid
	Act   ModifyRelationAct `json:"act"`    // 操作代码
	ReSrc int               `json:"re_src"` // 关注来源代码。空间：11。视频：14。文章：115。活动页面：222
}

// ModifyRelation 操作用户关系。
func (c *Client) ModifyRelation(ctx context.Context, param ModifyRelationParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/modify"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// BatchModifyRelationParam 是批量操作用户关系的请求参数。
type BatchModifyRelationParam struct {
	Fids  []int             `json:"fids"`   // 目标用户 mid 列表
	Act   ModifyRelationAct `json:"act"`    // 操作代码。仅可为 1 或 5，故只能进行批量关注和拉黑
	ReSrc int               `json:"re_src"` // 关注来源代码。同上
}

// BatchModifyRelation 批量操作用户关系。
func (c *Client) BatchModifyRelation(ctx context.Context, param BatchModifyRelationParam) (*BatchModifyRelationResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/batch/modify"
	)
	return bilibili.Execute[*BatchModifyRelationResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// GetUserRelationParam 是查询用户与自己关系的请求参数。
type GetUserRelationParam struct {
	Fid int `json:"fid"` // 目标用户 mid
}

// GetUserRelation 查询用户与自己关系（仅关注）。
func (c *Client) GetUserRelation(ctx context.Context, param GetUserRelationParam) (*GetUserRelationResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation"
	)
	return bilibili.Execute[*GetUserRelationResult](ctx, c.Client, method, url, param)
}

// GetUserRelation2Param 是 GetUserSpaceDetailParam 的别名。
type GetUserRelation2Param = GetUserSpaceDetailParam

// GetUserRelation2 查询用户与自己关系（互相关系）。
func (c *Client) GetUserRelation2(ctx context.Context, param GetUserRelation2Param) (*GetUserRelation2Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/acc/relation"
	)
	return bilibili.Execute[*GetUserRelation2Result](ctx, c.Client, method, url, param, c.FillWbi())
}

// BatchGetUserRelationParam 是批量查询用户与自己关系的请求参数。
type BatchGetUserRelationParam struct {
	Fids []int `json:"fids"` // 目标用户 mid
}

// BatchGetUserRelation 批量查询用户与自己关系。
func (c *Client) BatchGetUserRelation(ctx context.Context, param BatchGetUserRelationParam) (map[int]*BatchGetUserRelationDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/relations"
	)
	return bilibili.Execute[map[int]*BatchGetUserRelationDetail](ctx, c.Client, method, url, param)
}

// GetRelationTags 查询关注分组列表。
func (c *Client) GetRelationTags(ctx context.Context) ([]GetRelationTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/tags"
	)
	return bilibili.Execute[[]GetRelationTagsItem](ctx, c.Client, method, url, nil)
}
