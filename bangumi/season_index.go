package bangumi

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// 剧集索引相关接口。响应模型见 model_season_index.go。

// GetBangumiSeasonIndexParam 获取剧集索引的参数。
type GetBangumiSeasonIndexParam struct {
	SeasonType         int    `json:"season_type" request:"query"`                              // 类别。番剧：1，电影：2，纪录片：3，国创：4，电视剧：5，综艺：7
	Type               int    `json:"type" request:"query"`                                     // 未知作用。固定为 1
	Page               int    `json:"page,omitempty" request:"query,omitempty"`                 // 页码。默认为 1
	PageSize           int    `json:"pagesize,omitempty" request:"query,omitempty"`             // 每页项数。定义域 [1,∞)，超出最大项数则返回最大项数
	SeasonVersion      int    `json:"season_version,omitempty" request:"query,omitempty"`       // 类型。全部：-1，正片：1，电影：2，其他：3
	SpokenLanguageType int    `json:"spoken_language_type,omitempty" request:"query,omitempty"` // 配音。全部：-1，原声：1，中文配音：2
	Area               string `json:"area,omitempty" request:"query,omitempty"`                 // 地区代码列表，如日本：2，中国：1,6,7。-1 代表全部
	IsFinish           int    `json:"is_finish,omitempty" request:"query,omitempty"`            // 状态。全部：-1，完结：1，连载：0
	Copyright          string `json:"copyright,omitempty" request:"query,omitempty"`            // 版权。全部：-1，独家：3，其他：1,2,4
	SeasonStatus       string `json:"season_status,omitempty" request:"query,omitempty"`        // 付费类型。全部：-1，免费：1，付费：2,6，大会员：4,6
	SeasonMonth        int    `json:"season_month,omitempty" request:"query,omitempty"`         // 季度。全部：-1，一月：1，四月：4，七月：7，十月：10
	Year               string `json:"year,omitempty" request:"query,omitempty"`                 // 年份，适用于国创和番剧。以左闭右开区间表示，如 [2025,2026) 代表 2025 年，-1 代表全部
	ReleaseDate        string `json:"release_date,omitempty" request:"query,omitempty"`         // 年份，适用于电影、电视剧和纪录片。精确到秒，如 [2025-01-01 00:00:00,2026-01-01 00:00:00)
	StyleID            int    `json:"style_id,omitempty" request:"query,omitempty"`             // 风格。全部：-1，特定风格则值为该风格 id，如原创：10010
	ProducerID         int    `json:"producer_id,omitempty" request:"query,omitempty"`          // 出品方。全部：-1，特定出品方则值为该出品方 id，如央视：4
	Order              int    `json:"order,omitempty" request:"query,omitempty"`                // 排序指标。更新时间：0，弹幕数量：1，播放数量：2，追剧人数：3，最高评分：4，开播时间：5，上映时间：6
	Sort               int    `json:"sort,omitempty" request:"query,omitempty"`                 // 排序方向。增序：1，降序：0，默认为降序
	St                 int    `json:"st,omitempty" request:"query,omitempty"`                   // 未知作用。值与 season_type 相同
}

// GetBangumiSeasonIndex 获取剧集索引。
func (c *Client) GetBangumiSeasonIndex(ctx context.Context, param GetBangumiSeasonIndexParam) (*GetBangumiSeasonIndexResult, error) {
	const (
		method = resty.MethodGet
		url    = "https://api.bilibili.com/pgc/season/index/result"
	)
	return bilibili.Execute[*GetBangumiSeasonIndexResult](ctx, c.Client, method, url, param)
}
