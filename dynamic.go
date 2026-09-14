package bilibili

import (
	"context"
	"io"

	"github.com/go-resty/resty/v2"
)

// 动态相关接口。响应模型见 dynamic_model.go。

type SearchDynamicAtParam struct {
	UID     int    `json:"uid"`     // 自己的uid
	Keyword string `json:"keyword"` // 搜索关键字
}

// SearchDynamicAt 根据关键字搜索用户(at别人时的填充列表)
func (c *Client) SearchDynamicAt(ctx context.Context, param SearchDynamicAtParam) (*SearchDynamicAtResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_mix/v1/dynamic_mix/at_search"
	)
	return execute[*SearchDynamicAtResult](ctx, c, method, url, param)
}

type GetDynamicRepostDetailParam struct {
	DynamicID int `json:"dynamic_id"`                                 // 动态id
	Offset    int `json:"offset,omitempty" request:"query,omitempty"` // 偏移量
}

// GetDynamicRepostDetail 获取动态转发列表
//
// 见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/basicInfo.md
func (c *Client) GetDynamicRepostDetail(ctx context.Context, param GetDynamicRepostDetailParam) (*DynamicRepostDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_repost/v1/dynamic_repost/repost_detail"
	)
	return execute[*DynamicRepostDetail](ctx, c, method, url, param)
}

type GetDynamicLikeListParam struct {
	DynamicID int64 `json:"dynamic_id"`                             // 动态id
	Pn        int64 `json:"pn,omitempty" request:"query,omitempty"` // 页码
	Ps        int64 `json:"ps,omitempty" request:"query,omitempty"` // 每页数量。该值不得大于20
}

// GetDynamicLikeList 获取动态点赞列表
func (c *Client) GetDynamicLikeList(ctx context.Context, param GetDynamicLikeListParam) (*DynamicLikeList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_like/v1/dynamic_like/spec_item_likes"
	)
	return execute[*DynamicLikeList](ctx, c, method, url, param)
}

type GetDynamicLiveUserListParam struct {
	Size int `json:"size,omitempty" request:"query,omitempty"` // 每页显示数。默认为10
}

// GetDynamicLiveUserList 获取正在直播的已关注者
func (c *Client) GetDynamicLiveUserList(ctx context.Context, param GetDynamicLiveUserListParam) (*DynamicLiveUserList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/w_live_users"
	)
	return execute[*DynamicLiveUserList](ctx, c, method, url, param)
}

type GetDynamicUpListParam struct {
	TeenagersMode int `json:"teenagers_mode,omitempty" request:"query,omitempty"` // 是否开启青少年模式。否：0。是：1
}

// GetDynamicUpList 获取发布新动态的已关注者
func (c *Client) GetDynamicUpList(ctx context.Context, param GetDynamicUpListParam) (*DynamicUpList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/w_dyn_uplist"
	)
	return execute[*DynamicUpList](ctx, c, method, url, param)
}

type RemoveDynamicParam struct {
	DynamicID int `json:"dynamic_id"` // 动态id
}

// RemoveDynamic 删除动态
func (c *Client) RemoveDynamic(ctx context.Context, param RemoveDynamicParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/rm_dynamic"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type GetDynamicDetailParam struct {
	DynamicID int `json:"dynamic_id"` // 动态id
}

// GetDynamicDetail 获取特定动态卡片信息
func (c *Client) GetDynamicDetail(ctx context.Context, param GetDynamicDetailParam) (*DynamicDetail, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/get_dynamic_detail"
	)
	return execute[*DynamicDetail](ctx, c, method, url, param)
}

// GetDynamicPortal 获取最近更新UP主列表（其实就是获取自己的动态门户）
func (c *Client) GetDynamicPortal(ctx context.Context) (*DynamicPortal, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/portal"
	)
	return execute[*DynamicPortal](ctx, c, method, url, nil)
}

// UploadDynamicBfsParam 指定图片动态上传参数。文件本身通过 UploadDynamicBfs 的 file 参数传入。
type UploadDynamicBfsParam struct {
	Category string `json:"category" request:"query"` // 图片分类
}

// UploadDynamicBfs 为图片动态上传图片。
//
// 文件字段交由 Resty 生成 multipart boundary；CSRF 从请求 Cookie 快照填入 query。
func (c *Client) UploadDynamicBfs(ctx context.Context, param UploadDynamicBfsParam, fileName string, file io.Reader) (url string, size Size, err error) {
	const (
		method   = resty.MethodPost
		endpoint = "https://api.bilibili.com/x/dynamic/feed/draw/upload_bfs"
	)
	data, err := executeRequest[*UploadDynamicBfsResult](c, c.newRequest(ctx), method, endpoint, param,
		fillCsrf(c), func(r *resty.Request) error {
			r.SetFileReader("file_up", fileName, file)
			return nil
		})
	if err != nil {
		return "", Size{}, err
	}
	return data.ImageURL, Size{Width: data.ImageWidth, Height: data.ImageHeight}, nil
}

type CreateDynamicParam struct {
	DynamicID       int          `json:"dynamic_id"`                                            // 0
	Type            int          `json:"type"`                                                  // 4
	Rid             int          `json:"rid"`                                                   // 0
	Content         string       `json:"content"`                                               // 动态内容
	UpChooseComment int          `json:"up_choose_comment,omitempty" request:"query,omitempty"` // 0
	UpCloseComment  int          `json:"up_close_comment,omitempty" request:"query,omitempty"`  // 0
	Extension       string       `json:"extension,omitempty" request:"query,omitempty"`         // 位置信息，参考 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/dynamic/publish.md
	AtUIDs          string       `json:"at_uids,omitempty" request:"query,omitempty"`           // 动态中 at 到的用户的 uid。使用逗号,分隔
	Ctrl            []FormatCtrl `json:"ctrl,omitempty" request:"query,omitempty"`              // 特殊格式控制 (如 at 别人时的蓝字体和链接)
}

// CreateDynamic 发表纯文本动态
func (c *Client) CreateDynamic(ctx context.Context, param CreateDynamicParam) (*CreateDynamicResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/create"
	)
	return execute[*CreateDynamicResult](ctx, c, method, url, param, fillCsrf(c))
}

type GetUserSpaceDynamicParam struct {
	Offset         string `json:"offset"`          // 分页偏移量
	HostMid        string `json:"host_mid"`        // 用户UID
	TimezoneOffset int    `json:"timezone_offset"` // -480
	Features       string `json:"features"`        // itemOpusStyle
}

// GetUserSpaceDynamic 获取用户空间动态，mid就是用户UID，无需登录。
//
// 返回结构较为繁琐，见 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/dynamic/space.md
func (c *Client) GetUserSpaceDynamic(ctx context.Context, param GetUserSpaceDynamicParam) (*DynamicInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space"
	)
	return execute[*DynamicInfo](ctx, c, method, url, param)
}
