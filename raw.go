package bilibili

import (
	"context"
	"fmt"
	"net/http"
)

// executeRaw 发送请求并返回原始响应体，用于响应不是标准 code/message/data
// JSON 信封的接口（protobuf、XML、SVG、JS 等）。错误处理与 execute 保持一致：
// 参数编码错误带方法与方法路径前缀，网络错误保留原始错误链，非 200 返回 *HTTPError。
//
// 调用方负责解析响应体；Cookie 合并与统一请求头仍由请求链路处理。
func executeRaw(ctx context.Context, c *Client, method, endpoint string, in any, handlers ...paramHandler) ([]byte, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	r := c.newRequest(ctx)
	if err := withParams(r, in); err != nil {
		return nil, fmt.Errorf("%s %s: %w", method, safeEndpoint(endpoint), err)
	}
	for _, handler := range handlers {
		if err := handler(r); err != nil {
			return nil, err
		}
	}
	resp, err := c.sendRaw(r, method, endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, newHTTPError(method, endpoint, resp.StatusCode())
	}
	return resp.Body(), nil
}
