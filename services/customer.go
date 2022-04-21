package services

import (
	"sctrans/models"

	"gorm.io/gorm"
)

type customer struct {
	Db *gorm.DB
}

func (s customer) Save(u *models.Customer) (*models.Customer, error) {
	err := s.Db.Create(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s customer) FindByName(n string) (u *models.Customer, err error) {
	err = s.Db.First(&u, &models.Customer{CustomerName: n}).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s customer) AddAddress(id uint, a string) error {
	return s.Db.Create(&models.CustomerAddress{Address: a, CustomerId: id}).Error
}
