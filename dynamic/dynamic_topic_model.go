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

// GetRecommendTopicResult 推荐话题的响应数据。
type GetRecommendTopicResult struct {
	TopicItems []GetRecommendTopicItem `json:"topic_items"` // 话题列表
}

// GetRecommendTopicItem 是推荐话题列表中的一条话题。
type GetRecommendTopicItem struct {
	Discuss          int64  `json:"discuss"`            // 讨论数
	Dynamics         int64  `json:"dynamics"`           // 动态数
	ID               int64  `json:"id"`                 // 话题 id
	JumpURL          string `json:"jump_url"`           // 跳转 URL
	Name             string `json:"name"`               // 话题名
	ShowInteractData bool   `json:"show_interact_data"` // 是否显示互动数据
	View             int64  `json:"view"`               // 浏览数
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

// GetTopicDetailsResult 获取指定话题详细信息的响应数据。
type GetTopicDetailsResult struct {
	ClickAreaCard  GetTopicDetailsClickAreaCard  `json:"click_area_card"` // 点击区域卡片。文档未给出字段
	FunctionalCard GetTopicDetailsFunctionalCard `json:"functional_card"` // 功能卡片
	TopDetails     GetTopicDetailsTopDetails     `json:"top_details"`     // 话题详情
}

// GetTopicDetailsClickAreaCard 是点击区域卡片。文档未给出字段，保留为空结构体以忽略未知字段。
type GetTopicDetailsClickAreaCard struct{}

// GetTopicDetailsFunctionalCard 是话题页的功能卡片集合。
type GetTopicDetailsFunctionalCard struct {
	TrafficCard GetTopicDetailsTrafficCard `json:"traffic_card"` // 引流卡片
}

// GetTopicDetailsTrafficCard 是引流卡片。
type GetTopicDetailsTrafficCard struct {
	BenefitPoint string `json:"benefit_point"` // 奖励点
	CardDesc     string `json:"card_desc"`     // 卡片描述
	IconURL      string `json:"icon_url"`      // 图标 URL
	JumpTitle    string `json:"jump_title"`    // 跳转标题
	JumpURL      string `json:"jump_url"`      // 跳转 URL
	Name         string `json:"name"`          // 名称
}

// GetTopicDetailsTopDetails 是话题页顶部的详情信息。
type GetTopicDetailsTopDetails struct {
	ClosePubLayerEntry    bool                            `json:"close_pub_layer_entry"`   // 是否关闭发布层入口
	HasCreateJurisdiction bool                            `json:"has_create_jurisdiction"` // 是否有创建权限
	OperationContent      GetTopicDetailsOperationContent `json:"operation_content"`       // 操作内容。文档未给出字段
	TopicCreator          GetTopicDetailsCreator          `json:"topic_creator"`           // 话题创建者
	TopicItem             GetTopicDetailsDetailItem       `json:"topic_item"`              // 话题详情
	WordColor             int                             `json:"word_color"`              // 话题字体颜色
}

// GetTopicDetailsOperationContent 是操作内容。文档未给出字段，保留为空结构体以忽略未知字段。
type GetTopicDetailsOperationContent struct{}

// GetTopicDetailsCreator 是话题创建者信息。
type GetTopicDetailsCreator struct {
	Face string `json:"face"` // 头像 URL
	Name string `json:"name"` // 昵称
	UID  int64  `json:"uid"`  // mid (UID)
}

// GetTopicDetailsDetailItem 是话题的完整详情。
type GetTopicDetailsDetailItem struct {
	BackColor        string `json:"back_color"`         // 背景色
	Ctime            int64  `json:"ctime"`              // 创建时间
	Description      string `json:"description"`        // 话题描述
	Discuss          int64  `json:"discuss"`            // 讨论数
	Dynamics         int64  `json:"dynamics"`           // 动态数
	Fav              int64  `json:"fav"`                // 收藏数
	ID               int64  `json:"id"`                 // 话题 id
	JumpURL          string `json:"jump_url"`           // 跳转 URL
	Like             int64  `json:"like"`               // 点赞数
	Name             string `json:"name"`               // 话题名
	Share            int64  `json:"share"`              // 分享数
	SharePic         string `json:"share_pic"`          // 分享图片 URL。似乎就是话题的图标
	ShareURL         string `json:"share_url"`          // 分享的 URL
	ShowInteractData bool   `json:"show_interact_data"` // 是否显示互动数据
	View             int64  `json:"view"`               // 浏览数
}
