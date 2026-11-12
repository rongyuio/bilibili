package livestream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rongyuio/bilibili"
)

// 直播弹幕信息流连接所需的 HTTP 接口与身份信息。

var (
	errDanmakuInfoNoHost = errors.New("live danmaku info does not contain a host")
	errDanmakuBadRoomID  = errors.New("live danmaku room id must be positive")
	errDanmakuEmptyRoom  = errors.New("live danmaku room init returned no room id")
	errDanmakuInvalidUID = errors.New("live danmaku DedeUserID cookie is invalid")
)

type getDanmuInfoResult struct {
	Token    string             `json:"token"`
	HostList []getDanmuInfoHost `json:"host_list"`
}

type getDanmuInfoHost struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	WSPort  int    `json:"ws_port"`
	WSSPort int    `json:"wss_port"`
}

type getRoomInitResult struct {
	RoomID  int64 `json:"room_id"`
	ShortID int64 `json:"short_id"`
}

type getFingerSPIResult struct {
	B3 string `json:"b_3"`
	B4 string `json:"b_4"`
}

type getDanmuInfoParams struct {
	ID          int64  `json:"id"`
	Type        int    `json:"type"`
	WebLocation string `json:"web_location"`
}

type getRoomInitParams struct {
	ID int64 `json:"id"`
}

// ConnectOptions 指定建立直播弹幕信息流时的可选参数。
type ConnectOptions struct {
	// RoomID 是直播间短号或真实房间号；库会通过 room_init 解析为真实房间号。
	RoomID int64
	// UseTLS 为 true 时使用 WSS，默认 true；false 时使用明文 WS。
	UseTLS *bool
	// AuthProtover 是认证包声明的协议版本，默认 2（zlib）；可传 0、1、2、3。
	AuthProtover uint16
	// HeartbeatInterval 是心跳间隔，默认 30 秒；小于等于 0 时使用默认值。
	HeartbeatInterval time.Duration
	// EventBufferSize 是事件 channel 缓冲，默认 64；小于等于 0 时使用默认值。
	EventBufferSize int
	// Dialer 覆盖 WebSocket 拨号器；nil 时使用默认拨号器。
	Dialer WebSocketDialer
}

// danmakuConnectionInfo 是连接一次所需的服务端信息。
type danmakuConnectionInfo struct {
	RoomID   int64
	UID      int64
	Buvid3   string
	Token    string
	Host     getDanmuInfoHost
	UseTLS   bool
	Protover uint16
}

func (c *Client) connectInfo(ctx context.Context, options ConnectOptions) (danmakuConnectionInfo, error) {
	var info danmakuConnectionInfo
	if options.RoomID <= 0 {
		return info, errDanmakuBadRoomID
	}
	room, err := bilibili.Execute[*getRoomInitResult](ctx, c.Client, http.MethodGet,
		"https://api.live.bilibili.com/room/v1/Room/room_init", getRoomInitParams{ID: options.RoomID})
	if err != nil {
		return info, fmt.Errorf("resolve live room id: %w", err)
	}
	if room == nil || room.RoomID <= 0 {
		return info, fmt.Errorf("resolve live room id: %w", errDanmakuEmptyRoom)
	}
	info.RoomID = room.RoomID

	danmuInfo, err := bilibili.Execute[*getDanmuInfoResult](ctx, c.Client, http.MethodGet,
		"https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo",
		getDanmuInfoParams{ID: room.RoomID, Type: 0, WebLocation: "444.8"}, c.FillWbi())
	if err != nil {
		return info, fmt.Errorf("get live danmaku info: %w", err)
	}
	if danmuInfo == nil || len(danmuInfo.HostList) == 0 {
		return info, errDanmakuInfoNoHost
	}
	info.Token = danmuInfo.Token
	info.Host = danmuInfo.HostList[0]
	info.UseTLS = options.UseTLS == nil || *options.UseTLS
	info.Protover = options.AuthProtover
	if info.Protover > danmakuProtoBrotli {
		return info, fmt.Errorf("%w: %d", errDanmakuUnknownProto, info.Protover)
	}
	if info.Protover == 0 {
		info.Protover = danmakuDefaultAuthProto
	}

	info.UID, info.Buvid3, err = c.identity(ctx)
	if err != nil {
		return info, err
	}
	return info, nil
}

// identity 从会话 Cookie 读取 UID，缺失 buvid3 时向指纹接口补全。
func (c *Client) identity(ctx context.Context) (int64, string, error) {
	var uidValue, buvid3 string
	for _, cookie := range c.GetCookies() {
		switch cookie.Name {
		case "DedeUserID":
			uidValue = cookie.Value
		case "buvid3":
			buvid3 = cookie.Value
		}
	}
	uid := int64(0)
	if uidValue != "" {
		parsed, err := strconv.ParseInt(uidValue, 10, 64)
		if err != nil || parsed < 0 {
			return 0, "", errDanmakuInvalidUID
		}
		uid = parsed
	}
	if buvid3 == "" {
		spi, err := bilibili.Execute[*getFingerSPIResult](ctx, c.Client, http.MethodGet,
			"https://api.bilibili.com/x/frontend/finger/spi", nil)
		if err != nil {
			return 0, "", fmt.Errorf("get live danmaku buvid: %w", err)
		}
		if spi != nil {
			buvid3 = spi.B3
			c.SetCookie(&http.Cookie{Name: "buvid3", Value: spi.B3, Path: "/"})
		}
	}
	return uid, buvid3, nil
}

type danmakuAuthBody struct {
	UID      int64  `json:"uid"`
	RoomID   int64  `json:"roomid"`
	Protover uint16 `json:"protover"`
	Buvid    string `json:"buvid"`
	Platform string `json:"platform"`
	Type     int    `json:"type"`
	Key      string `json:"key"`
}

func authPacket(info danmakuConnectionInfo) ([]byte, error) {
	body, err := json.Marshal(danmakuAuthBody{
		UID:      info.UID,
		RoomID:   info.RoomID,
		Protover: info.Protover,
		Buvid:    info.Buvid3,
		Platform: "web",
		Type:     2,
		Key:      info.Token,
	})
	if err != nil {
		return nil, fmt.Errorf("encode live danmaku auth: %w", err)
	}
	return encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpAuth, body), nil
}

func heartbeatPacket() []byte {
	return encodeDanmakuPacket(danmakuProtoHeartbeat, danmakuOpHeartbeat, []byte("[object Object]"))
}
