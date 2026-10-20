package creativecenter

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 创作中心统计与数据接口。文档见 creativecenter/statistics&data.md。
// 统计与数据次日中午 12 点刷新。
//
// 各接口的返回体独立定义，即使字段相同也不复用。

// GetUpVideoStatResult 是 UP 主视频状态数据的 data 对象。
type GetUpVideoStatResult struct {
	IncCoin    int `json:"inc_coin"`    // 新增投币数
	IncElec    int `json:"inc_elec"`    // 新增充电数
	IncFav     int `json:"inc_fav"`     // 新增收藏数
	IncLike    int `json:"inc_like"`    // 新增点赞数
	IncShare   int `json:"inc_share"`   // 新增分享数
	IncrClick  int `json:"incr_click"`  // 新增播放数
	IncrDm     int `json:"incr_dm"`     // 新增弹幕数
	IncrFans   int `json:"incr_fans"`   // 新增粉丝数
	IncrReply  int `json:"incr_reply"`  // 新增评论数
	TotalClick int `json:"total_click"` // 总计播放数
	TotalCoin  int `json:"total_coin"`  // 总计投币数
	TotalDm    int `json:"total_dm"`    // 总计弹幕数
	TotalElec  int `json:"total_elec"`  // 总计充电数
	TotalFans  int `json:"total_fans"`  // 总计粉丝数
	TotalFav   int `json:"total_fav"`   // 总计收藏数
	TotalLike  int `json:"total_like"`  // 总计点赞数
	TotalReply int `json:"total_reply"` // 总计评论数
	TotalShare int `json:"total_share"` // 总计分享数
}

// GetUpVideoStat 获取 UP 主视频状态数据
func (c *Client) GetUpVideoStat(ctx context.Context) (*GetUpVideoStatResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/index/stat"
	)
	return bilibili.Execute[*GetUpVideoStatResult](ctx, c.Client, method, url, nil)
}

// GetArchiveDiagnoseCompareParam 是获取新版 UP 主视频数据比较的请求参数。
type GetArchiveDiagnoseCompareParam struct {
	T    int64 `json:"t,omitempty" request:"query,omitempty"` // 时间戳
	Size int   `json:"size" request:"query,default=5"`        // 比较最近的 N 条视频。默认为 5
}

// GetArchiveDiagnoseCompareResult 是获取新版 UP 主视频数据比较的 data 对象。
type GetArchiveDiagnoseCompareResult struct {
	List []GetArchiveDiagnoseCompareItem `json:"list"` // 最近的视频
}

// GetArchiveDiagnoseCompareItem 是获取新版 UP 主视频数据比较 list 数组中的对象。
//
// 各项比率字段一般为百分比，小数点后保留两位：100 代表 1%，10000 代表 100%。
type GetArchiveDiagnoseCompareItem struct {
	Aid                       int64  `json:"aid"`                           // av 号
	Bvid                      string `json:"bvid"`                          // bv 号
	Cover                     string `json:"cover"`                         // 封面 URL
	Title                     string `json:"title"`                         // 标题
	Pubtime                   int64  `json:"pubtime"`                       // 发布时间。UNIX 秒级时间戳
	Duration                  int    `json:"duration"`                      // 视频长度（秒）
	Play                      int    `json:"play"`                          // 播放数
	Vt                        int    `json:"vt"`                            // 作用尚不明确
	Like                      int    `json:"like"`                          // 点赞数
	Comment                   int    `json:"comment"`                       // 评论数
	Dm                        int    `json:"dm"`                            // 弹幕数
	Fav                       int    `json:"fav"`                           // 收藏数
	Coin                      int    `json:"coin"`                          // 投币数
	Share                     int    `json:"share"`                         // 分享数
	FullPlayRatio             int    `json:"full_play_ratio"`               // 完播比。用户平均在百分之多少退出
	PlayViewerRate            int    `json:"play_viewer_rate"`              // 游客播放数。这个视频有多少是游客播放
	ActiveFansRate            int    `json:"active_fans_rate"`              // 粉丝观看率。多少粉丝看了这个视频
	ActiveFansMed             int    `json:"active_fans_med"`               // 作用尚不明确
	TmRate                    int    `json:"tm_rate"`                       // 封标点击率
	TmRateMed                 int    `json:"tm_rate_med"`                   // 你自己平均封标点击率
	TmFanSimiRateMed          int    `json:"tm_fan_simi_rate_med"`          // 同类 UP 粉丝封标点击率
	TmViewerSimiRateMed       int    `json:"tm_viewer_simi_rate_med"`       // 同类 UP 游客封标点击率
	TmFanRate                 int    `json:"tm_fan_rate"`                   // 粉丝封标点击率
	TmViewerRate              int    `json:"tm_viewer_rate"`                // 游客封标点击率
	TmPassRate                int    `json:"tm_pass_rate"`                  // 封标点击率超过 n% 同类稿件
	TmFanPassRate             int    `json:"tm_fan_pass_rate"`              // 粉丝封标点击率超过 n% 同类稿件
	TmViewerPassRate          int    `json:"tm_viewer_pass_rate"`           // 游客封标点击率超过 n% 同类稿件
	CrashRate                 int    `json:"crash_rate"`                    // 3 秒退出率
	CrashRateMed              int    `json:"crash_rate_med"`                // 作用尚不明确
	CrashFanSimiRateMed       int    `json:"crash_fan_simi_rate_med"`       // 同类 UP 粉丝 3 秒退出率
	CrashViewerSimiRateMed    int    `json:"crash_viewer_simi_rate_med"`    // 同类 UP 游客 3 秒退出率
	CrashFanRate              int    `json:"crash_fan_rate"`                // 粉丝 3 秒退出率
	CrashViewerRate           int    `json:"crash_viewer_rate"`             // 游客 3 秒退出率
	InteractRate              int    `json:"interact_rate"`                 // 互动率
	InteractRateMed           int    `json:"interact_rate_med"`             // 作用尚不明确
	InteractFanSimiRateMed    int    `json:"interact_fan_simi_rate_med"`    // 同类 UP 粉丝互动率
	InteractViewerSimiRateMed int    `json:"interact_viewer_simi_rate_med"` // 同类 UP 游客互动率
	InteractFanRate           int    `json:"interact_fan_rate"`             // 粉丝互动率
	InteractViewerRate        int    `json:"interact_viewer_rate"`          // 游客互动率
	AvgPlayTime               int    `json:"avg_play_time"`                 // 平均播放时间。实测总是 0
	TotalNewAttentionCnt      int    `json:"total_new_attention_cnt"`       // 涨粉
	PlayTransFanRate          int    `json:"play_trans_fan_rate"`           // 播转粉率
	PlayTransFanRateMed       int    `json:"play_trans_fan_rate_med"`       // 其他 UP 平均播转粉率
}

// GetArchiveDiagnoseCompare 比较新版 UP 主视频数据
func (c *Client) GetArchiveDiagnoseCompare(ctx context.Context, param GetArchiveDiagnoseCompareParam) (*GetArchiveDiagnoseCompareResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/archive_diagnose/compare"
	)
	return bilibili.Execute[*GetArchiveDiagnoseCompareResult](ctx, c.Client, method, url, param)
}

// GetUpArticleStatResult 是 UP 主专栏状态数据的 data 对象。
type GetUpArticleStatResult struct {
	View      int `json:"view"`       // 总计阅读数
	Reply     int `json:"reply"`      // 总计评论数
	Like      int `json:"like"`       // 总计点赞数
	Coin      int `json:"coin"`       // 总计投币数
	Fav       int `json:"fav"`        // 总计收藏数
	Share     int `json:"share"`      // 总计分享数
	IncrView  int `json:"incr_view"`  // 新增阅读数
	IncrReply int `json:"incr_reply"` // 新增评论数
	IncrLike  int `json:"incr_like"`  // 新增点赞数
	IncrCoin  int `json:"incr_coin"`  // 新增投币数
	IncrFav   int `json:"incr_fav"`   // 新增收藏数
	IncrShare int `json:"incr_share"` // 新增分享数
}

// GetUpArticleStat 获取 UP 主专栏状态数据
func (c *Client) GetUpArticleStat(ctx context.Context) (*GetUpArticleStatResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/article"
	)
	return bilibili.Execute[*GetUpArticleStatResult](ctx, c.Client, method, url, nil)
}

// GetVideoDataIncrementTrendParam 是获取视频数据增量趋势的请求参数。
type GetVideoDataIncrementTrendParam struct {
	Type int `json:"type" request:"query"` // 目标数据类型。1：播放；2：弹幕；3：评论；4：分享；5：投币；6：收藏；7：充电；8：点赞
}

// GetVideoDataIncrementTrendItem 是获取视频数据增量趋势 data 数组中的对象。
type GetVideoDataIncrementTrendItem struct {
	DateKey  int64 `json:"date_key"`  // 对应时间。时间戳，前一天的 8:00
	TotalInc int   `json:"total_inc"` // 增加数量。意义由请求的数据类型决定
}

// GetVideoDataIncrementTrend 获取视频数据增量趋势。数据为前 30 天
func (c *Client) GetVideoDataIncrementTrend(ctx context.Context, param GetVideoDataIncrementTrendParam) ([]GetVideoDataIncrementTrendItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/pandect"
	)
	return bilibili.Execute[[]GetVideoDataIncrementTrendItem](ctx, c.Client, method, url, param)
}

// GetArticleDataIncrementTrendParam 是获取专栏数据增量趋势的请求参数。
type GetArticleDataIncrementTrendParam struct {
	Type int `json:"type" request:"query"` // 目标数据类型。1：阅读；2：评论；3：分享；4：投币；5：收藏；6：点赞
}

// GetArticleDataIncrementTrendItem 是获取专栏数据增量趋势 data 数组中的对象。
type GetArticleDataIncrementTrendItem struct {
	DateKey  int64 `json:"date_key"`  // 对应时间。时间戳，前一天的 8:00
	TotalInc int   `json:"total_inc"` // 增加数量。意义由请求的数据类型决定
}

// GetArticleDataIncrementTrend 获取专栏数据增量趋势。数据为前 30 天
func (c *Client) GetArticleDataIncrementTrend(ctx context.Context, param GetArticleDataIncrementTrendParam) ([]GetArticleDataIncrementTrendItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/article/thirty"
	)
	return bilibili.Execute[[]GetArticleDataIncrementTrendItem](ctx, c.Client, method, url, param)
}

// GetArchiveSourceSurveyParam 是获取稿件操作来源占比情况的请求参数。
type GetArchiveSourceSurveyParam struct {
	Type int `json:"type" request:"query"` // 目标数据类型。1：播放；2：弹幕；3：评论；4：分享；5：投币；6：收藏；7：充电；8：点赞
}

// GetArchiveSourceSurveyItem 是稿件操作来源占比情况中按日期索引的对象。
type GetArchiveSourceSurveyItem struct {
	ArcInc   []GetArchiveSourceSurveyArc `json:"arc_inc"`   // 稿件情况
	TotalInc int                         `json:"total_inc"` // 总计增长情况
	TypeRank map[string]int              `json:"type_rank"` // 分区排名情况。键为分区名
}

// GetArchiveSourceSurveyArc 是稿件操作来源占比情况 arc_inc 数组中的对象。
type GetArchiveSourceSurveyArc struct {
	Aid         int64  `json:"aid"`         // 稿件 avid
	Bvid        string `json:"bvid"`        // 稿件 bvid
	Daytime     int64  `json:"daytime"`     // 统计时间。UNIX 秒级时间戳
	Incr        int    `json:"incr"`        // 稿件增长情况数
	Interactive int    `json:"interactive"` // 0。作用尚不明确
	Ptime       int64  `json:"ptime"`       // 稿件发布时间。UNIX 秒级时间戳
	Title       string `json:"title"`       // 稿件标题
}

// GetArchiveSourceSurvey 获取稿件操作来源占比情况。数据为上一天的，返回值的键为日期（YYYYMMDD）
func (c *Client) GetArchiveSourceSurvey(ctx context.Context, param GetArchiveSourceSurveyParam) (map[string]GetArchiveSourceSurveyItem, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/survey"
	)
	return bilibili.Execute[map[string]GetArchiveSourceSurveyItem](ctx, c.Client, method, url, param)
}

// GetPlaySourceResult 是播放来源占比情况（平台及方式）的 data 对象。
type GetPlaySourceResult struct {
	PageSource     GetPlaySourcePageSource `json:"page_source"`     // 播放方式情况
	PlayProportion GetPlaySourceProportion `json:"play_proportion"` // 播放平台情况
}

// GetPlaySourcePageSource 是播放方式情况。
type GetPlaySourcePageSource struct {
	Dynamic      int `json:"dynamic"`       // 通过动态
	Other        int `json:"other"`         // 其他方式
	RelatedVideo int `json:"related_video"` // 通过推荐列表
	Search       int `json:"search"`        // 通过搜索
	Space        int `json:"space"`         // 空间列表播放
	Tenma        int `json:"tenma"`         // 天马（APP 推荐信息流）来源
}

// GetPlaySourceProportion 是播放平台情况。
type GetPlaySourceProportion struct {
	Android int `json:"android"` // 安卓端
	H5      int `json:"h5"`      // 移动端 h5 页面
	Ios     int `json:"ios"`     // ios 端
	Out     int `json:"out"`     // 站外
	PC      int `json:"pc"`      // 电脑版网页
}

// GetPlaySource 获取播放来源占比情况（平台及方式）
func (c *Client) GetPlaySource(ctx context.Context) (*GetPlaySourceResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/playsource"
	)
	return bilibili.Execute[*GetPlaySourceResult](ctx, c.Client, method, url, nil)
}

// GetPlayDistributionResult 是播放分布情况（粉丝与路人）的 data 对象。
type GetPlayDistributionResult struct {
	Period     GetPlayDistributionPeriod     `json:"period"`      // 提示信息
	ViewerArea GetPlayDistributionViewerArea `json:"viewer_area"` // 播放地区情况
	ViewerBase GetPlayDistributionViewerBase `json:"viewer_base"` // 播放数据情况
}

// GetPlayDistributionPeriod 是提示信息。
type GetPlayDistributionPeriod struct {
	ModuleOne   string `json:"module_one"`   // 作用尚不明确
	ModuleTwo   string `json:"module_two"`   // 作用尚不明确
	ModuleThree string `json:"module_three"` // 作用尚不明确
	ModuleFour  string `json:"module_four"`  // 作用尚不明确
}

// GetPlayDistributionViewerArea 是播放地区情况。键为行政区名。
type GetPlayDistributionViewerArea struct {
	Fan    map[string]int `json:"fan"`     // 粉丝播放地区情况
	NotFan map[string]int `json:"not_fan"` // 路人播放地区情况
}

// GetPlayDistributionViewerBase 是播放数据情况。
type GetPlayDistributionViewerBase struct {
	Fan    GetPlayDistributionBaseData `json:"fan"`     // 粉丝播放数据情况
	NotFan GetPlayDistributionBaseData `json:"not_fan"` // 路人播放数据情况
}

// GetPlayDistributionBaseData 是粉丝或路人的播放数据。
type GetPlayDistributionBaseData struct {
	Male         int `json:"male"`           // 男性播放数
	Female       int `json:"female"`         // 女性播放数
	AgeOne       int `json:"age_one"`        // 0-16 岁播放数
	AgeTwo       int `json:"age_two"`        // 16-25 岁播放数
	AgeThree     int `json:"age_three"`      // 25-40 岁播放数
	AgeFour      int `json:"age_four"`       // 40+ 岁播放数
	PlatPC       int `json:"plat_pc"`        // pc 网页端播放数
	PlatH5       int `json:"plat_h5"`        // 移动 h5 端播放数
	PlatOut      int `json:"plat_out"`       // 站外端播放数
	PlatIos      int `json:"plat_ios"`       // ios 端播放数
	PlatAndroid  int `json:"plat_android"`   // 安卓端播放数
	PlatOtherApp int `json:"plat_other_app"` // 其他端播放数
}

// GetPlayDistribution 获取播放分布情况（粉丝与路人）
func (c *Client) GetPlayDistribution(ctx context.Context) (*GetPlayDistributionResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://member.bilibili.com/x/web/data/base"
	)
	return bilibili.Execute[*GetPlayDistributionResult](ctx, c.Client, method, url, nil)
}
