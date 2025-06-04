package main

import (
	"pizzatech/config"
	"pizzatech/internal/di"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	engine, err := di.Build(cfg)
	if err != nil {
		panic(err)
	}
	engine.Run(":" + cfg.ServerPort)
}
