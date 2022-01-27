package layout

// Hall interface
type Hall interface {
	// Name returns the name of the venue.
	Name() string
	// Sections returns the sections of the venue.
	Sections() []Section
	// Copy returns a copy of the hall.
	Copy() Hall

	Stringer
}

// Section interface
type Section interface {
	// Name returns the name of the section.
	Name() string
	// Rows returns the rows of the section.
	Rows() []Row
	// Copy returns a copy of the section.
	Copy() Section

	Stringer
}

// Row interface
type Row interface {
	// Name returns the name of the row.
	Name() string
	// Order returns the order of the row.
	Order() int
	// Seats returns the seats of the row.
	Seats() []Seat
	// Copy returns a copy of the row.
	Copy() Row

	Stringer
}

// Seat interface
type Seat interface {
	// Position returns the position of the seat.
	Position() (int, int)
	// Rank returns the rank of the seat.
	Rank() int
	// Number returns the number of the seat.
	Number() int
	// Order returns the order of the seat.
	Order() int
	// Available returns true if the seat is available.
	Available() bool
	// Book books the seat.
	Book()
	// Dismiss dismisses the seat.
	Dismiss()
	// Feature returns the feature of the seat.
	Feature() SeatFeature
	// Type returns the type of the seat.
	HasFeature(feature SeatFeature) bool
	// Copy returns a copy of the seat.
	Copy() Seat

	Stringer
}

// Stringer interface for string representation
type Stringer interface {
	// String returns a string representation of the seat.
	String() string
}

// Filter interface for filtering seats
type Filter interface {
	// WithFeature sets the feature of the seat.
	WithFeature(feature SeatFeature) Filter
	// WithRank sets the rank of the seat.
	WithRank(rank int) Filter
	// WithAvailable sets the available of the seat.
	WithAvailable(available bool) Filter
	// FilterSeats filters the seats.
	FilterSeat(seat Seat) Seat
}

type filter struct {
	feature   SeatFeature
	rank      *int
	available *bool
}

// NewFilter creates a new filter.
func NewFilter() Filter {
	return &filter{}
}

// WithFeature sets the feature of the seat.
func (f *filter) WithFeature(feature SeatFeature) Filter {
	f.feature |= feature
	return f
}

// WithRank sets the rank of the seat.
func (f *filter) WithRank(rank int) Filter {
	f.rank = &rank
	return f
}

// WithAvailable sets the available of the seat.
func (f *filter) WithAvailable(available bool) Filter {
	f.available = &available
	return f
}

// FilterSeats filters the seats.
func (f *filter) FilterSeat(seat Seat) Seat {

	if f.feature > SeatFeatureDefault && !seat.HasFeature(f.feature) {
		return nil
	}

	if f.rank != nil && seat.Rank() != *f.rank {
		return nil
	}

	if f.available != nil && !seat.Available() {
		return nil
	}

	return seat
}
