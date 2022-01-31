package event

import (
	"time"

	"github.com/nozgurozturk/usher/internal/domain/layout"
)

// Event is an event.
type Event interface {
	ID() string
	// Name returns the name of the event.
	Name() string
	// Description returns the description of the event.
	Description() string
	// StartTime returns the start time of the event.
	StartAt() *time.Time
	// EndTime returns the end time of the event.
	EndAt() *time.Time
	// Hall returns the ID of the layout of the event.
	Hall() layout.Hall
	// SeatMap returns the seat map of the event.
	SeatMap() layout.SeatMap
	// ReserveSeatsWithSeatMap reserves seats with the seat map of the event.
	ReservedSeatsWithSeatMap()
}

// EventBuilder is an event builder.
type EventBuilder interface {
	// WithName sets the name of the event.
	WithID(id string) EventBuilder
	// WithName sets the name of the event.
	WithName(name string) EventBuilder
	// WithDescription sets the description of the event.
	WithDescription(description string) EventBuilder
	// WithStartDate sets the start time of the event.
	WithStartDate(startTime *time.Time) EventBuilder
	// WithEndDate sets the end time of the event.
	WithEndDate(endTime *time.Time) EventBuilder
	// WithHall sets the ID of the layout of the event.
	WithHall(hall layout.Hall) EventBuilder
	// WithSeatMap sets the available seat map of the event.
	WithSeatMap(seatMap layout.SeatMap) EventBuilder
	// From returns a new EventBuilder from an event.
	From(event Event) EventBuilder
	// Build returns the event.
	Build() Event
}

// NewEventBuilder returns a new EventBuilder.
func NewEventBuilder() EventBuilder {
	return &eventBuilder{}
}

// eventBuilder is a builder for an event.
type eventBuilder struct {
	id          string
	name        string
	description string
	hall        layout.Hall
	seatMap     layout.SeatMap
	startTime   *time.Time
	endTime     *time.Time
}

func (b *eventBuilder) WithID(id string) EventBuilder {
	b.id = id
	return b
}

func (b *eventBuilder) WithName(name string) EventBuilder {
	b.name = name
	return b
}

func (b *eventBuilder) WithDescription(description string) EventBuilder {
	b.description = description
	return b
}

func (b *eventBuilder) WithStartDate(startTime *time.Time) EventBuilder {
	b.startTime = startTime
	return b
}

func (b *eventBuilder) WithEndDate(endTime *time.Time) EventBuilder {
	b.endTime = endTime
	return b
}

func (b *eventBuilder) WithSeatMap(seatMap layout.SeatMap) EventBuilder {
	b.seatMap = seatMap
	return b
}

func (b *eventBuilder) WithHall(hall layout.Hall) EventBuilder {
	b.hall = hall
	return b
}

func (b *eventBuilder) From(event Event) EventBuilder {
	return &eventBuilder{
		id:          event.ID(),
		name:        event.Name(),
		description: event.Description(),
		hall:        event.Hall(),
		seatMap:     event.SeatMap(),
		startTime:   event.StartAt(),
		endTime:     event.EndAt(),
	}
}

// Build returns the event.
func (b *eventBuilder) Build() Event {
	return &event{
		id:          b.id,
		name:        b.name,
		description: b.description,
		hall:        b.hall,
		seatMap:     b.seatMap,
		startTime:   b.startTime,
		endTime:     b.endTime,
	}
}

type event struct {
	id          string
	name        string
	description string
	hall        layout.Hall
	seatMap     layout.SeatMap
	startTime   *time.Time
	endTime     *time.Time
}

func (e *event) ID() string {
	return e.id
}

func (e *event) Name() string {
	return e.name
}

func (e *event) Description() string {
	return e.description
}

func (e *event) StartAt() *time.Time {
	return e.startTime
}

func (e *event) EndAt() *time.Time {
	return e.endTime
}

func (e *event) Hall() layout.Hall {
	return e.hall
}

func (e *event) SeatMap() layout.SeatMap {
	return e.seatMap
}

func (e *event) ReservedSeatsWithSeatMap() {
	for i, sectionMap := range e.SeatMap() {
		for j, rowMap := range sectionMap {
			for k, seat := range rowMap {
				if seat == 1 {
					e.Hall().Sections()[i].Rows()[j].Seats()[k].Book()
				}
			}
		}
	}
}
