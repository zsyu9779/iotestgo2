// 客户端：连接 gRPC 服务端，发起 Unary RPC 调用
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("=== 03 Unary RPC 客户端 ===")
	fmt.Println()

	// 实际运行代码：
	// 1. 建立连接
	// conn, err := grpc.Dial("localhost:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	//
	// client := pb.NewGreeterClient(conn)
	//
	// 2. 带超时的 Context
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	//
	// 3. 发起 RPC 调用
	// resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Gopher"})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Response: %s", resp.GetMessage())

	fmt.Println("关键点：")
	fmt.Println("1. grpc.Dial 建立连接（支持负载均衡、连接池）")
	fmt.Println("2. context.WithTimeout 设置超时——超时控制贯穿整个调用链")
	fmt.Println("3. 客户端 Stub 自动生成的代码处理序列化/反序列化")
	fmt.Println("4. gRPC 默认使用 HTTP/2 + Protobuf 二进制传输")
	fmt.Println()
	fmt.Println("Java 对比：")
	fmt.Println("  Java: ManagedChannel + stub.blockingStub()")
	fmt.Println("  Go:   grpc.Dial + pb.NewXxxClient()")
	fmt.Println("  Go 的 context 传递超时/取消，Java 用 deadline 或拦截器")

	_ = grpc.Dial
	_ = insecure.NewCredentials
	_ = context.DeadlineExceeded
	_ = log.Println
	_ = time.Now
}
