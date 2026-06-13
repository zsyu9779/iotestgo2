package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func main() {
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		n := 38
		if value := r.URL.Query().Get("n"); value != "" {
			parsed, err := strconv.Atoi(value)
			if err != nil || parsed < 0 || parsed > 45 {
				http.Error(w, "n must be an integer between 0 and 45", http.StatusBadRequest)
				return
			}
			n = parsed
		}

		// Intentionally inefficient for local pprof demonstration.
		fmt.Fprintf(w, "fib(%d)=%d\n", n, fib(n))
	})

	log.Println("listening on :8892")
	log.Println("try: curl 'http://localhost:8892/fib?n=38'")
	log.Println("pprof: http://localhost:8892/debug/pprof/")
	log.Fatal(http.ListenAndServe(":8892", nil))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
