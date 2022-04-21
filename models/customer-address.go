package models

type CustomerAddress struct {
	BaseModel

	Address string `json:"address" gorm:"not null;unique"`

	CustomerId uint     `json:"-" gorm:"not null"`
	Customer   Customer `json:"-"`
}
