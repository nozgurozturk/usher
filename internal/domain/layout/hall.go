package layout

// HallBuilder is a builder for a hall.
type HallBuilder interface {
	// WithID sets the id of the hall.
	WithID(id string) HallBuilder
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
	id       string
	name     string
	sections []Section
}

// NewHallBuilder returns a new HallBuilder.
func NewHallBuilder() HallBuilder {
	return &hallBuilder{}
}

func (b *hallBuilder) WithID(id string) HallBuilder {
	b.id = id
	return b
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
		id:       b.id,
		name:     b.name,
		sections: b.sections,
	}
}

type hall struct {
	id       string
	name     string
	sections []Section
}

func (h *hall) ID() string {
	return h.id
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

func (h *hall) JSON() interface{} {
	sections := make([]interface{}, len(h.sections))
	for i, section := range h.sections {
		sections[i] = section.JSON()
	}
	return struct {
		ID       string        `json:"id"`
		Name     string        `json:"name"`
		Sections []interface{} `json:"sections"`
	}{
		ID:       h.ID(),
		Name:     h.Name(),
		Sections: sections,
	}
}

func (h *hall) Copy() Hall {
	return &hall{
		name:     h.name,
		sections: h.sections,
	}
}

func (h *hall) ToSeatMap() SeatMap {
	seatMap := make(SeatMap, len(h.sections))
	for i, section := range h.sections {
		sectionMap := make(SectionMap, len(section.Rows()))
		for j, row := range section.Rows() {
			rowMap := make(RowMap, len(row.Seats()))
			for k, seat := range row.Seats() {
				if seat.Available() {
					rowMap[k] = 1
				}
			}
			sectionMap[j] = rowMap
		}
		seatMap[i] = sectionMap
	}
	return seatMap

}
