package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"login-mongo-service/internal/model"
	"time"
)

// Method to check if user exist
func UserExist(email string) (model.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := MongoConnect.Database("twittor")
	col := db.Collection("user")

	condition := bson.M{"email": email}

	var result model.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
