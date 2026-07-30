package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// LogEntry 表示一条模拟日志。
type LogEntry struct {
	ID      int
	Content string
	Level   string // INFO、ERROR、WARN
}

// LogGenerator 模拟从大文件读取日志，并在取消时停止发送。
func LogGenerator(ctx context.Context, out chan<- LogEntry, count int) {
	defer close(out)
	for id := 1; id <= count; id++ {
		select {
		case <-ctx.Done():
			return
		default:
			level := "INFO"
			r := rand.Intn(10)
			if r > 8 {
				level = "ERROR"
			} else if r > 6 {
				level = "WARN"
			}

			log := LogEntry{
				ID:      id,
				Content: fmt.Sprintf("Log message content %d", id),
				Level:   level,
			}
			select {
			case <-ctx.Done():
				return
			case out <- log:
			}
		}
	}
}

// LogProcessor 兼容旧的课堂示例，处理不会被取消的输入。
func LogProcessor(id int, in <-chan LogEntry, errorsCh chan<- LogEntry, wg *sync.WaitGroup) {
	logProcessor(context.Background(), in, errorsCh, wg)
}

// logProcessor 处理输入，并响应 Context 取消。
func logProcessor(ctx context.Context, in <-chan LogEntry, errorsCh chan<- LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case log, ok := <-in:
			if !ok {
				return
			}
			if log.Level == "ERROR" {
				select {
				case <-ctx.Done():
					return
				case errorsCh <- log:
				}
			}
		}
	}
}

func CountErrors(entries []LogEntry, numProcessors int) int {
	if numProcessors < 1 {
		return 0
	}
	logsCh := make(chan LogEntry, len(entries))
	errorsCh := make(chan LogEntry, len(entries))

	var wg sync.WaitGroup
	for i := 1; i <= numProcessors; i++ {
		wg.Add(1)
		go LogProcessor(i, logsCh, errorsCh, &wg)
	}

	for _, entry := range entries {
		logsCh <- entry
	}
	close(logsCh)
	wg.Wait()
	close(errorsCh)

	count := 0
	for range errorsCh {
		count++
	}
	return count
}

// Analyze 消费日志输入，返回 ERROR 数量，并保证取消和关闭顺序明确。
func Analyze(ctx context.Context, source <-chan LogEntry, numProcessors int) (int, error) {
	if numProcessors < 1 {
		return 0, fmt.Errorf("worker 数量必须大于 0")
	}
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	errorsCh := make(chan LogEntry, numProcessors)
	var workers sync.WaitGroup
	for i := 0; i < numProcessors; i++ {
		workers.Add(1)
		go logProcessor(ctx, source, errorsCh, &workers)
	}

	countCh := make(chan int, 1)
	go func() {
		count := 0
		for range errorsCh {
			count++
		}
		countCh <- count
	}()

	workers.Wait()
	close(errorsCh)
	count := <-countCh
	if err := ctx.Err(); err != nil {
		return count, err
	}
	return count, nil
}

func RunPipeline(numProcessors int, logCount int) int {
	logsCh := make(chan LogEntry, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go LogGenerator(ctx, logsCh, logCount)
	count, _ := Analyze(ctx, logsCh, numProcessors)
	return count
}

func main() {
	start := time.Now()
	errs := RunPipeline(3, 100)
	fmt.Printf("Processed 100 logs, found %d errors in %v\n", errs, time.Since(start))
}
