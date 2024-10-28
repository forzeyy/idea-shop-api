package models

import "time"

type Comment struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:UserID"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Text      string  `json:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
