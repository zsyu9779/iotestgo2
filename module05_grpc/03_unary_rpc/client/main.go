// 03 Unary RPC 客户端：连接 gRPC 服务端，发起 SayHello 调用
//
// 用法：go run client/main.go [名字]
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "iotestgo/module05_grpc/03_unary_rpc/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 解析命令行参数
	name := "Gopher"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 2. 建立连接
	fmt.Println("正在连接 gRPC Server (localhost:50051)...")
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 3. 创建 Greeter 客户端
	client := pb.NewGreeterClient(conn)

	// 4. 设置超时 Context（超时控制贯穿整个调用链）
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 5. 发起 RPC 调用
	fmt.Printf("调用 SayHello(name=%q)...\n", name)
	start := time.Now()
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	elapsed := time.Since(start)

	if err != nil {
		log.Fatalf("RPC 调用失败: %v", err)
	}

	// 6. 打印响应
	fmt.Println()
	fmt.Println("=== 调用结果 ===")
	fmt.Printf("  响应: %s\n", resp.GetMessage())
	fmt.Printf("  耗时: %v\n", elapsed)
	fmt.Println()
	fmt.Println("关键点：")
	fmt.Println("  1. grpc.NewClient 建立 HTTP/2 连接（支持多路复用、连接池）")
	fmt.Println("  2. context.WithTimeout 设置 3 秒超时——超时控制贯穿调用链")
	fmt.Println("  3. 客户端 Stub 自动处理 Protobuf 序列化/反序列化")
	fmt.Println("  4. 底层使用 HTTP/2 + Protobuf 二进制传输")
}
