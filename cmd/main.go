package main

import (
	"log"
	"login-mongo-service/config"
	mongoDbConnection "login-mongo-service/db"
	"login-mongo-service/internal/server"
	"login-mongo-service/pkg/utils"
	"os"
)

func main() {

	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	// Initialize MongoDb
	mongoDB := mongoDbConnection.NewConnection(cfg.MongoDB.MongoURI)
	mongoClient := mongoDB.ConnectToDB()
	connectMongo := mongoDbConnection.ValidateConnection(mongoClient)
	if connectMongo == true {
		log.Printf("MongoDB connected")
	}

	s := server.NewServer(cfg, mongoClient)
	if err := s.RunServer(); err != nil {
		log.Fatal(err)
	}
}
