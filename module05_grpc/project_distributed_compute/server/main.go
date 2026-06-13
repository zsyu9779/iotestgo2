// 分布式计算项目 - 服务端
// 接收客户端流式数据，goroutine 池并发计算，流式返回结果
//
// 启动：go run server/main.go
package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	computeServer "iotestgo/module05_grpc/project_distributed_compute/internal/server"
	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDistributedComputeServer(s, computeServer.NewService(4))
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down...")
		s.GracefulStop()
	}()

	log.Println("=== 分布式计算服务 已启动 ===")
	log.Println("  监听端口: :50056")
	log.Println("  Worker 数量: 4")
	log.Println("  支持操作: sum, avg, max, min, stddev, median")
	log.Println("  测试: go run client/main.go")
	log.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
