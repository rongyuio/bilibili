package creativecenter

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 创作中心视频相关接口。文档见 creativecenter/videos.md。
// 用户空间的视频列表参见 video 子包的 GetUserVideos 等接口。
//
// 各接口的返回体独立定义，即使字段相同也不复用。

// GetCreativeCenterArchiveListParam 是获取稿件列表的请求参数。
type GetCreativeCenterArchiveListParam struct {
	Pn int `json:"pn" request:"query,default=1"`  // 页码。默认为 1
	Ps int `json:"ps" request:"query,default=10"` // 每页数量。默认为 10，留空为 100
}

// GetCreativeCenterArchiveListResult 是获取稿件列表的 data 对象。
type GetCreativeCenterArchiveListResult struct {
	Class     any                                 `json:"class"`      // 作用尚不明确。实测为 null
	ArcAudits []GetCreativeCenterArchiveListAudit `json:"arc_audits"` // 稿件列表
	Page      GetCreativeCenterArchiveListPage    `json:"page"`       // 分页信息
	PlayType  int                                 `json:"play_type"`  // 1。作用尚不明确
}

// GetCreativeCenterArchiveListPage 是稿件列表的分页信息。
type GetCreativeCenterArchiveListPage struct {
	Pn    int `json:"pn"`    // 页码
	Ps    int `json:"ps"`    // 每页数量
	Count int `json:"count"` // 稿件总数
}

// GetCreativeCenterArchiveListAudit 是稿件列表 arc_audits 数组中的对象。
type GetCreativeCenterArchiveListAudit struct {
	Archive        GetCreativeCenterArchiveListArchive `json:"Archive"`          // 稿件信息
	Videos         any                                 `json:"Videos"`           // 视频信息。实测为 null
	Stat           GetCreativeCenterArchiveListStat    `json:"stat"`             // 稿件统计信息
	StatePanel     int                                 `json:"state_panel"`      // 0。作用尚不明确
	ParentTname    string                              `json:"parent_tname"`     // 空
	Typename       string                              `json:"typename"`         // 空
	OpenAppeal     int                                 `json:"open_appeal"`      // 0。作用尚不明确
	Activity       any                                 `json:"activity"`         // 实测为 null
	SeasonAddState int                                 `json:"season_add_state"` // 是否可加入合集。0：可以。1：不可以
}

// GetCreativeCenterArchiveListArchive 是稿件列表中的稿件信息。
//
// 该接口返回的大多数字段为无实际意义的空值，字段说明按文档标注。
type GetCreativeCenterArchiveListArchive struct {
	Aid               int    `json:"aid"`                // av 号
	Bvid              string `json:"bvid"`               // bv 号
	Mid               int    `json:"mid"`                // 0
	Tid               int    `json:"tid"`                // 0
	TpInfo            any    `json:"tp_info"`            // 实测为 null
	Title             string `json:"title"`              // 稿件标题
	Author            string `json:"author"`             // 空
	Cover             string `json:"cover"`              // 封面图。位于 bfs，需要自行拼接 URL
	RejectReason      string `json:"reject_reason"`      // 空
	RejectReasonURL   string `json:"reject_reason_url"`  // 空
	Tag               string `json:"tag"`                // 空
	Duration          int    `json:"duration"`           // 0
	Copyright         int    `json:"copyright"`          // 0
	NoReprint         int    `json:"no_reprint"`         // 0
	UgcPay            int    `json:"ugcpay"`             // 0
	OrderID           int    `json:"order_id"`           // 0
	OrderName         string `json:"order_name"`         // 空
	AdorderID         int    `json:"adorder_id"`         // 0
	AdorderName       string `json:"adorder_name"`       // 空
	AdorderNo         string `json:"adorder_no"`         // 空
	OnlineTime        int    `json:"online_time"`        // 0
	NewAdorderInfo    any    `json:"new_adorder_info"`   // 实测为 null
	Desc              string `json:"desc"`               // 稿件简介
	MissionID         int    `json:"mission_id"`         // 0
	MissionName       string `json:"mission_name"`       // 空
	Attribute         int    `json:"attribute"`          // 0
	State             int    `json:"state"`              // 0
	StateDesc         string `json:"state_desc"`         // 空
	StatePanel        int    `json:"state_panel"`        // 0
	Source            string `json:"source"`             // 空
	DescFormatID      int    `json:"desc_format_id"`     // 0
	Attrs             any    `json:"attrs"`              // 实测为 null
	Porder            any    `json:"porder"`             // 实测为 null
	Dynamic           string `json:"dynamic"`            // 空
	PoiObject         any    `json:"poi_object"`         // 实测为 null
	Dtime             int    `json:"dtime"`              // 0
	Ptime             int    `json:"ptime"`              // 0
	Ctime             int    `json:"ctime"`              // 0
	UgcPayInfo        any    `json:"ugcpay_info"`        // 实测为 null
	Staffs            any    `json:"staffs"`             // 实测为 null
	Vote              any    `json:"vote"`               // 实测为 null
	Activity          any    `json:"activity"`           // 实测为 null
	Interactive       int    `json:"interactive"`        // 0
	Hl                any    `json:"hl"`                 // 实测为 null
	NoBackground      int    `json:"no_background"`      // 0
	DynamicVideo      int    `json:"dynamic_video"`      // 0
	NoPublic          int    `json:"no_public"`          // 0
	Is360             int    `json:"is_360"`             // 0
	IsDolby           int    `json:"is_dolby"`           // 0
	LosslessMusic     int    `json:"lossless_music"`     // 0
	BSEditor          int    `json:"bs_editor"`          // 0
	UpFrom            int    `json:"up_from"`            // 0
	DescV2            any    `json:"desc_v2"`            // 实测为 null
	DynamicV2         any    `json:"dynamic_v2"`         // 实测为 null
	TopicID           int    `json:"topic_id"`           // 0
	TopicName         string `json:"topic_name"`         // 空
	TopicStat         int    `json:"topic_stat"`         // 0
	Premiere          int    `json:"premiere"`           // 0
	IsUgcPayV2        int    `json:"is_ugcpay_v2"`       // 0
	Recreate          any    `json:"recreate"`           // 实测为 null
	PoliticalMedia    int    `json:"political_media"`    // 0。文档表格未列出，响应示例中存在
	PoliticalEditable int    `json:"political_editable"` // 0。文档表格未列出，响应示例中存在
	ChargingPay       int    `json:"charging_pay"`       // 0
	NeutralMark       string `json:"neutral_mark"`       // 空
	Preview           any    `json:"preview"`            // 实测为 null
	UpowerLevel       any    `json:"upower_level"`       // 实测为 null
	Cover43           string `json:"cover43"`            // 空
	HadPassed         bool   `json:"had_passed"`         // false
	IsStaff           int    `json:"is_staff"`           // 0
	IsPugv            int    `json:"is_pugv"`            // 0
	UpowerMode        int    `json:"upower_mode"`        // 0
	UpowerUnitPrice   int    `json:"upower_unit_price"`  // 0
}

// GetCreativeCenterArchiveListStat 是稿件列表中的稿件统计信息。
type GetCreativeCenterArchiveListStat struct {
	Aid      int `json:"aid"`      // av 号
	View     int `json:"view"`     // 播放量
	Danmaku  int `json:"danmaku"`  // 弹幕数
	Reply    int `json:"reply"`    // 评论数
	Favorite int `json:"favorite"` // 收藏数
	Coin     int `json:"coin"`     // 硬币数
	Share    int `json:"share"`    // 分享数
	NowRank  int `json:"now_rank"` // 当前排名
	HisRank  int `json:"his_rank"` // 历史最高排名
	Like     int `json:"like"`     // 点赞数
	Dislike  int `json:"dislike"`  // 点踩数。恒为 0
	Vt       int `json:"vt"`       // 0。作用尚不明确
	Vv       int `json:"vv"`       // 播放量
}

// GetCreativeCenterArchiveList 获取稿件列表
//
// 该接口返回内容大多为无实际意义的空值，具体内容建议通过常规接口获取。
func (c *Client) GetCreativeCenterArchiveList(ctx context.Context, param GetCreativeCenterArchiveListParam) (*GetCreativeCenterArchiveListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x2/creative/web/archives/sp"
	)
	return bilibili.Execute[*GetCreativeCenterArchiveListResult](ctx, c.Client, method, url, param)
}

// GetCreativeCenterVideoBasicInfoParam 是获取视频基础信息的请求参数。
type GetCreativeCenterVideoBasicInfoParam struct {
	Aid int `json:"aid" request:"query"` // av 号
}

// GetCreativeCenterVideoBasicInfoResult 是视频基础信息的 data 对象。
type GetCreativeCenterVideoBasicInfoResult struct {
	Archive GetCreativeCenterVideoBasicInfoArchive `json:"archive"` // 稿件信息
	Videos  []GetCreativeCenterVideoBasicInfoVideo `json:"videos"`  // 视频信息
}

// GetCreativeCenterVideoBasicInfoArchive 是视频基础信息中的稿件信息。
type GetCreativeCenterVideoBasicInfoArchive struct {
	Aid   int    `json:"aid"`   // av 号
	Bvid  string `json:"bvid"`  // bvid
	Title string `json:"title"` // 标题
}

// GetCreativeCenterVideoBasicInfoVideo 是视频基础信息中的分P信息。
type GetCreativeCenterVideoBasicInfoVideo struct {
	Cid      int    `json:"cid"`      // 分P cid
	Index    int    `json:"index"`    // 分P 序号
	Title    string `json:"title"`    // 分P 标题
	Duration int    `json:"duration"` // 0
}

// GetCreativeCenterVideoBasicInfo 视频基础信息
func (c *Client) GetCreativeCenterVideoBasicInfo(ctx context.Context, param GetCreativeCenterVideoBasicInfoParam) (*GetCreativeCenterVideoBasicInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/archive/videos"
	)
	return bilibili.Execute[*GetCreativeCenterVideoBasicInfoResult](ctx, c.Client, method, url, param)
}
