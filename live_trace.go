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
	Ets        int64  // 上一包（进房或上一跳）响应中的 timestamp（秒）
	Benchmark  string // 上一包响应中的 secret_key；留空无法计算签名，返回错误
	SecretRule []int  // 上一包响应中的 secret_rule，签名据此选择哈希规则
	Time       int64  // 上一包响应中的 heartbeat_interval（秒）。调用方应等待该间隔后再发下一包
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
	Time      int64  `json:"time"`      // 上一包响应中的 heartbeat_interval
	TS        int64  `json:"ts"`        // 毫秒时间戳
	UA        string `json:"ua"`        // User-Agent
	VisitID   string `json:"visit_id"`  // 访问标识
	Device    string `json:"device"`    // [LIVE_BUVID, uuid] 的 JSON 数组字符串
}

// SendLiveHeartBeat 发送 web 端直播心跳（x25Kn/X），CSRF 自动填入表单。
// Ets、Benchmark、SecretRule、Time 必须来自上一包（进房或上一跳）的响应，签名 s 由库计算。
func (c *Client) SendLiveHeartBeat(ctx context.Context, param SendLiveHeartBeatParam) (*LiveHeartBeatResult, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	if param.Benchmark == "" {
		return nil, errors.New("直播心跳缺少 secret_key：请先调用 EnterLiveRoom 并把响应中的 SecretKey 传入")
	}
	if param.Time <= 0 {
		return nil, errors.New("直播心跳缺少 heartbeat_interval：请把上一包响应中的 HeartbeatInterval 传入 Time")
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
	payload, err := liveHeartBeatSignatureJSON(liveHeartBeatSignature{
		ParentID: param.ParentID,
		AreaID:   param.AreaID,
		SeqID:    param.SeqID,
		RoomID:   param.RoomID,
		Buvid:    param.Buvid,
		UUID:     param.UUID,
		Ets:      param.Ets,
		Time:     param.Time,
		TS:       param.TS,
	})
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
		Time:      param.Time,
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
// 注意：平台实现（pc 算法）不含 ua/device/id 字段，ua 只出现在表单里。
type liveHeartBeatSignature struct {
	Platform string `json:"platform"`
	ParentID int64  `json:"parent_id"`
	AreaID   int64  `json:"area_id"`
	SeqID    int64  `json:"seq_id"`
	RoomID   int64  `json:"room_id"`
	Buvid    string `json:"buvid"`
	UUID     string `json:"uuid"`
	Ets      int64  `json:"ets"`
	Time     int64  `json:"time"`
	TS       int64  `json:"ts"`
}

// liveHeartBeatSignatureJSON 按固定字段顺序构造心跳签名明文 JSON（platform 固定为 web）。
func liveHeartBeatSignatureJSON(s liveHeartBeatSignature) (string, error) {
	s.Platform = "web"
	data, err := json.Marshal(s)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(data), nil
}

// sypderSign 对签名明文按 secret_rule 顺序做级联 HMAC：每轮均以 secretKey 为 key，
// 对上一轮输出的 hex 文本继续计算 HMAC，最终返回小写 hex。
// 规则映射：0-MD5（连续计算两次，与官方实现一致）、1-SHA1、2-SHA256、3-SHA224、
// 4-SHA512、5-SHA384，其它规则跳过。
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
		times := 1
		if rule == 0 {
			// 官方实现对 MD5 规则连续计算两次。
			times = 2
		}
		for range times {
			mac := hmac.New(newHash, []byte(secretKey))
			_, _ = mac.Write([]byte(text)) // hash.Write 永远不会返回错误
			text = hex.EncodeToString(mac.Sum(nil))
		}
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

// —— Web 端观看时长新链路（data.bilivideo.com 域）——
//
// 当前网页播放器已把观看时长上报迁移到独立端点：进房 reportEnterRoom、
// 心跳 reportHeartBeat，表单为 WebHeartBeatCheck 结构（JSON），心跳额外携带
// csn 签名。签名由官方 wasm（skynet）对同一 JSON 计算，本库不内置该算法，
// 调用方可用 wazero 等运行时执行官方 wasm 得到 csn 后传入。

// WebWatchReportParam 是观看时长上报（进房与心跳共用）的表单字段。
type WebWatchReportParam struct {
	UID          int64  `json:"uid"`           // 当前账号 UID，未登录传 0
	Buvid        string `json:"buvid"`         // Cookie 中 LIVE_BUVID 的值
	Platform     string `json:"platform"`      // 平台。一般为 web
	RoomID       int64  `json:"room_id"`       // 直播间号
	PlayURL      string `json:"play_url"`      // 当前播放的流地址，来自 getRoomPlayInfo
	QID          int64  `json:"qid"`           // 上报序号。进房为 0，心跳从 1 开始逐次递增
	SID          string `json:"sid"`           // 会话 id。进房响应下发，之后原样携带
	CTS          int64  `json:"cts"`           // 客户端毫秒时间戳
	STKY         string `json:"stky"`          // 心跳密钥。进房响应下发，每次心跳响应轮换
	ScreenStatus int    `json:"screen_status"` // 界面状态。官方为 1-100 随机数
	ClickStatus  int    `json:"click_status"`  // 点击状态。官方为 1-100 随机数
}

// WebWatchReportResult 是观看时长上报接口的响应。
type WebWatchReportResult struct {
	STKY string `json:"stky"` // 下一次心跳使用的心跳密钥（每次轮换）
	SID  string `json:"sid"`  // 会话 id，进房后原样携带
	HBIL int    `json:"hbil"` // 心跳间隔（秒）
}

// ReportWebWatchEnter 上报 web 端进入直播间（观看时长新链路）。
// 响应中的 STKY、SID、HBIL 用于后续心跳。
func (c *Client) ReportWebWatchEnter(ctx context.Context, param WebWatchReportParam) (*WebWatchReportResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://data.bilivideo.com/log/web/te9Kl"
	)
	return execute[*WebWatchReportResult](ctx, c, method, url, param)
}

// ReportWebWatchHeartBeatParam 指定 web 端观看心跳的参数，字段与进房一致，
// 另需携带上一包响应轮换出的 STKY 与 csn 签名。
type ReportWebWatchHeartBeatParam struct {
	UID          int64  `json:"uid"`           // 当前账号 UID，未登录传 0
	Buvid        string `json:"buvid"`         // Cookie 中 LIVE_BUVID 的值
	Platform     string `json:"platform"`      // 平台。一般为 web
	RoomID       int64  `json:"room_id"`       // 直播间号
	PlayURL      string `json:"play_url"`      // 当前播放的流地址
	QID          int64  `json:"qid"`           // 上报序号。心跳从 1 开始逐次递增
	SID          string `json:"sid"`           // 会话 id。进房响应下发，之后原样携带
	CTS          int64  `json:"cts"`           // 客户端毫秒时间戳
	STKY         string `json:"stky"`          // 上一包响应轮换出的心跳密钥
	ScreenStatus int    `json:"screen_status"` // 界面状态。官方为 1-100 随机数
	ClickStatus  int    `json:"click_status"`  // 点击状态。官方为 1-100 随机数
	Csn          string `json:"csn"`           // csn 签名。由官方 skynet wasm 对其余字段的 JSON 计算
}

// ReportWebWatchHeartBeat 发送 web 端观看心跳（观看时长新链路）。
// Csn 由官方 skynet wasm 对本结构其余字段的 JSON 计算，本库不内置该算法。
// 响应中的 STKY 为下一次心跳的密钥。
func (c *Client) ReportWebWatchHeartBeat(ctx context.Context, param ReportWebWatchHeartBeatParam) (*WebWatchReportResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://data.bilivideo.com/log/web/s82Tq"
	)
	return execute[*WebWatchReportResult](ctx, c, method, url, param)
}
