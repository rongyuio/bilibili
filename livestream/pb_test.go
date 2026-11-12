package livestream

import (
	"encoding/binary"
	"errors"
	"math"
	"reflect"
	"strings"
	"testing"
)

// 合成 wire-format 字节的辅助函数。

func pbTag(number uint64, wire uint64) []byte {
	return pbVarint(number<<3 | wire)
}

func pbVarint(value uint64) []byte {
	buf := make([]byte, 0, 10)
	for value >= 0x80 {
		buf = append(buf, byte(value)|0x80)
		value >>= 7
	}
	return append(buf, byte(value))
}

func pbFixed64(value uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, value)
	return buf
}

func pbFixed32(value uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, value)
	return buf
}

func pbBytesTag(number uint64, value []byte) []byte {
	buf := pbTag(number, pbWireBytes)
	buf = append(buf, pbVarint(uint64(len(value)))...)
	return append(buf, value...)
}

func pbJoin(chunks ...[]byte) []byte {
	var buf []byte
	for _, chunk := range chunks {
		buf = append(buf, chunk...)
	}
	return buf
}

func TestPBValidScalars(t *testing.T) {
	data := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(300),
		pbTag(2, pbWireFixed64), pbFixed64(math.MaxUint64),
		pbTag(3, pbWireFixed32), pbFixed32(0xdeadbeef),
		pbTag(4, pbWireVarint), pbVarint(math.MaxUint64),
		pbBytesTag(5, []byte("弹幕")),
	)
	fields, err := decodePBFields(data)
	if err != nil {
		t.Fatalf("decodePBFields 返回错误: %v", err)
	}
	want := []pbField{
		{Number: 1, Wire: pbWireVarint, Varint: 300},
		{Number: 2, Wire: pbWireFixed64, Fixed64: math.MaxUint64},
		{Number: 3, Wire: pbWireFixed32, Fixed32: 0xdeadbeef},
		{Number: 4, Wire: pbWireVarint, Varint: math.MaxUint64},
		{Number: 5, Wire: pbWireBytes, Bytes: []byte("弹幕")},
	}
	if !reflect.DeepEqual(fields, want) {
		t.Fatalf("字段解析不正确\n得到: %+v\n期望: %+v", fields, want)
	}
	if got := pbValue(fields[1]); got != math.MaxUint64 {
		t.Fatalf("pbValue(fixed64) = %d", got)
	}
	if got := pbValue(fields[2]); got != uint64(0xdeadbeef) {
		t.Fatalf("pbValue(fixed32) = %d", got)
	}
	if got := pbText(fields[4]); got != "弹幕" {
		t.Fatalf("pbText = %q", got)
	}
}

func TestPBLengthDelimited(t *testing.T) {
	nested := pbJoin(pbTag(9, pbWireVarint), pbVarint(42))
	data := pbJoin(
		pbBytesTag(1, []byte{}),
		pbBytesTag(2, []byte{0x00, 0xff, 0x80, 0x00}),
		pbBytesTag(3, nested),
	)
	fields, err := decodePBFields(data)
	if err != nil {
		t.Fatalf("decodePBFields 返回错误: %v", err)
	}
	if len(fields) != 3 {
		t.Fatalf("字段数 = %d，期望 3", len(fields))
	}
	if fields[0].Wire != pbWireBytes || len(fields[0].Bytes) != 0 {
		t.Fatalf("空字节字段不正确: %+v", fields[0])
	}
	if fields[1].Wire != pbWireBytes || !reflect.DeepEqual(fields[1].Bytes, []byte{0x00, 0xff, 0x80, 0x00}) {
		t.Fatalf("二进制载荷不正确: %+v", fields[1])
	}
	if fields[2].Wire != pbWireBytes || !reflect.DeepEqual(fields[2].Bytes, nested) {
		t.Fatalf("嵌套 message 载荷不正确: %+v", fields[2])
	}
	if got := pbText(fields[0]); got != "" {
		t.Fatalf("pbText(空载荷) = %q", got)
	}
}

func TestPBValueAndText(t *testing.T) {
	varints := pbField{Wire: pbWireVarint, Varint: 123}
	fixed64 := pbField{Wire: pbWireFixed64, Fixed64: 456}
	fixed32 := pbField{Wire: pbWireFixed32, Fixed32: 789}
	bytesField := pbField{Wire: pbWireBytes, Bytes: []byte("hi")}
	group := pbField{Wire: pbWireStartGroup, Number: 1}
	unknown := pbField{Wire: 7, Varint: 999}

	if pbValue(varints) != 123 || pbValue(fixed64) != 456 || pbValue(fixed32) != 789 {
		t.Fatal("数值型字段取值不正确")
	}
	if pbValue(bytesField) != 0 || pbValue(group) != 0 || pbValue(unknown) != 0 {
		t.Fatal("非数值型字段应返回 0")
	}
	if pbText(bytesField) != "hi" {
		t.Fatalf("pbText = %q", pbText(bytesField))
	}
	for _, field := range []pbField{varints, fixed64, fixed32, group, unknown} {
		if pbText(field) != "" {
			t.Fatalf("wire %d 应返回空串", field.Wire)
		}
	}
}

func TestPBVarintOverflow(t *testing.T) {
	// 第十字节为 2，超出 uint64 可表示范围。
	tenth := []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x02}
	cases := map[string][]byte{
		"标签溢出":          append(append([]byte{}, tenth...), 0x01),
		"载荷溢出":          pbJoin(pbTag(1, pbWireVarint), tenth),
		"长度溢出":          pbJoin(pbTag(1, pbWireBytes), tenth),
		"标签字段号超 uint32": append([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}, 0x01),
	}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := decodePBFields(data); !errors.Is(err, errPBVarintOverflow) && !errors.Is(err, errPBFieldNumberTooLarge) {
				t.Fatalf("错误 = %v，期望 varint 溢出或字段号超限", err)
			}
		})
	}
}

func TestPBTruncated(t *testing.T) {
	cases := map[string][]byte{
		"标签 varint 截断":   {0x80},
		"载荷 varint 截断":   pbJoin(pbTag(1, pbWireVarint), []byte{0x80}),
		"长度 varint 截断":   pbJoin(pbTag(1, pbWireBytes), []byte{0x80}),
		"fixed64 缺 1 字节": pbJoin(pbTag(1, pbWireFixed64), make([]byte, 7)),
		"fixed32 缺 1 字节": pbJoin(pbTag(1, pbWireFixed32), make([]byte, 3)),
		"length 声明超出剩余":  pbJoin(pbTag(1, pbWireBytes), pbVarint(5), []byte("ab")),
		"length 无载荷":     pbJoin(pbTag(1, pbWireBytes), pbVarint(1)),
		"十字节 varint 截断":  {0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80},
	}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := decodePBFields(data); err == nil {
				t.Fatal("截断数据应返回错误")
			}
		})
	}
}

func TestPBFieldNumberZero(t *testing.T) {
	cases := map[string][]byte{
		"varint":      pbJoin(pbTag(0, pbWireVarint), pbVarint(1)),
		"bytes":       pbJoin(pbTag(0, pbWireBytes), pbVarint(0)),
		"fixed64":     pbJoin(pbTag(0, pbWireFixed64), make([]byte, 8)),
		"start group": pbTag(0, pbWireStartGroup),
	}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := decodePBFields(data); !errors.Is(err, errPBFieldNumberZero) {
				t.Fatalf("错误 = %v，期望 errPBFieldNumberZero", err)
			}
		})
	}
}

func TestPBUnknownWireType(t *testing.T) {
	for wire := uint64(6); wire <= 7; wire++ {
		data := pbTag(1, wire)
		if _, err := decodePBFields(data); !errors.Is(err, errPBUnknownWireType) {
			t.Fatalf("wire %d 错误 = %v，期望 errPBUnknownWireType", wire, err)
		}
	}
	if _, err := decodePBFields([]byte{0x0e}); !errors.Is(err, errPBUnknownWireType) {
		t.Fatalf("错误 = %v，期望 errPBUnknownWireType", err)
	}
}

func TestPBGroupSkipping(t *testing.T) {
	// group 5 内含一个嵌套 group 6 和若干普通字段；顶层前后各留一个字段。
	data := pbJoin(
		pbTag(1, pbWireVarint), pbVarint(11),
		pbTag(5, pbWireStartGroup),
		pbTag(2, pbWireVarint), pbVarint(22),
		pbTag(6, pbWireStartGroup),
		pbBytesTag(3, []byte("inner")),
		pbTag(6, pbWireEndGroup),
		pbTag(4, pbWireFixed32), pbFixed32(44),
		pbTag(5, pbWireEndGroup),
		pbBytesTag(7, []byte("after")),
	)
	fields, err := decodePBFields(data)
	if err != nil {
		t.Fatalf("decodePBFields 返回错误: %v", err)
	}
	if len(fields) != 2 || fields[0].Number != 1 || fields[1].Number != 7 {
		t.Fatalf("顶层字段不正确: %+v", fields)
	}
	if fields[0].Varint != 11 || pbText(fields[1]) != "after" {
		t.Fatalf("顶层字段载荷不正确: %+v", fields)
	}
}

func TestPBGroupErrors(t *testing.T) {
	cases := map[string]struct {
		data []byte
		want error
	}{
		"孤立 end group": {
			data: pbTag(1, pbWireEndGroup),
			want: errPBEndGroupWithoutStart,
		},
		"未闭合 group": {
			data: pbJoin(pbTag(1, pbWireStartGroup), pbTag(2, pbWireVarint), pbVarint(1)),
			want: errPBUnclosedGroup,
		},
		"编号不匹配": {
			data: pbJoin(pbTag(1, pbWireStartGroup), pbTag(2, pbWireEndGroup)),
			want: errPBGroupNumberMismatch,
		},
		"嵌套未闭合": {
			data: pbJoin(pbTag(1, pbWireStartGroup), pbTag(2, pbWireStartGroup), pbTag(2, pbWireEndGroup)),
			want: errPBUnclosedGroup,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := decodePBFields(tc.data); !errors.Is(err, tc.want) {
				t.Fatalf("错误 = %v，期望 %v", err, tc.want)
			}
		})
	}
}

func TestPBLengthLimits(t *testing.T) {
	if _, err := decodePBFields(make([]byte, danmakuMaxPacketSize+1)); !errors.Is(err, errPBPacketTooLarge) {
		t.Fatalf("错误 = %v，期望 errPBPacketTooLarge", err)
	}
	// 长度声明超过包上限，且载荷并未实际附上。
	data := pbJoin(pbTag(1, pbWireBytes), pbVarint(uint64(danmakuMaxPacketSize)+1))
	if _, err := decodePBFields(data); !errors.Is(err, errPBBytesTooLarge) {
		t.Fatalf("错误 = %v，期望 errPBBytesTooLarge", err)
	}
	// 长度声明等于上限，但剩余字节不足。
	data = pbJoin(pbTag(1, pbWireBytes), pbVarint(uint64(danmakuMaxPacketSize)), []byte("x"))
	if _, err := decodePBFields(data); !errors.Is(err, errPBTruncatedBytes) {
		t.Fatalf("错误 = %v，期望 errPBTruncatedBytes", err)
	}
}

func TestPBTooManyFields(t *testing.T) {
	one := pbJoin(pbTag(1, pbWireVarint), pbVarint(0))
	data := []byte(strings.Repeat(string(one), pbMaxFieldCount+1))
	if _, err := decodePBFields(data); !errors.Is(err, errPBTooManyFields) {
		t.Fatalf("错误 = %v，期望 errPBTooManyFields", err)
	}
	data = []byte(strings.Repeat(string(one), pbMaxFieldCount))
	fields, err := decodePBFields(data)
	if err != nil {
		t.Fatalf("恰好等于上限时返回错误: %v", err)
	}
	if len(fields) != pbMaxFieldCount {
		t.Fatalf("字段数 = %d，期望 %d", len(fields), pbMaxFieldCount)
	}
}

func FuzzDecodePBFields(f *testing.F) {
	f.Add([]byte{})
	f.Add(pbJoin(pbTag(1, pbWireVarint), pbVarint(300)))
	f.Add(pbBytesTag(2, []byte("弹幕")))
	f.Add(pbJoin(pbTag(3, pbWireStartGroup), pbTag(4, pbWireVarint), pbVarint(1), pbTag(3, pbWireEndGroup)))
	f.Add([]byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x02})
	f.Add(pbTag(1, 7))
	f.Fuzz(func(t *testing.T, data []byte) {
		fields, err := decodePBFields(data)
		if err != nil {
			return
		}
		if len(fields) > pbMaxFieldCount {
			t.Fatalf("字段数超出上限: %d", len(fields))
		}
		for _, field := range fields {
			_ = pbValue(field)
			_ = pbText(field)
			if field.Wire == pbWireBytes && uint64(len(field.Bytes)) > uint64(len(data)) {
				t.Fatalf("length-delimited 载荷超出输入长度: %d > %d", len(field.Bytes), len(data))
			}
		}
	})
}
