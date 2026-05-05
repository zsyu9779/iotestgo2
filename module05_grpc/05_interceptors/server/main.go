// 05 Interceptor 拦截器：日志、计时、panic 恢复
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
	"syscall"
	"time"

	pb "iotestgo/module05_grpc/05_interceptors/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()

	// 故意 panic 触发恢复拦截器
	if name == "panic" {
		panic("故意的 panic，测试恢复拦截器")
	}

	if name == "" {
		name = "World"
	}
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", name)}, nil
}

func (s *server) SayHelloStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloStreamServer) error {
	name := req.GetName()
	if name == "" {
		name = "World"
	}
	for i := 1; i <= 3; i++ {
		stream.Send(&pb.HelloResponse{
			Message: fmt.Sprintf("Hello, %s! (stream #%d)", name, i),
		})
		time.Sleep(300 * time.Millisecond)
	}
	return nil
}

// ========== Unary 拦截器 ==========

// loggingInterceptor 日志 + 计时拦截器
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	// 前置处理
	log.Printf("[REQ] %s: %v", info.FullMethod, req)

	// 调用实际处理方法
	resp, err := handler(ctx, req)

	// 后置处理
	elapsed := time.Since(start)
	if err != nil {
		log.Printf("[ERR] %s: %v (elapsed: %v)", info.FullMethod, err, elapsed)
	} else {
		log.Printf("[OK]  %s (elapsed: %v)", info.FullMethod, elapsed)
	}

	return resp, err
}

// recoveryInterceptor panic 恢复拦截器
func recoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC] %s: %v", info.FullMethod, r)
			err = status.Errorf(codes.Internal, "internal server error: %v", r)
		}
	}()
	return handler(ctx, req)
}

// ========== Stream 拦截器 ==========

// streamLoggingInterceptor 流式拦截器（包装 ServerStream 记录收发）
func streamLoggingInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("[STREAM-START] %s", info.FullMethod)
	err := handler(srv, &loggingServerStream{ServerStream: ss})
	log.Printf("[STREAM-END] %s (err=%v)", info.FullMethod, err)
	return err
}

type loggingServerStream struct {
	grpc.ServerStream
}

func (w *loggingServerStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)
	log.Printf("[STREAM RECV] %T, err=%v", m, err)
	return err
}

func (w *loggingServerStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)
	log.Printf("[STREAM SEND] %T", m)
	return err
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 链式组合多个拦截器：先恢复 panic，再记录日志
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(recoveryInterceptor, loggingInterceptor),
		grpc.StreamInterceptor(streamLoggingInterceptor),
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

	log.Println("=== Interceptor Server 已启动 ===")
	log.Println("  监听端口: :50053")
	log.Println("  拦截器链: recoveryInterceptor → loggingInterceptor")
	log.Println("  测试: go run client/main.go")
	log.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
