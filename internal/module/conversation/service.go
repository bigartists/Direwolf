package conversation

import (
	"context"
)

type ConversationService struct {
	repo IConversationRepo
}

func (c *ConversationService) IsConversationExist(ctx context.Context, conversationSessionID string) (bool, error) {
	exist, err := c.repo.IsConversationExist(ctx, conversationSessionID)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (c *ConversationService) CreateConversation(ctx context.Context, req CreateConversationRequest) (*Conversation, error) {

	conversation, err := c.repo.CreateConversation(ctx, &req)
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

func (c *ConversationService) GetConversationDetail(ctx context.Context, userID int64, conversationID int64) (*Conversation, error) {
	detail, err := c.repo.GetConversationDetail(ctx, userID, conversationID)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

type IConversationService interface {
	CreateConversation(ctx context.Context, req CreateConversationRequest) (*Conversation, error)
	GetConversationHistory(ctx context.Context, userID int64) ([]*Conversation, error)
	GetConversationDetail(ctx context.Context, userID int64, conversationID int64) (*Conversation, error)
	IsConversationExist(ctx context.Context, conversationSessionID string) (bool, error)
}

func ProvideConversationService(repo IConversationRepo) IConversationService {
	return &ConversationService{repo: repo}
}
