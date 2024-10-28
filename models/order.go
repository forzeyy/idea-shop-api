package models

import "time"

type Order struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	UserID    uint        `json:"user_id"`
	User      User        `gorm:"foreignKey:UserID"`
	Items     []OrderItem `json:"order_items"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `gorm:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
