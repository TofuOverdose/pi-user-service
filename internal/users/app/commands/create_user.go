package commands

import (
	"context"
	"log"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

type CreateUserCommand struct {
	UserRepository user.UserRepository
	UserFactory    *user.UserFactory
}

type CreateUserCommandArgs struct {
	Name     string
	LastName string
	Age      uint8
}

func (c *CreateUserCommand) Execute(ctx context.Context, args CreateUserCommandArgs) (string, error) {
	usr, err := c.UserFactory.NewUser(args.Name, args.LastName, args.Age)
	if err != nil {
		switch e := err.(type) {
		case user.ModelValidationError:
			return "", ErrWrongInput{e}
		default:
			log.Println("ERROR: failed to create user", e.Error())
			return "", e
		}
	}
	id, err := c.UserRepository.CreateUser(usr)
	if err != nil {
		log.Println("ERROR: failed to persist new user", err.Error())
		return "", err
	}

	return id.Value, nil
}

type ErrWrongInput struct {
	user.ModelValidationError
}
