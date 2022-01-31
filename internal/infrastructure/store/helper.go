package store

import (
	"context"

	"github.com/nozgurozturk/usher/internal/infrastructure/store/ent"
	"github.com/nozgurozturk/usher/internal/infrastructure/store/seeder"
)

// CreateInitialData creates initial data for the store.
func CreateInitialData(ctx context.Context, c *ent.Client) {

	_ = seeder.CreateUsers(c, "John", "Jane", "Jack")

	layout1 := seeder.Layout{
		Name:      "Main Hall",
		Numbering: "sequential",
		Sections: []seeder.Section{
			{
				Name:    "A",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
				},
			},
			{
				Name:    "B",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
					{
						Seats: []seeder.Seat{"4", "4", "4", "4", "4", "4", "4", "4", "4", "4"},
					},
					{
						Seats: []seeder.Seat{"4", "4", "4", "4", "4", "4", "4", "4", "4", "4"},
					},
					{
						Seats: []seeder.Seat{"5", "5", "5", "5", "5", "5", "5", "5", "5", "5"},
					},
				},
			},
			{
				Name:    "C",
				Feature: "balcony",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
				},
			},
		},
	}

	layout2 := seeder.Layout{
		Name:      "Small Hall",
		Numbering: "odd-even",
		Sections: []seeder.Section{
			{
				Name:    "A",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
				},
			},
			{
				Name:    "B",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
					{
						Seats: []seeder.Seat{"4", "4", "4", "4", "4", "4", "4", "4", "4", "4"},
					},
					{
						Seats: []seeder.Seat{"5", "5", "5", "5", "5", "5", "5", "5", "5", "5"},
					},
				},
			},
			{
				Name:    "C",
				Feature: "balcony",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "1", "1", "1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "2", "2", "2", "2", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "3", "3", "3", "3", "3", "3"},
					},
				},
			},
		},
	}

	layout3 := seeder.Layout{
		Name:      "Medium Hall",
		Numbering: "odd-even",
		Sections: []seeder.Section{
			{
				Name:    "A",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"1", "1", "1", "1", "X1", "X1", "X1", "1", "1", "1"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "1", "1", "1", "1", "2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "3", "X2", "2", "3", "3", "3", "3"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "3", "X3", "3", "X3", "3", "3", "3", "3"},
					},
				},
			},
			{
				Name:    "B",
				Feature: "default",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"3", "3", "X3", "X3", "3", "3", "3", "3", "3", "3"},
					},
					{
						Seats: []seeder.Seat{"5", "5", "4", "4", "4", "X4", "X4", "4", "5", "5"},
					},
					{
						Seats: []seeder.Seat{"5", "5", "5", "5", "5", "5", "5", "5", "5", "5"},
					},
				},
			},
			{
				Name:    "C",
				Feature: "balcony",
				Rows: []seeder.Row{
					{
						Seats: []seeder.Seat{"2", "2", "1", "X1", "1", "X1", "1", "1", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"2", "2", "2", "3", "3", "3", "X3", "X2", "2", "2"},
					},
					{
						Seats: []seeder.Seat{"3", "3", "X3", "3", "X4", "4", "3", "3", "3", "3"},
					},
				},
			},
		},
	}

	ls := seeder.CreateLayouts(c, layout1, layout2, layout3)

	event1 := seeder.Event{
		Name:        "Phantom of the Opera",
		Description: "Phantom of the Opera is a musical comedy-drama film directed by Peter Jackson and written by Jackson's co-writer and producer, Ian McDonald. It is the first film in the Jackson family's second feature-length trilogy, and the first in the series to be produced by a studio other than Warner Bros. Pictures. The film is set in London, England, and depicts the story of the Phantom of the Opera, a young opera singer who is pursued by the Phantom of the Opera Society, a group of wealthy, immortal, immortalized ghosts who are determined to destroy the singer's career.",
		SeatMap:     "",
		LayoutID:    ls[0].ID,
	}

	event2 := seeder.Event{
		Name:        "Jesus Christ Superstar",
		Description: "Jesus Christ Superstar is a sung-through rock opera with music by Andrew Lloyd Webber and lyrics by Tim Rice. ",
		LayoutID:    ls[1].ID,
	}

	event3 := seeder.Event{
		Name:        "The Sound of Music",
		Description: "The Sound of Music is a musical comedy-drama film directed by Robert Wise and written by Robert Wise and Andrew Lloyd Webber. It is the second film in the Wise family's second feature-length trilogy, and the first in the series to be produced by a studio other than Warner Bros. Pictures. The film is set in the fictional Republic of New York, and depicts the story of the Sound of Music, a young musical singer who is pursued by the Sound of Music Society, a group of wealthy, immortal, immortalized ghosts who are determined to destroy the singer's career.",
		LayoutID:    ls[2].ID,
	}

	event4 := seeder.Event{
		Name:        "Rent",
		Description: "Rent is a rock musical with music, lyrics, and book by Jonathan Larson, loosely based on Giacomo Puccini's 1896 opera La Bohème",
		LayoutID:    ls[1].ID,
	}

	_ = seeder.CreateEvents(c, event1, event2, event3, event4)
}

func ResetDatabase(ctx context.Context, c *ent.Client) {
	c.Event.Delete().ExecX(ctx)
	c.Reservation.Delete().ExecX(ctx)
	c.Seat.Delete().ExecX(ctx)
	c.Row.Delete().ExecX(ctx)
	c.Section.Delete().ExecX(ctx)
	c.Layout.Delete().ExecX(ctx)
	c.Ticket.Delete().ExecX(ctx)
	c.User.Delete().ExecX(ctx)
}
