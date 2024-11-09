package conversation

import (
	"github.com/bigartists/Direwolf/internal/module/conversation_model"
	"time"
)

type ConversationStatus string

const (
	ConversationStatusActive   ConversationStatus = "active"
	ConversationStatusArchived ConversationStatus = "archived"
	ConversationStatusDeleted  ConversationStatus = "deleted"
)

type Conversation struct {
	ID            int64                                   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	SessionID     string                                  `json:"session_id" gorm:"column:session_id;not null;unique"`
	UserID        int64                                   `json:"user_id" gorm:"column:user_id;not null"`
	Title         string                                  `json:"title" gorm:"column:title"`
	Status        ConversationStatus                      `json:"status" gorm:"column:status;type:ENUM('active', 'archived', 'deleted');default:'active'"`
	Models        []*conversation_model.ConversationModel `json:"models" gorm:"foreignKey:ConversationSessionID;references:SessionID"`
	LastMessageAt time.Time                               `json:"last_message_at" gorm:"column:last_message_at"`
	TotalMessages int                                     `json:"total_messages" gorm:"column:total_messages;default:0"`
	CreateTime    time.Time                               `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateTime    time.Time                               `json:"update_time" gorm:"column:update_time;autoUpdateTime;type:datetime(0);"`
}

func (c *Conversation) TableName() string {
	return "conversations"
}

func NewConversation(attrs ...ConversationAttrFunc) *Conversation {
	c := &Conversation{}
	ConversationAttrFuncs(attrs).apply(c)
	return c
}

func (c *Conversation) Mutate(attrs ...ConversationAttrFunc) *Conversation {
	ConversationAttrFuncs(attrs).apply(c)
	return c
}
