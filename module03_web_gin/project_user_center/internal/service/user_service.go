package service

import (
	"errors"
	"iotestgo/module03_web_gin/project_user_center/internal/model"
	"iotestgo/module03_web_gin/project_user_center/internal/repository"
	"iotestgo/module03_web_gin/project_user_center/pkg/utils"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, password string) (*model.User, error) {
	// Simple validation
	if len(username) < 3 || len(password) < 6 {
		return nil, errors.New("invalid input")
	}

	u := &model.User{
		Username: username,
		Password: password, // In production, hash this!
	}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) Login(username, password string) (string, error) {
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	if u.Password != password { // In production, compare hash
		return "", ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(u.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
