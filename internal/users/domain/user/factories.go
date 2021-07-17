package user

import (
	"time"
)

type UserFactory struct {
	c UserPropsConstraints
}

// NewUserFactory creates user factory with constraints
func NewUserFactory(constraints UserPropsConstraints) UserFactory {
	return UserFactory{
		c: constraints,
	}
}

// NewUser creates new users according to the constraints
func (f UserFactory) NewUser(name string, lastName string, age uint8) (*User, error) {
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
