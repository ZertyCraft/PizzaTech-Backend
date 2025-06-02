package repositories

import (
    "pizzatech/internal/domain/models"
)

type OrderRepository interface {
    Create(o *models.Order) error
    GetByUser(userID uint) ([]models.Order, error)
    UpdateStatus(orderID uint, status models.OrderStatus) error
    GetAll() ([]models.Order, error)
}
