package mongo

import (
	"context"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"go.mongodb.org/mongo-driver/bson"
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
	res, err := r.userCollection().InsertOne(context.TODO(), modelToMongo(u, true))
	if err != nil {
		return user.UserId{}, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return user.UserId{Value: id}, nil
}

func (r *UserRepository) GetUserById(uid user.UserId) (*user.User, bool, error) {
	var u userModel
	objId, err := primitive.ObjectIDFromHex(uid.Value)
	if err != nil {
		return nil, false, err
	}
	err = r.userCollection().FindOne(
		context.TODO(),
		bson.D{{"_id", objId}},
	).Decode(&u)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, false, err
	}

	return mongoToModel(u), true, nil
}
func (r *UserRepository) ListUsers(query user.PaginatedListUsersQuery) ([]*user.User, error) {
	filter, opts := buildFindQuery(query)
	cursor, err := r.userCollection().Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	var results []userModel
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	out := make([]*user.User, len(results))
	for i, r := range results {
		out[i] = mongoToModel(r)
	}
	return out, nil
}

func (r *UserRepository) userCollection() *mongo.Collection {
	return r.client.Database(r.config.Database).Collection("users")
}

func buildFindQuery(query user.PaginatedListUsersQuery) (bson.M, *options.FindOptions) {
	opts := options.Find().
		SetSort(bson.D{{"recording_date", -1}}).
		SetSkip(int64(query.Offset)).
		SetLimit(int64(query.Limit))

	filter := bson.M{}
	if name := query.Query.Name; name != nil {
		filter["name"] = bson.M{"$eq": name}
	}
	if lastName := query.Query.LastName; lastName != nil {
		filter["last_name"] = bson.M{"$eq": lastName}
	}
	if ageQuery := query.Query.Age; ageQuery != nil {
		ageFilter := bson.M{}
		if ageQuery.LowerBound != nil {
			ageFilter["$gte"] = *ageQuery.LowerBound
		}
		if ageQuery.UpperBound != nil {
			ageFilter["$lte"] = *ageQuery.UpperBound
		}
		filter["age"] = ageFilter
	}

	return filter, opts
}

func modelToMongo(u *user.User, newId bool) userModel {
	var id primitive.ObjectID
	if newId {
		id = primitive.NewObjectID()
	} else {
		id, _ = primitive.ObjectIDFromHex(u.GetId().Value)
	}
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
		user.UserId{Value: u.Id.Hex()},
		user.UserProps{
			Name:          u.Name,
			LastName:      u.LastName,
			Age:           user.NewAge(uint8(u.Age)),
			RecordingDate: u.RecordingDate,
		},
	)
}
