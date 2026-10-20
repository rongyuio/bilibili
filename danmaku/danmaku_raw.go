package danmaku

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 弹幕原始响应接口。protobuf 与 XML 由调用方自行解析：
// protobuf 定义见接口文档仓库的 grpc_api/bilibili/community/service/dm/v1/dm.proto。

// danmakuOidParam 是仅带 oid 的查询参数，供只接受 cid 的 XML 接口使用。
type danmakuOidParam struct {
	Oid int64 `json:"oid" request:"query"` // 视频 cid
}

// GetDanmakuXML 获取实时弹幕（XML，web 端旧接口）。
//
// 响应经 deflate 压缩，调用方需要自行解压后按 XML 解析。
func (c *Client) GetDanmakuXML(ctx context.Context, cid int64) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v1/dm/list.so"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, danmakuOidParam{Oid: cid})
}

// GetDanmakuXMLByCid 通过 cid 直接获取实时弹幕 XML。
//
// 与 GetDanmakuXML 效果相同，响应同样经 deflate 压缩。
func (c *Client) GetDanmakuXMLByCid(ctx context.Context, cid int64) ([]byte, error) {
	const method = resty.MethodGet
	return bilibili.ExecuteRaw(ctx, c.Client, method, fmt.Sprintf("https://comment.bilibili.com/%d.xml", cid), nil)
}

// GetDanmakuProtoParam 指定要拉取的弹幕分包。
type GetDanmakuProtoParam struct {
	Type         int   `json:"type" request:"query"`                              // 弹幕类。1：视频弹幕。2：漫画弹幕
	Oid          int64 `json:"oid" request:"query"`                               // 视频 cid
	Pid          int   `json:"pid,omitempty" request:"query,omitempty"`           // 稿件 avid
	SegmentIndex int   `json:"segment_index,omitempty" request:"query,omitempty"` // 分包。6min 一包
	PullMode     int   `json:"pull_mode,omitempty" request:"query,omitempty"`     // 作用尚不明确
	Ps           int   `json:"ps,omitempty" request:"query,omitempty"`            // 作用尚不明确
	Pe           int   `json:"pe,omitempty" request:"query,omitempty"`            // 作用尚不明确
}

// GetDanmakuProto 获取实时弹幕（protobuf，web 端）。
func (c *Client) GetDanmakuProto(ctx context.Context, param GetDanmakuProtoParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/web/seg.so"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
}

// GetDanmakuProtoWBI 获取实时弹幕（protobuf，web 端 WBI 签名版）。
func (c *Client) GetDanmakuProtoWBI(ctx context.Context, param GetDanmakuProtoParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/wbi/web/seg.so"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param, c.FillWbi())
}

// GetDanmakuProtoApp 获取实时弹幕（protobuf，APP 端）。
func (c *Client) GetDanmakuProtoApp(ctx context.Context, param GetDanmakuProtoParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/list/seg.so"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
}

// GetDanmakuViewProtoParam 指定弹幕元数据的稿件。
type GetDanmakuViewProtoParam struct {
	Type int   `json:"type" request:"query"`                    // 弹幕类。1：视频弹幕
	Oid  int64 `json:"oid" request:"query"`                     // 视频 cid
	Pid  int   `json:"pid,omitempty" request:"query,omitempty"` // 稿件 avid
}

// GetDanmakuViewProto 获取弹幕个人配置、互动弹幕与 BAS（代码）弹幕专包（protobuf，web 端）。
func (c *Client) GetDanmakuViewProto(ctx context.Context, param GetDanmakuViewProtoParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/web/view"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
}

// GetHistoryDanmakuDatesParam 指定要查询的历史弹幕月份。
type GetHistoryDanmakuDatesParam struct {
	Type  int    `json:"type" request:"query"`  // 弹幕类。固定为 1
	Oid   int64  `json:"oid" request:"query"`   // 视频 cid
	Month string `json:"month" request:"query"` // 查询目标年月。YYYY-MM
}

// GetHistoryDanmakuDates 查询指定月份中存在历史弹幕的日期。
// 无弹幕时返回空切片（响应 data 为 null）。
func (c *Client) GetHistoryDanmakuDates(ctx context.Context, param GetHistoryDanmakuDatesParam) ([]string, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/history/index"
	)
	return bilibili.Execute[[]string](ctx, c.Client, method, url, param)
}

// GetHistoryDanmakuParam 指定要拉取的历史弹幕日期。
type GetHistoryDanmakuParam struct {
	Type int    `json:"type" request:"query"` // 弹幕类。固定为 1
	Oid  int64  `json:"oid" request:"query"`  // 视频 cid
	Date string `json:"date" request:"query"` // 弹幕日期。YYYY-MM-DD
}

// GetHistoryDanmakuProto 获取历史弹幕（protobuf）。
func (c *Client) GetHistoryDanmakuProto(ctx context.Context, param GetHistoryDanmakuParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/web/history/seg.so"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
}

// GetHistoryDanmakuXML 获取历史弹幕（XML）。
//
// 响应经 deflate 压缩，调用方需要自行解压。
func (c *Client) GetHistoryDanmakuXML(ctx context.Context, param GetHistoryDanmakuParam) ([]byte, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/dm/history"
	)
	return bilibili.ExecuteRaw(ctx, c.Client, method, url, param)
}
