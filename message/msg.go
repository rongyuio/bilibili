package message

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 消息通知相关接口。各接口的返回体独立定义，互不复用。

// GetMessageReplyParam 是获取"回复我的"信息的请求参数。
type GetMessageReplyParam struct {
	Build       int    `json:"build,omitempty" request:"query,omitempty"`        // 客户端内部版本号。默认为0
	MobiApp     string `json:"mobi_app,omitempty" request:"query,omitempty"`     // 平台标识。可为 web 等
	ID          int64  `json:"id,omitempty" request:"query,omitempty"`           // 起始 id。详情见返回的 cursor.id
	ReplyTime   int64  `json:"reply_time,omitempty" request:"query,omitempty"`   // 起始时间。详情见返回的 cursor.time
	Platform    string `json:"platform,omitempty" request:"query,omitempty"`     // 平台标识。可为 web 等
	WebLocation string `json:"web_location,omitempty" request:"query,omitempty"` // 暂时不知道是啥
}

// GetMessageReplyCursor 是下一次查询的指针。
type GetMessageReplyCursor struct {
	IsEnd bool  `json:"is_end"` // 是否结束。false 表示还有更多数据
	ID    int64 `json:"id"`     // 最后（旧）一条通知的 id，用作下次查询的起始 id
	Time  int64 `json:"time"`   // 最后一条通知的时间。unix 时间戳
}

// GetMessageReplyUser 是回复者或被@的用户信息。
type GetMessageReplyUser struct {
	Mid      int64  `json:"mid"`      // 用户 mid
	Fans     int    `json:"fans"`     // 粉丝数。固定为0，可能不返回实际数据
	Nickname string `json:"nickname"` // 用户昵称
	Avatar   string `json:"avatar"`   // 头像 URL
	MidLink  string `json:"mid_link"` // 用户主页链接。固定为空字符串
	Follow   bool   `json:"follow"`   // 是否关注该用户。false 表示未关注
}

// GetMessageReplyAtDetail 是通知详情中被@的用户信息。
type GetMessageReplyAtDetail struct {
	Mid      int64  `json:"mid"`      // 用户 mid
	Fans     int    `json:"fans"`     // 粉丝数。固定为0，可能不返回实际数据
	Nickname string `json:"nickname"` // 用户昵称
	Avatar   string `json:"avatar"`   // 头像 URL
	MidLink  string `json:"mid_link"` // 用户主页链接。固定为空字符串
	Follow   bool   `json:"follow"`   // 是否关注该用户。false 表示未关注
}

// GetMessageReplyDetail 是回复通知的详情。
type GetMessageReplyDetail struct {
	SubjectID          int64                     `json:"subject_id"`           // 主体 id
	RootID             int64                     `json:"root_id"`              // 根评论 id。最顶层的评论 ID
	SourceID           int64                     `json:"source_id"`            // 源评论 id。直接回复的评论 ID
	TargetID           int64                     `json:"target_id"`            // 目标评论 id。被回复的评论 ID
	Type               string                    `json:"type"`                 // 通知类型。固定为 "reply"（回复类型）
	BusinessID         int                       `json:"business_id"`          // 业务类型 id。1：视频评论。11：动态评论
	Business           string                    `json:"business"`             // 业务名称。"评论" 或 "视频"
	Title              string                    `json:"title"`                // 通知标题。摘要文本
	Desc               string                    `json:"desc"`                 // 描述。固定为空字符串
	Image              string                    `json:"image"`                // 图片 URL。固定为空字符串
	URI                string                    `json:"uri"`                  // 跳转链接。web 端跳转链接
	NativeURI          string                    `json:"native_uri"`           // 客户端跳转链接。客户端专用跳转链接
	DetailTitle        string                    `json:"detail_title"`         // 详细标题。固定为空字符串
	RootReplyContent   string                    `json:"root_reply_content"`   // 根评论内容。最顶层评论的文本内容
	SourceContent      string                    `json:"source_content"`       // 源评论内容。直接回复的评论内容
	TargetReplyContent string                    `json:"target_reply_content"` // 目标评论内容。被回复的评论内容
	AtDetails          []GetMessageReplyAtDetail `json:"at_details"`           // @的用户列表
	TopicDetails       []any                     `json:"topic_details"`        // 话题详情。固定为空数组
	HideReplyButton    bool                      `json:"hide_reply_button"`    // 是否隐藏回复按钮
	HideLikeButton     bool                      `json:"hide_like_button"`     // 是否隐藏点赞按钮
	LikeState          int                       `json:"like_state"`           // 点赞状态。0：未点赞
	Danmu              any                       `json:"danmu"`                // 弹幕信息。固定为 null
	Message            string                    `json:"message"`              // 消息内容。固定为空字符串
}

// GetMessageReplyItem 是一条回复通知。
type GetMessageReplyItem struct {
	ID        int64                 `json:"id"`         // 通知 id
	User      GetMessageReplyUser   `json:"user"`       // 回复者的用户信息
	Item      GetMessageReplyDetail `json:"item"`       // 通知详情
	Counts    int                   `json:"counts"`     // 通知计数。固定为1，可能表示单条通知
	IsMulti   int                   `json:"is_multi"`   // 是否多回复。固定为0，可能表示单条回复
	ReplyTime int64                 `json:"reply_time"` // 回复时间。unix 时间戳
}

// GetMessageReplyResult 是"回复我的"信息。
type GetMessageReplyResult struct {
	Cursor     GetMessageReplyCursor `json:"cursor"`       // 光标，下一次查询的指针
	Items      []GetMessageReplyItem `json:"items"`        // 通知列表，每个元素代表一条回复通知
	LastViewAt int64                 `json:"last_view_at"` // 上次查看的时间。unix 时间戳
}

// GetMessageReply 获取"回复我的"信息
func (c *Client) GetMessageReply(ctx context.Context, param GetMessageReplyParam) (*GetMessageReplyResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/msgfeed/reply"
	)
	return bilibili.Execute[*GetMessageReplyResult](ctx, c.Client, method, url, param)
}
