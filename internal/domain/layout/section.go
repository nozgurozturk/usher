package layout

// SectionBuilder is a builder for a section.
type SectionBuilder interface {
	// WithName sets the name of the section.
	WithName(name string) SectionBuilder
	// WithRow adds a row to the section.
	WithRow(row ...Row) SectionBuilder
	// Build returns the section.
	Build() Section
	// FromSection returns a new SectionBuilder from a section.
	From(section Section) SectionBuilder
}

type sectionBuilder struct {
	name string
	rows []Row
}

// NewSectionBuilder returns a new SectionBuilder.
func NewSectionBuilder() SectionBuilder {
	return &sectionBuilder{}
}

func (b *sectionBuilder) WithName(name string) SectionBuilder {
	b.name = name
	return b
}

func (b *sectionBuilder) WithRow(row ...Row) SectionBuilder {
	b.rows = append(b.rows, row...)
	return b
}

func (b *sectionBuilder) From(section Section) SectionBuilder {
	return &sectionBuilder{
		name: section.Name(),
		rows: section.Rows(),
	}
}

func (b *sectionBuilder) Build() Section {
	return &section{
		name: b.name,
		rows: b.rows,
	}
}

// section is a section instance.
type section struct {
	name string
	rows []Row
}

func (s *section) Name() string {
	return s.name
}

func (s *section) Rows() []Row {
	return s.rows
}

func (s *section) Copy() Section {
	return &section{
		name: s.name,
		rows: s.rows,
	}
}

func (s *section) String() string {
	str := "S-" + s.Name() + "\n"
	for _, row := range s.Rows() {
		str += row.String() + "\n"
	}
	return str
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
