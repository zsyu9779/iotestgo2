// 04 Streaming RPC 服务端：三种流模式实现
//
// 启动：go run server/main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "iotestgo/module05_grpc/04_streaming_rpc/proto/chatpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedChatRoomServer
}

// ========== 1. Server-side streaming（服务端流） ==========
// 客户端发一次请求 → 服务端返回多条响应
func (s *server) Subscribe(req *pb.SubscribeRequest, stream pb.ChatRoom_SubscribeServer) error {
	roomID := req.GetRoomId()
	log.Printf("[Subscribe] 客户端加入房间: %s, 开始推送消息...", roomID)

	// 模拟推送 5 条消息，每条间隔 500ms
	for i := 1; i <= 5; i++ {
		msg := &pb.ChatMessage{
			User:      "system",
			Text:      fmt.Sprintf("[%s] 消息 #%d - 欢迎来到聊天室", roomID, i),
			Timestamp: time.Now().Unix(),
		}
		if err := stream.Send(msg); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}

	log.Printf("[Subscribe] 推送完成，流正常结束")
	return nil // return nil 表示流正常结束，客户端 Recv() 收到 io.EOF
}

// ========== 2. Client-side streaming（客户端流） ==========
// 客户端发多条请求 → 服务端汇总返回一条响应
func (s *server) SendMessages(stream pb.ChatRoom_SendMessagesServer) error {
	var messages []string
	count := 0

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// 客户端发送完毕，返回汇总
			log.Printf("[SendMessages] 收到 %d 条消息，返回汇总", count)
			return stream.SendAndClose(&pb.SendSummary{
				MessageCount: int32(count),
				Status:       "ok",
			})
		}
		if err != nil {
			return err
		}
		count++
		messages = append(messages, fmt.Sprintf("%s: %s", msg.GetUser(), msg.GetText()))
		log.Printf("[SendMessages] 收到第 %d 条: <%s> %s", count, msg.GetUser(), msg.GetText())
	}
}

// ========== 3. Bidirectional streaming（双向流） ==========
// 双方独立收发，全双工通信
func (s *server) Chat(stream pb.ChatRoom_ChatServer) error {
	var wg sync.WaitGroup

	// 接收 goroutine：读取客户端消息并回显
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				log.Printf("[Chat] 客户端关闭发送端")
				return
			}
			if err != nil {
				log.Printf("[Chat] 接收错误: %v", err)
				return
			}
			log.Printf("[Chat] 收到: <%s> %s", msg.GetUser(), msg.GetText())

			// 回显消息给客户端
			echo := &pb.ChatMessage{
				User:      "echo-bot",
				Text:      fmt.Sprintf("你说: %s", msg.GetText()),
				Timestamp: time.Now().Unix(),
			}
			if err := stream.Send(echo); err != nil {
				log.Printf("[Chat] 发送错误: %v", err)
				return
			}
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatRoomServer(s, &server{})
	reflection.Register(s)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down...")
		s.GracefulStop()
	}()

	log.Println("=== Streaming RPC Server 已启动 ===")
	log.Println("  监听端口: :50052")
	log.Println("  三种流模式: Subscribe(服务端流) / SendMessages(客户端流) / Chat(双向流)")
	log.Println()
	fmt.Println("测试: go run client/main.go")
	fmt.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
