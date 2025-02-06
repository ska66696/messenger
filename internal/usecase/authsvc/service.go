package authsvc

import (
	"errors"
	"messenger/internal/entity"
	"messenger/internal/repository/user"
)

type service struct {
	userRepo user.Repository
}

func NewServiceAuth(repo user.Repository) UseCase {
	return &service{
		userRepo: repo,
	}
}

func (s *service) ResterUser(username, email, password string) (*entity.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("все поля должны быть заполнены")
	}

	hashedPassword := password

	newUser := &entity.User{
		ID:       "user-id" + username,
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	err := s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *service) LoginUser(username, password string) (*entity.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("все поля должны быть заполнены")
	}

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}
	if user.Password != password {
		return nil, errors.New("неверный пароль")
	}
	return user, err
}
