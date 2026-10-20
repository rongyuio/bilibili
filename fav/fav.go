package fav

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 收藏夹相关接口。各接口的返回体独立定义，互不复用。

// AddFavourFolderParam 是新建收藏夹的请求参数。
type AddFavourFolderParam struct {
	Title   string `json:"title"`                                       // 收藏夹标题
	Intro   string `json:"intro,omitempty" request:"query,omitempty"`   // 收藏夹简介。默认为空
	Privacy int    `json:"privacy,omitempty" request:"query,omitempty"` // 是否公开。默认为公开。0：公开。1：私密
	Cover   string `json:"cover,omitempty" request:"query,omitempty"`   // 封面图url。封面会被审核
}

// AddFavourFolderResult 是新建收藏夹接口的 data 对象。
type AddFavourFolderResult struct {
	ID              int                    `json:"id"`                // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid             int                    `json:"fid"`               // 收藏夹原始id
	Mid             int                    `json:"mid"`               // 创建者mid
	Attr            int                    `json:"attr"`              // 属性位（？）
	Title           string                 `json:"title"`             // 收藏夹标题
	Cover           string                 `json:"cover"`             // 收藏夹封面图片url
	Upper           AddFavourFolderUpper   `json:"upper"`             // 创建者信息
	CoverType       int                    `json:"cover_type"`        // 封面图类别（？）
	CntInfo         AddFavourFolderCntInfo `json:"cnt_info"`          // 收藏夹状态数
	Type            int                    `json:"type"`              // 类型（？）
	Intro           string                 `json:"intro"`             // 备注
	Ctime           int                    `json:"ctime"`             // 创建时间戳
	Mtime           int                    `json:"mtime"`             // 收藏时间戳
	State           int                    `json:"state"`             // 状态（？）
	FavState        int                    `json:"fav_state"`         // 收藏夹收藏状态，已收藏：1，未收藏：0
	LikeState       int                    `json:"like_state"`        // 点赞状态，已点赞：1，未点赞：0
	MediaCount      int                    `json:"media_count"`       // 收藏夹内容数量
	IsTop           bool                   `json:"is_top"`            // 是否置顶。实测为 false
	IsKidPlaylist   bool                   `json:"is_kid_playlist"`   // 是否为儿童播放列表。实测为 false
	KidPlaylistDesc string                 `json:"kid_playlist_desc"` // 儿童播放列表说明。实测为空
}

// AddFavourFolderUpper 是新建收藏夹接口返回的创建者信息。
type AddFavourFolderUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// AddFavourFolderCntInfo 是新建收藏夹接口返回的收藏夹状态数。
type AddFavourFolderCntInfo struct {
	Collect int `json:"collect"`  // 收藏数
	Play    int `json:"play"`     // 播放数
	ThumbUp int `json:"thumb_up"` // 点赞数
	Share   int `json:"share"`    // 分享数
}

// AddFavourFolder 新建收藏夹
func (c *Client) AddFavourFolder(ctx context.Context, param AddFavourFolderParam) (*AddFavourFolderResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/folder/add"
	)
	return bilibili.Execute[*AddFavourFolderResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// EditFavourFolderParam 是修改收藏夹的请求参数。
type EditFavourFolderParam struct {
	MediaID int    `json:"media_id"`                                    // 目标收藏夹mdid
	Title   string `json:"title"`                                       // 修改收藏夹标题
	Intro   string `json:"intro,omitempty" request:"query,omitempty"`   // 修改收藏夹简介
	Privacy int    `json:"privacy,omitempty" request:"query,omitempty"` // 是否公开。默认为公开。。0：公开。1：私密
	Cover   string `json:"cover,omitempty" request:"query,omitempty"`   // 封面图url。封面会被审核
}

// EditFavourFolderResult 是修改收藏夹接口的 data 对象。
type EditFavourFolderResult struct {
	ID              int                     `json:"id"`                // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid             int                     `json:"fid"`               // 收藏夹原始id
	Mid             int                     `json:"mid"`               // 创建者mid
	Attr            int                     `json:"attr"`              // 属性位（？）
	Title           string                  `json:"title"`             // 收藏夹标题
	Cover           string                  `json:"cover"`             // 收藏夹封面图片url
	Upper           EditFavourFolderUpper   `json:"upper"`             // 创建者信息
	CoverType       int                     `json:"cover_type"`        // 封面图类别（？）
	CntInfo         EditFavourFolderCntInfo `json:"cnt_info"`          // 收藏夹状态数
	Type            int                     `json:"type"`              // 类型（？）
	Intro           string                  `json:"intro"`             // 备注
	Ctime           int                     `json:"ctime"`             // 创建时间戳
	Mtime           int                     `json:"mtime"`             // 收藏时间戳
	State           int                     `json:"state"`             // 状态（？）
	FavState        int                     `json:"fav_state"`         // 收藏夹收藏状态，已收藏：1，未收藏：0
	LikeState       int                     `json:"like_state"`        // 点赞状态，已点赞：1，未点赞：0
	MediaCount      int                     `json:"media_count"`       // 收藏夹内容数量
	IsTop           bool                    `json:"is_top"`            // 是否置顶。实测为 false
	IsKidPlaylist   bool                    `json:"is_kid_playlist"`   // 是否为儿童播放列表。实测为 false
	KidPlaylistDesc string                  `json:"kid_playlist_desc"` // 儿童播放列表说明。实测为空
}

// EditFavourFolderUpper 是修改收藏夹接口返回的创建者信息。
type EditFavourFolderUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// EditFavourFolderCntInfo 是修改收藏夹接口返回的收藏夹状态数。
type EditFavourFolderCntInfo struct {
	Collect int `json:"collect"`  // 收藏数
	Play    int `json:"play"`     // 播放数
	ThumbUp int `json:"thumb_up"` // 点赞数
	Share   int `json:"share"`    // 分享数
}

// EditFavourFolder 修改收藏夹
func (c *Client) EditFavourFolder(ctx context.Context, param EditFavourFolderParam) (*EditFavourFolderResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/folder/edit"
	)
	return bilibili.Execute[*EditFavourFolderResult](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
}

// DeleteFavourFolderParam 是删除收藏夹的请求参数。
type DeleteFavourFolderParam struct {
	MediaIDs []int `json:"media_ids"` // 目标收藏夹mdid列表
}

// DeleteFavourFolder 删除收藏夹
func (c *Client) DeleteFavourFolder(ctx context.Context, param DeleteFavourFolderParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/folder/del"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// MoveFavourResourcesParam 是批量复制或移动收藏内容的请求参数。
type MoveFavourResourcesParam struct {
	SrcMediaID int      `json:"src_media_id"`                                 // 源收藏夹id
	TarMediaID int      `json:"tar_media_id"`                                 // 目标收藏夹id
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
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// MoveFavourResources 批量复制收藏内容
func (c *Client) MoveFavourResources(ctx context.Context, param MoveFavourResourcesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/move"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// DeleteFavourResourcesParam 是批量删除收藏内容的请求参数。
type DeleteFavourResourcesParam struct {
	Resources []int  `json:"resources"`                                    // 目标内容id列表。格式：{内容id}:{内容类型}。类型：2：视频稿件。12：音频。21：视频合集。内容id：。视频稿件：视频稿件avid。音频：音频auid。视频合集：视频合集id
	MediaID   int    `json:"media_id"`                                     // 目标收藏夹id
	Platform  string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web
}

// DeleteFavourResources 批量删除收藏内容
func (c *Client) DeleteFavourResources(ctx context.Context, param DeleteFavourResourcesParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/batch-del"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// MediaIDParam 是只指定目标收藏夹 id 的请求参数。
type MediaIDParam struct {
	MediaID int `json:"media_id"` // 目标收藏夹id
}

// CleanFavourResources 清空所有失效收藏内容
func (c *Client) CleanFavourResources(ctx context.Context, param MediaIDParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/v3/fav/resource/clean"
	)
	_, err := bilibili.Execute[any](ctx, c.Client, method, url, param, bilibili.FillCsrf(c.Client))
	return err
}

// GetFavourFolderInfoResult 是获取收藏夹元数据接口的 data 对象。
type GetFavourFolderInfoResult struct {
	ID              int                        `json:"id"`                // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid             int                        `json:"fid"`               // 收藏夹原始id
	Mid             int                        `json:"mid"`               // 创建者mid
	Attr            int                        `json:"attr"`              // 属性位（？）
	Title           string                     `json:"title"`             // 收藏夹标题
	Cover           string                     `json:"cover"`             // 收藏夹封面图片url
	Upper           GetFavourFolderInfoUpper   `json:"upper"`             // 创建者信息
	CoverType       int                        `json:"cover_type"`        // 封面图类别（？）
	CntInfo         GetFavourFolderInfoCntInfo `json:"cnt_info"`          // 收藏夹状态数
	Type            int                        `json:"type"`              // 类型（？）
	Intro           string                     `json:"intro"`             // 备注
	Ctime           int                        `json:"ctime"`             // 创建时间戳
	Mtime           int                        `json:"mtime"`             // 收藏时间戳
	State           int                        `json:"state"`             // 状态（？）
	FavState        int                        `json:"fav_state"`         // 收藏夹收藏状态，已收藏：1，未收藏：0
	LikeState       int                        `json:"like_state"`        // 点赞状态，已点赞：1，未点赞：0
	MediaCount      int                        `json:"media_count"`       // 收藏夹内容数量
	IsTop           bool                       `json:"is_top"`            // 是否置顶。实测为 false
	IsKidPlaylist   bool                       `json:"is_kid_playlist"`   // 是否为儿童播放列表。实测为 false
	KidPlaylistDesc string                     `json:"kid_playlist_desc"` // 儿童播放列表说明。实测为空
}

// GetFavourFolderInfoUpper 是获取收藏夹元数据接口返回的创建者信息。
type GetFavourFolderInfoUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// GetFavourFolderInfoCntInfo 是获取收藏夹元数据接口返回的收藏夹状态数。
type GetFavourFolderInfoCntInfo struct {
	Collect int `json:"collect"`  // 收藏数
	Play    int `json:"play"`     // 播放数
	ThumbUp int `json:"thumb_up"` // 点赞数
	Share   int `json:"share"`    // 分享数
}

// GetFavourFolderInfo 获取收藏夹元数据
func (c *Client) GetFavourFolderInfo(ctx context.Context, param MediaIDParam) (*GetFavourFolderInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/info"
	)
	return bilibili.Execute[*GetFavourFolderInfoResult](ctx, c.Client, method, url, param)
}

// GetAllFavourFolderInfoParam 是获取指定用户创建的所有收藏夹信息的请求参数。
type GetAllFavourFolderInfoParam struct {
	UpMid int `json:"up_mid"`                                   // 目标用户mid
	Type  int `json:"type,omitempty" request:"query,omitempty"` // 目标内容属性。默认为全部。0：全部。2：视频稿件
	Rid   int `json:"rid,omitempty" request:"query,omitempty"`  // 目标内容id。视频稿件：视频稿件avid
}

// GetAllFavourFolderInfoResult 是获取指定用户创建的所有收藏夹信息接口的 data 对象。
type GetAllFavourFolderInfoResult struct {
	Count int                          `json:"count"` // 创建的收藏夹总数
	List  []GetAllFavourFolderInfoItem `json:"list"`  // 创建的收藏夹列表
}

// GetAllFavourFolderInfoItem 是用户创建的一个收藏夹摘要。
type GetAllFavourFolderInfoItem struct {
	ID         int    `json:"id"`          // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid        int    `json:"fid"`         // 收藏夹原始id
	Mid        int    `json:"mid"`         // 创建者mid
	Attr       int    `json:"attr"`        // 属性位（？）
	Title      string `json:"title"`       // 收藏夹标题
	FavState   int    `json:"fav_state"`   // 目标id是否存在于该收藏夹，存在于该收藏夹：1，不存在于该收藏夹：0
	MediaCount int    `json:"media_count"` // 收藏夹内容数量
}

// GetAllFavourFolderInfo 获取指定用户创建的所有收藏夹信息
func (c *Client) GetAllFavourFolderInfo(ctx context.Context, param GetAllFavourFolderInfoParam) (*GetAllFavourFolderInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/created/list-all"
	)
	return bilibili.Execute[*GetAllFavourFolderInfoResult](ctx, c.Client, method, url, param)
}

// GetFavourInfoParam 是获取收藏内容的请求参数。
type GetFavourInfoParam struct {
	Resources []string `json:"resources"`                                    // 目标内容id列表。格式：{内容id}:{内容类型}。类型：2：视频稿件。12：音频。21：视频合集。内容id：视频稿件：视频稿件avid。音频：音频auid。视频合集：视频合集id。注意：一次最多只能请求100个内容id，超过100个内容id将不放回数据。
	Platform  string   `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web（影响内容列表类型）
}

// GetFavourInfoItem 是获取收藏内容接口 data 数组中的一项。
type GetFavourInfoItem struct {
	ID       int                      `json:"id"`       // 内容id
	Type     int                      `json:"type"`     // 内容类型
	Title    string                   `json:"title"`    // 标题
	Cover    string                   `json:"cover"`    // 封面url
	Intro    string                   `json:"intro"`    // 简介
	Page     int                      `json:"page"`     // 视频分P数
	Duration int                      `json:"duration"` // 音频/视频时长
	Upper    GetFavourInfoItemUpper   `json:"upper"`    // UP主信息
	Attr     int                      `json:"attr"`     // 属性位（？）
	CntInfo  GetFavourInfoItemCntInfo `json:"cnt_info"` // 状态数
	Link     string                   `json:"link"`     // 跳转uri
	Ctime    int                      `json:"ctime"`    // 投稿时间戳
	Pubtime  int                      `json:"pubtime"`  // 发布时间戳
	FavTime  int                      `json:"fav_time"` // 收藏时间戳
	BvID     string                   `json:"bv_id"`    // 视频稿件bvid
	Bvid     string                   `json:"bvid"`     // 视频稿件bvid
	Season   any                      `json:"season"`   // （？）
	Ugc      GetFavourInfoItemUgc     `json:"ugc"`      // 稿件元信息
}

// GetFavourInfoItemUpper 是获取收藏内容接口返回的 UP 主信息。
type GetFavourInfoItemUpper struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像url
}

// GetFavourInfoItemCntInfo 是获取收藏内容接口返回的计数信息。
type GetFavourInfoItemCntInfo struct {
	Collect    int    `json:"collect"`     // 收藏数
	Play       int    `json:"play"`        // 播放数
	Danmaku    int    `json:"danmaku"`     // 弹幕数
	PlaySwitch int    `json:"play_switch"` // 播放开关。实测为 0
	ViewText1  string `json:"view_text_1"` // 播放数文案。实测为 "285.9万"
}

// GetFavourInfoItemUgc 是获取收藏内容接口返回的稿件元信息。
type GetFavourInfoItemUgc struct {
	FirstCid int `json:"first_cid"` // 视频cid
}

// GetFavourInfo 获取收藏内容
func (c *Client) GetFavourInfo(ctx context.Context, param GetFavourInfoParam) ([]GetFavourInfoItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/resource/infos"
	)
	return bilibili.Execute[[]GetFavourInfoItem](ctx, c.Client, method, url, param)
}

// GetFavourListParam 是获取收藏夹内容明细列表的请求参数。
type GetFavourListParam struct {
	MediaID  int    `json:"media_id"`                                     // 目标收藏夹mlid（完整id）
	Tid      int    `json:"tid,omitempty" request:"query,omitempty"`      // 分区tid。默认为全部分区。0：全部分区
	Keyword  string `json:"keyword,omitempty" request:"query,omitempty"`  // 搜索关键字
	Order    string `json:"order,omitempty" request:"query,omitempty"`    // 排序方式。按收藏时间:mtime。按播放量: view。按投稿时间：pubtime
	Type     int    `json:"type,omitempty" request:"query,omitempty"`     // 查询范围。0：当前收藏夹（对应media_id）。 1：全部收藏夹
	Ps       int    `json:"ps"`                                           // 每页数量。定义域：1-20
	Pn       int    `json:"pn,omitempty" request:"query,omitempty"`       // 页码。默认为1
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web（影响内容列表类型）
}

// GetFavourListResult 是获取收藏夹内容明细列表接口的 data 对象。
type GetFavourListResult struct {
	Info    GetFavourListInfo   `json:"info"`     // 收藏夹元数据
	Medias  []GetFavourListItem `json:"medias"`   // 收藏夹内容
	HasMore bool                `json:"has_more"` // 收藏夹是否有下一页
}

// GetFavourListInfo 是获取收藏夹内容明细列表接口返回的收藏夹元数据。
type GetFavourListInfo struct {
	ID              int                      `json:"id"`                // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid             int                      `json:"fid"`               // 收藏夹原始id
	Mid             int                      `json:"mid"`               // 创建者mid
	Attr            int                      `json:"attr"`              // 属性位（？）
	Title           string                   `json:"title"`             // 收藏夹标题
	Cover           string                   `json:"cover"`             // 收藏夹封面图片url
	Upper           GetFavourListInfoUpper   `json:"upper"`             // 创建者信息
	CoverType       int                      `json:"cover_type"`        // 封面图类别（？）
	CntInfo         GetFavourListInfoCntInfo `json:"cnt_info"`          // 收藏夹状态数
	Type            int                      `json:"type"`              // 类型（？）
	Intro           string                   `json:"intro"`             // 备注
	Ctime           int                      `json:"ctime"`             // 创建时间戳
	Mtime           int                      `json:"mtime"`             // 收藏时间戳
	State           int                      `json:"state"`             // 状态（？）
	FavState        int                      `json:"fav_state"`         // 收藏夹收藏状态，已收藏：1，未收藏：0
	LikeState       int                      `json:"like_state"`        // 点赞状态，已点赞：1，未点赞：0
	MediaCount      int                      `json:"media_count"`       // 收藏夹内容数量
	IsTop           bool                     `json:"is_top"`            // 是否置顶。实测为 false
	IsKidPlaylist   bool                     `json:"is_kid_playlist"`   // 是否为儿童播放列表。实测为 false
	KidPlaylistDesc string                   `json:"kid_playlist_desc"` // 儿童播放列表说明。实测为空
}

// GetFavourListInfoUpper 是获取收藏夹内容明细列表接口返回的创建者信息。
type GetFavourListInfoUpper struct {
	Mid       int    `json:"mid"`        // UP 主 mid
	Name      string `json:"name"`       // 创建者昵称
	Face      string `json:"face"`       // 创建者头像url
	Followed  bool   `json:"followed"`   // 是否已关注创建者
	VipType   int    `json:"vip_type"`   // 会员类别，0：无，1：月大会员，2：年度及以上大会员
	VipStatue int    `json:"vip_statue"` // 0：无，1：有
}

// GetFavourListInfoCntInfo 是获取收藏夹内容明细列表接口返回的收藏夹状态数。
type GetFavourListInfoCntInfo struct {
	Collect int `json:"collect"`  // 收藏数
	Play    int `json:"play"`     // 播放数
	ThumbUp int `json:"thumb_up"` // 点赞数
	Share   int `json:"share"`    // 分享数
}

// GetFavourListItem 是收藏夹内的一条内容。
type GetFavourListItem struct {
	ID       int                      `json:"id"`       // 内容id，视频稿件：视频稿件avid，音频：音频auid，视频合集：视频合集id
	Type     int                      `json:"type"`     // 内容类型，2：视频稿件，12：音频，21：视频合集
	Title    string                   `json:"title"`    // 标题
	Cover    string                   `json:"cover"`    // 封面url
	Intro    string                   `json:"intro"`    // 简介
	Page     int                      `json:"page"`     // 视频分P数
	Duration int                      `json:"duration"` // 音频/视频时长
	Upper    GetFavourListItemUpper   `json:"upper"`    // UP主信息
	Attr     int                      `json:"attr"`     // 属性位（？）
	CntInfo  GetFavourListItemCntInfo `json:"cnt_info"` // 状态数
	Link     string                   `json:"link"`     // 跳转uri
	Ctime    int                      `json:"ctime"`    // 投稿时间戳
	Pubtime  int                      `json:"pubtime"`  // 发布时间戳
	FavTime  int                      `json:"fav_time"` // 收藏时间戳
	BvID     string                   `json:"bv_id"`    // 视频稿件bvid
	Bvid     string                   `json:"bvid"`     // 视频稿件bvid
	Ugc      GetFavourListItemUgc     `json:"ugc"`      // 稿件元信息
	// 媒体列表跳转链接。实测为 bilibili://music/playlist/playpage/... 形式
	MediaListLink string `json:"media_list_link"`
	Season        any    `json:"season"` // （？）
	Ogv           any    `json:"ogv"`    // (?)。实测为 null
}

// GetFavourListItemUpper 是收藏夹内容中的 UP 主信息。
type GetFavourListItemUpper struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像url
}

// GetFavourListItemCntInfo 是收藏夹内容中的计数信息。
type GetFavourListItemCntInfo struct {
	Collect    int    `json:"collect"`     // 收藏数
	Play       int    `json:"play"`        // 播放数
	Danmaku    int    `json:"danmaku"`     // 弹幕数
	PlaySwitch int    `json:"play_switch"` // 播放开关。实测为 0
	ViewText1  string `json:"view_text_1"` // 播放数文案。实测为 "285.9万"
}

// GetFavourListItemUgc 是收藏夹内容中的稿件元信息。
type GetFavourListItemUgc struct {
	FirstCid int `json:"first_cid"` // 视频cid
}

// GetFavourList 获取收藏夹内容明细列表
func (c *Client) GetFavourList(ctx context.Context, param GetFavourListParam) (*GetFavourListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/resource/list"
	)
	return bilibili.Execute[*GetFavourListResult](ctx, c.Client, method, url, param)
}

// GetFavourIDsParam 是获取收藏夹全部内容id的请求参数。
type GetFavourIDsParam struct {
	MediaID  int    `json:"media_id"`                                     // 目标收藏夹mlid（完整id）
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台标识。可为web（影响内容列表类型）
}

// GetFavourIDsItem 是获取收藏夹全部内容id接口 data 数组中的一项。
type GetFavourIDsItem struct {
	ID   int    `json:"id"`    // 内容id，视频稿件：视频稿件avid，音频：音频auid，视频合集：视频合集id
	Type int    `json:"type"`  // 内容类型，2：视频稿件，12：音频，21：视频合集
	BvID string `json:"bv_id"` // 视频稿件bvid
	Bvid string `json:"bvid"`  // 视频稿件bvid
}

// GetFavourIDs 获取收藏夹全部内容id
func (c *Client) GetFavourIDs(ctx context.Context, param GetFavourIDsParam) ([]GetFavourIDsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/resource/ids"
	)
	return bilibili.Execute[[]GetFavourIDsItem](ctx, c.Client, method, url, param)
}

// GetSelfFavourListItem 是获取自己的收藏夹列表接口 data 数组中的一项。
type GetSelfFavourListItem struct {
	ID                int64                                  `json:"id"`                // 收藏夹mlid（完整id）
	Name              string                                 `json:"name"`              // 收藏夹标题
	MediaListResponse GetSelfFavourListItemMediaListResponse `json:"mediaListResponse"` // 收藏夹内容列表
	URI               string                                 `json:"uri"`               // 收藏夹 uri
}

// GetSelfFavourListItemMediaListResponse 是自身收藏夹的内容列表。
type GetSelfFavourListItemMediaListResponse struct {
	Count   int                          `json:"count"`    // 内容数量
	List    []GetSelfFavourListItemMedia `json:"list"`     // 内容列表
	HasMore bool                         `json:"has_more"` // 是否有下一页
}

// GetSelfFavourListItemMedia 是自身收藏夹中的一条内容。
type GetSelfFavourListItemMedia struct {
	ID         int64                           `json:"id"`          // 内容id
	Fid        int                             `json:"fid"`         // 收藏夹原始id
	Mid        int                             `json:"mid"`         // 创建者mid
	Attr       int                             `json:"attr"`        // 属性位（？）
	AttrDesc   string                          `json:"attr_desc"`   // 属性描述
	Title      string                          `json:"title"`       // 标题
	Cover      string                          `json:"cover"`       // 封面url
	Upper      GetSelfFavourListItemMediaUpper `json:"upper"`       // UP主信息
	CoverType  int                             `json:"cover_type"`  // 封面图类别
	Intro      string                          `json:"intro"`       // 简介
	Ctime      int                             `json:"ctime"`       // 创建时间戳
	Mtime      int                             `json:"mtime"`       // 收藏时间戳
	State      int                             `json:"state"`       // 状态
	FavState   int                             `json:"fav_state"`   // 收藏夹收藏状态
	MediaCount int                             `json:"media_count"` // 收藏夹内容数量
	ViewCount  int                             `json:"view_count"`  // 播放数
	Vt         int                             `json:"vt"`          // 作用尚不明确
	IsTop      bool                            `json:"is_top"`      // 是否置顶。实测为 false
	RecentFav  any                             `json:"recent_fav"`  // 最近收藏。实测为 null
	PlaySwitch int                             `json:"play_switch"` // 播放开关。实测为 0
	Type       int                             `json:"type"`        // 内容类型
	Link       string                          `json:"link"`        // 跳转uri
	Bvid       string                          `json:"bvid"`        // 视频稿件bvid
}

// GetSelfFavourListItemMediaUpper 是自身收藏夹内容中的 UP 主信息。
type GetSelfFavourListItemMediaUpper struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像url
}

// GetSelfFavourList 获取自己的收藏夹列表
func (c *Client) GetSelfFavourList(ctx context.Context) ([]GetSelfFavourListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/list4navigate"
	)
	return bilibili.Execute[[]GetSelfFavourListItem](ctx, c.Client, method, url, nil)
}
