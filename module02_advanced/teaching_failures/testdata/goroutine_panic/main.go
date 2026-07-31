package main

func main() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		panic("child goroutine failed")
	}()
	<-done
}
