package models

import "time"

type Product struct {
	ID          uint                   `json:"id" gorm:"primaryKey"`
	Name        string                 `json:"name"`
	Price       uint                   `json:"price"`
	Description string                 `json:"description"`
	Specs       map[string]interface{} `json:"specs" gorm:"serializer:json"`
	CategoryID  uint                   `json:"category_id"`
	Category    Category               `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
