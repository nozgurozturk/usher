package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/group"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
)

type GetEventReservationsQuery struct {
	EventID string
}

type GetEventReservationsHandler struct {
	store *ent.Client
}

func NewGetEventReservationsHandler(s *ent.Client) GetEventReservationsHandler {
	return GetEventReservationsHandler{
		store: s,
	}
}

func (h GetEventReservationsHandler) Handle(q GetEventReservationsQuery) ([]group.Group, error) {
	ctx := context.Background()

	eventID, err := uuid.Parse(q.EventID)
	if err != nil {
		return nil, err
	}

	reservationDAOs, err := h.store.Reservation.Query().
		QueryEvent().
		Where(dao.ID(eventID)).
		QueryReservations().
		WithEvent().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, err
	}

	groups := make([]group.Group, len(reservationDAOs))
	for i, r := range reservationDAOs {
		groups[i] = group.NewGroup(r.ID.String(), r.Size, r.Rank, group.SeatPreference(r.Preference))
	}

	return groups, nil
}
