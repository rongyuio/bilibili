package video

import "github.com/rongyuio/bilibili"

// 视频在线人数与状态数接口的响应模型。

// GetVideoOnlineInfoResult 是「获取视频在线人数」的 data 对象。
type GetVideoOnlineInfoResult struct {
	Total      string                       `json:"total"`       // 所有终端总计人数。例如10万+
	Count      string                       `json:"count"`       // web端实时在线人数
	ShowSwitch GetVideoOnlineInfoShowSwitch `json:"show_switch"` // 数据显示控制
	Abtest     GetVideoOnlineInfoAbtest     `json:"abtest"`      // AB测试分组。实测为 {"group":"b"}
}

// GetVideoOnlineInfoShowSwitch 是「获取视频在线人数」show_switch 对象。
type GetVideoOnlineInfoShowSwitch struct {
	Total bool `json:"total"` // 展示所有终端总计人数
	Count bool `json:"count"` // 展示web端实时在线人数
}

// GetVideoOnlineInfoAbtest 是「获取视频在线人数」abtest 对象。
type GetVideoOnlineInfoAbtest struct {
	Group string `json:"group"` // 分组标识。实测为 b
}

// GetVideoStatusNumberResult 是「获取视频状态数」的 data 对象。
type GetVideoStatusNumberResult struct {
	Aid        int                     `json:"aid"`        // 稿件avid
	Bvid       string                  `json:"bvid"`       // 稿件bvid
	View       bilibili.NumberOrString `json:"view"`       // 正常：播放次数(num)。屏蔽："--"(str)
	Danmaku    int                     `json:"danmaku"`    // 弹幕条数
	Reply      int                     `json:"reply"`      // 评论条数
	Favorite   int                     `json:"favorite"`   // 收藏人数
	Coin       int                     `json:"coin"`       // 投币枚数
	Share      int                     `json:"share"`      // 分享次数
	NowRank    int                     `json:"now_rank"`   // 0。作用尚不明确
	HisRank    int                     `json:"his_rank"`   // 历史最高排行
	Like       int                     `json:"like"`       // 获赞次数
	Dislike    int                     `json:"dislike"`    // 0。作用尚不明确
	NoReprint  int                     `json:"no_reprint"` // 禁止转载标志。0：无。1：禁止
	Copyright  int                     `json:"copyright"`  // 版权标志。1：自制。2：转载
	ArgueMsg   string                  `json:"argue_msg"`  // 警告信息。默认为空
	Evaluation string                  `json:"evaluation"` // 视频评分。默认为空
}
