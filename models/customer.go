package models

type Customer struct {
	BaseModel

	CustomerName string `json:"customer_name" gorm:"unique;not null;type:char(50);check:customer_name <> ''"`

	CustomerAddresses []CustomerAddress `json:"customer_addresses"`
}
