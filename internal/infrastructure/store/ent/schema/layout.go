package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Layout holds the schema definition for the Layout entity.
type Layout struct {
	ent.Schema
}

// Fields of the Layout.
func (Layout) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
		field.String("name"),
	}
}

// Edges of the Layout.
func (Layout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type),
		edge.To("sections", Section.Type),
	}
}

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("layout", Layout.Type).Ref("sections").Unique(),
		edge.To("rows", Row.Type),
	}
}

// Row holds the schema definition for the Row entity.
type Row struct {
	ent.Schema
}

// Fields of the Row
func (Row) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.Int("order").
			Default(0),
	}
}

// Edges of the Row.
func (Row) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("section", Section.Type).Ref("rows").Unique(),
		edge.To("seats", Seat.Type),
	}
}

// Seat holds the schema definition for the Seat entity.
type Seat struct {
	ent.Schema
}

// Fields of the Seat.
func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int("number"),
		field.Int("row").
			Default(0),
		field.Int("col").
			Default(0),
		field.Int("rank").
			Default(0),
		field.Bool("is_available").
			Default(true),
		field.Int("feature").
			Default(0),
	}
}

// Edges of the Seat.
func (Seat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type),
		edge.From("rows", Row.Type).Ref("seats").Unique(),
	}
}
