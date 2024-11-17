package message

import (
	"fmt"
	"gorm.io/gorm"
)

type IMessageRepo interface {
	FindMessageByParams(conversationSessionID string, modelID int64, messageType MessageType) (*Message, error)
	GetMaxSequenceNumber(conversationSessionID string, messageType MessageType) (int, error)
	CreateMessage(message *Message) error
	FindMessageBySessionID(conversationSessionID string) ([]*Message, error)
}

type MessageRepo struct {
	db *gorm.DB
}

func (m *MessageRepo) FindMessageBySessionID(conversationSessionID string) ([]*Message, error) {
	fmt.Println("conversationSessionID===repo", conversationSessionID)
	var messages []*Message
	err := m.db.Where("conversation_session_id = ?", conversationSessionID).
		//Order("sequence_number DESC").
		Find(&messages).Error
	return messages, err
}

func (m *MessageRepo) CreateMessage(message *Message) error {
	return m.db.Create(message).Error
}

func (m *MessageRepo) FindMessageByParams(conversationSessionID string, modelID int64, messageType MessageType) (*Message, error) {
	var question Message
	err := m.db.Where("conversation_session_id = ? AND message_type = ? AND model_id = ?",
		conversationSessionID, messageType, modelID).
		Order("sequence_number DESC").
		First(&question).Error
	return &question, err
}

func (m *MessageRepo) GetMaxSequenceNumber(conversationSessionID string, messageType MessageType) (int, error) {
	var maxSeq int
	m.db.Model(&Message{}).
		Where("conversation_session_id = ? and message_type = ?", conversationSessionID, messageType).
		Select("COALESCE(MAX(sequence_number), 0)").
		Scan(&maxSeq)

	return maxSeq, nil
}

func ProvideMessageRepo(db *gorm.DB) IMessageRepo {
	return &MessageRepo{db: db}
}
