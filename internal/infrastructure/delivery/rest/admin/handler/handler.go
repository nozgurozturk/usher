package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	middlewareChi "github.com/go-chi/chi/v5/middleware"
	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/admin"
	"github.com/nozgurozturk/usher/internal/server/middleware"
)

type handlers struct {
	EventHandler
	LayoutHandler

	HelperHandler
}

func NewAPIHandler(app *application.Application) http.Handler {
	adminRouter := chi.NewRouter()

	adminRouter.Use(
		middlewareChi.AllowContentType("application/json"),
		middleware.CORS,
	)

	return admin.HandlerFromMux(handlers{
		EventHandler:  NewEventHandler(app),
		LayoutHandler: NewLayoutHandler(app),
		HelperHandler: NewHelperHandler(app),
	}, adminRouter)
}
