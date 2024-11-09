package conversation_model

import "time"

type ConversationModel struct {
	ID                    int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ConversationSessionID string    `json:"conversation_session_id" gorm:"column:conversation_session_id;not null"`
	ModelID               int64     `json:"model_id" gorm:"column:model_id;not null"`
	Status                string    `json:"status" gorm:"column:status;type:ENUM('active', 'inactive');default:'active'"`
	CreateTime            time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateTime            time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime;type:datetime(0);"`
}

func (cm *ConversationModel) TableName() string {
	return "conversation_models"
}
