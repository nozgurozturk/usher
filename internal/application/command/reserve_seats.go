package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/nozgurozturk/usher/internal/domain/group"
	"github.com/nozgurozturk/usher/internal/domain/seating"
	"github.com/nozgurozturk/usher/internal/infrastructure/store"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	dao "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/event"
	daoRes "github.com/nozgurozturk/usher/internal/infrastructure/store/ent/reservation"
)

type ReserverSeatsCommand struct {
	EventID string
}

type ReserveSeatsHandler struct {
	store *ent.Client
}

func NewReserveSeatsHandler(s *ent.Client) ReserveSeatsHandler {
	return ReserveSeatsHandler{
		store: s,
	}
}

func (h ReserveSeatsHandler) Handle(c ReserverSeatsCommand) error {
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
		}).
		WithReservations(
			func(q *ent.ReservationQuery) {
				q.Where(daoRes.IsActive(true)).WithUser()
			},
		).
		Only(ctx)
	if err != nil {
		return err
	}

	event := store.ToEventEntity(eventDAO)
	event.ReservedSeatsWithSeatMap()

	groups := make([]group.Group, len(eventDAO.Edges.Reservations))
	for i, reservation := range eventDAO.Edges.Reservations {
		groups[i] = group.NewGroup(reservation.ID.String(), reservation.Size, reservation.Rank, group.SeatPreference(reservation.Preference))
	}

	l, g, err := seating.ReserveSeatsForGroups(groups, event.Hall())
	if err != nil {
		return err
	}

	tx, err := h.store.Tx(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	totalUserCount := 0
	// convert group to map for easier access
	groupMap := make(map[string]group.Group, len(g))
	for _, g := range g {
		groupMap[g.ID()] = g
		totalUserCount += g.Size()
	}

	ticketCreates := make([]*ent.TicketCreate, 0, totalUserCount)

	for _, reservation := range eventDAO.Edges.Reservations {
		group := groupMap[reservation.ID.String()]
		for i := 0; i < group.Size(); i++ {
			seatID, err := uuid.Parse(group.Seats()[i].ID())
			if err != nil {
				if err := tx.Rollback(); err != nil {
					return err
				}
				return err
			}
			ticketCreates = append(ticketCreates, tx.Ticket.Create().
				SetUser(reservation.Edges.User).
				SetSeatID(seatID).
				SetEvent(eventDAO))

		}
	}
	_, err = tx.Reservation.Update().Where(daoRes.HasEventWith(dao.ID(eventID))).SetIsActive(false).Save(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	_, err = tx.Ticket.CreateBulk(ticketCreates...).Save(ctx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	_, err = tx.Event.UpdateOne(eventDAO).
		SetSeatMap(event.SeatMap().FromHall(l).String()).
		Save(ctx)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}
