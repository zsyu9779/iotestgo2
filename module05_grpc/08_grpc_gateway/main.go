// gRPC-Gateway：同时提供 gRPC 和 HTTP/JSON 接口
//
// 原理：通过 .proto 中的 google.api.http 注解，自动生成反向代理
// HTTP 请求 → Gateway → gRPC 请求 → gRPC Server
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("=== 08 gRPC-Gateway ===")
	fmt.Println()

	fmt.Println("架构图示：")
	fmt.Println("  ┌──────────┐     ┌───────────┐     ┌────────────┐")
	fmt.Println("  │  Browser  │────▶│  Gateway  │────▶│ gRPC Server │")
	fmt.Println("  │  (HTTP)   │     │  (:8080)  │     │  (:50051)   │")
	fmt.Println("  └──────────┘     └───────────┘     └────────────┘")
	fmt.Println("                   HTTP → gRPC 转换")
	fmt.Println()

	fmt.Println("--- 实现步骤 ---")
	fmt.Println()
	fmt.Println("1. proto 文件中添加 google.api.http 注解（见 hello.proto）")
	fmt.Println()
	fmt.Println("2. 安装 protoc-gen-grpc-gateway：")
	fmt.Println("   go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest")
	fmt.Println()
	fmt.Println("3. 生成 Gateway 代码：")
	fmt.Println("   protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative hello.proto")
	fmt.Println()
	fmt.Println("4. 服务端同时启动 gRPC 和 HTTP Gateway：")
	fmt.Println(`   func main() {
       // gRPC Server
       lis, _ := net.Listen("tcp", ":50051")
       grpcSrv := grpc.NewServer()
       pb.RegisterGreeterServer(grpcSrv, &server{})
       go grpcSrv.Serve(lis)

       // HTTP Gateway
       ctx := context.Background()
       mux := runtime.NewServeMux()
       opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
       pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)

       log.Println("HTTP Gateway listening on :8080")
       http.ListenAndServe(":8080", mux)
   }`)
	fmt.Println()

	fmt.Println("--- 效果 ---")
	fmt.Println("   curl -X POST http://localhost:8080/v1/hello \\")
	fmt.Println("        -H \"Content-Type: application/json\" \\")
	fmt.Println("        -d '{\"name\":\"Gopher\"}'")
	fmt.Println("   → {\"message\":\"Hello, Gopher!\"}")
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: Spring gRPC Starter / grpc-spring-boot-starter")
	fmt.Println("  Go:   grpc-gateway 独立代理进程或同进程多端口")
	fmt.Println("  Go 方案更轻量，不依赖 Spring 生态")

	_ = net.Listen
	_ = grpc.NewServer
	_ = http.ListenAndServe
	_ = context.Background
	_ = log.Println
}
