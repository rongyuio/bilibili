package bilibili

import (
	"context"
	"reflect"
)

// 本文件导出子包需要的基础设施：请求链路、参数 handler 与签名 helper。
// 业务接口按域拆分到子包（video、user、live 等），子包通过这里的函数
// 复用统一的 context 检查、Cookie 快照、参数编码、错误与解码逻辑。

// Handler 在请求发送前修改请求，用于 CSRF、WBI 签名与固定参数。
type Handler = paramHandler

// Execute 发送请求并把 data 解码到 Out。
//
// 与包内 execute 行为一致：context 作为首参、参数按 request 标签编码、
// 非 200 返回 *HTTPError、业务码非零返回 Error。
func Execute[Out any](ctx context.Context, c *Client, method, endpoint string, in any, handlers ...Handler) (Out, error) {
	return execute[Out](ctx, c, method, endpoint, in, handlers...)
}

// ExecuteRaw 发送请求并返回原始响应体，用于非标准 JSON 信封的接口。
func ExecuteRaw(ctx context.Context, c *Client, method, endpoint string, in any, handlers ...Handler) ([]byte, error) {
	return executeRaw(ctx, c, method, endpoint, in, handlers...)
}

// FillCsrf 把 CSRF 写入 query。
func FillCsrf(c *Client) Handler { return fillCsrf(c) }

// FillFormCsrf 把 CSRF 写入表单。
func FillFormCsrf(c *Client) Handler { return fillFormCsrf(c) }

// FillParam 追加一个固定 query 参数。
func FillParam(key, value string) Handler { return fillParam(key, value) }

// MoveFormParams 把指定 query 字段移入 URL 编码表单。
func MoveFormParams(keys ...string) Handler { return moveFormParams(keys...) }

// FillWbi 对 query 做 WBI 签名。
func (c *Client) FillWbi() Handler { return c.fillWbi() }

// NewDecodeError 构造解码错误，供子包手动解码非标准响应时使用。
//
// body 应为未剪枝的原始响应体，rootType 是解码目标类型，
// path 是 JSON 路径（例如 "$.result"），offset 是值在 body 中的 0 起始偏移。
func NewDecodeError(method, endpoint string, body []byte, rootType reflect.Type, path string, offset int64, err error) *DecodeError {
	return newDecodeError(method, endpoint, body, rootType, path, offset, err)
}
