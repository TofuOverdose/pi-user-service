package main

import (
	"fmt"
	"os"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/adapters/mongo"
	"github.com/TofuOverdose/pi-user-service/internal/users/app"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/commands"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/queries"
	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
)

func makeApp() *app.UserApp {
	// Setup all the dependencies
	userFactory := userFactory()
	userRepository := userRepository()

	// and build an app
	return &app.UserApp{
		Commands: app.UserAppCommands{
			CreateUser: commands.CreateUserCommand{
				UserFactory:    userFactory,
				UserRepository: userRepository,
			},
		},
		Queries: app.UserAppQueries{
			GetUserById: queries.GetUserByIdQuery{
				UserRepository: userRepository,
				DateTimeFormat: time.RFC3339,
			},
		},
	}
}

func userFactory() *user.UserFactory {
	return user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen: 4,
		NameMaxLen: 20,

		LastNameMinLen: 3,
		LastNameMaxLen: 30,

		MinAge: user.NewAge(16),
	})
}

func userRepository() user.UserRepository {
	repo, err := mongo.NewUserRepository(mongo.Config{
		ConnTimeout: 10 * time.Second,
		ConnUri:     getEnvString("MONGO_URI"),
		Database:    getEnvString("MONGO_DATABASE"),
	})
	if err != nil {
		panic(err)
	}
	return repo
}

func getEnvString(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Sprintf("Environment variable %s not found\n", key))
	}

	return value
}
