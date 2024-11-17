package conversation_maas

import "time"

type ConversationMaas struct {
	ID                    int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ConversationSessionID string    `json:"conversation_session_id" gorm:"column:conversation_session_id;not null"`
	MaasID                int64     `json:"maas_id" gorm:"column:maas_id;not null"`
	Status                string    `json:"status" gorm:"column:status;type:ENUM('active', 'inactive');default:'active'"`
	CreateTime            time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateTime            time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime;type:datetime(0);"`
}

func (cm *ConversationMaas) TableName() string {
	return "conversation_maas"
}
