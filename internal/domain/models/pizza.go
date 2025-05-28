package models

import "gorm.io/gorm"

type Pizza struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
}
