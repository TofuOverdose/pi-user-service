package queries

import (
	"context"
	"log"

	"example.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

type GetUserByIdQuery struct {
	UserRepository user.UserRepository
	DateTimeFormat string
}

func (c *GetUserByIdQuery) Execute(ctx context.Context, id string) (*User, error) {
	usr, found, err := c.UserRepository.GetUserById(user.UserId{Value: id})
	if err != nil {
		log.Println("ERROR: failed to get user by ID", err.Error())
		return nil, err
	}

	if !found {
		return nil, &ErrUserNotFound{}
	}

	props := usr.GetProps()
	return &User{
		Id:            usr.GetId().Value,
		Name:          props.Name,
		LastName:      props.LastName,
		Age:           props.Age.Value,
		RecordingDate: props.RecordingDate.Format(c.DateTimeFormat),
	}, nil
}
