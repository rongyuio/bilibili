package article

// CoinArticle 接口的响应模型。
//
// 本文件的类型只服务于 CoinArticle，不与其它接口共用；即使字段相同，
// 其它接口也在各自的 model_*.go 中单独定义。

// CoinArticleResult 是「投币文章」的 data 对象。
type CoinArticleResult struct {
	Like bool `json:"like"` // 是否点赞成功。true：成功。false：失败。已赞过则附加点赞失败
}
