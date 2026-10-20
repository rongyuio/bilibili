# 命名规范和编码风格

> [!NOTE]
> 本仓库处于 v0 阶段，破坏性变更随次版本号发布即可，但仍应尽量避免；确需破坏时在 `type` 后加 `!` 标注，并在[迁移指南](../docs/migration.md)中补充旧写法与新写法的对应条目。

## 关于命名规范

- 函数命名参考中文翻译，例如“获取全部直播间分区列表”翻译成`GetLiveAreaList`，每个词一一对应，个别词语在不影响理解的情况下可以省略，例如这里“全部”被省略了。
  - 可以参考一下上游接口文档的用词，例如专栏文集叫`Articles`，专栏文章叫`Article`。
- 参数结构体命名为`<方法名>Param`，返回值命名为`<方法名>Result`，列表项命名为`<方法名>Item`。例如直播子包的`GetLiveAreaList`没有参数，签名为`live.GetLiveAreaList(ctx)`，返回`[]live.GetLiveAreaListItem`；专栏子包的`GetArticleInfo`参数是`article.GetArticleInfoParam`，返回`*article.GetArticleInfoResult`。
  - 如果完全没有传入参数，没必要弄个空的结构体，不传或传`nil`即可。
  - 返回值按接口的实际形态实例化，可以是`*GetXxxResult`、`[]GetXxxItem`或`map[string]*GetXxxItem`，不要省略`Get`前缀写成`LiveAreaList`这类名字。
  - 传入参数中的`csrf`特殊处理一下，不用调用者手动填写，可以参考`live/live.go`。
- 子类型的名字以所属接口为前缀，例如`GetArticleInfoStats`、`GetLiveAreaListSub`，这样同一子包内不同接口之间也不会重名。
- 每个接口的`*Result`及其子类型只服务于该接口，即使与其它接口字段相同也不跨接口复用；不要为了复用把某个接口的返回体搬到公共文件。根包`type.go`只放真正跨业务域的基础类型（如`Size`、`FormatCtrl`），`number.go`中的`NumberOrString`同属此类。
- 可能为`null`的字段用指针（例如`*FansDetail`），避免`json.Unmarshal`失败。

例子（摘自`article/model_article_info.go`；这些类型只服务于`GetArticleInfo`，其它接口的相同字段在各自的`model_*.go`中单独定义）：

```go
// GetArticleInfoResult 是「获取专栏文章基本信息」的 data 对象。
type GetArticleInfoResult struct {
    Like            int                          `json:"like"`              // 是否点赞。0：未点赞。1：已点赞。需要登录(Cookie) 。未登录为0
    Attention       bool                         `json:"attention"`         // 是否关注文章作者。false：未关注。true：已关注。需要登录(Cookie) 。未登录为false
    Favorite        bool                         `json:"favorite"`          // 是否收藏。false：未收藏。true：已收藏。需要登录(Cookie) 。未登录为false
    Coin            int                          `json:"coin"`              // 为文章投币数
    Stats           GetArticleInfoStats          `json:"stats"`             // 状态数信息
    Title           string                       `json:"title"`             // 文章标题
    BannerURL       string                       `json:"banner_url"`        // 文章头图url
    Mid             int                          `json:"mid"`               // 文章作者mid
    AuthorName      string                       `json:"author_name"`       // 文章作者昵称
    IsAuthor        bool                         `json:"is_author"`         // true。作用尚不明确
    ImageURLs       []string                     `json:"image_urls"`        // 动态封面
    OriginImageURLs []string                     `json:"origin_image_urls"` // 封面图片
    Shareable       bool                         `json:"shareable"`         // true。作用尚不明确
    ShowLaterWatch  bool                         `json:"show_later_watch"`  // true。作用尚不明确
    ShowSmallWindow bool                         `json:"show_small_window"` // true。作用尚不明确
    InList          bool                         `json:"in_list"`           // 是否收于文集。false：否。true：是
    Pre             int                          `json:"pre"`               // 上一篇文章cvid。无为0
    Next            int                          `json:"next"`              // 下一篇文章cvid。无为0
    ShareChannels   []GetArticleInfoShareChannel `json:"share_channels"`    // 分享方式列表
    Type            int                          `json:"type"`              // 文章类别。0：文章。2：笔记
    DisableShare    bool                         `json:"disable_share"`     // 是否禁止分享。实测为 false
}

// GetArticleInfoStats 是文章的状态数信息，只属于 GetArticleInfo。
type GetArticleInfoStats struct {
    View     int `json:"view"`     // 阅读数
    Favorite int `json:"favorite"` // 收藏数
    Like     int `json:"like"`     // 点赞数
    Dislike  int `json:"dislike"`  // 点踩数
    Reply    int `json:"reply"`    // 评论数
    Share    int `json:"share"`    // 分享数
    Coin     int `json:"coin"`     // 投币数
    Dynamic  int `json:"dynamic"`  // 动态转发数
}

// GetArticleInfoShareChannel 是文章的一种分享方式，只属于 GetArticleInfo。
type GetArticleInfoShareChannel struct {
    Name         string `json:"name"`          // 分享名称
    Picture      string `json:"picture"`       // 分享图片url
    ShareChannel string `json:"share_channel"` // 分享代号
}
```

## 关于go文件的命名

仓库采用「根包 + 业务子包」结构：根包（仓库根目录）只放客户端、请求链路、参数编码、错误与解码、Cookie、WBI 等基础设施；业务接口按域拆分到`user/`、`video/`、`live/`等 34 个子包，每个子包通过`New(*bilibili.Client)`复用根包客户端。

子包内部：

- `client.go`定义`Client`（嵌入`*bilibili.Client`）与`New(*bilibili.Client) *Client`；子包不复制会话，也不自行构造 Resty 客户端。
- 接口方法与对应的`*Param`按接口或主题放在子包内的`info.go`、`live.go`、`topic.go`这类文件中，例如用户信息接口在`user/info.go`，直播接口在`live/live.go`。
- 该接口的响应模型放在同目录的`model_<接口>.go`或`<主题>_model.go`，例如`article/model_article_info.go`、`video/model_info.go`。

对于并非对应某个接口的方法（例如一些纯工具类方法），可自行找一个合适的`.go`文件。

## 关于编码风格

本项目不做太多的编码风格限制，就提这样一个建议吧：

- 如果是纯粹的接口调用函数，请参考现有的函数写法。
- 如果是其它函数，不限制编码风格，提交 pull request 时 CI 会运行 gofmt（要求`gofmt -s -l .`输出为空）、golangci-lint、`go test -race -v ./...`与`go build -v ./...`，请确保全部通过。

## 分支与提交流程

`master` 已开启分支保护：禁止强推与删除，改动必须经 Pull Request 合入。**不要直接向 `master` 提交或推送**，即使当前账号有权限绕过保护。

1. 从最新的 `master` 切分支，命名为 `<type>/<简短描述>`，`type` 与提交消息的 type 保持一致（如 `ci/race-and-release-guard`、`docs/contributing-and-contacts`）：

   ```bash
   git switch master && git pull
   git switch -c ci/release-guard
   ```

2. 在分支上提交，消息按下一节的规范书写。
3. 推送并开 PR，等 CI 通过后合并：

   ```bash
   git push -u origin ci/release-guard
   gh pr create --fill
   ```

若不慎已在本地 `master` 提交、且尚未推送，按下面方式补救，**不要强推**：

```bash
git branch <type>/<描述>        # 在当前位置建分支，保住提交
git reset --hard origin/master  # 本地 master 回退到远端状态
git switch <type>/<描述>        # 切到分支继续
```

## 关于提交消息

格式为 [Conventional Commits](https://www.conventionalcommits.org/) 加中文描述：

```text
<type>(<scope>): <中文描述>
```

- `type` 必填，`scope` 可选，冒号后必须有一个空格，描述使用中文。
- 一个提交只做一件事，描述用祈使句说清"做了什么"，避免"修改"、"更新"这类无信息量的表述。
- 标题不写句号，长度尽量控制在 72 字符以内。

常用 `type`：

| type | 用途 |
| --- | --- |
| `feat` | 新增接口或功能 |
| `feature` | `feat` 的别名，label 同样归入 `feat` |
| `fix` | 修复缺陷 |
| `refactor` | 重构，不改变对外行为 |
| `docs` | 文档、注释、示例 |
| `test` | 测试 |
| `chore` | 构建、依赖、工具链等杂项 |
| `build` | 构建系统或依赖变更，label 归入 `chore` |
| `ci` | CI 配置 |
| `style` | 仅格式调整，label 归入 `chore` |
| `perf` | 性能优化，label 归入 `chore` |

PR 标题同样使用上述格式：`.github/workflows/label.yml` 会根据前缀自动打 label——`feat`/`feature` 打 `feat`，`fix`、`docs`、`refactor`、`test`、`ci` 各打同名 label，`chore`/`build`/`style`/`perf` 统一打 `chore`；标题含 `!` 时追加 `breaking`。Release Notes 依据这些 label 分类——所以 PR 标题的 `type` 不要省略，也不要使用上表之外的 `type`，否则不会打 label，只会落到兜底的「其它变更」。

`scope` 使用受影响的模块或文件名，例如：

```text
refactor(api): 细分模型文件并下沉剩余响应模型
fix(params): 修正参数名推导，初始缩写字段的键名与旧写法保持一致
chore(lint): 清零 golangci-lint 告警并让门禁转为阻断
ci: 推送 v*.*.* 标签时自动创建 GitHub Release
docs: 补充第五轮重构迁移说明
```

破坏性变更在 `type` 后加 `!`（如 `refactor(api)!: ...`），并在提交正文中说明影响与迁移方式，同时必须在[迁移指南](../docs/migration.md)中补充对应条目。

## 关于发版

发布由推送 `v*.*.*` tag 触发（`release.yml` 的触发条件就是 `v*.*.*`，`v1`、`v0.3` 这类段数不足的 tag 不会触发工作流）。工作流会先校验 tag 名符合语义化版本格式、且指向的提交位于 `master`，否则拒绝发布；随后运行 `go test -race ./...` 与 `go build ./...`，最后自动创建 GitHub Release。打 tag 的完整流程见 README 的「发版」一节，版本号格式与递增规则见[版本策略](../docs/versioning.md)。

- tag 必须是 `v<major>.<minor>.<patch>` 形式，可带 `-rc.1`、`-beta.2` 这类预发布后缀；不使用 build metadata（`v0.4.0+build.1` 会被拒绝）。
- 带预发布后缀的 tag 发布为 Pre-release，正式版本标记为 Latest。
- 不要给未合入 `master` 的提交打 tag，工作流会直接失败。
- 破坏性变更随次版本号一并发布，不需要升主版本或改模块路径（v0 阶段）。
