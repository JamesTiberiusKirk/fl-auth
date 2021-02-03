package models

import "github.com/go-playground/validator/v10"

// LoginResponseDto is the struct that is getting retured to the user at login
type LoginResponseDto struct {
	Message string `json:"message"`
	Jwt     string `json:"jwt"`
}

// JwtDto is a DTO for the jwt alone
type JwtDto struct {
	Jwt string `json:"jwt" validate:"required"`
}

// IsValid validates an instance of JwtDto using the validator
func (j JwtDto) IsValid() error {
	validate := validator.New()
	return validate.Struct(j)
}
