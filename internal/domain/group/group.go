package group

import (
	"errors"
	"fmt"

	"github.com/nozgurozturk/usher/internal/domain/layout"
)

type SeatPreference byte

const PreferenceDefault SeatPreference = 0
const (
	PreferenceAisle SeatPreference = 1 << iota
	PreferenceHigh
	PreferenceFront
)

// Group represents a group of users.
type Group interface {
	ID() string
	// Size returns the size of the group.
	Size() int
	// SeatPreferences returns the preference of the group.
	SeatPreferences() SeatPreference
	// RankPreferences returns the preference of the group.
	RankPreference() int
	// Seats returns the seats of the group.
	Seats() []layout.Seat
	// AllocateSeats allocates seats for the group.
	AllocateSeats(s ...layout.Seat) error
	// IsSatisfied returns true if the group is satisfied.
	IsSatisfied() bool
	// String returns the string representation of the group.
	String() string
}

type group struct {
	id    string
	size  int
	rank  int
	seats []layout.Seat
	SeatPreference
}

// NewGroup returns a new group.
func NewGroup(id string, size int, rank int, preferences ...SeatPreference) Group {
	preference := PreferenceDefault
	for _, p := range preferences {
		preference |= p
	}

	return &group{
		id:             id,
		size:           size,
		rank:           rank,
		SeatPreference: preference,
	}
}

// Create unique identifier for the group with hash 3 int.
func (g *group) ID() string {
	return g.id
}

func (g *group) Size() int {
	return g.size
}

func (g *group) SeatPreferences() SeatPreference {
	return g.SeatPreference
}

func (g *group) RankPreference() int {
	return g.rank
}

func (g *group) Seats() []layout.Seat {
	return g.seats
}

func (g *group) AllocateSeats(s ...layout.Seat) error {
	for _, seat := range s {
		if !seat.Available() {
			return errors.New("seat is not available")
		}
		g.seats = append(g.seats, seat)
	}
	return nil
}

func (g *group) IsSatisfied() bool {
	return len(g.seats) == g.Size()
}

func (g *group) String() string {
	return fmt.Sprintf("ID:%s S:%d R:%d P:%d", g.ID(), g.Size(), g.RankPreference(), int(g.SeatPreferences()))
}
