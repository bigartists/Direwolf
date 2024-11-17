package conversation

import (
	"github.com/bigartists/Direwolf/internal/module/conversation_maas"
	"github.com/bigartists/Direwolf/internal/module/message"
	"time"
)

type CreateConversationRequest struct {
	//UserID   int64   `json:"user_id" binding:"required"`
	Title   string  `json:"title" binding:"required"`
	MaasIds []int64 `json:"maas_ids" binding:"required,min=1"`
}

type ConversationDetail struct {
	SessionID     string                                `json:"session_id"`
	Title         string                                `json:"title"`
	Maas          []*conversation_maas.ConversationMaas `json:"maas"`
	LastMessageAt time.Time                             `json:"last_message_at"`
	Messages      []*message.Message                    `json:"messages"`
}
