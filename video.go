package bilibili

import (
	"context"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// 视频相关接口。响应模型见 video_model.go。

type VideoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
}

// GetVideoDetailInfo 获取视频超详细信息
func (c *Client) GetVideoDetailInfo(ctx context.Context, param VideoParam) (*VideoDetailInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/view/detail"
	)
	return execute[*VideoDetailInfo](ctx, c, method, url, param)
}

// GetVideoRecommendList 获取单视频推荐列表
func (c *Client) GetVideoRecommendList(ctx context.Context, param VideoParam) ([]VideoInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/related"
	)
	return execute[[]VideoInfo](ctx, c, method, url, param)
}

// GetVideoInfo 获取视频详细信息
func (c *Client) GetVideoInfo(ctx context.Context, param VideoParam) (*VideoInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/view"
	)
	return execute[*VideoInfo](ctx, c, method, url, param)
}

// GetVideoDesc 获取视频简介
func (c *Client) GetVideoDesc(ctx context.Context, param VideoParam) (string, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/desc"
	)
	return execute[string](ctx, c, method, url, param)
}

// GetVideoPageList 获取视频分P列表
func (c *Client) GetVideoPageList(ctx context.Context, param VideoParam) ([]VideoPage, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/pagelist"
	)
	return execute[[]VideoPage](ctx, c, method, url, param)
}

// GetVideoTags 获取视频TAG
func (c *Client) GetVideoTags(ctx context.Context, param VideoParam) ([]VideoTag, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/tag/archive/tags"
	)
	return execute[[]VideoTag](ctx, c, method, url, param)
}

type VideoTagParam struct {
	Aid   int `json:"aid"`    // 稿件avid
	TagID int `json:"tag_id"` // tag_id
}

// LikeVideoTag 点赞视频TAG，重复请求为取消
func (c *Client) LikeVideoTag(ctx context.Context, param VideoTagParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/tag/archive/like2"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

// HateVideoTag 点踩视频TAG，重复访问为取消
func (c *Client) HateVideoTag(ctx context.Context, param VideoTagParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/tag/archive/hate2"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type LikeVideoParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件 avid。avid 与 bvid 任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件 bvid。avid 与 bvid 任选一个
	Like int    `json:"like"`                                     // 操作方式。1：点赞。2：取消赞
}

// LikeVideo 点赞视频
func (c *Client) LikeVideo(ctx context.Context, param LikeVideoParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/archive/like"
	)
	_, err := execute[any](ctx, c, method, url, param, fillCsrf(c))
	return err
}

type CoinVideoParam struct {
	Aid        int    `json:"aid,omitempty" request:"query,omitempty"`         // 稿件 avid。avid 与 bvid 任选一个
	Bvid       string `json:"bvid,omitempty" request:"query,omitempty"`        // 稿件 bvid。avid 与 bvid 任选一个
	Multiply   int    `json:"multiply"`                                        // 投币数量。上限为2
	SelectLike int    `json:"select_like,omitempty" request:"query,omitempty"` // 是否附加点赞。0：不点赞。1：同时点赞。默认为0
}

// CoinVideo 投币视频
func (c *Client) CoinVideo(ctx context.Context, param CoinVideoParam) (*CoinVideoResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/coin/add"
	)
	return execute[*CoinVideoResult](ctx, c, method, url, param, fillCsrf(c))
}

type FavourVideoParam struct {
	Rid         int   `json:"rid"`                                               // 稿件 avid
	Type        int   `json:"type"`                                              // 必须为2
	AddMediaIDs []int `json:"add_media_ids,omitempty" request:"query,omitempty"` // 需要加入的收藏夹 mlid。同时添加多个，用,（%2C）分隔
	DelMediaIDs []int `json:"del_media_ids,omitempty" request:"query,omitempty"` // 需要取消的收藏夹 mlid。同时取消多个，用,（%2C）分隔
}

// FavourVideo 收藏视频
func (c *Client) FavourVideo(ctx context.Context, param FavourVideoParam) (*FavourVideoResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/medialist/gateway/coll/resource/deal"
	)
	return execute[*FavourVideoResult](ctx, c, method, url, param, fillCsrf(c))
}

// LikeCoinFavourVideo 一键三连视频
func (c *Client) LikeCoinFavourVideo(ctx context.Context, param VideoParam) (*LikeCoinFavourResult, error) {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/web-interface/archive/like/triple"
	)
	return execute[*LikeCoinFavourResult](ctx, c, method, url, param, fillCsrf(c))
}

type VideoCidParam struct {
	Aid  int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件avid。avid与bvid任选一个
	Bvid string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件bvid。avid与bvid任选一个
	Cid  int    `json:"cid"`                                      // 视频cid。用于选择目标分P
}

// GetVideoOnlineInfo 获取视频在线人数
func (c *Client) GetVideoOnlineInfo(ctx context.Context, param VideoCidParam) (*VideoOnlineInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/online/total"
	)
	return execute[*VideoOnlineInfo](ctx, c, method, url, param)
}

// GetVideoStatusNumber 获取视频状态数视频
func (c *Client) GetVideoStatusNumber(ctx context.Context, param VideoParam) (*VideoStatusNumber, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/archive/stat"
	)
	return execute[*VideoStatusNumber](ctx, c, method, url, param)
}

type GetTopRecommendVideoParam struct {
	FreshType  int `json:"fresh_type,omitempty" request:"query,omitempty"`   // 相关性。默认为3 。 值越大推荐内容越相关
	Version    int `json:"version,omitempty" request:"query,omitempty"`      // web端新旧版本:0为旧版本1为新版本。默认为 0 。1,0分别为新旧web端
	Ps         int `json:"ps,omitempty" request:"query,omitempty"`           // pagesize 单页返回的记录条数默认为10或8。默认为10 。当version为1时默认为8
	FreshIdx   int `json:"fresh_idx,omitempty" request:"query,omitempty"`    // 翻页相关。默认为1 。 与翻页相关
	FreshIdx1H int `json:"fresh_idx_1h,omitempty" request:"query,omitempty"` // 翻页相关。默认为1 。 与翻页相关
}

// GetTopRecommendVideo 获取首页视频推荐列表
func (c *Client) GetTopRecommendVideo(ctx context.Context, param GetTopRecommendVideoParam) (*TopRecommendVideoList, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/web-interface/wbi/index/top/feed/rcmd"
	)
	return execute[*TopRecommendVideoList](ctx, c, method, url, param)
}

type GetVideoCollectionInfoParam struct {
	Mid         int  `json:"mid"`                                              // UP 主 ID
	SeasonID    int  `json:"season_id"`                                        // 视频合集 ID
	SortReverse bool `json:"sort_reverse,omitempty" request:"query,omitempty"` // 未知
	PageNum     int  `json:"page_num,omitempty" request:"query,omitempty"`     // 页码索引
	PageSize    int  `json:"page_size,omitempty" request:"query,omitempty"`    // 单页内容数量
}

// GetVideoCollectionInfo 获取视频合集信息 https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/video/collection.md#%E8%8E%B7%E5%8F%96%E8%A7%86%E9%A2%91%E5%90%88%E9%9B%86%E4%BF%A1%E6%81%AF
func (c *Client) GetVideoCollectionInfo(ctx context.Context, param GetVideoCollectionInfoParam) (*VideoCollectionInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/polymer/web-space/seasons_archives_list"
	)
	return execute[*VideoCollectionInfo](ctx, c, method, url, param)
}

type GetVideoByKeywordsParam struct {
	Mid      int    `json:"mid"`                                           // 用户 mid
	Keywords string `json:"keywords"`                                      // 关键词。可为空, 即获取所有视频
	Ps       int    `json:"ps,omitempty" request:"query,omitempty"`        // 每页视频数。默认为 0, 留空为 20
	Pn       int    `json:"pn,omitempty" request:"query,omitempty"`        // 页码。留空为 1
	Orderby  string `json:"orderby,omitempty" request:"query,omitempty"`   // 排序方式。最新发布: pubdate(默认)。最多播放: views。senddate: 最新发布
	SeriesID int    `json:"series_id,omitempty" request:"query,omitempty"` // 系列 ID。用于过滤结果, 即若某一视频包含在系列内则不返回该视频
}

// GetVideoByKeywords 根据关键词查找视频
//
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/video/collection.md#%E6%A0%B9%E6%8D%AE%E5%85%B3%E9%94%AE%E8%AF%8D%E6%9F%A5%E6%89%BE%E8%A7%86%E9%A2%91
func (c *Client) GetVideoByKeywords(ctx context.Context, param GetVideoByKeywordsParam) (*VideoCollectionByKeywordsInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/series/recArchivesByKeywords"
	)
	return execute[*VideoCollectionByKeywordsInfo](ctx, c, method, url, param)
}

type GetVideoSeriesInfoParam struct {
	Mid        int    `json:"mid"`                                             // UP 主 ID
	SeriesID   int    `json:"series_id"`                                       // 视频合集 ID
	Sort       string `json:"sort,omitempty" request:"query,omitempty"`        // 未知
	Pn         int    `json:"pn,omitempty" request:"query,omitempty"`          // 页码索引
	Ps         int    `json:"ps,omitempty" request:"query,omitempty"`          // 单页内容数量
	CurrentMid int    `json:"current_mid,omitempty" request:"query,omitempty"` // 单页内容数量
}

// GetVideoSeriesInfo 获取视频列表信息（在个人空间里创建的叫做视频列表，在创作中心里创建的叫合集，注意区分）
func (c *Client) GetVideoSeriesInfo(ctx context.Context, param GetVideoSeriesInfoParam) (*VideoCollectionInfo, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/series/archives"
	)
	return execute[*VideoCollectionInfo](ctx, c, method, url, param)
}

type GetVideoStreamParam struct {
	Avid        int    `json:"avid,omitempty" request:"query,omitempty"`         // 稿件 avid。avid 与 bvid 任选一个
	Bvid        string `json:"bvid,omitempty" request:"query,omitempty"`         // 稿件 bvid。avid 与 bvid 任选一个
	Cid         int    `json:"cid"`                                              // 视频 cid
	Qn          int    `json:"qn,omitempty" request:"query,omitempty"`           // 视频清晰度选择。未登录默认 32（480P），登录后默认 64（720P）。含义见 [上表](#qn视频清晰度标识)。DASH 格式时无效
	Fnval       int    `json:"fnval,omitempty" request:"query,omitempty"`        // 视频流格式标识。默认值为1（MP4 格式）。含义见 [ 上表](#fnval视频流格式标识)
	Fnver       int    `json:"fnver,omitempty" request:"query,omitempty"`        // 0
	Fourk       int    `json:"fourk,omitempty" request:"query,omitempty"`        // 是否允许 4K 视频。画质最高 1080P：0（默认）。画 质最高 4K：1
	Session     string `json:"session,omitempty" request:"query,omitempty"`      // 从视频播放页的 HTML 中获取
	Otype       string `json:"otype,omitempty" request:"query,omitempty"`        // 固定为json
	Type        string `json:"type,omitempty" request:"query,omitempty"`         // 目前为空
	Platform    string `json:"platform,omitempty" request:"query,omitempty"`     // pc：web播放（默认值，视频流存在 referer鉴权）。html5：移动端 HTML5 播放（仅支持 MP4 格式，无 referer 鉴权可以直接使用video标签播放）
	HighQuality int    `json:"high_quality,omitempty" request:"query,omitempty"` // 是否高画质。platform=html5时，此值 为1可使画质为1080p
}

// GetVideoStream 获取视频流地址_web端
func (c *Client) GetVideoStream(ctx context.Context, param GetVideoStreamParam) (*GetVideoStreamResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/player/wbi/playurl"
	)
	return execute[*GetVideoStreamResult](ctx, c, method, url, param, c.fillWbi())
}

// videoHeartbeatWebLocation 是 web 端播放心跳固定的来源标识。
// 该值对应「视频详情页播放器」，接口文档给出的示例取值同为 1315873，不随页面计算。
const videoHeartbeatWebLocation = "1315873"

// VideoPlayedComplete 是 played_time 的特殊取值，表示视频已播放完成。
const VideoPlayedComplete = -1

// ReportVideoWatchTimeParam 指定视频观看时长上报参数。
// 字段默认编码进 query，再由内部 handler 移入 URL 编码表单。
type ReportVideoWatchTimeParam struct {
	Aid           int    `json:"aid,omitempty" request:"query,omitempty"`  // 稿件 avid。avid 与 bvid 任选一个
	Bvid          string `json:"bvid,omitempty" request:"query,omitempty"` // 稿件 bvid。avid 与 bvid 任选一个
	Cid           int    `json:"cid"`                                      // 视频 cid，即当前分P，可用 GetVideoInfo 或 GetVideoPageList 获取
	Realtime      int    `json:"realtime"`                                 // 本次上报的观看时长，单位秒，必须大于 0
	PlayedTime    int    `json:"played_time"`                              // 播放进度，单位秒；不应超过视频总时长，VideoPlayedComplete（-1）表示已看完
	VideoDuration int    `json:"video_duration"`                           // 视频总时长，单位秒，可用 GetVideoInfo 的 Duration 获取
}

// ReportVideoWatchTime 上报视频观看时长（web 端播放心跳）。
//
// 需要登录态：mid 取自 Cookie DedeUserID，CSRF 取自 Cookie bili_jct，两者缺失都会返回错误。
// 请求使用 WBI 签名，只对带 w_ 前缀的 query 参数签名，业务参数走 URL 编码表单。
// 上报次数与间隔完全由调用方决定，库不做限制，也不自动重试。
func (c *Client) ReportVideoWatchTime(ctx context.Context, param ReportVideoWatchTimeParam) error {
	const (
		method = resty.MethodPost
		url    = "https://api.bilibili.com/x/click-interface/web/heartbeat"
	)
	aid, err := reportVideoWatchTimeAid(param)
	if err != nil {
		return err
	}
	_, err = execute[any](ctx, c, method, url, param,
		reportVideoWatchTimeHandler(param, aid), c.fillWbi())
	return err
}

// reportVideoWatchTimeAid 校验上报参数并返回稿件 avid。
func reportVideoWatchTimeAid(param ReportVideoWatchTimeParam) (int, error) {
	const root = "ReportVideoWatchTimeParam"
	if param.Cid <= 0 {
		return 0, parameterError(root, "Cid", "cid", "form", "cid 必须大于 0", nil)
	}
	if param.Realtime <= 0 {
		return 0, parameterError(root, "Realtime", "realtime", "form", "realtime 必须大于 0", nil)
	}
	if param.VideoDuration <= 0 {
		return 0, parameterError(root, "VideoDuration", "video_duration", "form", "video_duration 必须大于 0", nil)
	}
	if param.PlayedTime < 0 && param.PlayedTime != VideoPlayedComplete {
		return 0, parameterError(root, "PlayedTime", "played_time", "form", "played_time 不能为负数，仅允许 -1 表示播放完成", nil)
	}
	if param.Aid > 0 {
		return param.Aid, nil
	}
	if param.Bvid == "" {
		return 0, parameterError(root, "Aid/Bvid", "aid", "form", "avid 与 bvid 任选一个", nil)
	}
	// Bv2Av 对长度不为 12 的输入会 panic，必须先挡下来。
	if len(param.Bvid) != 12 {
		return 0, parameterError(root, "Bvid", "bvid", "form", "bvid 必须是 12 位", nil)
	}
	return Bv2Av(param.Bvid), nil
}

// reportVideoWatchTimeHandler 把业务字段从 query 搬进 URL 编码表单，
// 并在 query 上补一份带 w_ 前缀的副本供 WBI 签名。
// 必须在 c.fillWbi() 之前注册：fillWbi 会合并客户端级 query、
// 整体替换 r.QueryParam 完成签名，并清空 Referer。
func reportVideoWatchTimeHandler(param ReportVideoWatchTimeParam, aid int) paramHandler {
	return func(r *resty.Request) error {
		csrf, err := csrfValue(r)
		if err != nil {
			return err
		}
		mid := cookieValue(r.Cookies, "DedeUserID")
		if mid == "" {
			return errors.New("B站登录过期：缺少 DedeUserID")
		}

		aidText := strconv.Itoa(aid)
		realtime := strconv.Itoa(param.Realtime)
		played := strconv.Itoa(param.PlayedTime)
		duration := strconv.Itoa(param.VideoDuration)
		startTs := strconv.FormatInt(time.Now().Unix()-int64(param.Realtime), 10)

		// withParams 已把这些字段写进 query，这里移除，业务参数只出现在表单里。
		for _, key := range []string{"aid", "bvid", "cid", "realtime", "played_time", "video_duration"} {
			r.QueryParam.Del(key)
		}
		r.SetFormData(map[string]string{
			"start_ts":                startTs,
			"mid":                     mid,
			"aid":                     aidText,
			"cid":                     strconv.Itoa(param.Cid),
			"type":                    "3",
			"sub_type":                "0",
			"dt":                      "2",
			"play_type":               "0",
			"realtime":                realtime,
			"played_time":             played,
			"real_played_time":        realtime,
			"refer_url":               "https://www.bilibili.com/",
			"quality":                 strconv.Itoa(VideoQuality720P),
			"is_auto_qn":              "0",
			"video_duration":          duration,
			"last_play_progress_time": played,
			"max_play_progress_time":  duration,
			"outer":                   "0",
			"statistics":              `{"appId":100,"platform":5,"abtest":"","version":""}`,
			"mobi_app":                "web",
			"device":                  "web",
			"platform":                "web",
			"cur_language_vt":         "{}",
			"perfer_type":             "{}",
			"play_mode":               "1",
			"spmid":                   "333.788.0.0",
			"from_spmid":              "333.788.0.0",
			"csrf":                    csrf,
		})
		// WBI 要求每个 query 键只能有一个值；SetQueryParams 内部用 Set，符合该约束。
		r.SetQueryParams(map[string]string{
			"w_start_ts":                startTs,
			"w_mid":                     mid,
			"w_aid":                     aidText,
			"w_dt":                      "2",
			"w_realtime":                realtime,
			"w_played_time":             played,
			"w_real_played_time":        realtime,
			"w_video_duration":          duration,
			"w_last_play_progress_time": played,
			"web_location":              videoHeartbeatWebLocation,
		})
		r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		return nil
	}
}
