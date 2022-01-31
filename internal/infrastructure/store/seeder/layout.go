package seeder

import (
	"context"
	"strconv"
	"strings"

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

func CreateLayouts(c *ent.Client, ls ...Layout) []*ent.Layout {
	ctx := context.Background()
	halls := make([]layout.Hall, len(ls))

	for i, l := range ls {

		hallBuilder := layout.NewHallBuilder().WithName(l.Name)

		sections := make([]layout.Section, len(l.Sections))

		for sectionIndex, section := range l.Sections {

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
					if l.Numbering != "sequential" {
						number = number * 2
					}

					if l.Numbering == "odd-even" {
						if len(row.Seats)/2 > seatIndex {
							number -= 1
						} else {
							number = (len(row.Seats) - seatIndex) * 2
						}
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

					if rowIndex == 0 {
						seatBuilder.WithFeature(layout.SeatFeatureFront)
					}

					letters := strings.Split(string(seat), "")
					if len(letters) > 1 {
						if rank, err := strconv.Atoi(letters[1]); err == nil {
							seatBuilder.WithRank(rank)
						} else {
							seatBuilder.WithRank(1)
						}
					}

					s := seatBuilder.Build()

					if len(letters) > 1 {
						s.Book()
					}

					seats[seatIndex] = s

				}
				r := rowBuilder.WithName(rowName).WithSeat(seats...).Build()
				rows[rowIndex] = r
			}
			s := sectionBuilder.WithRow(rows...).Build()
			sections[sectionIndex] = s
		}
		hallBuilder.WithSection(sections...)

		halls[i] = hallBuilder.Build()
	}

	tx, err := c.Tx(ctx)
	if err != nil {
		panic(err)
	}

	layoutCreates := make([]*ent.LayoutCreate, len(halls))

	for layoutIndex, l := range halls {

		sectionCreates := make([]*ent.SectionCreate, len(l.Sections()))

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
					tx.Rollback()
					panic(err)
				}

				rowEntities[rowIndex] = tx.Row.Create().
					SetName(row.Name()).
					SetOrder(row.Order()).
					AddSeats(seats...)
			}

			rows, err := tx.Row.CreateBulk(rowEntities...).Save(ctx)
			if err != nil {
				tx.Rollback()
				panic(err)
			}

			sectionCreates[sectionIndex] = tx.Section.Create().
				SetName(section.Name()).
				AddRows(rows...)
		}

		entSections, err := tx.Section.CreateBulk(sectionCreates...).Save(ctx)
		if err != nil {
			tx.Rollback()
			panic(err)
		}

		layoutCreates[layoutIndex] = tx.Layout.Create().
			SetName(l.Name()).
			AddSections(entSections...)

	}

	layouts, err := tx.Layout.CreateBulk(layoutCreates...).Save(ctx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		panic(err)
	}
	return layouts
}
