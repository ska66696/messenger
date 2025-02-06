package chatsvc

import (
	"errors"
	"messenger/internal/entity"
	"messenger/internal/repository/chat"
	"messenger/internal/repository/messange"
	"time"
)

type service struct {
	chatRepo     chat.Repository
	messangeRepo messange.Repository
}

func NewServiceChat(chatRepo chat.Repository, messangeRepo messange.Repository) UseCase {
	return &service{
		chatRepo:     chatRepo,
		messangeRepo: messangeRepo,
	}
}

func (s *service) CreateChat(userIDs []string) (*entity.Chat, error) {
	if len(userIDs) < 2 {
		return nil, errors.New("для создания чата нужно минимум два пользователя")
	}
	for _, ususerID := range userIDs {
		if ususerID == "" {
			return nil, errors.New("ID не может быть пустым")
		}
	}
	newChat := &entity.Chat{
		ID:        "chat-id" + time.Now().String(),
		UserIDs:   userIDs,
		CreatedAt: time.Now().Unix(),
	}
	err := s.chatRepo.Create(newChat)
	if err != nil {
		return nil, err
	}
	return newChat, nil
}
func (s *service) SendMessage(chatID, senderID, text string) (*entity.Massege, error) {
	if chatID == "" || senderID == "" || text == "" {
		return nil, errors.New("ID, cahtID  и text не может быть пустым")
	}
	if len(text) > 4096 {
		return nil, errors.New("слишком длинное сообщение")
	}
	newMessange := &entity.Massege{
		ID:        "messande-id" + time.Now().String(),
		ChatID:    chatID,
		SenderID:  senderID,
		Text:      text,
		CreatedAt: time.Now().Unix(),
	}
	err := s.messangeRepo.Create(newMessange)
	if err != nil {
		return nil, err
	}
	return newMessange, nil
}
func (s *service) GetChatMessages(chatID string) ([]*entity.Massege, error) {
	if chatID == "" {
		return nil, errors.New("cahtID не может быть пустым")
	}
	messanges, err := s.messangeRepo.FindMessengesByChatID(chatID)
	if err != nil {
		return nil, err
	}
	return messanges, nil
}
func (s *service) GetUserChats(userID string) ([]*entity.Chat, error) {
	if userID == "" {
		return nil, errors.New("userID не может быть пустым")
	}
	chats, err := s.chatRepo.FindChatByUserID(userID)
	if err != nil {
		return nil, err
	}
	return chats, nil
}
func (s *service) FindChatByID(chatID string) (*entity.Massege, error) {
	return nil, errors.New("не реализовано")
}
