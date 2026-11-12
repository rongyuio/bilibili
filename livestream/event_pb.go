package livestream

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode/utf8"
)

// 直播弹幕 V2 事件的 protobuf 模型。
//
// 字段号来自 Akegarasu/blivedm-go 的 pb/events.proto（MIT 授权），其来源为
// Akegarasu/danmuji-next 的手写 wire-format 解码器。这里不生成 protobuf 代码，
// 只按字段号读取上游 base64 载荷；64 位 uid 全程用 uint64 保存，不截断。
// 不同命令的消息定义存在同号不同含义，按命令分别解码，不共用一套字段映射。

var (
	errDanmakuPBEventEmpty     = errors.New("live danmaku protobuf event has no fields")
	errDanmakuPBEventMalformed = errors.New("live danmaku protobuf event is malformed")
	errDanmakuPBValueTooLarge  = errors.New("live danmaku protobuf integer exceeds model range")
	errDanmakuPBPayloadMissing = errors.New("live danmaku protobuf payload is missing")
	errDanmakuPBBadBase64      = errors.New("live danmaku protobuf payload has invalid base64")
	errDanmakuPBBadText        = errors.New("live danmaku protobuf text is not valid utf-8")
)

// PBUser 是 V2 事件中的用户信息。不同来源填写的字段集合不同。
type PBUser struct {
	UID       uint64
	Name      string
	Face      string
	NameColor uint32
	IsMystery bool
	Medal     *PBMedal
	Wealth    *PBWealth
	Guard     *PBGuard
}

// PBMedal 是 V2 事件中的粉丝勋章。不同来源填写的字段集合不同。
type PBMedal struct {
	RUID      uint64
	Name      string
	Level     uint32
	Color     uint32
	ColorV2   uint32
	IsLighted *bool
	GuardLv   uint32
	RoomID    uint32
}

// PBWealth 是 V2 事件中的荣耀等级。
type PBWealth struct{ Level uint32 }

// PBGuard 是 V2 事件中的舰队信息。
type PBGuard struct {
	Level      uint32
	ExpiredStr string
}

// PBSendGift 是 SEND_GIFT_V2 的 protobuf 消息。
type PBSendGift struct {
	UID         uint64
	Uname       string
	Face        string
	GuardLevel  uint64
	MedalInfo   *PBMedal
	BlindGift   *PBBlindGift
	GiftItems   []PBGiftItem
	SenderUInfo *PBUser
}

// PBBlindGift 是盲盒礼物的原始配置与结果信息。
type PBBlindGift struct {
	ConfigID          uint64
	OriginalGiftID    uint64
	OriginalGiftName  string
	From              uint64
	GiftAction        string
	OriginalGiftPrice uint64
	GiftTipPrice      uint64
}

// PBGiftItem 是一次礼物赠送结果；十连盲盒可能产生多条。
type PBGiftItem struct {
	GiftID             uint32
	GiftName           string
	Num                uint32
	Price              uint64
	TotalCoin          uint64
	CoinType           string
	TID                string
	Timestamp          uint64
	SuperBatchGiftNum  uint64
	BatchComboID       string
	ComboResourcesID   uint64
	ComboTotalCoin     uint64
	ComboStayTime      uint64
	ShowBatchComboSend *bool
	Action             string
	GiftInfo           *PBGiftInfo
	GiftTipPrice       uint64
}

// PBGiftInfo 是礼物素材信息。
type PBGiftInfo struct {
	ImgBasic string
	Webp     string
	GIF      string
	EffectID *uint64
}

// PBInteractWord 是 INTERACT_WORD_V2 的 protobuf 消息。
type PBInteractWord struct {
	UID         uint64
	Username    string
	MsgType     uint32
	RoomID      uint32
	Timestamp   uint32
	TimestampMS uint64
	TriggerTime uint64
	GuardType   uint32
	Medal       *PBMedal
	UInfo       *PBUser
}

// PBOnlineRank 是 ONLINE_RANK_V2 / ONLINE_RANK_V3 的 protobuf 消息。
type PBOnlineRank struct {
	RankType string
	Users    []PBOnlineRankUser
}

// PBOnlineRankUser 是排行榜中的一位用户。
type PBOnlineRankUser struct {
	UID        uint64
	Face       string
	Score      string
	Uname      string
	Rank       uint32
	GuardLevel *uint32
	IsMystery  *bool
	UInfo      *PBUser
}

// PBDanmaku 是 DANMU_MSG 的 dm_v2 protobuf 消息。
type PBDanmaku struct {
	IDStr      string
	Mode       int32
	FontSize   int32
	Color      uint32
	MidHash    string
	Content    string
	Ctime      int64
	Weight     int32
	Rnd        int64
	Attr       int32
	BizScene   int32
	DMType     int32
	Emoticons  map[string]PBEmoticon
	SendFromMe bool
	User       *PBUser
	Room       *PBRoom
}

// PBEmoticon 是弹幕表情资源。
type PBEmoticon struct {
	Unique       string
	URL          string
	IsDynamic    bool
	InPlayerArea int32
	BulgeDisplay int32
	Height       int32
	Width        int32
}

// PBRoom 是 dm_v2 中的房间信息。
type PBRoom struct {
	UID  uint64
	Name string
}

// ParseSendGiftV2 解析 SEND_GIFT_V2.data.pb。GiftItems 保留一条通知内的
// 全部礼物结果（十连盲盒会有多条），顶层发送者与盲盒信息不丢弃。
// 字段缺失或 wire 格式非法时返回错误。
func ParseSendGiftV2(raw []byte) (PBSendGift, error) {
	var out PBSendGift
	fields, err := eventPBFields(raw)
	if err != nil {
		return out, err
	}
	if err = decodePBSendGift(fields, &out); err != nil {
		return PBSendGift{}, err
	}
	if len(out.GiftItems) == 0 {
		return PBSendGift{}, errDanmakuPBEventEmpty
	}
	return out, nil
}

// ParseInteractWordV2 解析 INTERACT_WORD_V2.data.pb。
func ParseInteractWordV2(raw []byte) (PBInteractWord, error) {
	var out PBInteractWord
	fields, err := eventPBFields(raw)
	if err != nil {
		return out, err
	}
	return out, decodePBInteractWord(fields, &out)
}

// ParseOnlineRankPB 解析 ONLINE_RANK_V2 / ONLINE_RANK_V3.data.pb。
func ParseOnlineRankPB(raw []byte) (PBOnlineRank, error) {
	var out PBOnlineRank
	fields, err := eventPBFields(raw)
	if err != nil {
		return out, err
	}
	return out, decodePBOnlineRank(fields, &out)
}

// ParseDanmakuPB 解析 DANMU_MSG.dm_v2 的 base64 protobuf 载荷。
func ParseDanmakuPB(raw []byte) (PBDanmaku, error) {
	var out PBDanmaku
	var envelope struct {
		DMV2 string `json:"dm_v2"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return out, fmt.Errorf("%w: %w", errDanmakuPBEventMalformed, err)
	}
	if envelope.DMV2 == "" {
		return out, errDanmakuPBPayloadMissing
	}
	data, err := base64.StdEncoding.DecodeString(envelope.DMV2)
	if err != nil {
		return out, fmt.Errorf("%w: %w", errDanmakuPBBadBase64, err)
	}
	fields, err := decodePBFields(data)
	if err != nil {
		return out, err
	}
	if len(fields) == 0 {
		return out, errDanmakuPBEventEmpty
	}
	return out, decodePBDanmaku(fields, &out)
}

func eventPBPayload(raw []byte) ([]byte, error) {
	var envelope struct {
		Data struct {
			PB *string `json:"pb"`
		} `json:"data"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		return nil, fmt.Errorf("%w: %w", errDanmakuPBEventMalformed, err)
	}
	if envelope.Data.PB == nil || *envelope.Data.PB == "" {
		return nil, errDanmakuPBPayloadMissing
	}
	data, err := base64.StdEncoding.DecodeString(*envelope.Data.PB)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errDanmakuPBBadBase64, err)
	}
	return data, nil
}

func eventPBFields(raw []byte) ([]pbField, error) {
	data, err := eventPBPayload(raw)
	if err != nil {
		return nil, err
	}
	fields, err := decodePBFields(data)
	if err != nil {
		return nil, err
	}
	if len(fields) == 0 {
		return nil, errDanmakuPBEventEmpty
	}
	return fields, nil
}

// pbFieldMap 把顶层字段按编号聚合。不同 wire 类型的同号字段并存时，
// 读取函数各自筛选匹配的 wire 类型，未知字段直接忽略。
func pbFieldMap(fields []pbField) (map[uint32][]pbField, error) {
	out := make(map[uint32][]pbField, len(fields))
	for _, field := range fields {
		if field.Wire == pbWireStartGroup || field.Wire == pbWireEndGroup {
			return nil, errDanmakuPBEventMalformed
		}
		out[field.Number] = append(out[field.Number], field)
	}
	return out, nil
}

func firstPBField(fields map[uint32][]pbField, number uint32, wires ...uint8) (pbField, bool) {
	for _, field := range fields[number] {
		if slices.Contains(wires, field.Wire) {
			return field, true
		}
	}
	return pbField{}, false
}

func pbUint64(fields map[uint32][]pbField, number uint32) uint64 {
	field, ok := firstPBField(fields, number, pbWireVarint, pbWireFixed64, pbWireFixed32)
	if !ok {
		return 0
	}
	return pbValue(field)
}

func pbUint32(fields map[uint32][]pbField, number uint32) (uint32, error) {
	field, ok := firstPBField(fields, number, pbWireVarint, pbWireFixed64, pbWireFixed32)
	if !ok {
		return 0, nil
	}
	value := pbValue(field)
	if value > math.MaxUint32 {
		return 0, fmt.Errorf("%w: field %d = %d", errDanmakuPBValueTooLarge, number, value)
	}
	return uint32(value & math.MaxUint32), nil
}

func pbInt32(fields map[uint32][]pbField, number uint32) (int32, error) {
	value, err := pbUint32(fields, number)
	if err != nil {
		return 0, err
	}
	return int32(value), nil //nolint:gosec // protobuf int32 按低 32 位补码还原
}

func pbInt64(fields map[uint32][]pbField, number uint32) int64 {
	return int64(pbUint64(fields, number)) //nolint:gosec // protobuf int64 按补码位还原
}

func pbBool(fields map[uint32][]pbField, number uint32) bool {
	return pbUint64(fields, number) != 0
}

func pbOptionalBool(fields map[uint32][]pbField, number uint32) *bool {
	if _, ok := firstPBField(fields, number, pbWireVarint, pbWireFixed64, pbWireFixed32); !ok {
		return nil
	}
	value := pbBool(fields, number)
	return &value
}

func pbOptionalUint64(fields map[uint32][]pbField, number uint32) *uint64 {
	if _, ok := firstPBField(fields, number, pbWireVarint, pbWireFixed64, pbWireFixed32); !ok {
		return nil
	}
	value := pbUint64(fields, number)
	return &value
}

func pbString(fields map[uint32][]pbField, number uint32) (string, error) {
	field, ok := firstPBField(fields, number, pbWireBytes)
	if !ok {
		return "", nil
	}
	if !utf8.Valid(field.Bytes) {
		return "", fmt.Errorf("%w: field %d", errDanmakuPBBadText, number)
	}
	return pbText(field), nil
}

func pbMessage(fields map[uint32][]pbField, number uint32) ([]pbField, bool, error) {
	field, ok := firstPBField(fields, number, pbWireBytes)
	if !ok {
		return nil, false, nil
	}
	nested, err := decodePBFields(field.Bytes)
	if err != nil {
		return nil, false, err
	}
	return nested, true, nil
}

// decodePBEventUserInfo 解析 events.proto 的 EventUserInfo / EventOnlineRankUser.uinfo。
func decodePBEventUserInfo(fields []pbField, out *PBUser) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.UID = pbUint64(m, 1)
	base, ok, err := pbMessage(m, 2)
	if err != nil {
		return err
	}
	if ok {
		if err = decodePBEventUserBase(base, out); err != nil {
			return err
		}
	}
	medal, ok, err := pbMessage(m, 3)
	if err != nil {
		return err
	}
	if ok {
		out.Medal = &PBMedal{}
		if err = decodePBEventUserMedal(medal, out.Medal); err != nil {
			return err
		}
	}
	wealth, ok, err := pbMessage(m, 4)
	if err != nil {
		return err
	}
	if ok {
		out.Wealth = &PBWealth{}
		if err = decodePBWealth(wealth, out.Wealth); err != nil {
			return err
		}
	}
	guard, ok, err := pbMessage(m, 6)
	if err != nil {
		return err
	}
	if ok {
		out.Guard = &PBGuard{}
		if err = decodePBGuard(guard, out.Guard); err != nil {
			return err
		}
	}
	return nil
}

func decodePBEventUserBase(fields []pbField, out *PBUser) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.Name, err = pbString(m, 1); err != nil {
		return err
	}
	if out.Face, err = pbString(m, 2); err != nil {
		return err
	}
	if out.NameColor, err = pbUint32(m, 3); err != nil {
		return err
	}
	out.IsMystery = pbBool(m, 4)
	return nil
}
func decodePBEventUserMedal(fields []pbField, out *PBMedal) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.Name, err = pbString(m, 1); err != nil {
		return err
	}
	if out.Level, err = pbUint32(m, 2); err != nil {
		return err
	}
	if out.Color, err = pbUint32(m, 6); err != nil {
		return err
	}
	out.RUID = pbUint64(m, 10)
	out.ColorV2, err = pbUint32(m, 18)
	return err
}

func decodePBWealth(fields []pbField, out *PBWealth) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.Level, err = pbUint32(m, 1)
	return err
}

func decodePBGuard(fields []pbField, out *PBGuard) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.Level, err = pbUint32(m, 1); err != nil {
		return err
	}
	out.ExpiredStr, err = pbString(m, 2)
	return err
}

func decodePBSendGift(fields []pbField, out *PBSendGift) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.UID = pbUint64(m, 1)
	if out.Uname, err = pbString(m, 2); err != nil {
		return err
	}
	if out.Face, err = pbString(m, 3); err != nil {
		return err
	}
	out.GuardLevel = pbUint64(m, 5)
	medal, ok, err := pbMessage(m, 8)
	if err != nil {
		return err
	}
	if ok {
		out.MedalInfo = &PBMedal{}
		if err = decodePBGiftMedal(medal, out.MedalInfo); err != nil {
			return err
		}
	}
	blind, ok, err := pbMessage(m, 9)
	if err != nil {
		return err
	}
	if ok {
		out.BlindGift = &PBBlindGift{}
		if err = decodePBBlindGift(blind, out.BlindGift); err != nil {
			return err
		}
	}
	for _, field := range m[10] {
		if field.Wire != pbWireBytes {
			return fmt.Errorf("%w: gift item has wire %d", errDanmakuPBEventMalformed, field.Wire)
		}
		item, err := decodePBGiftItem(field.Bytes)
		if err != nil {
			return err
		}
		out.GiftItems = append(out.GiftItems, item)
	}
	sender, ok, err := pbMessage(m, 15)
	if err != nil {
		return err
	}
	if ok {
		out.SenderUInfo = &PBUser{}
		if err = decodePBEventUserInfo(sender, out.SenderUInfo); err != nil {
			return err
		}
	}
	return nil
}

func decodePBGiftMedal(fields []pbField, out *PBMedal) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.RUID = pbUint64(m, 1)
	if out.Level, err = pbUint32(m, 5); err != nil {
		return err
	}
	if out.Name, err = pbString(m, 6); err != nil {
		return err
	}
	out.IsLighted = pbOptionalBool(m, 11)
	return err
}

func decodePBBlindGift(fields []pbField, out *PBBlindGift) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	values := []struct {
		dst    *uint64
		number uint32
	}{
		{&out.ConfigID, 1}, {&out.OriginalGiftID, 2}, {&out.From, 4},
		{&out.OriginalGiftPrice, 6}, {&out.GiftTipPrice, 7},
	}
	for _, item := range values {
		*item.dst = pbUint64(m, item.number)
	}
	if out.OriginalGiftName, err = pbString(m, 3); err != nil {
		return err
	}
	out.GiftAction, err = pbString(m, 5)
	return err
}

func decodePBGiftItem(data []byte) (PBGiftItem, error) {
	var out PBGiftItem
	fields, err := decodePBFields(data)
	if err != nil {
		return out, err
	}
	m, err := pbFieldMap(fields)
	if err != nil {
		return out, err
	}
	if out.GiftID, err = pbUint32(m, 1); err != nil {
		return out, err
	}
	if out.GiftName, err = pbString(m, 2); err != nil {
		return out, err
	}
	if out.Num, err = pbUint32(m, 3); err != nil {
		return out, err
	}
	if out.Price = pbUint64(m, 5); err != nil {
		return out, err
	}
	if out.TotalCoin = pbUint64(m, 7); err != nil {
		return out, err
	}
	if out.CoinType, err = pbString(m, 8); err != nil {
		return out, err
	}
	if out.TID, err = pbString(m, 9); err != nil {
		return out, err
	}
	if out.Timestamp = pbUint64(m, 10); err != nil {
		return out, err
	}
	if out.SuperBatchGiftNum = pbUint64(m, 11); err != nil {
		return out, err
	}
	if out.BatchComboID, err = pbString(m, 12); err != nil {
		return out, err
	}
	if out.ComboResourcesID = pbUint64(m, 13); err != nil {
		return out, err
	}
	if out.ComboTotalCoin = pbUint64(m, 14); err != nil {
		return out, err
	}
	if out.ComboStayTime = pbUint64(m, 15); err != nil {
		return out, err
	}
	if out.ShowBatchComboSend = pbOptionalBool(m, 17); err != nil {
		return out, err
	}
	if out.Action, err = pbString(m, 18); err != nil {
		return out, err
	}
	info, ok, err := pbMessage(m, 35)
	if err != nil {
		return out, err
	}
	if ok {
		out.GiftInfo = &PBGiftInfo{}
		if err = decodePBGiftInfo(info, out.GiftInfo); err != nil {
			return out, err
		}
	}
	out.GiftTipPrice = pbUint64(m, 36)
	return out, err
}

func decodePBGiftInfo(fields []pbField, out *PBGiftInfo) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.ImgBasic, err = pbString(m, 1); err != nil {
		return err
	}
	if out.Webp, err = pbString(m, 2); err != nil {
		return err
	}
	if out.GIF, err = pbString(m, 5); err != nil {
		return err
	}
	out.EffectID = pbOptionalUint64(m, 3)
	return err
}

func decodePBInteractWord(fields []pbField, out *PBInteractWord) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.UID = pbUint64(m, 1)
	if out.Username, err = pbString(m, 2); err != nil {
		return err
	}
	if out.MsgType, err = pbUint32(m, 5); err != nil {
		return err
	}
	if out.RoomID, err = pbUint32(m, 6); err != nil {
		return err
	}
	if out.Timestamp, err = pbUint32(m, 7); err != nil {
		return err
	}
	out.TimestampMS = pbUint64(m, 8)
	out.TriggerTime = pbUint64(m, 15)
	if out.GuardType, err = pbUint32(m, 16); err != nil {
		return err
	}
	medal, ok, err := pbMessage(m, 9)
	if err != nil {
		return err
	}
	if ok {
		out.Medal = &PBMedal{}
		if err = decodePBInteractMedal(medal, out.Medal); err != nil {
			return err
		}
	}
	uinfo, ok, err := pbMessage(m, 22)
	if err != nil {
		return err
	}
	if ok {
		out.UInfo = &PBUser{}
		if err = decodePBEventUserInfo(uinfo, out.UInfo); err != nil {
			return err
		}
	}
	return nil
}

func decodePBInteractMedal(fields []pbField, out *PBMedal) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.RUID = pbUint64(m, 1)
	if out.Level, err = pbUint32(m, 2); err != nil {
		return err
	}
	if out.Name, err = pbString(m, 3); err != nil {
		return err
	}
	out.IsLighted = pbOptionalBool(m, 8)
	if out.GuardLv, err = pbUint32(m, 9); err != nil {
		return err
	}
	out.RoomID, err = pbUint32(m, 12)
	return err
}

func decodePBOnlineRank(fields []pbField, out *PBOnlineRank) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.RankType, err = pbString(m, 1); err != nil {
		return err
	}
	for _, field := range m[3] {
		if field.Wire != pbWireBytes {
			return fmt.Errorf("%w: rank user has wire %d", errDanmakuPBEventMalformed, field.Wire)
		}
		user, err := decodePBOnlineRankUser(field.Bytes)
		if err != nil {
			return err
		}
		out.Users = append(out.Users, user)
	}
	return nil
}

func decodePBOnlineRankUser(data []byte) (PBOnlineRankUser, error) {
	var out PBOnlineRankUser
	fields, err := decodePBFields(data)
	if err != nil {
		return out, err
	}
	m, err := pbFieldMap(fields)
	if err != nil {
		return out, err
	}
	if out.UID = pbUint64(m, 1); err != nil {
		return out, err
	}
	if out.Face, err = pbString(m, 2); err != nil {
		return out, err
	}
	if out.Score, err = pbString(m, 3); err != nil {
		return out, err
	}
	if out.Uname, err = pbString(m, 4); err != nil {
		return out, err
	}
	if out.Rank, err = pbUint32(m, 5); err != nil {
		return out, err
	}
	if _, ok := firstPBField(m, 6, pbWireVarint, pbWireFixed64, pbWireFixed32); ok {
		value, err := pbUint32(m, 6)
		if err != nil {
			return out, err
		}
		out.GuardLevel = &value
	}
	if out.IsMystery = pbOptionalBool(m, 7); err != nil {
		return out, err
	}
	uinfo, ok, err := pbMessage(m, 8)
	if err != nil {
		return out, err
	}
	if ok {
		out.UInfo = &PBUser{}
		if err = decodePBEventUserInfo(uinfo, out.UInfo); err != nil {
			return out, err
		}
	}
	return out, nil
}

func decodePBDanmaku(fields []pbField, out *PBDanmaku) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.IDStr, err = pbString(m, 1); err != nil {
		return err
	}
	if out.Mode, err = pbInt32(m, 2); err != nil {
		return err
	}
	if out.FontSize, err = pbInt32(m, 3); err != nil {
		return err
	}
	if out.Color, err = pbUint32(m, 4); err != nil {
		return err
	}
	if out.MidHash, err = pbString(m, 5); err != nil {
		return err
	}
	if out.Content, err = pbString(m, 6); err != nil {
		return err
	}
	out.Ctime = pbInt64(m, 7)
	if out.Weight, err = pbInt32(m, 8); err != nil {
		return err
	}
	out.Rnd = pbInt64(m, 9)
	if out.Attr, err = pbInt32(m, 10); err != nil {
		return err
	}
	if out.BizScene, err = pbInt32(m, 11); err != nil {
		return err
	}
	if out.DMType, err = pbInt32(m, 13); err != nil {
		return err
	}
	out.SendFromMe = pbBool(m, 18)
	for _, field := range m[14] {
		if field.Wire != pbWireBytes {
			return fmt.Errorf("%w: emoticon entry has wire %d", errDanmakuPBEventMalformed, field.Wire)
		}
		entry, err := decodePBEmoticonEntry(field.Bytes)
		if err != nil {
			return err
		}
		if strings.TrimSpace(entry.key) == "" {
			continue
		}
		if out.Emoticons == nil {
			out.Emoticons = make(map[string]PBEmoticon)
		}
		out.Emoticons[entry.key] = entry.value
	}
	user, ok, err := pbMessage(m, 20)
	if err != nil {
		return err
	}
	if ok {
		out.User = &PBUser{}
		if err = decodePBDanmakuUser(user, out.User); err != nil {
			return err
		}
	}
	room, ok, err := pbMessage(m, 21)
	if err != nil {
		return err
	}
	if ok {
		out.Room = &PBRoom{}
		if err = decodePBRoom(room, out.Room); err != nil {
			return err
		}
	}
	return nil
}

type pbEmoticonEntry struct {
	key   string
	value PBEmoticon
}

func decodePBEmoticonEntry(data []byte) (pbEmoticonEntry, error) {
	var out pbEmoticonEntry
	fields, err := decodePBFields(data)
	if err != nil {
		return out, err
	}
	m, err := pbFieldMap(fields)
	if err != nil {
		return out, err
	}
	if out.key, err = pbString(m, 1); err != nil {
		return out, err
	}
	nested, ok, err := pbMessage(m, 2)
	if err != nil || !ok {
		return out, err
	}
	em, err := pbFieldMap(nested)
	if err != nil {
		return out, err
	}
	if out.value.Unique, err = pbString(em, 1); err != nil {
		return out, err
	}
	if out.value.URL, err = pbString(em, 2); err != nil {
		return out, err
	}
	out.value.IsDynamic = pbBool(em, 3)
	if out.value.InPlayerArea, err = pbInt32(em, 4); err != nil {
		return out, err
	}
	if out.value.BulgeDisplay, err = pbInt32(em, 5); err != nil {
		return out, err
	}
	if out.value.Height, err = pbInt32(em, 6); err != nil {
		return out, err
	}
	out.value.Width, err = pbInt32(em, 7)
	return out, err
}

// decodePBDanmakuUser 解析 dm_v2 的 EventDmUser，字段布局与 EventUserInfo 不同。
func decodePBDanmakuUser(fields []pbField, out *PBUser) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.UID = pbUint64(m, 1)
	if out.Name, err = pbString(m, 2); err != nil {
		return err
	}
	if out.Face, err = pbString(m, 4); err != nil {
		return err
	}
	medal, ok, err := pbMessage(m, 11)
	if err != nil {
		return err
	}
	if ok {
		out.Medal = &PBMedal{}
		if err = decodePBDanmakuMedal(medal, out.Medal); err != nil {
			return err
		}
	}
	wealth, ok, err := pbMessage(m, 15)
	if err != nil {
		return err
	}
	if ok {
		out.Wealth = &PBWealth{}
		if err = decodePBWealth(wealth, out.Wealth); err != nil {
			return err
		}
	}
	return nil
}

func decodePBDanmakuMedal(fields []pbField, out *PBMedal) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	if out.Level, err = pbUint32(m, 1); err != nil {
		return err
	}
	if out.Name, err = pbString(m, 2); err != nil {
		return err
	}
	out.Color, err = pbUint32(m, 4)
	return err
}

func decodePBRoom(fields []pbField, out *PBRoom) error {
	m, err := pbFieldMap(fields)
	if err != nil {
		return err
	}
	out.UID = pbUint64(m, 1)
	out.Name, err = pbString(m, 2)
	return err
}
