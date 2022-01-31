package application

import (
	"github.com/nozgurozturk/usher/internal/application/command"
	"github.com/nozgurozturk/usher/internal/application/query"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateLayout      command.CreateLayoutHandler
	CreateReservation command.CreateReservationHandler
	ReserveSeats      command.ReserveSeatsHandler
	ResetDatabase     command.ResetDatabaseHandler
}
type Queries struct {
	GetEvent             query.GetEventHandler
	GetEvents            query.GetEventsHandler
	CheckEventSeats      query.CheckEventSeatsHandler
	GetUserTickets       query.GetUserTicketsHandler
	GetUserTicket        query.GetUserTicketHandler
	GetLayouts           query.GetLayoutsHandler
	GetLayout            query.GetLayoutHandler
	GetUsers             query.GetUsersHandler
	GetEventReservations query.GetEventReservationsHandler
	GetEventTickets      query.GetEventTicketsHandler
}

func NewApplication(s *ent.Client) *Application {
	return &Application{
		Commands: Commands{
			CreateLayout:      command.NewCreateLayoutHandler(s),
			CreateReservation: command.NewCreateReservationHandler(s),
			ReserveSeats:      command.NewReserveSeatsHandler(s),
			ResetDatabase:     command.NewResetDatabaseHandler(s),
		},
		Queries: Queries{
			GetEvent:             query.NewGetEventHandler(s),
			GetEvents:            query.NewGetEventsHandler(s),
			CheckEventSeats:      query.NewCheckEventSeatsHandler(s),
			GetEventTickets:      query.NewGetEventTicketsHandler(s),
			GetEventReservations: query.NewGetEventReservationsHandler(s),
			GetUserTickets:       query.NewGetUserTicketsHandler(s),
			GetUserTicket:        query.NewGetUserTicketHandler(s),
			GetLayouts:           query.NewGetLayoutsHandler(s),
			GetLayout:            query.NewGetLayoutHandler(s),
			GetUsers:             query.NewGetUsersHandler(s),
		},
	}
}
