package models

type Order struct {
	BaseModel

	CustomerAddressID uint `json:"customer_address_id"`
	CustomerAddress   CustomerAddress

	Products       []Product       `json:"products,omitempty" gorm:"many2many:product_order;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	PaymentMethods []PaymentMethod `json:"payment_methods,omitempty" gorm:"many2many:product_payment_methods;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
