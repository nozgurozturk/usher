// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/row"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/schema"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/seat"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/section"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/ticket"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescStartAt is the schema descriptor for start_at field.
	eventDescStartAt := eventFields[4].Descriptor()
	// event.DefaultStartAt holds the default value on creation for the start_at field.
	event.DefaultStartAt = eventDescStartAt.Default.(func() time.Time)
	// eventDescEndAt is the schema descriptor for end_at field.
	eventDescEndAt := eventFields[5].Descriptor()
	// event.DefaultEndAt holds the default value on creation for the end_at field.
	event.DefaultEndAt = eventDescEndAt.Default.(func() time.Time)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() uuid.UUID)
	layoutFields := schema.Layout{}.Fields()
	_ = layoutFields
	// layoutDescID is the schema descriptor for id field.
	layoutDescID := layoutFields[0].Descriptor()
	// layout.DefaultID holds the default value on creation for the id field.
	layout.DefaultID = layoutDescID.Default.(func() uuid.UUID)
	reservationFields := schema.Reservation{}.Fields()
	_ = reservationFields
	// reservationDescSize is the schema descriptor for size field.
	reservationDescSize := reservationFields[1].Descriptor()
	// reservation.DefaultSize holds the default value on creation for the size field.
	reservation.DefaultSize = reservationDescSize.Default.(int)
	// reservationDescRank is the schema descriptor for rank field.
	reservationDescRank := reservationFields[2].Descriptor()
	// reservation.DefaultRank holds the default value on creation for the rank field.
	reservation.DefaultRank = reservationDescRank.Default.(int)
	// reservationDescPreference is the schema descriptor for preference field.
	reservationDescPreference := reservationFields[3].Descriptor()
	// reservation.DefaultPreference holds the default value on creation for the preference field.
	reservation.DefaultPreference = reservationDescPreference.Default.(int)
	// reservationDescIsActive is the schema descriptor for is_active field.
	reservationDescIsActive := reservationFields[4].Descriptor()
	// reservation.DefaultIsActive holds the default value on creation for the is_active field.
	reservation.DefaultIsActive = reservationDescIsActive.Default.(bool)
	// reservationDescID is the schema descriptor for id field.
	reservationDescID := reservationFields[0].Descriptor()
	// reservation.DefaultID holds the default value on creation for the id field.
	reservation.DefaultID = reservationDescID.Default.(func() uuid.UUID)
	rowFields := schema.Row{}.Fields()
	_ = rowFields
	// rowDescOrder is the schema descriptor for order field.
	rowDescOrder := rowFields[2].Descriptor()
	// row.DefaultOrder holds the default value on creation for the order field.
	row.DefaultOrder = rowDescOrder.Default.(int)
	// rowDescID is the schema descriptor for id field.
	rowDescID := rowFields[0].Descriptor()
	// row.DefaultID holds the default value on creation for the id field.
	row.DefaultID = rowDescID.Default.(func() uuid.UUID)
	seatFields := schema.Seat{}.Fields()
	_ = seatFields
	// seatDescRow is the schema descriptor for row field.
	seatDescRow := seatFields[2].Descriptor()
	// seat.DefaultRow holds the default value on creation for the row field.
	seat.DefaultRow = seatDescRow.Default.(int)
	// seatDescCol is the schema descriptor for col field.
	seatDescCol := seatFields[3].Descriptor()
	// seat.DefaultCol holds the default value on creation for the col field.
	seat.DefaultCol = seatDescCol.Default.(int)
	// seatDescRank is the schema descriptor for rank field.
	seatDescRank := seatFields[4].Descriptor()
	// seat.DefaultRank holds the default value on creation for the rank field.
	seat.DefaultRank = seatDescRank.Default.(int)
	// seatDescIsAvailable is the schema descriptor for is_available field.
	seatDescIsAvailable := seatFields[5].Descriptor()
	// seat.DefaultIsAvailable holds the default value on creation for the is_available field.
	seat.DefaultIsAvailable = seatDescIsAvailable.Default.(bool)
	// seatDescFeature is the schema descriptor for feature field.
	seatDescFeature := seatFields[6].Descriptor()
	// seat.DefaultFeature holds the default value on creation for the feature field.
	seat.DefaultFeature = seatDescFeature.Default.(int)
	// seatDescID is the schema descriptor for id field.
	seatDescID := seatFields[0].Descriptor()
	// seat.DefaultID holds the default value on creation for the id field.
	seat.DefaultID = seatDescID.Default.(func() uuid.UUID)
	sectionFields := schema.Section{}.Fields()
	_ = sectionFields
	// sectionDescID is the schema descriptor for id field.
	sectionDescID := sectionFields[0].Descriptor()
	// section.DefaultID holds the default value on creation for the id field.
	section.DefaultID = sectionDescID.Default.(func() uuid.UUID)
	ticketFields := schema.Ticket{}.Fields()
	_ = ticketFields
	// ticketDescID is the schema descriptor for id field.
	ticketDescID := ticketFields[0].Descriptor()
	// ticket.DefaultID holds the default value on creation for the id field.
	ticket.DefaultID = ticketDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
