package grpc

import (
	"context"
	"net"

	"github.com/TofuOverdose/pi-user-service/internal/users/app"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/commands"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/queries"
	proto "github.com/TofuOverdose/pi-user-service/internal/users/ports/grpc/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	proto.UnimplementedUsersServiceServer
	App *app.UserApp
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	uid, err := s.App.Commands.CreateUser.Execute(ctx, commands.CreateUserCommandArgs{
		Name:     req.Name,
		LastName: req.LastName,
		Age:      uint8(req.Age),
	})
	if err != nil {
		switch e := err.(type) {
		case commands.ErrWrongInput:
			return nil, status.Error(codes.InvalidArgument, e.Error())
		default:
			return nil, status.Error(codes.Unknown, "Server Error")
		}
	}
	return &proto.CreateUserResponse{
		UserId: uid,
	}, nil
}

func (s *server) GetUserById(ctx context.Context, req *proto.GetUserByIdRequest) (*proto.User, error) {
	usr, err := s.App.Queries.GetUserById.Execute(ctx, req.UserId)
	if err != nil {
		switch e := err.(type) {
		case queries.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, e.Error())
		default:
			return nil, status.Error(codes.Unknown, "Server Error")
		}
	}
	return &proto.User{
		Id:            usr.Id,
		Name:          usr.Name,
		LastName:      usr.LastName,
		Age:           uint32(usr.Age),
		RecordingDate: usr.RecordingDate,
	}, nil
}

func (s *server) ListUsers(context.Context, *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Endpoint Not Implemented")
}

func Serve(app *app.UserApp, config ServerConfig) error {
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	proto.RegisterUsersServiceServer(srv, &server{
		App: app,
	})
	if config.UseReflection {
		reflection.Register(srv)
	}
	return srv.Serve(lis)
}

type ServerConfig struct {
	Port          string
	UseReflection bool
}
