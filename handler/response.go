package handler

import "login-mongo-service/model"

const (
	Error   = "error"
	Message = "message"
)

func NewResponse(messageType string, message string, data interface{}) model.Response {
	return model.Response{
		MessageType: messageType,
		Message:     message,
		Data:        data,
	}
}
