package models

import "github.com/shopspring/decimal"

type Product struct {
	BaseModel

	Name  int             `json:"name"`
	Price decimal.Decimal `json:"price" gorm:"type:decimal"`
}
