package livestream

import (
	"errors"
	"io"

	"github.com/gorilla/websocket"
)

// danmakuReadResult 是一次异步读取的结果。
type danmakuReadResult struct {
	data []byte
	err  error
}

// readAllLimited 读取一帧，超出大小上限时拒绝，避免服务端下发异常包导致无界内存占用。
func readAllLimited(reader io.Reader) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(reader, danmakuMaxPacketSize+1))
	if err != nil {
		return nil, err
	}
	if len(data) > danmakuMaxPacketSize {
		return nil, errDanmakuPacketTooLarge
	}
	return data, nil
}

// readDanmakuPacket 把阻塞读取交给独立 goroutine，使心跳不依赖服务端是否下发消息。
func readDanmakuPacket(conn *websocket.Conn, result chan<- danmakuReadResult) {
	messageType, reader, err := conn.NextReader()
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			result <- danmakuReadResult{err: errDanmakuClosed}
			return
		}
		result <- danmakuReadResult{err: err}
		return
	}
	if messageType != websocket.BinaryMessage {
		result <- danmakuReadResult{err: errors.New("live danmaku expected binary message")}
		return
	}
	data, err := readAllLimited(reader)
	result <- danmakuReadResult{data: data, err: err}
}
