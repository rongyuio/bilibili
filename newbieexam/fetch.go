package newbieexam

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 新用户答题题目拉取相关接口。响应模型见 fetch_model.go。

// GetNewbieExamBaseQuestion 拉取基础题。
//
// 该接口没有请求参数，需要登录（Cookie SESSDATA 或 APP 认证）。
// 基础题未通过时无法拉取附加题，附加题未通过时无法拉取自选题。
func (c *Client) GetNewbieExamBaseQuestion(ctx context.Context) (*GetNewbieExamBaseQuestionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/base"
	)
	return bilibili.Execute[*GetNewbieExamBaseQuestionResult](ctx, c.Client, method, url, nil)
}

// GetNewbieExamExtraQuestion 拉取附加题。
//
// 该接口没有请求参数，需要登录（Cookie SESSDATA 或 APP 认证）。
func (c *Client) GetNewbieExamExtraQuestion(ctx context.Context) (*GetNewbieExamExtraQuestionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/extra"
	)
	return bilibili.Execute[*GetNewbieExamExtraQuestionResult](ctx, c.Client, method, url, nil)
}

// GetNewbieExamProQuestion 拉取自选题。
//
// 该接口没有请求参数，需要登录（Cookie SESSDATA 或 APP 认证）。
func (c *Client) GetNewbieExamProQuestion(ctx context.Context) (*GetNewbieExamProQuestionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/answer/v4/pro"
	)
	return bilibili.Execute[*GetNewbieExamProQuestionResult](ctx, c.Client, method, url, nil)
}
