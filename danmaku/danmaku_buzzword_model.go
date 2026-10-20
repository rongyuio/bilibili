package danmaku

// 名词解释相关响应模型。各接口的返回体独立定义，互不复用。

// GetBuzzwordListPinyin 拼音信息。键为名词中的单个字符，值为该字符的拼音字串。
type GetBuzzwordListPinyin map[string]string

// GetBuzzwordListType 名词类型信息。
type GetBuzzwordListType struct {
	ID         int                   `json:"id"`          // 类型id
	Name       string                `json:"name"`        // 类型名称
	NamePinyin GetBuzzwordListPinyin `json:"name_pinyin"` // 拼音信息。详情接口中可能为null
	Level      int                   `json:"level"`       // 权重？实测为0
}

// GetBuzzwordListItem 名词列表中的名词条目。
type GetBuzzwordListItem struct {
	ID         int                   `json:"id"`          // 名词id。用于查询名词解释详情
	Name       string                `json:"name"`        // 名词
	NamePinyin GetBuzzwordListPinyin `json:"name_pinyin"` // 拼音信息
	Picture    string                `json:"picture"`     // 名词图片url
}

// GetBuzzwordListPage 名词列表的分页信息。
type GetBuzzwordListPage struct {
	Num           int `json:"num"`             // 当前页码
	Size          int `json:"size"`            // 每页项数
	Total         int `json:"total"`           // 总计项数
	Cnt           int `json:"cnt"`             // (?)
	Pe            int `json:"pe"`              // (?)。实测为0
	SegStartMills int `json:"seg_start_mills"` // (?)。实测为0
	SegEndMills   int `json:"seg_end_mills"`   // (?)。实测为0
	SegmentRules  any `json:"segment_rules"`   // (?)。实测为null
}

// GetBuzzwordListResult 拉取名词解释列表的响应数据。
type GetBuzzwordListResult struct {
	Type      GetBuzzwordListType   `json:"type"`      // 类型信息
	Buzzwords []GetBuzzwordListItem `json:"buzzwords"` // 词语列表
	Page      GetBuzzwordListPage   `json:"page"`      // 页面信息
}

// GetBuzzwordDetailPinyin 拼音信息。键为名词中的单个字符，值为该字符的拼音字串。
type GetBuzzwordDetailPinyin map[string]string

// GetBuzzwordDetailType 名词类型信息。
type GetBuzzwordDetailType struct {
	ID         int                     `json:"id"`          // 类型id
	Name       string                  `json:"name"`        // 类型名称
	NamePinyin GetBuzzwordDetailPinyin `json:"name_pinyin"` // 拼音信息。详情接口中可能为null
	Level      int                     `json:"level"`       // 权重？实测为0
}

// GetBuzzwordDetailUpperOfficial 名词相关UP主的认证信息。
type GetBuzzwordDetailUpperOfficial struct {
	Role  int    `json:"role"`  // 成员认证级别
	Title string `json:"title"` // 成员认证名。无为空
	Desc  string `json:"desc"`  // 成员认证备注。无为空
	Type  int    `json:"type"`  // 成员认证类型。-1：无。0：有
}

// GetBuzzwordDetailUpper 名词相关的UP主信息。
type GetBuzzwordDetailUpper struct {
	ID        int                            `json:"id"`         // UP主mid
	Recommend string                         `json:"recommend"`  // (?)。实测为空
	Name      string                         `json:"name"`       // UP主昵称
	Face      string                         `json:"face"`       // UP主头像url
	Official  GetBuzzwordDetailUpperOfficial `json:"official"`   // UP主认证信息
	Followed  bool                           `json:"followed"`   // 是否关注。需要登录(Cookie或APP)
	FansCount int                            `json:"fans_count"` // 粉丝数
}

// GetBuzzwordDetailUgcVideo 名词相关的ugc视频。
type GetBuzzwordDetailUgcVideo struct {
	Aid      int    `json:"aid"`      // 稿件avid
	Cid      int    `json:"cid"`      // 视频cid
	Name     string `json:"name"`     // 视频标题
	Link     string `json:"link"`     // 空
	Play     int    `json:"play"`     // 播放数
	Danmaku  int    `json:"danmaku"`  // 弹幕数
	Duration int    `json:"duration"` // 视频时长。单位为秒
	Cover    string `json:"cover"`    // 视频封面url
}

// GetBuzzwordDetailOgvVideo 名词相关的pgc视频。
type GetBuzzwordDetailOgvVideo struct {
	SeasonID   int    `json:"season_id"`   // 剧集ssid
	Name       string `json:"name"`        // 剧集标题
	Play       int    `json:"play"`        // 播放数
	Followed   int    `json:"followed"`    // 追剧/番数
	IsFollowed bool   `json:"is_followed"` // 是否追剧/番
	NewEp      string `json:"new_ep"`      // 最新一集信息
	Cover      string `json:"cover"`       // 剧集封面url
	Episodes   any    `json:"episodes"`    // 实测为null
	FollowInfo any    `json:"follow_info"` // 追剧信息。文档未给出字段
}

// GetBuzzwordDetailContributeTag 推荐创作信息中的创作跳转按钮。
type GetBuzzwordDetailContributeTag struct {
	Title string `json:"title"` // 按钮文案
	Link  string `json:"link"`  // 跳转uri。跳转至必剪或客户端创作入口
}

// GetBuzzwordDetailGuides 名词的推荐创作信息。
type GetBuzzwordDetailGuides struct {
	Description    string                           `json:"description"`     // 创作推介文案
	ContributeTags []GetBuzzwordDetailContributeTag `json:"contribute_tags"` // 创作跳转按钮信息
}

// GetBuzzwordDetailRelatedItem 推荐名词列表中的名词条目。
type GetBuzzwordDetailRelatedItem struct {
	ID         int                      `json:"id"`          // 名词id。用于查询名词解释详情
	Name       string                   `json:"name"`        // 名词
	NamePinyin *GetBuzzwordDetailPinyin `json:"name_pinyin"` // 实测为null
	Picture    string                   `json:"picture"`     // 空
}

// GetBuzzwordDetailTopic 名词的关联话题信息。
type GetBuzzwordDetailTopic struct {
	TopicID   int    `json:"topic_id"`   // 关联话题id
	TopicName string `json:"topic_name"` // 关联话题名称
}

// GetBuzzwordDetailResult 查询名词解释详情的响应数据。接口无数据时 data 为 null。
type GetBuzzwordDetailResult struct {
	ID               int                            `json:"id"`                // 名词id
	Name             string                         `json:"name"`              // 名词
	Description      string                         `json:"description"`       // 名词解释
	Picture          string                         `json:"picture"`           // 名词图片url
	Types            []GetBuzzwordDetailType        `json:"types"`             // 父分类列表。该名词可能隶属多个父分类
	RelatedWords     []string                       `json:"related_words"`     // 近义词/相关词
	Uppers           []GetBuzzwordDetailUpper       `json:"uppers"`            // 相关UP列表。该名词(梗)出处。可能为null
	RelatedUgc       []GetBuzzwordDetailUgcVideo    `json:"related_ugc"`       // 相关ugc视频列表。该名词(梗)出处。可能为null
	RelatedOgv       []GetBuzzwordDetailOgvVideo    `json:"related_ogv"`       // 相关pgc视频列表。该名词(梗)出处。可能为null
	Guides           GetBuzzwordDetailGuides        `json:"guides"`            // 推荐创作信息
	RelatedBuzzwords []GetBuzzwordDetailRelatedItem `json:"related_buzzwords"` // 推荐名词列表。可能为null
	Topic            GetBuzzwordDetailTopic         `json:"topic"`             // 关联话题信息
	ShowDynamic      bool                           `json:"show_dynamic"`      // 是否展示话题动态聚合
}
