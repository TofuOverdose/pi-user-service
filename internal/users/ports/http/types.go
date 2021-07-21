package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TofuOverdose/pi-user-service/internal/users/app/queries"
	"github.com/TofuOverdose/pi-user-service/internal/users/domain/user"
	"github.com/go-chi/render"
)

type UserResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	LastName      string `json:"lastName"`
	Age           int    `json:"age"`
	RecordingDate int64  `json:"recordingDate"`
}

func (d *UserResponse) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, 200)
	return nil
}

func marshalUserData(usr *queries.User) *UserResponse {
	return &UserResponse{
		Id:            usr.Id,
		Name:          usr.Name,
		LastName:      usr.LastName,
		Age:           int(usr.Age),
		RecordingDate: usr.RecordingDate,
	}
}

type UserListResponse struct {
	Data []*UserResponse `json:"data"`
}

func (d *UserListResponse) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, 200)
	return nil
}

func marshalListUserData(usrs *queries.ListUsersQueryRes) *UserListResponse {
	out := make([]*UserResponse, len(usrs.Data))
	for i, usr := range usrs.Data {
		out[i] = marshalUserData(usr)
	}
	return &UserListResponse{
		Data: out,
	}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}

func (d *CreateUserRequest) Bind(req *http.Request) error {
	return nil
}

type CreateUserResponse struct {
	UserId string `json:"userId"`
}

func (d *CreateUserResponse) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, 201)
	return nil
}

/* Error Response Payloads */

type ErrorResponse struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("HTTP Error %d: %s", e.Code, e.Message)
}

func (e *ErrorResponse) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, e.Code)
	return nil
}

var ServerErrorResponse = &ErrorResponse{Code: 500, Message: "Server Error"}

var NotFoundErrorResponse = &ErrorResponse{Code: 404, Message: "Resource not found"}

func CreateUserInputErrorResponse(err *user.ModelValidationError) *ErrorResponse {
	data := make(map[string]string)
	for _, fe := range err.FieldErrors {
		data[strings.ToLower(fe.Field)] = fe.Message
	}

	return &ErrorResponse{
		Code:    400,
		Message: "Bad Request",
		Data:    data,
	}
}
