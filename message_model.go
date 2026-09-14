package bilibili

// 私信相关响应模型。

type UnreadMessage struct {
	At     int `json:"at"`      // 未读at数
	Chat   int `json:"chat"`    // 0。作用尚不明确
	Like   int `json:"like"`    // 未读点赞数
	Reply  int `json:"reply"`   // 未读回复数
	SysMsg int `json:"sys_msg"` // 未读系统通知数
	Up     int `json:"up"`      // UP主助手信息数
}

type UnreadPrivateMessage struct {
	UnfollowUnread int `json:"unfollow_unread"` // 未关注用户未读私信数
	FollowUnread   int `json:"follow_unread"`   // 已关注用户未读私信数
	Gt             int `json:"_gt_"`            // 0
}

type SendPrivateMessageResult struct {
	MsgKey      int    `json:"msg_key"`       // 消息唯一id
	MsgContent  string `json:"msg_content"`   // 发送的消息
	KeyHitInfos any    `json:"key_hit_infos"` // 作用尚不明确
}

type Message struct {
	SenderUID      int    `json:"sender_uid"`       // 发送者uid。注意名称是sender_uid
	ReceiverType   int    `json:"receiver_type"`    // 与session_type对应。1为用户，2为粉丝团
	ReceiverID     int    `json:"receiver_id"`      // 接收者uid。注意名称是receiver_id
	MsgType        int    `json:"msg_type"`         // 消息类型。1:文字消息。2:图片消息。5:撤回的消息。12、13:通知
	Content        string `json:"content"`          // 消息内容。此处存在设计缺陷
	MsgSeqno       int    `json:"msg_seqno"`        // 消息序列号，保证按照时间顺序从小到大
	Timestamp      int    `json:"timestamp"`        // 消息发送时间戳
	AtUIDs         []int  `json:"at_uids"`          // 未知
	MsgKey         int    `json:"msg_key"`          // 未知
	MsgStatus      int    `json:"msg_status"`       // 消息状态。0
	NotifyCode     string `json:"notify_code"`      // 未知
	NewFaceVersion int    `json:"new_face_version"` // 表情包版本。0或者没有是旧版，此时b站会自动转换成新版表情包，例如[doge] -> [tv_doge]；1是新版
}

type EInfo struct {
	Text string `json:"text"` // 表情名称
	URI  string `json:"uri"`  // 表情链接
	Size int    `json:"size"` // 表情尺寸。1
}

type PrivateMessageRecords struct {
	Messages []Message `json:"messages"`  // 聊天记录列表
	HasMore  int       `json:"has_more"`  // 0
	MinSeqno uint64    `json:"min_seqno"` // 所有消息最小的序列号（最早）
	MaxSeqno uint64    `json:"max_seqno"` // 所有消息最大的序列号（最晚）
	EInfos   []EInfo   `json:"e_infos"`   // 聊天表情列表
}

// PrivateMessageLastMsg 是会话中最后一条消息。
type PrivateMessageLastMsg struct {
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

// PrivateMessageAccountInfo 是会话对方的账号信息。
type PrivateMessageAccountInfo struct {
	Name   string `json:"name"`
	PicURL string `json:"pic_url"`
}

// PrivateMessageSession 是一个私信会话。
type PrivateMessageSession struct {
	TalkerID          int64                     `json:"talker_id"`
	SessionType       int                       `json:"session_type"`
	AtSeqno           int                       `json:"at_seqno"`
	TopTs             int                       `json:"top_ts"`
	GroupName         string                    `json:"group_name"`
	GroupCover        string                    `json:"group_cover"`
	IsFollow          int                       `json:"is_follow"`
	IsDnd             int                       `json:"is_dnd"`
	AckSeqno          int64                     `json:"ack_seqno"`
	AckTs             int64                     `json:"ack_ts"`
	SessionTs         int64                     `json:"session_ts"`
	UnreadCount       int                       `json:"unread_count"`
	LastMsg           PrivateMessageLastMsg     `json:"last_msg"`
	GroupType         int                       `json:"group_type"`
	CanFold           int                       `json:"can_fold"`
	Status            int                       `json:"status"`
	MaxSeqno          int64                     `json:"max_seqno"`
	NewPushMsg        int                       `json:"new_push_msg"`
	Setting           int                       `json:"setting"`
	IsGuardian        int                       `json:"is_guardian"`
	IsIntercept       int                       `json:"is_intercept"`
	IsTrust           int                       `json:"is_trust"`
	SystemMsgType     int                       `json:"system_msg_type"`
	LiveStatus        int                       `json:"live_status"`
	BizMsgUnreadCount int                       `json:"biz_msg_unread_count"`
	AccountInfo       PrivateMessageAccountInfo `json:"account_info,omitempty"`
}

type PrivateMessageList struct {
	SessionList         []PrivateMessageSession `json:"session_list"`
	HasMore             int                     `json:"has_more"`
	AntiDisturbCleaning bool                    `json:"anti_disturb_cleaning"`
	IsAddressListEmpty  int                     `json:"is_address_list_empty"`
	SystemMsg           map[string]int64        `json:"system_msg"`
	ShowLevel           bool                    `json:"show_level"`
}
