// 流式 RPC 服务端演示：三种流模式
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("=== 04 Streaming RPC 服务端 ===")
	fmt.Println()
	fmt.Println("三种流模式：")
	fmt.Println()
	fmt.Println("1. Server-side streaming（服务端流）：")
	fmt.Println("   客户端发一次请求 → 服务端返回多条响应（如实时推送、日志订阅）")
	fmt.Println("   实现：方法参数 (req, stream) 返回 error")
	fmt.Println("   代码模式：for { stream.Send(&msg) } → return nil")
	fmt.Println()
	showServerStreamExample()
	fmt.Println()

	fmt.Println("2. Client-side streaming（客户端流）：")
	fmt.Println("   客户端发多条请求 → 服务端汇总后返回一条响应（如上传文件、批量处理）")
	fmt.Println("   实现：方法参数 stream 返回 (response, error)")
	fmt.Println("   代码模式：for { req, err := stream.Recv() ... } → stream.SendAndClose(&resp)")
	fmt.Println()
	showClientStreamExample()
	fmt.Println()

	fmt.Println("3. Bidirectional streaming（双向流）：")
	fmt.Println("   双方独立收发，不互相阻塞（如聊天、实时协作）")
	fmt.Println("   实现：方法参数 stream 返回 error")
	fmt.Println("   代码模式：两个 goroutine 分别读/写，或用 select 交错处理")
	fmt.Println()
	showBidiStreamExample()

	fmt.Println("=== Java 对比 ===")
	fmt.Println("  Java: StreamObserver<Req> + StreamObserver<Resp> 回调模式")
	fmt.Println("  Go:   stream.Send() / stream.Recv() 同步但非阻塞")

	_ = net.Listen
	_ = grpc.NewServer
	_ = io.EOF
}

func showServerStreamExample() {
	fmt.Println("  // Server-side streaming 伪代码:")
	fmt.Println("  func (s *server) Subscribe(req *SubscribeRequest, stream ChatRoom_SubscribeServer) error {")
	fmt.Println("      for i := 0; i < 10; i++ {")
	fmt.Println("          stream.Send(&ChatMessage{User: \"system\", Text: fmt.Sprintf(\"msg %d\", i)})")
	fmt.Println("          time.Sleep(time.Second)")
	fmt.Println("      }")
	fmt.Println("      return nil  // return nil 表示流正常结束")
	fmt.Println("  }")

	_ = time.Now
	_ = log.Println
}

func showClientStreamExample() {
	fmt.Println("  // Client-side streaming 伪代码:")
	fmt.Println("  func (s *server) SendMessages(stream ChatRoom_SendMessagesServer) error {")
	fmt.Println("      count := 0")
	fmt.Println("      for {")
	fmt.Println("          msg, err := stream.Recv()")
	fmt.Println("          if err == io.EOF {")
	fmt.Println("              // 客户端发送完毕，返回汇总")
	fmt.Println("              return stream.SendAndClose(&SendSummary{MessageCount: int32(count), Status: \"ok\"})")
	fmt.Println("          }")
	fmt.Println("          if err != nil { return err }")
	fmt.Println("          count++")
	fmt.Println("          log.Printf(\"received: %s\", msg.GetText())")
	fmt.Println("      }")
	fmt.Println("  }")

	_ = io.EOF
}

func showBidiStreamExample() {
	fmt.Println("  // Bidirectional streaming 伪代码:")
	fmt.Println("  func (s *server) Chat(stream ChatRoom_ChatServer) error {")
	fmt.Println("      var wg sync.WaitGroup")
	fmt.Println("      // 接收 goroutine")
	fmt.Println("      wg.Add(1)")
	fmt.Println("      go func() {")
	fmt.Println("          defer wg.Done()")
	fmt.Println("          for {")
	fmt.Println("              msg, err := stream.Recv()")
	fmt.Println("              if err == io.EOF { return }")
	fmt.Println("              if err != nil { return }")
	fmt.Println("              // 回显给当前用户")
	fmt.Println("              stream.Send(&ChatMessage{User: \"echo\", Text: msg.GetText()})")
	fmt.Println("          }")
	fmt.Println("      }()")
	fmt.Println("      wg.Wait()")
	fmt.Println("      return nil")
	fmt.Println("  }")

	_ = sync.WaitGroup{}
}
