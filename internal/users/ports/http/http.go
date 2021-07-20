package http

import (
	"net/http"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	App *app.UserApp
}

func (s *server) CreateUser(res http.ResponseWriter, req *http.Request) {

}

func (s *server) GetUserById(res http.ResponseWriter, req *http.Request) {

}

func (s *server) ListUsers(res http.ResponseWriter, req *http.Request) {

}

func Serve(app *app.UserApp, config ServerConfig) error {
	router := chi.NewRouter()
	srv := server{
		App: app,
	}

	router.Use(
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
