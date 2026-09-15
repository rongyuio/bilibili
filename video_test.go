package bilibili

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

func TestReportVideoWatchTimeAid(t *testing.T) {
	valid := ReportVideoWatchTimeParam{Cid: 22222, Realtime: 3600, PlayedTime: 120, VideoDuration: 600}

	cases := []struct {
		name    string
		mutate  func(*ReportVideoWatchTimeParam)
		wantAid int
		wantErr bool
	}{
		{name: "aid 有效", mutate: func(p *ReportVideoWatchTimeParam) { p.Aid = 12345 }, wantAid: 12345},
		{name: "bvid 有效", mutate: func(p *ReportVideoWatchTimeParam) { p.Bvid = Av2Bv(54321) }, wantAid: 54321},
		{name: "aid 与 bvid 都缺", mutate: func(*ReportVideoWatchTimeParam) {}, wantErr: true},
		{name: "aid 为负且无 bvid", mutate: func(p *ReportVideoWatchTimeParam) { p.Aid = -1 }, wantErr: true},
		{name: "bvid 长度不足时不 panic", mutate: func(p *ReportVideoWatchTimeParam) { p.Bvid = "BV1xx" }, wantErr: true},
		{name: "bvid 为空字符串", mutate: func(p *ReportVideoWatchTimeParam) { p.Bvid = "" }, wantErr: true},
		{name: "cid 非法", mutate: func(p *ReportVideoWatchTimeParam) { p.Aid = 1; p.Cid = 0 }, wantErr: true},
		{name: "realtime 非法", mutate: func(p *ReportVideoWatchTimeParam) { p.Aid = 1; p.Realtime = 0 }, wantErr: true},
		{name: "video_duration 非法", mutate: func(p *ReportVideoWatchTimeParam) { p.Aid = 1; p.VideoDuration = 0 }, wantErr: true},
		{name: "played_time 为负", mutate: func(p *ReportVideoWatchTimeParam) { p.Aid = 1; p.PlayedTime = -1 }, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			param := valid
			tc.mutate(&param)

			aid, err := reportVideoWatchTimeAid(param)
			if !tc.wantErr {
				if err != nil {
					t.Fatalf("期望成功，实际错误: %v", err)
				}
				if aid != tc.wantAid {
					t.Fatalf("aid = %d，期望 %d", aid, tc.wantAid)
				}
				return
			}
			if err == nil {
				t.Fatalf("期望返回错误，实际 aid = %d", aid)
			}
			var paramErr *ParamError
			if !errors.As(err, &paramErr) {
				t.Fatalf("期望返回 *ParamError，实际 %T", err)
			}
		})
	}
}

func TestReportVideoWatchTimeHandlerRequiresLogin(t *testing.T) {
	param := ReportVideoWatchTimeParam{Aid: 12345, Cid: 22222, Realtime: 3600, PlayedTime: 120, VideoDuration: 600}

	cases := []struct {
		name    string
		cookies []*http.Cookie
	}{
		{name: "缺少 bili_jct", cookies: []*http.Cookie{{Name: "DedeUserID", Value: "10001"}}},
		{name: "缺少 DedeUserID", cookies: []*http.Cookie{{Name: "bili_jct", Value: "csrf-value"}}},
		{name: "两个都缺", cookies: nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := resty.New().R().SetCookies(tc.cookies)
			if err := withParams(r, param); err != nil {
				t.Fatalf("withParams 失败: %v", err)
			}
			if err := reportVideoWatchTimeHandler(param, param.Aid)(r); err == nil {
				t.Fatal("期望返回错误，实际为 nil")
			}
		})
	}
}

func TestReportVideoWatchTimeHandlerLayout(t *testing.T) {
	const (
		mid      = "10001"
		csrf     = "csrf-value"
		aid      = 12345
		realtime = 3600
	)
	param := ReportVideoWatchTimeParam{
		Aid:           aid,
		Cid:           22222,
		Realtime:      realtime,
		PlayedTime:    120,
		VideoDuration: 600,
	}

	r := resty.New().R().SetCookies([]*http.Cookie{
		{Name: "DedeUserID", Value: mid},
		{Name: "bili_jct", Value: csrf},
	})
	if err := withParams(r, param); err != nil {
		t.Fatalf("withParams 失败: %v", err)
	}

	before := time.Now().Unix()
	if err := reportVideoWatchTimeHandler(param, aid)(r); err != nil {
		t.Fatalf("handler 失败: %v", err)
	}
	after := time.Now().Unix()

	// WBI 只签名 query，且要求每个键只有一个值。
	wantQuery := map[string]string{
		"w_mid":                     mid,
		"w_aid":                     strconv.Itoa(aid),
		"w_dt":                      "2",
		"w_realtime":                strconv.Itoa(realtime),
		"w_played_time":             "120",
		"w_real_played_time":        strconv.Itoa(realtime),
		"w_video_duration":          "600",
		"w_last_play_progress_time": "120",
		"web_location":              videoHeartbeatWebLocation,
	}
	if len(r.QueryParam) != len(wantQuery)+1 { // +1 为 w_start_ts
		t.Fatalf("query 键数量 = %d，期望 %d；实际 %v", len(r.QueryParam), len(wantQuery)+1, r.QueryParam)
	}
	for key, want := range wantQuery {
		values, ok := r.QueryParam[key]
		if !ok {
			t.Fatalf("query 缺少键 %s", key)
		}
		if len(values) != 1 {
			t.Fatalf("query 键 %s 有 %d 个值，WBI 要求单值", key, len(values))
		}
		if values[0] != want {
			t.Fatalf("query %s = %q，期望 %q", key, values[0], want)
		}
	}

	// 业务字段不应留在 query 里。
	for _, key := range []string{"aid", "bvid", "cid", "realtime", "played_time", "video_duration"} {
		if _, ok := r.QueryParam[key]; ok {
			t.Fatalf("业务键 %s 不应出现在 query 中", key)
		}
	}

	startText := r.QueryParam.Get("w_start_ts")
	start, err := strconv.ParseInt(startText, 10, 64)
	if err != nil {
		t.Fatalf("解析 w_start_ts 失败: %v", err)
	}
	if start < before-int64(realtime) || start > after-int64(realtime) {
		t.Fatalf("w_start_ts = %d，期望落在 [%d, %d]", start, before-int64(realtime), after-int64(realtime))
	}
	if got := r.FormData.Get("start_ts"); got != startText {
		t.Fatalf("表单 start_ts = %q，与 query 的 %q 不一致", got, startText)
	}

	wantForm := map[string]string{
		"mid":                     mid,
		"aid":                     strconv.Itoa(aid),
		"cid":                     "22222",
		"type":                    "3",
		"sub_type":                "0",
		"dt":                      "2",
		"play_type":               "0",
		"realtime":                strconv.Itoa(realtime),
		"played_time":             "120",
		"real_played_time":        strconv.Itoa(realtime),
		"refer_url":               "https://www.bilibili.com/",
		"quality":                 "64",
		"is_auto_qn":              "0",
		"video_duration":          "600",
		"last_play_progress_time": "120",
		"max_play_progress_time":  "600",
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
	}
	for key, want := range wantForm {
		if got := r.FormData.Get(key); got != want {
			t.Fatalf("表单 %s = %q，期望 %q", key, got, want)
		}
	}
	if got := r.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
		t.Fatalf("Content-Type = %q，期望 application/x-www-form-urlencoded", got)
	}
}
