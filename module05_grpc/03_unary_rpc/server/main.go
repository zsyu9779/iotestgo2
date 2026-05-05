// 03 Unary RPC 服务端：实现 GreeterServer 接口，启动 gRPC 服务
//
// 启动：go run server/main.go
// 测试：grpcurl -plaintext -d '{"name":"Gopher"}' localhost:50051 hello.Greeter/SayHello
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "iotestgo/module05_grpc/03_unary_rpc/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server 实现 GreeterServer 接口
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 Unary RPC 方法
// 客户端发一个请求 → 服务端返回一个响应
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 1. 检查 context 是否已取消（客户端超时/主动取消）
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// 2. 提取请求参数
	name := req.GetName()
	if name == "" {
		name = "World"
	}

	// 3. 模拟业务处理
	time.Sleep(10 * time.Millisecond)

	// 4. 构造响应
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s! (from gRPC server)", name),
	}, nil
}

func main() {
	// 1. 监听 TCP 端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. 创建 gRPC Server
	s := grpc.NewServer()

	// 3. 注册 Greeter 服务
	pb.RegisterGreeterServer(s, &server{})

	// 4. 注册反射服务（方便 grpcurl 等工具调试）
	reflection.Register(s)

	// 5. 优雅退出：捕获信号后停止接受新请求
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigCh
		log.Printf("Received signal %v, shutting down...", sig)
		s.GracefulStop()
	}()

	// 6. 启动服务
	log.Println("=== gRPC Server 已启动 ===")
	log.Println("  监听端口: :50051")
	log.Println("  测试命令: grpcurl -plaintext -d '{\"name\":\"Gopher\"}' localhost:50051 hello.Greeter/SayHello")
	log.Println("  或运行: go run client/main.go")
	log.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
