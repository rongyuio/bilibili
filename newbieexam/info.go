package newbieexam

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 新用户答题信息查询相关接口。响应模型见 info_model.go。

// GetNewbieExamStatus 查询答题状态。
//
// 需要登录（Cookie SESSDATA 或 APP 认证）。
func (c *Client) GetNewbieExamStatus(ctx context.Context) (*GetNewbieExamStatusResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/status"
	)
	return bilibili.Execute[*GetNewbieExamStatusResult](ctx, c.Client, method, url, nil)
}

// GetNewbieExamProTypeList 查询自选题分类。
//
// 需要登录（Cookie SESSDATA 或 APP 认证）。
// 返回的父分类随答题版本（edition）不同而不同：旧版 edition=0 细分子分类，新版 edition=2 每个父分类仅含一个子分类。
func (c *Client) GetNewbieExamProTypeList(ctx context.Context) ([]GetNewbieExamProTypeListItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/pro/type"
	)
	return bilibili.Execute[[]GetNewbieExamProTypeListItem](ctx, c.Client, method, url, nil)
}

// GetNewbieExamResultParam 是查询答题结果的请求参数。
type GetNewbieExamResultParam struct {
	Hid int64 `json:"hid" request:"query"` // 答题会话id
}

// GetNewbieExamResult 查询答题结果。
//
// 认证方式为无；hid 取自查询答题状态接口。
// 返回类型名为 GetNewbieExamResultResult：方法名本身已含 Result，按本包
// 「接口名 + Result」的命名规则拼接而成。
func (c *Client) GetNewbieExamResult(ctx context.Context, param GetNewbieExamResultParam) (*GetNewbieExamResultResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/result"
	)
	return bilibili.Execute[*GetNewbieExamResultResult](ctx, c.Client, method, url, param)
}
