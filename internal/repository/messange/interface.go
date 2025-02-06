package messange

import "messenger/internal/entity"

type Repository interface {
	Create(user *entity.Massege) error
	GetByID(id string) (*entity.Massege, error)
	FindMessengesByChatID(chatID string) ([]*entity.Massege, error)
}
