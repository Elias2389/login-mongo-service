package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"login-mongo-service/db"
	"login-mongo-service/internal/auth"
	"login-mongo-service/internal/model"
	"net/http"
)

type login struct {
	MongoClient *mongo.Client
}

// Initialize login
func NewLogin(client *mongo.Client) *login {
	return &login{MongoClient: client}
}

// Method to do login
func (l *login) Login(c echo.Context) error {
	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, "estructura no válida", nil)
		return c.JSON(http.StatusOK, resp)
	}

	user, userExist, id := db.UserExist(data.Email, l.MongoClient)

	if !userExist {
		resp := NewResponse(Error, "Ocurrio un error", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := auth.GenerateToken(&data)
	if err != nil {
		resp := NewResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	response := model.LoginResponse{
		ID:    id,
		User:  &user,
		Token: token,
	}

	resp := NewResponse(Message, "Ok", response)
	return c.JSON(http.StatusOK, resp)
}
