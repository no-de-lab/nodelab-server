package model

import (
	"github.com/go-playground/validator"
)

func NewValidator() validator.Validate {
	v := validator.New()

	// Add custom validators
	return *v
}
