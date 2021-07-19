package mongo

import (
	"testing"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

func testUser() *user.User {
	u, _ := user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     0,
		NameMaxLen:     100,
		LastNameMinLen: 0,
		LastNameMaxLen: 100,
		MinAge:         user.NewAge(30),
	}).NewUser("John", "Doe", 30)
	return u
}

func TestNewUserRepository(t *testing.T) {
	// t.Parallel()
}
