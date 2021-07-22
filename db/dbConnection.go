package db

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	MongoURI string
}

func NewConnection(mongoUri string) *MongoDb {
	return &MongoDb{MongoURI: mongoUri}
}

// Connect to Mongo DB
func (m MongoDb) ConnectToDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(m.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Connected to DB")
	return client
}

// Validate Connection by PING
func ValidateConnection(client *mongo.Client) bool {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
