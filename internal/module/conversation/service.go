package conversation

import (
	"context"
	"github.com/bigartists/Direwolf/internal/module/message"
	"gorm.io/gorm"
	"time"
)

type ConversationService struct {
	repo        IConversationRepo
	messageRepo message.IMessageRepo
	db          *gorm.DB
}

func (c *ConversationService) IsConversationExist(ctx context.Context, conversationSessionID string) (bool, error) {
	exist, err := c.repo.IsConversationExist(ctx, conversationSessionID)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (c *ConversationService) CreateConversation(ctx context.Context, req CreateConversationRequest, userId int64) (*Conversation, error) {

	conversation, err := c.repo.CreateConversation(ctx, &req, userId)
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (c *ConversationService) GetConversationHistory(ctx context.Context, userID int64) ([]*Conversation, error) {
	history, err := c.repo.GetConversationHistory(ctx, userID)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (c *ConversationService) GetConversationDetail(ctx context.Context, userID int64, session_id string) (*ConversationDetail, error) {

	detail, err := c.repo.GetConversationDetail(ctx, userID, session_id)
	if err != nil {
		return nil, err
	}
	messages, err := c.messageRepo.FindMessageBySessionID(session_id)
	if err != nil {
		return nil, err
	}

	println("messages===============>", messages)

	conversationDetail := &ConversationDetail{}
	conversationDetail.SessionID = detail.SessionID
	conversationDetail.Title = detail.Title
	conversationDetail.Maas = detail.Maas
	conversationDetail.LastMessageAt = detail.LastMessageAt
	conversationDetail.Messages = messages

	return conversationDetail, nil
}

func (c *ConversationService) CreateModelAnswerMessage(newAnswer *message.Message) error {

	return c.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(newAnswer).Error
		if err != nil {
			return err
		}
		// 更新 conversation表中 的 last_message_at
		return tx.Model(&Conversation{}).
			Where("session_id = ?", newAnswer.ConversationSessionID).
			Updates(map[string]interface{}{
				"last_message_at": time.Now(),
			}).Error
	})
}

type IConversationService interface {
	CreateConversation(ctx context.Context, req CreateConversationRequest, userId int64) (*Conversation, error)
	GetConversationHistory(ctx context.Context, userID int64) ([]*Conversation, error)
	GetConversationDetail(ctx context.Context, userID int64, session_id string) (*ConversationDetail, error)
	IsConversationExist(ctx context.Context, conversationSessionID string) (bool, error)
	CreateModelAnswerMessage(newAnswer *message.Message) error
}

func ProvideConversationService(repo IConversationRepo, messageRepo message.IMessageRepo, db *gorm.DB) IConversationService {
	return &ConversationService{repo: repo, messageRepo: messageRepo, db: db}
}
