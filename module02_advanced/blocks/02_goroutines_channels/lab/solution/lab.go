package lab

import "sync"

// CollectSquares 使用 worker 并发计算输入数字的平方，并保持输入顺序。
func CollectSquares(values []int, workers int) []int {
	if workers < 1 {
		workers = 1
	}
	results := make([]int, len(values))
	jobs := make(chan int)
	var wg sync.WaitGroup

	if workers > len(values) {
		workers = len(values)
	}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range jobs {
				results[index] = values[index] * values[index]
			}
		}()
	}

	for index := range values {
		jobs <- index
	}
	close(jobs)
	wg.Wait()
	return results
}
