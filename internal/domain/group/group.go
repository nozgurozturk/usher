package group

import (
	"fmt"
	"math/rand"
	"strconv"
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
	// String returns a string representation of the group.
	String() string
}

type group struct {
	id   string
	size int
	rank int
	SeatPreference
}

// NewGroup returns a new group.
func NewGroup(size int, rank int, preferences ...SeatPreference) Group {
	preference := PreferenceDefault
	for _, p := range preferences {
		preference |= p
	}

	return &group{
		id:             strconv.FormatInt(rand.Int63(), 16),
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

func (g *group) String() string {
	return fmt.Sprintf("ID:%s S:%d R:%d P:%d", g.ID(), g.Size(), g.RankPreference(), int(g.SeatPreferences()))
}
