package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	if w.Code != 200 {
		t.Fatalf("status = %d; want 200", w.Code)
	}
}
