package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
)

type GetEventQuery struct {
	EventID string
}

type GetEventHandler struct {
	store *ent.Client
}

func NewGetEventHandler(s *ent.Client) GetEventHandler {
	return GetEventHandler{
		store: s,
	}
}

func (h GetEventHandler) Handle(q GetEventQuery) (event.Event, error) {
	ctx := context.Background()

	eventID, err := uuid.Parse(q.EventID)
	if err != nil {
		return nil, err
	}

	eventDAO, err := h.store.Event.Query().Where(dao.ID(eventID)).WithLayout(
		func(q *ent.LayoutQuery) {
			q.WithSections(func(sq *ent.SectionQuery) {
				sq.WithRows(func(rq *ent.RowQuery) {
					rq.WithSeats()
				})
			})
		}).Only(ctx)
	if err != nil {
		return nil, err
	}

	event := store.ToEventEntity(eventDAO)
	event.ReservedSeatsWithSeatMap()

	return event, nil
}
