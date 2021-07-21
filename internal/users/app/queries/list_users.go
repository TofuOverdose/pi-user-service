package queries

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

type ListUsersQuery struct {
	UserRepository user.UserRepository
}

func (c *ListUsersQuery) Execute(ctx context.Context, args ListUsersQueryArgs) (*ListUsersQueryRes, error) {
	query, err := argsToQuery(args)
	if err != nil {
		return nil, err
	}
	res, err := c.UserRepository.ListUsers(*query)
	if err != nil {
		log.Println("ERROR: failed to list users", err.Error())
		return nil, err
	}
	data := make([]*User, len(res.Users))
	for i, u := range res.Users {
		data[i] = marshalUser(u)
	}
	return &ListUsersQueryRes{
		Data: data,
	}, nil
}

type ListUsersQueryArgs struct {
	SortBy string
	Order  string
}

type ListUsersQueryRes struct {
	Data []*User
}

func argsToQuery(args ListUsersQueryArgs) (*user.ListUsersQuery, error) {

	sortField := user.SortFieldRecordingDate
	switch strings.ToLower(args.SortBy) {
	case "recording_date":
		sortField = user.SortFieldRecordingDate
	case "age":
		sortField = user.SortFieldAge
	}

	sortOrder := user.SortDesc
	switch strings.ToLower(args.Order) {
	case "desc":
		sortOrder = user.SortDesc
	case "asc":
		sortOrder = user.SortAsc
	}

	return &user.ListUsersQuery{
		SortSpec: user.SortSpec{
			Field: sortField,
			Order: sortOrder,
		},
	}, nil
}

var ErrIncorrectQuery = errors.New("Incorrect query input")
