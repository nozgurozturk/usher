package layout

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
func (l *Layout) AddSection(section *Section) {
	l.sections = append(l.sections, section)
}

// Section is a section of seats.
type Section struct {
	name string
	rows []Row
}

// NewSection creates a new section.
func NewSection() *Section {
	return &Section{}
}

// Rows is a row of seats.
func (s *Section) Rows() []Row {
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
func (s *Section) AddRow(row Row) {
	s.rows = append(s.rows, row)
}

// Seat is a seat.
type Seat struct {
	rank   int
	number int
}

// NewSeat creates a new seat.
func NewSeat(rank, number int) *Seat {
	return &Seat{
		rank:   rank,
		number: number,
	}
}

// Rank returns the rank of the seat.
func (s *Seat) Rank() int {
	return s.rank
}

// Number returns the number of the seat.
func (s *Seat) Number() int {
	return s.number
}

// Row is a row of seats.
type Row []*Seat

// CreateRow creates a new row with given array of integer that represents ranks and numbering style.
//
// Example
// r := layout.CreateRow([]int{1, 1, 2, 3, 2, 1, 1})
//
// fmt.Println(r)
//
// Output:
// [{1 1} {1 2} {2 3} {3 4} {2 5} {1 6} {1 7}]

func CreateRow(ranks []int) Row {
	row := make(Row, 0, len(ranks))

	for i, rank := range ranks {
		row = append(row, NewSeat(rank, i+1))
	}

	return row
}
