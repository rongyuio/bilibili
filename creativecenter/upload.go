package creativecenter

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 创作中心投稿上传相关接口。文档见 creativecenter/upload.md。
//
// 各接口的返回体独立定义，即使字段相同也不复用。

// GetUploadTemplateListParam 指定获取上传模板列表的查询条件。
type GetUploadTemplateListParam struct {
	T int64 `json:"t,omitempty" request:"query,omitempty"` // 当前 UNIX 毫秒时间戳
}

// GetUploadTemplateListItem 是上传模板列表中的一项。
type GetUploadTemplateListItem struct {
	Tid         int    `json:"tid"`         // 模板 ID
	Name        string `json:"name"`        // 模板名称
	Typeid      int    `json:"typeid"`      // 分区 ID
	Title       string `json:"title"`       // 标题
	Tags        string `json:"tags"`        // 标签
	Description string `json:"description"` // 描述
	Copyright   int    `json:"copyright"`   // 版权类型。1：自制；2：转载
	Attribute   int    `json:"attribute"`   // 属性。0 或其他
	IsDefault   int    `json:"is_default"`  // 是否默认。0：否；1：是
}

// GetUploadTemplateList 获取上传模板列表。
func (c *Client) GetUploadTemplateList(ctx context.Context, param GetUploadTemplateListParam) ([]GetUploadTemplateListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/tpls"
	)
	return bilibili.Execute[[]GetUploadTemplateListItem](ctx, c.Client, method, url, param)
}

// GetTopicListParam 指定查询投稿话题的参数。
type GetTopicListParam struct {
	TypeID int    `json:"type_id,omitempty" request:"query,omitempty"` // 分区 ID
	Pn     int    `json:"pn" request:"query"`                          // 页码。从 0 开始
	Ps     int    `json:"ps" request:"query"`                          // 每页个数
	Title  string `json:"title,omitempty" request:"query,omitempty"`   // 视频标题
	T      int64  `json:"t,omitempty" request:"query,omitempty"`       // 当前 UNIX 毫秒时间戳
}

// GetTopicListResult 是查询投稿话题的 data 部分。
type GetTopicListResult struct {
	Topics    []GetTopicListItem `json:"topics"`     // 话题列表
	Tags      any                `json:"tags"`       // 标签。实测为 null
	Maxpage   int                `json:"maxpage"`    // 最大页数。实测为 210
	RequestID string             `json:"request_id"` // 请求 ID
}

// GetTopicListItem 是投稿话题列表中的一项。
type GetTopicListItem struct {
	TopicID             int    `json:"topic_id"`             // 话题 ID
	TopicName           string `json:"topic_name"`           // 话题名称
	Description         string `json:"description"`          // 话题描述
	MissionID           int    `json:"mission_id"`           // 任务 ID
	ActivityText        string `json:"activity_text"`        // 活动文本
	ActivityDescription string `json:"activity_description"` // 活动描述
	ArcPlayVV           int    `json:"arc_play_vv"`          // 稿件总播放量
}

// GetTopicList 查询投稿时可选的话题。
func (c *Client) GetTopicList(ctx context.Context, param GetTopicListParam) (*GetTopicListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/topic/type"
	)
	return bilibili.Execute[*GetTopicListResult](ctx, c.Client, method, url, param)
}

// SearchTopicParam 指定话题搜索的参数。
type SearchTopicParam struct {
	PageSize int    `json:"page_size,omitempty" request:"query,omitempty"` // 每页个数
	Offset   int    `json:"offset,omitempty" request:"query,omitempty"`    // 个数偏移。并非页数
	Keywords string `json:"keywords,omitempty" request:"query,omitempty"`  // 关键字
	T        int64  `json:"t,omitempty" request:"query,omitempty"`         // 当前 UNIX 毫秒时间戳
}

// SearchTopicResult 是话题搜索的 data 部分。
type SearchTopicResult struct {
	Result SearchTopicResultContent `json:"result"` // 搜索结果
}

// SearchTopicResultContent 是话题搜索的 result 对象。
type SearchTopicResultContent struct {
	HasCreateJurisdiction bool                `json:"has_create_jurisdiction"` // 是否拥有创建话题的权限
	IsNewTopic            bool                `json:"is_new_topic"`            // 是否为新话题
	Tips                  string              `json:"tips"`                    // 提示信息
	PageInfo              SearchTopicPageInfo `json:"page_info"`               // 分页信息
	Topics                []SearchTopicItem   `json:"topics"`                  // 话题列表
}

// SearchTopicPageInfo 是话题搜索的分页信息。
type SearchTopicPageInfo struct {
	HasMore    bool `json:"has_more"`    // 是否有更多
	Offset     int  `json:"offset"`      // 偏移量
	PageNumber int  `json:"page_number"` // 页码
}

// SearchTopicItem 是话题搜索结果中的一项。
type SearchTopicItem struct {
	ActProtocol  string `json:"act_protocol"`  // 作用尚不明确
	ActivitySign string `json:"activity_sign"` // 作用尚不明确
	Description  string `json:"description"`   // 话题描述
	ID           int    `json:"id"`            // 话题 ID
	MissionID    int    `json:"mission_id"`    // 任务 ID
	Name         string `json:"name"`          // 话题名称
	State        int    `json:"state"`         // 状态。作用尚不明确
	Uname        string `json:"uname"`         // 作用尚不明确
}

// SearchTopic 搜索投稿话题。
func (c *Client) SearchTopic(ctx context.Context, param SearchTopicParam) (*SearchTopicResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/topic/search"
	)
	return bilibili.Execute[*SearchTopicResult](ctx, c.Client, method, url, param)
}

// CheckTagAvailabilityParam 指定标签可用性检查的参数。
type CheckTagAvailabilityParam struct {
	Tag string `json:"tag" request:"query"` // 需要检查的标签
}

// CheckTagAvailabilityResult 是标签可用性检查的 data 部分。
type CheckTagAvailabilityResult struct {
	Code    int    `json:"code"`    // 标签状态。0：可用；其他：不可用
	Content string `json:"content"` // 错误返回
}

// CheckTagAvailability 检查投稿标签是否可用。
func (c *Client) CheckTagAvailability(ctx context.Context, param CheckTagAvailabilityParam) (*CheckTagAvailabilityResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/topic/tag/check"
	)
	return bilibili.Execute[*CheckTagAvailabilityResult](ctx, c.Client, method, url, param)
}

// GetArchiveDescFormatParam 指定获取简介相关信息的参数。
type GetArchiveDescFormatParam struct {
	TypeID    int   `json:"typeid" request:"query"`                // 分区 ID
	Copyright int   `json:"copyright" request:"query"`             // 版权类型。1：自制；2：转载
	T         int64 `json:"t,omitempty" request:"query,omitempty"` // 当前 UNIX 毫秒时间戳
}

// GetArchiveDescFormatResult 是投稿简介相关信息的 data 部分。
type GetArchiveDescFormatResult struct {
	Typeid     int    `json:"typeid"`     // 分区 ID
	ID         int    `json:"id"`         // ID
	Lang       int    `json:"lang"`       // 未知
	Copyright  int    `json:"copyright"`  // 版权类型。1：自制；2：转载
	Components string `json:"components"` // 简介输入框提示文字。JSON 字符串
}

// GetArchiveDescFormat 获取投稿简介的相关信息，包括分区提示与简介输入框配置。
// data 过去可能为 null（此时简介上限 250 字），因此返回指针。
func (c *Client) GetArchiveDescFormat(ctx context.Context, param GetArchiveDescFormatParam) (*GetArchiveDescFormatResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/archive/desc/format"
	)
	return bilibili.Execute[*GetArchiveDescFormatResult](ctx, c.Client, method, url, param)
}

// GetUploadLinesResult 是获取上传线路接口的根对象。
type GetUploadLinesResult struct {
	OK        int                  `json:"OK"`         // 返回值。1：成功；其他：失败
	Lines     []GetUploadLinesLine `json:"lines"`      // 上传线路列表
	Probe     GetUploadLinesProbe  `json:"probe"`      // 探测信息
	QuicLines []GetUploadLinesLine `json:"quic_lines"` // QUIC 上传线路列表。结构同 lines
}

// GetUploadLinesLine 是一条上传线路。
type GetUploadLinesLine struct {
	OS       string `json:"os"`        // 操作系统
	Query    string `json:"query"`     // 查询参数
	ProbeURL string `json:"probe_url"` // 探测 URL
}

// GetUploadLinesProbe 是上传线路的探测信息。
type GetUploadLinesProbe struct {
	Post float64 `json:"post"` // 作用尚不明确。固定为 0.1
}

// GetUploadLines 获取视频上传线路。
//
// 该接口不使用标准的 code/message/data 信封，故走 ExecuteRaw 并手动解码。
func (c *Client) GetUploadLines(ctx context.Context) (*GetUploadLinesResult, error) {
	const (
		method   = resty.MethodGet
		endpoint = "https://member.bilibili.com/preupload"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, endpoint, nil, bilibili.FillParam("r", "probe"))
	if err != nil {
		return nil, err
	}
	var result GetUploadLinesResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, bilibili.NewDecodeError(method, endpoint, body, reflect.TypeFor[GetUploadLinesResult](), "$", 0, err)
	}
	return &result, nil
}

// GetRecommendedTagsParam 指定预测稿件标签的参数。
type GetRecommendedTagsParam struct {
	UploadID    string `json:"upload_id,omitempty" request:"query,omitempty"`   // 上传 ID
	SubtypeID   int    `json:"subtype_id,omitempty" request:"query,omitempty"`  // 子分区 ID
	Title       string `json:"title,omitempty" request:"query,omitempty"`       // 视频标题
	Filename    string `json:"filename,omitempty" request:"query,omitempty"`    // 视频文件名。无后缀名
	Description string `json:"description,omitempty" request:"query,omitempty"` // 视频简介
	CoverURL    string `json:"cover_url,omitempty" request:"query,omitempty"`   // 视频封面 URL。不含 https: 或 http: 字串
	T           int64  `json:"t,omitempty" request:"query,omitempty"`           // 当前 UNIX 毫秒时间戳
}

// GetRecommendedTagsItem 是预测稿件标签结果中的一项。
type GetRecommendedTagsItem struct {
	Tag       string `json:"tag"`        // 标签名称
	Checked   int    `json:"checked"`    // 0
	RequestID string `json:"request_id"` // 请求 ID
}

// GetRecommendedTags 根据稿件信息预测投稿标签。
func (c *Client) GetRecommendedTags(ctx context.Context, param GetRecommendedTagsParam) ([]GetRecommendedTagsItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/tag/recommend"
	)
	return bilibili.Execute[[]GetRecommendedTagsItem](ctx, c.Client, method, url, param)
}

// GetHumanTypeListParam 指定获取新分区 ID 列表的参数。
type GetHumanTypeListParam struct {
	T int64 `json:"t,omitempty" request:"query,omitempty"` // 当前 UNIX 毫秒时间戳
}

// GetHumanTypeListResult 是获取新分区 ID 列表的 data 部分。
type GetHumanTypeListResult struct {
	TypeList []GetHumanTypeListItem `json:"type_list"` // 新分区列表
}

// GetHumanTypeListItem 是一个新分区。
type GetHumanTypeListItem struct {
	ID   int    `json:"id"`   // 分区 ID
	Name string `json:"name"` // 分区名
}

// GetHumanTypeList 获取投稿页使用的新分区 ID 列表。
func (c *Client) GetHumanTypeList(ctx context.Context, param GetHumanTypeListParam) (*GetHumanTypeListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/vupre/web/archive/human/type2/list"
	)
	return bilibili.Execute[*GetHumanTypeListResult](ctx, c.Client, method, url, param)
}

// GetUploadMetadataParam 指定获取上传元数据（预上传）的参数。
type GetUploadMetadataParam struct {
	Name         string `json:"name" request:"query"`                              // 文件名。会影响返回的上传地址
	R            string `json:"r" request:"query"`                                 // 上传区域。upos
	Profile      string `json:"profile" request:"query"`                           // 上传配置。普通视频：ugcfx/bup；提交反馈：feedback/bup
	ProbeVersion int    `json:"probe_version,omitempty" request:"query,omitempty"` // 上传版本。20221109
	Upcdn        string `json:"upcdn,omitempty" request:"query,omitempty"`         // 上传 CDN。txa
	Zone         string `json:"zone,omitempty" request:"query,omitempty"`          // 上传区域。cs
	SSL          int    `json:"ssl,omitempty" request:"query,omitempty"`           // 是否使用 SSL。0
	Version      string `json:"version,omitempty" request:"query,omitempty"`       // 上传版本。2.14.0.0
	Build        string `json:"build,omitempty" request:"query,omitempty"`         // 上传版本。2140000
	Size         int64  `json:"size,omitempty" request:"query,omitempty"`          // 视频文件大小。单位字节
	WebVersion   string `json:"webVersion,omitempty" request:"query,omitempty"`    // 上传版本。2.13.0
}

// GetUploadMetadataResult 是获取上传元数据（预上传）接口的根对象。
type GetUploadMetadataResult struct {
	OK              int      `json:"OK"`                // 返回值。1：成功
	Auth            string   `json:"auth"`              // 上传凭证。作为后续请求的请求头，有效期 5 天
	BizID           int64    `json:"biz_id"`            // 业务 ID
	ChunkRetry      int      `json:"chunk_retry"`       // 重试次数
	ChunkRetryDelay int      `json:"chunk_retry_delay"` // 重试延迟
	ChunkSize       int64    `json:"chunk_size"`        // 分块大小
	Endpoint        string   `json:"endpoint"`          // 上传节点
	Endpoints       []string `json:"endpoints"`         // 上传节点列表
	ExposeParams    any      `json:"expose_params"`     // 作用尚不明确。实测为 null
	PutQuery        string   `json:"put_query"`         // 上传参数
	Threads         int      `json:"threads"`           // 上传线程数
	Timeout         int      `json:"timeout"`           // 超时时间
	UIP             string   `json:"uip"`               // 你的 IP
	UposURI         string   `json:"upos_uri"`          // 上传地址
}

// GetUploadMetadata 获取上传元数据（预上传）。
//
// 该接口不使用标准的 code/message/data 信封，故走 ExecuteRaw 并手动解码。
// 返回的 Auth 有效期 5 天，Endpoint 与 UposURI 用于拼接后续的上传地址。
func (c *Client) GetUploadMetadata(ctx context.Context, param GetUploadMetadataParam) (*GetUploadMetadataResult, error) {
	const (
		method   = resty.MethodGet
		endpoint = "https://member.bilibili.com/preupload"
	)
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, endpoint, param)
	if err != nil {
		return nil, err
	}
	var result GetUploadMetadataResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, bilibili.NewDecodeError(method, endpoint, body, reflect.TypeFor[GetUploadMetadataResult](), "$", 0, err)
	}
	return &result, nil
}
