package models

import (
	"github.com/go-playground/validator/v10"
)

type UserLoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7,max=255"`
}

type User struct {
	Email    string   `json:"email" validate:"required,email"`
	Username string   `json:"username" validate:"required,min=3,max=100"`
	Password string   `json:"password" validate:"required,min=7,max=255"`
	Roles    []string `json:"roles"`
}

func (u User) IsValid() error {
	validate := validator.New()
	return validate.Struct(u)
}
