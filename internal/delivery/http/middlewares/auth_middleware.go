package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"pizzatech/config"
	"pizzatech/internal/domain/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(cfg *config.Config, requiredRoles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		role := models.Role(claims["role"].(string))
		ok := false
		for _, r := range requiredRoles {
			if r == role {
				fmt.Println("Found role : ", role, "in ", claims["role"], "")
				ok = true
			}
		}
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Set("userID", uint(claims["sub"].(float64)))
		c.Next()
	}
}
