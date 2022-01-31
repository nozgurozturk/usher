package handler

import (
	"net/http"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/command"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
)

type HelperHandler interface {
	// (POST /reset)
	Reset(w http.ResponseWriter, r *http.Request)
}

// HelperHandler is a REST handler for helper
type helperHandler struct {
	app *application.Application
}

// NewHelperHandler creates a new HelperHandler
func NewHelperHandler(app *application.Application) HelperHandler {
	return helperHandler{
		app: app,
	}
}

// Reset resets the application
func (h helperHandler) Reset(w http.ResponseWriter, r *http.Request) {
	h.app.Commands.ResetDatabase.Handle(command.ResetDatabaseCommand{})
	rest.JSON(w, http.StatusOK, nil)
}
