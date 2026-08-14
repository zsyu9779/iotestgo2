package main

import (
	"flag"
	"log"
	"net/http"
)

/*
go run . -mode=server -proto=http
*/
func main() {
	flag.Parse()
	runServer(":8080")

}

func runServer(addr string) {
	log.Printf("Starting HTTP server on %s", addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	mux.HandleFunc("/slow", Slow)
	// Apply middleware
	handler := Logging(mux)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}

}
