package services

import (
	"sctrans/models"

	"gorm.io/gorm"
)

type order struct {
	Db *gorm.DB
}

func (s order) Save(m *models.Order) (*models.Order, error) {
	err := s.Db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s order) FindAll() (ms *[]models.Order, err error) {
	err = s.Db.Find(&ms).Error
	if err != nil {
		return nil, err
	}

	return
}
