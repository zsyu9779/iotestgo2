# Module 05: gRPC 开发

本模块专注于使用 gRPC 构建高性能微服务通信。

> 教学说明：本模块代码优先服务课堂理解，部分实现会使用内存数据、固定 token、简化错误处理或本地模拟组件。讲课时应明确这些是教学简化，不是生产写法。

## 目录结构

### 01_protobuf_basics/
- Protobuf3 语法速览、message/enum/service 定义、与 JSON 对比

### 02_codegen/
- protoc 编译、protoc-gen-go/protoc-gen-go-grpc 插件、生成 .pb.go 和 _grpc.pb.go 解读

### 03_unary_rpc/
- 一元 RPC：服务端实现、客户端 Dial、Context 超时传递

### 04_streaming_rpc/
- 流式 RPC：Server-side / Client-side / Bidirectional streaming、EOF 处理

### 05_interceptors/
- UnaryInterceptor / StreamInterceptor：日志、panic 恢复

### 06_metadata_auth/
- Metadata 读写、Bearer Token 认证

### 07_error_handling/
- status 包、codes、自定义 Error Detail

### 08_grpc_gateway/
- gRPC + HTTP 双暴露、HTTP ↔ gRPC 转换

### project_distributed_compute/
- 分布式计算实战：客户端流式发数据 → 服务端实时计算 → 流式返回结果

## 建议讲课顺序

1. `01_protobuf_basics`: 只讲 IDL 和 message。
2. `02_codegen`: 讲 `protoc` 到 `*.pb.go` / `*_grpc.pb.go`。
3. `03_unary_rpc`: 讲最小 request/response。
4. `04_streaming_rpc`: 讲三种 stream。
5. `05_interceptors`: 讲横切能力。
6. `06_metadata_auth`: 讲 metadata 和认证。
7. `07_error_handling`: 讲 status code。
8. `08_grpc_gateway`: 讲 annotation-based HTTP bridge。
9. `project_distributed_compute`: 讲综合项目，重点是双向流 + worker pool + auth interceptor + testable engine。

## 生成代码规则

不要手改 `*.pb.go`、`*_grpc.pb.go`、`*.gw.pb.go`。修改 `.proto` 后运行对应 `gen.sh`。

## 学习目标

1. 掌握 Protobuf3 语法与代码生成流程
2. 熟练编写 Unary 和 Streaming RPC
3. 理解 Interceptor 机制并实现日志/认证/恢复
4. 掌握 Metadata 传递与 Token 认证
5. 正确处理 gRPC 错误与状态码
6. 能够通过 gRPC-Gateway 同时提供 HTTP 和 gRPC 接口
7. 具备构建分布式计算系统的能力

## 前置条件

- 安装 protoc：`brew install protobuf` 或从 GitHub 下载
- 安装 Go 插件：
  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
