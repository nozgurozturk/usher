package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/ticket"
	"github.com/nozgurozturk/usher/internal/domain/user"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	ticketDAO "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/ticket"
	userDAO "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/user"
)

type GetUserTicketsQuery struct {
	UserID string
}

type GetUserTicketsHandler struct {
	store *ent.Client
}

func NewGetUserTicketsHandler(s *ent.Client) GetUserTicketsHandler {
	return GetUserTicketsHandler{
		store: s,
	}
}

func (h GetUserTicketsHandler) Handle(q GetUserTicketsQuery) ([]ticket.Ticket, error) {
	ctx := context.Background()

	userID, err := uuid.Parse(q.UserID)
	if err != nil {
		return nil, err
	}

	ticketDAOs, err := h.store.Ticket.Query().
		Where(ticketDAO.HasUserWith(userDAO.ID(userID))).
		WithEvent(func(eq *ent.EventQuery) {
			eq.WithLayout()
		}).
		WithSeat().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, err
	}

	tickets := make([]ticket.Ticket, len(ticketDAOs))
	for i := range ticketDAOs {
		seatDAO := ticketDAOs[i].Edges.Seat
		tickets[i] = ticket.Ticket{
			User: user.User{
				ID:   ticketDAOs[i].Edges.User.ID.String(),
				Name: ticketDAOs[i].Edges.User.Name,
			},
			ID:    ticketDAOs[i].ID.String(),
			Event: store.ToEventEntity(ticketDAOs[i].Edges.Event),
			Seat:  layout.NewSeatBuilder().WithPosition(seatDAO.Row, seatDAO.Col).WithNumber(seatDAO.Number).WithRank(seatDAO.Rank).Build(),
		}
	}

	return tickets, nil
}
