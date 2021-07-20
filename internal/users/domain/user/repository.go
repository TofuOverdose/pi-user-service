package user

type UserRepository interface {
	CreateUser(*User) (UserId, error)
	GetUserById(UserId) (*User, bool, error)
	ListUsers(ListUsersQuery) (*ListUsersResult, error)
}

type ListUsersResult struct {
	Users []*User
}

// ListUserQuery specifies the filters and sorting for the users listing
type ListUsersQuery struct {
	SortSpec SortSpec
}

type SortSpec struct {
	Field SortField
	Order SortOrder
}

type SortField string

var (
	SortFieldAge           SortField = "age"
	SortFieldRecordingDate SortField = "recording_date"
)

type SortOrder int

var (
	SortDesc SortOrder = -1
	SortAsc  SortOrder = 1
)
