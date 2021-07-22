package server

import (
	"github.com/labstack/echo/v4"
	authHttp "login-mongo-service/internal/auth/delivery/http"
	authHandler "login-mongo-service/internal/auth/delivery/http/handler"
	authRepository "login-mongo-service/internal/auth/repository"
	authUseCase "login-mongo-service/internal/auth/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Setup routes
	//authRouter := s.echo.Group("v1/user")

	v1 := e.Group("/api/v1")

	userGroup := v1.Group("/user")
	mongoRepo := authRepository.NewMongoRepository(s.mongo)
	authUC := authUseCase.NewAuthUC(s.cfg, mongoRepo)
	authH := authHandler.NewAuthHandler(s.cfg, authUC)
	authHttp.MapAuthRoutes(userGroup, authH)

	return nil
}
