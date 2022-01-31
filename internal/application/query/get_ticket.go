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
)

type GetUserTicketQuery struct {
	TicketID string
}

type GetUserTicketHandler struct {
	store *ent.Client
}

func NewGetUserTicketHandler(s *ent.Client) GetUserTicketHandler {
	return GetUserTicketHandler{
		store: s,
	}
}

func (h GetUserTicketHandler) Handle(q GetUserTicketQuery) (*ticket.Ticket, error) {
	ctx := context.Background()

	ticketID, err := uuid.Parse(q.TicketID)
	if err != nil {
		return nil, err
	}

	t, err := h.store.Ticket.Query().
		Where(ticketDAO.ID(ticketID)).
		WithEvent().
		WithSeat().
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	s := t.Edges.Seat
	e := t.Edges.Event
	u := t.Edges.User
	return &ticket.Ticket{
		ID:    t.ID.String(),
		Event: store.ToEventEntity(e),
		Seat:  layout.NewSeatBuilder().WithPosition(s.Row, s.Col).WithNumber(s.Number).WithRank(s.Rank).Build(),
		User: user.User{
			ID:   u.ID.String(),
			Name: u.Name,
		},
	}, nil
}
