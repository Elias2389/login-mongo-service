package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"login-mongo-service/dbMongo"
)

// Validate connection to DB
func CheckDb(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if dbMongo.CheckConnection() == 0 {
			log.Fatal("Error DB")
		}
		return f(c)
	}
}
