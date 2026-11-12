package livestream

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"testing"
)

// 测试用 protobuf 字节按 wire format 手工拼接，字段号与 events.proto 一致。

func TestParseSendGiftV2(t *testing.T) {
	item := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(31039),
		pbBytesTag(2, []byte("小花花")),
		pbTag(3, pbWireVarint), pbVarint(2),
		pbTag(5, pbWireVarint), pbVarint(1000),
		pbTag(7, pbWireVarint), pbVarint(2000),
		pbBytesTag(8, []byte("gold")),
		pbBytesTag(9, []byte("txn-1")),
		pbTag(10, pbWireVarint), pbVarint(1700000000),
		pbBytesTag(18, []byte("投喂")),
		pbTag(36, pbWireVarint), pbVarint(2500),
	)
	item2 := pbJoin(pbTag(1, pbWireVarint), pbVarint(999))
	base := pbBytesTag(1, []byte("嵌套用户"))
	base = append(base, pbBytesTag(2, []byte("https://example.invalid/face.jpg"))...)
	sender := pbJoin(pbTag(1, pbWireVarint), pbVarint(42), pbBytesTag(2, base))
	blind := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(139),
		pbTag(2, pbWireVarint), pbVarint(200),
		pbBytesTag(3, []byte("心动盲盒")),
		pbTag(6, pbWireVarint), pbVarint(2000),
	)
	message := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(7),
		pbBytesTag(2, []byte("顶层用户")),
		pbBytesTag(9, blind),
		pbBytesTag(10, item),
		pbBytesTag(10, item2),
		pbBytesTag(15, sender),
	)
	raw := eventPBJSON("SEND_GIFT_V2", message)
	event, err := ParseSendGiftV2(raw)
	if err != nil {
		t.Fatalf("ParseSendGiftV2 返回错误: %v", err)
	}
	if len(event.GiftItems) != 2 {
		t.Fatalf("礼物结果数 = %d，期望 2", len(event.GiftItems))
	}
	first := event.GiftItems[0]
	if first.GiftID != 31039 || first.GiftName != "小花花" || first.Num != 2 ||
		first.TotalCoin != 2000 || first.CoinType != "gold" || first.TID != "txn-1" ||
		first.GiftTipPrice != 2500 {
		t.Fatalf("礼物结果不正确: %+v", first)
	}
	if event.GiftItems[1].GiftID != 999 {
		t.Fatalf("第二个礼物结果不正确: %+v", event.GiftItems[1])
	}
	if event.SenderUInfo == nil || event.SenderUInfo.UID != 42 ||
		event.SenderUInfo.Name != "嵌套用户" || event.SenderUInfo.Face != "https://example.invalid/face.jpg" {
		t.Fatalf("嵌套发送者不正确: %+v", event.SenderUInfo)
	}
	if event.BlindGift == nil || event.BlindGift.ConfigID != 139 ||
		event.BlindGift.OriginalGiftName != "心动盲盒" {
		t.Fatalf("盲盒信息不正确: %+v", event.BlindGift)
	}
	parsed, err := Parse(Event{Cmd: "SEND_GIFT_V2", Raw: raw})
	if err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if _, ok := parsed.(PBSendGift); !ok {
		t.Fatalf("Parse 返回类型不正确: %T", parsed)
	}
	if err := parseSendGiftV2Error(t, pbJoin(pbTag(1, pbWireVarint), pbVarint(1))); !errors.Is(err, errDanmakuPBEventEmpty) {
		t.Fatalf("无礼物结果应返回空事件错误，实际为 %v", err)
	}
}

func parseSendGiftV2Error(t *testing.T, data []byte) error {
	t.Helper()
	_, err := ParseSendGiftV2(eventPBJSON("SEND_GIFT_V2", data))
	return err
}

func TestParseInteractWordV2(t *testing.T) {
	base := pbJoin(
		pbBytesTag(1, []byte("访客")),
		pbBytesTag(2, []byte("https://example.invalid/face.png")),
	)
	uinfo := pbJoin(pbTag(1, pbWireVarint), pbVarint(9000000000), pbBytesTag(2, base))
	data := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(9000000000),
		pbBytesTag(2, []byte("老用户")),
		pbTag(5, pbWireVarint), pbVarint(1),
		pbTag(6, pbWireVarint), pbVarint(263605),
		pbTag(7, pbWireVarint), pbVarint(1700000000),
		pbTag(8, pbWireVarint), pbVarint(1700000000123),
		pbTag(15, pbWireVarint), pbVarint(1700000000000000000),
		pbTag(16, pbWireVarint), pbVarint(3),
		pbBytesTag(22, uinfo),
	)
	event, err := ParseInteractWordV2(eventPBJSON("INTERACT_WORD_V2", data))
	if err != nil {
		t.Fatalf("ParseInteractWordV2 返回错误: %v", err)
	}
	if event.UID != 9000000000 || event.Username != "老用户" || event.MsgType != 1 ||
		event.RoomID != 263605 || event.Timestamp != 1700000000 ||
		event.TimestampMS != 1700000000123 || event.GuardType != 3 {
		t.Fatalf("解析结果不正确: %+v", event)
	}
	if event.UInfo == nil || event.UInfo.UID != 9000000000 ||
		event.UInfo.Name != "访客" || event.UInfo.Face != "https://example.invalid/face.png" {
		t.Fatalf("嵌套用户信息不正确: %+v", event.UInfo)
	}
}

func TestParseOnlineRankPB(t *testing.T) {
	user := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(42),
		pbBytesTag(2, []byte("https://example.invalid/face.webp")),
		pbBytesTag(3, []byte("100")),
		pbBytesTag(4, []byte("榜一")),
		pbTag(5, pbWireVarint), pbVarint(1),
		pbTag(6, pbWireVarint), pbVarint(2),
		pbTag(7, pbWireVarint), pbVarint(1),
	)
	data := pbJoin(
		pbBytesTag(1, []byte("online_rank")),
		pbBytesTag(3, user),
	)
	event, err := ParseOnlineRankPB(eventPBJSON("ONLINE_RANK_V3", data))
	if err != nil {
		t.Fatalf("ParseOnlineRankPB 返回错误: %v", err)
	}
	if event.RankType != "online_rank" || len(event.Users) != 1 {
		t.Fatalf("榜单解析结果不正确: %+v", event)
	}
	first := event.Users[0]
	if first.UID != 42 || first.Uname != "榜一" || first.Rank != 1 ||
		first.GuardLevel == nil || *first.GuardLevel != 2 ||
		first.IsMystery == nil || !*first.IsMystery {
		t.Fatalf("榜单用户不正确: %+v", first)
	}
}

func TestParseDanmakuPB(t *testing.T) {
	emoticon := pbJoin(
		pbBytesTag(1, []byte("emo-1")),
		pbBytesTag(2, []byte("https://example.invalid/emo.png")),
		pbTag(3, pbWireVarint), pbVarint(1),
		pbTag(6, pbWireVarint), pbVarint(32),
		pbTag(7, pbWireVarint), pbVarint(32),
	)
	entry := pbJoin(pbBytesTag(1, []byte("[微笑]")), pbBytesTag(2, emoticon))
	user := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(9000000000),
		pbBytesTag(2, []byte("ry")),
		pbBytesTag(4, []byte("https://example.invalid/face.jpg")),
		pbBytesTag(11, pbJoin(pbTag(1, pbWireVarint), pbVarint(20), pbBytesTag(2, []byte("粉丝牌")))),
		pbBytesTag(15, pbJoin(pbTag(1, pbWireVarint), pbVarint(30))),
	)
	data := pbJoin(
		pbBytesTag(1, []byte("dm-1")),
		pbTag(2, pbWireVarint), pbVarint(1),
		pbTag(4, pbWireVarint), pbVarint(16777215),
		pbBytesTag(6, []byte("打卡")),
		pbTag(13, pbWireVarint), pbVarint(1),
		pbBytesTag(14, entry),
		pbBytesTag(20, user),
		pbBytesTag(21, pbJoin(pbTag(1, pbWireVarint), pbVarint(5055), pbBytesTag(2, []byte("直播间")))),
	)
	raw, err := json.Marshal(map[string]any{
		"cmd":   "DANMU_MSG:4:0:2:2:2:0",
		"dm_v2": base64.StdEncoding.EncodeToString(data),
	})
	if err != nil {
		t.Fatalf("构造测试 JSON 失败: %v", err)
	}
	event, err := ParseDanmakuPB(raw)
	if err != nil {
		t.Fatalf("ParseDanmakuPB 返回错误: %v", err)
	}
	if event.IDStr != "dm-1" || event.Content != "打卡" || event.DMType != 1 ||
		event.Color != 16777215 || event.Emoticons["[微笑]"].URL != "https://example.invalid/emo.png" {
		t.Fatalf("dm_v2 解析结果不正确: %+v", event)
	}
	if event.User == nil || event.User.UID != 9000000000 || event.User.Name != "ry" ||
		event.User.Medal == nil || event.User.Medal.Level != 20 ||
		event.User.Wealth == nil || event.User.Wealth.Level != 30 {
		t.Fatalf("dm_v2 用户不正确: %+v", event.User)
	}
	if event.Room == nil || event.Room.UID != 5055 || event.Room.Name != "直播间" {
		t.Fatalf("dm_v2 房间不正确: %+v", event.Room)
	}
}

func TestParsePBMissingPayload(t *testing.T) {
	if _, err := ParseInteractWordV2([]byte(`{"cmd":"INTERACT_WORD_V2","data":{}}`)); !errors.Is(err, errDanmakuPBPayloadMissing) {
		t.Fatalf("缺少 pb 字段的错误 = %v", err)
	}
	if _, err := ParseDanmakuPB([]byte(`{"cmd":"DANMU_MSG"}`)); !errors.Is(err, errDanmakuPBPayloadMissing) {
		t.Fatalf("缺少 dm_v2 字段的错误 = %v", err)
	}
	if _, err := ParseInteractWordV2([]byte(`{"cmd":"INTERACT_WORD_V2","data":{"pb":"!!!"}}`)); !errors.Is(err, errDanmakuPBBadBase64) {
		t.Fatalf("非法 base64 的错误 = %v", err)
	}
}

func TestParseCommandWithSuffix(t *testing.T) {
	raw := []byte(`{"cmd":"DANMU_MSG:4:0:2:2:2:0","info":[[0,1,25,16777215],"打卡",[42,"ry"]]}`)
	parsed, err := Parse(Event{Cmd: "DANMU_MSG:4:0:2:2:2:0", Raw: raw})
	if err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	message, ok := parsed.(DanmakuMessage)
	if !ok || message.Text != "打卡" {
		t.Fatalf("带后缀命令解析结果不正确: %+v", parsed)
	}
}

func eventPBJSON(cmd string, data []byte) []byte {
	raw, _ := json.Marshal(map[string]any{
		"cmd":  cmd,
		"data": map[string]any{"pb": base64.StdEncoding.EncodeToString(data)},
	})
	return raw
}
