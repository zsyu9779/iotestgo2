package repository

import (
	"errors"
	"iotestgo/module03_web_gin/project_user_center/internal/model"
	"sync"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByID(id int) (*model.User, error)
}

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*model.User
	lastID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *InMemoryUserRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Username]; exists {
		return ErrUserAlreadyExists
	}

	r.lastID++
	user.ID = r.lastID
	// Store a copy or pointer? Here we store pointer but ensure we don't mutate external state unexpectedly
	// For in-memory, let's just store the pointer
	r.users[user.Username] = user
	return nil
}

func (r *InMemoryUserRepository) FindByUsername(username string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[username]; exists {
		return user, nil
	}
	return nil, ErrUserNotFound
}

func (r *InMemoryUserRepository) FindByID(id int) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}
