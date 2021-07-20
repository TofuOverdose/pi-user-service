package http

import (
	"net/http"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/app"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/commands"
	"github.com/TofuOverdose/pi-user-service/internal/users/app/queries"
	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type server struct {
	App *app.UserApp
}

func (s *server) CreateUser(res http.ResponseWriter, req *http.Request) {
	data := &CreateUserRequest{}
	err := render.Bind(req, data)
	if err != nil {
		render.Render(res, req, ServerErrorResponse)
		return
	}

	uid, err := s.App.Commands.CreateUser.Execute(req.Context(), commands.CreateUserCommandArgs{
		Name:     data.Name,
		LastName: data.LastName,
		Age:      uint8(data.Age),
	})
	if err != nil {
		switch e := err.(type) {
		case *user.ModelValidationError:
			render.Render(res, req, CreateUserInputErrorResponse(e))
			return
		default:
			render.Render(res, req, ServerErrorResponse)
			return
		}
	}

	render.Render(res, req, &CreateUserResponse{
		UserId: uid,
	})
}

func (s *server) GetUserById(res http.ResponseWriter, req *http.Request) {
	uid := chi.URLParam(req, "userId")
	usr, err := s.App.Queries.GetUserById.Execute(req.Context(), uid)
	if err != nil {
		switch err.(type) {
		case queries.ErrUserNotFound:
			render.Render(res, req, NotFoundErrorResponse)
			return
		default:
			render.Render(res, req, ServerErrorResponse)
			return
		}
	}
	data := marshalUserData(usr)
	render.Render(res, req, data)
}

func (s *server) ListUsers(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	results, err := s.App.Queries.ListUsers.Execute(req.Context(), queries.ListUsersQueryArgs{
		SortBy: query.Get("sortby"),
		Order:  query.Get("order"),
	})
	if err != nil {
		render.Render(res, req, ServerErrorResponse)
		return
	}
	render.Render(res, req, marshalListUserData(results))
}

func Serve(app *app.UserApp, config ServerConfig) error {
	router := chi.NewRouter()
	srv := server{
		App: app,
	}

	router.Use(
		middleware.AllowContentType("application/json"),
		middleware.Logger,
		middleware.Timeout(30*time.Second),
		middleware.CleanPath,
		middleware.StripSlashes,
	)

	router.Post("/users", srv.CreateUser)
	router.Get("/users/{userId}", srv.GetUserById)
	router.Get("/users", srv.ListUsers)

	return http.ListenAndServe(config.Port, router)
}

type ServerConfig struct {
	Port string
}
