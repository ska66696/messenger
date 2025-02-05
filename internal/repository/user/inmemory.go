package user

import (
	"errors"
	"messenger/internal/entity"
	"sync"
)

type inMemoryRepository struct {
	users map[string]*entity.User
	mu    sync.RWMutex
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *inMemoryRepository) Create(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("пользователь с таким Id уже существует")
	}

	r.users[user.ID] = user
	return nil
}

func (r *inMemoryRepository) GetByID(id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("пользователь не найден")
	}

	return user, nil
}

func (r *inMemoryRepository) GetByUsername(username string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, errors.New("пользователь не найден")
}

func (r *inMemoryRepository) GetByEmail(email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("пользователь не найден")
}
