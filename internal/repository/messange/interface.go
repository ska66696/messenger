package messange

import "messenger/internal/entity"

type Repository interface {
	Create(user *entity.Massege) error
	GetByID(id string) (*entity.Massege, error)
	FindMessengesByChatID(ChatID string) ([]*entity.Massege, error)
}
