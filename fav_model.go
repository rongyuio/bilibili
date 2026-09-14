package bilibili

// 收藏夹相关响应模型。

type CntInfo struct {
	Collect int `json:"collect"`  // 收藏数
	Play    int `json:"play"`     // 播放数
	ThumbUp int `json:"thumb_up"` // 点赞数
	Share   int `json:"share"`    // 分享数
}

type FavourFolderInfo struct {
	ID         int     `json:"id"`          // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid        int     `json:"fid"`         // 收藏夹原始id
	Mid        int     `json:"mid"`         // 创建者mid
	Attr       int     `json:"attr"`        // 属性位（？）
	Title      string  `json:"title"`       // 收藏夹标题
	Cover      string  `json:"cover"`       // 收藏夹封面图片url
	Upper      Upper   `json:"upper"`       // 创建者信息
	CoverType  int     `json:"cover_type"`  // 封面图类别（？）
	CntInfo    CntInfo `json:"cnt_info"`    // 收藏夹状态数
	Type       int     `json:"type"`        // 类型（？）
	Intro      string  `json:"intro"`       // 备注
	Ctime      int     `json:"ctime"`       // 创建时间戳
	Mtime      int     `json:"mtime"`       // 收藏时间戳
	State      int     `json:"state"`       // 状态（？）
	FavState   int     `json:"fav_state"`   // 收藏夹收藏状态，已收藏：1，未收藏：0
	LikeState  int     `json:"like_state"`  // 点赞状态，已点赞：1，未点赞：0
	MediaCount int     `json:"media_count"` // 收藏夹内容数量
}

type AllFavourFolderInfo struct {
	Count int                   `json:"count"` // 创建的收藏夹总数
	List  []AllFavourFolderItem `json:"list"`  // 创建的收藏夹列表
}

// AllFavourFolderItem 是用户创建的一个收藏夹摘要。
type AllFavourFolderItem struct {
	ID         int    `json:"id"`          // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid        int    `json:"fid"`         // 收藏夹原始id
	Mid        int    `json:"mid"`         // 创建者mid
	Attr       int    `json:"attr"`        // 属性位（？）
	Title      string `json:"title"`       // 收藏夹标题
	FavState   int    `json:"fav_state"`   // 目标id是否存在于该收藏夹，存在于该收藏夹：1，不存在于该收藏夹：0
	MediaCount int    `json:"media_count"` // 收藏夹内容数量
}

// FavourUpper 是 Owner 的别名（收藏夹内容创建者的简要信息）。
type FavourUpper = Owner

// FavourResourceCntInfo 是收藏内容的计数信息（含弹幕数）。
type FavourResourceCntInfo struct {
	Collect int `json:"collect"`
	Play    int `json:"play"`
	Danmaku int `json:"danmaku"`
}

// FavourUgc 是收藏内容的稿件元信息。
type FavourUgc struct {
	FirstCid int `json:"first_cid"` // 视频cid
}

type FavourInfo struct {
	ID       int                   `json:"id"`
	Type     int                   `json:"type"`
	Title    string                `json:"title"`
	Cover    string                `json:"cover"`
	Intro    string                `json:"intro"`
	Page     int                   `json:"page"`
	Duration int                   `json:"duration"`
	Upper    FavourUpper           `json:"upper"`
	Attr     int                   `json:"attr"`
	CntInfo  FavourResourceCntInfo `json:"cnt_info"`
	Link     string                `json:"link"`
	Ctime    int                   `json:"ctime"`
	Pubtime  int                   `json:"pubtime"`
	FavTime  int                   `json:"fav_time"`
	BvID     string                `json:"bv_id"`
	Bvid     string                `json:"bvid"`
	Season   any                   `json:"season"`
	Ugc      FavourUgc             `json:"ugc"`
}

// FavourFolderDetail 是收藏夹元数据（GetFavourList 的 info 字段）。
type FavourFolderDetail struct {
	ID         int     `json:"id"`          // 收藏夹mlid（完整id），收藏夹原始id+创建者mid尾号2位
	Fid        int     `json:"fid"`         // 收藏夹原始id
	Mid        int     `json:"mid"`         // 创建者mid
	Attr       int     `json:"attr"`        // 属性，0：正常，1：失效
	Title      string  `json:"title"`       // 收藏夹标题
	Cover      string  `json:"cover"`       // 收藏夹封面图片url
	Upper      Upper   `json:"upper"`       // 创建者信息
	CoverType  int     `json:"cover_type"`  // 封面图类别（？）
	CntInfo    CntInfo `json:"cnt_info"`    // 收藏夹状态数
	Type       int     `json:"type"`        // 类型（？），一般是11
	Intro      string  `json:"intro"`       // 备注
	Ctime      int     `json:"ctime"`       // 创建时间戳
	Mtime      int     `json:"mtime"`       // 收藏时间戳
	State      int     `json:"state"`       // 状态（？），一般为0
	FavState   int     `json:"fav_state"`   // 收藏夹收藏状态，已收藏收藏夹：1，未收藏收藏夹：0
	LikeState  int     `json:"like_state"`  // 点赞状态，已点赞：1，未点赞：0
	MediaCount int     `json:"media_count"` // 收藏夹内容数量
}

// FavourMedia 是收藏夹内的一条内容。
type FavourMedia struct {
	ID       int                   `json:"id"`       // 内容id，视频稿件：视频稿件avid，音频：音频auid，视频合集：视频合集id
	Type     int                   `json:"type"`     // 内容类型，2：视频稿件，12：音频，21：视频合集
	Title    string                `json:"title"`    // 标题
	Cover    string                `json:"cover"`    // 封面url
	Intro    string                `json:"intro"`    // 简介
	Page     int                   `json:"page"`     // 视频分P数
	Duration int                   `json:"duration"` // 音频/视频时长
	Upper    FavourUpper           `json:"upper"`    // UP主信息
	Attr     int                   `json:"attr"`     // 属性位（？）
	CntInfo  FavourResourceCntInfo `json:"cnt_info"` // 状态数
	Link     string                `json:"link"`     // 跳转uri
	Ctime    int                   `json:"ctime"`    // 投稿时间戳
	Pubtime  int                   `json:"pubtime"`  // 发布时间戳
	FavTime  int                   `json:"fav_time"` // 收藏时间戳
	BvID     string                `json:"bv_id"`    // 视频稿件bvid
	Bvid     string                `json:"bvid"`     // 视频稿件bvid
	Ugc      FavourUgc             `json:"ugc"`
}

type FavourList struct {
	Info    FavourFolderDetail `json:"info"`   // 收藏夹元数据
	Medias  []FavourMedia      `json:"medias"` // 收藏夹内容
	HasMore bool               `json:"has_more"`
}

type FavourID struct {
	ID   int    `json:"id"`    // 内容id，视频稿件：视频稿件avid，音频：音频auid，视频合集：视频合集id
	Type int    `json:"type"`  // 内容类型，2：视频稿件，12：音频，21：视频合集
	BvID string `json:"bv_id"` // 视频稿件bvid
	Bvid string `json:"bvid"`  // 视频稿件bvid
}

// SelfFavourMediaListResponse 是自身收藏夹内容列表。
type SelfFavourMediaListResponse struct {
	Count   int                   `json:"count"`
	List    []SelfFavourMediaItem `json:"list"`
	HasMore bool                  `json:"has_more"`
}

// SelfFavourMediaItem 是自身收藏夹中的一条内容。
type SelfFavourMediaItem struct {
	ID         int64       `json:"id"`
	Fid        int         `json:"fid"`
	Mid        int         `json:"mid"`
	Attr       int         `json:"attr"`
	AttrDesc   string      `json:"attr_desc"`
	Title      string      `json:"title"`
	Cover      string      `json:"cover"`
	Upper      FavourUpper `json:"upper"`
	CoverType  int         `json:"cover_type"`
	Intro      string      `json:"intro"`
	Ctime      int         `json:"ctime"`
	Mtime      int         `json:"mtime"`
	State      int         `json:"state"`
	FavState   int         `json:"fav_state"`
	MediaCount int         `json:"media_count"`
	ViewCount  int         `json:"view_count"`
	Vt         int         `json:"vt"`
	IsTop      bool        `json:"is_top"`
	RecentFav  any         `json:"recent_fav"`
	PlaySwitch int         `json:"play_switch"`
	Type       int         `json:"type"`
	Link       string      `json:"link"`
	Bvid       string      `json:"bvid"`
}

type SelfFavourList struct {
	ID                int64                       `json:"id"`
	Name              string                      `json:"name"`
	MediaListResponse SelfFavourMediaListResponse `json:"mediaListResponse"`
	URI               string                      `json:"uri"`
}
