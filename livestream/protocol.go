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

	// 默认连接配置。
	danmakuDefaultAuthProto = danmakuProtoZlib
	danmakuHeartbeatPeriod  = 30 // 秒
	danmakuMaxPacketSize    = 16 << 20
	danmakuWriteTimeout     = 10 // 秒
	danmakuReadBufferSize   = 4096
	danmakuEventBufferSize  = 64

	danmakuUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0"
)
