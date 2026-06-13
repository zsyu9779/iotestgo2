# Course Runbook

## 课前 10 分钟检查

```bash
go version
make fmt-check
make test-basic
```

Makefile targets are introduced by the next course-improvement task. Until then, use the per-module `go run` / `go test` commands from module READMEs.

Go 1.25.x is required for full verification. The recommended lab baseline is toolchain go1.25.0 or a newer patch in the 1.25 series. If a local machine still has Go 1.20.6, older Go may fail before downloading the requested toolchain if that Go version does not support the `toolchain` directive.

## 本机验证记录

```bash
$ go version
go version go1.25.11 darwin/arm64

$ go env GOVERSION GOMOD
go1.25.11
/Users/zhangshiyu/iotestgo2/go.mod
```

## 外部依赖

- Module 04 需要 MySQL。
- Module 05 需要 `protoc`、`protoc-gen-go`、`protoc-gen-go-grpc`，Gateway 需要 `protoc-gen-grpc-gateway`。
- Module 06 的电商综合项目需要 Docker Compose 启动 MySQL、Redis、Etcd、Prometheus、Grafana。

## 课堂运行顺序

1. 基础语法：`make run-basic-hello`
2. 并发：`make run-log-analyzer`
3. Gin 用户中心：`make run-user-center`
4. GORM 博客 API：`make run-blog-api`
5. gRPC 双向流：`make run-compute-server`，另开终端 `make run-compute-client`
6. go-zero 电商链路：`make ecommerce-up` 后按 README 启动服务。

## 现场故障处理

- `invalid go version`: 检查 Go 版本是否符合 `go.mod` 与 `toolchain`。本课程完整验证要求 Go 1.25.x；推荐课堂基线是 toolchain go1.25.0 或 1.25 系列更新补丁版本。如果本机仍是 Go 1.20.6，旧版本可能在下载 toolchain 前失败，需要先升级 Go。
- `connection refused`: 先确认对应 server 是否启动，端口是否被占用。
- `protoc: command not found`: 使用课程安装页安装 `protoc` 和 Go 插件。
- MySQL 连接失败：检查 Docker 容器、端口、DSN。
