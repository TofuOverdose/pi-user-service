package queries

import (
	"context"
	"log"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

type GetUserByIdQuery struct {
	UserRepository user.UserRepository
}

func (c *GetUserByIdQuery) Execute(ctx context.Context, id string) (*User, error) {
	usr, found, err := c.UserRepository.GetUserById(user.UserId{Value: id})
	if err != nil {
		log.Println("ERROR: failed to get user by ID", err.Error())
		return nil, err
	}

	if !found {
		return nil, ErrUserNotFound{}
	}

	return marshalUser(usr), nil
}
