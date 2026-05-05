// 03 RPC 服务 - UserRpc 服务端（gRPC）
//
// 启动：go run userrpc/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "iotestgo/module06_gozero/03_rpc_service/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// 内存用户数据
var users = map[int64]*pb.User{
	1: {Id: 1, Username: "gopher", Email: "gopher@example.com", Status: 1},
	2: {Id: 2, Username: "alice", Email: "alice@example.com", Status: 1},
	3: {Id: 3, Username: "bob", Email: "bob@example.com", Status: 2},
}

type server struct {
	pb.UnimplementedUserRpcServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, ok := users[req.GetUserId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %d not found", req.GetUserId())
	}
	log.Printf("[UserRpc] GetUser: userId=%d → username=%s", req.GetUserId(), user.GetUsername())
	return &pb.GetUserResponse{User: user}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserRpcServer(s, &server{})
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down...")
		s.GracefulStop()
	}()

	fmt.Println("=== UserRpc 服务已启动 ===")
	fmt.Println("  监听端口: :9091 (gRPC)")
	fmt.Println("  测试: go run userapi/main.go")
	fmt.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
