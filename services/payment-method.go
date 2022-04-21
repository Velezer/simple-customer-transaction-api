package services

import (
	"sctrans/models"

	"gorm.io/gorm"
)

type paymentMethod struct {
	Db *gorm.DB
}

func (s paymentMethod) Save(m *models.PaymentMethod) (*models.PaymentMethod, error) {
	err := s.Db.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s paymentMethod) FindAll() (ps *[]models.PaymentMethod, err error) {
	err = s.Db.Find(&ps).Error
	if err != nil {
		return nil, err
	}

	return
}
