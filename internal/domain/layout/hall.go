package layout

// HallBuilder is a builder for a hall.
type HallBuilder interface {
	// WithName sets the name of the hall.
	WithName(name string) HallBuilder
	// WithSection adds a section to the hall.
	WithSection(section ...Section) HallBuilder
	// Build returns the hall.
	Build() Hall
	// FromHall returns a new HallBuilder from a hall.
	From(hall Hall) HallBuilder
}

type hallBuilder struct {
	name     string
	sections []Section
}

// NewHallBuilder returns a new HallBuilder.
func NewHallBuilder() HallBuilder {
	return &hallBuilder{}
}

func (b *hallBuilder) WithName(name string) HallBuilder {
	b.name = name
	return b
}

func (b *hallBuilder) WithSection(section ...Section) HallBuilder {
	b.sections = append(b.sections, section...)
	return b
}

func (b *hallBuilder) From(hall Hall) HallBuilder {
	return &hallBuilder{
		name:     hall.Name(),
		sections: hall.Sections(),
	}
}

func (b *hallBuilder) Build() Hall {
	return &hall{
		name:     b.name,
		sections: b.sections,
	}
}

type hall struct {
	name     string
	sections []Section
}

func (h *hall) Name() string {
	return h.name
}

func (h *hall) Sections() []Section {
	return h.sections
}

func (h *hall) String() string {
	str := "H-" + h.Name() + "\n"
	for _, section := range h.Sections() {
		str += section.String()
	}
	return str
}

func (h *hall) Copy() Hall {
	return &hall{
		name:     h.name,
		sections: h.sections,
	}
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
