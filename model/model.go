package model

import "errors"

var (
	// ErrPersonCanNotBeNil
	ErrPersonCanNotBeNil = errors.New("user can't be null")
	// ErrIDPersonDoesNotExists
	ErrIDPersonDoesNotExists = errors.New("user not found")
)
