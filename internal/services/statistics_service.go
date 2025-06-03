package services

import (
    "pizzatech/internal/domain/repositories"
)

type StatisticsService interface {
    TotalOrders() (int64, error)
}

type statsService struct {
    orders repositories.OrderRepository
}

func NewStatisticsService(r repositories.OrderRepository) StatisticsService {
    return &statsService{orders: r}
}

func (s *statsService) TotalOrders() (int64, error) {
    all, err := s.orders.GetAll()
    return int64(len(all)), err
}
