package middleware

import (
	"github.com/labstack/echo/v4"
)

// Validate connection to DB
func CheckDb(f echo.HandlerFunc) {
	//return func(c echo.Context) error {
	//	if db.CheckConnection() == 0 {
	//		log.Fatal("Error DB")
	//	}
	//	return f(c)
	//}
}
