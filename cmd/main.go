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
	err := auth.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Certificates not found: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	if dbMongo.CheckConnection() == 0 {
		log.Fatal("Can't connect to DB")
	}

	handler.RouteLogin(e)

	log.Printf("Init server in port: 9191")
	err = e.Start(":9191")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
