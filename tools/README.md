# 生成工具的使用方法

在仓库根目录执行：

```bash
go run tools/gen_struct.go
```

会得到这样的提示：

```console
请输入Markdown表格，在最后一行之后输入ok表示结束：（退出请输入exit）
```

把以下Markdown表格复制粘贴输入进去并回车：

```md
| 字段         | 类型  | 内容           | 备注                                                         |
| ------------ | ----- | -------------- | ------------------------------------------------------------ |
| id           | num   | 专栏cvid       |                                                              |
| title        | str   | 文章标题       |                                                              |
| state        | num   | 0              | 作用尚不明确                                                 |
| publish_time | num   | 发布时间       | 秒时间戳                                                     |
| words        | num   | 文章字数       |                                                              |
| image_urls   | array | 文章封面       |                                                              |
| category     | obj   | 文章标签       |                                                              |
| categories   | array | 文章标签列表   |                                                              |
| summary      | str   | 文章摘要       |                                                              |
| stats        | obj   | 文章状态数信息 |                                                              |
| like_state   | num   | 是否点赞       | 0：未点赞<br />1：已点赞<br />需要登录(Cookie) <br />未登录为0 |
```

然后输入：

```console
ok
```

并回车，你就会得到：

```go
type T struct {
    Id int `json:"id"` // 专栏cvid
    Title string `json:"title"` // 文章标题
    State int `json:"state"` // 0。作用尚不明确
    PublishTime int `json:"publish_time"` // 发布时间。秒时间戳
    Words int `json:"words"` // 文章字数
    ImageUrls []ImageUrl `json:"image_urls"` // 文章封面
    Category Category `json:"category"` // 文章标签
    Categories []Categorie `json:"categories"` // 文章标签列表
    Summary string `json:"summary"` // 文章摘要
    Stats Stats `json:"stats"` // 文章状态数信息
    LikeState int `json:"like_state"` // 是否点赞。0：未点赞。1：已点赞。需要登录(Cookie) 。未登录为0
}
```

自行复制到go代码中去，把`T`改个名即可。

---

# 和接口文档对字段（sync_docs）

`gen_struct.go` 是「文档 → 代码」的手动单向粘贴，粘贴完两边就脱钩了：文档加了字段你不会知道，
代码改了名文档也不会知道。`tools/sync_docs` 补这一环 —— 它把本库的结构体和
接口文档的字段表对起来。

先在旁边检出接口文档仓库，然后：

```bash
go run ./tools/sync_docs -check       # 只报漂移，有漂移退出码非 0
go run ./tools/sync_docs -write       # 把「文档有、Go 没有」的字段补进结构体
go run ./tools/sync_docs -v           # 打印每个锚点实际匹到哪张表
go run ./tools/sync_docs -coverage    # 还有多少结构体没挂锚点
go run ./tools/sync_docs -docs <路径>  # 接口文档检出位置，默认 ../bilibili-api-docs
```

退出码：`0` 一致，`1` 有漂移，`2` 用法/IO 出错。

## 为什么是「锚点表 + 工具」而不是全自动

文档里的对象叫 `data`、`list[]`、`data`中的`control`对象，而 Go 结构体叫
`CommentsControl`、`VipUserVip` —— **名字对不上**，没法推。靠字段重合度去猜会误配
（实测把 `vip_type` 匹到 `VipPrivilege`、把根对象的 `code` 匹到 `SearchRespData`），
而**往错的结构体加字段编译器不报错**，会静默上线。

所以锚点写死在 `tools/sync_docs/docmap.go` 里：一个结构体挂到文档的哪一节、哪张表。
那是人的判断，机器做不了，但做完一次就能一直用。

## 这个工具能做什么、不能做什么

**能**：报「文档有、Go 没有」（新增字段漏了）、报「Go 有、文档没有」（拼错或改名）、
按文档补齐字段、报锚点失效。

**不能**：判断 `num` 该是 int 还是指针还是 bool、判断两个结构体该不该复用、
决定新结构体叫什么名字。`-write` 对拿不准的类型会写 `any`，跑完要**看一眼**再提交。
它也不碰文法和注释质量。

## docmap.go 怎么维护

新增接口时给它的每个响应结构体挂一条锚点。没挂的结构体 `-check` 盯不到，
`-coverage` 会把它们列出来（现在 526 个结构体里挂了 38 个）。

一个结构体可以挂**多处**，比对时取并集 —— `Label` 被视频、用户两边的文档各写了一张表，
`VideoInfo` 的字段一部分在「获取视频基本信息」的 `data` 表、一部分在「获取视频超详细信息」的
`Related[]` 表（文档原话是「基本同 data 对象，另有以下字段」）。

⚠️ 锚点的 `Label` 是**包含匹配**，且**只看表前面两行**（文档的表头常写成
「标签行 + 说明行」两行）。写锚点时对着文档原文复制，别按印象写 ——
`vip`中的`label`对象 在用户文档里成立，在视频文档里原文是
`staff`数组中的对象中的`vip`对象中的`label`对象，少两个「对象」就匹不中。
匹不中会报 `✗ 锚点读不到表`，不会静默跳过。

## docKnown：已知但不在表里的字段

有些字段确实存在，但当前锚点的表里没有：文档表漏列（`Article` 漏了 16 个）、
字段在别的章节有表、或者只在本库的实机响应里见过。这些记在 `docmap.go` 的
`docKnown` 里，每条带一句为什么，`-check` 会从「Go 有、文档没有」里扣掉。

记在这里而不是留着红灯，是因为红灯久了就没人看 —— 新漂移会淹在已知的里面。
**删字段前先翻这里**：那些是确认存在的，不是残留。

## CI

`-check` 需要接口文档的检出，跨仓库拉取不值当，
所以没接进 CI。同步接口文档之后手动跑一次 `-check`，或者本地起个
workflow_dispatch 自己传 `-docs` 路径。
