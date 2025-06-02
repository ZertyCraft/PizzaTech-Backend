package services

import (
    "pizzatech/internal/domain/models"
    "pizzatech/internal/domain/repositories"
)

type PizzaService interface {
    Create(p *models.Pizza) error
    List() ([]models.Pizza, error)
    Get(id uint) (*models.Pizza, error)
    Update(id uint, p *models.Pizza) error
    Delete(id uint) error
}

type pizzaService struct {
    repo repositories.PizzaRepository
}

func NewPizzaService(r repositories.PizzaRepository) PizzaService {
    return &pizzaService{repo: r}
}

func (s *pizzaService) Create(p *models.Pizza) error {
    return s.repo.Create(p)
}

func (s *pizzaService) List() ([]models.Pizza, error) {
    return s.repo.GetAll()
}

func (s *pizzaService) Get(id uint) (*models.Pizza, error) {
    return s.repo.GetByID(id)
}

func (s *pizzaService) Update(id uint, p *models.Pizza) error {
    p.ID = id
    return s.repo.Update(p)
}

func (s *pizzaService) Delete(id uint) error {
    return s.repo.Delete(id)
}
