package di

import (
	"pizzatech/config"
	"pizzatech/internal/delivery/http/handlers"
	"pizzatech/internal/delivery/http/middlewares"
	"pizzatech/internal/domain/models"
	"pizzatech/internal/infrastructure/logger"
	"pizzatech/internal/infrastructure/persistence"
	"pizzatech/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func Build(cfg *config.Config) (*gin.Engine, error) {
	container := dig.New()

	container.Provide(func() *config.Config { return cfg })
	container.Provide(logger.New)
	container.Provide(persistence.NewDB)
	container.Provide(persistence.NewUserRepository)
	container.Provide(persistence.NewPizzaRepository)
	container.Provide(persistence.NewOrderRepository)
	container.Provide(services.NewAuthService)
	container.Provide(services.NewPizzaService)
	container.Provide(services.NewOrderService)
	container.Provide(services.NewStatisticsService)
	container.Provide(handlers.NewAuthHandler)
	container.Provide(handlers.NewPizzaHandler)
	container.Provide(handlers.NewOrderHandler)
	container.Provide(handlers.NewProfileHandler)
	container.Provide(handlers.NewStatsHandler)

	var engine *gin.Engine
	err := container.Invoke(func(
		cfg *config.Config,
		log *logrus.Logger,
		ah *handlers.AuthHandler,
		ph *handlers.PizzaHandler,
		oh *handlers.OrderHandler,
		prh *handlers.ProfileHandler,
		sh *handlers.StatsHandler,
	) {
		engine = gin.New()
		engine.Use(middlewares.Logger(log))
		api := engine.Group("/api")
		api.POST("/register", ah.Register)
		api.POST("/login", ah.Login)

		authAll := api.Group("", middlewares.Auth(cfg, models.RoleCustomer, models.RoleAdmin, models.RoleWorker))
		authAll.GET("/pizzas", ph.List)
		authAll.GET("/pizzas/:id", ph.Get)

		authAll.POST("/orders", oh.Create)
		authAll.GET("/orders", oh.List)

		workerOrAdmin := api.Group("", middlewares.Auth(cfg, models.RoleWorker, models.RoleAdmin))
		workerOrAdmin.POST("/pizzas", ph.Create)
		workerOrAdmin.PUT("/pizzas/:id", ph.Update)
		workerOrAdmin.DELETE("/pizzas/:id", ph.Delete)

		workerOrAdmin.PATCH("/orders/:id/status", oh.UpdateStatus)

		adminOnly := api.Group("", middlewares.Auth(cfg, models.RoleAdmin))
		adminOnly.GET("/stats", sh.TotalOrders)

		customerOnly := api.Group("", middlewares.Auth(cfg, models.RoleCustomer))
		customerOnly.GET("/profile/history", prh.History)
	})

	return engine, err
}
