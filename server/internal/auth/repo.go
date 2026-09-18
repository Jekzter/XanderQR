package auth

import (
	"sync"

	"github.com/google/uuid"
)

type Repository interface {
	FindByEmail(email string) (User, error)
	Create(user User) (User, error)
}

type repository struct {
	mu    sync.Mutex
	users map[string]User
}

func NewRepository() Repository {
	return &repository{
		users: make(map[string]User),
	}
}

func (r *repository) FindByEmail(email string) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[email]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Email]; exists {
		return User{}, ErrEmailAlreadyExist
	}

	user.ID = uuid.NewString()
	r.users[user.Email] = user
	return user, nil
}
