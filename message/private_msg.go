package message

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 私信会话相关接口。各接口的返回体独立定义，互不复用。

// GetUnreadGroupPrivateMessageParam 是获取未读粉丝团私信数的请求参数。
type GetUnreadGroupPrivateMessageParam struct {
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetUnreadGroupPrivateMessageResult 是获取未读粉丝团私信数接口的 data 对象。
type GetUnreadGroupPrivateMessageResult struct {
	UnreadCount int `json:"unread_count"` // 未读粉丝团私信数
}

// GetUnreadGroupPrivateMessage 获取未读粉丝团私信数
func (c *Client) GetUnreadGroupPrivateMessage(ctx context.Context, param GetUnreadGroupPrivateMessageParam) (*GetUnreadGroupPrivateMessageResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/session_svr/v1/session_svr/my_group_unread"
	)
	return bilibili.Execute[*GetUnreadGroupPrivateMessageResult](ctx, c.Client, method, url, param)
}

// GetNewPrivateMessageSessionsParam 是获取新会话列表的请求参数。
type GetNewPrivateMessageSessionsParam struct {
	BeginTs int64  `json:"begin_ts,omitempty" request:"query,omitempty"` // 起始时间。微秒级时间戳
	Size    int    `json:"size,omitempty" request:"query,omitempty"`     // 返回的会话数。默认为 20，最大为 100
	Build   int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetNewPrivateMessageSessionsLastMsg 是新会话中最后一条消息。
type GetNewPrivateMessageSessionsLastMsg struct {
	SenderUID      int64  `json:"sender_uid"`
	ReceiverType   int    `json:"receiver_type"`
	ReceiverID     int    `json:"receiver_id"`
	MsgType        int    `json:"msg_type"`
	Content        string `json:"content"`
	MsgSeqno       int64  `json:"msg_seqno"`
	Timestamp      int    `json:"timestamp"`
	MsgKey         int64  `json:"msg_key"`
	MsgStatus      int    `json:"msg_status"`
	NotifyCode     string `json:"notify_code"`
	NewFaceVersion int    `json:"new_face_version,omitempty"`
}

// GetNewPrivateMessageSessionsAccountInfo 是新会话对方的账号信息。
type GetNewPrivateMessageSessionsAccountInfo struct {
	Name   string `json:"name"`
	PicURL string `json:"pic_url"`
}

// GetNewPrivateMessageSessionsSession 是一个新会话。
type GetNewPrivateMessageSessionsSession struct {
	TalkerID          int64                                   `json:"talker_id"`
	SessionType       int                                     `json:"session_type"`
	AtSeqno           int                                     `json:"at_seqno"`
	TopTs             int                                     `json:"top_ts"`
	GroupName         string                                  `json:"group_name"`
	GroupCover        string                                  `json:"group_cover"`
	IsFollow          int                                     `json:"is_follow"`
	IsDnd             int                                     `json:"is_dnd"`
	AckSeqno          int64                                   `json:"ack_seqno"`
	AckTs             int64                                   `json:"ack_ts"`
	SessionTs         int64                                   `json:"session_ts"`
	UnreadCount       int                                     `json:"unread_count"`
	LastMsg           GetNewPrivateMessageSessionsLastMsg     `json:"last_msg"`
	GroupType         int                                     `json:"group_type"`
	CanFold           int                                     `json:"can_fold"`
	Status            int                                     `json:"status"`
	MaxSeqno          int64                                   `json:"max_seqno"`
	NewPushMsg        int                                     `json:"new_push_msg"`
	Setting           int                                     `json:"setting"`
	IsGuardian        int                                     `json:"is_guardian"`
	IsIntercept       int                                     `json:"is_intercept"`
	IsTrust           int                                     `json:"is_trust"`
	SystemMsgType     int                                     `json:"system_msg_type"`
	LiveStatus        int                                     `json:"live_status"`
	BizMsgUnreadCount int                                     `json:"biz_msg_unread_count"`
	AccountInfo       GetNewPrivateMessageSessionsAccountInfo `json:"account_info,omitempty"`
}

// GetNewPrivateMessageSessionsResult 是获取新会话列表接口的 data 对象。
type GetNewPrivateMessageSessionsResult struct {
	SessionList         []GetNewPrivateMessageSessionsSession `json:"session_list"`          // 会话列表。无会话时为 null，按发送时间顺序逆向排序
	HasMore             int                                   `json:"has_more"`              // 是否有更多会话
	AntiDistrubCleaning bool                                  `json:"anti_distrub_cleaning"` // 是否开启了一键防骚扰功能。文档字段名如此拼写
	AntiDisturbCleaning bool                                  `json:"anti_disturb_cleaning"` // 是否开启免打扰清理。实测为 false
	IsAddressListEmpty  int                                   `json:"is_address_list_empty"` // 0。作用尚不明确
	ShowLevel           bool                                  `json:"show_level"`            // 是否在会话列表中显示用户等级。目前恒为 false
}

// GetNewPrivateMessageSessions 获取指定时间之后的新会话列表
func (c *Client) GetNewPrivateMessageSessions(ctx context.Context, param GetNewPrivateMessageSessionsParam) (*GetNewPrivateMessageSessionsResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/session_svr/v1/session_svr/new_sessions"
	)
	return bilibili.Execute[*GetNewPrivateMessageSessionsResult](ctx, c.Client, method, url, param)
}

// GetPrivateMessageSessionDetailParam 是获取会话详细信息的请求参数。
type GetPrivateMessageSessionDetailParam struct {
	TalkerID    int64  `json:"talker_id"`                                    // 聊天对象的 id。session_type 为 1 时表示用户 mid，为 2 时表示粉丝团 id
	SessionType int    `json:"session_type"`                                 // 聊天对象的类型。1：用户，2：粉丝团
	Build       int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp     string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetPrivateMessageSessionDetailLastMsg 是会话中最后一条消息。
type GetPrivateMessageSessionDetailLastMsg struct {
	SenderUID      int64  `json:"sender_uid"`
	ReceiverType   int    `json:"receiver_type"`
	ReceiverID     int    `json:"receiver_id"`
	MsgType        int    `json:"msg_type"`
	Content        string `json:"content"`
	MsgSeqno       int64  `json:"msg_seqno"`
	Timestamp      int    `json:"timestamp"`
	MsgKey         int64  `json:"msg_key"`
	MsgStatus      int    `json:"msg_status"`
	NotifyCode     string `json:"notify_code"`
	NewFaceVersion int    `json:"new_face_version,omitempty"`
}

// GetPrivateMessageSessionDetailAccountInfo 是会话对方的账号信息。
type GetPrivateMessageSessionDetailAccountInfo struct {
	Name   string `json:"name"`
	PicURL string `json:"pic_url"`
}

// GetPrivateMessageSessionDetailResult 是获取会话详细信息接口的 data 对象。
//
// 若从未与该会话聊过天，接口返回错误且 data 为 null，此时结果为 nil。
type GetPrivateMessageSessionDetailResult struct {
	TalkerID          int64                                     `json:"talker_id"`
	SessionType       int                                       `json:"session_type"`
	AtSeqno           int                                       `json:"at_seqno"`
	TopTs             int                                       `json:"top_ts"`
	GroupName         string                                    `json:"group_name"`
	GroupCover        string                                    `json:"group_cover"`
	IsFollow          int                                       `json:"is_follow"`
	IsDnd             int                                       `json:"is_dnd"`
	AckSeqno          int64                                     `json:"ack_seqno"`
	AckTs             int64                                     `json:"ack_ts"`
	SessionTs         int64                                     `json:"session_ts"`
	UnreadCount       int                                       `json:"unread_count"`
	LastMsg           GetPrivateMessageSessionDetailLastMsg     `json:"last_msg"`
	GroupType         int                                       `json:"group_type"`
	CanFold           int                                       `json:"can_fold"`
	Status            int                                       `json:"status"`
	MaxSeqno          int64                                     `json:"max_seqno"`
	NewPushMsg        int                                       `json:"new_push_msg"`
	Setting           int                                       `json:"setting"`
	IsGuardian        int                                       `json:"is_guardian"`
	IsIntercept       int                                       `json:"is_intercept"`
	IsTrust           int                                       `json:"is_trust"`
	SystemMsgType     int                                       `json:"system_msg_type"`
	LiveStatus        int                                       `json:"live_status"`
	BizMsgUnreadCount int                                       `json:"biz_msg_unread_count"`
	AccountInfo       GetPrivateMessageSessionDetailAccountInfo `json:"account_info,omitempty"`
}

// GetPrivateMessageSessionDetail 获取会话详细信息。
//
// 若从未与该会话聊过天，则返回 1000004 入口节点已存在错误；此时 data 为 null。
func (c *Client) GetPrivateMessageSessionDetail(ctx context.Context, param GetPrivateMessageSessionDetailParam) (*GetPrivateMessageSessionDetailResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/session_svr/v1/session_svr/session_detail"
	)
	return bilibili.Execute[*GetPrivateMessageSessionDetailResult](ctx, c.Client, method, url, param)
}

// GetPrivateMessageSessionLimitParam 是获取会话限制状态的请求参数。
type GetPrivateMessageSessionLimitParam struct {
	UID  int64 `json:"uid"`  // 聊天对象 mid
	Type int   `json:"type"` // 1
}

// GetPrivateMessageSessionLimitResult 是获取会话限制状态接口的 data 对象。
type GetPrivateMessageSessionLimitResult struct {
	IsLimit     int `json:"is_limit"`     // 用户是否被封禁
	ReportLimit int `json:"report_limit"` // 自己是否被限制举报该会话。常见于自己被封禁时出现该情况
}

// GetPrivateMessageSessionLimit 获取会话限制状态。仅支持用户会话。
func (c *Client) GetPrivateMessageSessionLimit(ctx context.Context, param GetPrivateMessageSessionLimitParam) (*GetPrivateMessageSessionLimitResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_setting/v1/link_setting/is_limit"
	)
	return bilibili.Execute[*GetPrivateMessageSessionLimitResult](ctx, c.Client, method, url, param)
}

// GetPrivateMessageSessionDndParam 是获取会话免打扰状态的请求参数。
type GetPrivateMessageSessionDndParam struct {
	OwnUID   int64  `json:"own_uid"`                                       // 自己的 mid
	UIDs     int64  `json:"uids,omitempty" request:"query,omitempty"`      // 用户 mid。仅支持填入 1 个 mid
	GroupIDs int64  `json:"group_ids,omitempty" request:"query,omitempty"` // 粉丝团 id。仅支持填入 1 个 id
	Build    int    `json:"build,omitempty" request:"query,omitempty"`     // 客户端内部版本号。默认为 0
	MobiApp  string `json:"mobi_app,omitempty" request:"query,omitempty"`  // 平台标识。可为 web 等
}

// GetPrivateMessageSessionDndSetting 是单个会话的免打扰状态。
type GetPrivateMessageSessionDndSetting struct {
	ID      int64 `json:"id"`      // 用户 mid 或粉丝团 id
	Setting int   `json:"setting"` // 免打扰状态。0：关闭免打扰，1：开启免打扰
}

// GetPrivateMessageSessionDndResult 是获取会话免打扰状态接口的 data 对象。
type GetPrivateMessageSessionDndResult struct {
	UIDSettings   []GetPrivateMessageSessionDndSetting `json:"uid_settings"`   // 用户免打扰状态。仅当请求参数 uids 存在时有此项
	GroupSettings []GetPrivateMessageSessionDndSetting `json:"group_settings"` // 粉丝团免打扰状态。仅当请求参数 group_ids 存在时有此项
}

// GetPrivateMessageSessionDnd 获取会话免打扰状态
func (c *Client) GetPrivateMessageSessionDnd(ctx context.Context, param GetPrivateMessageSessionDndParam) (*GetPrivateMessageSessionDndResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_setting/v1/link_setting/get_msg_dnd"
	)
	return bilibili.Execute[*GetPrivateMessageSessionDndResult](ctx, c.Client, method, url, param)
}

// GetPrivateMessageSessionPushSettingParam 是获取会话推送设置的请求参数。
type GetPrivateMessageSessionPushSettingParam struct {
	TalkerUID int64  `json:"talker_uid"`                                   // 聊天对象 mid
	Build     int    `json:"build,omitempty" request:"query,omitempty"`    // 客户端内部版本号。默认为 0
	MobiApp   string `json:"mobi_app,omitempty" request:"query,omitempty"` // 平台标识。可为 web 等
}

// GetPrivateMessageSessionPushSettingResult 是获取会话推送设置接口的 data 对象。
type GetPrivateMessageSessionPushSettingResult struct {
	FollowStatus    int `json:"follow_status"`     // 对方对于自己的关注属性。0：未关注，2：已关注，6：已互粉，128：已拉黑
	Special         int `json:"special"`           // 自己是否特别关注了对方
	PushSetting     int `json:"push_setting"`      // 推送设置。0：接收推送，1：不接收推送
	ShowPushSetting int `json:"show_push_setting"` // 是否显示推送设置
}

// GetPrivateMessageSessionPushSetting 获取会话推送设置。仅支持用户会话。
func (c *Client) GetPrivateMessageSessionPushSetting(ctx context.Context, param GetPrivateMessageSessionPushSettingParam) (*GetPrivateMessageSessionPushSettingResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/link_setting/v1/link_setting/get_session_ss"
	)
	return bilibili.Execute[*GetPrivateMessageSessionPushSettingResult](ctx, c.Client, method, url, param)
}

// GetPrivateMessageFeedInfoParam 是获取私信分享内容信息的请求参数。
type GetPrivateMessageFeedInfoParam struct {
	AIDs       []int64 `json:"aids,omitempty" request:"query,omitempty"`        // 视频 AV 号列表。与 ep_ids、article_ids 中至少存在一个，最多 50 个
	EpIDs      []int64 `json:"ep_ids,omitempty" request:"query,omitempty"`      // 番剧 epid 列表。与 aids、article_ids 中至少存在一个，最多 50 个
	ArticleIDs []int64 `json:"article_ids,omitempty" request:"query,omitempty"` // 专栏 CV 号列表。与 aids、ep_ids 中至少存在一个，无成员限制
	Build      int     `json:"build,omitempty" request:"query,omitempty"`       // 客户端内部版本号。默认为 0
	MobiApp    string  `json:"mobi_app"`                                        // 平台标识。可为 web 等
}

// GetPrivateMessageFeedInfoArchive 私信信息中的视频信息
type GetPrivateMessageFeedInfoArchive struct {
	Bvid      string `json:"bvid"`       // 视频 BV 号
	Aid       int64  `json:"aid"`        // 视频 AV 号
	Title     string `json:"title"`      // 视频标题。若视频失效则为「内容已失效」
	Pic       string `json:"pic"`        // 视频封面。若视频失效则为空文本
	Param     string `json:"param"`      // 提供的参数。即 AV 号的文本形式
	URI       string `json:"uri"`        // 跳转 url。bilibili://video/{视频AV号}
	Goto      string `json:"goto"`       // av
	Duration  int    `json:"duration"`   // 视频时长。以秒为单位，若视频失效则为 0
	UpName    string `json:"up_name"`    // 视频 UP 主昵称
	View      int64  `json:"view"`       // 视频播放量
	Danmaku   int64  `json:"danmaku"`    // 视频弹幕数
	Status    int    `json:"status"`     // 视频状态。0：正常，-1：已失效
	IsStarted int    `json:"is_started"` // 1。作用尚不明确
}

// GetPrivateMessageFeedInfoArticle 私信信息中的专栏信息
type GetPrivateMessageFeedInfoArticle struct {
	ID         int64    `json:"id"`          // 专栏 CV 号
	Title      string   `json:"title"`       // 专栏标题。若专栏失效则为「内容已失效」
	Summary    string   `json:"summary"`     // 专栏内容概要。若专栏失效则为空文本
	TemplateID int      `json:"template_id"` // 作用尚不明确
	UpName     string   `json:"up_name"`     // 专栏 UP 主昵称。若专栏失效则为空文本
	ImageURLs  []string `json:"image_urls"`  // 专栏封面列表。若专栏失效则为空数组
	ViewNum    int64    `json:"view_num"`    // 专栏观看数。若专栏失效则为 0
	LikeNum    int64    `json:"like_num"`    // 专栏点赞数。若专栏失效则为 0
	ReplyNum   int64    `json:"reply_num"`   // 专栏评论数。若专栏失效则为 0
	Status     int      `json:"status"`      // 专栏状态。0：正常，-1：已失效
}

// GetPrivateMessageFeedInfoPgc 私信信息中的番剧信息
type GetPrivateMessageFeedInfoPgc struct {
	EpID     int64  `json:"ep_id"`    // 番剧 epid
	Cover    string `json:"cover"`    // 番剧封面 url
	Title    string `json:"title"`    // 番剧分享标题。如《{番剧名}》 第{n}话 {单集标题} 等
	Duration int    `json:"duration"` // 番剧时长。以秒为单位
	View     int64  `json:"view"`     // 番剧播放量
	Danmaku  int64  `json:"danmaku"`  // 番剧弹幕数
	URL      string `json:"url"`      // 跳转 url。https://www.bilibili.com/bangumi/play/ep{番剧epid}
}

// GetPrivateMessageFeedInfoResult 是获取私信分享内容信息接口的 data 对象。
type GetPrivateMessageFeedInfoResult struct {
	Archive []GetPrivateMessageFeedInfoArchive `json:"archive"` // 视频信息列表。仅在指定了 aids 参数时存在此项
	Article []GetPrivateMessageFeedInfoArticle `json:"article"` // 专栏信息列表。仅在指定了 article_ids 参数时存在此项
	Pgc     []GetPrivateMessageFeedInfoPgc     `json:"pgc"`     // 番剧信息列表。仅在指定了 ep_ids 参数时存在此项
}

// GetPrivateMessageFeedInfo 获取多个视频、番剧、专栏的信息。
//
// 可用于获取私信中分享或推送的多个视频、番剧、专栏的信息。
func (c *Client) GetPrivateMessageFeedInfo(ctx context.Context, param GetPrivateMessageFeedInfoParam) (*GetPrivateMessageFeedInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/x/im/feed/infoweb"
	)
	return bilibili.Execute[*GetPrivateMessageFeedInfoResult](ctx, c.Client, method, url, param)
}
