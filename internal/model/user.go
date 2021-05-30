package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Role struct {
	Name string `json:"name"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Lastname string             `bson:"lastname" json:"lastname,omitempty"`
	Date     time.Time          `bson:"time" json:"time,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Website  string             `bson:"website" json:"website,omitempty"`
}

// Compare Passwords
func (u User) ComparePasswords(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}
