// 服务端：实现 Greeter 服务，监听 50051 端口
//
// 运行前需先生成 proto 代码：
//   cd module05_grpc/03_unary_rpc/proto
//   protoc --go_out=. --go_opt=paths=source_relative \
//          --go-grpc_out=. --go-grpc_opt=paths=source_relative hello.proto
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 实际项目中 import 路径类似：
// pb "iotestgo2/module05_grpc/03_unary_rpc/proto/hellopb"
// 这里为教学演示使用伪代码结构

// server 实现 GreeterServer 接口
type server struct {
	// pb.UnimplementedGreeterServer  // 实际项目中嵌入此结构体以保证向前兼容
}

// SayHello 实现一元 RPC 方法
// ctx 携带截止时间、取消信号、metadata 等
func (s *server) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	// 检查 ctx 是否已超时或取消
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	name := req.GetName()
	if name == "" {
		name = "World"
	}

	// 模拟业务处理耗时
	time.Sleep(10 * time.Millisecond)

	return &HelloResponse{
		Message: fmt.Sprintf("Hello, %s! (from gRPC server)", name),
	}, nil
}

// ========== 以下是伪代码演示结构，实际需配合 proto 生成代码运行 ==========

func main() {
	fmt.Println("=== 03 Unary RPC 服务端 ===")
	fmt.Println()
	fmt.Println("步骤：")
	fmt.Println("1. 定义 .proto 文件 -> protoc 生成 .pb.go + _grpc.pb.go")
	fmt.Println("2. 实现 GreeterServer 接口（SayHello 方法）")
	fmt.Println("3. 创建 gRPC Server -> 注册服务 -> net.Listen + Serve")
	fmt.Println()

	// 实际运行代码：
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	//
	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	// reflection.Register(s) // 可选：允许 grpcurl 等工具发现服务
	//
	// log.Println("gRPC server listening on :50051")
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	// 确保编译通过（防止 import 未使用报错）
	_ = net.Listen
	_ = grpc.NewServer
	_ = reflection.Register
	_ = context.Canceled
	_ = log.Println
}

// 占位结构体（实际由 protoc 生成）
type HelloRequest struct {
	Name string
}

func (r *HelloRequest) GetName() string { return r.Name }

type HelloResponse struct {
	Message string
}
