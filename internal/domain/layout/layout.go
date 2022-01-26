package layout

import (
	"strconv"
)

// Layout is a layout of seats.
type Layout struct {
	sections []*Section
}

// NewLayout creates a new layout.
func NewLayout() *Layout {
	return &Layout{}
}

// Sections is a section of seats.
func (l *Layout) Sections() []*Section {
	return l.sections
}

// AddSection adds a section to the layout.
func (l *Layout) AddSection(section ...*Section) {
	l.sections = append(l.sections, section...)
}

// String returns the string representation of the layout.
func (l *Layout) String() string {
	var str string
	for _, s := range l.Sections() {
		str += s.String()
	}
	str += "\n"
	return str
}

func (s *Section) String() string {
	str := "S-" + s.Name() + "\n"
	for _, row := range s.Rows() {
		str += row.String() + "\n"
	}
	return str
}

// ConsecutiveAvailableSeatsByRank returns the available consecutive seats by rank in a row.
// The seats are not sorted by number and rank.
// For consecutiveness the seats must be in the same rank and must follow previous order.
// Example:
//	[Seat1, Seat2, Seat3, Seat4, Seat5, Seat6, Seat7]
// 	Seat 2,3,5,6,7 are available
// 	Seat 2,3,5,6 ranks are the same
//	So, the result is [[Seat2, Seat3], [Seat5, Seat6]]
func (l *Layout) ConsecutiveAvailableSeatsByRank(rank int) [][]*Seat {
	availableSeatsByRank := make([][]*Seat, 0, len(l.sections))

	for _, section := range l.sections {
		availableSeatsByRank = append(availableSeatsByRank, section.consecutiveAvailableSeatsByRank(rank)...)
	}

	return availableSeatsByRank
}

// AvailableSeatsByRank returns the available seats by rank in a row.
func (l *Layout) AvailableSeatsByRank(rank int) []*Seat {
	availableSeatsByRank := make([]*Seat, 0, len(l.sections))

	for _, section := range l.sections {
		availableSeatsByRank = append(availableSeatsByRank, section.availableSeatsByRank(rank)...)
	}

	return availableSeatsByRank
}

// Section is a section of seats.
type Section struct {
	name string
	rows []*Row
}

// NewSection creates a new section.
func NewSection() *Section {
	return &Section{}
}

// Rows is a row of seats.
func (s *Section) Rows() []*Row {
	return s.rows
}

// Name returns the name of the section.
func (s *Section) Name() string {
	return s.name
}

// SetName sets the name of the section.
func (s *Section) SetName(name string) {
	s.name = name
}

// AddRow adds a row to the section.
func (s *Section) AddRow(row ...*Row) {
	for i, r := range row {
		rowName := s.Name() + strconv.Itoa(i)
		r.SetOrder(i)
		r.SetName(rowName)
		for j, seat := range r.Seats() {
			// set position of the seats
			seat.SetPosition(i, j)
		}
	}

	s.rows = append(s.rows, row...)
}

// availableSeatsByRank returns the available seats by rank in a section.
func (s *Section) availableSeatsByRank(rank int) []*Seat {
	availableSeatsByRank := make([]*Seat, 0, len(s.rows))

	for _, row := range s.rows {
		availableSeatsByRank = append(availableSeatsByRank, row.availableSeatsByRank(rank)...)
	}

	return availableSeatsByRank
}

// consecutiveAvailableSeatsByRank returns the available consecutive seats by rank in a section.
func (s *Section) consecutiveAvailableSeatsByRank(rank int) [][]*Seat {
	availableSeatsByRank := make([][]*Seat, 0, len(s.rows))

	for _, row := range s.rows {
		availableSeatsByRank = append(availableSeatsByRank, row.consecutiveAvailableSeatsByRank(rank)...)
	}

	return availableSeatsByRank
}

// Row is a row of seats.
type Row struct {
	name  string
	order int
	seats []*Seat
}

// NewRow creates a new row.
func NewRow() *Row {
	return &Row{}
}

// Name returns the name of the row.
func (r *Row) Name() string {
	return r.name
}

// Order returns the order of the row.
func (r *Row) Order() int {
	return r.order
}

// Seats is a row of seats.
func (r *Row) Seats() []*Seat {
	return r.seats
}

// String returns the string representation of the row.
func (r *Row) String() string {
	var str string
	for _, seat := range r.Seats() {
		str += seat.String()
	}
	return str
}

// SetName sets the name of the row.
func (r *Row) SetName(name string) *Row {
	r.name = name
	return r
}

// SetOrder sets the order of the row.
func (r *Row) SetOrder(order int) *Row {
	r.order = order
	return r
}

// AddSeat adds a seat to the row.
func (r *Row) AddSeat(seat ...*Seat) {
	r.seats = append(r.seats, seat...)
}

// availableSeatsByRank returns the available seats by rank in a row.
func (r *Row) availableSeatsByRank(rank int) []*Seat {
	availableSeatsByRank := make([]*Seat, 0, len(r.seats))

	for _, seat := range r.seats {
		if seat.Rank() == rank && seat.Available() {
			availableSeatsByRank = append(availableSeatsByRank, seat)
		}
	}

	return availableSeatsByRank
}

// ConsecutiveAvailableSeatsByRank returns the available consecutive seats by rank in a row.
func (r *Row) consecutiveAvailableSeatsByRank(rank int) [][]*Seat {

	availableSeatsByRank := make([][]*Seat, 0, len(r.Seats()))

	for i, seat := range r.Seats() {
		if seat.Rank() != rank || !seat.Available() {
			continue
		}

		if i == 0 {
			availableSeatsByRank = append(availableSeatsByRank, []*Seat{seat})
			continue
		}

		prevSeat := r.Seats()[i-1]
		if prevSeat.Rank() == rank && prevSeat.Available() && prevSeat.Order()+1 == seat.Order() {
			index := len(availableSeatsByRank) - 1
			availableSeatsByRank[index] = append(availableSeatsByRank[index], seat)
		} else {
			availableSeatsByRank = append(availableSeatsByRank, []*Seat{seat})
		}
	}

	return availableSeatsByRank
}

// Seat is a seat.
type Seat struct {
	row       int
	col       int
	rank      int
	order     int
	number    int
	available bool
}

// NewSeat creates a new seat.
func NewSeat(rank, order, number int) *Seat {
	return &Seat{
		row:       0,
		col:       0,
		rank:      rank,
		order:     order,
		number:    number,
		available: true,
	}
}

// Position returns the position of the seat.
func (s *Seat) Position() (int, int) {
	return s.row, s.col
}

// SetPosition sets the position of the seat.
func (s *Seat) SetPosition(row, col int) *Seat {
	s.row = row
	s.col = col

	return s
}

// Rank returns the rank of the seat.
func (s *Seat) Rank() int {
	return s.rank
}

// Number returns the number of the seat.
func (s *Seat) Number() int {
	return s.number
}

func (s *Seat) Order() int {
	return s.order
}

func (s *Seat) Available() bool {
	return s.available
}

func (s *Seat) Book() *Seat {
	s.available = false
	return s
}

func (s *Seat) Dismiss() *Seat {
	s.available = true
	return s
}

func (s *Seat) String() string {
	if !s.Available() {
		return "■"
	} else {
		return "□"
	}
}
