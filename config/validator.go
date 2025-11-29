package config

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

// InitValidator initializes validator
func InitValidator() *validator.Validate {
	validate := validator.New()

	// Register custom validations here if needed
	// validate.RegisterValidation("custom_tag", customValidationFunc)

	Validator = validate
	return validate
}
