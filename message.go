package bilibili

import (
	"context"
	"crypto/rand"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// 私信相关接口。响应模型见 message_model.go。

// GetUnreadMessage 获取未读消息数
func (c *Client) GetUnreadMessage(ctx context.Context) (*UnreadMessage, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/msgfeed/unread"
	)
	return execute[*UnreadMessage](ctx, c, method, url, nil)
}

// GetUnreadPrivateMessage 获取未读私信数
func (c *Client) GetUnreadPrivateMessage(ctx context.Context) (*UnreadPrivateMessage, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/session_svr/v1/session_svr/single_unread"
	)
	return execute[*UnreadPrivateMessage](ctx, c, method, url, nil)
}

var deviceID string

func init() {
	b := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	s := []byte("xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx")
	randBytes := make([]byte, len(s))
	_, _ = rand.Read(randBytes)
	for i := range s {
		if s[i] == '-' || s[i] == '4' {
			continue
		}
		j := randBytes[i] % 16
		if s[i] == 'x' {
			s[i] = b[j]
		} else {
			s[i] = b[3&j|8]
		}
	}
	deviceID = string(s)
}

type SendPrivateMessageParam struct {
	SenderUID      int         `json:"msg[sender_uid]"`                                           // 发送者mid
	ReceiverID     int         `json:"msg[receiver_id]"`                                          // 接收者mid
	ReceiverType   int         `json:"msg[receiver_type]"`                                        // 1。固定为1
	MsgType        int         `json:"msg[msg_type]"`                                             // 消息类型。1:发送文字。2:发送图片。5:撤回消息
	MsgStatus      int         `json:"msg[msg_status],omitempty" request:"query,omitempty"`       // 0
	Timestamp      int         `json:"msg[timestamp]"`                                            // 时间戳（秒）
	NewFaceVersion int         `json:"msg[new_face_version],omitempty" request:"query,omitempty"` // 表情包版本
	Content        json.Number `json:"msg[content]"`                                              // 消息内容。发送文字时：str<br />撤回消息时：num
	DeviceID       string      `json:"-" request:"query,field=msg[dev_id]"`                       // 设备 id，留空时由 SendPrivateMessage 填入进程内随机生成的设备标识
}

// SendPrivateMessage 发送私信（文字消息）
func (c *Client) SendPrivateMessage(ctx context.Context, param SendPrivateMessageParam) (*SendPrivateMessageResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/web_im/v1/web_im/send_msg"
	)
	if param.DeviceID == "" {
		param.DeviceID = deviceID
	}
	return execute[*SendPrivateMessageResult](ctx, c, method, url, param, fillCsrf(c))
}

type GetPrivateMessageRecordsParam struct {
	TalkerID       int    `json:"talker_id"`                                            // 聊天对象的uid
	SenderDeviceID int    `json:"sender_device_id,omitempty" request:"query,omitempty"` // 发送者设备。1
	SessionType    int    `json:"session_type"`                                         // 聊天对象的类型。1为用户，2为粉丝团
	Size           int    `json:"size,omitempty" request:"query,omitempty"`             // 列出消息条数。默认是20，最大为200
	Build          int    `json:"build,omitempty" request:"query,omitempty"`            // 未知。默认是0
	MobiApp        string `json:"mobi_app,omitempty" request:"query,omitempty"`         // 设备。web
	BeginSeqno     int    `json:"begin_seqno,omitempty" request:"query,omitempty"`      // 开始的序列号。默认0为全部
	EndSeqno       int    `json:"end_seqno,omitempty" request:"query,omitempty"`        // 结束的序列号。默认0为全部
}

// GetPrivateMessageRecords 获取与聊天对象的私信消息记录
func (c *Client) GetPrivateMessageRecords(ctx context.Context, param GetPrivateMessageRecordsParam) (*PrivateMessageRecords, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/svr_sync/v1/svr_sync/fetch_session_msgs"
	)
	return execute[*PrivateMessageRecords](ctx, c, method, url, param)
}

type GetPrivateMessageListParam struct {
	SessionType int    `json:"session_type"`                                 // 1：系统，2：用户，3：应援团
	MobiApp     string `json:"mobi_app,omitempty" request:"query,omitempty"` // 设备
}

// GetPrivateMessageList 获取消息列表 session_type，1：系统，2：用户，3：应援团
//
// 参照 https://github.com/CuteReimu/bilibili/issues/8
func (c *Client) GetPrivateMessageList(ctx context.Context, param GetPrivateMessageListParam) (*PrivateMessageList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/session_svr/v1/session_svr/get_sessions"
	)
	return execute[*PrivateMessageList](ctx, c, method, url, param)
}
