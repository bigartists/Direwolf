package message

import (
	"gorm.io/gorm"
)

type IMessageService interface {
	CreateQuestionMessage(newQuestion *Message) (*Message, error)
	GetMaxSequenceNumber(conversationSessionID string, messageType MessageType) (int, error)
}

type MessageService struct {
	repo IMessageRepo
	db   *gorm.DB
}

func (m *MessageService) GetMaxSequenceNumber(conversationSessionID string, messageType MessageType) (int, error) {
	return m.repo.GetMaxSequenceNumber(conversationSessionID, messageType)
}

func (m *MessageService) CreateQuestionMessage(newQuestion *Message) (*Message, error) {
	err := m.repo.CreateMessage(newQuestion)
	if err != nil {
		return nil, err
	}
	return newQuestion, nil
}

func ProvideMessageService(repo IMessageRepo, db *gorm.DB) IMessageService {
	return &MessageService{repo: repo, db: db}
}
