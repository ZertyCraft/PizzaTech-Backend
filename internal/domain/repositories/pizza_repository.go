package repositories

import "pizzatech/internal/domain/models"

type PizzaRepository interface {
	Create(p *models.Pizza) error
	GetAll() ([]models.Pizza, error)
	GetByID(id uint) (*models.Pizza, error)
	Update(p *models.Pizza) error
	Delete(id uint) error
}
