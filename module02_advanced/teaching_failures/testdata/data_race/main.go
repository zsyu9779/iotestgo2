package main

import "sync"

func main() {
	var value int
	var workers sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < 4; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			<-start
			for j := 0; j < 100000; j++ {
				value++
			}
		}()
	}
	close(start)
	workers.Wait()
	_ = value
}
