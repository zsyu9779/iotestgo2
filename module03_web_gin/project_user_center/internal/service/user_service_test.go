package service

import (
	"iotestgo/module03_web_gin/project_user_center/internal/repository"
	"testing"
)

func TestUserService_RegisterAndLogin(t *testing.T) {
	// Setup
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)

	// Test Register
	username := "testuser"
	password := "password123"

	user, err := svc.Register(username, password)
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}
	if user.Username != username {
		t.Errorf("expected username %s, got %s", username, user.Username)
	}

	// Test Login
	token, err := svc.Login(username, password)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	if token == "" {
		t.Error("expected token, got empty string")
	}

	// Test Invalid Login
	_, err = svc.Login(username, "wrongpassword")
	if err != ErrInvalidCredentials {
		t.Errorf("expected ErrInvalidCredentials, got %v", err)
	}
}
