package persistence

import (
    "fmt"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "pizzatech/config"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
    )
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
