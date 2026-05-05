// 07 gRPC 错误处理：根据请求参数返回不同的 gRPC status codes
//
// 启动：go run server/main.go
package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "iotestgo/module05_grpc/07_error_handling/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 根据请求 name 返回不同的结果或错误
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()

	switch name {
	case "":
		// InvalidArgument：参数错误
		return nil, status.Errorf(codes.InvalidArgument, "name is required (field violation: name must not be empty)")

	case "notfound":
		// NotFound：资源不存在
		return nil, status.Errorf(codes.NotFound, "user %q not found", name)

	case "exists":
		// AlreadyExists：资源已存在（冲突）
		return nil, status.Errorf(codes.AlreadyExists, "user %q already exists", name)

	case "timeout":
		// 模拟慢处理 → 客户端 context deadline 超时
		time.Sleep(2 * time.Second)
		// 检查 context 是否已取消
		if ctx.Err() != nil {
			return nil, status.FromContextError(ctx.Err()).Err()
		}
		return &pb.HelloResponse{Message: "Hello, " + name + "!"}, nil

	case "panic":
		panic("deliberate panic for testing")

	default:
		// OK
		return &pb.HelloResponse{Message: "Hello, " + name + "!"}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down...")
		s.GracefulStop()
	}()

	log.Println("=== Error Handling Server 已启动 ===")
	log.Println("  监听端口: :50055")
	log.Println("  测试:")
	log.Println("    go run client/main.go              # 全部场景")
	log.Println("    go run client/main.go Gopher       # 正常")
	log.Println("    go run client/main.go ''           # InvalidArgument")
	log.Println("    go run client/main.go notfound     # NotFound")
	log.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
