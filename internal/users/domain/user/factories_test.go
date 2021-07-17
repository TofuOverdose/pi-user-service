package user_test

import (
	"testing"
	"time"

	"example.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	t.Log("Creating user successfully within constraints")
	factory := user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     4,
		NameMaxLen:     5,
		LastNameMinLen: 6,
		LastNameMaxLen: 7,
		MinAge:         user.NewAge(16),
	})

	var (
		name           = "John"
		lastName       = "Barber"
		age      uint8 = 16
	)
	user, err := factory.NewUser(name, lastName, age)
	t.Log(user, err)
	assert.Nil(t, err, "Should create user without error")
	assert.NotNil(t, user, "User model should be returned")

	props := user.GetProps()
	assert.Equal(t, name, props.Name)
	assert.Equal(t, lastName, props.LastName)
	assert.Equal(t, age, props.Age.Value)
	assert.True(t, time.Now().After(props.RecordingDate), "RecordingDate should be in the past")
}

func TestNewUser_invalidName(t *testing.T) {
	t.Parallel()

	factory := user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     4,
		NameMaxLen:     100,
		LastNameMinLen: 0,
		LastNameMaxLen: 100,
		MinAge:         user.NewAge(0),
	})
	_, err := factory.NewUser("Don", "", 0)
	assert.NotNil(t, err)
	assert.IsType(t, &user.ModelValidationError{}, err)
	verr, _ := err.(*user.ModelValidationError)
	assert.Equal(t, "Name", verr.FieldErrors[0].Field, "Error should refer to the Name field")

	factory = user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     0,
		NameMaxLen:     6,
		LastNameMinLen: 0,
		LastNameMaxLen: 100,
		MinAge:         user.NewAge(0),
	})
	_, err = factory.NewUser("Maximilian", "", 0)
	assert.NotNil(t, err)
	assert.IsType(t, &user.ModelValidationError{}, err)
	verr, _ = err.(*user.ModelValidationError)
	assert.Equal(t, "Name", verr.FieldErrors[0].Field, "Error should refer to the Name field")
}

func TestNewUser_invalidLastName(t *testing.T) {
	t.Parallel()

	factory := user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     0,
		NameMaxLen:     100,
		LastNameMinLen: 4,
		LastNameMaxLen: 100,
		MinAge:         user.NewAge(0),
	})
	_, err := factory.NewUser("", "Lee", 0)
	assert.NotNil(t, err)
	assert.IsType(t, &user.ModelValidationError{}, err)
	verr, _ := err.(*user.ModelValidationError)
	assert.Equal(t, "LastName", verr.FieldErrors[0].Field, "Error should refer to the LastName field")

	factory = user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     0,
		NameMaxLen:     100,
		LastNameMinLen: 0,
		LastNameMaxLen: 6,
		MinAge:         user.NewAge(0),
	})
	_, err = factory.NewUser("", "Peterson", 0)
	assert.NotNil(t, err)
	assert.IsType(t, &user.ModelValidationError{}, err)
	verr, _ = err.(*user.ModelValidationError)
	assert.Equal(t, "LastName", verr.FieldErrors[0].Field, "Error should refer to the LastName field")
}

func TestNewUser_invalidAge(t *testing.T) {
	t.Parallel()

	factory := user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     0,
		NameMaxLen:     100,
		LastNameMinLen: 0,
		LastNameMaxLen: 100,
		MinAge:         user.NewAge(16),
	})
	_, err := factory.NewUser("", "", 15)
	assert.NotNil(t, err)
	assert.IsType(t, &user.ModelValidationError{}, err)
	verr, _ := err.(*user.ModelValidationError)
	assert.Equal(t, "Age", verr.FieldErrors[0].Field, "Error should refer to the Age field")
}
