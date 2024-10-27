package models

import "time"

type Cart struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"foreignKey:UserID"`
	Items     []CartItem `json:"cart_items"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
