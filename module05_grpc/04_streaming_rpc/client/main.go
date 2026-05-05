// 04 Streaming RPC 客户端：演示三种流模式
//
// 用法：go run client/main.go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pb "iotestgo/module05_grpc/04_streaming_rpc/proto/chatpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatRoomClient(conn)
	ctx := context.Background()

	fmt.Println("=== Streaming RPC 客户端演示 ===")
	fmt.Println()

	// ========== 1. Server-side streaming ==========
	fmt.Println("--- 1. 服务端流 (Server-side streaming) ---")
	fmt.Println("   客户端发送一次请求 → 服务端推送多条消息")
	demoServerStream(ctx, client)
	fmt.Println()

	// ========== 2. Client-side streaming ==========
	fmt.Println("--- 2. 客户端流 (Client-side streaming) ---")
	fmt.Println("   客户端发送多条消息 → 服务端汇总返回一条响应")
	demoClientStream(ctx, client)
	fmt.Println()

	// ========== 3. Bidirectional streaming ==========
	fmt.Println("--- 3. 双向流 (Bidirectional streaming) ---")
	fmt.Println("   双方独立收发，全双工通信")
	demoBidiStream(ctx, client)

	fmt.Println()
	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: StreamObserver<Req> + StreamObserver<Resp> 回调模式")
	fmt.Println("  Go:   stream.Send() / stream.Recv() 同步调用模式")
}

func demoServerStream(ctx context.Context, client pb.ChatRoomClient) {
	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{RoomId: "go-room"})
	if err != nil {
		log.Printf("Subscribe 失败: %v", err)
		return
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("   [流正常结束: 收到 io.EOF]")
			break
		}
		if err != nil {
			log.Printf("Recv 错误: %v", err)
			return
		}
		fmt.Printf("   收到: [%s] %s\n", msg.GetUser(), msg.GetText())
	}
}

func demoClientStream(ctx context.Context, client pb.ChatRoomClient) {
	stream, err := client.SendMessages(ctx)
	if err != nil {
		log.Printf("SendMessages 失败: %v", err)
		return
	}

	messages := []string{"你好", "我在学 gRPC 流式传输", "很有意思"}
	for i, text := range messages {
		msg := &pb.ChatMessage{
			User:      "gopher",
			Text:      text,
			Timestamp: time.Now().Unix(),
		}
		if err := stream.Send(msg); err != nil {
			log.Printf("Send 错误: %v", err)
			return
		}
		fmt.Printf("   发送第 %d 条: %s\n", i+1, text)
	}

	// CloseAndRecv：关闭客户端发送端，等待服务端汇总响应
	summary, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("CloseAndRecv 错误: %v", err)
		return
	}
	fmt.Printf("   服务端汇总: 共收到 %d 条消息, 状态=%s\n", summary.GetMessageCount(), summary.GetStatus())
}

func demoBidiStream(ctx context.Context, client pb.ChatRoomClient) {
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Printf("Chat 失败: %v", err)
		return
	}

	var wg sync.WaitGroup

	// 接收 goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("   [双向流接收完成]")
				return
			}
			if err != nil {
				log.Printf("Recv 错误: %v", err)
				return
			}
			fmt.Printf("   收到回显: [%s] %s\n", msg.GetUser(), msg.GetText())
		}
	}()

	// 发送 3 条消息
	for _, text := range []string{"Hello", "Bidirectional", "Goodbye"} {
		stream.Send(&pb.ChatMessage{
			User:      "gopher",
			Text:      text,
			Timestamp: time.Now().Unix(),
		})
		fmt.Printf("   发送: %s\n", text)
		time.Sleep(200 * time.Millisecond)
	}

	// CloseSend：告知服务端不再发送，但可以继续接收
	stream.CloseSend()
	wg.Wait()
}
