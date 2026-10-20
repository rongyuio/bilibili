package blackroom

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 风纪委员会用户基本信息相关接口。各接口的返回体独立定义，互不复用。

// GetJuryInfoParam 是获取风纪委员会用户基本数据的请求参数。
type GetJuryInfoParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必要
}

// GetJuryInfoResult 是获取风纪委员会用户基本数据的 data 对象。
type GetJuryInfoResult struct {
	AllowApply  bool   `json:"allow_apply"`  // 尚不明确（用户从未担任风纪委员时也为 true）
	ApplyStatus int    `json:"apply_status"` // 任期审核状态。-1：资格失效，且未申请。0：刚申请连任时。5：申请后，等待审核。3：申请连任成功，status 为 1
	CaseTotal   int    `json:"case_total"`   // 总众裁数
	ErrMsg      string `json:"err_msg"`      // 尚不明确
	Face        string `json:"face"`         // 用户头像 url
	Status      int    `json:"status"`       // 当前资格状态。0：未曾拥有资格。1：任期内。2：资格失效
	TermEnd     int64  `json:"term_end"`     // 任期结束时间。时间戳（秒级），无任期时为 0
	Uname       string `json:"uname"`        // 用户昵称
}

// GetJuryInfo 获取风纪委员会用户基本数据
//
// 需要登录（Cookie 或 APP）。
func (c *Client) GetJuryInfo(ctx context.Context, param GetJuryInfoParam) (*GetJuryInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/jury"
	)
	return bilibili.Execute[*GetJuryInfoResult](ctx, c.Client, method, url, param)
}

// GetJuryKpiParam 是获取风纪委员会统计信息（新版 API）的请求参数。
type GetJuryKpiParam struct {
	TermID int64 `json:"term_id,omitempty" request:"query,omitempty"` // 任期 id。不传或传 0 时只返回最近一次已结束的任期
}

// GetJuryKpiResult 是获取风纪委员会统计信息（新版 API）的 data 对象。
type GetJuryKpiResult struct {
	Mid          int64             `json:"mid"`           // 用户 UID
	Uname        string            `json:"uname"`         // 留空
	Face         string            `json:"face"`          // 留空
	TermID       int               `json:"term_id"`       // 任期 ID。未传入 term_id 时，返回最近一次已结束任期的 id
	TermStart    int64             `json:"term_start"`    // 任期开始时间。时间戳，精确到秒
	TermEnd      int64             `json:"term_end"`      // 任期结束时间。时间戳，精确到秒
	CaseTotal    int               `json:"case_total"`    // 任内总投票数
	ActiveDays   int               `json:"active_days"`   // 活跃天数
	LikeNum      int               `json:"like_num"`      // 发表观点，被点赞数
	AccuracyRate int               `json:"accuracy_rate"` // 投中率。所选观点与大多数风纪委员一致，记为一次“投中”
	Pass         int               `json:"pass"`          // 任期是否合格。0：不合格。1：合格
	Status       int               `json:"status"`        // 当前资格状态。0：未曾拥有资格。1：任期内。2：资格失效
	ApplyStatus  int               `json:"apply_status"`  // 当前资格申请状态。-1：资格失效，且未申请。0：刚申请连任时。5：申请后，等待审核。3：申请连任成功，status 为 1
	PrevTermID   int               `json:"prev_term_id"`  // 上一任期的任期 ID。新版风纪委员启用后的第一个任期为 0
	NextTermID   int               `json:"next_term_id"`  // 下一任期的任期 ID。当前任期尚未结束时为 0
	Rewards      GetJuryKpiRewards `json:"rewards"`       // 任期奖励
}

// GetJuryKpiRewards 是风纪委员会统计信息（新版 API）中的任期奖励。
type GetJuryKpiRewards struct {
	Pendant int `json:"pendant"` // 头像挂件礼包 ID。合格后一般为 1，不合格为 0
	Coin    int `json:"coin"`    // 硬币礼包 ID。合格后一般为 3，不合格为 0
}

// GetJuryKpi 获取风纪委员会统计信息（新版 API）
//
// 只收录 2021 年 10 月风纪委员改版之后的数据，每次只返回一次任期的数据。
// 需要登录（Cookie）。业务错误码：25018 不能进行此操作（还没有成为新风纪委员，或 term_id 不属于自己）。
func (c *Client) GetJuryKpi(ctx context.Context, param GetJuryKpiParam) (*GetJuryKpiResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/kpi"
	)
	return bilibili.Execute[*GetJuryKpiResult](ctx, c.Client, method, url, param)
}

// GetJuryKpiLegacyParam 是获取风纪委员会统计信息（旧版 API）的请求参数。
type GetJuryKpiLegacyParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必要
}

// GetJuryKpiLegacyItem 是风纪委员会某个月的统计信息（旧版 API）。
type GetJuryKpiLegacyItem struct {
	ID           int   `json:"id"`            // 0
	Mid          int64 `json:"mid"`           // 用户 id
	Number       int   `json:"number"`        // 风纪委员编号
	Day          int64 `json:"day"`           // 数据生成时间（也是任期结束时间）。时间戳，精确到秒
	Rate         int   `json:"rate"`          // 任期完成度。1：A。2：S。3：S+。4：S++
	Rank         int   `json:"rank"`          // 尚不明确
	RankPer      int   `json:"rankper"`       // 尚不明确
	RankTotal    int   `json:"rankTotal"`     // 尚不明确
	Point        int   `json:"point"`         // 32767
	ActiveDays   int   `json:"activeDays"`    // 活跃天数
	VoteTotal    int   `json:"voteTotal"`     // 投票总数
	VoteRadio    int   `json:"voteRadio"`     // 投中率
	BlockedTotal int   `json:"blockedTotal"`  // 尚不明确
	TermStart    int64 `json:"termStart"`     // 任期开始时间。时间戳，精确到秒
	TermEnd      int64 `json:"termEnd"`       // 任期结束时间（与数据生成时间一致）。时间戳，精确到秒
	OpinionLikes int   `json:"opinion_likes"` // 观点获赞
}

// GetJuryKpiLegacy 获取风纪委员会统计信息（旧版 API）
//
// 只统计到 2021 年 9 月任期（风纪委员改版之前），改版之后的数据未收录。
// 需要登录（Cookie 或 APP）。
func (c *Client) GetJuryKpiLegacy(ctx context.Context, param GetJuryKpiLegacyParam) ([]GetJuryKpiLegacyItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/jury/kpi"
	)
	return bilibili.Execute[[]GetJuryKpiLegacyItem](ctx, c.Client, method, url, param)
}

// GetJuryRequirementParam 是检查是否具有申请风纪委员会资格的请求参数。
type GetJuryRequirementParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必要
}

// GetJuryRequirementResult 是申请风纪委员会的资格检查结果。
type GetJuryRequirementResult struct {
	Blocked bool  `json:"blocked"` // 是否被封禁。true：是。false：否
	Cert    bool  `json:"cert"`    // 是否实名认证
	Level   *bool `json:"level"`   // 等级是否 >= 4。仅旧版 API 返回，新版无该字段
	Rule    bool  `json:"rule"`    // 是否 90 天内无违规
}

// GetJuryRequirement 检查是否具有申请风纪委员会的资格
//
// 使用新版 API（无等级限制）。需要登录（Cookie 或 APP）。
func (c *Client) GetJuryRequirement(ctx context.Context, param GetJuryRequirementParam) (*GetJuryRequirementResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/requirement"
	)
	return bilibili.Execute[*GetJuryRequirementResult](ctx, c.Client, method, url, param)
}
