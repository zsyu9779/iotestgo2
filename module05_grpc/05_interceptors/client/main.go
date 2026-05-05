// 05 Interceptor 客户端：演示拦截器服务端日志 + panic 恢复
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "iotestgo/module05_grpc/05_interceptors/proto/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient("localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	fmt.Println("=== Interceptor 客户端演示 ===")
	fmt.Println()

	// 1. 正常调用
	fmt.Println("--- 1. 正常调用 SayHello ---")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel1()
	resp, err := client.SayHello(ctx1, &pb.HelloRequest{Name: "Gopher"})
	if err != nil {
		log.Printf("错误: %v", err)
	} else {
		fmt.Printf("  响应: %s\n", resp.GetMessage())
	}
	fmt.Println()

	// 2. 触发 panic（演示恢复拦截器）
	fmt.Println("--- 2. 触发 panic（测试 recoveryInterceptor） ---")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()
	_, err = client.SayHello(ctx2, &pb.HelloRequest{Name: "panic"})
	if err != nil {
		st := status.Convert(err)
		fmt.Printf("  gRPC 错误码: %s\n", st.Code())
		fmt.Printf("  错误消息: %s\n", st.Message())
		fmt.Println("  → panic 被 recoveryInterceptor 成功捕获，转换为 gRPC Internal 错误")
	}
	fmt.Println()

	// 3. 服务端流式调用（演示 Stream Interceptor）
	fmt.Println("--- 3. 服务端流 SayHelloStream（测试 StreamInterceptor） ---")
	ctx3, cancel3 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel3()
	stream, err := client.SayHelloStream(ctx3, &pb.HelloRequest{Name: "Streamer"})
	if err != nil {
		log.Printf("错误: %v", err)
	} else {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Recv 错误: %v", err)
				break
			}
			fmt.Printf("  收到: %s\n", msg.GetMessage())
		}
	}
	fmt.Println()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: ServerInterceptor 接口 → interceptCall() 方法")
	fmt.Println("  Go:   grpc.UnaryInterceptor / grpc.StreamInterceptor 函数")
	fmt.Println("  Go 的拦截器是函数式，更轻量；用 grpc.ChainUnaryInterceptor 链式组合")
}
