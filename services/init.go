package services

import (
	"sctrans/config"

	"gorm.io/gorm"
)

var CustomerService customerIface
var ProductService productIface
var PaymentMethodService paymentMethodIface
var OrderService orderIface

var Db *gorm.DB

func init() {
	if Db == nil {
		Db = config.ConnectDatabase()
	}

	if CustomerService == nil {
		CustomerService = customer{Db}
	}
	if ProductService == nil {
		ProductService = product{Db}
	}
	if PaymentMethodService == nil {
		PaymentMethodService = paymentMethod{Db}
	}
	if OrderService == nil {
		OrderService = order{Db}
	}

}
