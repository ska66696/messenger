package usersvc

import "messenger/internal/entity"

type UseCase interface {
	FindUserByUsername(username string) (*entity.User, error)
}
