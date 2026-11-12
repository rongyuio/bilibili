package livestream

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/andybalholm/brotli"
)

// 直播弹幕信息流的二进制帧编解码。

var (
	errDanmakuPacketTooLarge = errors.New("live danmaku packet exceeds size limit")
	errDanmakuPacketTooShort = errors.New("live danmaku packet is shorter than its header")
	errDanmakuFrameTooShort  = errors.New("live danmaku frame is shorter than its declared length")
	errDanmakuUnknownProto   = errors.New("live danmaku packet uses an unsupported protocol version")
)

// danmakuPacket 是信息流上的一帧。
type danmakuPacket struct {
	Length    uint32
	Header    uint16
	Protover  uint16
	Operation uint32
	Sequence  uint32
	Body      []byte
}

// encodeDanmakuPacket 把业务体打包成发送帧。
func encodeDanmakuPacket(protover uint16, operation uint32, body []byte) []byte {
	packet := make([]byte, danmakuHeaderSize+len(body))
	binary.BigEndian.PutUint32(packet[0:4], uint32(len(packet)&0xffffffff))
	binary.BigEndian.PutUint16(packet[4:6], danmakuHeaderSize)
	binary.BigEndian.PutUint16(packet[6:8], protover)
	binary.BigEndian.PutUint32(packet[8:12], operation)
	binary.BigEndian.PutUint32(packet[12:16], danmakuSequence)
	copy(packet[danmakuHeaderSize:], body)
	return packet
}

// decodeDanmakuPackets 解析一帧，并展开压缩帧内的子帧。
func decodeDanmakuPackets(data []byte) ([]danmakuPacket, error) {
	packet, err := decodeDanmakuPacket(data)
	if err != nil {
		return nil, err
	}
	switch packet.Protover {
	case danmakuProtoPlain, danmakuProtoHeartbeat:
		return []danmakuPacket{packet}, nil
	case danmakuProtoZlib:
		body, err := decompressZlib(packet.Body)
		if err != nil {
			return nil, fmt.Errorf("decompress live danmaku zlib packet: %w", err)
		}
		return decodeDanmakuPacketStream(body)
	case danmakuProtoBrotli:
		body, err := readAllLimited(brotli.NewReader(bytes.NewReader(packet.Body)))
		if err != nil {
			return nil, fmt.Errorf("decompress live danmaku brotli packet: %w", err)
		}
		return decodeDanmakuPacketStream(body)
	default:
		return nil, fmt.Errorf("%w: %d", errDanmakuUnknownProto, packet.Protover)
	}
}

func decompressZlib(body []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer func() { _ = reader.Close() }()
	return readAllLimited(reader)
}

// decodeDanmakuPacketStream 顺序解析压缩体内连续的二进制帧。
func decodeDanmakuPacketStream(data []byte) ([]danmakuPacket, error) {
	var packets []danmakuPacket
	for len(data) > 0 {
		if len(data) < danmakuHeaderSize {
			return nil, errDanmakuFrameTooShort
		}
		length := int64(binary.BigEndian.Uint32(data[0:4]))
		if length < danmakuHeaderSize || length > int64(len(data)) {
			return nil, errDanmakuFrameTooShort
		}
		packet, err := decodeDanmakuPacket(data[:length])
		if err != nil {
			return nil, err
		}
		if packet.Protover != danmakuProtoPlain && packet.Protover != danmakuProtoHeartbeat {
			return nil, fmt.Errorf("%w: nested %d", errDanmakuUnknownProto, packet.Protover)
		}
		packets = append(packets, packet)
		data = data[length:]
	}
	return packets, nil
}

func decodeDanmakuPacket(data []byte) (danmakuPacket, error) {
	if len(data) < danmakuHeaderSize {
		return danmakuPacket{}, errDanmakuPacketTooShort
	}
	length := int64(binary.BigEndian.Uint32(data[0:4]))
	if length < danmakuHeaderSize || length > danmakuMaxPacketSize {
		return danmakuPacket{}, errDanmakuPacketTooLarge
	}
	if length > int64(len(data)) {
		return danmakuPacket{}, errDanmakuFrameTooShort
	}
	header := int64(binary.BigEndian.Uint16(data[4:6]))
	if header < danmakuHeaderSize || header > length {
		return danmakuPacket{}, errDanmakuFrameTooShort
	}
	return danmakuPacket{
		Length:    uint32(length),
		Header:    uint16(header & 0xffff),
		Protover:  binary.BigEndian.Uint16(data[6:8]),
		Operation: binary.BigEndian.Uint32(data[8:12]),
		Sequence:  binary.BigEndian.Uint32(data[12:16]),
		Body:      data[header:length],
	}, nil
}

// danmakuCommand 只读取消息体中的 cmd，未知命令仍以 Raw 事件交给调用方。
type danmakuCommand struct {
	Cmd string `json:"cmd"`
}

func danmakuCommandOf(body []byte) string {
	var command danmakuCommand
	if json.Unmarshal(body, &command) != nil {
		return ""
	}
	return commandName(command.Cmd)
}
