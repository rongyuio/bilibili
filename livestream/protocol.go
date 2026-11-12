package livestream

// 直播弹幕信息流协议常量。

const (
	// 帧头固定 16 字节：包长、帧头长、协议版本、操作码、sequence。
	danmakuHeaderSize = 16
	danmakuSequence   = 1

	// 协议版本。
	danmakuProtoPlain     = 0 // JSON 明文
	danmakuProtoHeartbeat = 1 // 心跳与认证响应
	danmakuProtoZlib      = 2 // zlib 压缩后的多包
	danmakuProtoBrotli    = 3 // brotli 压缩后的多包

	// 操作码。
	danmakuOpHeartbeat      = 2
	danmakuOpHeartbeatReply = 3
	danmakuOpMessage        = 5
	danmakuOpAuth           = 7
	danmakuOpAuthReply      = 8

	// 默认连接配置。认证声明 brotli，与网页端一致；服务端仍按帧内 protover 下发。
	danmakuDefaultAuthProto = danmakuProtoBrotli
	danmakuHeartbeatPeriod  = 30 // 秒
	danmakuJoinTimeout      = 10 // 秒
	danmakuReadIdleTimeout  = 75 // 秒
	danmakuReconnectDelay   = 3  // 秒
	danmakuMaxPacketSize    = 16 << 20
	danmakuWriteTimeout     = 10 // 秒
	danmakuReadBufferSize   = 4096
	danmakuEventBufferSize  = 64

	// 认证回复中需要刷新 token 或降级公共节点的业务码。
	danmakuAuthCodeTokenExpired = -101
	danmakuAuthCodeRiskControl  = -352 // 部分节点返回 +352，判断时两个都算

	// 服务端未下发可用节点、或触发风控时使用的公共节点。
	danmakuFallbackHost = "broadcastlv.chat.bilibili.com"

	// 防止服务端持续返回 -101 或 -352 时无限重连。
	danmakuMaxTokenRefresh = 3
	danmakuMaxRiskRetries  = 3

	danmakuUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0"
)
