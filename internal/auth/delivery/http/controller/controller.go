package controller

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/config"
	"login-mongo-service/internal/auth/usecase"
	"login-mongo-service/internal/handler"
	"login-mongo-service/internal/model"
	"net/http"
)

// Auth handlers
type authController struct {
	cfg    *config.Config
	authUC usecase.UseCase
}

func NewAuthHandler(cfg *config.Config, authUC usecase.UseCase) *authController {
	return &authController{cfg: cfg, authUC: authUC}
}

// Handler Register
func (h *authController) RegisterUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := &model.User{}

		if err := ctx.Bind(user); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		createUser, err := h.authUC.RegisterUser(ctx, user)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		return ctx.JSON(http.StatusOK, createUser)
	}
}

// Handler Login
func (h *authController) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestData := model.Login{}
		if err := ctx.Bind(&requestData); err != nil {
			response := handler.NewResponse(handler.Error, "Invalid Data", nil)
			return ctx.JSON(http.StatusBadRequest, response)
		}

		user := h.authUC.Login(ctx, &requestData)
		if user == nil {
			resp := handler.NewResponse(handler.Error, "Error", nil)
			return ctx.JSON(http.StatusBadRequest, resp)
		}

		response := handler.NewResponse(handler.Message, "Success", user)
		return ctx.JSON(http.StatusOK, response)
	}
}
