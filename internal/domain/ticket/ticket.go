package ticket

import (
	"github.com/nozgurozturk/usher/internal/domain/event"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/user"
)

// Ticket is a struct that represents a ticket.
type Ticket struct {
	ID    string
	Event event.Event
	Seat  layout.Seat
	User  user.User
}
