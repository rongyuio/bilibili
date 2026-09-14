package bilibili

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// 用户与关系相关接口。响应模型见 user_model.go。

type GetUserVideosParam struct {
	Mid     int    `json:"mid"`                                         // 目标用户mid
	Order   string `json:"order,omitempty" request:"query,omitempty"`   // 排序方式。默认为pubdate。最新发布：pubdate。最多播放：click。最多收藏：stow
	Tid     int    `json:"tid,omitempty" request:"query,omitempty"`     // 筛选目标分区。默认为0。0：不进行分区筛选。分区tid为所筛选的分区
	Keyword string `json:"keyword,omitempty" request:"query,omitempty"` // 关键词筛选。用于使用关键词搜索该UP主视频稿件
	Pn      int    `json:"pn,omitempty" request:"query,omitempty"`      // 页码。默认为 1
	Ps      int    `json:"ps,omitempty" request:"query,omitempty"`      // 每页项数。默认为 30
}

// GetUserVideos 查询用户投稿视频明细
func (c *Client) GetUserVideos(ctx context.Context, param GetUserVideosParam) (*UserVideos, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/arc/search"
	)
	return execute[*UserVideos](ctx, c, method, url, param, c.fillWbi())
}

// GetUserSpaceDetail 获取用户空间详细信息
func (c *Client) GetUserSpaceDetail(ctx context.Context, param GetUserSpaceDetailParam) (*UserSpaceDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/acc/info"
	)
	return execute[*UserSpaceDetail](ctx, c, method, url, param, c.fillWbi())
}

type GetUserCardParam struct {
	Mid   int  `json:"mid"`                                       // 目标用户mid
	Photo bool `json:"photo,omitempty" request:"query,omitempty"` // 是否请求用户主页头图。true：是。false：否
}

// GetUserCard 获取用户用户名片 免登录
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/user/info.md#%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87%E4%BF%A1%E6%81%AF
func (c *Client) GetUserCard(ctx context.Context, param GetUserCardParam) (*UserCard, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/card"
	)
	return execute[*UserCard](ctx, c, method, url, param)
}

// GetMyUserSpaceDetail 获取登录用户空间详细信息
func (c *Client) GetMyUserSpaceDetail(ctx context.Context) (*MyUserSpaceDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/myinfo"
	)
	return execute[*MyUserSpaceDetail](ctx, c, method, url, nil)
}

type CheckNickNameParam struct {
	Nickname string `json:"nickName"` // 目标昵称。最长为16字符
}

// CheckNickName 检查昵称是否可注册
func (c *Client) CheckNickName(ctx context.Context, param CheckNickNameParam) error {
	const (
		method = resty.MethodGet
		url    = "https://passport.bilibili.com/web/generic/check/nickname"
	)
	_, err := execute[any](ctx, c, method, url, param)
	return err
}

type JoinOldFansParam struct {
	Aid      string `json:"aid,omitempty" request:"query,omitempty"`      // 空串
	UpMid    string `json:"up_mid"`                                       // UP主UID
	Source   string `json:"source,omitempty" request:"query,omitempty"`   // "4"
	Scene    string `json:"scene,omitempty" request:"query,omitempty"`    // "105"
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // "web"
	MobiApp  string `json:"mobi_app,omitempty" request:"query,omitempty"` // "pc"
}

// JoinOldFans 加入老粉计划
func (c *Client) JoinOldFans(ctx context.Context, param JoinOldFansParam) (*JoinOldFansResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v1/contract/add_contract"
	)
	return execute[*JoinOldFansResult](ctx, c, method, url, param, fillCsrf(c))
}

type FansSendMessageParam struct {
	Aid     string `json:"aid,omitempty" request:"query,omitempty"`    // 空串
	UpMid   string `json:"up_mid"`                                     // UP主UID
	Source  string `json:"source,omitempty" request:"query,omitempty"` // "4"
	Scene   string `json:"scene,omitempty" request:"query,omitempty"`  // "105"
	Content string `json:"content"`                                    // 留言内容
}

// FansSendMessage 老粉计划发送留言
func (c *Client) FansSendMessage(ctx context.Context, param FansSendMessageParam) (*FansSendMessageResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v1/contract/add_message"
	)
	return execute[*FansSendMessageResult](ctx, c, method, url, param, fillCsrf(c))
}

type BatchGetUserCardsParam struct {
	Uids []int `json:"uids"` // 目标用户的UID列表
}

// BatchGetUserCards 获取多用户详细信息
func (c *Client) BatchGetUserCards(ctx context.Context, param BatchGetUserCardsParam) ([]*BatchGetUserCardsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/account/v1/user/cards"
	)
	return execute[[]*BatchGetUserCardsResult](ctx, c, method, url, param)
}

type GetUserFollowersParam struct {
	Vmid int `json:"vmid"`                                   // 目标用户 mid
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 50
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1。仅可查看前 1000 名粉丝
}

// GetUserFollowers 查询用户粉丝明细（需要登录）
func (c *Client) GetUserFollowers(ctx context.Context, param GetUserFollowersParam) (*GetUserFollowersResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followers"
	)
	return execute[*GetUserFollowersResult](ctx, c, method, url, param)
}

type GetUserFollowingsParam struct {
	Vmid      int    `json:"vmid"`                                           // 目标用户 mid
	OrderType string `json:"order_type,omitempty" request:"query,omitempty"` // 排序方式。当目标用户为自己时有效。按照关注顺序排列：留空。按照最常访问排列：attention
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页项数。默认为 50
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为 1。其他用户仅可查看前 100 个
}

// GetUserFollowings 查询用户关注明细（需要登录）
func (c *Client) GetUserFollowings(ctx context.Context, param GetUserFollowingsParam) (*GetUserFollowingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followings"
	)
	return execute[*GetUserFollowingsResult](ctx, c, method, url, param)
}

type GetUserFollowings2Param struct {
	Vmid  int    `json:"vmid"`                                      // 目标用户 mid
	Order string `json:"order,omitempty" request:"query,omitempty"` // 排序方式。按照降序排列：desc。按照升序排列：asc。默认降序排列
	Ps    int    `json:"ps,omitempty" request:"query,omitempty"`    // 每页项数。默认为 50
	Pn    int    `json:"pn,omitempty" request:"query,omitempty"`    // 页码。默认为 1。仅可查看前 5 页
}

// GetUserFollowings2 查询用户关注明细2
//
// 仅可查看前 5 页，对于已设置可见性隐私关注列表的用户，则返回的List为nil，Total为0
func (c *Client) GetUserFollowings2(ctx context.Context, param GetUserFollowings2Param) (*GetUserFollowings2Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://app.biliapi.net/x/v2/relation/followings"
	)
	return execute[*GetUserFollowings2Result](ctx, c, method, url, param)
}

type GetUserFollowings3Param struct {
	Vmid int `json:"vmid"`                                   // 目标用户mid
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为20
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
}

// GetUserFollowings3 查询用户关注明细3
//
// 对于设置了可见性隐私关注列表的用户会返回空列表
func (c *Client) GetUserFollowings3(ctx context.Context, param GetUserFollowings3Param) (*GetUserFollowings3Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://line3-h5-mobile-api.biligame.com/game/center/h5/user/relationship/following_list"
	)
	return execute[*GetUserFollowings3Result](ctx, c, method, url, param)
}

type SearchUserFollowingsParam struct {
	Vmid string `json:"vmid"`                                     // 目标用户 mid
	Name string `json:"name,omitempty" request:"query,omitempty"` // 搜索关键词
	Ps   int    `json:"ps,omitempty" request:"query,omitempty"`   // 每页项数。默认为 50
	Pn   int    `json:"pn,omitempty" request:"query,omitempty"`   // 页码。默认为 1
}

// SearchUserFollowings 搜索关注明细
func (c *Client) SearchUserFollowings(ctx context.Context, param SearchUserFollowingsParam) (*SearchUserFollowingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/followings/search"
	)
	return execute[*SearchUserFollowingsResult](ctx, c, method, url, param)
}

type GetSameFollowingsParam struct {
	Vmid int `json:"vmid"`                                   // 目标用户 mid
	Ps   int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 50
	Pn   int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
}

// GetSameFollowings 查询共同关注明细
func (c *Client) GetSameFollowings(ctx context.Context, param GetSameFollowingsParam) (*GetSameFollowingsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/same/followings"
	)
	return execute[*GetSameFollowingsResult](ctx, c, method, url, param)
}

// GetWhispers 查询悄悄关注明细
func (c *Client) GetWhispers(ctx context.Context) (*GetWhispersResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/whispers"
	)
	return execute[*GetWhispersResult](ctx, c, method, url, nil)
}

// GetFriends 查询互相关注明细
func (c *Client) GetFriends(ctx context.Context) (*GetFriendsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/friends"
	)
	return execute[*GetFriendsResult](ctx, c, method, url, nil)
}

type GetBlacksParam struct {
	Ps int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数。默认为 50
	Pn int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
}

// GetBlacks 查询黑名单明细
func (c *Client) GetBlacks(ctx context.Context, param GetBlacksParam) (*GetBlacksResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/blacks"
	)
	return execute[*GetBlacksResult](ctx, c, method, url, param)
}

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

type ModifyRelationParam struct {
	Fid   int               `json:"fid"`    // 目标用户mid
	Act   ModifyRelationAct `json:"act"`    // 操作代码
	ReSrc int               `json:"re_src"` // 关注来源代码。空间：11。视频：14。文章：115。活动页面：222
}

// ModifyRelation 操作用户关系
func (c *Client) ModifyRelation(ctx context.Context, param ModifyRelationParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/modify"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type BatchModifyRelationParam struct {
	Fids  []int             `json:"fids"`   // 目标用户 mid 列表
	Act   ModifyRelationAct `json:"act"`    // 操作代码。仅可为 1 或 5，故只能进行批量关注和拉黑
	ReSrc int               `json:"re_src"` // 关注来源代码。同上
}

// BatchModifyRelation 批量操作用户关系
func (c *Client) BatchModifyRelation(ctx context.Context, param BatchModifyRelationParam) (*BatchModifyRelationResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/relation/batch/modify"
	)
	return execute[*BatchModifyRelationResult](ctx, c, method, url, param, fillCsrf(c))
}

type GetUserRelationParam struct {
	Fid int `json:"fid"` // 目标用户 mid
}

// GetUserRelation 查询用户与自己关系（仅关注）
func (c *Client) GetUserRelation(ctx context.Context, param GetUserRelationParam) (*RelationDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation"
	)
	return execute[*RelationDetail](ctx, c, method, url, param)
}

type GetUserRelation2Param struct {
	Mid int `json:"mid"` // 目标用户mid
}

// GetUserRelation2 查询用户与自己关系（互相关系）
func (c *Client) GetUserRelation2(ctx context.Context, param GetUserRelation2Param) (*GetUserRelation2Result, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/space/wbi/acc/relation"
	)
	return execute[*GetUserRelation2Result](ctx, c, method, url, param, c.fillWbi())
}

type BatchGetUserRelationParam struct {
	Fids []int `json:"fids"` // 目标用户 mid
}

// BatchGetUserRelation 批量查询用户与自己关系
func (c *Client) BatchGetUserRelation(ctx context.Context, param BatchGetUserRelationParam) (map[int]*RelationDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/relations"
	)
	return execute[map[int]*RelationDetail](ctx, c, method, url, param)
}

// GetRelationTags 查询关注分组列表
func (c *Client) GetRelationTags(ctx context.Context) ([]RelationTag, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/relation/tags"
	)
	return execute[[]RelationTag](ctx, c, method, url, nil)
}
