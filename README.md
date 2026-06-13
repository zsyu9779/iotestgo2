# iotestgo2

面向计算机专业本科生的 Go 后端课程演示项目。

## 课程定位

本项目用于课堂演示和课后练习，不是生产项目模板。Java 对比只用于连接学生已有课程经验，核心目标是建立 Go 后端开发的概念地图和动手能力。

## 模块

| 模块 | 内容 | 运行依赖 |
|---|---|---|
| module01_basics | Go 基础、指针、结构体、泛型、CLI 项目 | Go |
| module02_advanced | 接口、错误、并发、测试、文件 IO、反射 | Go |
| module03_web_gin | Gin、JWT、中间件、httptest、用户中心 | Go |
| module04_gorm | GORM、事务、Preload、Raw SQL、博客 API | Go + MySQL |
| module05_grpc | Protobuf、gRPC、stream、metadata、gateway | Go + protoc + grpcurl |
| module06_gozero | go-zero、API/RPC、Etcd、缓存、MQ、观测、部署 | Go + Docker |

## 快速开始

```bash
go version
make test-basic
make run-user-center
```

The root `Makefile` is the preferred classroom entry point. Use module READMEs for deeper per-lesson `go run` / `go test` commands.

When running DB/cache/RPC demos, copy `.env.example` to `.env` as shared classroom defaults. These defaults are not a guaranteed central config consumed by every lesson.

## 教学边界

部分示例刻意简化：明文密码、内存存储、硬编码 token、模拟消息队列和模拟 RPC 都只用于解释概念。正式工程需要替换为安全配置、持久化存储、真实 RPC client、真实消息系统和可观测性方案。
