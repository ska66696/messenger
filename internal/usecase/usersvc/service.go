package usersvc

import (
	"errors"
	"messenger/internal/entity"
	"messenger/internal/repository/user"
)

type service struct {
	userRepo user.Repository
}

func NewServiceUser(repo user.Repository) UseCase {
	return &service{
		userRepo: repo,
	}
}

func (s *service) FindUserByUsername(username string) (*entity.User, error) {
	if username == "" {
		return nil, errors.New("имя не может быть пустым")
	}
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
