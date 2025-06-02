package persistence

import (
    "pizzatech/internal/domain/models"
    "pizzatech/internal/domain/repositories"

    "gorm.io/gorm"
)

type gormPizzaRepo struct {
    db *gorm.DB
}

func NewPizzaRepository(db *gorm.DB) repositories.PizzaRepository {
    db.AutoMigrate(&models.Pizza{})
    return &gormPizzaRepo{db}
}

func (r *gormPizzaRepo) Create(p *models.Pizza) error {
    return r.db.Create(p).Error
}

func (r *gormPizzaRepo) GetAll() ([]models.Pizza, error) {
    var list []models.Pizza
    err := r.db.Find(&list).Error
    return list, err
}

func (r *gormPizzaRepo) GetByID(id uint) (*models.Pizza, error) {
    var p models.Pizza
    err := r.db.First(&p, id).Error
    return &p, err
}

func (r *gormPizzaRepo) Update(p *models.Pizza) error {
    return r.db.Save(p).Error
}

func (r *gormPizzaRepo) Delete(id uint) error {
    return r.db.Delete(&models.Pizza{}, id).Error
}
