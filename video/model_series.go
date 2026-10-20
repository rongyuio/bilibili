package video

// 用户空间视频列表（系列 / 合集）与视频列表管理接口的响应模型。
//
// 两个列表接口的形状相近，但按接口隔离的原则各自定义，互不复用。

// GetUserSpaceSeriesListResult 是「只获取系列视频」的 data 对象。
type GetUserSpaceSeriesListResult struct {
	ItemsLists GetUserSpaceSeriesListItemsLists `json:"items_lists"` // 内容列表。套了个娃
}

// GetUserSpaceSeriesListItemsLists 是「只获取系列视频」的 data.items_lists 对象。
type GetUserSpaceSeriesListItemsLists struct {
	Page        GetUserSpaceSeriesListPage         `json:"page"`         // 分页信息
	SeasonsList []GetUserSpaceSeriesListSeasonItem `json:"seasons_list"` // 合集列表。只获取系列视频时为空
	SeriesList  []GetUserSpaceSeriesListSeriesItem `json:"series_list"`  // 系列列表
}

// GetUserSpaceSeriesListPage 是「只获取系列视频」的分页信息。
type GetUserSpaceSeriesListPage struct {
	PageNum  int `json:"page_num"`  // 分页页码
	PageSize int `json:"page_size"` // 单页个数
	Total    int `json:"total"`     // 合集内视频数量
}

// GetUserSpaceSeriesListSeasonItem 是「只获取系列视频」seasons_list 数组中的对象。
type GetUserSpaceSeriesListSeasonItem struct {
	Archives   []GetUserSpaceSeriesListArchive  `json:"archives"`    // 合集视频列表
	Meta       GetUserSpaceSeriesListSeasonMeta `json:"meta"`        // 合集元数据
	RecentAids []int64                          `json:"recent_aids"` // 合集视频 aid 列表
}

// GetUserSpaceSeriesListSeriesItem 是「只获取系列视频」series_list 数组中的对象。
type GetUserSpaceSeriesListSeriesItem struct {
	Archives   []GetUserSpaceSeriesListArchive  `json:"archives"`    // 系列视频列表
	Meta       GetUserSpaceSeriesListSeriesMeta `json:"meta"`        // 系列元数据
	RecentAids []int64                          `json:"recent_aids"` // 系列视频 aid 列表
}

// GetUserSpaceSeriesListArchive 是「只获取系列视频」archives 数组中的对象。
type GetUserSpaceSeriesListArchive struct {
	Aid              int                               `json:"aid"`               // 稿件avid
	Bvid             string                            `json:"bvid"`              // 稿件bvid
	Ctime            int                               `json:"ctime"`             // 创建时间。Unix 时间戳
	Duration         int                               `json:"duration"`          // 视频时长。单位为秒
	EnableVt         any                               `json:"enable_vt"`         // int or bool
	InteractiveVideo bool                              `json:"interactive_video"` // false
	Pic              string                            `json:"pic"`               // 封面 URL
	PlaybackPosition int                               `json:"playback_position"` // 会随着播放时间增长，播放完成后为 -1 。单位未知
	PubDate          int                               `json:"pubdate"`           // 发布日期。Unix 时间戳
	Stat             GetUserSpaceSeriesListArchiveStat `json:"stat"`              // 稿件信息
	State            int                               `json:"state"`             // 0
	Title            string                            `json:"title"`             // 稿件标题
	UgcPay           int                               `json:"ugc_pay"`           // 0
	VtDisplay        string                            `json:"vt_display"`
	UpMid            int                               `json:"upMid"` // UP主mid。实测为 39665558
}

// GetUserSpaceSeriesListArchiveStat 是「只获取系列视频」archives[].stat 对象。
type GetUserSpaceSeriesListArchiveStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 0
}

// GetUserSpaceSeriesListSeriesMeta 是「只获取系列视频」series_list[].meta 对象。
type GetUserSpaceSeriesListSeriesMeta struct {
	Category     int      `json:"category"`       // 1
	Cover        string   `json:"cover"`          // 系列封面 URL
	Creator      string   `json:"creator"`        // auto
	Ctime        int64    `json:"ctime"`          // 创建时间。Unix 时间戳
	Description  string   `json:"description"`    // 系列描述
	Keywords     []string `json:"keywords"`       // 系列关键词列表
	LastUpdateTs int64    `json:"last_update_ts"` // 最近更新时间。Unix 时间戳
	Mid          int      `json:"mid"`            // UP 主 ID
	Mtime        int64    `json:"mtime"`          // 修改时间。Unix 时间戳
	Name         string   `json:"name"`           // 系列标题
	RawKeywords  string   `json:"raw_keywords"`   // 原始系列关键词
	SeriesID     int      `json:"series_id"`      // 系列 ID
	State        int      `json:"state"`          // 2
	Total        int      `json:"total"`          // 系列视频数量
}

// GetUserSpaceSeriesListSeasonMeta 是「只获取系列视频」seasons_list[].meta 对象。
type GetUserSpaceSeriesListSeasonMeta struct {
	Category    int    `json:"category"`    // 0
	Cover       string `json:"cover"`       // 合集封面 URL
	Description string `json:"description"` // 合集描述
	Mid         int    `json:"mid"`         // UP 主 ID
	Name        string `json:"name"`        // 合集标题
	Ptime       int64  `json:"ptime"`       // 发布时间。Unix 时间戳
	SeasonID    int    `json:"season_id"`   // 合集 ID
	Total       int    `json:"total"`       // 合集内视频数量
}

// GetUserSpaceSeasonsSeriesListResult 是「获取系列和合集视频」的 data 对象。
type GetUserSpaceSeasonsSeriesListResult struct {
	ItemsLists GetUserSpaceSeasonsSeriesListItemsLists `json:"items_lists"` // 内容列表。套了个娃
}

// GetUserSpaceSeasonsSeriesListItemsLists 是「获取系列和合集视频」的 data.items_lists 对象。
type GetUserSpaceSeasonsSeriesListItemsLists struct {
	Page        GetUserSpaceSeasonsSeriesListPage         `json:"page"`         // 分页信息
	SeasonsList []GetUserSpaceSeasonsSeriesListSeasonItem `json:"seasons_list"` // 合集列表
	SeriesList  []GetUserSpaceSeasonsSeriesListSeriesItem `json:"series_list"`  // 系列列表
}

// GetUserSpaceSeasonsSeriesListPage 是「获取系列和合集视频」的分页信息。
type GetUserSpaceSeasonsSeriesListPage struct {
	PageNum  int `json:"page_num"`  // 分页页码
	PageSize int `json:"page_size"` // 单页个数
	Total    int `json:"total"`     // 合集内视频数量
}

// GetUserSpaceSeasonsSeriesListSeasonItem 是「获取系列和合集视频」seasons_list 数组中的对象。
type GetUserSpaceSeasonsSeriesListSeasonItem struct {
	Archives   []GetUserSpaceSeasonsSeriesListArchive  `json:"archives"`    // 合集视频列表
	Meta       GetUserSpaceSeasonsSeriesListSeasonMeta `json:"meta"`        // 合集元数据
	RecentAids []int64                                 `json:"recent_aids"` // 合集视频 aid 列表
}

// GetUserSpaceSeasonsSeriesListSeriesItem 是「获取系列和合集视频」series_list 数组中的对象。
type GetUserSpaceSeasonsSeriesListSeriesItem struct {
	Archives   []GetUserSpaceSeasonsSeriesListArchive  `json:"archives"`    // 系列视频列表
	Meta       GetUserSpaceSeasonsSeriesListSeriesMeta `json:"meta"`        // 系列元数据
	RecentAids []int64                                 `json:"recent_aids"` // 系列视频 aid 列表
}

// GetUserSpaceSeasonsSeriesListArchive 是「获取系列和合集视频」archives 数组中的对象。
type GetUserSpaceSeasonsSeriesListArchive struct {
	Aid              int                                      `json:"aid"`               // 稿件avid
	Bvid             string                                   `json:"bvid"`              // 稿件bvid
	Ctime            int                                      `json:"ctime"`             // 创建时间。Unix 时间戳
	Duration         int                                      `json:"duration"`          // 视频时长。单位为秒
	EnableVt         any                                      `json:"enable_vt"`         // int or bool
	InteractiveVideo bool                                     `json:"interactive_video"` // false
	Pic              string                                   `json:"pic"`               // 封面 URL
	PlaybackPosition int                                      `json:"playback_position"` // 会随着播放时间增长，播放完成后为 -1 。单位未知
	PubDate          int                                      `json:"pubdate"`           // 发布日期。Unix 时间戳
	Stat             GetUserSpaceSeasonsSeriesListArchiveStat `json:"stat"`              // 稿件信息
	State            int                                      `json:"state"`             // 0
	Title            string                                   `json:"title"`             // 稿件标题
	UgcPay           int                                      `json:"ugc_pay"`           // 0
	VtDisplay        string                                   `json:"vt_display"`
	UpMid            int                                      `json:"upMid"` // UP主mid。实测为 39665558
}

// GetUserSpaceSeasonsSeriesListArchiveStat 是「获取系列和合集视频」archives[].stat 对象。
type GetUserSpaceSeasonsSeriesListArchiveStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 0
}

// GetUserSpaceSeasonsSeriesListSeriesMeta 是「获取系列和合集视频」series_list[].meta 对象。
type GetUserSpaceSeasonsSeriesListSeriesMeta struct {
	Category     int      `json:"category"`       // 1
	Cover        string   `json:"cover"`          // 系列封面 URL
	Creator      string   `json:"creator"`        // auto
	Ctime        int64    `json:"ctime"`          // 创建时间。Unix 时间戳
	Description  string   `json:"description"`    // 系列描述
	Keywords     []string `json:"keywords"`       // 系列关键词列表
	LastUpdateTs int64    `json:"last_update_ts"` // 最近更新时间。Unix 时间戳
	Mid          int      `json:"mid"`            // UP 主 ID
	Mtime        int64    `json:"mtime"`          // 修改时间。Unix 时间戳
	Name         string   `json:"name"`           // 系列标题
	RawKeywords  string   `json:"raw_keywords"`   // 原始系列关键词
	SeriesID     int      `json:"series_id"`      // 系列 ID
	State        int      `json:"state"`          // 2
	Total        int      `json:"total"`          // 系列视频数量
}

// GetUserSpaceSeasonsSeriesListSeasonMeta 是「获取系列和合集视频」seasons_list[].meta 对象。
type GetUserSpaceSeasonsSeriesListSeasonMeta struct {
	Category    int    `json:"category"`    // 0
	Cover       string `json:"cover"`       // 合集封面 URL
	Description string `json:"description"` // 合集描述
	Mid         int    `json:"mid"`         // UP 主 ID
	Name        string `json:"name"`        // 合集标题
	Ptime       int64  `json:"ptime"`       // 发布时间。Unix 时间戳
	SeasonID    int    `json:"season_id"`   // 合集 ID
	Total       int    `json:"total"`       // 合集内视频数量
}

// GetSeriesInfoResult 是「查询指定系列」的 data 对象。
type GetSeriesInfoResult struct {
	Meta       GetSeriesInfoMeta `json:"meta"`        // 系列信息
	RecentAids []int64           `json:"recent_aids"` // 系列 aid 列表
}

// GetSeriesInfoMeta 是「查询指定系列」的 meta 对象。
type GetSeriesInfoMeta struct {
	Category     int      `json:"category"`       // 1
	Cover        string   `json:"cover"`          // 系列封面 URL
	Creator      string   `json:"creator"`        // auto
	Ctime        int64    `json:"ctime"`          // 创建时间。Unix 时间戳
	Description  string   `json:"description"`    // 系列描述
	Keywords     []string `json:"keywords"`       // 系列关键词列表
	LastUpdateTs int64    `json:"last_update_ts"` // 最近更新时间。Unix 时间戳
	Mid          int      `json:"mid"`            // UP 主 ID
	Mtime        int64    `json:"mtime"`          // 修改时间。Unix 时间戳
	Name         string   `json:"name"`           // 系列标题
	RawKeywords  string   `json:"raw_keywords"`   // 原始系列关键词
	SeriesID     int      `json:"series_id"`      // 系列 ID
	State        int      `json:"state"`          // 2
	Total        int      `json:"total"`          // 系列视频数量
}

// CreateVideoSeriesResult 是「创建视频列表」的 data 对象。
type CreateVideoSeriesResult struct {
	SeriesID int64 `json:"series_id"` // 视频列表 ID
}
