package livestream

import (
	"io"
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
