package models

import "time"

type Product struct {
	ID          uint                   `json:"id" gorm:"primaryKey"`
	Name        string                 `json:"name"`
	Price       uint                   `json:"price"`
	ImageUrl    string                 `json:"image_url"`
	Description string                 `json:"description"`
	Specs       map[string]interface{} `json:"specs" gorm:"serializer:json"`
	Categories  []Category             `json:"categories" gorm:"many2many:product_categories;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
