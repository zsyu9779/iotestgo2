// 分布式计算项目 - 服务端
// 接收客户端流式数据，goroutine 池并发计算，流式返回结果
//
// 启动：go run server/main.go
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"

	pb "iotestgo/module05_grpc/project_distributed_compute/proto/computepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedDistributedComputeServer
	workerCount int
}

// Process 实现双向流 RPC：接收任务 → 并发计算 → 返回结果
func (s *server) Process(stream pb.DistributedCompute_ProcessServer) error {
	tasksCh := make(chan *pb.ComputeTask, 100)
	var wg sync.WaitGroup

	// 启动 worker 池并发处理任务
	for i := 0; i < s.workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range tasksCh {
				result := compute(task)
				log.Printf("[Worker-%d] 完成: task=%s op=%s status=%s", workerID, task.GetTaskId(), task.GetOperation(), result.GetStatus())
				if err := stream.Send(result); err != nil {
					log.Printf("[Worker-%d] 发送结果失败: %v", workerID, err)
				}
			}
		}(i)
	}

	// 接收 goroutine：读取客户端发送的任务
	log.Println("服务端开始接收任务...")
	for {
		task, err := stream.Recv()
		if err == io.EOF {
			log.Println("客户端发送完毕 (EOF)，等待所有任务处理完成...")
			close(tasksCh)
			wg.Wait()
			log.Println("所有任务处理完成，流正常结束")
			return nil // return nil 表示流正常结束
		}
		if err != nil {
			log.Printf("接收错误: %v", err)
			close(tasksCh)
			return err
		}
		log.Printf("收到任务: id=%s op=%s numbers=%v", task.GetTaskId(), task.GetOperation(), task.GetNumbers())
		tasksCh <- task
	}
}

// compute 计算引擎
func compute(task *pb.ComputeTask) *pb.ComputeResult {
	r := &pb.ComputeResult{
		TaskId:    task.GetTaskId(),
		Operation: task.GetOperation(),
		Status:    "done",
	}

	numbers := task.GetNumbers()
	if len(numbers) == 0 {
		r.Status = "error"
		r.Message = "no numbers provided"
		return r
	}

	switch task.GetOperation() {
	case "sum":
		var sum int64
		for _, n := range numbers {
			sum += n
		}
		r.Value = float64(sum)
	case "avg":
		var sum int64
		for _, n := range numbers {
			sum += n
		}
		r.Value = float64(sum) / float64(len(numbers))
	case "max":
		r.Value = float64(numbers[0])
		for _, n := range numbers[1:] {
			if float64(n) > r.Value {
				r.Value = float64(n)
			}
		}
	case "min":
		r.Value = float64(numbers[0])
		for _, n := range numbers[1:] {
			if float64(n) < r.Value {
				r.Value = float64(n)
			}
		}
	case "stddev":
		mean := float64(0)
		for _, n := range numbers {
			mean += float64(n)
		}
		mean /= float64(len(numbers))
		variance := float64(0)
		for _, n := range numbers {
			diff := float64(n) - mean
			variance += diff * diff
		}
		variance /= float64(len(numbers))
		r.Value = math.Sqrt(variance)
	case "median":
		sorted := make([]int64, len(numbers))
		copy(sorted, numbers)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
		mid := len(sorted) / 2
		if len(sorted)%2 == 0 {
			r.Value = float64(sorted[mid-1]+sorted[mid]) / 2
		} else {
			r.Value = float64(sorted[mid])
		}
	default:
		r.Status = "error"
		r.Message = fmt.Sprintf("unknown operation: %s", task.GetOperation())
	}

	return r
}

func main() {
	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDistributedComputeServer(s, &server{workerCount: 4})
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

	_ = status.New
	_ = codes.OK
}
