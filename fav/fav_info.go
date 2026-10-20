package fav

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 查询用户收藏的视频收藏夹。返回体独立定义，不与其它接口复用。

// GetCollectedFavourFoldersParam 是查询用户收藏的视频收藏夹的请求参数。
type GetCollectedFavourFoldersParam struct {
	Ps       int    `json:"ps"`                                           // 每页项数。定义域：1-大于70
	Pn       int    `json:"pn"`                                           // 页码
	UpMid    int    `json:"up_mid"`                                       // 目标用户mid
	Platform string `json:"platform,omitempty" request:"query,omitempty"` // 平台类型。填写 web 时返回值才会包含用户收藏的视频合集
}

// GetCollectedFavourFoldersResult 是用户收藏的收藏夹列表（data 对象）。
// 收藏夹隐藏时 data 为 null，此时返回的指针为 nil。
type GetCollectedFavourFoldersResult struct {
	Count int                             `json:"count"` // 创建的收藏夹数
	List  []GetCollectedFavourFoldersItem `json:"list"`  // 收藏夹列表。无收藏夹时为 null
}

// GetCollectedFavourFoldersItem 是用户收藏的一个收藏夹。
type GetCollectedFavourFoldersItem struct {
	ID              int                                `json:"id"`                // 收藏夹mlid（完整id）
	Fid             int                                `json:"fid"`               // 原始收藏夹mlid，去除两位mid尾号
	Mid             int                                `json:"mid"`               // 创建用户mid
	Attr            int                                `json:"attr"`              // 收藏夹属性
	Title           string                             `json:"title"`             // 收藏夹标题
	Cover           string                             `json:"cover"`             // 收藏夹封面图片url
	Upper           GetCollectedFavourFoldersItemUpper `json:"upper"`             // 收藏夹创建用户信息
	CoverType       int                                `json:"cover_type"`        // 作用尚不明确。实测为 2
	Intro           string                             `json:"intro"`             // 作用尚不明确。实测为空
	Ctime           int                                `json:"ctime"`             // 创建时间戳
	Mtime           int                                `json:"mtime"`             // 审核时间戳
	State           int                                `json:"state"`             // 状态。0：正常。1：收藏夹已失效
	FavState        int                                `json:"fav_state"`         // 作用尚不明确。实测为 0
	MediaCount      int                                `json:"media_count"`       // 收藏夹总计视频数
	IsTop           bool                               `json:"is_top"`            // 是否置顶。实测为 false
	RecentFav       any                                `json:"recent_fav"`        // 最近收藏。实测为 null
	PlaySwitch      int                                `json:"play_switch"`       // 播放开关。实测为 0
	IsKidPlaylist   bool                               `json:"is_kid_playlist"`   // 是否为儿童播放列表。实测为 false
	KidPlaylistDesc string                             `json:"kid_playlist_desc"` // 儿童播放列表说明。实测为空
}

// GetCollectedFavourFoldersItemUpper 是用户收藏的收藏夹的创建用户信息。
type GetCollectedFavourFoldersItemUpper struct {
	Mid  int    `json:"mid"`  // 创建人mid
	Name string `json:"name"` // 创建人昵称
	Face string `json:"face"` // 空。作用尚不明确
}

// GetCollectedFavourFolders 查询用户收藏的视频收藏夹
func (c *Client) GetCollectedFavourFolders(ctx context.Context, param GetCollectedFavourFoldersParam) (*GetCollectedFavourFoldersResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v3/fav/folder/collected/list"
	)
	return bilibili.Execute[*GetCollectedFavourFoldersResult](ctx, c.Client, method, url, param)
}
