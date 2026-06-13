package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"iotestgo/module03_web_gin/project_user_center/internal/repository"
	"iotestgo/module03_web_gin/project_user_center/internal/service"

	"github.com/gin-gonic/gin"
)

func TestRegisterAndLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	userHandler := NewUserHandler(svc)

	router := gin.New()
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	registerBody := bytes.NewBufferString(`{"username":"alice","password":"secret123"}`)
	registerReq := httptest.NewRequest(http.MethodPost, "/register", registerBody)
	registerReq.Header.Set("Content-Type", "application/json")
	registerResp := httptest.NewRecorder()

	router.ServeHTTP(registerResp, registerReq)
	if registerResp.Code != http.StatusCreated {
		t.Fatalf("expected register status %d, got %d: %s", http.StatusCreated, registerResp.Code, registerResp.Body.String())
	}

	var registerPayload struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
	if err := json.Unmarshal(registerResp.Body.Bytes(), &registerPayload); err != nil {
		t.Fatalf("decode register response: %v", err)
	}
	if registerPayload.ID != 1 || registerPayload.Username != "alice" {
		t.Fatalf("unexpected register payload: %+v", registerPayload)
	}

	loginBody := bytes.NewBufferString(`{"username":"alice","password":"secret123"}`)
	loginReq := httptest.NewRequest(http.MethodPost, "/login", loginBody)
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp := httptest.NewRecorder()

	router.ServeHTTP(loginResp, loginReq)
	if loginResp.Code != http.StatusOK {
		t.Fatalf("expected login status %d, got %d: %s", http.StatusOK, loginResp.Code, loginResp.Body.String())
	}

	var loginPayload struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(loginResp.Body.Bytes(), &loginPayload); err != nil {
		t.Fatalf("decode login response: %v", err)
	}
	if loginPayload.Token == "" {
		t.Fatal("expected login token")
	}
}
