package livestream

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestParseDanmakuMessage(t *testing.T) {
	raw := json.RawMessage(`{"cmd":"DANMU_MSG","info":[[0,1,25,16777215],"打卡",[42,"ry",0,1], [0,0,0,0,0,0,0,0,0,0,3]]}`)
	parsed, err := Parse(Event{Cmd: "DANMU_MSG", Raw: raw})
	if err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	message, ok := parsed.(DanmakuMessage)
	if !ok {
		t.Fatalf("Parse 返回类型不正确: %T", parsed)
	}
	if message.Text != "打卡" || message.UID != 42 || message.Uname != "ry" ||
		message.Color != 16777215 || message.Mode != 1 || message.Guard != 3 {
		t.Fatalf("弹幕解析结果不正确: %+v", message)
	}
}

func TestParseGiftMessage(t *testing.T) {
	raw := json.RawMessage(`{"cmd":"SEND_GIFT","data":{"giftName":"辣条","giftId":1,"num":2,"uid":42,"uname":"ry","price":100,"coin_type":"silver","action":"赠送","timestamp":123}}`)
	parsed, err := Parse(Event{Cmd: "SEND_GIFT", Raw: raw})
	if err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	gift, ok := parsed.(GiftMessage)
	if !ok {
		t.Fatalf("Parse 返回类型不正确: %T", parsed)
	}
	if gift.GiftName != "辣条" || gift.Num != 2 || gift.UID != 42 {
		t.Fatalf("礼物解析结果不正确: %+v", gift)
	}
}

func TestParseUnsupportedCommand(t *testing.T) {
	if _, err := Parse(Event{Cmd: "UNKNOWN", Raw: json.RawMessage(`{}`)}); !errors.Is(err, errDanmakuDecodeEvent) {
		t.Fatalf("未知命令应返回可判定的错误: %v", err)
	}
}
