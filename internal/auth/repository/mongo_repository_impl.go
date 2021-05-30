package repository

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"login-mongo-service/internal/model"
	"time"
)

const (
	dbName         = "twittor"
	collectionName = "user"
)

type MongoRepoImpl struct {
	client *mongo.Client
}

func NewMongoRepository(client *mongo.Client) *MongoRepoImpl {
	return &MongoRepoImpl{client: client}
}

func (r MongoRepoImpl) GetUserByEmail(ctx context.Context, email string) *model.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := provideDB(r.client)
	col := provideCollection(db)

	filter := bson.M{"email": email}

	var response model.User
	err := col.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &response
}

func provideDB(client *mongo.Client) *mongo.Database {
	return client.Database(dbName)
}

func provideCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection(collectionName)
}
