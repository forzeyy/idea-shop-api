package models

import "time"

type Comment struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	Text      string `json:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
