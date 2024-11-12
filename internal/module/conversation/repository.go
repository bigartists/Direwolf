package conversation

import (
	"context"
	"github.com/bigartists/Direwolf/internal/module/conversation_maas"
	"github.com/bigartists/Direwolf/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type IConversationRepo interface {
	CreateConversation(ctx context.Context, req *CreateConversationRequest, userId int64) (*Conversation, error)
	GetConversationHistory(ctx context.Context, userID int64) ([]*Conversation, error)
	GetConversationDetail(ctx context.Context, userID int64, conversationID int64) (*Conversation, error)
	IsConversationExist(ctx context.Context, conversationSessionID string) (bool, error)
}

type ConversationRepo struct {
	db *gorm.DB
}

func (c *ConversationRepo) IsConversationExist(ctx context.Context, conversationSessionID string) (bool, error) {
	var count int64
	err := c.db.Model(&Conversation{}).Where("session_id = ?", conversationSessionID).Count(&count).Error
	return count > 0, err
}

func (c *ConversationRepo) CreateConversation(ctx context.Context, req *CreateConversationRequest, userId int64) (*Conversation, error) {
	conversation := &Conversation{
		SessionID:     utils.GenerateUUID(),
		UserID:        userId,
		Title:         req.Title,
		Status:        ConversationStatusActive,
		LastMessageAt: time.Now(),
	}

	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(conversation).Error; err != nil {
			return err
		}
		// 关联模型创建
		for _, maasID := range req.MaasIds {
			convModel := &conversation_maas.ConversationModel{
				ConversationSessionID: conversation.SessionID,
				MaasID:                maasID,
				Status:                "active",
			}
			if err := tx.Create(convModel).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (c *ConversationRepo) GetConversationHistory(ctx context.Context, userID int64) ([]*Conversation, error) {
	var conversations []*Conversation
	err := c.db.Where("user_id=? AND status=?", userID, ConversationStatusActive).Order("last_message_at DESC").Find(&conversations).Error
	return conversations, err
}

func (c *ConversationRepo) GetConversationDetail(ctx context.Context, userID int64, conversationID int64) (*Conversation, error) {
	var conversation Conversation
	err := c.db.Preload("Models").First(&conversation, conversationID).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func ProvideConversationRepo(db *gorm.DB) IConversationRepo {
	return &ConversationRepo{db: db}
}
