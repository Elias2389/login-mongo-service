package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"login-mongo-service/internal/model"
	"time"
)

// Method to check if user exist
func UserExist(email string, mongoClient *mongo.Client) (model.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := mongoClient.Database("twittor")
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
