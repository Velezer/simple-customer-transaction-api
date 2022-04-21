package services

import (
	"sctrans/models"

	"gorm.io/gorm"
)

type order struct {
	Db *gorm.DB
}

func (s order) Save(m *models.Order) (*models.Order, error) {
	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s order) FindById(id uint) (m *models.Order, err error) {
	err = s.Db.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return
}
