package chat

import "messenger/internal/entity"

type UseCase interface {
	CreateChat(userIDs string) (*entity.Chat, error)
	SendMessage(chatID, senderID, text string) (*entity.Massege, error)
	GetChatMessages(chatID string) ([]*entity.Massege, error)
	GetUserChats(userID string) ([]*entity.Chat, error)
	FindChatByID(chatID string) (*entity.Massege, error)
}
