package usecase

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/internal/model"
)

// User AuthUseCase
type UseCase interface {
	RegisterUser(ctx echo.Context, user *model.User) (*model.User, error)
	Login(ctx echo.Context, user *model.Login) *model.LoginResponse
}
