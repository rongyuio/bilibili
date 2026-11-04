# AGENTS.md

封装 Bilibili HTTP API 的 Go 客户端库，模块路径 `github.com/rongyuio/bilibili`。上游 `CuteReimu/bilibili` 已被作者删除，本仓库接手维护；当前处于 v0 阶段，破坏性变更随次版本号发布，并记录到 `docs/migration.md`。要求 Go 1.27.3+：`go.mod` 的 `go` 指令表示最低支持版本，提升需说明理由，见 `docs/versioning.md`。新增或修改的文档与注释用中文；根包核心文件里既有的英文注释保持原样，不要为统一语言改写。

## 构建与验证

```bash
make check                        # 门禁全量：gencheck + fmtcheck + vet + race + build + lint
make help                         # 列出全部 target
go build ./...                    # 编译
go vet ./...                      # 静态检查
gofmt -s -w <file.go>             # 格式化；CI 要求 gofmt -s -l . 输出为空
golangci-lint run                 # 读取 .golangci.yml（v2 schema），CI 版本 v2.15.0
go test ./...                     # 全部测试
go test -run 'TestCookie|TestSignMap' ./...   # 按正则跑一组测试
go test -race -v ./...            # CI 使用的形式
go generate ./...                 # 重新生成 bili/client_gen.go
```

CI 位于 `.github/workflows/`，四个文件：

- `gofmt.yml`（GoFmt）—— 向 `master` 的 push / PR，要求 `gofmt -s -l .` 输出为空。
- `golangci-lint.yml`（Go）—— 向 `master` 的 push / PR，依次 `go generate ./...` → `git diff --exit-code` + `git status --porcelain`（生成物必须干净）→ golangci-lint → `go test -race -v ./...` → `go build -v ./...`。
- `label.yml`（Label）—— 仅 PR，按标题前缀打 label，供 Release Notes 分类。
- `release.yml`（Release）—— 仅 `v*.*.*` tag。

除 `label.yml` 外都通过 `go-version-file: go.mod` 取 Go 版本，因此改 `go.mod` 的 `go` 指令即可切换验证环境。

## 发版

纯库，没有部署环节。推 `v<major>.<minor>.<patch>` tag 触发 `release.yml`：校验 tag 名（semver，禁 build metadata）与「提交位于 `master`」→ `go test -race ./...` → `go build ./...` → `gh release create --generate-notes`（带预发布后缀的发为 Pre-release）。tag 格式与递增规则见 `docs/versioning.md`。

## 分支与推送

提交、分支、推送按全局约定。本仓库额外：`master` 有分支保护（禁强推与删除，必须走 PR），不要直接向 `master` 提交或推送，即使当前账号能绕过。误在本地 `master` 提交且尚未推送时，先用 `git branch <新分支>` 保住提交，确认工作区干净（`git status --porcelain` 为空，否则先 stash），再 `git reset --hard origin/master`；绝不强推。流程细节见 `.github/CONTRIBUTING.md`。

## 结构

- **根包 `bilibili`（仓库根目录）** —— 只放跨业务域的基础设施：构造与会话、请求链路、参数编码、错误类型与解码诊断、Cookie、WBI 签名，以及 `export.go` 中向子包导出的设施；不放业务方法和业务模型。`type.go` 只放真正跨业务域的基础类型（如 `Size`、`FormatCtrl`）。
- **业务子包** —— 按域拆分（`video/`、`user/`、`live/` 等，清单见 README「包结构」）。每个子包的 `client.go` 定义 `Client`（嵌入 `*bilibili.Client`）与 `New(*bilibili.Client) *Client`；不复制会话，也不自建 Resty 客户端。文件命名与命名规范见 `.github/CONTRIBUTING.md`。
- **聚合包 `bili`** —— `client.go` 手写 `New` / `NewWith` / `NewAnonymous`，`client_gen.go` 由 `internal/cmd/genall` 生成域字段。只做包装，不定义业务方法，也不持有独立会话；域级状态放根客户端或聚合层，不要塞进子包。
- **硬约束**：新增子包或改 `New` 构造签名后必须 `go generate ./...` 并提交 `bili/client_gen.go`，CI 以 diff 拦截过期生成物；复合包名需要切词时更新生成器的 `wordBreaks`。
- **返回体独立定义**：一个接口的 `*Result` 及其子类型只服务该接口，字段再像也不跨接口复用，也不要为了复用把返回体搬进根包。

## 请求链路

```
子包方法 -> Execute[Out]（export.go）-> newRequest（只取一次 Cookie 快照）
  -> withParams（request 标签 -> query/json/multipart）
  -> handlers（FillCsrf / FillParam / MoveFormParams / FillWbi）
  -> send -> sendRaw（出错也合并响应 Cookie）-> decodeResponse
```

子包可用 `export.go` 的 `Execute` / `ExecuteRaw`（非标准信封走后者）、`Handler`、`FillCsrf` / `FillFormCsrf`、`FillParam` / `MoveFormParams`、`(*Client).FillWbi`、`NewDecodeError`。需要读 302 `Location` 或 HTML 的接口直接走 `Resty()` 逃生通道（如 `(*miscellaneous.Client).UnwrapShortURL`）。

改动链路时必须守住的约束：

- **context 永远首参**。`checkContext` 在任何工作开始前拒绝 nil / 已结束的 context，WBI 密钥刷新也用它。
- **每次请求只取一次 Cookie 快照**（`newRequest` 深拷贝 `GetCookies`）；响应 Cookie 在 `sendRaw` 合并，HTTP 与业务错误时同样合并。`Client` 首次使用后不可复制；普通请求可并发，登录 / 会话切换 / `SetDroppedFieldHandler` / Resty 配置变更必须在请求之外串行。
- **解码到全新值**，失败不部分写入。容错只剪枝 JSON kind 与 Go 类型不兼容的叶子，`data` 根节点与实现了自定义 unmarshaler 的类型是不透明边界；诊断基于未剪枝的原始字节。WBI nav 是例外，走 `decodeWBIResponse`，忽略非零 code 且不容错。
- **参数位置只由 `request` 标签决定**（与 `json` 标签无关）：query / json / form-data / `-` / `field=name` / `default=v`。`omitempty` 优先于 `default`；同字段声明多个位置、或 JSON 与 multipart 混用，返回 `ParamError`。
- **WBI 只签 query**，每个键必须且只能有一个值；`FillWbi` 会显式清空 `Referer`，WBI 请求带 Referer 必失败。
- **CSRF 只从请求 Cookie 快照的 `bili_jct` 取**，缺失返回稳定的「登录过期」错误；`FillCsrf` / `FillFormCsrf` 同时写 `csrf` 与 `csrf_token`。
- **新增错误必须用包级静态变量**（`var errXxx = errors.New(...)`）。golangci 的 `err113` 与 staticcheck `ST1005` 只认标准库 `errors`；历史代码那些内联写法走的是 `github.com/pkg/errors`，不要照抄，中文错误串（「B站…」的 B 是拉丁大写）尤其容易触发 ST1005。
- 错误类型 `*HTTPError` / `Error` / `*ParamError` / `*DecodeError` 为 `errors.Is` / `errors.As` 设计，分类与诊断字段见 `docs/request.md`。`ParamError.Err`、响应值、Cookie 与凭证不要进日志。

客户端构造：`New()` 离线创建；`NewAnonymousClient(ctx)` 取游客 Cookie；`NewWithClient` 独占 Resty 客户端，把其 `Cookies` 复制进 Client 存储并关闭 Jar。直接调 `Resty()` 发请求会绕过签名、Cookie 存储与统一错误处理。

## 接口实现规则

- **不实现 APP 域名接口**（`app.bilibili.com`、`app.biliapi.com`、`app.biliapi.net`、`club.bilibili.com`）：它们要求 appkey 签名，库不计算。web 接口上的可选 `access_key` 不受此限。
- 没实测过的接口，不得在注释、文档或提交信息里写成「已验证」；要标注验证状态，只写有可追溯实测记录支撑的结论，不能凭「只读 GET」之类的属性反推。
- 新接口在对应子包加方法、`XxxParam` / `XxxResult`，并为它单独定义返回模型。
- `csrf` 通常不作为调用方参数；`RefreshCookieParam.Csrf` 这类字段允许显式覆盖，缺省才回退到 Cookie 快照。

## 测试

测试不访问真实 API、不需要凭证，全是纯逻辑用例（签名、参数编码、解码容错等）。`test/` 是独立 Go 模块（自带 `go.mod` + `replace`），被 gitignore，也被 `.golangci.yml` 的 `exclusions.paths` 排除，其依赖与 Go 版本不影响主模块。

## 行尾

`.gitattributes` 强制 `*.go` 与 `Makefile` 为 LF。不要让 `core.autocrlf` 检出 CRLF：gofmt CI 会全量报错，Makefile 的 `\` 续行也会散架。

## 被忽略的目录

`.gitignore` 忽略 `test/`、`.idea/` 与 `tools/sync_docs/`（早期的本地接口文档比对工具，已从仓库移除）。不要提交这些目录的内容。

## 文档

索引见 README。细节在 `docs/`：`usage.md`（常用接口示例）、`authentication.md`（登录与会话）、`request.md`（`Client.Do` 契约、参数标签、错误分类）、`migration.md`（迁移记录）、`versioning.md`（tag 与版本规则）。
