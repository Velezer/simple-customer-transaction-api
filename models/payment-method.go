package models

type PaymentMethod struct {
	BaseModel

	Name     string `json:"name"`
	IsActive int    `json:"is_active" gorm:"tinyint(1)"`
}
