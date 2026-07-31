package main

import "sync"

func main() {
	values := make(map[int]int)
	start := make(chan struct{})
	var workers sync.WaitGroup
	for i := 0; i < 16; i++ {
		workers.Add(1)
		go func(id int) {
			defer workers.Done()
			<-start
			for j := 0; j < 100000; j++ {
				values[j] = id
			}
		}(i)
	}
	close(start)
	workers.Wait()
}
