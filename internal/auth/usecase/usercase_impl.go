package usecase

import (
	"context"
	"github.com/labstack/echo/v4"
	"login-mongo-service/config"
	"login-mongo-service/internal/auth/repository"
	"login-mongo-service/internal/model"
	"login-mongo-service/pkg/logger"
)

// Auth handlers
type authUC struct {
	cfg       *config.Config
	mongoRepo repository.MongoRepository
	logger    logger.Logger
}

// Constructor
func NewAuthUC(cfg *config.Config, mongoRepo repository.MongoRepository, logger logger.Logger) *authUC {
	return &authUC{cfg: cfg, mongoRepo: mongoRepo, logger: logger}
}

func (a authUC) Register(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (a authUC) Login(ctx echo.Context, user *model.Login) *model.LoginResponse {
	userFound := a.mongoRepo.GetUserByEmail(ctx.Request().Context(), user.Email)
	if userFound != nil {
		if err := userFound.ComparePasswords(user.Password); err != nil {
			response := &model.LoginResponse{
				ID:    userFound.ID.Hex(),
				User:  userFound,
				Token: "nil",
			}

			return response
		}
	}
	return nil
}
