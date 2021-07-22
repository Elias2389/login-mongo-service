package usecase

import (
	"github.com/labstack/echo/v4"
	"login-mongo-service/config"
	"login-mongo-service/internal/auth/repository"
	"login-mongo-service/internal/model"
)

// Auth handlers
type authUC struct {
	cfg       *config.Config
	mongoRepo repository.MongoRepository
}

// Constructor
func NewAuthUC(cfg *config.Config, mongoRepo repository.MongoRepository) *authUC {
	return &authUC{cfg: cfg, mongoRepo: mongoRepo}
}

// RegisterUser
func (a authUC) RegisterUser(ctx echo.Context, user *model.User) (*model.User, error) {
	existUser := a.mongoRepo.GetUserByEmail(ctx.Request().Context(), user.Email)
	if existUser != nil {
		return nil, nil
	}

	createUser := a.mongoRepo.RegisterUser(ctx.Request().Context(), user)
	if createUser != nil {
		response := &model.User{
			ID:       createUser.ID,
			Name:     createUser.Name,
			Lastname: createUser.Lastname,
			Email:    createUser.Email,
		}
		return response, nil
	}

	return nil, nil
}

// Login user
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
