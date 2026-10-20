package audio

// 音频歌单与音频收藏夹相关响应模型。
//
// 各接口的 data 层与嵌套对象均独立定义，即使字段相同也不复用。

// GetAudioCollectionListResult 是查询自己创建的歌单的响应 data 层。
type GetAudioCollectionListResult struct {
	CurPage   int                          `json:"curPage"`   // 当前页码
	PageCount int                          `json:"pageCount"` // 总计页数
	TotalSize int                          `json:"totalSize"` // 总计收藏夹数
	PageSize  int                          `json:"pageSize"`  // 当前页面项数
	Data      []GetAudioCollectionListItem `json:"data"`      // 歌单列表
}

// GetAudioCollectionListItem 是歌单列表中的一项，即一个音频收藏夹（歌单）。
type GetAudioCollectionListItem struct {
	ID        int                                 `json:"id"`        // 音频收藏夹 mlid
	UID       int                                 `json:"uid"`       // 创建用户 mid
	Uname     string                              `json:"uname"`     // 创建用户昵称
	Title     string                              `json:"title"`     // 歌单标题
	Type      int                                 `json:"type"`      // 收藏夹属性。0：普通收藏夹，1：默认收藏夹
	Published int                                 `json:"published"` // 是否公开。0：不公开，1：公开
	Cover     string                              `json:"cover"`     // 歌单封面图片 url
	Ctime     int                                 `json:"ctime"`     // 歌单创建时间。时间戳
	Song      int                                 `json:"song"`      // 歌单中的音乐数量
	Desc      string                              `json:"desc"`      // 歌单备注信息
	Sids      []int                               `json:"sids"`      // 歌单中的音乐 auid 列表，按歌单顺序排列
	MenuID    int                                 `json:"menuId"`    // 音频收藏夹对应的歌单 amid
	Statistic GetAudioCollectionListItemStatistic `json:"statistic"` // 歌单状态数信息
}

// GetAudioCollectionListItemStatistic 是歌单列表项的状态数信息。
type GetAudioCollectionListItemStatistic struct {
	Sid     int  `json:"sid"`     // 音频收藏夹对应的歌单 amid
	Play    int  `json:"play"`    // 播放数
	Collect int  `json:"collect"` // 收藏数
	Comment *int `json:"comment"` // 评论数。可能为 null
	Share   int  `json:"share"`   // 分享数
}

// GetAudioCollectionInfoResult 是查询音频收藏夹（默认歌单）信息的响应 data 层。
type GetAudioCollectionInfoResult struct {
	ID        int                             `json:"id"`        // 音频收藏夹 mlid
	UID       int                             `json:"uid"`       // 创建用户 mid
	Uname     string                          `json:"uname"`     // 创建用户昵称
	Title     string                          `json:"title"`     // 歌单标题
	Type      int                             `json:"type"`      // 收藏夹属性。0：普通收藏夹，1：默认收藏夹
	Published int                             `json:"published"` // 是否公开。0：不公开，1：公开
	Cover     string                          `json:"cover"`     // 歌单封面图片 url
	Ctime     int                             `json:"ctime"`     // 歌单创建时间。时间戳
	Song      int                             `json:"song"`      // 歌单中的音乐数量
	Desc      string                          `json:"desc"`      // 歌单备注信息
	Sids      []int                           `json:"sids"`      // 歌单中的音乐 auid 列表，按歌单顺序排列
	MenuID    int                             `json:"menuId"`    // 音频收藏夹对应的歌单 amid
	Statistic GetAudioCollectionInfoStatistic `json:"statistic"` // 歌单状态数信息
}

// GetAudioCollectionInfoStatistic 是音频收藏夹（默认歌单）的状态数信息。
type GetAudioCollectionInfoStatistic struct {
	Sid     int  `json:"sid"`     // 音频收藏夹对应的歌单 amid
	Play    int  `json:"play"`    // 播放数
	Collect int  `json:"collect"` // 收藏数
	Comment *int `json:"comment"` // 评论数。可能为 null
	Share   int  `json:"share"`   // 分享数
}

// GetAudioHotMenuListResult 是查询热门歌单的响应 data 层。
type GetAudioHotMenuListResult struct {
	CurPage   int                   `json:"curPage"`   // 当前页码
	PageCount int                   `json:"pageCount"` // 总计页数
	TotalSize int                   `json:"totalSize"` // 总计收藏夹数
	PageSize  int                   `json:"pageSize"`  // 当前页面项数
	Data      []GetAudioHotMenuItem `json:"data"`      // 歌单列表
}

// GetAudioHotMenuItem 是热门歌单列表中的一个歌单。
type GetAudioHotMenuItem struct {
	MenuID       int                          `json:"menuId"`       // 音频收藏夹对应的歌单 amid
	UID          int                          `json:"uid"`          // 创建用户 mid
	Uname        string                       `json:"uname"`        // 创建用户昵称
	Title        string                       `json:"title"`        // 歌单标题
	Cover        string                       `json:"cover"`        // 歌单封面图片 url
	Intro        string                       `json:"intro"`        // 歌单介绍
	Type         int                          `json:"type"`         // 歌单属性。1：普通歌单，2：置顶歌单，5：PGC 歌单
	Off          int                          `json:"off"`          // 歌单是否公开。0：公开，1：私密
	Ctime        int                          `json:"ctime"`        // 歌单创建时间。时间戳
	Curtime      int                          `json:"curtime"`      // 当前时间。时间戳
	Statistic    GetAudioHotMenuItemStatistic `json:"statistic"`    // 歌单状态数信息
	Snum         int                          `json:"snum"`         // 歌单包含歌曲个数
	IsDefault    int                          `json:"isDefault"`    // 是否为默认歌单。实测为 0
	CollectionID int                          `json:"collectionId"` // 收藏夹 id。实测为 0
}

// GetAudioHotMenuItemStatistic 是热门歌单的状态数信息。
type GetAudioHotMenuItemStatistic struct {
	Sid     int `json:"sid"`     // 音频收藏夹对应的歌单 amid
	Play    int `json:"play"`    // 播放数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

// GetAudioRankMenuListResult 是查询热门榜的响应 data 层。
type GetAudioRankMenuListResult struct {
	CurPage   int                    `json:"curPage"`   // 当前页码
	PageCount int                    `json:"pageCount"` // 总计页数
	TotalSize int                    `json:"totalSize"` // 总计收藏夹数
	PageSize  int                    `json:"pageSize"`  // 当前页面项数
	Data      []GetAudioRankMenuItem `json:"data"`      // 歌单列表
}

// GetAudioRankMenuItem 是热门榜中的一个歌单。
type GetAudioRankMenuItem struct {
	MenuID       int                           `json:"menuId"`       // 音频收藏夹对应的歌单 amid
	UID          int                           `json:"uid"`          // 创建用户 mid
	Uname        string                        `json:"uname"`        // 创建用户昵称
	Title        string                        `json:"title"`        // 歌单标题
	Cover        string                        `json:"cover"`        // 歌单封面图片 url
	Intro        string                        `json:"intro"`        // 歌单介绍
	Type         int                           `json:"type"`         // 歌单属性。1：普通歌单，2：置顶歌单，5：PGC 歌单
	Off          int                           `json:"off"`          // 歌单是否公开。0：公开，1：私密
	Ctime        int                           `json:"ctime"`        // 歌单创建时间。时间戳
	Curtime      int                           `json:"curtime"`      // 当前时间。时间戳
	Statistic    GetAudioRankMenuItemStatistic `json:"statistic"`    // 歌单状态数信息
	Snum         int                           `json:"snum"`         // 歌单包含歌曲个数
	Audios       []GetAudioRankMenuItemAudio   `json:"audios"`       // 歌单中的音乐信息（部分）
	IsDefault    int                           `json:"isDefault"`    // 是否为默认歌单。实测为 0
	CollectionID int                           `json:"collectionId"` // 收藏夹 id。实测为 0
}

// GetAudioRankMenuItemStatistic 是热门榜歌单的状态数信息。
type GetAudioRankMenuItemStatistic struct {
	Sid     int `json:"sid"`     // 音频收藏夹对应的歌单 amid
	Play    int `json:"play"`    // 播放数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

// GetAudioRankMenuItemAudio 是热门榜歌单中一首音乐的简要信息。
type GetAudioRankMenuItemAudio struct {
	ID       int    `json:"id"`       // 音频 id
	Title    string `json:"title"`    // 音频标题
	Duration int    `json:"duration"` // 音频时长。单位：秒
}
