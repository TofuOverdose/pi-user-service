package user

import (
	"fmt"
	"strings"
)

type ModelValidationError struct {
	FieldErrors []FieldValidationError
}

func (e ModelValidationError) Error() string {
	errs := make([]string, 0)
	for _, fe := range e.FieldErrors {
		errs = append(errs, fmt.Sprintf("%s (%s)", fe.Field, fe.Message))
	}
	return fmt.Sprintf("Validation errors: %s", strings.Join(errs, ", "))
}

type FieldValidationError struct {
	Field   string
	Message string
}

func (e FieldValidationError) Error() string {
	return fmt.Sprintf("Validation for field %s failed: %s", e.Field, e.Message)
}
