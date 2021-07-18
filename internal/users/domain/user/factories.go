package user

import (
	"time"
)

type UserFactory struct {
	c UserPropsConstraints
}

// NewUserFactory creates user factory with constraints
func NewUserFactory(constraints UserPropsConstraints) *UserFactory {
	return &UserFactory{
		c: constraints,
	}
}

// NewUser creates new users according to the constraints
func (f *UserFactory) NewUser(name string, lastName string, age uint8) (*User, error) {
	props := UserProps{
		Name:          name,
		LastName:      lastName,
		Age:           NewAge(age),
		RecordingDate: time.Now(),
	}
	if err := ValidateUserProps(props, f.c); err != nil {
		return nil, err
	}

	return &User{props: props}, nil
}

/*
	BuildUser recreates user model with given Id and Props without validation
	This function should only be used by repositories to transfer their internal representation to domain model
	Application-level code must use factory to create user
*/
func BuildUser(id UserId, props UserProps) *User {
	return &User{
		id:    id,
		props: props,
	}
}
