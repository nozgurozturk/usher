package layout

import "strconv"

type SeatFeature byte

const (
	SeatFeatureDefault SeatFeature = 1 << iota
	SeatFeatureAisle
	SeatFeatureHigh
	SeatFeatureFront
)

type SeatBuilder interface {
	// WithPosition sets the position of the seat.
	WithPosition(row, col int) SeatBuilder
	// WithRank sets the rank of the seat.
	WithRank(rank int) SeatBuilder
	// WithNumber sets the number of the seat.
	WithNumber(number int) SeatBuilder
	// WithFeature sets the feature of the seat.
	WithFeature(feature SeatFeature) SeatBuilder
	// Build returns the seat.
	Build() Seat
	// From returns a new SeatBuilder from the seat.
	From(seat Seat) SeatBuilder
}

type seatBuilder struct {
	row     int
	col     int
	rank    int
	number  int
	feature SeatFeature
}

// NewSeatBuilder returns a new SeatBuilder.
func NewSeatBuilder() SeatBuilder {
	return &seatBuilder{
		feature: SeatFeatureDefault,
	}
}

func (b *seatBuilder) WithPosition(row, col int) SeatBuilder {
	b.row = row
	b.col = col
	return b
}

func (b *seatBuilder) WithRank(rank int) SeatBuilder {
	b.rank = rank
	return b
}

func (b *seatBuilder) WithNumber(number int) SeatBuilder {
	b.number = number
	return b
}

func (b *seatBuilder) WithFeature(feature SeatFeature) SeatBuilder {
	b.feature |= feature
	return b
}

func (b *seatBuilder) From(seat Seat) SeatBuilder {
	row, col := seat.Position()
	return &seatBuilder{
		row:     row,
		col:     col,
		rank:    seat.Rank(),
		number:  seat.Number(),
		feature: seat.Feature(),
	}
}

func (b *seatBuilder) Build() Seat {
	return &seat{
		row:       b.row,
		col:       b.col,
		rank:      b.rank,
		number:    b.number,
		feature:   b.feature,
		available: true,
	}
}

type seat struct {
	row       int
	col       int
	rank      int
	number    int
	available bool
	feature   SeatFeature
}

func (s *seat) Position() (int, int) {
	return s.row, s.col
}

func (s *seat) Rank() int {
	return s.rank
}

func (s *seat) Number() int {
	return s.number
}

func (s *seat) Order() int {
	return s.col
}

func (s *seat) Available() bool {
	return s.available
}

func (s *seat) Book() {
	s.available = false
}

func (s *seat) Dismiss() {
	s.available = true
}

func (s *seat) HasFeature(feature SeatFeature) bool {
	return s.feature&feature != 0
}

func (s *seat) Feature() SeatFeature {
	return s.feature
}

const (
	// Red color
	red = "\033[31m"
	// Green color
	green = "\033[32m"
	// Reset color
	reset = "\033[0m"
)

func (s *seat) String() string {
	rep := strconv.Itoa(s.Rank())
	if !s.Available() {
		rep = "X"
	}
	return "[" + rep + "]"
}

func (s *seat) Copy() Seat {
	return &seat{
		row:       s.row,
		col:       s.col,
		rank:      s.rank,
		number:    s.number,
		available: s.available,
		feature:   s.feature,
	}
}
