package bilibili

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// 收藏夹相关接口。响应模型见 fav_model.go。

type AddFavourFolderParam struct {
	Title   string `json:"title"`                                       // 收藏夹标题
	Intro   string `json:"intro,omitempty" request:"query,omitempty"`   // 收藏夹简介。默认为空
	Privacy int    `json:"privacy,omitempty" request:"query,omitempty"` // 是否公开。默认为公开。0：公开。1：私密
	Cover   string `json:"cover,omitempty" request:"query,omitempty"`   // 封面图url。封面会被审核
}

// AddFavourFolder 新建收藏夹
func (c *Client) AddFavourFolder(ctx context.Context, param AddFavourFolderParam) (*FavourFolderInfo, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/folder/add"
	)
	return execute[*FavourFolderInfo](ctx, c, method, url, param, fillCsrf(c))
}

type EditFavourFolderParam struct {
	MediaId int    `json:"media_id"`                                    // 目标收藏夹mdid
	Title   string `json:"title"`                                       // 修改收藏夹标题
	Intro   string `json:"intro,omitempty" request:"query,omitempty"`   // 修改收藏夹简介
	Privacy int    `json:"privacy,omitempty" request:"query,omitempty"` // 是否公开。默认为公开。。0：公开。1：私密
	Cover   string `json:"cover,omitempty" request:"query,omitempty"`   // 封面图url。封面会被审核
}

// EditFavourFolder 修改收藏夹
func (c *Client) EditFavourFolder(ctx context.Context, param EditFavourFolderParam) (*FavourFolderInfo, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/folder/edit"
	)
	return execute[*FavourFolderInfo](ctx, c, method, url, param, fillCsrf(c))
}

type DeleteFavourFolderParam struct {
	MediaIds []int `json:"media_ids"` // 目标收藏夹mdid列表
}

// DeleteFavourFolder 删除收藏夹
func (c *Client) DeleteFavourFolder(ctx context.Context, param DeleteFavourFolderParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/folder/del"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type MoveFavourResourcesParam struct {
	SrcMediaId int      `json:"src_media_id"`                                 // 源收藏夹id
	TarMediaId int      `json:"tar_media_id"`                                 // 目标收藏夹id
	Mid        int      `json:"mid"`                                          // 当前用户mid
	Resources  []string `json:"resources"`                                    // 目标内容id列表。格式：{内容id}:{内容类型}。类型：2：视频稿件。12：音频。21：视频合集。内容id：。视频稿件：视频稿件avid。音频：音频auid。视频合集：视频合集id
	Platform   string   `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web
}

// CopyFavourResources 批量复制收藏内容
func (c *Client) CopyFavourResources(ctx context.Context, param MoveFavourResourcesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/copy"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// MoveFavourResources 批量复制收藏内容
func (c *Client) MoveFavourResources(ctx context.Context, param MoveFavourResourcesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/move"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type DeleteFavourResourcesParam struct {
	Resources []int  `json:"resources"`                                    // 目标内容id列表。格式：{内容id}:{内容类型}。类型：2：视频稿件。12：音频。21：视频合集。内容id：。视频稿件：视频稿件avid。音频：音频auid。视频合集：视频合集id
	MediaId   int    `json:"media_id"`                                     // 目标收藏夹id
	Platform  string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web
}

// DeleteFavourResources 批量删除收藏内容
func (c *Client) DeleteFavourResources(ctx context.Context, param DeleteFavourResourcesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/batch-del"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type MediaIdParam struct {
	MediaId int `json:"media_id"` // 目标收藏夹id
}

// CleanFavourResources 清空所有失效收藏内容
func (c *Client) CleanFavourResources(ctx context.Context, param MediaIdParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/clean"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// GetFavourFolderInfo 获取收藏夹元数据
func (c *Client) GetFavourFolderInfo(ctx context.Context, param MediaIdParam) (*FavourFolderInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/info"
	)
	return execute[*FavourFolderInfo](ctx, c, method, url, param)
}

type GetAllFavourFolderInfoParam struct {
	UpMid int `json:"up_mid"`                                   // 目标用户mid
	Type  int `json:"type,omitempty" request:"query,omitempty"` // 目标内容属性。默认为全部。0：全部。2：视频稿件
	Rid   int `json:"rid,omitempty" request:"query,omitempty"`  // 目标内容id。视频稿件：视频稿件avid
}

// GetAllFavourFolderInfo 获取指定用户创建的所有收藏夹信息
func (c *Client) GetAllFavourFolderInfo(ctx context.Context, param GetAllFavourFolderInfoParam) (*AllFavourFolderInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/created/list-all"
	)
	return execute[*AllFavourFolderInfo](ctx, c, method, url, param)
}

type GetFavourInfoParam struct {
	Resources []string `json:"resources"`                                    // 目标内容id列表。格式：{内容id}:{内容类型}。类型：2：视频稿件。12：音频。21：视频合集。内容id：视频稿件：视频稿件avid。音频：音频auid。视频合集：视频合集id。注意：一次最多只能请求100个内容id，超过100个内容id将不放回数据。
	Platform  string   `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web（影响内容列表类型）
}

// GetFavourInfo 获取收藏内容
func (c *Client) GetFavourInfo(ctx context.Context, param GetFavourInfoParam) ([]FavourInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/resource/infos"
	)
	return execute[[]FavourInfo](ctx, c, method, url, param)
}

type GetFavourListParam struct {
	MediaId  int    `json:"media_id"`                                     // 目标收藏夹mlid（完整id）
	Tid      int    `json:"tid,omitempty" request:"query,omitempty"`      // 分区tid。默认为全部分区。0：全部分区
	Keyword  string `json:"keyword,omitempty" request:"query,omitempty"`  // 搜索关键字
	Order    string `json:"order,omitempty" request:"query,omitempty"`    // 排序方式。按收藏时间:mtime。按播放量: view。按投稿时间：pubtime
	Type     int    `json:"type,omitempty" request:"query,omitempty"`     // 查询范围。0：当前收藏夹（对应media_id）。 1：全部收藏夹
	Ps       int    `json:"ps"`                                           // 每页数量。定义域：1-20
	Pn       int    `json:"pn,omitempty" request:"query,omitempty"`       // 页码。默认为1
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web（影响内容列表类型）
}

// GetFavourList 获取收藏夹内容明细列表
func (c *Client) GetFavourList(ctx context.Context, param GetFavourListParam) (*FavourList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/resource/list"
	)
	return execute[*FavourList](ctx, c, method, url, param)
}

type GetFavourIdsParam struct {
	MediaId  int    `json:"media_id"`                                     // 目标收藏夹mlid（完整id）
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web（影响内容列表类型）
}

// GetFavourIds 获取收藏夹全部内容id
func (c *Client) GetFavourIds(ctx context.Context, param GetFavourIdsParam) ([]FavourId, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/resource/ids"
	)
	return execute[[]FavourId](ctx, c, method, url, param)
}

// GetSelfFavourList 获取自己的收藏夹列表
func (c *Client) GetSelfFavourList(ctx context.Context) ([]SelfFavourList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/list4navigate"
	)
	return execute[[]SelfFavourList](ctx, c, method, url, nil)
}
