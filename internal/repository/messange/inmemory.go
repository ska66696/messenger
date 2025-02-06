package messange

import (
	"errors"
	"messenger/internal/entity"
	"sync"
)

type inMemoryRepository struct {
	messanges map[string]*entity.Massege
	mu        sync.RWMutex
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		messanges: make(map[string]*entity.Massege),
	}
}

func (r *inMemoryRepository) Create(messange *entity.Massege) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.messanges[messange.ID]; exists {
		return errors.New("сообщение с таким Id уже существует")
	}

	r.messanges[messange.ID] = messange
	return nil
}

func (r *inMemoryRepository) GetByID(id string) (*entity.Massege, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	messange, exists := r.messanges[id]
	if !exists {
		return nil, errors.New("сообщение не найдено")
	}

	return messange, nil
}

func (r *inMemoryRepository) FindMessengesByChatID(chatID string) ([]*entity.Massege, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var chatMessages []*entity.Massege
	for _, message := range r.messanges {
		if message.ChatID == chatID {
			chatMessages = append(chatMessages, message)
		}
	}

	return chatMessages, nil
}
