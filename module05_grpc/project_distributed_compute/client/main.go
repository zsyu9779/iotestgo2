// 分布式计算客户端：流式发送计算任务，流式接收结果
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("=== 分布式计算项目 - 客户端 ===")
	fmt.Println()
	fmt.Println("客户端流程：")
	fmt.Println("1. 建立连接 → 创建 Process stream")
	fmt.Println("2. 启动接收 goroutine（接收结果并打印）")
	fmt.Println("3. 逐个发送计算任务")
	fmt.Println("4. CloseSend() 通知服务端不再发送")
	fmt.Println("5. 等待结果接收完毕")
	fmt.Println()

	fmt.Println("--- 核心伪代码 ---")
	fmt.Println()
	fmt.Println(`func runClient() {
    conn, _ := grpc.Dial("localhost:50051", ...)
    client := pb.NewDistributedComputeClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    stream, err := client.Process(ctx)
    if err != nil { log.Fatal(err) }

    // 接收 goroutine
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            result, err := stream.Recv()
            if err == io.EOF { return }
            if err != nil { log.Fatal(err) }
            fmt.Printf("Result: task=%s op=%s value=%.2f status=%s\n",
                result.TaskId, result.Operation, result.Value, result.Status)
        }
    }()

    // 发送任务
    tasks := []*pb.ComputeTask{
        {TaskId: "t1", Numbers: []int64{1,2,3,4,5}, Operation: "sum"},
        {TaskId: "t2", Numbers: []int64{1,2,3,4,5}, Operation: "avg"},
        {TaskId: "t3", Numbers: []int64{3,1,4,1,5,9,2,6}, Operation: "max"},
    }
    for _, task := range tasks {
        if err := stream.Send(task); err != nil {
            log.Printf("send error: %v", err)
        }
    }

    stream.CloseSend()  // 告知服务端：发送完毕
    wg.Wait()           // 等待接收完成
}`)
	fmt.Println()

	fmt.Println("--- 关键技术点 ---")
	fmt.Println("1. Bidirectional stream: Send() 和 Recv() 独立，不互相阻塞")
	fmt.Println("2. CloseSend() 关闭发送端（客户端无法再发送，但可以继续接收）")
	fmt.Println("3. 服务端处理完所有任务后 return nil，客户端 Recv() 收到 io.EOF")
	fmt.Println("4. Context 超时控制整个 stream 的生命周期")
	fmt.Println("5. 生产环境建议：处理 backpressure（服务端慢时客户端不要发太快）")

	_ = context.Background
	_ = io.EOF
	_ = log.Println
	_ = time.Now
}
