package audio

// 音频榜单相关响应模型。
//
// 各接口的 data 层与嵌套对象均独立定义，即使字段相同也不复用。

// GetAudioRankPeriodListResult 是获取音频榜单每期列表的响应 data 层。
type GetAudioRankPeriodListResult struct {
	List map[string][]GetAudioRankPeriodListItem `json:"list"` // 年份索引。键为年份，值为该年份的每期信息
}

// GetAudioRankPeriodListItem 是音频榜单的单期信息。
type GetAudioRankPeriodListItem struct {
	ID          int   `json:"ID"`           // 榜单id
	Priod       int   `json:"priod"`        // 榜单期数。文档字段名如此拼写
	PublishTime int64 `json:"publish_time"` // 发布时间。秒时间戳
}

// GetAudioRankDetailResult 是查询音频榜单单期信息的响应 data 层。
type GetAudioRankDetailResult struct {
	ListenFid          int    `json:"listen_fid"`            // 畅听版歌单收藏夹原始id。非真实收藏夹mlid，需在后方拼接用户mid的后两位
	AllFid             int    `json:"all_fid"`               // 完整版歌单收藏夹原始id。非真实收藏夹mlid，算法同上
	FavMid             int64  `json:"fav_mid"`               // 绑定收藏夹用户的mid
	CoverURL           string `json:"cover_url"`             // 榜单封面url
	IsSubscribe        bool   `json:"is_subscribe"`          // 是否已订阅榜单。需要登录，未登录恒为false
	ListenCount        int    `json:"listen_count"`          // 平台有版权音频的数量
	ListenMusicSwitch  bool   `json:"listen_music_switch"`   // 作用尚不明确。实测为false
	ShootVideoSwitch   bool   `json:"shoot_video_switch"`    // 作用尚不明确。实测为false
	FirstMusicCoverURL string `json:"first_music_cover_url"` // 第一名音乐封面url
}

// GetAudioRankMusicListResult 是获取音频榜单单期内容的响应 data 层。
type GetAudioRankMusicListResult struct {
	List   []GetAudioRankMusicListItem `json:"list"`    // 内容列表
	ListID int                         `json:"list_id"` // 榜单id
}

// GetAudioRankMusicListItem 是音频榜单中的单条内容。
type GetAudioRankMusicListItem struct {
	MusicID          string   `json:"music_id"`           // 音频MAID。例如 MA409252256362326366
	MusicTitle       string   `json:"music_title"`        // 音频标题
	Singer           string   `json:"singer"`             // 音频作者
	Album            string   `json:"album"`              // 音频专辑
	MVAid            int      `json:"mv_aid"`             // 音频MV的avid。无MV时为0
	MVBvid           string   `json:"mv_bvid"`            // 音频MV的bvid
	MVCover          string   `json:"mv_cover"`           // 音频封面url
	Heat             int      `json:"heat"`               // 热度值
	Rank             int      `json:"rank"`               // 排序值。1为最高排序
	CanListen        bool     `json:"can_listen"`         // 平台是否有版权
	Recommendation   string   `json:"recommendation"`     // 作用尚不明确
	CreationAid      int      `json:"creation_aid"`       // 关联稿件avid
	CreationBvid     string   `json:"creation_bvid"`      // 关联稿件bvid
	CreationCover    string   `json:"creation_cover"`     // 关联稿件封面url
	CreationTitle    string   `json:"creation_title"`     // 关联稿件标题
	CreationUp       int      `json:"creation_up"`        // 关联稿件UP主mid
	CreationNickname string   `json:"creation_nickname"`  // 关联稿件UP主昵称
	CreationDuration int      `json:"creation_duration"`  // 关联稿件时长。单位为秒
	CreationPlay     int      `json:"creation_play"`      // 关联稿件播放量
	CreationReason   string   `json:"creation_reason"`    // 关联稿件二级分区名
	Achievements     []string `json:"achievements"`       // 获得成就
	MaterialID       int      `json:"material_id"`        // 作用尚不明确
	MaterialUseNum   int      `json:"material_use_num"`   // 作用尚不明确
	MaterialDuration int      `json:"material_duration"`  // 作用尚不明确
	MaterialShow     int      `json:"material_show"`      // 作用尚不明确
	SongType         int      `json:"song_type"`          // 作用尚不明确
	CreationFirstCid int64    `json:"creation_first_cid"` // 作用尚不明确
	IsVt             int      `json:"is_vt"`              // 作用尚不明确
	LastRank         int      `json:"last_rank"`          // 上期名次
	ArcList          any      `json:"arc_list"`           // 作用尚不明确。实测为null
}
