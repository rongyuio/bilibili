package audio

// 音频基本信息相关响应模型。

// GetAudioInfoStatistic 是音频基本信息中的状态数。
type GetAudioInfoStatistic struct {
	Sid     int `json:"sid"`     // 音频auid
	Play    int `json:"play"`    // 播放次数
	Collect int `json:"collect"` // 收藏数
	Comment int `json:"comment"` // 评论数
	Share   int `json:"share"`   // 分享数
}

// GetAudioInfoVipInfo 是音频基本信息中的 UP 主会员状态。
type GetAudioInfoVipInfo struct {
	Type       int   `json:"type"`         // 会员类型。0：无。1：月会员。2：年会员
	Status     int   `json:"status"`       // 会员状态。0：无。1：有
	DueDate    int64 `json:"due_date"`     // 会员到期时间。时间戳，单位毫秒
	VipPayType int   `json:"vip_pay_type"` // 会员开通状态。0：无。1：有
}

// GetAudioInfoResult 是查询音频基本信息接口的 data 对象。
type GetAudioInfoResult struct {
	ID         int                   `json:"id"`         // 音频auid
	UID        int                   `json:"uid"`        // UP主mid
	Uname      string                `json:"uname"`      // UP主昵称
	Author     string                `json:"author"`     // 作者名
	Title      string                `json:"title"`      // 歌曲标题
	Cover      string                `json:"cover"`      // 封面图片url
	Intro      string                `json:"intro"`      // 歌曲简介
	Lyric      string                `json:"lyric"`      // lrc歌词url
	Crtype     int                   `json:"crtype"`     // 1。作用尚不明确
	Duration   int                   `json:"duration"`   // 歌曲时间长度。单位为秒
	Passtime   int64                 `json:"passtime"`   // 歌曲发布时间。时间戳
	Curtime    int64                 `json:"curtime"`    // 当前请求时间。时间戳
	Aid        int                   `json:"aid"`        // 关联稿件avid。无为0
	Bvid       string                `json:"bvid"`       // 关联稿件bvid。无为空
	Cid        int64                 `json:"cid"`        // 关联视频cid。无为0
	Msid       int                   `json:"msid"`       // 0。作用尚不明确
	Attr       int                   `json:"attr"`       // 0。作用尚不明确
	Limit      int                   `json:"limit"`      // 0。作用尚不明确
	ActivityID int                   `json:"activityId"` // 0。作用尚不明确
	Limitdesc  string                `json:"limitdesc"`  // 空。作用尚不明确
	CTime      *int                  `json:"ctime"`      // 作用尚不明确。文档标注为 null
	Statistic  GetAudioInfoStatistic `json:"statistic"`  // 状态数
	VipInfo    GetAudioInfoVipInfo   `json:"vipInfo"`    // UP主会员状态
	CollectIDs []int                 `json:"collectIds"` // 歌曲所在的收藏夹mlid。需要登录（SESSDATA）
	CoinNum    int                   `json:"coin_num"`   // 投币数
}

// GetAudioTagsItem 是查询音频 TAG 列表接口返回的一项。
type GetAudioTagsItem struct {
	Type    string `json:"type"`    // song。作用尚不明确
	Subtype int    `json:"subtype"` // 作用尚不明确
	Key     int    `json:"key"`     // TAG id。作用尚不明确
	Info    string `json:"info"`    // TAG名
}

// GetAudioMembersGroup 是查询音频创作成员列表接口返回的一个成员分组。
type GetAudioMembersGroup struct {
	List []GetAudioMembersGroupMember `json:"list"` // 成员列表
	Type int                          `json:"type"` // 成员类型代码。1：歌手。2：作词。3：作曲。4：编曲。5：后期/混音。7：封面制作。8：音源。9：调音。10：演奏。11：乐器。127：UP主
}

// GetAudioMembersGroupMember 是成员分组中的一位成员。
type GetAudioMembersGroupMember struct {
	Mid      int    `json:"mid"`       // 0。作用尚不明确
	Name     string `json:"name"`      // 成员名
	MemberID int    `json:"member_id"` // 成员id。作用尚不明确
}
