package user

import "messenger/internal/entity"

type Repository interface {
	Create(user *entity.User) error
	GetByID(id string) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}
