package handler

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/config"
	"login-mongo-service/internal/auth/usecase"
	"login-mongo-service/internal/handler"
	"login-mongo-service/internal/model"
	"login-mongo-service/pkg/logger"
	"net/http"
)

// Auth handlers
type authHandler struct {
	cfg    *config.Config
	authUC usecase.UseCase
	logger logger.Logger
}

func NewAuthHandler(cfg *config.Config, authUC usecase.UseCase, logger logger.Logger) *authHandler {
	return &authHandler{cfg: cfg, authUC: authUC, logger: logger}
}

// Handler Register
func (h *authHandler) Register() echo.HandlerFunc {
	return nil
}

// Handler Login
func (h *authHandler) Login() echo.HandlerFunc {
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
