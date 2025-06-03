package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "pizzatech/internal/services"
)

type ProfileHandler struct {
    orders services.OrderService
}

func NewProfileHandler(o services.OrderService) *ProfileHandler {
    return &ProfileHandler{orders: o}
}

func (h *ProfileHandler) History(c *gin.Context) {
    uid := c.GetUint("userID")
    list, err := h.orders.ListByUser(uid)
    if err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusOK, list)
}
