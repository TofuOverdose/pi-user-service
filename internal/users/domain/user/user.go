package user

import (
	"time"
)

type User struct {
	id    UserId
	props UserProps
}

type UserId struct {
	Value string
}

type UserProps struct {
	Name          string
	LastName      string
	Age           Age
	RecordingDate time.Time
}

/* Getters */

func (user *User) GetId() UserId {
	return user.id
}

func (user *User) GetProps() UserProps {
	return user.props
}

// Age is a value object of User model
type Age struct {
	Value uint8
}

// NewAge creates new instance of age value
func NewAge(value uint8) Age {
	return Age{Value: value}
}

func (a *Age) MoreThan(age Age) bool {
	return a.Value > age.Value
}

func (a *Age) LessThan(age Age) bool {
	return a.Value < age.Value
}

func (a *Age) EqualTo(age Age) bool {
	return a.Value == age.Value
}
