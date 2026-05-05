// 分布式计算项目 - 客户端
// 批量流式发送计算任务，实时接收计算结果
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

	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50056",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewDistributedComputeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := client.Process(ctx)
	if err != nil {
		log.Fatalf("创建 Process stream 失败: %v", err)
	}

	fmt.Println("=== 分布式计算客户端 ===")
	fmt.Println()

	// 准备测试任务
	tasks := []*pb.ComputeTask{
		{TaskId: "t1-sum", Numbers: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Operation: "sum"},
		{TaskId: "t2-avg", Numbers: []int64{10, 20, 30, 40, 50}, Operation: "avg"},
		{TaskId: "t3-max", Numbers: []int64{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}, Operation: "max"},
		{TaskId: "t4-min", Numbers: []int64{8, 3, 2, 7, 0, 4}, Operation: "min"},
		{TaskId: "t5-stddev", Numbers: []int64{2, 4, 4, 4, 5, 5, 7, 9}, Operation: "stddev"},
		{TaskId: "t6-median", Numbers: []int64{3, 5, 1, 4, 2, 6}, Operation: "median"},
		{TaskId: "t7-avg", Numbers: []int64{100, 200, 300}, Operation: "avg"},
	}

	var wg sync.WaitGroup
	start := time.Now()

	// 接收 goroutine：实时打印计算结果
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				fmt.Println()
				fmt.Println("  流接收完成 (EOF)")
				return
			}
			if err != nil {
				log.Printf("接收错误: %v", err)
				return
			}
			fmt.Printf("  ✓ %-10s | %-8s | %.4f | %s\n",
				result.GetTaskId(), result.GetOperation(), result.GetValue(), result.GetStatus())
		}
	}()

	// 发送任务
	fmt.Printf("发送 %d 个计算任务...\n\n", len(tasks))
	fmt.Println("  Task ID    | 操作     | 结果")
	fmt.Println("  -----------|----------|-----------")

	for _, task := range tasks {
		if err := stream.Send(task); err != nil {
			log.Printf("发送失败: %v", err)
			break
		}
	}

	// 通知服务端：客户端不再发送
	stream.CloseSend()
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println()
	fmt.Printf("=== 完成：%d 个任务，耗时 %v ===\n", len(tasks), elapsed)
	fmt.Println()
	fmt.Println("关键技术点：")
	fmt.Println("  1. Bidirectional streaming: Send() 和 Recv() 独立不阻塞")
	fmt.Println("  2. Worker 池：4 个 goroutine 并发处理计算任务")
	fmt.Println("  3. CloseSend() 关闭发送端 → 服务端 Recv() 收到 io.EOF")
	fmt.Println("  4. 服务端处理完所有任务后 return nil → 客户端收到 io.EOF")
}
