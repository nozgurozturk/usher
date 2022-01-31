package query

import (
	"context"

	"github.com/nozgurozturk/usher/internal/domain/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type GetEventsQuery struct{}

type GetEventsHandler struct {
	store *ent.Client
}

func NewGetEventsHandler(s *ent.Client) GetEventsHandler {
	return GetEventsHandler{
		store: s,
	}
}

func (h GetEventsHandler) Handle(q GetEventsQuery) ([]event.Event, error) {
	ctx := context.Background()

	eventDAOs, err := h.store.Event.Query().WithLayout(
		func(lq *ent.LayoutQuery) {
			lq.Select("id")
		},
	).All(ctx)
	if err != nil {
		return nil, err
	}

	events := make([]event.Event, len(eventDAOs))
	for i := range eventDAOs {
		events[i] = store.ToEventEntity(eventDAOs[i])
	}

	return events, nil
}
