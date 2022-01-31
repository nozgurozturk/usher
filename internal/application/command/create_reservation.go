package command

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/group"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/seating"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
)

type CreateReservationCommand struct {
	Count       int
	EventID     string
	UserID      string
	Preferences struct {
		Features *int
		Rank     *int
	}
}

type CreateReservationHandler struct {
	store *ent.Client
}

func NewCreateReservationHandler(s *ent.Client) CreateReservationHandler {
	return CreateReservationHandler{
		store: s,
	}
}

func (h CreateReservationHandler) Handle(c CreateReservationCommand) error {
	ctx := context.Background()

	eventID, err := uuid.Parse(c.EventID)
	if err != nil {
		return err
	}

	eventDAO, err := h.store.Event.Query().Where(dao.ID(eventID)).WithLayout(
		func(q *ent.LayoutQuery) {
			q.WithSections(func(sq *ent.SectionQuery) {
				sq.WithRows(func(rq *ent.RowQuery) {
					rq.WithSeats()
				})
			})
		}).Only(ctx)
	if err != nil {
		return err
	}

	event := store.ToEventEntity(eventDAO)
	event.ReservedSeatsWithSeatMap()

	filter := layout.NewFilter()
	if c.Preferences.Rank != nil {
		filter.WithRank(*c.Preferences.Rank)
	}
	if c.Preferences.Features != nil {
		filter.WithFeature(layout.SeatFeature(*c.Preferences.Features))
	}
	filter.WithAvailable(true)

	seatBlocks := layout.ConsecutiveFilteredSeatsInHall(event.Hall(), filter)
	if len(seatBlocks) < 1 {
		return errors.New("no seats available")
	}

	var rank int
	if c.Preferences.Rank != nil {
		rank = *c.Preferences.Rank
	}

	var features int
	if c.Preferences.Features != nil {
		features = *c.Preferences.Features
	}

	g := group.NewGroup("", c.Count, rank, group.SeatPreference(features))

	seats, _ := seating.FindClosestSeatBlock(g, seatBlocks)
	if len(seats) < 1 {
		return errors.New("no seats available")
	}

	for _, seat := range seats {
		seat.Book()
	}

	event.SeatMap().FromHall(event.Hall())
	event.ReservedSeatsWithSeatMap()

	userID, err := uuid.Parse(c.UserID)
	if err != nil {
		return err
	}

	tx, err := h.store.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Reservation.Create().
		SetUserID(userID).
		SetRank(rank).
		SetSize(c.Count).
		SetPreference(features).
		SetEvent(eventDAO).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Event.UpdateOne(eventDAO).SetSeatMap(event.SeatMap().String()).Save(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()

}
