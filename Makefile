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

.PHONY: help build test race vet fmt lint fmtcheck check

help:
	@echo "make build       - build all packages"
	@echo "make test        - run tests"
	@echo "make race        - go test -race -count=1"
	@echo "make vet         - go vet"
	@echo "make fmt         - gofmt -s -w ."
	@echo "make fmtcheck    - fail if anything is unformatted"
	@echo "make lint        - golangci-lint run"
	@echo "make check       - the whole CI gate: fmtcheck + vet + race + build + lint"

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
