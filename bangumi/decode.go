package bangumi

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/rongyuio/bilibili"
)

// resultEnvelope 是番剧类接口的根对象，业务数据位于 result 而非 data。
type resultEnvelope struct {
	Code    int             `json:"code"`    // 返回值。0：成功
	Message string          `json:"message"` // 错误信息
	Result  json.RawMessage `json:"result"`  // 数据本体。可能缺失或为 null
}

// executeResult 发送结果位于 result 字段的请求并解码。
//
// 行为与根包 Execute 对齐：context 检查、参数编码、非 200 状态码与业务码
// 均返回对应错误；result 缺失或为 null 时返回 nil。
func executeResult[T any](ctx context.Context, c *Client, method, endpoint string, param any, handlers ...bilibili.Handler) (*T, error) {
	body, err := bilibili.ExecuteRaw(ctx, c.Client, method, endpoint, param, handlers...)
	if err != nil {
		return nil, err
	}
	var envelope resultEnvelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, bilibili.NewDecodeError(method, endpoint, body, reflect.TypeFor[resultEnvelope](), "$", 0, err)
	}
	if envelope.Code != 0 {
		return nil, fmt.Errorf("%s %s: %w", method, endpoint, bilibili.Error{Code: envelope.Code, Message: envelope.Message})
	}
	if len(envelope.Result) == 0 || string(envelope.Result) == "null" {
		// result 为 null 表示接口没有返回数据，不是错误。
		return nil, nil //nolint:nilnil
	}
	var result T
	if err := json.Unmarshal(envelope.Result, &result); err != nil {
		return nil, bilibili.NewDecodeError(method, endpoint, envelope.Result, reflect.TypeFor[T](), "$.result", 0, err)
	}
	return &result, nil
}
