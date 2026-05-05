// 分布式计算服务端：接收客户端流式数据，实时计算，流式返回结果
//
// 教学要点：
// - Bidirectional streaming 实战
// - 并发安全的计算引擎
// - Context 取消处理
// - goroutine + channel 协作
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"sort"
	"sync"
)

func main() {
	fmt.Println("=== 分布式计算项目 - 服务端 ===")
	fmt.Println()
	fmt.Println("项目描述：客户端批量流式发送计算任务 → 服务端流式返回计算结果")
	fmt.Println()

	fmt.Println("--- 架构设计 ---")
	fmt.Println("  客户端                          服务端")
	fmt.Println("  ┌──────────┐                   ┌──────────────────┐")
	fmt.Println("  │  发送 Task1 │ ──────────────▶ │ Process()        │")
	fmt.Println("  │  发送 Task2 │ ──────────────▶ │   ├─ 解析任务       │")
	fmt.Println("  │  发送 Task3 │ ──────────────▶ │   ├─ 并发计算       │")
	fmt.Println("  │  ← 结果1   │ ◀────────────── │   └─ Send(result)  │")
	fmt.Println("  │  ← 结果2   │ ◀────────────── │                   │")
	fmt.Println("  │  ← 结果3   │ ◀────────────── │                   │")
	fmt.Println("  │  关闭发送    │ ──────────────▶ │  return nil       │")
	fmt.Println("  └──────────┘                   └──────────────────┘")
	fmt.Println()

	fmt.Println("--- 核心伪代码 ---")
	fmt.Println()
	fmt.Println(`func (s *Server) Process(stream DistributedCompute_ProcessServer) error {
    // 接收 goroutine：从客户端读任务，放入 channel
    tasksCh := make(chan *ComputeTask, 100)
    go func() {
        defer close(tasksCh)
        for {
            task, err := stream.Recv()
            if err == io.EOF { return }
            if err != nil { return }
            tasksCh <- task
        }
    }()

    // 处理逻辑：从 channel 读任务，计算结果，发送回客户端
    for task := range tasksCh {
        result := compute(task)
        if err := stream.Send(result); err != nil {
            return err
        }
    }
    return nil
}`)
	fmt.Println()

	// 演示计算引擎
	showComputeEngine()
}

// showComputeEngine 演示计算逻辑
func showComputeEngine() {
	fmt.Println("--- 计算引擎 ---")
	fmt.Println()

	tasks := []ComputeTask{
		{TaskId: "t1", Numbers: []int64{1, 2, 3, 4, 5}, Operation: "sum"},
		{TaskId: "t2", Numbers: []int64{1, 2, 3, 4, 5}, Operation: "avg"},
		{TaskId: "t3", Numbers: []int64{3, 1, 4, 1, 5, 9, 2, 6}, Operation: "max"},
		{TaskId: "t4", Numbers: []int64{3, 1, 4, 1, 5, 9, 2, 6}, Operation: "min"},
		{TaskId: "t5", Numbers: []int64{1, 2, 3, 4, 5}, Operation: "stddev"},
	}

	// 模拟并发处理
	var wg sync.WaitGroup
	results := make(chan ComputeResult, len(tasks))

	for _, t := range tasks {
		wg.Add(1)
		go func(t ComputeTask) {
			defer wg.Done()
			results <- compute(t)
		}(t)
	}

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Printf("  Task[%s] %s(%v) = %.2f [%s]\n",
			r.TaskID, r.Operation, r.TaskID, r.Value, r.Status)
	}

	_ = io.EOF
	_ = context.Canceled
	_ = log.Println
}

// ========== 计算引擎实现 ==========

type ComputeTask struct {
	TaskID    string
	Numbers   []int64
	Operation string
}

type ComputeResult struct {
	TaskID    string
	Operation string
	Value     float64
	Status    string
	Message   string
}

func compute(task ComputeTask) ComputeResult {
	r := ComputeResult{
		TaskID:    task.TaskID,
		Operation: task.Operation,
		Status:    "done",
	}

	if len(task.Numbers) == 0 {
		r.Status = "error"
		r.Message = "no numbers provided"
		return r
	}

	switch task.Operation {
	case "sum":
		var sum int64
		for _, n := range task.Numbers {
			sum += n
		}
		r.Value = float64(sum)
	case "avg":
		var sum int64
		for _, n := range task.Numbers {
			sum += n
		}
		r.Value = float64(sum) / float64(len(task.Numbers))
	case "max":
		r.Value = float64(task.Numbers[0])
		for _, n := range task.Numbers[1:] {
			if float64(n) > r.Value {
				r.Value = float64(n)
			}
		}
	case "min":
		r.Value = float64(task.Numbers[0])
		for _, n := range task.Numbers[1:] {
			if float64(n) < r.Value {
				r.Value = float64(n)
			}
		}
	case "stddev":
		mean := float64(0)
		for _, n := range task.Numbers {
			mean += float64(n)
		}
		mean /= float64(len(task.Numbers))

		variance := float64(0)
		for _, n := range task.Numbers {
			diff := float64(n) - mean
			variance += diff * diff
		}
		variance /= float64(len(task.Numbers))
		r.Value = math.Sqrt(variance)
	case "median":
		sorted := make([]int64, len(task.Numbers))
		copy(sorted, task.Numbers)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
		mid := len(sorted) / 2
		if len(sorted)%2 == 0 {
			r.Value = float64(sorted[mid-1]+sorted[mid]) / 2
		} else {
			r.Value = float64(sorted[mid])
		}
	default:
		r.Status = "error"
		r.Message = fmt.Sprintf("unknown operation: %s", task.Operation)
	}

	return r
}

_ = math.Sqrt // 确保 math 被引用
