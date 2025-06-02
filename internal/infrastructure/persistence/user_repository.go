package persistence

import (
    "pizzatech/internal/domain/models"
    "pizzatech/internal/domain/repositories"

    "gorm.io/gorm"
)

type gormUserRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
    db.AutoMigrate(&models.User{})
    return &gormUserRepo{db}
}

func (r *gormUserRepo) Create(u *models.User) error {
    return r.db.Create(u).Error
}

func (r *gormUserRepo) FindByEmail(email string) (*models.User, error) {
    var u models.User
    err := r.db.Where("email = ?", email).First(&u).Error
    return &u, err
}

func (r *gormUserRepo) FindByID(id uint) (*models.User, error) {
    var u models.User
    err := r.db.First(&u, id).Error
    return &u, err
}
