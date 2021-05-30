package http

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/internal/auth/delivery"
)

// Map handlers about User Auth
func MapAuthRoutes(authGroup *echo.Group, h delivery.Handlers) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
}
