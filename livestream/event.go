package livestream

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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
// V2 命令返回 protobuf 解析结果；未识别字段与该命令的完整 JSON 仍保留在 Raw 中。
func Parse(event Event) (any, error) {
	cmd := commandName(event.Cmd)
	switch cmd {
	case "DANMU_MSG", "DANMU_MSG_MIRROR":
		return parseDanmakuMessage(event.Raw)
	case "SEND_GIFT":
		return parseGiftMessage(event.Raw)
	case "SEND_GIFT_V2":
		return ParseSendGiftV2(event.Raw)
	case "GUARD_BUY", "USER_TOAST_MSG", "USER_TOAST_MSG_V2":
		return parseGuardMessage(event.Raw)
	case "SUPER_CHAT_MESSAGE":
		return parseSuperChatMessage(event.Raw)
	case "INTERACT_WORD", "ENTRY_EFFECT", "LIKE_INFO_V3_CLICK":
		return parseInteractMessage(event.Raw)
	case "INTERACT_WORD_V2":
		return ParseInteractWordV2(event.Raw)
	case "ONLINE_RANK_V2", "ONLINE_RANK_V3":
		return ParseOnlineRankPB(event.Raw)
	case "LIVE", "PREPARING":
		return event.Raw, nil
	default:
		return nil, fmt.Errorf("%w: unsupported command %q", errDanmakuDecodeEvent, event.Cmd)
	}
}

// commandName 去掉 DANMU_MSG:4:0:2:2:2:0 这类命令后缀参数。
func commandName(cmd string) string {
	name, _, _ := strings.Cut(cmd, ":")
	return name
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
	user := parseDanmakuUser(envelope.Info)
	message.UID, message.Uname, message.Admin, message.VIP = user.UID, user.Uname, user.Admin, user.VIP
	message.Guard = parseDanmakuGuard(envelope.Info)
	style := parseDanmakuStyle(envelope.Info)
	message.Color, message.Mode = style.Color, style.Mode
	// info[0][15].user 是网页端展开 dm_v2 后的用户信息，比 info[2] 更完整。
	// 只在旧字段缺失时补齐，避免覆盖 JSON 主结构已有的值。
	if message.UID == 0 {
		uid, name := parseDanmakuNestedUser(envelope.Info)
		message.UID = uid
		if message.Uname == "" {
			message.Uname = name
		}
	}
	return message, nil
}

type danmakuJSONUser struct {
	UID   int64
	Uname string
	Admin bool
	VIP   bool
}

func parseDanmakuUser(info []json.RawMessage) danmakuJSONUser {
	var user []json.RawMessage
	if len(info) < 3 || json.Unmarshal(info[2], &user) != nil {
		return danmakuJSONUser{}
	}
	var uid int64
	var uname string
	var admin, vip bool
	assign := func(index int, dst any) {
		if len(user) > index {
			_ = json.Unmarshal(user[index], dst)
		}
	}
	assign(0, &uid)
	assign(1, &uname)
	assign(2, &admin)
	assign(3, &vip)
	return danmakuJSONUser{UID: uid, Uname: uname, Admin: admin, VIP: vip}
}

func parseDanmakuGuard(info []json.RawMessage) int {
	var medal []json.RawMessage
	if len(info) < 4 || json.Unmarshal(info[3], &medal) != nil {
		return 0
	}
	var guard int
	if len(medal) > 10 {
		_ = json.Unmarshal(medal[10], &guard)
	}
	return guard
}

type danmakuJSONStyle struct {
	Color int
	Mode  int
}

func parseDanmakuStyle(info []json.RawMessage) (style danmakuJSONStyle) {
	var basic []json.RawMessage
	if len(info) == 0 || json.Unmarshal(info[0], &basic) != nil || len(basic) <= 3 {
		return style
	}
	_ = json.Unmarshal(basic[3], &style.Color)
	_ = json.Unmarshal(basic[1], &style.Mode)
	return style
}

func parseDanmakuNestedUser(info []json.RawMessage) (int64, string) {
	var basic []json.RawMessage
	if len(info) == 0 || json.Unmarshal(info[0], &basic) != nil || len(basic) <= 15 {
		return 0, ""
	}
	var extra struct {
		User struct {
			UID  int64 `json:"uid"`
			Base struct {
				Name string `json:"name"`
			} `json:"base"`
		} `json:"user"`
	}
	if json.Unmarshal(basic[15], &extra) != nil {
		return 0, ""
	}
	return extra.User.UID, extra.User.Base.Name
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
