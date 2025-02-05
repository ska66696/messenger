package chat

import "messenger/internal/entity"

type Repository interface {
	Create(user *entity.Chat) error
	GetByID(id string) (*entity.Chat, error)
	FindChatByUserID(userID string) ([]*entity.Chat, error)
}
