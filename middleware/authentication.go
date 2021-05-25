package middleware

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/auth"
	"net/http"
)

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := auth.ValidateToken(token)
		if err != nil {
			return forbidden(c)
		}

		return f(c)
	}
}

func forbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden, getError())
}

func getError() map[string]string {
	return map[string]string{
		"error": "Don't have auth",
	}
}
