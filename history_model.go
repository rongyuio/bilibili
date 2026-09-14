package bilibili

// 历史记录与稍后再看响应模型。

type HistoryCursor struct {
	Max      int    `json:"max"`      // 最后一项目标 id。**见请求参数**
	ViewAt   int    `json:"view_at"`  // 最后一项时间节点。时间戳
	Business string `json:"business"` // 最后一项业务类型。**见请求参数**
	Ps       int    `json:"ps"`       // 每页项数
}

type HistoryTab struct {
	Type string `json:"type"` // 类型
	Name string `json:"name"` // 类型名
}

type HistoryDetail struct {
	Oid      int    `json:"oid"`      // 目标id。稿件视频&剧集（当business=archive或business=pgc时）：稿件avid。直播（当business=live时）：直播间id。文章（当business=article时）：文章cvid。文集（当business=article-list时）：文集rlid
	Epid     int    `json:"epid"`     // 剧集epid。仅用于剧集
	Bvid     string `json:"bvid"`     // 稿件bvid。仅用于稿件视频
	Page     int    `json:"page"`     // 观看到的视频分P数。仅用于稿件视频
	Cid      int    `json:"cid"`      // 观看到的对象id。稿件视频&剧集（当business=archive或business=pgc时）：视频cid。文集（当business=article-list时）：文章cvid
	Part     string `json:"part"`     // 观看到的视频分 P 标题。仅用于稿件视频
	Business string `json:"business"` // 业务类型。**见请求参数**
	Dt       int    `json:"dt"`       // 记录查看的平台代码。1 3 5 7：手机端。2：web端。4 6：pad端。33：TV端。0：其他
}

type HistoryList struct {
	Title      string        `json:"title"`       // 条目标题
	LongTitle  string        `json:"long_title"`  // 条目副标题
	Cover      string        `json:"cover"`       // 条目封面图 url。用于专栏以外的条目
	Covers     []string      `json:"covers"`      // 条目封面图组。仅用于专栏
	Uri        string        `json:"uri"`         // 重定向 url。仅用于剧集和直播
	History    HistoryDetail `json:"history"`     // 条目详细信息
	Videos     int           `json:"videos"`      // 视频分 P 数目。仅用于稿件视频
	AuthorName string        `json:"author_name"` // UP 主昵称
	AuthorFace string        `json:"author_face"` // UP 主头像 url
	AuthorMid  int           `json:"author_mid"`  // UP 主 mid
	ViewAt     int           `json:"view_at"`     // 查看时间。时间戳
	Progress   int           `json:"progress"`    // 视频观看进度。单位为秒。用于稿件视频或剧集
	Badge      string        `json:"badge"`       // 角标文案。稿件视频 / 剧集 / 笔记
	ShowTitle  string        `json:"show_title"`  // 分 P 标题。用于稿件视频或剧集
	Duration   int           `json:"duration"`    // 视频总时长。用于稿件视频或剧集
	Current    string        `json:"current"`     // (?)
	Total      int           `json:"total"`       // 总计分集数。仅用于剧集
	NewDesc    string        `json:"new_desc"`    // 最新一话 / 最新一 P 标识。用于稿件视频或剧集
	IsFinish   int           `json:"is_finish"`   // 是否已完结。仅用于剧集。0：未完结。1：已完结
	IsFav      int           `json:"is_fav"`      // 是否收藏。0：未收藏。1：已收藏
	Kid        int           `json:"kid"`         // 条目目标 id。**详细内容见参数**
	TagName    string        `json:"tag_name"`    // 子分区名。用于稿件视频和直播
	LiveStatus int           `json:"live_status"` // 直播状态。仅用于直播。0：未开播。1：已开播
}

type HistoryInfo struct {
	Cursor HistoryCursor `json:"cursor"` // 历史记录页面信息
	Tab    []HistoryTab  `json:"tab"`    // 历史记录筛选类型
	List   []HistoryList `json:"list"`   // 分段历史记录列表
}

type ToViewInfo struct {
	Count int            `json:"count"` // 稍后再看视频数
	List  []ToViewDetail `json:"list"`  // 稍后再看视频列表
}

type ToViewDetail struct {
	Aid       int         `json:"aid"`       // 稿件avid
	Videos    int         `json:"videos"`    // 稿件分P总数。默认为1
	Tid       int         `json:"tid"`       // 分区tid
	Tname     string      `json:"tname"`     // 子分区名称
	Copyright int         `json:"copyright"` // 是否转载。1：原创。2：转载
	Pic       string      `json:"pic"`       // 稿件封面图片url
	Title     string      `json:"title"`     // 稿件标题
	Pubdate   int         `json:"pubdate"`   // 稿件发布时间。时间戳
	Ctime     int         `json:"ctime"`     // 用户提交稿件的时间。时间戳
	Desc      string      `json:"desc"`      // 视频简介
	State     int         `json:"state"`     // 视频状态。略，见[获取视频详细信息（web端）](../video/info.md#获取视频详细信息（web端）)中的state备注
	Duration  int         `json:"duration"`  // 稿件总时长（所有分P）。单位为秒
	Rights    VideoRights `json:"rights"`    // 稿件属性标志。略，见[获取视频详细信息（web端）](../video/info.md#获取视频详细信息（web端）)中的rights对象
	Owner     Owner       `json:"owner"`     // 稿件UP主信息。略，见[获取视频详细信息（web端）](../video/info.md#获取视频详细信息 （web端）)中的owner对象
	Stat      VideoStat   `json:"stat"`      // 稿件状态数。略，见[获取视频详细信息（web端）](../video/info.md#获取视频详细信息（web 端）)中的stat对象
	Dynamic   string      `json:"dynamic"`   // 视频同步发布的的动态的文字内容。无为空
	Dimension Dimension   `json:"dimension"` // 稿件1P分辨率。略，见[获取视频详细信息（web端）](../video/info.md#获取 视频详细信息（web端）)中的dimension对象
	Count     int         `json:"count"`     // 稿件分P数。非投稿视频无此项
	Cid       int         `json:"cid"`       // 视频cid
	Progress  int         `json:"progress"`  // 观看进度时间。单位为秒
	AddAt     int         `json:"add_at"`    // 添加时间。时间戳
	Bvid      string      `json:"bvid"`      // 稿件bvid
}
