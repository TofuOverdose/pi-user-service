package app

import (
	"example.com/TofuOverdose/pi-user-service/internal/users/app/commands"
	"example.com/TofuOverdose/pi-user-service/internal/users/app/queries"
)

type UserApp struct {
}

type UserAppQueries struct {
	CreateUser commands.CreateUserCommand
}

type UserAppCommands struct {
	GetUserById queries.GetUserByIdQuery
}
