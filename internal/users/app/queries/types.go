package queries

type User struct {
	Id            string
	Name          string
	LastName      string
	Age           uint8
	RecordingDate string
}

type PaginatedUserList struct {
	Page    uint
	PerPage uint
	Data    []User
}

type ErrUserNotFound struct {}

func (e *ErrUserNotFound) Error() string {
	return "User with given ID not found"
}
