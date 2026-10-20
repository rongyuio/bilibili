package message

// GetIMWebUnreadResult 是获取未读消息数（IM web 接口）的 data 对象。
type GetIMWebUnreadResult struct {
	At          int `json:"at"`            // 未读 at 数（旧）
	Chat        int `json:"chat"`          // 0。仅旧接口，作用尚不明确
	Coin        int `json:"coin"`          // 未读投币数
	Danmu       int `json:"danmu"`         // 未读弹幕数
	Favorite    int `json:"favorite"`      // 未读收藏数
	Like        int `json:"like"`          // 未读点赞数（旧）
	RecvLike    int `json:"recv_like"`     // 未读收到喜欢数（新），同 like
	RecvReply   int `json:"recv_reply"`    // 未读回复与 at 数（新），即 reply 与 at 之和
	Reply       int `json:"reply"`         // 未读回复数（旧）
	SysMsg      int `json:"sys_msg"`       // 未读系统通知数
	SysMsgStyle int `json:"sys_msg_style"` // 1。仅新接口，作用尚不明确
	Up          int `json:"up"`            // 未读 UP 主助手信息数
}
