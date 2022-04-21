package services

import "sctrans/models"

type customerIface interface {
	Save(u *models.Customer) (*models.Customer, error)
	AddAddress(id uint, a string) error
	FindByName(n string) (u *models.Customer, err error)
}

type productIface interface {
	FindAll() (ps *[]models.Product, err error)
	Save(s *models.Product) (*models.Product, error)
}
type paymentMethodIface interface {
	FindAll() (ps *[]models.PaymentMethod, err error)
	Save(s *models.PaymentMethod) (*models.PaymentMethod, error)
}

type orderIface interface {
	FindAll() (*[]models.Order, error)
	Save(m *models.Order) (*models.Order, error)
}
