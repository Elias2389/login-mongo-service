package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

// Role Model
type Role struct {
	Name string `json:"name"`
}

// User Model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Lastname  string             `bson:"lastname" json:"lastname,omitempty"`
	Date      time.Time          `bson:"time" json:"time,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}

// Find user query
type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// Generate Hash from password
func (u *User) GenerateHashPassword() error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashPass)
	return nil
}

// Compare Passwords
func (u *User) ComparePasswords(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}

// Setup info about user to register
func (u *User) SetupToRegister() error {
	u.Password = strings.TrimSpace(u.Password)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	if err := u.GenerateHashPassword(); err != nil {
		return err
	}
	return nil
}
