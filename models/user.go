package models

import (
	"github.com/go-playground/validator/v10"
)

// UserLoginForm is the input from user on login.
type UserLoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7,max=255"`
}

// User is a direct representation of the mongo document.
type User struct {
	ID       string   `json:"id" bson:"_id"`
	Email    string   `json:"email" bson:"email" validate:"required,email"`
	Username string   `json:"username" bson:"username" validate:"required,min=3,max=100"`
	Password string   `json:"password" bson:"password" validate:"required,min=7,max=255"`
	Roles    []string `json:"roles" bson:"roles"`
}

// IsValid checks if instance of User is valid using the validator.
func (u User) IsValid() error {
	validate := validator.New()
	return validate.Struct(u)
}
