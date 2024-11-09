package maas

import (
	"time"
)

type MaasStatus string

const (
	MaasStatusActive     MaasStatus = "active"
	MaasStatusInactive   MaasStatus = "inactive"
	MaasStatusDeprecated MaasStatus = "deprecated"
)

// User 创建 Users struct
type Maas struct {
	ID          int64      `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	Status      MaasStatus `json:"status" gorm:"type:ENUM('active', 'inactive', 'deprecated');default:'active'"`
	Model       string     `json:"model" gorm:"column:maas;unique" binding:"required"`
	BaseUrl     string     `json:"base_url" gorm:"column:base_url;not null" binding:"required"`
	ApiKey      string     `json:"api_key" gorm:"column:api_key;not null" binding:"required"`
	ModelType   string     `json:"model_type" gorm:"column:model_type;not null" binding:"required"`
	Avatar      string     `json:"avatar" gorm:"column:avatar"`
	// 自动维护时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime;type:datetime(0);"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoCreateTime;<-:false;type:datetime(0);"`
	CreateBy   int64     `json:"create_by" gorm:"not null;default:0"`
	UpdateBy   int64     `json:"-" gorm:"column:update_by"`
}

func (u *Maas) TableName() string {
	return "model"
}

func NewMaas(attrs ...MaasAttrFunc) *Maas {
	u := &Maas{}
	MaasAttrFuncs(attrs).apply(u)
	return u
}

func (u *Maas) Mutate(attrs ...MaasAttrFunc) *Maas {
	MaasAttrFuncs(attrs).apply(u)
	return u
}
