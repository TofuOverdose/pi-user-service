package app

import (
	"github.com/TofuOverdose/pi-user-service/internal/users/app/commands"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/queries"
)

type UserApp struct {
	Commands UserAppCommands
	Queries  UserAppQueries
}

type UserAppCommands struct {
	CreateUser commands.CreateUserCommand
}

type UserAppQueries struct {
	GetUserById queries.GetUserByIdQuery
	ListUsers   queries.ListUsersQuery
}
