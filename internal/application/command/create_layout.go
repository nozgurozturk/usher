package command

import (
	"context"
	"strconv"

	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

type Seat string

type Row struct {
	Seats []Seat
}

type Section struct {
	Name    string
	Feature string
	Rows    []Row
}

type Layout struct {
	Name      string
	Numbering string
	Sections  []Section
}

// CreateLayoutCommand is a command to create a layout.
type CreateLayoutCommand struct {
	Name      string
	Numbering string
	Sections  []Section
}

// CreateLayoutCommandHandler is a handler for CreateLayoutCommand.
type CreateLayoutHandler struct {
	store *ent.Client
}

// NewCreateLayoutHandler return nil,s a new CreateLayoutHandler.
func NewCreateLayoutHandler(s *ent.Client) CreateLayoutHandler {
	return CreateLayoutHandler{
		store: s,
	}
}

func (h CreateLayoutHandler) Handle(c CreateLayoutCommand) (layout.Hall, error) {
	ctx := context.Background()
	hallBuilder := layout.NewHallBuilder().WithName(c.Name)

	sections := make([]layout.Section, len(c.Sections))

	for sectionIndex, section := range c.Sections {

		sectionBuilder := layout.NewSectionBuilder().WithName(section.Name)
		rows := make([]layout.Row, len(section.Rows))

		for rowIndex, row := range section.Rows {
			rowName := strconv.Itoa(rowIndex + 1)
			rowBuilder := layout.NewRowBuilder().WithName(rowName).WithOrder(rowIndex)
			seats := make([]layout.Seat, len(row.Seats))

			for seatIndex, seat := range row.Seats {
				// If numbering style is sequential like "1", "2", "3", ...
				number := seatIndex + 1

				// if numbering style is not sequential "1", "3", "5", "6", "4", "2" (Odd-Even) or "2", "4", "6", "5", "3", "1" (Even-Odd)
				if c.Numbering != "sequential" {
					number = number * 2
				}

				if c.Numbering == "odd-even" && len(row.Seats)/2 <= seatIndex {
					number -= 1
				}

				if c.Numbering == "even-odd" && len(row.Seats)/2 > seatIndex {
					number -= 1
				}

				seatBuilder := layout.NewSeatBuilder().WithPosition(rowIndex, seatIndex).WithNumber(number)

				if rank, err := strconv.Atoi(string(seat)); err == nil {
					seatBuilder.WithRank(rank)
				}

				if seatIndex == 0 || len(row.Seats)-1 == seatIndex {
					seatBuilder.WithFeature(layout.SeatFeatureAisle)
				}

				if section.Feature == "balcony" {
					seatBuilder.WithFeature(layout.SeatFeatureHigh)
				}

				s := seatBuilder.Build()

				seats[seatIndex] = s

			}
			r := rowBuilder.WithName(rowName).WithSeat(seats...).Build()
			rows[rowIndex] = r
		}
		s := sectionBuilder.WithRow(rows...).Build()
		sections[sectionIndex] = s
	}
	hallBuilder.WithSection(sections...)

	l := hallBuilder.Build()

	tx, err := h.store.Tx(ctx)
	if err != nil {
		return nil, err
	}

	sectionEntities := make([]*ent.SectionCreate, len(l.Sections()))

	for sectionIndex, section := range l.Sections() {
		rowEntities := make([]*ent.RowCreate, len(section.Rows()))

		for rowIndex, row := range section.Rows() {
			seatEntities := make([]*ent.SeatCreate, len(row.Seats()))

			for seatIndex, seat := range row.Seats() {
				row, col := seat.Position()
				seatEntities[seatIndex] = tx.Seat.Create().
					SetRow(row).
					SetCol(col).
					SetNumber(seat.Number()).
					SetFeature(int(seat.Feature())).
					SetRank(seat.Rank()).
					SetIsAvailable(seat.Available())
			}

			seats, err := tx.Seat.CreateBulk(seatEntities...).Save(ctx)
			if err != nil {
				return nil, tx.Rollback()
			}

			rowEntities[rowIndex] = tx.Row.Create().
				SetName(row.Name()).
				SetOrder(row.Order()).
				AddSeats(seats...)
		}

		rows, err := tx.Row.CreateBulk(rowEntities...).Save(ctx)
		if err != nil {
			return nil, tx.Rollback()
		}

		sectionEntities[sectionIndex] = tx.Section.Create().
			SetName(section.Name()).
			AddRows(rows...)
	}

	entSections, err := tx.Section.CreateBulk(sectionEntities...).Save(ctx)
	if err != nil {
		return nil, tx.Rollback()
	}

	_, err = tx.Layout.Create().AddSections(entSections...).Save(ctx)
	if err != nil {
		return nil, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return l, nil
}
