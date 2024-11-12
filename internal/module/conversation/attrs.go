package conversation

import (
	"github.com/bigartists/Direwolf/internal/module/conversation_maas"
	"time"
)

type ConversationAttrFunc func(c *Conversation)

type ConversationAttrFuncs []ConversationAttrFunc

func WithID(id int64) ConversationAttrFunc {
	return func(c *Conversation) {
		c.ID = id
	}
}

func WithUserID(userID int64) ConversationAttrFunc {
	return func(c *Conversation) {
		c.UserID = userID
	}
}

func WithTitle(title string) ConversationAttrFunc {
	return func(c *Conversation) {
		c.Title = title
	}
}

func WithStatus(status ConversationStatus) ConversationAttrFunc {
	return func(c *Conversation) {
		c.Status = status
	}
}

func WithModels(models []*conversation_maas.ConversationModel) ConversationAttrFunc {
	return func(c *Conversation) {
		c.Models = models
	}
}

func WithLastMessageAt(lastMessageAt time.Time) ConversationAttrFunc {
	return func(c *Conversation) {
		c.LastMessageAt = lastMessageAt
	}
}

func WithTotalMessages(totalMessages int) ConversationAttrFunc {
	return func(c *Conversation) {
		c.TotalMessages = totalMessages
	}
}

func (this ConversationAttrFuncs) apply(c *Conversation) {
	for _, f := range this {
		f(c)
	}
}
