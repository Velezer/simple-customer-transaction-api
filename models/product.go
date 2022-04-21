package models

type Product struct {
	BaseModel

	Name  string  `json:"name"`
	Price float64 `json:"price" gorm:"type:decimal(16,2)"`
}
