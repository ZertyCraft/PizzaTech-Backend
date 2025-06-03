package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "pizzatech/internal/services"
)

type StatsHandler struct {
    stats services.StatisticsService
}

func NewStatsHandler(s services.StatisticsService) *StatsHandler {
    return &StatsHandler{stats: s}
}

func (h *StatsHandler) TotalOrders(c *gin.Context) {
    total, err := h.stats.TotalOrders()
    if err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusOK, gin.H{"total_orders": total})
}
