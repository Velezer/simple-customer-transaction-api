package models

type Customer struct {
	BaseModel

	CustomerName string `json:"customer_name" gorm:"unique;not null;type:char(50);check:username <> ''"`

	CustomerAddresses []CustomerAddress
}
