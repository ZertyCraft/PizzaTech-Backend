package models

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPreparing OrderStatus = "preparing"
	StatusReady     OrderStatus = "ready"
	StatusDone      OrderStatus = "done"
)

type Order struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	User   User
	Status OrderStatus `gorm:"not null;default:'pending'"`
	Items  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID  uint `gorm:"not null"`
	PizzaID  uint `gorm:"not null"`
	Pizza    Pizza
	Quantity int `gorm:"not null"`
}
