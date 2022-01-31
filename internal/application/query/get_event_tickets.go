package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/ticket"
	"github.com/nozgurozturk/usher/internal/domain/user"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
)

type GetEventTicketsQuery struct {
	EventID string
}

type GetEventTicketsHandler struct {
	store *ent.Client
}

func NewGetEventTicketsHandler(s *ent.Client) GetEventTicketsHandler {
	return GetEventTicketsHandler{
		store: s,
	}
}

func (h GetEventTicketsHandler) Handle(q GetEventTicketsQuery) ([]ticket.Ticket, error) {
	ctx := context.Background()

	eventID, err := uuid.Parse(q.EventID)
	if err != nil {
		return nil, err
	}

	ticketDAOs, err := h.store.Ticket.Query().
		QueryEvent().
		Where(dao.ID(eventID)).
		QueryTickets().
		WithEvent().
		WithUser().
		WithSeat().
		All(ctx)
	if err != nil {
		return nil, err
	}

	tickets := make([]ticket.Ticket, len(ticketDAOs))
	for i := range ticketDAOs {
		seatDAO := ticketDAOs[i].Edges.Seat
		userDAO := ticketDAOs[i].Edges.User
		tickets[i] = ticket.Ticket{
			ID: ticketDAOs[i].ID.String(),
			User: user.User{
				ID:   userDAO.ID.String(),
				Name: userDAO.Name,
			},
			// Event: store.ToEventEntity(ticketDAOs[i].Edges.Event),
			Seat: layout.NewSeatBuilder().WithPosition(seatDAO.Row, seatDAO.Col).WithNumber(seatDAO.Number).WithRank(seatDAO.Rank).Build(),
		}
	}

	return tickets, nil
}
