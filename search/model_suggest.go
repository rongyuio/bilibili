package search

// 搜索建议响应模型。

// GetSearchSuggestResult 搜索建议结果。
type GetSearchSuggestResult struct {
	Tag []GetSearchSuggestTag `json:"tag"` // 建议关键词列表，按相关程度与热度排序，最多 10 项
}

// GetSearchSuggestTag 单条搜索建议关键词。
type GetSearchSuggestTag struct {
	Value       string `json:"value"`        // 关键词内容
	Ref         int    `json:"ref"`          // 作用尚不明确，实测为 0
	Name        string `json:"name"`         // 显示内容，带有 <em class="suggest_high_light"> 的 XML 标签
	Spid        int    `json:"spid"`         // 作用尚不明确，实测为 5
	Type        string `json:"type"`         // 类型，实测为空
	Term        string `json:"term"`         // 搜索词，实测与 value 相同
	ItemFeature string `json:"item_feature"` // 作用尚不明确，实测为空
}
