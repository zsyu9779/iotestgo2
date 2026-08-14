package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHello(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	srv := httptest.NewServer(Logging(mux))
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/hello")
	if err != nil {
		t.Fatalf("http get: %v", err)
	}
	b, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if string(b) == "" {
		t.Fatalf("empty body")
	}
}
