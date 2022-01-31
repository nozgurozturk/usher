package handler

import (
	"net/http"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/query"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
	. "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public"
)

// UserHandler represents all user handlers.
type UserHandler interface {
	// (GET /user/{userID})
	GetUser(w http.ResponseWriter, r *http.Request, userID string)
	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	app *application.Application
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(app *application.Application) userHandler {
	return userHandler{
		app: app,
	}
}

// TODO: GetUser
func (h userHandler) GetUser(w http.ResponseWriter, r *http.Request, userID string) {}

func (h userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	us, err := h.app.Queries.GetUsers.Handle(query.GetUsersQuery{})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting users",
			Message: err.Error(),
		})
		return
	}

	var resp UsersResponse
	for _, u := range us {
		name := u.Name
		resp = append(resp, User{
			Id:   u.ID,
			Name: &name,
		})
	}

	rest.JSON(w, http.StatusOK, resp)
}
