package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "pizzatech/internal/domain/models"
    "pizzatech/internal/services"
)

type PizzaHandler struct {
    srv services.PizzaService
}

func NewPizzaHandler(s services.PizzaService) *PizzaHandler {
    return &PizzaHandler{srv: s}
}

func (h *PizzaHandler) Create(c *gin.Context) {
    var p models.Pizza
    if c.BindJSON(&p) != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    if err := h.srv.Create(&p); err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusCreated)
}

func (h *PizzaHandler) List(c *gin.Context) {
    list, err := h.srv.List()
    if err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusOK, list)
}

func (h *PizzaHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    p, err := h.srv.Get(uint(id))
    if err != nil {
        c.Status(http.StatusNotFound)
        return
    }
    c.JSON(http.StatusOK, p)
}

func (h *PizzaHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var p models.Pizza
    if c.BindJSON(&p) != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    if err := h.srv.Update(uint(id), &p); err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusOK)
}

func (h *PizzaHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.srv.Delete(uint(id)); err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusNoContent)
}
