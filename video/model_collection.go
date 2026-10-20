package video

// 视频合集、视频列表与关键词检索接口的响应模型。
//
// 三个接口的 data 形状相近，但按接口隔离的原则各自定义，互不复用。

// GetVideoCollectionInfoResult 是「获取视频合集信息」的 data 对象。
type GetVideoCollectionInfoResult struct {
	Aids     []int                           `json:"aids"`     // 稿件avid。对应下方数组中内容 aid
	Archives []GetVideoCollectionInfoArchive `json:"archives"` // 合集中的视频
	Meta     GetVideoCollectionInfoMeta      `json:"meta"`     // 合集元数据
	Page     GetVideoCollectionInfoPage      `json:"page"`     // 分页信息
}

// GetVideoCollectionInfoArchive 是「获取视频合集信息」archives 数组中的对象。
type GetVideoCollectionInfoArchive struct {
	Aid              int                               `json:"aid"`               // 稿件avid
	Bvid             string                            `json:"bvid"`              // 稿件bvid
	Ctime            int                               `json:"ctime"`             // 创建时间。Unix 时间戳
	Duration         int                               `json:"duration"`          // 视频时长。单位为秒
	EnableVt         any                               `json:"enable_vt"`         // int or bool
	InteractiveVideo bool                              `json:"interactive_video"` // 是否是互动视频
	Pic              string                            `json:"pic"`               // 封面 URL
	PlaybackPosition int                               `json:"playback_position"` // 播放进度，播放完成后为 -1
	PubDate          int                               `json:"pubdate"`           // 发布日期。Unix 时间戳
	Stat             GetVideoCollectionInfoArchiveStat `json:"stat"`              // 稿件信息
	State            int                               `json:"state"`             // 0
	Title            string                            `json:"title"`             // 稿件标题
	UgcPay           int                               `json:"ugc_pay"`           // UGC 付费? 0: 否
	VtDisplay        string                            `json:"vt_display"`        // 空
	UpMid            int                               `json:"upMid"`             // UP主mid
}

// GetVideoCollectionInfoArchiveStat 是「获取视频合集信息」archives[].stat 对象。
type GetVideoCollectionInfoArchiveStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 0
}

// GetVideoCollectionInfoMeta 是「获取视频合集信息」meta 对象。
type GetVideoCollectionInfoMeta struct {
	Category    int    `json:"category"`    // 0
	Covr        string `json:"covr"`        // 合集封面 URL
	Description string `json:"description"` // 合集描述
	Mid         int    `json:"mid"`         // UP 主 ID
	Name        string `json:"name"`        // 合集标题
	Ptime       int    `json:"ptime"`       // 发布时间。Unix 时间戳
	SeasonID    int    `json:"season_id"`   // 合集 ID
	Total       int    `json:"total"`       // 合集内视频数量
}

// GetVideoCollectionInfoPage 是「获取视频合集信息」page 对象。
type GetVideoCollectionInfoPage struct {
	PageNum  int `json:"page_num"`  // 分页页码
	PageSize int `json:"page_size"` // 单页个数
	Total    int `json:"total"`     // 合集内视频数量
}

// GetVideoByKeywordsResult 是「根据关键词查找视频」的 data 对象。
type GetVideoByKeywordsResult struct {
	Archives []GetVideoByKeywordsArchive `json:"archives"` // 视频列表
	Page     GetVideoByKeywordsPage      `json:"page"`     // 页码信息
}

// GetVideoByKeywordsArchive 是「根据关键词查找视频」archives 数组中的对象。
type GetVideoByKeywordsArchive struct {
	Aid              int                           `json:"aid"`               // 稿件avid
	Bvid             string                        `json:"bvid"`              // 稿件bvid
	Ctime            int                           `json:"ctime"`             // 创建时间。Unix 时间戳
	Duration         int                           `json:"duration"`          // 视频时长。单位为秒
	EnableVt         any                           `json:"enable_vt"`         // int or bool
	InteractiveVideo bool                          `json:"interactive_video"` // 是否是互动视频
	Pic              string                        `json:"pic"`               // 封面 URL
	PlaybackPosition int                           `json:"playback_position"` // 播放进度，播放完成后为 -1
	PubDate          int                           `json:"pubdate"`           // 发布日期。Unix 时间戳
	Stat             GetVideoByKeywordsArchiveStat `json:"stat"`              // 稿件信息
	State            int                           `json:"state"`             // 0
	Title            string                        `json:"title"`             // 稿件标题
	UgcPay           int                           `json:"ugc_pay"`           // UGC 付费? 0: 否
	VtDisplay        string                        `json:"vt_display"`        // 空
	UpMid            int                           `json:"upMid"`             // UP主mid
}

// GetVideoByKeywordsArchiveStat 是「根据关键词查找视频」archives[].stat 对象。
type GetVideoByKeywordsArchiveStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 0
}

// GetVideoByKeywordsPage 是「根据关键词查找视频」page 对象。
//
// 文档字段表称与合集 page 相同，但本节 JSON 示例用的是 num/size/total，
// 与合集接口的 page_num/page_size 不同，这里按示例取键名。
type GetVideoByKeywordsPage struct {
	Num   int `json:"num"`   // 当前页码
	Size  int `json:"size"`  // 每页数量
	Total int `json:"total"` // 视频总数
}

// GetVideoSeriesInfoResult 是「获取指定系列视频」的 data 对象。
type GetVideoSeriesInfoResult struct {
	Aids     []int                       `json:"aids"`     // 视频 aid 列表
	Archives []GetVideoSeriesInfoArchive `json:"archives"` // 视频信息列表
	Page     GetVideoSeriesInfoPage      `json:"page"`     // 页码信息
}

// GetVideoSeriesInfoArchive 是「获取指定系列视频」archives 数组中的对象。
type GetVideoSeriesInfoArchive struct {
	Aid              int                           `json:"aid"`               // 稿件avid
	Bvid             string                        `json:"bvid"`              // 稿件bvid
	Ctime            int                           `json:"ctime"`             // 创建时间。Unix 时间戳
	Duration         int                           `json:"duration"`          // 视频时长。单位为秒
	EnableVt         any                           `json:"enable_vt"`         // int or bool
	InteractiveVideo bool                          `json:"interactive_video"` // 是否是互动视频
	Pic              string                        `json:"pic"`               // 封面 URL
	PlaybackPosition int                           `json:"playback_position"` // 播放进度，播放完成后为 -1
	PubDate          int                           `json:"pubdate"`           // 发布日期。Unix 时间戳
	Stat             GetVideoSeriesInfoArchiveStat `json:"stat"`              // 稿件信息
	State            int                           `json:"state"`             // 0
	Title            string                        `json:"title"`             // 稿件标题
	UgcPay           int                           `json:"ugc_pay"`           // UGC 付费? 0: 否
	VtDisplay        string                        `json:"vt_display"`        // 空
	UpMid            int                           `json:"upMid"`             // UP主mid
}

// GetVideoSeriesInfoArchiveStat 是「获取指定系列视频」archives[].stat 对象。
type GetVideoSeriesInfoArchiveStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 0
}

// GetVideoSeriesInfoPage 是「获取指定系列视频」page 对象。
//
// 与合集接口不同，这里的键名是 num/size/total。
type GetVideoSeriesInfoPage struct {
	Num   int `json:"num"`   // 当前页码
	Size  int `json:"size"`  // 每页数量
	Total int `json:"total"` // 视频总数
}
