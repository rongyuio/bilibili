package history

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 旧版 web 端视频观看历史记录接口。返回体按接口独立定义。

// GetHistoryLegacyParam 是获取旧版 web 端观看历史记录的请求参数。
type GetHistoryLegacyParam struct {
	Pn int `json:"pn,omitempty" request:"query,omitempty"` // 页码。默认为1
	Ps int `json:"ps,omitempty" request:"query,omitempty"` // 每页项数
}

// GetHistoryLegacyRights 是历史记录条目的视频属性标志。
type GetHistoryLegacyRights struct {
	Bp            int `json:"bp"`              // 是否允许承包
	Elec          int `json:"elec"`            // 是否支持充电
	Download      int `json:"download"`        // 是否允许下载
	Movie         int `json:"movie"`           // 是否电影
	Pay           int `json:"pay"`             // 是否PGC付费
	Hd5           int `json:"hd5"`             // 是否有高码率
	NoReprint     int `json:"no_reprint"`      // 是否显示“禁止转载”标志
	Autoplay      int `json:"autoplay"`        // 是否自动播放
	UgcPay        int `json:"ugc_pay"`         // 是否UGC付费
	IsCooperation int `json:"is_cooperation"`  // 是否为联合投稿
	UgcPayPreview int `json:"ugc_pay_preview"` // 0。作用尚不明确
	NoBackground  int `json:"no_background"`   // 0。作用尚不明确
	CleanMode     int `json:"clean_mode"`      // 0。作用尚不明确
	IsSteinGate   int `json:"is_stein_gate"`   // 是否为互动视频
	Is360         int `json:"is_360"`          // 是否为全景视频
	NoShare       int `json:"no_share"`        // 0。作用尚不明确
	ArcPay        int `json:"arc_pay"`         // 0。作用尚不明确
	FreeWatch     int `json:"free_watch"`      // 0。作用尚不明确
}

// GetHistoryLegacyOwner 是历史记录条目的 UP 主信息。
type GetHistoryLegacyOwner struct {
	Mid  int    `json:"mid"`  // UP主mid
	Name string `json:"name"` // UP主昵称
	Face string `json:"face"` // UP主头像
}

// GetHistoryLegacyStat 是历史记录条目的视频状态数。
type GetHistoryLegacyStat struct {
	Aid        int    `json:"aid"`        // 稿件avid
	View       int    `json:"view"`       // 播放数
	Danmaku    int    `json:"danmaku"`    // 弹幕数
	Reply      int    `json:"reply"`      // 评论数
	Favorite   int    `json:"favorite"`   // 收藏数
	Coin       int    `json:"coin"`       // 投币数
	Share      int    `json:"share"`      // 分享数
	NowRank    int    `json:"now_rank"`   // 当前排名
	HisRank    int    `json:"his_rank"`   // 历史最高排行
	Like       int    `json:"like"`       // 获赞数
	Dislike    int    `json:"dislike"`    // 点踩数。恒为0
	Evaluation string `json:"evaluation"` // 视频评分
	Vt         int    `json:"vt"`         // 作用尚不明确。恒为0
}

// GetHistoryLegacyDimension 是历史记录条目 1P 的分辨率信息。
type GetHistoryLegacyDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetHistoryLegacyPageDimension 是最后观看分 P 的分辨率信息。
type GetHistoryLegacyPageDimension struct {
	Width  int `json:"width"`  // 当前分P 宽度
	Height int `json:"height"` // 当前分P 高度
	Rotate int `json:"rotate"` // 是否将宽高对换。0：正常。1：对换
}

// GetHistoryLegacySeason 是历史记录条目中番剧/影视的剧集信息。
type GetHistoryLegacySeason struct {
	IsFinish      int    `json:"is_finish"`       // 是否完结。0：连载中。1：完结
	NewestEpID    int    `json:"newest_ep_id"`    // 最细一话epid
	NewestEpIndex string `json:"newest_ep_index"` // 最细一话集数
	SeasonID      int    `json:"season_id"`       // 番剧ssid
	SeasonStatus  int    `json:"season_status"`   // ？？？
	SeasonType    int    `json:"season_type"`     // ？？？
	Title         string `json:"title"`           // 番剧标题
	TotalCount    int    `json:"total_count"`     // 总集数
}

// GetHistoryLegacyBangumi 是历史记录条目中的番剧/影视信息。
type GetHistoryLegacyBangumi struct {
	Cover         string                 `json:"cover"`          // 番剧封面图片url
	EpID          int                    `json:"ep_id"`          // 番剧epid
	EpisodeStatus int                    `json:"episode_status"` // ？？？
	Follow        int                    `json:"follow"`         // 0。作用尚不明确
	LongTitle     string                 `json:"long_title"`     // 单集标题
	Season        GetHistoryLegacySeason `json:"season"`         // 剧集信息
	Title         string                 `json:"title"`          // 集数
}

// GetHistoryLegacyCheese 是历史记录条目中的课程信息。
type GetHistoryLegacyCheese struct {
	Cover      string `json:"cover"`       // 课程封面图片url
	LongTitle  string `json:"long_title"`  // 完整标题
	Number     string `json:"number"`      // 本集数
	SeasonID   int    `json:"season_id"`   // 课程ssid
	UpdateInfo string `json:"update_info"` // 总集数
}

// GetHistoryLegacyPage 是历史记录条目中最后观看的分P信息。
type GetHistoryLegacyPage struct {
	Cid         int                           `json:"cid"`          // 视频cid
	Page        int                           `json:"page"`         // 分P页码。实测为1
	From        string                        `json:"from"`         // 视频来源。实测为 vupload
	Part        string                        `json:"part"`         // 分P标题
	Duration    int                           `json:"duration"`     // 分P时长。单位秒
	Vid         string                        `json:"vid"`          // (?) 实测为空
	Weblink     string                        `json:"weblink"`      // (?) 实测为空
	Dimension   GetHistoryLegacyPageDimension `json:"dimension"`    // 分P分辨率信息。含 width / height / rotate
	FirstFrame  string                        `json:"first_frame"`  // 分P首帧图片url
	Ctime       int                           `json:"ctime"`        // 分P创建时间。时间戳
	AsrLanguage string                        `json:"asr_language"` // 字幕语言。实测为 cmn，非必要，实测5项中只有2项有此项
}

// GetHistoryLegacyItem 是旧版 web 端历史记录列表中的一条记录。
type GetHistoryLegacyItem struct {
	Aid          int                       `json:"aid"`           // 稿件avid
	Videos       int                       `json:"videos"`        // 视频分P总数。默认为1
	Tid          int                       `json:"tid"`           // 分区tid
	Tname        string                    `json:"tname"`         // 子分区名称
	Copyright    int                       `json:"copyright"`     // 是否转载。1：原创。2：转载。3：未填写
	Pic          string                    `json:"pic"`           // 视频封面图片url
	Title        string                    `json:"title"`         // 稿件标题
	Pubdate      int                       `json:"pubdate"`       // 稿件发布时间。时间戳
	Ctime        int                       `json:"ctime"`         // 用户提交稿件的时间。时间戳
	Desc         string                    `json:"desc"`          // 视频简介
	State        int                       `json:"state"`         // 视频状态。略，见获取视频详细信息（web端）中的state备注
	Duration     int                       `json:"duration"`      // 视频总计持续时长（所有分P）。单位为秒
	Rights       GetHistoryLegacyRights    `json:"rights"`        // 视频属性标志。略，见获取视频详细信息（web端）中的rights对象
	Owner        GetHistoryLegacyOwner     `json:"owner"`         // 视频UP主信息。略，见获取视频详细信息（web端）中的owner对象
	Stat         GetHistoryLegacyStat      `json:"stat"`          // 视频状态数。略，见获取视频详细信息（web端）中的stat对象
	Dynamic      string                    `json:"dynamic"`       // 视频同步发布的的动态的文字内容。无为空
	Cid          int                       `json:"cid"`           // 视频1P cid
	Dimension    GetHistoryLegacyDimension `json:"dimension"`     // 视频1P分辨率。略，见获取视频详细信息（web端）中的dimension对象
	Bangumi      *GetHistoryLegacyBangumi  `json:"bangumi"`       // 番剧/影视信息。非番剧/影视无此项
	Cheese       *GetHistoryLegacyCheese   `json:"cheese"`        // 课程信息。非课程无此项
	Favorite     bool                      `json:"favorite"`      // 是否已收藏。true：已收藏。false：未收藏
	Type         int                       `json:"type"`          // 视频属性。3：普通视频。4：剧集。10：课程
	SubType      int                       `json:"sub_type"`      // 附视频属性。0：普通视频。1：番剧。2：电影。3：纪录片。4：国创。5：电视剧。7：综艺
	Device       int                       `json:"device"`        // 观看平台代码。1 3 5 7：手机端。2：web端。4 6：pad端。33：TV端。0：其他
	Page         GetHistoryLegacyPage      `json:"page"`          // 最后观看的分P信息
	Count        int                       `json:"count"`         // 分P数。非投稿视频无此项
	Progress     int                       `json:"progress"`      // 观看进度。单位为秒
	ViewAt       int                       `json:"view_at"`       // 观看时间。时间戳
	Kid          int                       `json:"kid"`           // 稿件avid
	Business     string                    `json:"business"`      // 视频类型标识。archive：用户投稿视频。pgc：番剧/影视。cheese：课程
	RedirectLink string                    `json:"redirect_link"` // 重定向url
	Bvid         string                    `json:"bvid"`          // 稿件bvid
}

// GetHistoryLegacy 获取视频观看历史记录（web端旧版）。本接口也可以返回已失效稿件的信息。
func (c *Client) GetHistoryLegacy(ctx context.Context, param GetHistoryLegacyParam) ([]GetHistoryLegacyItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/x/v2/history"
	)
	return bilibili.Execute[[]GetHistoryLegacyItem](ctx, c.Client, method, url, param)
}
