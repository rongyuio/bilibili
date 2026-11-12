package livestream

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 直播弹幕常见命令的结构化事件。未知命令保留 Event.Raw，由调用方自行解析。

var errDanmakuDecodeEvent = errors.New("decode live danmaku event")

// DanmakuMessage 是 DANMU_MSG 的文本内容。
type DanmakuMessage struct {
	Text  string `json:"text"`
	UID   int64  `json:"uid"`
	Uname string `json:"uname"`
	Admin bool   `json:"admin"`
	VIP   bool   `json:"vip"`
	Guard int    `json:"guard"`
	Color int    `json:"color"`
	Mode  int    `json:"mode"`
}

// GiftMessage 是 SEND_GIFT 的主要内容。
type GiftMessage struct {
	GiftName  string `json:"giftName"`
	GiftID    int64  `json:"giftId"`
	Num       int    `json:"num"`
	UID       int64  `json:"uid"`
	Uname     string `json:"uname"`
	Price     int    `json:"price"`
	CoinType  string `json:"coin_type"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
}

// GuardMessage 是 USER_TOAST_MSG / GUARD_BUY 的主要字段集合。
type GuardMessage struct {
	UID        int64  `json:"uid"`
	Username   string `json:"username"`
	GuardLevel int    `json:"guard_level"`
	Num        int    `json:"num"`
	Price      int    `json:"price"`
	Unit       string `json:"unit"`
}

// SuperChatMessage 是 SUPER_CHAT_MESSAGE 的主要内容。
type SuperChatMessage struct {
	UID       int64  `json:"uid"`
	UserInfo  string `json:"user_info"`
	Message   string `json:"message"`
	Price     int    `json:"price"`
	Time      int    `json:"time"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	ID        int64  `json:"id"`
}

// InteractMessage 是 INTERACT_WORD / ENTRY_EFFECT / LIKE_INFO_V3_CLICK 等交互事件的主要字段集合。
type InteractMessage struct {
	UID       int64  `json:"uid"`
	Uname     string `json:"uname"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
	MsgType   int    `json:"msg_type"`
}

// Parse 把原始事件解码为常见命令的弱结构化对象。
// 字段随 B 站命令版本变化；未识别字段与该命令的完整 JSON 仍保留在 Raw 中。
func Parse(event Event) (any, error) {
	switch event.Cmd {
	case "DANMU_MSG":
		return parseDanmakuMessage(event.Raw)
	case "SEND_GIFT":
		return parseGiftMessage(event.Raw)
	case "GUARD_BUY", "USER_TOAST_MSG":
		return parseGuardMessage(event.Raw)
	case "SUPER_CHAT_MESSAGE":
		return parseSuperChatMessage(event.Raw)
	case "INTERACT_WORD", "ENTRY_EFFECT", "LIKE_INFO_V3_CLICK":
		return parseInteractMessage(event.Raw)
	default:
		return nil, fmt.Errorf("%w: unsupported command %q", errDanmakuDecodeEvent, event.Cmd)
	}
}

func parseDanmakuMessage(raw []byte) (DanmakuMessage, error) {
	var envelope struct {
		Info []json.RawMessage `json:"info"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return DanmakuMessage{}, fmt.Errorf("%w: %w", errDanmakuDecodeEvent, err)
	}
	if len(envelope.Info) < 2 {
		return DanmakuMessage{}, fmt.Errorf("%w: DANMU_MSG has too few fields", errDanmakuDecodeEvent)
	}
	var message DanmakuMessage
	if err := json.Unmarshal(envelope.Info[1], &message.Text); err != nil {
		return DanmakuMessage{}, fmt.Errorf("%w: %w", errDanmakuDecodeEvent, err)
	}
	if len(envelope.Info) > 2 {
		var user []json.RawMessage
		if json.Unmarshal(envelope.Info[2], &user) == nil && len(user) > 0 {
			_ = json.Unmarshal(user[0], &message.UID)
			if len(user) > 1 {
				_ = json.Unmarshal(user[1], &message.Uname)
			}
			if len(user) > 2 {
				_ = json.Unmarshal(user[2], &message.Admin)
			}
			if len(user) > 3 {
				_ = json.Unmarshal(user[3], &message.VIP)
			}
		}
	}
	if len(envelope.Info) > 3 {
		var medal []json.RawMessage
		if json.Unmarshal(envelope.Info[3], &medal) == nil && len(medal) > 10 {
			_ = json.Unmarshal(medal[10], &message.Guard)
		}
	}
	if len(envelope.Info) > 0 {
		var basic []json.RawMessage
		if json.Unmarshal(envelope.Info[0], &basic) == nil && len(basic) > 3 {
			_ = json.Unmarshal(basic[3], &message.Color)
			_ = json.Unmarshal(basic[1], &message.Mode)
		}
	}
	return message, nil
}

func parseGiftMessage(raw []byte) (GiftMessage, error) {
	var envelope struct {
		Data GiftMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return GiftMessage{}, fmt.Errorf("%w: %w", errDanmakuDecodeEvent, err)
	}
	return envelope.Data, nil
}

func parseGuardMessage(raw []byte) (GuardMessage, error) {
	var envelope struct {
		Data GuardMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return GuardMessage{}, fmt.Errorf("%w: %w", errDanmakuDecodeEvent, err)
	}
	return envelope.Data, nil
}

func parseSuperChatMessage(raw []byte) (SuperChatMessage, error) {
	var envelope struct {
		Data SuperChatMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return SuperChatMessage{}, fmt.Errorf("%w: %w", errDanmakuDecodeEvent, err)
	}
	return envelope.Data, nil
}

func parseInteractMessage(raw []byte) (InteractMessage, error) {
	var envelope struct {
		Data InteractMessage `json:"data"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return InteractMessage{}, fmt.Errorf("%w: %w", errDanmakuDecodeEvent, err)
	}
	return envelope.Data, nil
}
