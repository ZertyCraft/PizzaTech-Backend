package models

import "gorm.io/gorm"

type Role string

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
	RoleWorker   Role = "worker"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     Role   `gorm:"not null"`
}
