package store

import (
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/event"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

func ToEventEntity(e *ent.Event) event.Event {
	b := event.NewEventBuilder().
		WithID(e.ID.String()).
		WithName(e.Name).
		WithDescription(e.Description).
		WithStartDate(&e.StartAt).
		WithEndDate(&e.EndAt).
		WithSeatMap(layout.SeatMap{}.FromString(e.SeatMap))

	if e.Edges.Layout != nil {
		b.WithHall(ToLayoutEntity(e.Edges.Layout))
	}

	return b.Build()
}

func ToEventDAO(d event.Event) *ent.Event {
	eventID, err := uuid.Parse(d.ID())
	if err != nil {
		return nil
	}
	return &ent.Event{
		ID:          eventID,
		Name:        d.Name(),
		Description: d.Description(),
		StartAt:     d.StartAt().UTC(),
		EndAt:       d.EndAt().UTC(),
		SeatMap:     d.SeatMap().String(),
	}
}
