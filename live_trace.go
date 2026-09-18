package bilibili

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// Web 端直播数据上报相关接口（live-trace 域），包括进房上报与心跳。
// 响应模型见 live_trace_model.go。

// liveHeartBeatInterval 是心跳接口签名明文中固定的 time 字段值。
const liveHeartBeatInterval = 60

// EnterLiveRoomParam 指定 web 端进房数据上报的参数。
type EnterLiveRoomParam struct {
	ParentID int64  // 直播分区 id（父分区）
	AreaID   int64  // 直播分区 id（子分区）
	Seq      int64  // 包序号。进房一般为 0，之后每包递增
	RoomID   int64  // 直播间号
	RUID     int64  // 主播 UID
	TS       int64  // 发送时刻的毫秒时间戳；留空自动取当前时间
	UA       string // User-Agent；留空时发送客户端默认 User-Agent
	Buvid    string // Cookie 中 LIVE_BUVID 的值；留空时自动从请求 Cookie 快照读取，仍缺失则发送空串
	UUID     string // 设备 uuid；留空时自动生成
	VisitID  string // 访问标识；按源实现发送空串
}

// liveEnterRoomForm 是进房接口的表单字段。id 与 device 为 JSON 数组字符串。
type liveEnterRoomForm struct {
	ID        string `json:"id"`         // [parent_id,area_id,seq,room_id] 的 JSON 数组字符串
	RUID      int64  `json:"ruid"`       // 主播 UID
	TS        int64  `json:"ts"`         // 毫秒时间戳
	IsPatch   int    `json:"is_patch"`   // 固定为 0
	HeartBeat string `json:"heart_beat"` // 固定为 "[]"
	UA        string `json:"ua"`         // User-Agent
	VisitID   string `json:"visit_id"`   // 访问标识
	Device    string `json:"device"`     // [LIVE_BUVID, uuid] 的 JSON 数组字符串
}

// EnterLiveRoom 上报 web 端进房数据（x25Kn/E），CSRF 自动填入表单。
// 响应中的 SecretKey、SecretRule、Timestamp 需由调用方带入下一包心跳（见 SendLiveHeartBeat）。
func (c *Client) EnterLiveRoom(ctx context.Context, param EnterLiveRoomParam) (*LiveHeartBeatResult, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	r := c.newRequest(ctx)
	if param.TS == 0 {
		param.TS = time.Now().UnixMilli()
	}
	if param.UA == "" {
		param.UA = defaultUserAgent
	}
	if param.UUID == "" {
		uuid, err := randomLiveUUID()
		if err != nil {
			return nil, err
		}
		param.UUID = uuid
	}
	if param.Buvid == "" {
		param.Buvid = cookieValue(r.Cookies, "LIVE_BUVID")
	}
	id, err := liveTraceIDJSON(param.ParentID, param.AreaID, param.Seq, param.RoomID)
	if err != nil {
		return nil, err
	}
	device, err := liveTraceDeviceJSON(param.Buvid, param.UUID)
	if err != nil {
		return nil, err
	}
	const (
		method = resty.MethodPost
		url    = "https://live-trace.bilibili.com/xlive/data-interface/v1/x25Kn/E"
	)
	// withParams 的默认分支已将 Content-Type 设为 application/x-www-form-urlencoded。
	return executeRequest[*LiveHeartBeatResult](c, r, method, url, liveEnterRoomForm{
		ID:        id,
		RUID:      param.RUID,
		TS:        param.TS,
		IsPatch:   0,
		HeartBeat: "[]",
		UA:        param.UA,
		VisitID:   param.VisitID,
		Device:    device,
	}, moveFormParams("id", "ruid", "ts", "is_patch", "heart_beat", "ua", "visit_id", "device"),
		fillFormCsrf(c))
}

// SendLiveHeartBeatParam 指定 web 端直播心跳的参数。签名 s 由库内部按 secret_rule 级联 HMAC 计算。
type SendLiveHeartBeatParam struct {
	ParentID   int64  // 直播分区 id（父分区）
	AreaID     int64  // 直播分区 id（子分区）
	SeqID      int64  // 包序号。与进房时的 Seq 保持一致并逐包递增
	RoomID     int64  // 直播间号
	Ets        int64  // 上一包（进房或上一跳）响应中的 timestamp
	Benchmark  string // 上一包响应中的 secret_key；留空无法计算签名，返回错误
	SecretRule []int  // 上一包响应中的 secret_rule，签名据此选择哈希规则
	TS         int64  // 发送时刻的毫秒时间戳；留空自动取当前时间
	UA         string // User-Agent；留空时发送客户端默认 User-Agent
	Buvid      string // Cookie 中 LIVE_BUVID 的值；留空时自动从请求 Cookie 快照读取，仍缺失则发送空串
	UUID       string // 本次心跳的设备 uuid；留空时自动生成，每包应使用新的 uuid
	VisitID    string // 访问标识；按源实现发送空串
}

// liveHeartBeatForm 是心跳接口的表单字段。id 与 device 为 JSON 数组字符串。
type liveHeartBeatForm struct {
	S         string `json:"s"`         // 按 secret_rule 级联 HMAC 计算的签名，hex 小写
	ID        string `json:"id"`        // [parent_id,area_id,seq_id,room_id] 的 JSON 数组字符串
	Ets       int64  `json:"ets"`       // 上一包响应中的 timestamp
	Benchmark string `json:"benchmark"` // 上一包响应中的 secret_key
	Time      int    `json:"time"`      // 固定为 60
	TS        int64  `json:"ts"`        // 毫秒时间戳
	UA        string `json:"ua"`        // User-Agent
	VisitID   string `json:"visit_id"`  // 访问标识
	Device    string `json:"device"`    // [LIVE_BUVID, uuid] 的 JSON 数组字符串
}

// SendLiveHeartBeat 发送 web 端直播心跳（x25Kn/X），CSRF 自动填入表单。
// Ets、Benchmark、SecretRule 必须来自上一包（进房或上一跳）的响应，签名 s 由库计算。
func (c *Client) SendLiveHeartBeat(ctx context.Context, param SendLiveHeartBeatParam) (*LiveHeartBeatResult, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	if param.Benchmark == "" {
		return nil, errors.New("直播心跳缺少 secret_key：请先调用 EnterLiveRoom 并把响应中的 SecretKey 传入")
	}
	r := c.newRequest(ctx)
	if param.TS == 0 {
		param.TS = time.Now().UnixMilli()
	}
	if param.UA == "" {
		param.UA = defaultUserAgent
	}
	if param.UUID == "" {
		uuid, err := randomLiveUUID()
		if err != nil {
			return nil, err
		}
		param.UUID = uuid
	}
	if param.Buvid == "" {
		param.Buvid = cookieValue(r.Cookies, "LIVE_BUVID")
	}
	id, err := liveTraceIDJSON(param.ParentID, param.AreaID, param.SeqID, param.RoomID)
	if err != nil {
		return nil, err
	}
	device, err := liveTraceDeviceJSON(param.Buvid, param.UUID)
	if err != nil {
		return nil, err
	}
	payload, err := liveHeartBeatSignatureJSON(param.ParentID, param.AreaID, param.SeqID,
		param.RoomID, param.Buvid, param.UUID, param.Ets, param.TS)
	if err != nil {
		return nil, err
	}
	const (
		method = resty.MethodPost
		url    = "https://live-trace.bilibili.com/xlive/data-interface/v1/x25Kn/X"
	)
	// withParams 的默认分支已将 Content-Type 设为 application/x-www-form-urlencoded。
	return executeRequest[*LiveHeartBeatResult](c, r, method, url, liveHeartBeatForm{
		S:         sypderSign(payload, param.SecretRule, param.Benchmark),
		ID:        id,
		Ets:       param.Ets,
		Benchmark: param.Benchmark,
		Time:      liveHeartBeatInterval,
		TS:        param.TS,
		UA:        param.UA,
		VisitID:   param.VisitID,
		Device:    device,
	}, moveFormParams("s", "id", "ets", "benchmark", "time", "ts", "ua", "visit_id", "device"),
		fillFormCsrf(c))
}

// liveTraceIDJSON 把 [parent_id,area_id,seq,room_id] 序列化为接口要求的 JSON 数组字符串。
func liveTraceIDJSON(parentID, areaID, seq, roomID int64) (string, error) {
	data, err := json.Marshal([]int64{parentID, areaID, seq, roomID})
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(data), nil
}

// liveTraceDeviceJSON 把 [buvid, uuid] 序列化为接口要求的 JSON 数组字符串。
func liveTraceDeviceJSON(buvid, uuid string) (string, error) {
	data, err := json.Marshal([]string{buvid, uuid})
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(data), nil
}

// liveHeartBeatSignature 是心跳签名明文的结构。字段顺序是协议要求，不可调整。
type liveHeartBeatSignature struct {
	Platform string `json:"platform"`
	ParentID int64  `json:"parent_id"`
	AreaID   int64  `json:"area_id"`
	SeqID    int64  `json:"seq_id"`
	RoomID   int64  `json:"room_id"`
	Buvid    string `json:"buvid"`
	UUID     string `json:"uuid"`
	Ets      int64  `json:"ets"`
	Time     int    `json:"time"`
	TS       int64  `json:"ts"`
}

// liveHeartBeatSignatureJSON 按固定字段顺序构造心跳签名明文 JSON。
func liveHeartBeatSignatureJSON(parentID, areaID, seqID, roomID int64, buvid, uuid string, ets, ts int64) (string, error) {
	data, err := json.Marshal(liveHeartBeatSignature{
		Platform: "web",
		ParentID: parentID,
		AreaID:   areaID,
		SeqID:    seqID,
		RoomID:   roomID,
		Buvid:    buvid,
		UUID:     uuid,
		Ets:      ets,
		Time:     liveHeartBeatInterval,
		TS:       ts,
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(data), nil
}

// sypderSign 对签名明文按 secret_rule 顺序做级联 HMAC：每轮均以 secretKey 为 key，
// 对上一轮输出的 hex 文本继续计算 HMAC，最终返回小写 hex。
// 规则映射：0-MD5、1-SHA1、2-SHA256、3-SHA224、4-SHA512、5-SHA384，其它规则跳过。
func sypderSign(payload string, secretRule []int, secretKey string) string {
	text := payload
	for _, rule := range secretRule {
		var newHash func() hash.Hash
		switch rule {
		case 0:
			newHash = md5.New
		case 1:
			newHash = sha1.New
		case 2:
			newHash = sha256.New
		case 3:
			newHash = sha256.New224
		case 4:
			newHash = sha512.New
		case 5:
			newHash = sha512.New384
		default:
			continue
		}
		mac := hmac.New(newHash, []byte(secretKey))
		_, _ = mac.Write([]byte(text)) // hash.Write 永远不会返回错误
		text = hex.EncodeToString(mac.Sum(nil))
	}
	return text
}

// randomLiveUUID 用 crypto/rand 生成 8-4-4-4-12 格式的 uuid，用于进房与心跳的 device 字段。
func randomLiveUUID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", errors.WithStack(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}
