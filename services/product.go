package services

import (
	"sctrans/models"

	"gorm.io/gorm"
)

type product struct {
	Db *gorm.DB
}

func (s product) Save(m *models.Product) (*models.Product, error) {
	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s product) FindAll() (ps *[]models.Product, err error) {
	err = s.Db.Find(&ps).Error
	if err != nil {
		return nil, err
	}

	return
}

func (s product) FindById(id uint) (p *models.Product, err error) {
	err = s.Db.First(&p).Error
	if err != nil {
		return nil, err
	}

	return
}
