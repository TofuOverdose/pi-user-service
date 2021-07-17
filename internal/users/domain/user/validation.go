package user

import "unicode/utf8"

type UserPropsConstraints struct {
	NameMinLen     uint8
	NameMaxLen     uint8
	LastNameMinLen uint8
	LastNameMaxLen uint8
	MinAge         Age
}

// ValidateUserProps checks user props for meething the constraints
func ValidateUserProps(props UserProps, constraints UserPropsConstraints) *ModelValidationError {
	errors := make([]FieldValidationError, 0)

	if utf8.RuneCountInString(props.Name) < int(constraints.NameMinLen) {
		errors = append(errors, FieldValidationError{Field: "Name", Message: "Name is too short"})
	}

	if utf8.RuneCountInString(props.Name) > int(constraints.NameMaxLen) {
		errors = append(errors, FieldValidationError{Field: "Name", Message: "Name is too long"})
	}

	if utf8.RuneCountInString(props.LastName) < int(constraints.LastNameMinLen) {
		errors = append(errors, FieldValidationError{Field: "LastName", Message: "LastName is too short"})
	}

	if utf8.RuneCountInString(props.LastName) > int(constraints.LastNameMaxLen) {
		errors = append(errors, FieldValidationError{Field: "LastName", Message: "LastName is too long"})
	}

	if props.Age.LessThan(constraints.MinAge) {
		errors = append(errors, FieldValidationError{Field: "Age", Message: "User is too young"})
	}

	if len(errors) == 0 {
		return nil
	}

	return &ModelValidationError{FieldErrors: errors}
}
