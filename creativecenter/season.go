package creativecenter

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 创作中心合集管理接口。文档见 creativecenter/season.md。
// 用户空间的合集及视频列表参见 video 子包的 GetVideoCollectionInfo 等接口。
//
// 各接口的返回体独立定义，即使字段相同也不复用。

// GetCreativeCenterSeasonListParam 是获取合集列表的请求参数。
type GetCreativeCenterSeasonListParam struct {
	Pn    int    `json:"pn" request:"query,default=1"`              // 页码。默认为 1
	Ps    int    `json:"ps" request:"query,default=30"`             // 每页数量。默认为 30
	Order string `json:"order,omitempty" request:"query,omitempty"` // 排序方式。创建时间: ctime。修改时间: mtime
	Sort  string `json:"sort,omitempty" request:"query,omitempty"`  // 排序方向。创建时间: asc。修改时间: desc
	Draft int    `json:"draft,omitempty" request:"query,omitempty"` // 1。作用尚不明确
}

// GetCreativeCenterSeasonListResult 是获取合集列表的 data 对象。
type GetCreativeCenterSeasonListResult struct {
	Seasons  []GetCreativeCenterSeasonListItem `json:"seasons"`   // 合集列表
	Tip      GetCreativeCenterSeasonListTip    `json:"tip"`       // 提示信息。title 与 url 均为空
	Total    int                               `json:"total"`     // 合集总数
	PlayType int                               `json:"play_type"` // 1。作用尚不明确
	SortMode string                            `json:"sort_mode"` // 排序方式。实测为 default
}

// GetCreativeCenterSeasonListTip 是合集列表的 tip 对象。
type GetCreativeCenterSeasonListTip struct {
	Title string `json:"title"` // 提示标题。实测为空
	URL   string `json:"url"`   // 提示链接。实测为空
}

// GetCreativeCenterSeasonListItem 是合集列表 seasons 数组中的对象。
type GetCreativeCenterSeasonListItem struct {
	Season       GetCreativeCenterSeasonListSeasonInfo `json:"season"`        // 合集信息
	Course       any                                   `json:"course"`        // 课程信息。实测为 null
	Checkin      GetCreativeCenterSeasonListCheckin    `json:"checkin"`       // 审核信息
	SeasonStat   GetCreativeCenterSeasonListSeasonStat `json:"seasonStat"`    // 合集统计信息
	Sections     GetCreativeCenterSeasonListSections   `json:"sections"`      // 小节列表
	PartEpisodes []GetCreativeCenterSeasonListEpisode  `json:"part_episodes"` // 合集视频列表。可能为 null
}

// GetCreativeCenterSeasonListSeasonInfo 是合集列表中 season 对象。
type GetCreativeCenterSeasonListSeasonInfo struct {
	ID             int    `json:"id"`               // 合集 ID
	Title          string `json:"title"`            // 合集标题
	Desc           string `json:"desc"`             // 合集描述
	Cover          string `json:"cover"`            // 合集封面 URL
	IsEnd          int    `json:"isEnd"`            // 是否已完结。0：未完结。1：已完结
	Mid            int    `json:"mid"`              // 合集作者 ID
	IsAct          int    `json:"isAct"`            // 是否为活动合集。0：否。1：是
	IsPay          int    `json:"is_pay"`           // 是否付费。0：否。1：是
	State          int    `json:"state"`            // 合集状态。0：正常显示。-6：正在审核
	PartState      int    `json:"partState"`        // 合集分段状态。0：正常
	SignState      int    `json:"signState"`        // 合集签名状态。0：正常
	RejectReason   string `json:"rejectReason"`     // 合集拒绝原因
	Ctime          int    `json:"ctime"`            // 创建时间。UNIX 时间戳
	Mtime          int    `json:"mtime"`            // 修改时间。UNIX 时间戳
	NoSection      int    `json:"no_section"`       // 是否设小节。1：不设小节
	Forbid         int    `json:"forbid"`           // 合集是否禁止。0：否。1：是
	ProtocolID     string `json:"protocol_id"`      // 协议 ID。实测为空
	EpNum          int    `json:"ep_num"`           // 视频数量
	SeasonPrice    int    `json:"season_price"`     // 合集价格。0：免费
	IsOpened       int    `json:"is_opened"`        // 是否公开。1：公开。0：不公开
	HasChargingPay int    `json:"has_charging_pay"` // 是否充电付费。0：否。1：是
	HasPugvPay     int    `json:"has_pugv_pay"`     // 是否 PUGV 付费。0：否。1：是
}

// GetCreativeCenterSeasonListCheckin 是合集列表 seasons 数组中的 checkin 对象。
type GetCreativeCenterSeasonListCheckin struct {
	Status       int    `json:"status"`        // 状态。0：正常显示
	StatusReason string `json:"status_reason"` // 状态原因
	SeasonStatus int    `json:"season_status"` // 合集审核状态。1：审核通过
}

// GetCreativeCenterSeasonListSeasonStat 是合集列表 seasons 数组中的 seasonStat 对象。
type GetCreativeCenterSeasonListSeasonStat struct {
	View         int `json:"view"`         // 播放量
	Danmaku      int `json:"danmaku"`      // 弹幕数
	Reply        int `json:"reply"`        // 评论数
	Fav          int `json:"fav"`          // 收藏数
	Coin         int `json:"coin"`         // 硬币数
	Share        int `json:"share"`        // 分享数
	NowRank      int `json:"nowRank"`      // 当前排名
	HisRank      int `json:"hisRank"`      // 历史最高排名
	Like         int `json:"like"`         // 点赞数
	Subscription int `json:"subscription"` // 订阅数
	Vt           int `json:"vt"`           // 0。作用尚不明确
}

// GetCreativeCenterSeasonListSections 是合集列表 seasons 数组中的 sections 对象。
// 文档称「套了个娃」：外层对象里还有一个 sections 数组。
type GetCreativeCenterSeasonListSections struct {
	Sections []GetCreativeCenterSeasonListSection `json:"sections"` // 小节列表
}

// GetCreativeCenterSeasonListSection 是合集列表中的合集小节信息。
//
// 获取合集列表只返回前 14 个字段；order 文档标为 str、cover 标为 num，
// 但实测分别是数字与字符串，按实测取值。
type GetCreativeCenterSeasonListSection struct {
	ID             int    `json:"id"`               // 小节 ID
	Type           int    `json:"type"`             // 小节类型。实测为 1
	SeasonID       int    `json:"seasonId"`         // 合集 ID
	Title          string `json:"title"`            // 小节标题
	Order          int    `json:"order"`            // 排序编号
	IsEnd          int    `json:"isEnd"`            // 是否已完结。0：未完结
	State          int    `json:"state"`            // 合集状态。0：正常显示。-6：正在审核
	PartState      int    `json:"partState"`        // 合集分段状态。0
	RejectReason   string `json:"rejectReason"`     // 合集拒绝原因
	Ctime          int    `json:"ctime"`            // 创建时间。UNIX 时间戳
	Mtime          int    `json:"mtime"`            // 修改时间。UNIX 时间戳
	EpCount        int    `json:"epCount"`          // 小节视频总数
	Cover          string `json:"cover"`            // 合集封面 URL
	Show           int    `json:"show"`             // 是否公开。1：公开
	HasChargingPay int    `json:"has_charging_pay"` // 是否充电付费。0：否
	Episodes       any    `json:"Episodes"`         // 实测为 null
	HasPugvPay     int    `json:"has_pugv_pay"`     // 是否 PUGV 付费。0：否
	CaptchaToken   string `json:"captcha_token"`    // 验证码 token。实测为空
}

// GetCreativeCenterSeasonListEpisode 是合集列表 part_episodes 数组中的合集内视频。
type GetCreativeCenterSeasonListEpisode struct {
	ID           int    `json:"id"`           // 合集内视频 ID
	Title        string `json:"title"`        // 视频标题
	Aid          int    `json:"aid"`          // 视频 aid
	Bvid         string `json:"bvid"`         // 视频 bvid
	Cid          int    `json:"cid"`          // 视频 cid
	SeasonID     int    `json:"seasonId"`     // 合集 ID
	SectionID    int    `json:"sectionId"`    // 小节 ID
	Order        int    `json:"order"`        // 排序编号
	VideoTitle   string `json:"videoTitle"`   // 视频标题。合集列表接口实测为空
	ArchiveTitle string `json:"archiveTitle"` // 稿件标题
	ArchiveState int    `json:"archiveState"` // 稿件状态。实测为 0
	RejectReason string `json:"rejectReason"` // 拒绝理由
	State        int    `json:"state"`        // 状态。实测为 0
	Cover        string `json:"cover"`        // 封面 URL
	IsFree       int    `json:"is_free"`      // 是否免费。0：免费
	AidOwner     bool   `json:"aid_owner"`    // 是否视频所有者。true：是所有者
	ChargingPay  int    `json:"charging_pay"` // 充电付费。0：否
	MemberFirst  int    `json:"member_first"` // 是否为会员抢先看。实测为 0
	PugvPay      int    `json:"pugv_pay"`     // 是否为付费课程。实测为 0
	LimitedFree  bool   `json:"limited_free"` // 是否限时免费。实测为 false
}

// GetCreativeCenterSeasonList 获取合集列表
func (c *Client) GetCreativeCenterSeasonList(ctx context.Context, param GetCreativeCenterSeasonListParam) (*GetCreativeCenterSeasonListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x2/creative/web/seasons"
	)
	return bilibili.Execute[*GetCreativeCenterSeasonListResult](ctx, c.Client, method, url, param)
}

// GetCreativeCenterSeasonSectionVideosParam 是获取合集小节中视频的请求参数。
type GetCreativeCenterSeasonSectionVideosParam struct {
	ID int `json:"id" request:"query"` // 合集 ID
}

// GetCreativeCenterSeasonSectionVideosResult 是获取合集小节中的视频的 data 对象。
type GetCreativeCenterSeasonSectionVideosResult struct {
	Section  GetCreativeCenterSeasonSectionVideosSection   `json:"section"`  // 小节信息
	Episodes []GetCreativeCenterSeasonSectionVideosEpisode `json:"episodes"` // 小节中的视频
}

// GetCreativeCenterSeasonSectionVideosSection 是获取合集小节中的视频的小节信息。
//
// 该接口还会返回 isEnd、show、has_pugv_pay、captcha_token。
type GetCreativeCenterSeasonSectionVideosSection struct {
	ID             int    `json:"id"`               // 小节 ID
	Type           int    `json:"type"`             // 小节类型。实测为 1
	SeasonID       int    `json:"seasonId"`         // 合集 ID
	Title          string `json:"title"`            // 小节标题
	Order          int    `json:"order"`            // 排序编号
	IsEnd          int    `json:"isEnd"`            // 是否已完结。0：未完结
	State          int    `json:"state"`            // 合集状态。0：正常显示。-6：正在审核
	PartState      int    `json:"partState"`        // 合集分段状态。0
	RejectReason   string `json:"rejectReason"`     // 合集拒绝原因
	Ctime          int    `json:"ctime"`            // 创建时间。UNIX 时间戳
	Mtime          int    `json:"mtime"`            // 修改时间。UNIX 时间戳
	EpCount        int    `json:"epCount"`          // 小节视频总数
	Cover          string `json:"cover"`            // 合集封面 URL
	Show           int    `json:"show"`             // 是否公开。1：公开
	HasChargingPay int    `json:"has_charging_pay"` // 是否充电付费。0：否
	Episodes       any    `json:"Episodes"`         // 实测为 null
	HasPugvPay     int    `json:"has_pugv_pay"`     // 是否 PUGV 付费。0：否
	CaptchaToken   string `json:"captcha_token"`    // 验证码 token。实测为空
}

// GetCreativeCenterSeasonSectionVideosEpisode 是获取合集小节中的视频的 episodes 数组元素。
//
// member_first、pugv_pay、limited_free 只在小节接口返回。
type GetCreativeCenterSeasonSectionVideosEpisode struct {
	ID           int    `json:"id"`           // 合集内视频 ID
	Title        string `json:"title"`        // 视频标题
	Aid          int    `json:"aid"`          // 视频 aid
	Bvid         string `json:"bvid"`         // 视频 bvid
	Cid          int    `json:"cid"`          // 视频 cid
	SeasonID     int    `json:"seasonId"`     // 合集 ID
	SectionID    int    `json:"sectionId"`    // 小节 ID
	Order        int    `json:"order"`        // 排序编号
	VideoTitle   string `json:"videoTitle"`   // 视频标题。合集列表接口实测为空
	ArchiveTitle string `json:"archiveTitle"` // 稿件标题
	ArchiveState int    `json:"archiveState"` // 稿件状态。实测为 0
	RejectReason string `json:"rejectReason"` // 拒绝理由
	State        int    `json:"state"`        // 状态。实测为 0
	Cover        string `json:"cover"`        // 封面 URL
	IsFree       int    `json:"is_free"`      // 是否免费。0：免费
	AidOwner     bool   `json:"aid_owner"`    // 是否视频所有者。true：是所有者
	ChargingPay  int    `json:"charging_pay"` // 充电付费。0：否
	MemberFirst  int    `json:"member_first"` // 是否为会员抢先看。实测为 0
	PugvPay      int    `json:"pugv_pay"`     // 是否为付费课程。实测为 0
	LimitedFree  bool   `json:"limited_free"` // 是否限时免费。实测为 false
}

// GetCreativeCenterSeasonSectionVideos 获取合集小节中的视频
//
// 可以在给合集排序时使用，可以获取别人的合集。
func (c *Client) GetCreativeCenterSeasonSectionVideos(ctx context.Context, param GetCreativeCenterSeasonSectionVideosParam) (*GetCreativeCenterSeasonSectionVideosResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x2/creative/web/season/section"
	)
	return bilibili.Execute[*GetCreativeCenterSeasonSectionVideosResult](ctx, c.Client, method, url, param)
}

// GetCreativeCenterSeasonByAidParam 是通过视频 aid 反查所属合集的请求参数。
type GetCreativeCenterSeasonByAidParam struct {
	ID int `json:"id" request:"query"` // 视频 aid
}

// GetCreativeCenterSeasonByAidResult 是通过视频 aid 反查所属合集的 data 对象。
//
// has_pugv_pay 只在该接口返回。
type GetCreativeCenterSeasonByAidResult struct {
	ID             int    `json:"id"`               // 合集 ID
	Title          string `json:"title"`            // 合集标题
	Desc           string `json:"desc"`             // 合集描述
	Cover          string `json:"cover"`            // 合集封面 URL
	IsEnd          int    `json:"isEnd"`            // 是否已完结。0：未完结。1：已完结
	Mid            int    `json:"mid"`              // 合集作者 ID
	IsAct          int    `json:"isAct"`            // 是否为活动合集。0：否。1：是
	IsPay          int    `json:"is_pay"`           // 是否付费。0：否。1：是
	State          int    `json:"state"`            // 合集状态。0：正常显示。-6：正在审核
	PartState      int    `json:"partState"`        // 合集分段状态。0：正常
	SignState      int    `json:"signState"`        // 合集签名状态。0：正常
	RejectReason   string `json:"rejectReason"`     // 合集拒绝原因
	Ctime          int    `json:"ctime"`            // 创建时间。UNIX 时间戳
	Mtime          int    `json:"mtime"`            // 修改时间。UNIX 时间戳
	NoSection      int    `json:"no_section"`       // 是否设小节。1：不设小节
	Forbid         int    `json:"forbid"`           // 合集是否禁止。0：否。1：是
	ProtocolID     string `json:"protocol_id"`      // 协议 ID。实测为空
	EpNum          int    `json:"ep_num"`           // 视频数量
	SeasonPrice    int    `json:"season_price"`     // 合集价格。0：免费
	IsOpened       int    `json:"is_opened"`        // 是否公开。1：公开。0：不公开
	HasChargingPay int    `json:"has_charging_pay"` // 是否充电付费。0：否。1：是
	HasPugvPay     int    `json:"has_pugv_pay"`     // 是否 PUGV 付费。0：否。1：是。仅 aid 反查接口返回
}

// GetCreativeCenterSeasonByAid 通过视频 aid 反查所属合集
func (c *Client) GetCreativeCenterSeasonByAid(ctx context.Context, param GetCreativeCenterSeasonByAidParam) (*GetCreativeCenterSeasonByAidResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x2/creative/web/season/aid"
	)
	return bilibili.Execute[*GetCreativeCenterSeasonByAidResult](ctx, c.Client, method, url, param)
}
