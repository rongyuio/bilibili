package video

// 视频交互接口的响应模型。

// CoinVideoResult 是「投币视频」的 data 对象。
type CoinVideoResult struct {
	Like bool `json:"like"` // 是否点赞成功。true：成功。false：失败。已赞过则附加点赞失败
}

// FavourVideoResult 是「收藏视频」的 data 对象。
type FavourVideoResult struct {
	Prompt bool `json:"prompt"` // 是否为未关注用户收藏。false：否。true：是
}

// LikeCoinFavourVideoResult 是「一键三连视频」的 data 对象。
type LikeCoinFavourVideoResult struct {
	Like     bool `json:"like"`     // 是否点赞成功。true：成功。false：失败
	Coin     bool `json:"coin"`     // 是否投币成功。true：成功。false：失败
	Fav      bool `json:"fav"`      // 是否收藏成功。true：成功。false：失败
	Multiply int  `json:"multiply"` // 投币枚数。默认为2
}
