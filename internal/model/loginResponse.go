package model

// Response of login
type LoginResponse struct {
	ID    string `json:"id"`
	User  *User  `json:"user"`
	Token string `json:"token"`
}
