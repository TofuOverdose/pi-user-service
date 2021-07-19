package ports

import (
	"net/http"
	"time"

	"github.com/TofuOverdose/pi-user-service/internal/users/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	App *app.UserApp
}

func (s *HttpServer) CreateUser(res http.ResponseWriter, req *http.Request) {

}

func (s *HttpServer) GetUserById(res http.ResponseWriter, req *http.Request) {

}

func (s *HttpServer) ListUsers(res http.ResponseWriter, req *http.Request) {

}

func ServeHttp(port string, app *app.UserApp) error {
	router := chi.NewRouter()
	server := HttpServer{
		App: app,
	}

	router.Use(
		middleware.Logger,
		middleware.Timeout(30*time.Second),
		middleware.CleanPath,
		middleware.StripSlashes,
	)

	router.Post("/users", server.CreateUser)
	router.Get("/users/{userId}", server.GetUserById)
	router.Get("/users", server.ListUsers)

	return http.ListenAndServe(port, router)
}
