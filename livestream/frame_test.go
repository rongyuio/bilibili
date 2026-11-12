package livestream

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"testing"

	"github.com/andybalholm/brotli"
)

func TestEncodeDecodeDanmakuPacket(t *testing.T) {
	packet := encodeDanmakuPacket(danmakuProtoPlain, danmakuOpMessage, []byte(`{"cmd":"DANMU_MSG"}`))
	decoded, err := decodeDanmakuPacket(packet)
	if err != nil {
		t.Fatalf("decodeDanmakuPacket 返回错误: %v", err)
	}
	if decoded.Length != uint32(len(packet)) || decoded.Protover != danmakuProtoPlain ||
		decoded.Operation != danmakuOpMessage || danmakuCommandOf(decoded.Body) != "DANMU_MSG" {
		t.Fatalf("帧解析结果不正确: %+v", decoded)
	}
}

func TestDecodeDanmakuZlibPackets(t *testing.T) {
	inner := encodeDanmakuPacket(danmakuProtoPlain, danmakuOpMessage, []byte(`{"cmd":"DANMU_MSG","data":1}`))
	inner = append(inner, encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpHeartbeatReply, []byte(`[1]`))...)
	var compressed bytes.Buffer
	writer := zlib.NewWriter(&compressed)
	if _, err := writer.Write(inner); err != nil {
		t.Fatalf("压缩测试帧失败: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("关闭 zlib writer 失败: %v", err)
	}
	outer := encodeDanmakuPacket(danmakuProtoZlib, danmakuOpMessage, compressed.Bytes())
	packets, err := decodeDanmakuPackets(outer)
	if err != nil {
		t.Fatalf("decodeDanmakuPackets 返回错误: %v", err)
	}
	if len(packets) != 2 || danmakuCommandOf(packets[0].Body) != "DANMU_MSG" ||
		packets[1].Operation != danmakuOpHeartbeatReply {
		t.Fatalf("zlib 帧展开结果不正确: %+v", packets)
	}
}

func TestDecodeDanmakuBrotliPackets(t *testing.T) {
	inner := encodeDanmakuPacket(danmakuProtoPlain, danmakuOpMessage, []byte(`{"cmd":"SEND_GIFT"}`))
	var compressed bytes.Buffer
	writer := brotli.NewWriter(&compressed)
	if _, err := writer.Write(inner); err != nil {
		t.Fatalf("压缩测试帧失败: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("关闭 brotli writer 失败: %v", err)
	}
	outer := encodeDanmakuPacket(danmakuProtoBrotli, danmakuOpMessage, compressed.Bytes())
	packets, err := decodeDanmakuPackets(outer)
	if err != nil {
		t.Fatalf("decodeDanmakuPackets 返回错误: %v", err)
	}
	if len(packets) != 1 || danmakuCommandOf(packets[0].Body) != "SEND_GIFT" {
		t.Fatalf("brotli 帧展开结果不正确: %+v", packets)
	}
}

func TestDecodeDanmakuPacketRejectsInvalidLength(t *testing.T) {
	packet := make([]byte, danmakuHeaderSize)
	binary.BigEndian.PutUint32(packet[0:4], uint32(danmakuMaxPacketSize+1))
	if _, err := decodeDanmakuPacket(packet); err == nil {
		t.Fatal("超过大小上限的帧应返回错误")
	}
	short := make([]byte, danmakuHeaderSize)
	binary.BigEndian.PutUint32(short[0:4], 8)
	if _, err := decodeDanmakuPacket(short); err == nil {
		t.Fatal("包长小于帧头的帧应返回错误")
	}
}

func TestDanmakuAuthBody(t *testing.T) {
	packet, err := authPacket(danmakuConnectionInfo{
		RoomID:   23058,
		UID:      42,
		Buvid3:   "buvid",
		Token:    "token",
		Protover: danmakuProtoZlib,
	})
	if err != nil {
		t.Fatalf("authPacket 返回错误: %v", err)
	}
	decoded, err := decodeDanmakuPacket(packet)
	if err != nil {
		t.Fatalf("认证包不是合法帧: %v", err)
	}
	if decoded.Operation != danmakuOpAuth || decoded.Protover != danmakuProtoHeartbeat {
		t.Fatalf("认证包操作码或版本不正确: %+v", decoded)
	}
	var body danmakuAuthBody
	if err := json.Unmarshal(decoded.Body, &body); err != nil {
		t.Fatalf("认证包不是 JSON: %v", err)
	}
	if body.RoomID != 23058 || body.UID != 42 || body.Buvid != "buvid" ||
		body.Key != "token" || body.Protover != danmakuProtoZlib || body.Platform != "web" {
		t.Fatalf("认证包字段不正确: %+v", body)
	}
}
