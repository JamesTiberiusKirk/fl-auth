package models

import "github.com/go-playground/validator/v10"

type LoginResponseDto struct {
	Message string `json:"message"`
	Jwt     string `json:"jwt"`
}

type JwtDto struct {
	Jwt string `json:"jwt" validate:"required"`
}

func (j JwtDto) IsValid() error {
	validate := validator.New()
	return validate.Struct(j)
}
