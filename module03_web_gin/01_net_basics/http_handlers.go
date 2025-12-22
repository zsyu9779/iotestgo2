package main

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("method=%s path=%s status=%d dur_ms=%d", r.Method, r.URL.Path, 200, time.Since(start).Milliseconds())
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"hello"}`))
}

func Slow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(300 * time.Millisecond):
		w.Write([]byte(`{"done":true}`))
	case <-ctx.Done():
		http.Error(w, "canceled", 499)
	}
}
