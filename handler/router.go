package handler

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/middleware"
)

// RouteLogin .
func RouteLogin(e *echo.Echo) {
	userRouter := e.Group("v1/user")
	h := NewLogin()

	userRouter.Use(middleware.CheckDb)
	userRouter.POST("/login", h.Login)
}
