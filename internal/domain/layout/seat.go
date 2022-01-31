package layout

type SeatFeature byte

const SeatFeatureDefault SeatFeature = 0
const (
	SeatFeatureAisle SeatFeature = 1 << iota
	SeatFeatureHigh
	SeatFeatureFront
)

func (s SeatFeature) Is(feature SeatFeature) bool {
	return s&feature != 0
}

type SeatBuilder interface {
	WithID(id string) SeatBuilder
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
	id      string
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

func (sb *seatBuilder) WithID(id string) SeatBuilder {
	sb.id = id
	return sb
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
		id:        b.id,
		row:       b.row,
		col:       b.col,
		rank:      b.rank,
		number:    b.number,
		feature:   b.feature,
		available: true,
	}
}

type seat struct {
	id        string
	row       int
	col       int
	rank      int
	number    int
	available bool
	feature   SeatFeature
}

func (s *seat) ID() string {
	return s.id
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

func (s *seat) String() string {
	rep := "□"
	if !s.Available() {
		rep = "■"
	}

	return rep
}

func (s *seat) JSON() interface{} {
	return struct {
		Row       int  `json:"row"`
		Col       int  `json:"col"`
		Rank      int  `json:"rank"`
		Number    int  `json:"number"`
		Available bool `json:"available"`
		Feature   int  `json:"feature"`
	}{
		Row:       s.row,
		Col:       s.col,
		Rank:      s.rank,
		Number:    s.number,
		Available: s.available,
		Feature:   int(s.Feature()),
	}
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

func FilteredSeatBlock(seatBlock []Seat, filter Filter) []Seat {
	availableSeats := make([]Seat, 0, len(seatBlock))

	for _, seat := range seatBlock {
		if filter.FilterSeat(seat) != nil {
			availableSeats = append(availableSeats, seat)
		}
	}

	return availableSeats
}
