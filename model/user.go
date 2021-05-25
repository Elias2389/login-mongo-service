package model

type Role struct {
	Name string `json:"name"`
}

type Communities []Role

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

type People []User
