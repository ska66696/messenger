package auth

import "messenger/internal/entity"

type UseCase interface {
	ResterUser(username, email, password string) (*entity.User, error)
	LoginUser(username, password string) (*entity.User, error)
}
