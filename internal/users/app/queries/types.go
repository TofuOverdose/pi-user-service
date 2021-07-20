package queries

import "github.com/TofuOverdose/pi-user-service/internal/users/domain/user"

type User struct {
	Id            string
	Name          string
	LastName      string
	Age           uint8
	RecordingDate string
}

func marshalUser(usr *user.User, dateFormat string) *User {
	props := usr.GetProps()
	return &User{
		Id:            usr.GetId().Value,
		Name:          props.Name,
		LastName:      props.LastName,
		Age:           props.Age.Value,
		RecordingDate: props.RecordingDate.Format(dateFormat),
	}
}

type PaginatedUserList struct {
	Page    uint
	PerPage uint
	Data    []User
}

type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "User with given ID not found"
}
