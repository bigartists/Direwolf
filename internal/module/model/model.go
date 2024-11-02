package model

import (
	"time"
)

type ModelStatus string

const (
	ModelStatusActive     ModelStatus = "active"
	ModelStatusInactive   ModelStatus = "inactive"
	ModelStatusDeprecated ModelStatus = "deprecated"
)

// User 创建 Users struct
type Model struct {
	ID          int64       `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Name        string      `json:"name" gorm:"not null"`
	Description string      `json:"description"`
	Status      ModelStatus `json:"status" gorm:"type:ENUM('active', 'inactive', 'deprecated');default:'active'"`
	Model       string      `json:"model" gorm:"column:model;unique" binding:"required"`
	BaseUrl     string      `json:"base_url" gorm:"column:base_url;not null" binding:"required"`
	ApiKey      string      `json:"api_key" gorm:"column:api_key;not null" binding:"required"`
	ModelType   string      `json:"model_type" gorm:"column:model_type;not null" binding:"required"`
	Avatar      string      `json:"avatar" gorm:"column:avatar"`
	// 自动维护时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
	CreateBy   int64     `json:"create_by" gorm:"not null;default:0"`
	UpdateBy   int64     `json:"-" gorm:"column:update_by"`
}

func (u *Model) TableName() string {
	return "model"
}

func NewModel(attrs ...ModelAttrFunc) *Model {
	u := &Model{}
	ModelAttrFuncs(attrs).apply(u)
	return u
}

func (u *Model) Mutate(attrs ...ModelAttrFunc) *Model {
	ModelAttrFuncs(attrs).apply(u)
	return u
}
