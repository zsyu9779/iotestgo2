package main

import (
	"context"
	"errors"
	"sync"
)

var ErrNotImplemented = errors.New("not implemented")

type LogEntry struct {
	ID      int
	Content string
	Level   string
}

func LogGenerator(ctx context.Context, out chan<- LogEntry, count int) {
	defer close(out)
}

func LogProcessor(id int, in <-chan LogEntry, errorsCh chan<- LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
}

func CountErrors(entries []LogEntry, numProcessors int) int { return 0 }

func Analyze(ctx context.Context, source <-chan LogEntry, numProcessors int) (int, error) {
	return 0, ErrNotImplemented
}

func RunPipeline(numProcessors, logCount int) int { return 0 }

func main() {}
