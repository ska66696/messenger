package chat

import (
	"errors"
	"messenger/internal/entity"
	"sync"
)

type inMemoryRepository struct {
	chats map[string]*entity.Chat
	mu    sync.RWMutex
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		chats: make(map[string]*entity.Chat),
	}
}

func (r *inMemoryRepository) Create(chat *entity.Chat) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.chats[chat.ID]; exists {
		return errors.New("чат с таким Id уже существует")
	}

	r.chats[chat.ID] = chat
	return nil
}

func (r *inMemoryRepository) GetByID(id string) (*entity.Chat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	chat, exists := r.chats[id]
	if !exists {
		return nil, errors.New("чат не найден")
	}

	return chat, nil
}

func (r *inMemoryRepository) FindChatByUserID(userID string) ([]*entity.Chat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var userChats []*entity.Chat
	for _, chat := range r.chats {
		for _, chatUserID := range chat.UserIDs {
			if chatUserID == userID {
				userChats = append(userChats, chat)
				break
			}
		}
	}
	return userChats, nil
}
