package mongo

import (
	"context"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	client *mongo.Client
	config Config
}

type Config struct {
	Database    string
	ConnUri     string
	ConnTimeout time.Duration
}

func NewUserRepository(config Config) (*UserRepository, error) {
	opts := options.Client().SetConnectTimeout(config.ConnTimeout).
		ApplyURI(config.ConnUri)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		client: client,
		config: config,
	}, nil
}

type userModel struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	LastName      string             `bson:"last_name"`
	Age           int                `bson:"age"`
	RecordingDate time.Time          `bson:"recording_date"`
}

func (r *UserRepository) CreateUser(u *user.User) (user.UserId, error) {
	panic("Not implemented")
}

func (r *UserRepository) GetUserById(uid user.UserId) (*user.User, bool, error) {
	panic("Not implemented")
}
func (r *UserRepository) ListUsers(user.PaginatedListUsersQuery) ([]*user.User, error) {
	panic("Not implemented")
}

func (r *UserRepository) userCollection() *mongo.Collection {
	return r.client.Database(r.config.Database).Collection("users")
}

func modelToMongo(u *user.User) userModel {
	id, _ := primitive.ObjectIDFromHex(u.GetId().Value)
	props := u.GetProps()
	return userModel{
		Id:            id,
		Name:          props.Name,
		LastName:      props.LastName,
		Age:           int(props.Age.Value),
		RecordingDate: props.RecordingDate,
	}
}

func mongoToModel(u userModel) *user.User {
	return user.BuildUser(
		user.UserId{Value: u.Id.String()},
		user.UserProps{
			Name:          u.Name,
			LastName:      u.LastName,
			Age:           user.NewAge(uint8(u.Age)),
			RecordingDate: u.RecordingDate,
		},
	)
}
