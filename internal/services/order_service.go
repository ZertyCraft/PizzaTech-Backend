package services

import (
    "pizzatech/internal/domain/models"
    "pizzatech/internal/domain/repositories"
)

type OrderService interface {
    Create(o *models.Order) error
    ListByUser(userID uint) ([]models.Order, error)
    UpdateStatus(orderID uint, status models.OrderStatus) error
}

type orderService struct {
    repo repositories.OrderRepository
}

func NewOrderService(r repositories.OrderRepository) OrderService {
    return &orderService{repo: r}
}

func (s *orderService) Create(o *models.Order) error {
    return s.repo.Create(o)
}

func (s *orderService) ListByUser(userID uint) ([]models.Order, error) {
    return s.repo.GetByUser(userID)
}

func (s *orderService) UpdateStatus(orderID uint, status models.OrderStatus) error {
    return s.repo.UpdateStatus(orderID, status)
}
