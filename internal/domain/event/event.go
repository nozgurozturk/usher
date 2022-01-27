package event

import (
	"time"

	"github.com/nozgurozturk/usher/internal/domain/layout"
)

type Event interface {
	// Name returns the name of the event.
	Name() string
	// Description returns the description of the event.
	Description() string
	// StartTime returns the start time of the event.
	StartAt() time.Time
	// EndTime returns the end time of the event.
	EndAt() time.Time
}

type EventBuilder interface {
	// WithName sets the name of the event.
	WithName(name string) EventBuilder
	// WithDescription sets the description of the event.
	WithDescription(description string) EventBuilder
	// WithStartTime sets the start time of the event.
	WithStartTime(startTime time.Time) EventBuilder
	// WithEndTime sets the end time of the event.
	WithEndTime(endTime time.Time) EventBuilder
	// Build returns the event.
	Build() Event
}

// NewEventBuilder returns a new EventBuilder.
func NewEventBuilder() EventBuilder {
	return &eventBuilder{}
}

// eventBuilder is a builder for an event.
type eventBuilder struct {
	name        string
	description string
	hall        layout.Hall
	startTime   time.Time
	endTime     time.Time
}

// WithName sets the name of the event.
func (b *eventBuilder) WithName(name string) EventBuilder {
	b.name = name
	return b
}

// WithDescription sets the description of the event.
func (b *eventBuilder) WithDescription(description string) EventBuilder {
	b.description = description
	return b
}

// WithHall sets the location of the event.
func (b *eventBuilder) WithHall(hall layout.Hall) EventBuilder {
	b.hall = hall
	return b
}

// WithStartTime sets the start time of the event.
func (b *eventBuilder) WithStartTime(startTime time.Time) EventBuilder {
	b.startTime = startTime
	return b
}

// WithEndTime sets the end time of the event.
func (b *eventBuilder) WithEndTime(endTime time.Time) EventBuilder {
	b.endTime = endTime
	return b
}

// FromEvent returns a new EventBuilder from an event.
func (b *eventBuilder) From(event Event) EventBuilder {
	return &eventBuilder{
		name:        event.Name(),
		description: event.Description(),
		startTime:   event.StartAt(),
		endTime:     event.EndAt(),
	}
}

// Build returns the event.
func (b *eventBuilder) Build() Event {
	return &event{
		name:        b.name,
		description: b.description,
		startTime:   b.startTime,
		endTime:     b.endTime,
	}
}

type event struct {
	name        string
	description string
	startTime   time.Time
	endTime     time.Time
}

func (e *event) Name() string {
	return e.name
}

func (e *event) Description() string {
	return e.description
}

func (e *event) StartAt() time.Time {
	return e.startTime
}

func (e *event) EndAt() time.Time {
	return e.endTime
}
