package model

type Response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}
