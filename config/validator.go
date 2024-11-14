package config

import "github.com/go-playground/validator/v10"

func InitValidator() *validator.Validate {
	validator := validator.New()
	return validator
}
