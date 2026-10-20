package blackroom

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 风纪委员仲裁案件信息相关接口。各接口的返回体独立定义，互不复用。

// GetJuryCaseInfoParam 是获取单个案件信息的请求参数。
type GetJuryCaseInfoParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必要
	CaseID    string `json:"case_id"`                                        // 案件 id
}

// GetJuryCaseInfoResult 是单个仲裁案件信息（新版 API）的 data 对象。
type GetJuryCaseInfoResult struct {
	CaseID      string                    `json:"case_id"`      // 案件 id
	CaseType    int                       `json:"case_type"`    // 任务类型。1：单条评论。2：评论氛围。3：单条弹幕。4：弹幕氛围
	JuryState   int                       `json:"jury_state"`   // 1。改版后未见过其它值
	Participate int                       `json:"participate"`  // 是否投票。0：未投票。1：已投票
	VoteItems   []GetJuryCaseInfoVoteItem `json:"vote_items"`   // 投票选项
	DefaultVote int                       `json:"default_vote"` // 默认投票选项。4 或 14 即「无法判断」
	Status      int                       `json:"status"`       // 裁决状态（新版）。0：进行中。1：已结束
	OriginStart int                       `json:"origin_start"` // 案件关联视频，播放进度条起始坐标。单位：秒
	Avid        int64                     `json:"avid"`         // 视频 av 号
	Cid         int64                     `json:"cid"`          // 弹幕 dmid
	VoteCd      int                       `json:"vote_cd"`      // -1。暂不明确
	Result      int                       `json:"result"`       // 众议结论
	ResultText  string                    `json:"result_text"`  // 众议结论文本
	Title       string                    `json:"title"`        // 视频标题
	CaseInfo    GetJuryCaseInfoContent    `json:"case_info"`    // 案件信息
	MyPoint     *GetJuryCaseInfoPoint     `json:"my_point"`     // 我的观点。仅自己投过票才有该字段
	VoteInfo    GetJuryCaseInfoVoteInfo   `json:"vote_info"`    // 投票信息
}

// GetJuryCaseInfoVoteItem 是单个仲裁案件信息中的投票选项。
type GetJuryCaseInfoVoteItem struct {
	Vote     int    `json:"vote"`      // 投票选项码
	VoteText string `json:"vote_text"` // 投票选项文本
}

// GetJuryCaseInfoContent 是单个仲裁案件信息中的案件信息本体，不同 case_type 下只有对应字段有值。
type GetJuryCaseInfoContent struct {
	Comment     *GetJuryCaseInfoComment  `json:"comment"`      // 单条评论信息。仅 case_type 为 1 时有值
	DanmuImg    string                   `json:"danmu_img"`    // 弹幕截图 url。仅 case_type 为 4 时有值
	Comments    []GetJuryCaseInfoComment `json:"comments"`     // 若干条评论。仅 case_type 为 2 时有值
	SingleDanmu *GetJuryCaseInfoDanmu    `json:"single_danmu"` // 单条弹幕信息。仅 case_type 为 3 时有值
}

// GetJuryCaseInfoComment 是单个仲裁案件信息中涉及的评论信息。
type GetJuryCaseInfoComment struct {
	Mid           int64                    `json:"mid"`            // 用户 uid
	Uname         string                   `json:"uname"`          // 用户名
	Face          string                   `json:"face"`           // 用户头像 url
	Content       string                   `json:"content"`        // 评论内容
	ChildComments []GetJuryCaseInfoComment `json:"child_comments"` // 子评论（对该评论的回复）。单条评论和子评论下无该字段
}

// GetJuryCaseInfoDanmu 是单个仲裁案件信息中涉及的单条弹幕信息。
type GetJuryCaseInfoDanmu struct {
	IDStr    string `json:"id_str"`   // 弹幕 dmid
	Content  string `json:"content"`  // 弹幕内容
	Progress int    `json:"progress"` // 尚不明确
}

// GetJuryCaseInfoPoint 是单个仲裁案件信息中风纪委员的投票（观点）信息。
type GetJuryCaseInfoPoint struct {
	Opid       int64  `json:"opid"`        // 观点编号。每当某个风纪委员给某个仲裁案件发表观点，就 +1
	Mid        int64  `json:"mid"`         // 风纪委员用户 id
	Uname      string `json:"uname"`       // 风纪委员用户名。勾选匿名投票则显示「匿名用户」
	Face       string `json:"face"`        // 头像 url
	Vote       int    `json:"vote"`        // 投票选项。合适题为 1~4，氛围题为 11~14
	VoteText   string `json:"vote_text"`   // 内容为空。现在系统已不再展示其他风纪委所投的票
	Content    string `json:"content"`     // 观点内容
	Anonymous  int    `json:"anonymous"`   // 0。尚不明确
	Like       int    `json:"like"`        // 被点赞数量
	Hate       int    `json:"hate"`        // 被点踩数量
	LikeStatus int    `json:"like_status"` // 本账号的点赞（踩）状态。0：没做操作。1：点赞。2：点踩
	VoteTime   int64  `json:"vote_time"`   // 投票时间。时间戳，精确到秒
	Insiders   int    `json:"insiders"`    // 是否观看此类视频。1：平时观看。0：平时不观看
}

// GetJuryCaseInfoVoteInfo 是单个仲裁案件信息中的案件投票统计。
type GetJuryCaseInfoVoteInfo struct {
	AllCount       int   `json:"all_count"`       // 总投票数
	Counts         []int `json:"counts"`          // 观点分布。格式：[合适, 一般, 不合适, 无法判断]
	InsidersCounts []int `json:"insiders_counts"` // 行为分布。格式：[平时会看此类视频, 平时不看此类视频]
}

// GetJuryCaseInfo 获取单个案件信息
//
// 可查询任意案件，无论自己是否参与投票。需要登录（Cookie 或 APP）。
func (c *Client) GetJuryCaseInfo(ctx context.Context, param GetJuryCaseInfoParam) (*GetJuryCaseInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/case/info"
	)
	return bilibili.Execute[*GetJuryCaseInfoResult](ctx, c.Client, method, url, param)
}

// GetJuryCaseInfoLegacyParam 是查询案件众裁信息（旧版 API）的请求参数。
type GetJuryCaseInfoLegacyParam struct {
	CID int64 `json:"cid"` // 案件 id
}

// GetJuryCaseInfoLegacyResult 是旧版 API 的仲裁案件信息。
type GetJuryCaseInfoLegacyResult struct {
	ID            int64  `json:"id"`            // 案件 id
	Mid           int64  `json:"mid"`           // 被举报用户 mid
	Status        int    `json:"status"`        // 裁决状态。1：投票中。2：停止发放。3：复核中（结案中）。4：已裁决。5：待重启。6：未裁决。7：冻结中。8：队列中
	StatusTitle   string `json:"statusTitle"`   // 封禁时间文字。judgeType=1 时有效
	OriginType    int    `json:"originType"`    // 来源类型
	ReasonType    int    `json:"reasonType"`    // 原因类型
	OriginContent string `json:"originContent"` // 被举报的原文
	PunishResult  int    `json:"punishResult"`  // 处理结果。0：无。1：封禁 3 天。2：封禁 7 天。3：永封。4：其他封禁天数。5：封禁 15 天。6：仅删除不封禁
	PunishTitle   string `json:"punishTitle"`   // 处罚原因
	JudgeType     int    `json:"judgeType"`     // 裁决类型。status=4 时有效。0：未裁决。1：违规。2：不违规
	OriginURL     string `json:"originUrl"`     // 来源 url
	BlockedDays   int    `json:"blockedDays"`   // 封禁时间。judgeType=1 时有效
	PutTotal      int    `json:"putTotal"`      // 作用尚不明确
	VoteRule      int    `json:"voteRule"`      // 投票不违规人数
	VoteBreak     int    `json:"voteBreak"`     // 投票建议封禁人数
	VoteDelete    int    `json:"voteDelete"`    // 投票建议删除人数
	StartTime     int64  `json:"startTime"`     // 裁决起始时间。时间戳
	EndTime       int64  `json:"endTime"`       // 裁决截止时间。时间戳，未结案固定为 24 小时
	Ctime         int64  `json:"ctime"`         // 举报创建时间。时间戳
	Mtime         int64  `json:"mtime"`         // 开始拉取时间。时间戳
	OriginTitle   string `json:"originTitle"`   // 来源标题
	RelationID    string `json:"relationId"`    // 违规关系标识字。{评论 rpid}-{违规类型 id}-{来源对象 id}
	Face          string `json:"face"`          // 被举报用户头像 url
	Uname         string `json:"uname"`         // 被举报用户昵称
	Vote          int    `json:"vote"`          // 我的投票
	VoteTime      int64  `json:"voteTime"`      // 我的裁决时间。时间戳，单位毫秒。仅旧版列表返回
	CaseType      int    `json:"case_type"`     // 众裁类型。0：小众众裁。1：大众众裁
}

// GetJuryCaseInfoLegacy 查询案件众裁信息（旧版 API）
//
// 该接口无需登录。data 在案件无效时为 null。
func (c *Client) GetJuryCaseInfoLegacy(ctx context.Context, param GetJuryCaseInfoLegacyParam) (*GetJuryCaseInfoLegacyResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/jury/caseInfo"
	)
	return bilibili.Execute[*GetJuryCaseInfoLegacyResult](ctx, c.Client, method, url, param)
}

// GetJuryCaseListParam 是查询我的众裁记录的请求参数。
type GetJuryCaseListParam struct {
	Pn int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
	Ps int `json:"ps,omitempty" request:"query,omitempty"` // 每页显示数量。默认为 30，最大 30
}

// GetJuryCaseListResult 是我的众裁记录（新版 API）的 data 对象。
type GetJuryCaseListResult struct {
	Total int                   `json:"total"` // 仲裁案例数
	List  []GetJuryCaseListItem `json:"list"`  // 仲裁记录列表，按时间顺序
}

// GetJuryCaseListItem 是我的众裁记录条目。
type GetJuryCaseListItem struct {
	CaseID   string `json:"case_id"`   // 案件 id
	CaseType int    `json:"case_type"` // 案件类型。1：单条评论。2：评论氛围。3：单条弹幕。4：弹幕氛围
	Status   int    `json:"status"`    // 裁决状态。0：投票中。1：已结束
	Vote     int    `json:"vote"`      // 投票选项。单条评论（弹幕）为 1~4，评论（弹幕）氛围为 11~14
	VoteText string `json:"vote_text"` // 投票选项文本
	VoteTime int64  `json:"vote_time"` // 投票时间。时间戳，精确到秒
}

// GetJuryCaseList 查询我的众裁记录
//
// 需要登录（Cookie）。
func (c *Client) GetJuryCaseList(ctx context.Context, param GetJuryCaseListParam) (*GetJuryCaseListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/case/list"
	)
	return bilibili.Execute[*GetJuryCaseListResult](ctx, c.Client, method, url, param)
}

// GetJuryCaseListLegacyParam 是查询我的众裁记录（旧版 API）的请求参数。
type GetJuryCaseListLegacyParam struct {
	AccessKey string `json:"access_key,omitempty" request:"query,omitempty"` // APP 登录 Token。APP 方式必要
	Pn        int    `json:"pn,omitempty" request:"query,omitempty"`         // 页码。默认为 1
	Ps        int    `json:"ps,omitempty" request:"query,omitempty"`         // 每页显示数量。默认为 30，最大 30
}

// GetJuryCaseListLegacyItem 是旧版 API 的一条众裁记录。
type GetJuryCaseListLegacyItem struct {
	ID            int64  `json:"id"`            // 案件 id
	Mid           int64  `json:"mid"`           // 被举报用户 mid
	Status        int    `json:"status"`        // 裁决状态。1：投票中。2：停止发放。3：复核中（结案中）。4：已裁决。5：待重启。6：未裁决。7：冻结中。8：队列中
	StatusTitle   string `json:"statusTitle"`   // 封禁时间文字。judgeType=1 时有效
	OriginType    int    `json:"originType"`    // 来源类型
	ReasonType    int    `json:"reasonType"`    // 原因类型
	OriginContent string `json:"originContent"` // 被举报的原文
	PunishResult  int    `json:"punishResult"`  // 处理结果。0：无。1：封禁 3 天。2：封禁 7 天。3：永封。4：其他封禁天数。5：封禁 15 天。6：仅删除不封禁
	PunishTitle   string `json:"punishTitle"`   // 处罚原因
	JudgeType     int    `json:"judgeType"`     // 裁决类型。status=4 时有效。0：未裁决。1：违规。2：不违规
	OriginURL     string `json:"originUrl"`     // 来源 url
	BlockedDays   int    `json:"blockedDays"`   // 封禁时间。judgeType=1 时有效
	PutTotal      int    `json:"putTotal"`      // 作用尚不明确
	VoteRule      int    `json:"voteRule"`      // 投票不违规人数
	VoteBreak     int    `json:"voteBreak"`     // 投票建议封禁人数
	VoteDelete    int    `json:"voteDelete"`    // 投票建议删除人数
	StartTime     int64  `json:"startTime"`     // 裁决起始时间。时间戳
	EndTime       int64  `json:"endTime"`       // 裁决截止时间。时间戳，未结案固定为 24 小时
	Ctime         int64  `json:"ctime"`         // 举报创建时间。时间戳
	Mtime         int64  `json:"mtime"`         // 开始拉取时间。时间戳
	OriginTitle   string `json:"originTitle"`   // 来源标题
	RelationID    string `json:"relationId"`    // 违规关系标识字。{评论 rpid}-{违规类型 id}-{来源对象 id}
	Face          string `json:"face"`          // 被举报用户头像 url
	Uname         string `json:"uname"`         // 被举报用户昵称
	Vote          int    `json:"vote"`          // 我的投票
	VoteTime      int64  `json:"voteTime"`      // 我的裁决时间。时间戳，单位毫秒。仅旧版列表返回
	CaseType      int    `json:"case_type"`     // 众裁类型。0：小众众裁。1：大众众裁
}

// GetJuryCaseListLegacy 查询我的众裁记录（旧版 API）
//
// 需要登录（Cookie 或 APP）。
func (c *Client) GetJuryCaseListLegacy(ctx context.Context, param GetJuryCaseListLegacyParam) ([]GetJuryCaseListLegacyItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/jury/caseList"
	)
	return bilibili.Execute[[]GetJuryCaseListLegacyItem](ctx, c.Client, method, url, param)
}

// GetJuryCaseOpinionParam 是获取众议观点的请求参数。
type GetJuryCaseOpinionParam struct {
	CaseID string `json:"case_id"`                                // 仲裁案件 id
	Pn     int    `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为 1
	Ps     int    `json:"ps,omitempty" request:"query,omitempty"` // 每页显示数量。默认为 10，最大 20
}

// GetJuryCaseOpinionResult 是众议观点（新版 API）的 data 对象。
type GetJuryCaseOpinionResult struct {
	Total int                      `json:"total"` // 当前案件众议观点总数
	List  []GetJuryCaseOpinionItem `json:"list"`  // 观点列表。无数据时为空数组
}

// GetJuryCaseOpinionItem 是众议观点列表中的一条观点。
type GetJuryCaseOpinionItem struct {
	Opid       int64  `json:"opid"`        // 观点编号。每当某个风纪委员给某个仲裁案件发表观点，就 +1
	Mid        int64  `json:"mid"`         // 风纪委员用户 id
	Uname      string `json:"uname"`       // 风纪委员用户名。勾选匿名投票则显示「匿名用户」
	Face       string `json:"face"`        // 头像 url
	Vote       int    `json:"vote"`        // 投票选项。合适题为 1~4，氛围题为 11~14
	VoteText   string `json:"vote_text"`   // 内容为空。现在系统已不再展示其他风纪委所投的票
	Content    string `json:"content"`     // 观点内容
	Anonymous  int    `json:"anonymous"`   // 0。尚不明确
	Like       int    `json:"like"`        // 被点赞数量
	Hate       int    `json:"hate"`        // 被点踩数量
	LikeStatus int    `json:"like_status"` // 本账号的点赞（踩）状态。0：没做操作。1：点赞。2：点踩
	VoteTime   int64  `json:"vote_time"`   // 投票时间。时间戳，精确到秒
	Insiders   int    `json:"insiders"`    // 是否观看此类视频。1：平时观看。0：平时不观看
}

// GetJuryCaseOpinion 获取众议观点
//
// 需要登录（Cookie）。ps 大于 20 时返回 -400。
func (c *Client) GetJuryCaseOpinion(ctx context.Context, param GetJuryCaseOpinionParam) (*GetJuryCaseOpinionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/credit/v2/jury/case/opinion"
	)
	return bilibili.Execute[*GetJuryCaseOpinionResult](ctx, c.Client, method, url, param)
}
