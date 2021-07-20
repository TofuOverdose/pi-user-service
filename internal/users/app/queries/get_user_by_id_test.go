package queries_test

import (
	"context"
	"testing"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/app/queries"
	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repoMock struct {
	mock.Mock
}

func (r *repoMock) CreateUser(*user.User) (user.UserId, error) {
	return user.UserId{}, nil
}

const (
	existingUserId    = "existing-user-id"
	nonExistingUserId = "non-existing-user-id"
)

func (r *repoMock) GetUserById(uid user.UserId) (*user.User, bool, error) {
	r.Called(uid)
	if uid.Value == existingUserId {
		return user.BuildUser(user.UserId{Value: existingUserId}, user.UserProps{
			Name:          "John",
			LastName:      "Doe",
			Age:           user.NewAge(30),
			RecordingDate: time.Now(),
		}), true, nil
	}

	return nil, false, nil
}

func (r *repoMock) ListUsers(user.PaginatedListUsersQuery) ([]*user.User, error) {
	return nil, nil
}

func TestGetUserById(t *testing.T) {
	t.Parallel()

	uid := user.UserId{Value: existingUserId}
	repo := &repoMock{}
	repo.On("GetUserById", uid)

	cmd := queries.GetUserByIdQuery{
		UserRepository: repo,
	}
	usr, err := cmd.Execute(context.TODO(), existingUserId)
	repo.AssertCalled(t, "GetUserById", uid)
	assert.Nil(t, err)
	assert.NotNil(t, usr)
	assert.Equal(t, existingUserId, usr.Id)
}

func TestGetUserById_WithWrongId(t *testing.T) {
	t.Parallel()

	uid := user.UserId{Value: nonExistingUserId}
	repo := &repoMock{}
	repo.On("GetUserById", uid)

	cmd := queries.GetUserByIdQuery{
		UserRepository: repo,
	}
	usr, err := cmd.Execute(context.TODO(), nonExistingUserId)
	repo.AssertCalled(t, "GetUserById", uid)
	assert.Nil(t, usr)
	assert.Error(t, err)
	assert.IsType(t, queries.ErrUserNotFound{}, err)
}
