package query

import (
	"context"

	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/ticket"
	"github.com/nozgurozturk/usher/internal/domain/user"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type GetAllTicketsQuery struct{}

type GetAllTicketsHandler struct {
	store *ent.Client
}

func NewGetAllTicketsHandler(s *ent.Client) GetAllTicketsHandler {
	return GetAllTicketsHandler{
		store: s,
	}
}

func (h GetAllTicketsHandler) Handle(q GetAllTicketsQuery) ([]ticket.Ticket, error) {
	ctx := context.Background()

	ticketDAOs, err := h.store.Ticket.Query().
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
			Event: store.ToEventEntity(ticketDAOs[i].Edges.Event),
			Seat:  layout.NewSeatBuilder().WithPosition(seatDAO.Row, seatDAO.Col).WithNumber(seatDAO.Number).WithRank(seatDAO.Rank).Build(),
		}
	}

	return tickets, nil
}
