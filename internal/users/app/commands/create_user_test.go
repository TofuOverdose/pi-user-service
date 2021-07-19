package commands_test

import (
	"context"
	"testing"

	"github.com/TofuOverdose/pi-user-service/internal/users/app/commands"
	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const mockUserid = "test-user-id"

type repoMock struct {
	mock.Mock
}

func (r *repoMock) CreateUser(u *user.User) (user.UserId, error) {
	props := u.GetProps()
	r.Called(props.Name, props.LastName, props.Age)
	return user.UserId{Value: mockUserid}, nil
}

func (r *repoMock) GetUserById(user.UserId) (*user.User, bool, error) {
	return nil, false, nil
}

func (r *repoMock) ListUsers(user.PaginatedListUsersQuery) ([]*user.User, error) {
	return nil, nil
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	var (
		name     string = "John"
		lastName string = "Doe"
		age      uint8  = 30
	)

	repo := &repoMock{}
	repo.On("CreateUser", name, lastName, user.NewAge(age))

	cmd := commands.CreateUserCommand{
		UserRepository: repo,
		UserFactory: user.NewUserFactory(user.UserPropsConstraints{
			NameMinLen:     0,
			NameMaxLen:     100,
			LastNameMinLen: 0,
			LastNameMaxLen: 100,
			MinAge:         user.NewAge(0),
		}),
	}

	id, err := cmd.Execute(context.TODO(), commands.CreateUserCommandArgs{
		Name:     name,
		LastName: lastName,
		Age:      age,
	})
	repo.AssertCalled(t, "CreateUser", name, lastName, user.NewAge(age))
	assert.Nil(t, err)
	assert.Equal(t, mockUserid, id)
}

func TestCreateUser_WithInvalidInput(t *testing.T) {
	t.Parallel()

	var (
		name     string = "John"
		lastName string = "Doe"
		age      uint8  = 30
	)

	repo := &repoMock{}
	repo.On("CreateUser", name, lastName, user.NewAge(age))

	// Should fail all validation rules
	cmd := commands.CreateUserCommand{
		UserRepository: repo,
		UserFactory: user.NewUserFactory(user.UserPropsConstraints{
			NameMinLen:     100,
			NameMaxLen:     100,
			LastNameMinLen: 100,
			LastNameMaxLen: 100,
			MinAge:         user.NewAge(100),
		}),
	}
	id, err := cmd.Execute(context.TODO(), commands.CreateUserCommandArgs{
		Name:     name,
		LastName: lastName,
		Age:      age,
	})
	repo.AssertNotCalled(t, "CreateUser", name, lastName, user.NewAge(age))
	assert.Equal(t, "", id)
	assert.Error(t, err)
	assert.IsType(t, &user.ModelValidationError{}, err)

	failedFields := make(map[string]bool)
	ferrs := err.(*user.ModelValidationError).FieldErrors
	assert.NotEmpty(t, ferrs)
	for _, e := range ferrs {
		failedFields[e.Field] = true
	}
	assert.True(t, failedFields["Name"], "Should fail validation for name")
	assert.True(t, failedFields["LastName"], "Should fail validation for last name")
	assert.True(t, failedFields["Age"], "Should fail validation for age")
}
