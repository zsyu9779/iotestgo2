// 流式 RPC 客户端演示
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("=== 04 Streaming RPC 客户端 ===")
	fmt.Println()

	fmt.Println("1. Server-side streaming 接收：")
	fmt.Println("   stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{RoomId: \"go-room\"})")
	fmt.Println("   for {")
	fmt.Println("       msg, err := stream.Recv()")
	fmt.Println("       if err == io.EOF { break }  // 流正常结束")
	fmt.Println("       // 处理每条消息")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("2. Client-side streaming 发送：")
	fmt.Println("   stream, err := client.SendMessages(ctx)")
	fmt.Println("   for _, msg := range messages {")
	fmt.Println("       stream.Send(msg)")
	fmt.Println("   }")
	fmt.Println("   resp, err := stream.CloseAndRecv()  // 获取服务端汇总响应")
	fmt.Println()

	fmt.Println("3. Bidirectional streaming：")
	fmt.Println("   stream, err := client.Chat(ctx)")
	fmt.Println("   // 两个 goroutine，一个收一个发")
	fmt.Println("   go func() {")
	fmt.Println("       for { msg, _ := stream.Recv(); ... }  // 接收")
	fmt.Println("   }()")
	fmt.Println("   for { stream.Send(msg); ... }             // 发送")
	fmt.Println()

	fmt.Println("EOF 处理要点：")
	fmt.Println("  - 服务端 return nil = 正常结束，客户端 Recv() 收到 io.EOF")
	fmt.Println("  - 服务端 return err  = 异常结束，客户端 Recv() 收到对应 error")
	fmt.Println("  - 双向流中任意一方都可先关闭自己的发送端")

	_ = context.Background
	_ = io.EOF
	_ = log.Println
	_ = grpc.Dial
	_ = time.Now
}
