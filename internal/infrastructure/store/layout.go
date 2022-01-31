package store

import (
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
)

func ToLayoutEntity(e *ent.Layout) layout.Hall {

	h := layout.NewHallBuilder()
	h.WithID(e.ID.String()).
		WithName(e.Name)

	if e.Edges.Sections != nil {
		for _, section := range e.Edges.Sections {

			sec := layout.NewSectionBuilder().
				WithName(section.Name)

			for _, row := range section.Edges.Rows {

				r := layout.NewRowBuilder().
					WithOrder(row.Order).
					WithName(row.Name)

				for _, seat := range row.Edges.Seats {
					s := layout.NewSeatBuilder().
						WithID(seat.ID.String()).
						WithPosition(seat.Row, seat.Col).
						WithNumber(seat.Number).
						WithFeature(layout.SeatFeature(seat.Feature)).
						WithRank(seat.Rank).
						Build()
					if !seat.IsAvailable {
						s.Book()
					}
					r.WithSeat(s)

				}
				sec.WithRow(r.Build())

			}
			h.WithSection(sec.Build())

		}
	}
	return h.Build()

}
