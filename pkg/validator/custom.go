package validator

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator implements echo.Validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the request data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewCustomValidator creates a new instance of CustomValidator
func NewCustomValidator() *CustomValidator {
	v := validator.New()

	// Add any custom validation functions here if needed
	// Example: v.RegisterValidation("custom_tag", customValidationFunction)

	return &CustomValidator{validator: v}
}
