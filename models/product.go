package models

import "time"

type Product struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	Name          string            `json:"name"`
	Price         uint              `json:"price"`
	Description   string            `json:"description"`
	Specs         map[string]string `json:"specs"`
	CategoryRefer uint              `json:"category_id"`
	Category      Category          `gorm:"foreignKey:CategoryRefer"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
