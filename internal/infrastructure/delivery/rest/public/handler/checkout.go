package handler

import (
	"net/http"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/command"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
	. "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public"
)

// CheckoutHandler represents all checkout handlers.
type CheckoutHandler interface {
	// (POST /checkout)
	Checkout(w http.ResponseWriter, r *http.Request)
}

type checkoutHandler struct {
	app *application.Application
}

func NewCheckoutHandler(app *application.Application) CheckoutHandler {
	return checkoutHandler{
		app: app,
	}
}

func (h checkoutHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var body CheckoutRequest
	if err := rest.ParseBody(r, &body); err != nil {
		rest.JSON(w, http.StatusUnprocessableEntity, ErrorResponse{
			Error:   "Error parsing body",
			Message: err.Error(),
		})
		return
	}

	if err := h.app.Commands.CreateReservation.Handle(command.CreateReservationCommand{
		Count:   body.Count,
		EventID: body.EventID,
		Preferences: struct {
			Features *int
			Rank     *int
		}{body.Preferences.Features, body.Preferences.Rank},
		UserID: body.UserID,
	}); err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error creating reservation",
			Message: err.Error(),
		})
	}

	rest.JSON(w, http.StatusOK, CheckoutResponse{
		Count:       body.Count,
		EventID:     body.EventID,
		Preferences: body.Preferences,
		UserID:      body.UserID,
	})

}
