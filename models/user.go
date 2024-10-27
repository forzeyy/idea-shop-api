package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Phone     string `json:"phone" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
