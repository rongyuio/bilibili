package dynamic

// 动态发布相关的话题响应模型。信封中的 code/message/ttl 由库统一处理。
//
// 本包不复用根包或其他子包的响应体：每个接口的 data 层以及所有嵌套对象、
// 数组元素都以接口方法名为前缀独立定义，即使字段完全相同也各自定义，
// 便于后续单个接口的字段变更不影响其他调用。

// SearchPublishTopicResult 搜索发布话题的响应数据。
type SearchPublishTopicResult struct {
	NewTopic              SearchPublishTopicNewTopic `json:"new_topic"`               // 新建话题
	HasCreateJurisdiction bool                       `json:"has_create_jurisdiction"` // 是否有创建权限
	TopicItems            []SearchPublishTopicItem   `json:"topic_items"`             // 话题列表
	RequestID             string                     `json:"request_id"`              // 请求 id
	PageInfo              SearchPublishTopicPageInfo `json:"page_info"`               // 页信息
}

// SearchPublishTopicNewTopic 是搜索关键词对应的新建话题信息。
type SearchPublishTopicNewTopic struct {
	Name string `json:"name"` // 请求参数中的 keywords
}

// SearchPublishTopicItem 是搜索发布话题结果中的一条话题。
type SearchPublishTopicItem struct {
	ID               int64  `json:"id"`                 // 话题 id
	Name             string `json:"name"`               // 话题名
	View             int64  `json:"view"`               // 浏览数
	Discuss          int64  `json:"discuss"`            // 讨论数
	StatDesc         string `json:"stat_desc"`          // 状态描述
	Description      string `json:"description"`        // 话题描述
	ShowInteractData bool   `json:"show_interact_data"` // 是否显示互动数据
}

// SearchPublishTopicPageInfo 是话题搜索的分页信息。
type SearchPublishTopicPageInfo struct {
	Offset  int   `json:"offset"`   // 下次请求可用的偏移，为请求参数中的 offset + page_size
	HasMore *bool `json:"has_more"` // 是否有更多数据。没有更多时该字段可能不存在
}

// GetRecommendSearchTopicResult 推荐搜索话题的响应数据。
// 文档将 topic_items 的元素类型标为 unknown，故保留为 any。
type GetRecommendSearchTopicResult struct {
	TopicItems []any  `json:"topic_items"` // 话题列表。实测为空
	RequestID  string `json:"request_id"`  // 请求 id。传入 keywords 时不为空
}

// CheckTopicExistedResult 检查话题是否存在的响应数据。
type CheckTopicExistedResult struct {
	AlreadyExisted bool                          `json:"already_existed"` // 是否已经存在
	SynonymTopic   CheckTopicExistedSynonymTopic `json:"synonym_topic"`   // 类似话题
}

// CheckTopicExistedSynonymTopic 是与查询话题类似的话题集合。
type CheckTopicExistedSynonymTopic struct {
	TopicItems []CheckTopicExistedSynonymItem `json:"topic_items"` // 话题列表
}

// CheckTopicExistedSynonymItem 是类似话题列表中的一条话题。
type CheckTopicExistedSynonymItem struct {
	Description      string `json:"description"`        // 话题描述
	Discuss          int64  `json:"discuss"`            // 讨论数
	Dynamics         int64  `json:"dynamics"`           // 动态数
	ID               int64  `json:"id"`                 // 话题 id
	JumpURL          string `json:"jump_url"`           // 跳转 URL
	Name             string `json:"name"`               // 话题名
	ShowInteractData bool   `json:"show_interact_data"` // 是否显示互动数据
	View             int64  `json:"view"`               // 浏览数
}
