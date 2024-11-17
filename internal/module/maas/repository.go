package maas

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type IModelRepo interface {
	FindMaasList(userId int64) ([]*Maas, error)
	FindMaasById(id int64, model *Maas) (*Maas, error)
	FindMaasByModel(model string) (*Maas, error)
	ImportMaas(model *Maas) error
	UpdateMaas(id int, model *Maas) error
	DeleteMaas(id int) error
	CheckDuplicateMaasBaseURL(ctx context.Context, model string, baseURL string) (bool, error)
}

type MaasRepo struct {
	db *gorm.DB
}

func ProvideMaasRepo(db *gorm.DB) IModelRepo {
	return &MaasRepo{db: db}
}

func (m MaasRepo) FindMaasList(userId int64) ([]*Maas, error) {
	var maas []*Maas
	err := m.db.Where("create_by=?", userId).Find(&maas).Error
	return maas, err
}

func (m MaasRepo) CheckDuplicateMaasBaseURL(ctx context.Context, model string, baseURL string) (bool, error) {
	var count int64
	err := m.db.Model(&Maas{}).Where("maas=? AND base_url=?", model, baseURL).Count(&count).Error
	return count > 0, err
}

func (m MaasRepo) FindMaasById(id int64, model *Maas) (*Maas, error) {
	db := m.db.Where("id=?", id).Find(model)
	if db.Error != nil || db.RowsAffected == 0 {
		return nil, fmt.Errorf("maas not found, id=%d", id)
	}
	return model, nil
}

func (m MaasRepo) FindMaasByModel(model string) (*Maas, error) {
	var modelObj Maas
	err := m.db.Where("maas=?", model).Find(&modelObj).Error
	return &modelObj, err
}

func (m MaasRepo) ImportMaas(model *Maas) error {
	return m.db.Create(model).Error
}

func (m MaasRepo) UpdateMaas(id int, model *Maas) error {
	return m.db.Where("id=?", id).Updates(model).Error
}

func (m MaasRepo) DeleteMaas(id int) error {
	return m.db.Where("id=?", id).Delete(&Maas{}).Error
}
