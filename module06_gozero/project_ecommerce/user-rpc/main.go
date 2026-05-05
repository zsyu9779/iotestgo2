// 电商项目 - User RPC 服务
// 职责：管理用户数据，提供 GetUser、CheckUserStatus 接口
//
// 启动：go run user-rpc/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// 简化版 proto message 定义（实际项目用 protoc + goctl 生成）

type GetUserRequest struct {
	UserId int64
}
type GetUserResponse struct {
	UserId   int64
	Username string
	Email    string
	Status   int32
}

type CheckUserRequest struct {
	UserId int64
}
type CheckUserResponse struct {
	Valid  bool
	Status int32
	Reason string
}

// UserServiceServer 接口
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CheckUserStatus(context.Context, *CheckUserRequest) (*CheckUserResponse, error)
}

type userServer struct {
	mu    sync.RWMutex
	users map[int64]*GetUserResponse
}

func newUserServer() *userServer {
	return &userServer{
		users: map[int64]*GetUserResponse{
			1: {UserId: 1, Username: "gopher", Email: "gopher@example.com", Status: 1},
			2: {UserId: 2, Username: "alice", Email: "alice@example.com", Status: 1},
			3: {UserId: 3, Username: "bob", Email: "bob@example.com", Status: 2}, // 禁用
		},
	}
}

func (s *userServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %d not found", req.UserId)
	}
	log.Printf("[UserRpc] GetUser: userId=%d → username=%s", req.UserId, user.Username)
	return user, nil
}

func (s *userServer) CheckUserStatus(ctx context.Context, req *CheckUserRequest) (*CheckUserResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[req.UserId]
	if !ok {
		return &CheckUserResponse{Valid: false, Reason: "user not found"}, nil
	}
	if user.Status == 2 {
		return &CheckUserResponse{Valid: false, Status: user.Status, Reason: "user disabled"}, nil
	}
	return &CheckUserResponse{Valid: true, Status: user.Status}, nil
}

// ========== gRPC Service Descriptor (简化，不用 proto 生成) ==========

var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				req := &GetUserRequest{}
				if err := dec(req); err != nil {
					return nil, err
				}
				return srv.(UserServiceServer).GetUser(ctx, req)
			},
		},
		{
			MethodName: "CheckUserStatus",
			Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				req := &CheckUserRequest{}
				if err := dec(req); err != nil {
					return nil, err
				}
				return srv.(UserServiceServer).CheckUserStatus(ctx, req)
			},
		},
	},
}

func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	s.RegisterService(&UserService_ServiceDesc, newUserServer())
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down UserRpc...")
		s.GracefulStop()
	}()

	fmt.Println("=== User RPC 服务 已启动 ===")
	fmt.Println("  gRPC 端口: :9091")
	fmt.Println("  接口: GetUser, CheckUserStatus")
	fmt.Println()
	fmt.Println("  用户数据：ID=1 gopher(正常) ID=2 alice(正常) ID=3 bob(禁用)")
	fmt.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
