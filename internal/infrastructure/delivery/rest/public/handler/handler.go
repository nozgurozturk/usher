package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	middlewareChi "github.com/go-chi/chi/v5/middleware"
	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public"
	"github.com/nozgurozturk/usher/internal/server/middleware"
)

type handlers struct {
	EventHandler
	CheckoutHandler
	UserHandler
	TicketHandler
}

func NewAPIHandler(app *application.Application) http.Handler {
	apiRouter := chi.NewRouter()

	apiRouter.Use(
		middlewareChi.Logger,
		middlewareChi.Recoverer,
		middlewareChi.AllowContentType("application/json"),
		middleware.CORS,
	)

	return public.HandlerFromMux(handlers{
		EventHandler:    NewEventHandler(app),
		CheckoutHandler: NewCheckoutHandler(app),
		UserHandler:     NewUserHandler(app),
		TicketHandler:   NewTicketHandler(app),
	}, apiRouter)
}
