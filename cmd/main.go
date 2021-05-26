package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"login-mongo-service/auth"
	"login-mongo-service/dbMongo"
	"login-mongo-service/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))

	if dbMongo.CheckConnection() == 0 {
		log.Fatal("Can't connect to DB")
	}

	err := auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Certificates not found: %v", err)
	}

	handler.RouteLogin(e)

	log.Printf("Init server in port: 9191")
	err = e.Start(":9191")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
