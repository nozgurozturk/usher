package handler

import (
	"net/http"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/query"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
	. "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public"
)

// TicketHandler represents all ticket handlers.
type TicketHandler interface {
	// (GET /user/{userID}/tickets)
	GetUserTickets(w http.ResponseWriter, r *http.Request, userID string)
}

type ticketHandler struct {
	app *application.Application
}

// NewTicketHandler creates a new ticketHandler
func NewTicketHandler(app *application.Application) ticketHandler {
	return ticketHandler{
		app: app,
	}
}

func (h ticketHandler) GetUserTickets(w http.ResponseWriter, r *http.Request, userID string) {
	ts, err := h.app.Queries.GetUserTickets.Handle(query.GetUserTicketsQuery{
		UserID: userID,
	})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting tickets",
			Message: err.Error(),
		})
		return
	}

	resp := UserTicketsResponse{}
	for _, t := range ts {
		event := Event{
			Id:          t.Event.ID(),
			Name:        t.Event.Name(),
			Description: t.Event.Description(),
			StartDate:   t.Event.StartAt(),
			EndDate:     t.Event.EndAt(),
		}

		row, col := t.Seat.Position()
		seat := &Seat{
			Number: t.Seat.Number(),
			Position: struct {
				Col int "json:\"col\""
				Row int "json:\"row\""
			}{col, row},
			Rank: t.Seat.Rank(),
		}

		userName := t.User.Name
		user := User{
			Id:   t.User.ID,
			Name: &userName,
		}

		ticket := Ticket{
			Id:   t.ID,
			Seat: seat,
			User: user,
		}

		ticket.Event = struct {
			Event    "yaml:\",inline\""
			Location struct {
				Id   string "json:\"id\""
				Name string "json:\"name\""
			} "json:\"location\""
		}{Event: event}
		ticket.Event.Location.Id = t.Event.Hall().ID()
		ticket.Event.Location.Name = t.Event.Hall().Name()

		resp = append(resp, ticket)

	}

	rest.JSON(w, http.StatusOK, resp)

}
