package delivery

import "github.com/labstack/echo/v4"

type Handlers interface {
	RegisterUser() echo.HandlerFunc
	Login() echo.HandlerFunc
}
