package repository

import (
	"context"
	"login-mongo-service/internal/model"
)

// Repository
type MongoRepository interface {
	GetUserByEmail(ctx context.Context, email string) *model.User
}
