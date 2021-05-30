package db

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userMongoDb string = "twittor_user"
const passMongoDb string = "TestUser123"

var MongoConnect = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://" + userMongoDb + ":" + passMongoDb + "@twittor.blv2u.mongodb.net/myFirstDatabase")

func ConnectDB() *mongo.Client {
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

func CheckConnection() int {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func ValidateConnection(client *mongo.Client) bool {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
