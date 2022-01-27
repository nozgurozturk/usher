package layout

// RowBuilder is a builder for a row.
type RowBuilder interface {
	// WithName sets the name of the row.
	WithName(name string) RowBuilder
	// WithOrder sets the order of the row.
	WithOrder(order int) RowBuilder
	// WithSeat adds a seat to the row.
	WithSeat(seat ...Seat) RowBuilder
	// Build returns the row.
	Build() Row
	// FromRow returns a new RowBuilder from a row.
	From(row Row) RowBuilder
}

type rowBuilder struct {
	name  string
	order int
	seats []Seat
}

// NewRowBuilder returns a new RowBuilder.
func NewRowBuilder() RowBuilder {
	return &rowBuilder{}
}

func (b *rowBuilder) WithName(name string) RowBuilder {
	b.name = name
	return b
}

func (b *rowBuilder) WithOrder(order int) RowBuilder {
	b.order = order
	return b
}

func (b *rowBuilder) WithSeat(seat ...Seat) RowBuilder {
	b.seats = append(b.seats, seat...)
	return b
}

func (b *rowBuilder) From(row Row) RowBuilder {
	return &rowBuilder{
		name:  row.Name(),
		order: row.Order(),
		seats: row.Seats(),
	}
}

func (b *rowBuilder) Build() Row {
	return &row{
		name:  b.name,
		order: b.order,
		seats: b.seats,
	}
}

type row struct {
	name  string
	order int
	seats []Seat
}

func (r *row) Name() string {
	return r.name
}

func (r *row) Order() int {
	return r.order
}

func (r *row) Seats() []Seat {
	return r.seats
}

func (r *row) String() string {
	var str string
	for _, seat := range r.seats {
		str += seat.String()
	}
	return str
}

func (r *row) Copy() Row {
	return &row{
		name:  r.name,
		order: r.order,
		seats: r.seats,
	}
}

func FilteredSeatsInRow(row Row, filter Filter) []Seat {
	availableSeats := make([]Seat, 0, len(row.Seats()))

	for _, seat := range row.Seats() {
		if filter.FilterSeat(seat) != nil {
			availableSeats = append(availableSeats, seat)
		}
	}

	return availableSeats
}

func ConsecutiveFilteredSeatsInRow(row Row, filter Filter) [][]Seat {
	consecutiveAvailableSeats := make([][]Seat, 0, len(row.Seats()))

	for i, seat := range row.Seats() {
		if filter.FilterSeat(seat) == nil {
			continue
		}

		if i == 0 {
			consecutiveAvailableSeats = append(consecutiveAvailableSeats, []Seat{seat})
			continue
		}

		prevSeat := row.Seats()[i-1]

		// Previous seat is same as filtering result with current.
		if filter.FilterSeat(prevSeat) != nil {
			consecutiveAvailableSeats[len(consecutiveAvailableSeats)-1] = append(consecutiveAvailableSeats[len(consecutiveAvailableSeats)-1], seat)
			continue
		}

		// Previous seat is different from filtering result with current.
		consecutiveAvailableSeats = append(consecutiveAvailableSeats, []Seat{seat})
	}

	return consecutiveAvailableSeats
}
