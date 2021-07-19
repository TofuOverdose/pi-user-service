package queries

import (
	"context"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

type ListUsersQuery struct {
	UserRepository *user.UserRepository
}

type ListUsersQueryArgs struct{}

func (c *ListUsersQuery) Execute(ctx context.Context, args ListUsersQueryArgs) ([]User, error) {
	panic("Not implemented")
}
