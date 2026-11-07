# AGENTS.md

本文件约定本仓库的任务执行、代码修改与交付要求。只修改当前目标所需的内容，保留其他已有工作。

## 项目定位

- 本项目是封装 Bilibili HTTP API 的 Go 客户端库，模块路径为 `github.com/rongyuio/bilibili`。原上游 `CuteReimu/bilibili` 已被作者删除，本仓库接手维护。
- 当前处于 v0 阶段。破坏性变更随次版本号发布，提交使用 `type!: 描述`，并在 `docs/migration.md` 中记录旧写法与新写法。
- 最低 Go 版本以 `go.mod` 的 `go` 指令为准，当前为 `1.27.3`。提升下限须说明必要性，并同步 README 和 `docs/versioning.md`；不要把它当作跟随最新工具链的配置。
- 新增或修改的文档与注释使用中文。根包核心文件中既有的英文注释保持原样，不为统一语言单独改写。
- 本项目是纯库，没有部署环节。

## 任务执行

1. 开始前检查当前分支、工作区和暂存区，识别已有改动。新任务使用 `feat/<描述>`、`fix/<描述>` 或 `chore/<描述>` 独立分支；继续同一任务时复用对应分支。
2. 先阅读相关实现、测试和文档。影响目标、范围、用户可见行为、数据处理或验收标准的歧义，先与用户确认；明确范围内的实现细节按项目惯例处理。
3. 只完成当前目标，不顺手重构、升级依赖或扩展功能。发现无关问题时单独说明。
4. 复杂且可并行的任务，先拆分范围、交付物和验收标准，再分派 subagent。显式指定思考等级为 `high`；工具不支持时如实说明。并行修改必须分配互不重叠的文件或模块，否则串行执行。
5. 分派后只继续不依赖结果的必要工作；没有此类工作时结束回合，等待通知，不重复执行已委派任务。优先使用完成通知；不支持时采用有间隔、有上限的状态检查。
6. 潜在长时任务设置合理的执行超时，区分等待返回与任务实际超时。超时后先检查状态，不重复启动或无限重试。
7. 交付前独立复核全部差异，包括 subagent 的修改，并完成与风险相符的验证。清理自己创建且不再需要的 shell、subagent、临时服务和浏览器标签；只关闭能确认归属的资源，不按进程名批量关闭。保留给用户验收的资源须说明用途及关闭方式。

## Git 与 GitHub

- 禁止直接向 `master` 提交或推送，即使当前账号能绕过分支保护；所有合入必须走 PR。
- 每次提交只包含一个明确目的，按文件或改动块暂存，提交前检查暂存区差异并运行必要验证。提交格式为 `type: 描述`，type 使用 `feat`、`fix`、`docs`、`refactor`、`test` 或 `chore`；破坏性变更加 `!`。
- 未经用户明确授权，不改写历史、不强制推送、不执行丢弃改动的操作。误在本地 `master` 提交且尚未推送时，先创建符合命名规则的分支保住提交；恢复 `master` 前检查工作区与暂存区，涉及 `reset --hard` 等丢弃操作必须已有明确授权，不自动 stash 或清理他人的改动。
- 推送、创建 PR、合入主分支、发版须在用户明确通过验收后，按其授权的动作依次执行。验收通过不等于授权后续全部动作；同一范围内已经明确的授权不重复询问。
- GitHub 操作使用 `gh` CLI。连接失败时可尝试本机已有 Clash 代理，仅对当前命令或会话生效，不修改持久或全局代理配置。
- 命名与协作流程详见 `.github/CONTRIBUTING.md`。

## 包结构与职责

| 位置 | 职责与约束 |
| --- | --- |
| 根包 `bilibili` | 只放跨业务域基础设施：构造与会话、请求链路、参数编码、错误与解码诊断、Cookie、WBI 签名，以及 `export.go` 中对子包导出的设施。不放业务方法或业务模型。 |
| 根包 `type.go` | 只放真正跨业务域的基础类型，如 `Size`、`FormatCtrl`。 |
| 业务子包，如 `video/`、`user/`、`live/` | 按业务域组织。每个 `client.go` 定义嵌入 `*bilibili.Client` 的 `Client` 和 `New(*bilibili.Client) *Client`；不复制会话、不自建 Resty 客户端。完整清单见 README。 |
| 聚合包 `bili` | `client.go` 手写 `New`、`NewWith`、`NewAnonymous`；`client_gen.go` 由 `internal/cmd/genall` 生成域字段。只做包装，不定义业务方法、不持有独立会话。 |

- 域级状态放在根客户端或聚合层，不放进业务子包。
- 新增业务子包或修改 `New` 构造签名后，必须运行 `go generate ./...`，并将生成物变更纳入同一提交；不要手改 `bili/client_gen.go`。复合包名需要切词时，同步更新生成器的 `wordBreaks`。
- 一个接口的 `*Result` 及其子类型只服务该接口；字段相同也不跨接口复用，不为复用将返回体搬进根包或公共文件。

## 接口实现

- 在对应业务子包添加方法和参数、返回类型，命名遵循 `.github/CONTRIBUTING.md`：`<方法名>Param`、`<方法名>Result`、`<方法名>Item`，子类型也以所属接口为前缀。无参数接口不创建空参数结构体，返回类型按接口实际形态定义。
- 不实现 APP 域名接口：`app.bilibili.com`、`app.biliapi.com`、`app.biliapi.net`、`club.bilibili.com`。这些接口要求 appkey 签名，本库不计算该签名。Web 接口上的可选 `access_key` 不受此限制。
- `csrf` 通常不暴露为调用方参数。`RefreshCookieParam.Csrf` 这类允许显式覆盖的字段，缺省时才回退到请求 Cookie 快照。
- 未经真实接口实测，不得在注释、文档或提交信息中写成「已验证」。验证状态必须有可追溯的实测记录，不能依据「只读 GET」或本地测试通过反推。

## 客户端与请求链路

```text
子包方法
  -> Execute[Out]（export.go）
  -> newRequest（一次 Cookie 快照）
  -> withParams（request 标签决定 query / json / multipart）
  -> handlers（FillCsrf / FillParam / MoveFormParams / FillWbi）
  -> send -> sendRaw（出错也合并响应 Cookie）
  -> decodeResponse
```

子包使用 `export.go` 提供的 `Execute`、`ExecuteRaw`、`Handler`、`FillCsrf`、`FillFormCsrf`、`FillParam`、`MoveFormParams`、`(*Client).FillWbi` 和 `NewDecodeError`。非标准响应信封使用 `ExecuteRaw`；需要读取 302 `Location` 或 HTML 时可直接使用 `Resty()`，例如 `(*miscellaneous.Client).UnwrapShortURL`。

### 构造与会话

- `New()` 离线创建客户端；`NewAnonymousClient(ctx)` 请求游客 Cookie。
- `NewWithClient` 独占传入的 Resty 客户端，将其 `Cookies` 复制到 Client 存储并关闭 Jar。不要重新启用 Jar 或直接设置 Resty 的 Cookies；会话 Cookie 使用 Client 的接口管理。
- `Client` 首次使用后不可复制。普通请求可并发；登录、会话切换、`SetDroppedFieldHandler` 和 Resty 配置变更必须在请求之外串行进行。
- 直接通过 `Resty()` 发请求会绕过签名、Client Cookie 存储和统一错误处理，使用前须明确承担这些处理职责。

### 请求与签名约束

- **Context**：需要 context 的 API 将其作为首参。`checkContext` 在请求准备前拒绝 nil 或已结束的 context，WBI 密钥刷新同样适用。
- **Cookie**：每次请求只取一次快照，`newRequest` 通过 `GetCookies` 取得深拷贝。`sendRaw` 合并响应 Cookie，HTTP 错误和业务错误也不能跳过合并。
- **参数**：位置只由 `request` 标签决定，与 `json` 标签无关。支持 `query`、`json`、`form-data`、`-`、`field=name`、`default=v` 和 `omitempty`；`omitempty` 优先于 `default`。同字段声明多个位置或 JSON 与 multipart 混用时返回 `ParamError`。
- **WBI**：只签 query，每个键必须且只能有一个值。`FillWbi` 显式清空 `Referer`，保留此行为，WBI 请求不得携带 Referer。
- **CSRF**：自动填充只从本次请求 Cookie 快照的 `bili_jct` 读取，缺失时返回稳定的「登录过期」错误。`FillCsrf` 和 `FillFormCsrf` 同时写入 `csrf` 与 `csrf_token`。

### 解码与错误约束

- 解码到全新值，失败时不部分写入目标。
- 容错只剪枝 JSON kind 与 Go 类型不兼容的叶子；`data` 根节点和实现自定义 unmarshaler 的类型是不透明边界。诊断必须基于未剪枝的原始字节。
- WBI nav 使用独立的 `decodeWBIResponse`：忽略非零 code，不执行容错剪枝。
- 新增错误必须使用包级静态变量，例如 `var errXxx = errors.New(...)`。遵守 `err113` 和 staticcheck `ST1005`；不要照抄历史代码借助 `github.com/pkg/errors` 内联创建错误的写法。错误串以拉丁大写字母开头（如「B站…」）也可能触发 ST1005。
- 保持 `*HTTPError`、`Error`、`*ParamError`、`*DecodeError` 的 `errors.Is` / `errors.As` 契约，分类与诊断字段见 `docs/request.md`。
- `ParamError.Err`、响应值、Cookie 和凭证不得写入日志。密钥、令牌、密码等敏感信息不得进入提交、截图或交付文档；排查时只展示必要且脱敏的信息。

## 验证与交付

按改动风险选择相关测试、构建或实际操作验证。纯文档改动检查内容与实现、配置是否一致，以及链接、差异和格式；代码改动运行相关测试与必要门禁。验证失败时修复本次引入的问题，无关的已有失败如实说明。未运行的检查不能写成通过，工具调用成功不能代替实际效果验证。

### 本地命令

```bash
make help                                  # 查看全部 target
make check                                 # gencheck + fmtcheck + vet + race + build + lint

go generate ./...                          # 生成 bili/client_gen.go
gofmt -s -w <file.go>                       # 只格式化本次涉及的 Go 文件
gofmt -s -l .                              # 格式检查，预期无输出
go vet ./...                               # 静态检查
go test ./...                              # 全部测试
go test -run 'TestCookie|TestSignMap' ./...  # 按名称选择测试
go test -race -v ./...                      # CI 的竞态检测形式
go build ./...                             # 编译
golangci-lint run                           # 使用 .golangci.yml 的 v2 配置
```

- 仓库测试不访问真实 API、不需要凭证，使用签名、参数编码、解码容错等本地用例。本地测试通过不代表真实接口已实测。
- `test/` 是被忽略的独立 Go 模块，有自己的 `go.mod` 和 `replace`，也被 lint 的 `exclusions.paths` 排除；其依赖与 Go 版本不影响主模块。
- `make gencheck` 在生成后检查 `bili/client_gen.go` 相对 Git 的差异。若该生成物是本次有意修改的一部分，提交前出现差异需人工核对，不能为使检查通过而丢弃改动。
- 交付时说明修改内容、实际验证结果、尚未完成或受阻的事项。

### CI 门禁

| 工作流 | 触发条件 | 检查或操作 |
| --- | --- | --- |
| `.github/workflows/gofmt.yml` | 向 `master` 的 push / PR | `gofmt -s -l .` 必须无输出。 |
| `.github/workflows/golangci-lint.yml` | 向 `master` 的 push / PR | 依次生成代码、检查 Git diff 与工作区干净、lint、竞态测试、构建。当前 lint 版本为 `v2.15.0`。 |
| `.github/workflows/label.yml` | PR 打开、编辑、重新打开或同步 | 按标题前缀打标签，供 Release Notes 分类。 |
| `.github/workflows/release.yml` | 推送匹配 `v*.*.*` 的 tag | 校验版本与提交归属，测试、构建并创建 GitHub Release。 |

除 `label.yml` 外，以上工作流通过 `go-version-file: go.mod` 选择 Go 版本。CI 在已提交的检出副本上执行生成物检查，要求生成后 `git diff --exit-code` 成功且 `git status --porcelain` 为空。

## 发布

发版前先满足前述验收与授权要求，再按 `docs/versioning.md` 选择版本。推送 `v<major>.<minor>.<patch>` tag 后，`release.yml` 会：

1. 校验语义化版本格式：允许预发布后缀，禁止 build metadata。
2. 确认 tag 指向的提交已包含在 `master` 中。
3. 运行 `go test -race ./...` 和 `go build ./...`。
4. 通过 `gh release create --generate-notes` 发布；带预发布后缀的版本标为 Pre-release，正式版本标为 Latest。重跑时工作流会删除并重建同名 GitHub Release。

## 文件与行尾

- `.gitattributes` 强制 `*.go` 和 `Makefile` 使用 LF。保持该设置，不让 `core.autocrlf` 引入 CRLF。
- 修改 Makefile 时遵循其现有约束：recipe 输出使用 ASCII，不使用多行 `\` 续行，以兼容 Windows 下的 make / MSYS sh；注释可用中文。
- 不提交 `.gitignore` 忽略的 `test/`、`.idea/` 和 `tools/sync_docs/`。后者是已从仓库移除的早期本地接口文档比对工具。

## 文档入口

| 文件 | 内容 |
| --- | --- |
| `README.md` | 项目入口、包结构与文档索引。 |
| `.github/CONTRIBUTING.md` | 方法、类型、文件命名与贡献流程。 |
| `docs/usage.md` | 常用接口示例。 |
| `docs/authentication.md` | 登录与会话。 |
| `docs/request.md` | `Client.Do` 契约、参数标签、错误分类与诊断。 |
| `docs/migration.md` | 破坏性变更的迁移记录。 |
| `docs/versioning.md` | 版本递增、tag 格式与最低 Go 版本策略。 |
