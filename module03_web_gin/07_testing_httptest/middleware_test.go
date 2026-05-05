package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// LoggingMiddleware wraps an http.Handler with request logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("method=%s path=%s duration_ms=%d", r.Method, r.URL.Path, time.Since(start).Milliseconds())
	})
}

// CorrelationMiddleware adds a request ID header
func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", "test-request-id")
		next.ServeHTTP(w, r)
	})
}

func TestLoggingMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	wrapped := LoggingMiddleware(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	wrapped.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("status: got %d, want 200", resp.StatusCode)
	}
}

func TestCorrelationMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	wrapped := CorrelationMiddleware(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	wrapped.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("status: got %d, want 200", w.Code)
	}
	if w.Header().Get("X-Request-ID") != "test-request-id" {
		t.Errorf("X-Request-ID: got %q, want 'test-request-id'", w.Header().Get("X-Request-ID"))
	}
}

func TestHelloHandlerViaMiddleware(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	if w.Code != 200 {
		t.Errorf("status: got %d, want 200", w.Code)
	}
	body := w.Body.String()
	if body != "Hello from test handler\n" {
		t.Errorf("body: got %q, want %q", body, "Hello from test handler\n")
	}
}

func TestMiddlewareChain(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	// Chain: Logging -> Correlation -> handler
	chain := LoggingMiddleware(CorrelationMiddleware(handler))

	req := httptest.NewRequest("GET", "/chain", nil)
	w := httptest.NewRecorder()
	chain.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("status: got %d, want 200", w.Code)
	}
	// Correlation middleware should set the header
	if w.Header().Get("X-Request-ID") != "test-request-id" {
		t.Error("middleware chain: expected X-Request-ID header")
	}
}
