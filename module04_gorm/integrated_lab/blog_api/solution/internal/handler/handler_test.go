package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"iotestgo/module04_gorm/integrated_lab/blog_api/solution/internal/service"
)

type fakeService struct{ deleteErr error }

func (fakeService) CreatePost(context.Context, string, string, []string) (*service.PostResponse, error) {
	return &service.PostResponse{ID: 1, Title: "Title", Content: "Content", Tags: []service.TagResponse{}, Comments: []service.CommentResponse{}}, nil
}
func (fakeService) CreatePostWithComment(context.Context, string, string, string, []string) (*service.PostResponse, error) {
	return &service.PostResponse{ID: 1}, nil
}
func (fakeService) ListPosts(context.Context) ([]service.PostResponse, error) {
	return []service.PostResponse{}, nil
}
func (f fakeService) DeletePost(context.Context, uint) error { return f.deleteErr }

func TestCreateAndList(t *testing.T) {
	router := Router(fakeService{})
	request := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewBufferString(`{"title":"Title","content":"Content","tags":["go"]}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusCreated {
		t.Fatalf("POST status=%d body=%s", response.Code, response.Body.String())
	}
	response = httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/posts", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("GET status=%d", response.Code)
	}
}

func TestDeleteStatuses(t *testing.T) {
	tests := []struct {
		path string
		err  error
		want int
	}{{"/posts/nope", nil, 400}, {"/posts/1", service.ErrNotFound, 404}, {"/posts/1", errors.New("db"), 500}, {"/posts/1", nil, 204}}
	for _, test := range tests {
		response := httptest.NewRecorder()
		Router(fakeService{deleteErr: test.err}).ServeHTTP(response, httptest.NewRequest(http.MethodDelete, test.path, nil))
		if response.Code != test.want {
			t.Fatalf("DELETE %s err=%v status=%d, want %d", test.path, test.err, response.Code, test.want)
		}
	}
}
