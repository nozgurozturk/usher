package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Reservation holds the schema definition for the Reservation entity.
type Reservation struct {
	ent.Schema
}

// Fields of the Reservation.
func (Reservation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("size").Default(0),
		field.Int("rank").Default(0),
		field.Int("preference").Default(0),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the Reservation.
func (Reservation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("reservations").Required().Unique(),
		edge.From("user", User.Type).Ref("reservations").Required().Unique(),
	}
}
