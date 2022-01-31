package handler

import (
	"net/http"

	"github.com/nozgurozturk/usher/internal/application"
	"github.com/nozgurozturk/usher/internal/application/command"
	"github.com/nozgurozturk/usher/internal/application/query"
	"github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest"
	. "github.com/nozgurozturk/usher/internal/infrastructure/delivery/rest/admin"
)

// EventHandler is a REST handler for event
type EventHandler interface {
	// (GET /events)
	GetEvents(w http.ResponseWriter, r *http.Request)

	// (GET /events/{eventID})
	GetEvent(w http.ResponseWriter, r *http.Request, eventID string)

	// (POST /events)
	CreateEvent(w http.ResponseWriter, r *http.Request)

	// (POST /events/{eventID}/reserve)
	ReserveEvent(w http.ResponseWriter, r *http.Request, eventID string)
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
		sections := make([]Section, len(event.Hall().Sections()))
		for i, section := range event.Hall().Sections() {
			sections[i].Name = section.Name()
			rows := make([]Row, len(section.Rows()))
			for j, row := range section.Rows() {
				rows[j].Name = row.Name()
				seats := make([]Seat, len(row.Seats()))
				for k, seat := range row.Seats() {
					c, r := seat.Position()
					f := int(seat.Feature())
					seats[k].Id = seat.ID()
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
	event := Event{
		Description: e.Description(),
		EndDate:     e.EndAt(),
		Id:          e.ID(),
		Name:        e.Name(),
		StartDate:   e.StartAt(),
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

	layout := Hall{
		Id:       e.Hall().ID(),
		Name:     e.Hall().Name(),
		Sections: sections,
	}

	groups, err := h.app.Queries.GetEventReservations.Handle(query.GetEventReservationsQuery{
		EventID: eventID,
	})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting event reservations",
			Message: err.Error(),
		})
		return
	}

	reservations := make([]Reservation, len(groups))

	for i, group := range groups {
		feat := int(group.SeatPreferences())
		reservations[i] = Reservation{
			Event: event,
			Id:    group.ID(),
			Preferences: struct {
				Features *int "json:\"features,omitempty\""
				Rank     int  "json:\"rank\""
			}{
				Features: &feat,
				Rank:     group.RankPreference(),
			},
			Size: group.Size(),
			User: User{},
		}
	}

	t, err := h.app.Queries.GetEventTickets.Handle(query.GetEventTicketsQuery{
		EventID: eventID,
	})

	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error getting event tickets",
			Message: err.Error(),
		})
		return
	}

	tickets := make([]Ticket, len(t))

	for i, ticket := range t {
		row, col := ticket.Seat.Position()
		feat := int(ticket.Seat.Feature())
		seat := Seat{
			Available: ticket.Seat.Available(),
			Features:  &feat,
			Number:    ticket.Seat.Number(),
			Position: struct {
				Col int "json:\"col\""
				Row int "json:\"row\""
			}{
				Col: col,
				Row: row,
			},
			Rank: ticket.Seat.Rank(),
		}
		tickets[i] = Ticket{
			Event: event,
			Id:    ticket.ID,
			Seat:  &seat,
			User: User{
				Id: ticket.User.ID,
			},
		}
	}

	resp := EventResponse{
		Event:        event,
		Layout:       layout,
		Reservations: reservations,
		Tickets:      tickets,
	}

	rest.JSON(w, http.StatusOK, resp)
}

// CheckEventSeats checks seats for an event
func (h eventHandler) ReserveEvent(w http.ResponseWriter, r *http.Request, eventID string) {
	err := h.app.Commands.ReserveSeats.Handle(command.ReserverSeatsCommand{
		EventID: eventID,
	})
	if err != nil {
		rest.JSON(w, http.StatusInternalServerError, ErrorResponse{
			Error:   "Error reserving event",
			Message: err.Error(),
		})
		return
	}

	rest.JSON(w, http.StatusOK, nil)
}

func (h eventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {}
