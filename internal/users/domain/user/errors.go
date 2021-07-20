package user

import "fmt"

type ModelValidationError struct {
	FieldErrors []FieldValidationError
}

func (e ModelValidationError) Error() string {
	msg := "Validation errors:"
	for _, fe := range e.FieldErrors {
		msg += fmt.Sprintf(" %s (%s),", fe.Field, fe.Message)
	}
	return msg
}

type FieldValidationError struct {
	Field   string
	Message string
}

func (e FieldValidationError) Error() string {
	return fmt.Sprintf("Validation for field %s failed: %s", e.Field, e.Message)
}
