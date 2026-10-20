package audio

// 音频状态数相关响应模型。

// GetAudioStatusNumberResult 是查询歌曲状态数接口的 data 对象。
type GetAudioStatusNumberResult struct {
	Sid     int `json:"sid"`     // 音频auid
	Play    int `json:"play"`    // 播放次数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}
