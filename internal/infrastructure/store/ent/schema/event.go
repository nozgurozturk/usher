package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("name"),
		field.String("description"),
		field.String("seat_map"),
		field.Time("start_at").Default(time.Now().UTC),
		field.Time("end_at").Default(time.Now().Add(time.Hour * 24 * 30).UTC),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reservations", Reservation.Type),
		edge.To("tickets", Ticket.Type),
		edge.From("layout", Layout.Type).
			Ref("events").
			Unique(),
	}
}
