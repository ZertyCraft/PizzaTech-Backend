package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "pizzatech/internal/domain/models"
    "pizzatech/internal/services"
)

type AuthHandler struct {
    auth services.AuthService
}

func NewAuthHandler(a services.AuthService) *AuthHandler {
    return &AuthHandler{auth: a}
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req struct {
        Email    string      `json:"email"`
        Password string      `json:"password"`
        Role     models.Role `json:"role"`
    }
    if c.BindJSON(&req) != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    if err := h.auth.Register(req.Email, req.Password, req.Role); err != nil {
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusCreated)
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if c.BindJSON(&req) != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    token, err := h.auth.Login(req.Email, req.Password)
    if err != nil {
        c.Status(http.StatusUnauthorized)
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
