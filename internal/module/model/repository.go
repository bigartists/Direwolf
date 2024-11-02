package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type IModelRepo interface {
	FindModelList(userId int64) ([]*Model, error)
	FindModelById(id int64, model *Model) (*Model, error)
	FindModelByModel(model string) (*Model, error)
	CreateModel(model *Model) error
	UpdateModel(id int, model *Model) error
	DeleteModel(id int) error
	CheckDuplicateModelBaseURL(ctx context.Context, model string, baseURL string) (bool, error)
}

type ModelRepo struct {
	db *gorm.DB
}

func ProvideModelRepo(db *gorm.DB) IModelRepo {
	return &ModelRepo{db: db}
}

func (m ModelRepo) FindModelList(userId int64) ([]*Model, error) {
	var models []*Model
	err := m.db.Where("create_by=?", userId).Find(&models).Error
	return models, err
}

func (m ModelRepo) CheckDuplicateModelBaseURL(ctx context.Context, model string, baseURL string) (bool, error) {
	var count int64
	err := m.db.Model(&Model{}).Where("model=? AND base_url=?", model, baseURL).Count(&count).Error
	return count > 0, err
}

func (m ModelRepo) FindModelById(id int64, model *Model) (*Model, error) {
	db := m.db.Where("id=?", id).Find(model)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("model not found, id=%d", id)
	}
	return model, nil
}

func (m ModelRepo) FindModelByModel(model string) (*Model, error) {
	var modelObj Model
	err := m.db.Where("model=?", model).Find(&modelObj).Error
	return &modelObj, err
}

func (m ModelRepo) CreateModel(model *Model) error {
	return m.db.Create(model).Error
}

func (m ModelRepo) UpdateModel(id int, model *Model) error {
	return m.db.Where("id=?", id).Updates(model).Error
}

func (m ModelRepo) DeleteModel(id int) error {
	return m.db.Where("id=?", id).Delete(&Model{}).Error
}
