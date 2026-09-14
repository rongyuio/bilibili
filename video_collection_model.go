package bilibili

// 视频合集相关响应模型。

type CollectionVideoStat struct {
	View int `json:"view"` // 稿件播放量
	Vt   int `json:"vt"`   // 0
}

type CollectionVideo struct {
	Aid              int                 `json:"aid"`               // 稿件avid
	Bvid             string              `json:"bvid"`              // 稿件bvid
	Ctime            int                 `json:"ctime"`             // 创建时间。Unix 时间戳
	Duration         int                 `json:"duration"`          // 视频时长。单位为秒
	EnableVt         any                 `json:"enable_vt"`         // int or bool
	InteractiveVideo bool                `json:"interactive_video"` // false
	Pic              string              `json:"pic"`               // 封面 URL
	PlaybackPosition int                 `json:"playback_position"` // 会随着播放时间增长，播放完成后为 -1 。单位未知
	PubDate          int                 `json:"pubdate"`           // 发布日期。Unix 时间戳
	Stat             CollectionVideoStat `json:"stat"`              // 稿件信息
	State            int                 `json:"state"`             // 0
	Title            string              `json:"title"`             // 稿件标题
	UgcPay           int                 `json:"ugc_pay"`           // 0
	VtDisplay        string              `json:"vt_display"`
}

type CollectionMeta struct {
	Category    int    `json:"category"`    // 0
	Covr        string `json:"covr"`        // 合集封面 URL
	Description string `json:"description"` // 合集描述
	Mid         int    `json:"mid"`         // UP 主 ID
	Name        string `json:"name"`        // 合集标题
	Ptime       int    `json:"ptime"`       // 发布时间。Unix 时间戳
	SeasonID    int    `json:"season_id"`   // 合集 ID
	Total       int    `json:"total"`       // 合集内视频数量
}

// CollectionPage 视频合集分页信息。
//
// 字段名为 page_num/page_size/total，与空间投稿、分区列表的分页字段不同，
// 因此保留独立类型，不做跨接口合并。
type CollectionPage struct {
	PageNum  int `json:"page_num"`  // 分页页码
	PageSize int `json:"page_size"` // 单页个数
	Total    int `json:"total"`     // 合集内视频数量
}

type VideoCollectionInfo struct {
	Aids     []int             `json:"aids"`           // 稿件avid。对应下方数组中内容 aid
	Archives []CollectionVideo `json:"archives"`       // 合集中的视频
	Meta     CollectionMeta    `json:"meta,omitempty"` // 合集元数据
	Page     CollectionPage    `json:"page"`           // 分页信息
}

type VideoCollectionByKeywordsInfo struct {
	Archives []CollectionVideo `json:"archives"` // 视频列表
	Page     CollectionPage    `json:"page"`     // 页码信息
}
