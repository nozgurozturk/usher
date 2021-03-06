// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/user"
)

// Reservation is the model entity for the Reservation schema.
type Reservation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Rank holds the value of the "rank" field.
	Rank int `json:"rank,omitempty"`
	// Preference holds the value of the "preference" field.
	Preference int `json:"preference,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReservationQuery when eager-loading is set.
	Edges              ReservationEdges `json:"edges"`
	event_reservations *uuid.UUID
	user_reservations  *uuid.UUID
}

// ReservationEdges holds the relations/edges for other nodes in the graph.
type ReservationEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// The edge event was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReservationEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reservation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case reservation.FieldIsActive:
			values[i] = new(sql.NullBool)
		case reservation.FieldSize, reservation.FieldRank, reservation.FieldPreference:
			values[i] = new(sql.NullInt64)
		case reservation.FieldID:
			values[i] = new(uuid.UUID)
		case reservation.ForeignKeys[0]: // event_reservations
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case reservation.ForeignKeys[1]: // user_reservations
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Reservation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reservation fields.
func (r *Reservation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reservation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case reservation.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				r.Size = int(value.Int64)
			}
		case reservation.FieldRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank", values[i])
			} else if value.Valid {
				r.Rank = int(value.Int64)
			}
		case reservation.FieldPreference:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field preference", values[i])
			} else if value.Valid {
				r.Preference = int(value.Int64)
			}
		case reservation.FieldIsActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_active", values[i])
			} else if value.Valid {
				r.IsActive = value.Bool
			}
		case reservation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field event_reservations", values[i])
			} else if value.Valid {
				r.event_reservations = new(uuid.UUID)
				*r.event_reservations = *value.S.(*uuid.UUID)
			}
		case reservation.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_reservations", values[i])
			} else if value.Valid {
				r.user_reservations = new(uuid.UUID)
				*r.user_reservations = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the Reservation entity.
func (r *Reservation) QueryEvent() *EventQuery {
	return (&ReservationClient{config: r.config}).QueryEvent(r)
}

// QueryUser queries the "user" edge of the Reservation entity.
func (r *Reservation) QueryUser() *UserQuery {
	return (&ReservationClient{config: r.config}).QueryUser(r)
}

// Update returns a builder for updating this Reservation.
// Note that you need to call Reservation.Unwrap() before calling this method if this Reservation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reservation) Update() *ReservationUpdateOne {
	return (&ReservationClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Reservation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reservation) Unwrap() *Reservation {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reservation is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reservation) String() string {
	var builder strings.Builder
	builder.WriteString("Reservation(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", r.Size))
	builder.WriteString(", rank=")
	builder.WriteString(fmt.Sprintf("%v", r.Rank))
	builder.WriteString(", preference=")
	builder.WriteString(fmt.Sprintf("%v", r.Preference))
	builder.WriteString(", is_active=")
	builder.WriteString(fmt.Sprintf("%v", r.IsActive))
	builder.WriteByte(')')
	return builder.String()
}

// Reservations is a parsable slice of Reservation.
type Reservations []*Reservation

func (r Reservations) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
