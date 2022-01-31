package layout

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

	// fmt.Println("filter", f.feature, seat.Feature(), seat.HasFeature(f.feature), seat.String())

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

func FilteredSeatsInSection(section Section, filter Filter) []Seat {
	var seats []Seat

	for _, row := range section.Rows() {
		seats = append(seats, FilteredSeatsInRow(row, filter)...)
	}

	return seats
}

func ConsecutiveFilteredSeatsInSection(section Section, filter Filter) [][]Seat {
	var seats [][]Seat

	for _, row := range section.Rows() {
		seats = append(seats, ConsecutiveFilteredSeatsInRow(row, filter)...)
	}

	return seats
}

func FilteredSeatsInHall(hall Hall, filter Filter) []Seat {
	var seats []Seat
	for _, section := range hall.Sections() {
		seats = append(seats, FilteredSeatsInSection(section, filter)...)
	}
	return seats
}

func ConsecutiveFilteredSeatsInHall(hall Hall, filter Filter) [][]Seat {
	var seats [][]Seat
	for _, section := range hall.Sections() {
		seats = append(seats, ConsecutiveFilteredSeatsInSection(section, filter)...)
	}
	return seats
}
