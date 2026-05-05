# Module 05: gRPC 开发

本模块专注于使用 gRPC 构建高性能微服务通信。

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
