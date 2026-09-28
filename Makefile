# bilibili Go 客户端库 —— 常用开发命令
#
# 只收「跨好几个子命令」或「参数不好记」的入口；单条 go 命令直接敲就好。
#
# ⚠️ recipe 里的 echo 一律用 ASCII。不是风格问题：Windows 版 make 把命令串
# 交给 MSYS sh 时，中文字节过不了 `if ... then ... fi` 这类结构（实测报
# `unexpected EOF while looking for matching '"'`），而同样的中文放在纯
# `echo` 里又是好的。注释可以随便中文，运行时输出不行。
# 另：recipe 也不要用多行 `\` 续行 —— 同样会被拆散，写成单行。

GO     ?= go
LINTER ?= golangci-lint

# `make sync-docs` 要用的接口文档检出位置。
#
# **不写死目录名** —— 它跟本地布局绑定、跟代码本身无关，写进仓库会一路带进
# 公开历史。所以这里按**特征文件**在同级目录里探测：谁有 docs/video/info.md，
# 谁就是那份检出。
#
# 探测不中（检出不在同级、或同时命中多个）时显式指定：
#
#     make sync-docs DOCS=../某个目录
#     BAC_DOCS=/绝对路径 make sync-docs
DOCS ?= $(or $(BAC_DOCS),$(patsubst %/docs/video/info.md,%,$(firstword $(wildcard ../*/docs/video/info.md))))

.PHONY: help build test race vet fmt lint fmtcheck check sync-docs

help:
	@echo "make build       - build all packages"
	@echo "make test        - run tests"
	@echo "make race        - go test -race -count=1 (what CI runs)"
	@echo "make vet         - go vet"
	@echo "make fmt         - gofmt -s -w ."
	@echo "make fmtcheck    - fail if anything is unformatted"
	@echo "make lint        - golangci-lint run"
	@echo "make check       - the whole CI gate: fmtcheck + vet + race + build + lint"
	@echo "make sync-docs   - diff structs against the API docs field tables"

build:
	$(GO) build ./...

test:
	$(GO) test ./...

race:
	$(GO) test -race -count=1 ./...

vet:
	$(GO) vet ./...

fmt:
	gofmt -s -w .

fmtcheck:
	@out=$$(gofmt -s -l .); if [ -n "$$out" ]; then echo "$$out"; echo "gofmt: unformatted files listed above (run: make fmt)"; exit 1; fi

lint:
	$(LINTER) run

check: fmtcheck vet race build lint

sync-docs:
	@test -n "$(DOCS)" || { echo "no API docs checkout found. Point at it explicitly:"; echo "    make sync-docs DOCS=../some-dir"; exit 1; }
	$(GO) run ./tools/sync_docs -check -docs "$(DOCS)"
