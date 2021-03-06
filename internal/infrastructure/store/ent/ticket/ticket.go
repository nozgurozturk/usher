// Code generated by entc, DO NOT EDIT.

package ticket

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ticket type in the database.
	Label = "ticket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeSeat holds the string denoting the seat edge name in mutations.
	EdgeSeat = "seat"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the ticket in the database.
	Table = "tickets"
	// SeatTable is the table that holds the seat relation/edge.
	SeatTable = "tickets"
	// SeatInverseTable is the table name for the Seat entity.
	// It exists in this package in order to avoid circular dependency with the "seat" package.
	SeatInverseTable = "seats"
	// SeatColumn is the table column denoting the seat relation/edge.
	SeatColumn = "seat_tickets"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "tickets"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_tickets"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "tickets"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_tickets"
)

// Columns holds all SQL columns for ticket fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tickets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_tickets",
	"seat_tickets",
	"user_tickets",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
