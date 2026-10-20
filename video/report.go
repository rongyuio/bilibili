package video

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 视频观看时长上报接口。

// videoHeartbeatWebLocation 是 web 端播放心跳固定的来源标识。
// 该值对应「视频详情页播放器」，接口文档给出的示例取值同为 1315873，不随页面计算。
const videoHeartbeatWebLocation = "1315873"

var (
	// errReportCid 表示上报参数缺少有效的 cid。
	errReportCid = errors.New("参数 Cid 必须大于 0")
	// errReportRealtime 表示上报的本次观看时长无效。
	errReportRealtime = errors.New("参数 Realtime 必须大于 0")
	// errReportVideoDuration 表示上报的视频总时长无效。
	errReportVideoDuration = errors.New("参数 VideoDuration 必须大于 0")
	// errReportPlayedTime 表示播放进度取值无效。
	errReportPlayedTime = errors.New("参数 PlayedTime 不能为负数，仅允许 -1 表示播放完成")
	// errReportAidBvid 表示 avid 与 bvid 都未提供。
	errReportAidBvid = errors.New("参数 Aid 与 Bvid 必须任选一个")
	// errReportBvidLength 表示 bvid 长度不符合 Bv2Av 的要求。
	errReportBvidLength = errors.New("参数 Bvid 必须是 12 位")
	// errMissingCSRF 表示请求 Cookie 快照中缺少 bili_jct。
	errMissingCSRF = errors.New("登录过期：缺少 bili_jct")
	// errMissingDedeUserID 表示请求 Cookie 快照中缺少 DedeUserID。
	errMissingDedeUserID = errors.New("登录过期：缺少 DedeUserID")
)

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
	_, err = bilibili.Execute[any](ctx, c.Client, method, url, param,
		reportVideoWatchTimeHandler(param, aid), c.FillWbi())
	return err
}

// reportVideoWatchTimeAid 校验上报参数并返回稿件 avid。
func reportVideoWatchTimeAid(param ReportVideoWatchTimeParam) (int, error) {
	if param.Cid <= 0 {
		return 0, errReportCid
	}
	if param.Realtime <= 0 {
		return 0, errReportRealtime
	}
	if param.VideoDuration <= 0 {
		return 0, errReportVideoDuration
	}
	if param.PlayedTime < 0 && param.PlayedTime != VideoPlayedComplete {
		return 0, errReportPlayedTime
	}
	if param.Aid > 0 {
		return param.Aid, nil
	}
	if param.Bvid == "" {
		return 0, errReportAidBvid
	}
	// Bv2Av 对长度不为 12 的输入会 panic，必须先挡下来。
	if len(param.Bvid) != 12 {
		return 0, errReportBvidLength
	}
	return bv2Av(param.Bvid), nil
}

// reportVideoWatchTimeHandler 把业务字段从 query 搬进 URL 编码表单，
// 并在 query 上补一份带 w_ 前缀的副本供 WBI 签名。
// 必须在 c.FillWbi() 之前注册：fillWbi 会合并客户端级 query、
// 整体替换 r.QueryParam 完成签名，并清空 Referer。
func reportVideoWatchTimeHandler(param ReportVideoWatchTimeParam, aid int) bilibili.Handler {
	return func(r *resty.Request) error {
		csrf := requestCookieValue(r.Cookies, "bili_jct")
		if csrf == "" {
			return errMissingCSRF
		}
		mid := requestCookieValue(r.Cookies, "DedeUserID")
		if mid == "" {
			return errMissingDedeUserID
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

// requestCookieValue 从请求 Cookie 快照中按名称取值，找不到返回空串。
func requestCookieValue(cookies []*http.Cookie, name string) string {
	for _, cookie := range cookies {
		if cookie != nil && cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
