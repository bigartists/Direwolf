package message

import (
	"github.com/bigartists/Direwolf/internal/module/conversation"
	"gorm.io/gorm"
	"time"
)

type IMessageService interface {
	CreateQuestionMessage(newQuestion *Message) (*Message, error)
	CreateModelAnswerMessage(newAnswer *Message) error
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

func (m *MessageService) CreateModelAnswerMessage(newAnswer *Message) error {

	return m.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(newAnswer).Error
		if err != nil {
			return err
		}

		// 更新 conversation表中 的 last_message_at
		return tx.Model(&conversation.Conversation{}).
			Where("session_id = ?", newAnswer.ConversationSessionID).
			Updates(map[string]interface{}{
				"last_message_at": time.Now(),
			}).Error
	})

}

func ProvideMessageService(repo IMessageRepo, db *gorm.DB) IMessageService {
	return &MessageService{repo: repo, db: db}
}
