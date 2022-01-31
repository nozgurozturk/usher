package seeder

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

//
type Event struct {
	Name        string
	Description string
	SeatMap     string
	LayoutID    uuid.UUID
}

func CreateEvents(client *ent.Client, events ...Event) []*ent.Event {
	eventCreates := make([]*ent.EventCreate, len(events))

	for i, e := range events {
		eventCreates[i] = client.Event.Create().
			SetName(e.Name).
			SetDescription(e.Description).
			SetSeatMap(e.SeatMap).
			SetLayoutID(e.LayoutID)
	}

	return client.Event.CreateBulk(eventCreates...).SaveX(context.Background())
}
