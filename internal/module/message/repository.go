package message

import "gorm.io/gorm"

type IMessageRepo interface {
	FindMessageByParams(conversationSessionID string, modelID int64, messageType MessageType) (*Message, error)
	GetMaxSequenceNumber(conversationSessionID string) (int, error)
	CreateMessage(message *Message) error
}

type MessageRepo struct {
	db *gorm.DB
}

func (m *MessageRepo) CreateMessage(message *Message) error {
	return m.db.Create(message).Error
}

func (m *MessageRepo) FindMessageByParams(conversationSessionID string, modelID int64, messageType MessageType) (*Message, error) {
	var question Message
	err := m.db.Where("conversation_session_id = ? AND message_type = ? AND model_id = ?",
		conversationSessionID, messageType, modelID).
		Order("created_at DESC").
		First(&question).Error
	return &question, err
}

func (m *MessageRepo) GetMaxSequenceNumber(conversationSessionID string) (int, error) {
	var maxSeq int
	m.db.Model(&Message{}).
		Where("conversation_session_id = ?", conversationSessionID).
		Select("COALESCE(MAX(sequence_number), 0)").
		Scan(&maxSeq)

	return maxSeq, nil
}

func ProvideMessageRepo(db *gorm.DB) IMessageRepo {
	return &MessageRepo{db: db}
}
