package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	daoReservation "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
)

// CheckEventSeatsQuery is a query for checking event seats
type CheckEventSeatsQuery struct {
	EventID  string
	Count    int
	Features *int
	Rank     *int
}

// CheckEventSeatsHandler is a handler for checking event seats
type CheckEventSeatsHandler struct {
	store *ent.Client
}

// NewCheckEventSeatsHandler creates a new CheckEventSeatsHandler
func NewCheckEventSeatsHandler(s *ent.Client) CheckEventSeatsHandler {
	return CheckEventSeatsHandler{
		store: s,
	}
}

// Handle checks event seats
func (h CheckEventSeatsHandler) Handle(q CheckEventSeatsQuery) (remaining int, err error) {
	ctx := context.Background()
	eventID, err := uuid.Parse(q.EventID)
	if err != nil {
		return 0, err
	}

	var rank int
	if q.Rank != nil {
		rank = *q.Rank
	}

	var pref int
	if q.Features != nil {
		pref = *q.Features
	}

	eventDAO, err := h.store.Event.Query().Where(dao.ID(eventID)).WithLayout(
		func(q *ent.LayoutQuery) {
			q.WithSections(func(sq *ent.SectionQuery) {
				sq.WithRows(func(rq *ent.RowQuery) {
					rq.WithSeats()
				})
			})
		}).
		WithReservations(func(rq *ent.ReservationQuery) {
			rq.Where(daoReservation.IsActive(true), daoReservation.Rank(rank), daoReservation.Preference(pref))
		}).
		Only(ctx)
	if err != nil {
		return 0, err
	}

	event := store.ToEventEntity(eventDAO)
	event.ReservedSeatsWithSeatMap()

	filter := layout.NewFilter()
	if rank > 0 {
		filter.WithRank(rank)
	}
	if pref > 0 {
		filter.WithFeature(layout.SeatFeature(pref))
	}

	filter.WithAvailable(true)

	totalUserCount := 0
	for _, r := range eventDAO.Edges.Reservations {
		totalUserCount += r.Size
	}
	seatBlocks := layout.ConsecutiveFilteredSeatsInHall(event.Hall(), filter)

	maxAvailableTotalSeat := 0
	maxAvailableTotalBlock := 0

	for _, block := range seatBlocks {
		if len(block) > maxAvailableTotalBlock {
			maxAvailableTotalSeat += len(block)
			maxAvailableTotalBlock = len(block)
		}
	}
	var min int
	if maxAvailableTotalSeat-totalUserCount < maxAvailableTotalBlock {
		min = maxAvailableTotalSeat - totalUserCount
	} else {
		min = maxAvailableTotalBlock
	}

	return min, nil
}
