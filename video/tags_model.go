package video

// 视频 TAG（新版）响应模型。

// GetVideoDetailTagsItem 是获取视频TAG信息（新）返回的单个 TAG。
//
// 字段有效性与 TagType 相关：普通标签使用 TagID，背景音乐使用 MusicID，
// 话题与背景音乐使用 JumpURL。
type GetVideoDetailTagsItem struct {
	TagID   int    `json:"tag_id"`   // tag_id。当 tag_type 不为 bgm 时有效
	TagName string `json:"tag_name"` // TAG名称
	MusicID string `json:"music_id"` // 背景音乐id。当 tag_type 为 bgm 时有效，以 MA 开头
	TagType string `json:"tag_type"` // TAG类型。old_channel：普通标签；topic：话题；bgm：背景音乐
	JumpURL string `json:"jump_url"` // 跳转url。当 tag_type 为 topic 或 bgm 时有效
}
