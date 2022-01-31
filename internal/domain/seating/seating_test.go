package seating_test

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/group"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/seating"
)

func TestReserveSeatsForGroups(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		sectionRowSeatSizes [3]int
		alreadyBooked       [][3]int
		groups              []group.Group
		rank                int
		err                 error
	}{
		{
			"simple",
			[3]int{2, 5, 6},
			[][3]int{
				{0, 1, 0},
				{5, 1, 0},
				{1, 2, 0},
				{3, 2, 0},
				{2, 0, 1},
				{3, 0, 1},
				{4, 1, 1},
			},
			[]group.Group{group.NewGroup(randomID(), 3, 1), group.NewGroup(randomID(), 2, 1), group.NewGroup(randomID(), 2, 1), group.NewGroup(randomID(), 1, 1), group.NewGroup(randomID(), 3, 1), group.NewGroup(randomID(), 1, 1), group.NewGroup(randomID(), 1, 1)},
			1,
			nil,
		},
		{
			"overflow",
			[3]int{2, 2, 2},
			[][3]int{
				{0, 1, 0},
			},
			[]group.Group{group.NewGroup(randomID(), 2, 1), group.NewGroup(randomID(), 2, 1), group.NewGroup(randomID(), 2, 1), group.NewGroup(randomID(), 2, 1)},
			1,
			seating.ErrNotEnoughSpace,
		},
		{
			"not enough seats",
			[3]int{2, 2, 2},
			[][3]int{
				{0, 1, 0},
			},
			[]group.Group{group.NewGroup(randomID(), 3, 1)},
			1,
			seating.ErrNotEnoughSeats,
		},
	}

	for _, test := range tests {
		totalSeats := test.sectionRowSeatSizes[0] * test.sectionRowSeatSizes[1] * test.sectionRowSeatSizes[2]
		bookedSeats := len(test.alreadyBooked)
		remainingSize := totalSeats - bookedSeats - totalSize(test.groups)

		testLayout := createLayout(test.sectionRowSeatSizes[0], test.sectionRowSeatSizes[1], test.sectionRowSeatSizes[2])
		bookSeats(testLayout, test.alreadyBooked)

		filter := layout.NewFilter().WithRank(test.rank).WithAvailable(true)
		t.Run(test.name, func(t *testing.T) {
			l, g, err := seating.ReserveSeatsForGroups(test.groups, testLayout)

			if !errors.Is(err, test.err) {
				t.Errorf("got %v, want %v", err, test.err)
			}

			if err == nil {
				for _, group := range g {
					if !group.IsSatisfied() {
						t.Errorf("group %s is not satisfied: %d - %v", group.ID(), group.Size(), group.Seats())
					}
				}
				if remainingSize != len(layout.FilteredSeatsInHall(l, filter)) {
					t.Errorf("got %d, want %d", len(layout.FilteredSeatsInHall(l, filter)), remainingSize)
				}
			}

		})

	}
}

func totalSize(group []group.Group) int {
	total := 0

	for _, g := range group {
		total += g.Size()
	}

	return total
}

// bookSeats is a helper function to book seats in a row.
//
// It takes a slice of [3]int, where the first element is the seat number, the second element is the row number, and the third element is the section number.
// [numberOfBookedSeats][seat, row, section]
func bookSeats(l layout.Hall, seats [][3]int) layout.Hall {
	for _, seat := range seats {
		l.Sections()[seat[2]].Rows()[seat[1]].Seats()[seat[0]].Book()
	}

	return l

}

func createLayout(section, row, seat int) layout.Hall {

	sections := make([]layout.Section, section)
	for i := 0; i < section; i++ {
		secB := layout.NewSectionBuilder()
		rows := make([]layout.Row, row)
		for j := 0; j < row; j++ {
			rb := layout.NewRowBuilder()
			seats := make([]layout.Seat, seat)
			for k := 0; k < seat; k++ {
				seat := layout.NewSeatBuilder().WithPosition(j, k).WithRank(1).WithNumber(k + 1).Build()
				seats[k] = seat
			}
			rb.WithSeat(seats...)
			rows[j] = rb.Build()
		}
		secB.WithRow(rows...)
		sections[i] = secB.Build()
	}

	return layout.NewHallBuilder().WithSection(sections...).Build()
}

func randomID() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
