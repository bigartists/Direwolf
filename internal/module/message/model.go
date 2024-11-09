package message

import "time"

type MessageType string

var MessageTypes = map[MessageType]string{
	"User":      "user",
	"Assistant": "assistant",
}

type Message struct {
	ID                    int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ConversationSessionID string    `json:"conversation_session_id" gorm:"column:conversation_session_id;not null"`
	MessageType           string    `json:"message_type" gorm:"column:message_type;type:ENUM('user', 'assistant');not null"`
	Content               string    `json:"content" gorm:"column:content;not null"`
	Context               string    `json:"context" gorm:"column:context;"`
	Expand                string    `json:"expand" gorm:"column:expand;"`
	ModelID               *int64    `json:"model_id" gorm:"column:model_id"`
	SequenceNumber        int       `json:"sequence_number" gorm:"column:sequence_number;not null"`
	CreateTime            time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	ParentQuestionID      int64     `json:"parent_question_id" gorm:"column:parent_question_id"`
}

func (m *Message) TableName() string {
	return "messages"
}

func NewMessage(attrs ...MessageAttrFunc) *Message {
	u := &Message{}
	MessageAttrFuncs(attrs).apply(u)
	return u
}

func (m *Message) Mutate(attrs ...MessageAttrFunc) *Message {
	MessageAttrFuncs(attrs).apply(m)
	return m
}
