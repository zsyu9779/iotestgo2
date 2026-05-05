// 06 Metadata & Auth 服务端：通过 metadata 提取和验证 Bearer Token
//
// 启动：go run server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pb "iotestgo/module05_grpc/06_metadata_auth/proto/authpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 需要认证才能调用
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 从 context 中获取认证后注入的 user_id
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		userID = "unknown"
	}

	name := req.GetName()
	if name == "" {
		name = "World"
	}

	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s! (user: %s)", name, userID),
	}, nil
}

// authInterceptor 认证拦截器：从 metadata 提取 Bearer Token 并验证
func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 1. 从 incoming context 中提取 metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// 2. 获取 authorization header
	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
	}

	// 3. 验证 Bearer Token 格式
	token := authHeaders[0]
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, status.Errorf(codes.Unauthenticated, "authorization must be Bearer <token>")
	}
	token = strings.TrimPrefix(token, "Bearer ")

	// 4. 校验 token 值
	if token != "valid-token-12345" {
		return nil, status.Errorf(codes.PermissionDenied, "invalid token: %s", token)
	}

	// 5. 认证通过，将用户信息注入 context 供业务方法使用
	ctx = context.WithValue(ctx, "user_id", "gopher-12345")
	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
	)
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down...")
		s.GracefulStop()
	}()

	log.Println("=== Auth Server 已启动 ===")
	log.Println("  监听端口: :50054")
	log.Println("  有效 Token: Bearer valid-token-12345")
	log.Println("  测试: go run client/main.go")
	log.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
