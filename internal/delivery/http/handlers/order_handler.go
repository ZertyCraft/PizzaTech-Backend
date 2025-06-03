package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "pizzatech/internal/domain/models"
    "pizzatech/internal/services"
)

type OrderHandler struct {
    srv services.OrderService
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
    return &OrderHandler{srv: s}
}

func (h *OrderHandler) Create(c *gin.Context) {
    var req struct {
        Items []struct {
            PizzaID  uint `json:"pizza_id"`
            Quantity int  `json:"quantity"`
        } `json:"items"`
    }
    if c.BindJSON(&req) != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    uid := c.GetUint("userID")
    order := &models.Order{UserID: uid}
    for _, it := range req.Items {
        order.Items = append(order.Items, models.OrderItem{PizzaID: it.PizzaID, Quantity: it.Quantity})
    }
    if err := h.srv.Create(order); err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusCreated)
}

func (h *OrderHandler) List(c *gin.Context) {
    uid := c.GetUint("userID")
    orders, err := h.srv.ListByUser(uid)
    if err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var req struct {
        Status models.OrderStatus `json:"status"`
    }
    if c.BindJSON(&req) != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    if err := h.srv.UpdateStatus(uint(id), req.Status); err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusOK)
}
