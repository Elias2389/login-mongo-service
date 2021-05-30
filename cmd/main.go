package main

import (
	"log"
	"login-mongo-service/config"
	"login-mongo-service/db"
	"login-mongo-service/internal/server"
	"login-mongo-service/pkg/logger"
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

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	mongoClient := db.ConnectDB()
	connectMongo := db.ValidateConnection(mongoClient)
	if connectMongo == true {
		appLogger.Info("MongoDB connected")
	}

	s := server.NewServer(cfg, mongoClient, appLogger)
	if err := s.RunServer(); err != nil {
		log.Fatal(err)
	}
}
