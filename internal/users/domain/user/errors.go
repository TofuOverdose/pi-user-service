package user

import "fmt"

type ModelValidationError struct {
	FieldErrors []FieldValidationError
}

func (e ModelValidationError) Error() string {
	return "Validation for model failed"
}

type FieldValidationError struct {
	Field   string
	Message string
}

func (e FieldValidationError) Error() string {
	return fmt.Sprintf("Validation for field %s failed: %s", e.Field, e.Message)
}
