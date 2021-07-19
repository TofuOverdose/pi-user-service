package mongo_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/adapters/mongo"
	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"github.com/stretchr/testify/assert"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	moptions "go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewUserRepository(t *testing.T) {
	repo, err := mongo.NewUserRepository(repoConfig())
	assert.Nil(t, err)
	assert.NotNil(t, repo)
}

func TestCreateUser(t *testing.T) {
	genUser := makeGenTestUser(testUserFactory())
	usr := genUser()
	repo, drop := testRepo("TestCreateUser")
	defer drop()

	uid, err := repo.CreateUser(usr)
	assert.Nil(t, err)
	assert.NotEqual(t, "", uid.Value)
}

func TestGetUserById(t *testing.T) {
	t.Parallel()

	genUser := makeGenTestUser(testUserFactory())
	repo, drop := testRepo("TestGetUserById")
	defer drop()

	users := make([]user.UserId, 3)
	for i := 0; i < cap(users); i++ {
		uid, err := repo.CreateUser(genUser())
		if err != nil {
			t.Fatal(err)
		}
		users[i] = uid
	}

	for i := 0; i < len(users); i++ {
		uid := users[i].Value
		usr, found, err := repo.GetUserById(user.UserId{Value: uid})
		assert.Nil(t, err)
		assert.True(t, found)
		assert.NotNil(t, usr)
		assert.Equal(t, uid, usr.GetId().Value)
	}
}

func TestListUsers_Pagination(t *testing.T) {
	t.Parallel()

	genUser := makeGenTestUser(testUserFactory())
	repo, drop := testRepo("TestListUsers_Pagination")
	defer drop()

	usersCount := 20
	users := make([]user.UserId, usersCount)
	for i := 0; i < cap(users); i++ {
		uid, err := repo.CreateUser(genUser())
		if err != nil {
			t.Fatal(err)
		}
		users[i] = uid
	}

	// All users should be listed properly
	query := user.PaginatedListUsersQuery{
		Offset: 0,
		Limit:  uint(usersCount + 1),
	}
	res, err := repo.ListUsers(query)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(users), len(res))
}

func testUser() *user.User {
	u, _ := user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen:     0,
		NameMaxLen:     100,
		LastNameMinLen: 0,
		LastNameMaxLen: 100,
		MinAge:         user.NewAge(30),
	}).NewUser("John", "Doe", 30)
	return u
}

func testUserFactory() *user.UserFactory {
	return user.NewUserFactory(user.UserPropsConstraints{
		NameMinLen: 0,
		NameMaxLen: 255,

		LastNameMinLen: 0,
		LastNameMaxLen: 255,

		MinAge: user.NewAge(0),
	})
}

func makeGenTestUser(factory *user.UserFactory) func() *user.User {
	id := 0

	f := func() *user.User {
		u, _ := factory.NewUser(
			fmt.Sprintf("name_%d", id),
			fmt.Sprintf("last-name_%d", id),
			uint8(20+id),
		)
		id += 1
		return u
	}

	return f
}

func repoConfig() mongo.Config {
	return mongo.Config{
		ConnUri:     os.Getenv("MONGO_URI"),
		ConnTimeout: 5 * time.Second,
		Database:    os.Getenv("MONGO_DATABASE"),
	}
}

func testRepo(database string) (*mongo.UserRepository, func()) {
	cfg := repoConfig()
	cfg.Database = database
	repo, _ := mongo.NewUserRepository(cfg)

	dropFunc := func() {
		conn, err := mongodb.Connect(context.TODO(), moptions.Client().ApplyURI(cfg.ConnUri))
		if err != nil {
			panic(err)
		}
		err = conn.Database(cfg.Database).Drop(context.TODO())
		if err != nil {
			panic(err)
		}
	}

	return repo, dropFunc
}
