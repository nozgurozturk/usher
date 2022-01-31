package handler

import (
	"net/http"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/query"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
	. "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/public"
)

// EventHandler is a REST handler for event
type EventHandler interface {
	// (GET /events)
	GetEvents(w http.ResponseWriter, r *http.Request)

	// (GET /events/{eventID})
	GetEvent(w http.ResponseWriter, r *http.Request, eventID string)

	// (POST /events/{eventID}/check)
	CheckEventSeats(w http.ResponseWriter, r *http.Request, eventID string)
}

type eventHandler struct {
	app *application.Application
}

// NewEventHandler creates a new EventHandler
func NewEventHandler(app *application.Application) EventHandler {
	return eventHandler{
		app: app,
	}
}

// GetEvents gets all events
func (h eventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	e, err := h.app.Queries.GetEvents.Handle(query.GetEventsQuery{})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting events",
			Message: err.Error(),
		})
		return
	}

	resp := make(EventsResponse, len(e))
	for i, event := range e {
		resp[i].Id = event.ID()
		resp[i].Name = event.Name()
		resp[i].Description = event.Description()
		resp[i].StartDate = (event).StartAt()
		resp[i].EndDate = event.EndAt()
		if event.Hall() != nil {
			resp[i].LocationID = event.Hall().ID()
		}
	}

	rest.JSON(w, http.StatusOK, resp)
}

// GetEvent gets an event
func (h eventHandler) GetEvent(w http.ResponseWriter, r *http.Request, eventID string) {
	e, err := h.app.Queries.GetEvent.Handle(query.GetEventQuery{
		EventID: eventID,
	})
	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting event",
			Message: err.Error(),
		})
		return
	}

	sections := make([]Section, len(e.Hall().Sections()))
	for i, section := range e.Hall().Sections() {
		sections[i].Name = section.Name()
		rows := make([]Row, len(section.Rows()))
		for j, row := range section.Rows() {
			rows[j].Name = row.Name()
			seats := make([]Seat, len(row.Seats()))
			for k, seat := range row.Seats() {
				c, r := seat.Position()
				f := int(seat.Feature())
				seats[k].Rank = seat.Rank()
				seats[k].Available = seat.Available()
				seats[k].Number = seat.Number()
				seats[k].Position.Col = c
				seats[k].Position.Row = r
				seats[k].Features = &f

			}
			rows[j].Seats = seats
		}
		sections[i].Rows = rows
	}

	resp := EventResponse{
		Event: Event{
			Description: e.Description(),
			EndDate:     e.EndAt(),
			Id:          e.ID(),
			Name:        e.Name(),
			StartDate:   e.StartAt(),
		},
		Location: Hall{
			Name:     e.Hall().Name(),
			Sections: sections,
		},
	}

	rest.JSON(w, http.StatusOK, resp)
}

// CheckEventSeats checks seats for an event
func (h eventHandler) CheckEventSeats(w http.ResponseWriter, r *http.Request, eventID string) {

	var body CheckEventSeatsRequest
	if err := rest.ParseBody(r, &body); err != nil {
		rest.JSON(w, http.StatusUnprocessableEntity, ErrorResponse{
			Error:   "Error parsing body",
			Message: err.Error(),
		})
		return
	}

	remaining, err := h.app.Queries.CheckEventSeats.Handle(query.CheckEventSeatsQuery{
		EventID:  eventID,
		Count:    body.Count,
		Features: body.Features,
		Rank:     body.Rank,
	})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error checking event seats",
			Message: err.Error(),
		})
		return
	}

	resp := CheckEventSeatsResponse{
		Remaining: remaining,
	}

	rest.JSON(w, http.StatusOK, resp)

}
