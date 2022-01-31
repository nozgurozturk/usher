package layout

// Hall interface
type Hall interface {
	// ID returns the ID of the hall.
	ID() string
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
	// ID returns the ID of the seat.
	ID() string
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
	// JSON returns a JSON representation of the seat.
	JSON() interface{}
}
