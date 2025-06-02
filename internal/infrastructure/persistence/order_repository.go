package persistence

import (
    "pizzatech/internal/domain/models"
    "pizzatech/internal/domain/repositories"

    "gorm.io/gorm"
)

type gormOrderRepo struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
    db.AutoMigrate(&models.Order{}, &models.OrderItem{})
    return &gormOrderRepo{db}
}

func (r *gormOrderRepo) Create(o *models.Order) error {
    return r.db.Create(o).Error
}

func (r *gormOrderRepo) GetByUser(userID uint) ([]models.Order, error) {
    var orders []models.Order
    err := r.db.Preload("Items.Pizza").Where("user_id = ?", userID).Find(&orders).Error
    return orders, err
}

func (r *gormOrderRepo) UpdateStatus(orderID uint, status models.OrderStatus) error {
    return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *gormOrderRepo) GetAll() ([]models.Order, error) {
    var orders []models.Order
    err := r.db.Preload("Items.Pizza").Find(&orders).Error
    return orders, err
}
