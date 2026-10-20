package activity

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// GetActivitySubjectInfoParam 是获取活动主题信息的参数。
type GetActivitySubjectInfoParam struct {
	Sid  int    `json:"sid" request:"query"`                      // 活动 id
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 来源视频 bvid
}

// GetActivitySubjectInfoResult 是活动主题信息的 data 对象。
type GetActivitySubjectInfoResult struct {
	ID         int    `json:"id"`            // 活动 id
	Oid        int    `json:"oid"`           // 0
	Type       int    `json:"type"`          // 13
	State      int    `json:"state"`         // 1
	Stime      int    `json:"stime"`         // 开始时间。UNIX 秒级时间戳
	Etime      int    `json:"etime"`         // 结束时间。UNIX 秒级时间戳
	Ctime      int    `json:"ctime"`         // 创建时间。UNIX 秒级时间戳
	Mtime      int    `json:"mtime"`         // 修改时间。UNIX 秒级时间戳
	Name       string `json:"name"`          // 活动名称
	ActURL     string `json:"act_url"`       // 活动链接
	Lstime     int    `json:"lstime"`        // 作用尚不明确
	Letime     int    `json:"letime"`        // 作用尚不明确
	Cover      string `json:"cover"`         // 封面图片
	Dic        string `json:"dic"`           // 简介
	H5Cover    string `json:"h5_cover"`      // H5 封面
	AndroidURL string `json:"android_url"`   // Android 端活动链接
	IosURL     string `json:"ios_url"`       // iOS 端活动链接
	ChildSids  string `json:"child_sids"`    // 子活动 id?
	Calendar   string `json:"calendar"`      // 日历。空 或 {}
	Lid        *int   `json:"lid,omitempty"` // 作用尚不明确。仅在传入 bvid 时存在
}

// GetActivitySubjectInfo 获取活动主题信息。
func (c *Client) GetActivitySubjectInfo(ctx context.Context, param GetActivitySubjectInfoParam) (*GetActivitySubjectInfoResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/activity/subject/info"
	)
	return bilibili.Execute[*GetActivitySubjectInfoResult](ctx, c.Client, method, url, param)
}

// GetActivityPageListParam 是获取活动列表的参数。
type GetActivityPageListParam struct {
	Plat string `json:"plat" request:"query"`                     // 活动平台类型。可选范围 [1, 3]，以半角逗号分隔，默认 1,3
	Mold int    `json:"mold,omitempty" request:"query,omitempty"` // 0。作用尚不明确
	HTTP int    `json:"http,omitempty" request:"query,omitempty"` // 3。作用尚不明确
	Pn   int    `json:"pn,omitempty" request:"query,omitempty"`   // 目标页码。默认为 1
	Ps   int    `json:"ps,omitempty" request:"query,omitempty"`   // 每页条数。默认为 15
}

// GetActivityPageListResult 是活动列表接口的 data 对象。
type GetActivityPageListResult struct {
	List  []GetActivityPageListItem `json:"list"`  // 活动列表
	Num   int                       `json:"num"`   // 当前页码
	Size  int                       `json:"size"`  // 每页条数
	Total int                       `json:"total"` // 总条数
}

// GetActivityPageListItem 是活动列表中的一项。
type GetActivityPageListItem struct {
	ID       int    `json:"id"`        // 活动 ID
	State    int    `json:"state"`     // 1。作用尚不明确
	Stime    int64  `json:"stime"`     // 开始时间。UNIX 秒级时间戳
	Etime    int64  `json:"etime"`     // 结束时间。UNIX 秒级时间戳
	Ctime    int64  `json:"ctime"`     // 创建时间。UNIX 秒级时间戳，可能为 0
	Mtime    int64  `json:"mtime"`     // 修改时间。UNIX 秒级时间戳，可能为 0
	Name     string `json:"name"`      // 活动名称
	Author   string `json:"author"`    // 空。作用尚不明确
	PcURL    string `json:"pc_url"`    // 空。作用尚不明确
	Rank     int    `json:"rank"`      // 0。作用尚不明确
	H5URL    string `json:"h5_url"`    // 活动链接
	PcCover  string `json:"pc_cover"`  // 空。作用尚不明确
	H5Cover  string `json:"h5_cover"`  // 活动封面
	PageName string `json:"page_name"` // 页面名称
	Plat     int    `json:"plat"`      // 活动平台类型。即 URL 中 plat 参数
	Desc     string `json:"desc"`      // 活动描述
	Click    int    `json:"click"`     // 0。作用尚不明确
	Type     int    `json:"type"`      // 0。作用尚不明确
	Mold     int    `json:"mold"`      // 0。作用尚不明确
	Series   int    `json:"series"`    // 0。作用尚不明确
	Dept     int    `json:"dept"`      // 0。作用尚不明确
	ReplyID  int    `json:"reply_id"`  // 0。作用尚不明确
	TpID     int    `json:"tp_id"`     // 0。作用尚不明确
	Ptime    int64  `json:"ptime"`     // 0。作用尚不明确
	Catalog  int    `json:"catalog"`   // 0。作用尚不明确
	Creator  string `json:"creator"`   // 空。作用尚不明确
	SpmID    string `json:"spm_id"`    // 空。作用尚不明确
}

// GetActivityPageList 获取活动列表。
func (c *Client) GetActivityPageList(ctx context.Context, param GetActivityPageListParam) (*GetActivityPageListResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/activity/page/list"
	)
	return bilibili.Execute[*GetActivityPageListResult](ctx, c.Client, method, url, param)
}
