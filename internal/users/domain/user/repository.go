package user

import "errors"

type UserRepository interface {
	CreateUser(User) (UserId, error)
	GetUserById(UserId) (*User, error)
	ListUsers(PaginatedListUsersQuery) ([]*User, error)
}

// ListUserQuery specifies the filters for listing the users
type ListUsersQuery struct {
	Name     *string
	LastName *string
	Age      *AgeQuery
}

type PaginatedListUsersQuery struct {
	Offset uint
	Limit  uint
	Query  ListUsersQuery
}

type AgeQuery struct {
	UpperBound *Age
	LowerBound *Age
}

func AgeLessThan(age Age) AgeQuery {
	return AgeQuery{
		UpperBound: &age,
	}
}

func AgeMoreThan(age Age) AgeQuery {
	return AgeQuery{
		LowerBound: &age,
	}
}

func AgeBetween(lowerBound, upperBound Age) (*AgeQuery, error) {
	if lowerBound.LessThan(upperBound) {
		return nil, errors.New("Invalid query")
	}

	return &AgeQuery{
		LowerBound: &lowerBound,
		UpperBound: &upperBound,
	}, nil
}
