package models

import "time"

type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	AdminName string `json:"admin_name" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
