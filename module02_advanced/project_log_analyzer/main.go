package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulated Log Entry
type LogEntry struct {
	ID      int
	Content string
	Level   string // INFO, ERROR, WARN
}

// LogGenerator simulates reading from a large file
func LogGenerator(ctx context.Context, out chan<- LogEntry, count int) {
	defer close(out)
	for id := 1; id <= count; id++ {
		select {
		case <-ctx.Done():
			return
		default:
			// Simulate random logs
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
			out <- log
			// time.Sleep(1 * time.Millisecond) // Removed sleep for benchmark
		}
	}
}

// LogProcessor processes logs
func LogProcessor(id int, in <-chan LogEntry, errorsCh chan<- LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for log := range in {
		// Simulate processing
		// time.Sleep(1 * time.Millisecond) // Removed sleep for benchmark
		if log.Level == "ERROR" {
			// fmt.Printf("[Processor %d] Found ERROR: %s\n", id, log.Content) // Remove print for benchmark
			errorsCh <- log
		}
	}
}

func RunPipeline(numProcessors int, logCount int) int {
	// Channels
	logsCh := make(chan LogEntry, 100)
	errorsCh := make(chan LogEntry, 100)

	// Context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Start Generator (Producer)
	go LogGenerator(ctx, logsCh, logCount)

	// 2. Start Processors (Consumers)
	var wg sync.WaitGroup
	for i := 1; i <= numProcessors; i++ {
		wg.Add(1)
		go LogProcessor(i, logsCh, errorsCh, &wg)
	}

	// 3. Error Collector
	errorCount := 0
	var collectorWg sync.WaitGroup
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		for range errorsCh {
			errorCount++
		}
	}()

	// Wait for processors to finish
	wg.Wait()
	close(errorsCh) // Close error channel so collector can finish

	collectorWg.Wait() // Wait for collector
	return errorCount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	errs := RunPipeline(3, 100)
	fmt.Printf("Processed 100 logs, found %d errors in %v\n", errs, time.Since(start))
}
