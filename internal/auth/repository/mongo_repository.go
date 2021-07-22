package repository

import (
	"context"
	"login-mongo-service/internal/model"
)

// Repository
type MongoRepository interface {
	RegisterUser(ctx context.Context, user *model.User) *model.User
	GetUserByEmail(ctx context.Context, email string) *model.User
	GetUserById(ctx context.Context, id string) *model.User
}
