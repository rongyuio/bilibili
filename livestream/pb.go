package livestream

import (
	"encoding/binary"
	"errors"
)

// 手写 protobuf wire-format 解码，供直播弹幕 V2 事件解析使用。
// 只还原字段号与载荷，不依赖 schema，也不引入 protobuf 运行时。

// wire 类型与解码上限。
const (
	pbWireVarint     = 0
	pbWireFixed64    = 1
	pbWireBytes      = 2
	pbWireStartGroup = 3
	pbWireEndGroup   = 4
	pbWireFixed32    = 5

	pbMaxVarintBytes = 10        // uint64 的 varint 最多 10 字节
	pbMaxFieldNumber = 1<<32 - 1 // 字段号必须能放进 uint32
	pbMaxFieldCount  = 65536     // 单条 message 的顶层字段数上限
)

var (
	errPBPacketTooLarge       = errors.New("protobuf message exceeds size limit")
	errPBBytesTooLarge        = errors.New("protobuf length-delimited value exceeds size limit")
	errPBTooManyFields        = errors.New("protobuf message has too many fields")
	errPBFieldNumberZero      = errors.New("protobuf field number is zero")
	errPBFieldNumberTooLarge  = errors.New("protobuf field number exceeds uint32")
	errPBVarintTruncated      = errors.New("protobuf varint is truncated")
	errPBVarintOverflow       = errors.New("protobuf varint overflows uint64")
	errPBTruncatedFixed64     = errors.New("protobuf fixed64 value is truncated")
	errPBTruncatedFixed32     = errors.New("protobuf fixed32 value is truncated")
	errPBTruncatedBytes       = errors.New("protobuf length-delimited value is truncated")
	errPBUnknownWireType      = errors.New("protobuf wire type is unknown")
	errPBEndGroupWithoutStart = errors.New("protobuf end group has no matching start group")
	errPBGroupNumberMismatch  = errors.New("protobuf end group field number does not match start group")
	errPBUnclosedGroup        = errors.New("protobuf start group is not closed")
)

// pbField 是一条 message 的顶层字段。按 wire 类型只填写对应的载荷字段。
type pbField struct {
	Number  uint32
	Wire    uint8
	Varint  uint64
	Fixed64 uint64
	Bytes   []byte
	Fixed32 uint32
}

// decodePBFields 解析一条 message 的顶层字段。
// wire 3/4 构成的 group 整组跳过，组内字段不会出现在结果里。
func decodePBFields(data []byte) ([]pbField, error) {
	if len(data) > danmakuMaxPacketSize {
		return nil, errPBPacketTooLarge
	}
	fields := make([]pbField, 0, 8)
	groups := make([]uint32, 0, 4)
	for pos := 0; pos < len(data); {
		field, next, err := decodePBField(data, pos)
		if err != nil {
			return nil, err
		}
		pos = next
		switch field.Wire {
		case pbWireStartGroup:
			groups = append(groups, field.Number)
			continue
		case pbWireEndGroup:
			if len(groups) == 0 {
				return nil, errPBEndGroupWithoutStart
			}
			last := groups[len(groups)-1]
			groups = groups[:len(groups)-1]
			if last != field.Number {
				return nil, errPBGroupNumberMismatch
			}
			continue
		}
		if len(groups) > 0 {
			continue
		}
		if len(fields) >= pbMaxFieldCount {
			return nil, errPBTooManyFields
		}
		fields = append(fields, field)
	}
	if len(groups) > 0 {
		return nil, errPBUnclosedGroup
	}
	return fields, nil
}

// decodePBField 从 pos 解析一个字段，返回字段与下一个读取位置。
// group 起始/结束标记不消耗载荷，只把字段号带回给调用方配对。
func decodePBField(data []byte, pos int) (pbField, int, error) {
	tag, next, err := decodePBVarint(data, pos)
	if err != nil {
		return pbField{}, pos, err
	}
	pos = next
	number := tag >> 3
	wire := tag & 7
	if number == 0 {
		return pbField{}, pos, errPBFieldNumberZero
	}
	if number > pbMaxFieldNumber {
		return pbField{}, pos, errPBFieldNumberTooLarge
	}
	field := pbField{Number: uint32(number)}
	switch wire {
	case pbWireVarint:
		field.Wire = uint8(pbWireVarint)
		value, end, err := decodePBVarint(data, pos)
		if err != nil {
			return pbField{}, pos, err
		}
		field.Varint = value
		return field, end, nil
	case pbWireFixed64:
		field.Wire = uint8(pbWireFixed64)
		if len(data)-pos < 8 {
			return pbField{}, pos, errPBTruncatedFixed64
		}
		field.Fixed64 = binary.LittleEndian.Uint64(data[pos:])
		return field, pos + 8, nil
	case pbWireBytes:
		field.Wire = uint8(pbWireBytes)
		length, end, err := decodePBVarint(data, pos)
		if err != nil {
			return pbField{}, pos, err
		}
		if length > uint64(danmakuMaxPacketSize) {
			return pbField{}, end, errPBBytesTooLarge
		}
		rest := data[end:]
		if length > uint64(len(rest)) {
			return pbField{}, end, errPBTruncatedBytes
		}
		chunk := rest[:length]
		field.Bytes = chunk
		return field, end + len(chunk), nil
	case pbWireStartGroup:
		field.Wire = uint8(pbWireStartGroup)
		return field, pos, nil
	case pbWireEndGroup:
		field.Wire = uint8(pbWireEndGroup)
		return field, pos, nil
	case pbWireFixed32:
		field.Wire = uint8(pbWireFixed32)
		if len(data)-pos < 4 {
			return pbField{}, pos, errPBTruncatedFixed32
		}
		field.Fixed32 = binary.LittleEndian.Uint32(data[pos:])
		return field, pos + 4, nil
	default:
		return pbField{}, pos, errPBUnknownWireType
	}
}

// decodePBVarint 读取一个 varint。第十字节只允许 0 或 1，否则视为溢出。
func decodePBVarint(data []byte, pos int) (uint64, int, error) {
	var value uint64
	for i := range pbMaxVarintBytes - 1 {
		if pos >= len(data) {
			return 0, pos, errPBVarintTruncated
		}
		octet := data[pos]
		pos++
		value |= uint64(octet&0x7f) << (7 * i)
		if octet&0x80 == 0 {
			return value, pos, nil
		}
	}
	if pos >= len(data) {
		return 0, pos, errPBVarintTruncated
	}
	octet := data[pos]
	if octet > 1 {
		return 0, pos, errPBVarintOverflow
	}
	return value | uint64(octet)<<63, pos + 1, nil
}

// pbValue 返回数值型字段的值；其它 wire 类型返回 0。
func pbValue(f pbField) uint64 {
	switch f.Wire {
	case pbWireVarint:
		return f.Varint
	case pbWireFixed64:
		return f.Fixed64
	case pbWireFixed32:
		return uint64(f.Fixed32)
	default:
		return 0
	}
}

// pbText 返回 length-delimited 字段的文本；其它 wire 类型返回空串。
func pbText(f pbField) string {
	if f.Wire != pbWireBytes {
		return ""
	}
	return string(f.Bytes)
}
