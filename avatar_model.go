package bilibili

// 头像 fallback_layers 渲染树的共享片段。
//
// 整体动态（DynamicAuthorAvatar）、原动态（DynamicOriginalAuthorAvatar）与话题
// （TopicAuthorAvatar）三处头像共用 general_spec 子树；整体动态与原动态还共用
// web_css_style 的四字段形态。其余部分（tags、resource）三处字段集各不相同，
// 保留在各自文件内，避免臆造字段。

// AvatarLayerPosSpec 是头像渲染层的位置规格。
type AvatarLayerPosSpec struct {
	AxisX         float64 `json:"axis_x"`
	AxisY         float64 `json:"axis_y"`
	CoordinatePos int     `json:"coordinate_pos"`
}

// AvatarLayerRenderSpec 是头像渲染层的渲染规格。
type AvatarLayerRenderSpec struct {
	Opacity int `json:"opacity"`
}

// AvatarLayerSizeSpec 是头像渲染层的尺寸规格。
type AvatarLayerSizeSpec struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

// AvatarLayerGeneralSpec 是头像渲染层的通用规格，三个头像模型共用。
type AvatarLayerGeneralSpec struct {
	PosSpec    AvatarLayerPosSpec    `json:"pos_spec"`
	RenderSpec AvatarLayerRenderSpec `json:"render_spec"`
	SizeSpec   AvatarLayerSizeSpec   `json:"size_spec"`
}

// AvatarLayerWebCssStyle 是头像渲染层的 CSS 样式（整体动态与原动态共用；话题仅返回 borderRadius）。
type AvatarLayerWebCssStyle struct {
	BorderRadius    string `json:"borderRadius"`
	BackgroundColor string `json:"background-color,omitempty"`
	Border          string `json:"border,omitempty"`
	BoxSizing       string `json:"boxSizing,omitempty"`
}
