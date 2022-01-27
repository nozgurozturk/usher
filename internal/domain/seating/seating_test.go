package seating_test

import (
	"errors"
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/packing"
	"github.com/nozgurozturk/usher/internal/domain/seating"
)

type mockGroup struct {
	size int
}

func (g *mockGroup) Size() int {
	return g.size
}

func createMockGroup(size int) packing.Group {
	return &mockGroup{size}
}
func TestAllocateSeats(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		users      []int
		isReversed bool
		want       [][]int
		err        error
	}{
		{"simple",
			[]int{1, 3, 4, 4, 5, 1, 2, 4},
			false,
			[][]int{
				{1, 2, 2, 2, 3, 3, 3, 3},
				{4, 4, 4, 4, 5, 5, 5, 5},
				{5, 6, 7, 7, 8, 8, 8, 8},
			},
			nil,
		},
		{"reversed order",
			[]int{1, 3, 4, 4, 5, 1, 2, 4},
			true,
			[][]int{
				{1, 2, 2, 2, 3, 3, 3, 3},
				{5, 5, 5, 5, 4, 4, 4, 4},
				{5, 6, 7, 7, 8, 8, 8, 8},
			},
			nil,
		},
		{"overflow",
			[]int{2, 3, 4, 4, 5, 1, 2, 4},
			false,
			nil,
			seating.ErrOverflow,
		},
		{"partially filled",
			[]int{1, 3, 4, 4, 5, 1, 2, 1},
			false,
			[][]int{
				{1, 2, 2, 2, 3, 3, 3, 3},
				{4, 4, 4, 4, 5, 5, 5, 5},
				{5, 6, 7, 7, 8, 0, 0, 0},
			},
			nil,
		},
		{"partially reversed filled",
			[]int{1, 3, 4, 4, 2},
			true,
			[][]int{
				{1, 2, 2, 2, 3, 3, 3, 3},
				{0, 0, 5, 5, 4, 4, 4, 4},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := seating.AllocateSeats(test.users, test.isReversed)
			if err != test.err {
				t.Errorf("got error %v, want nil", err)
			}
			compareMatrix(t, got, test.want)
		})
	}
}

func TestAllocateSeatsInLayout(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		sectionRowSeatSizes [3]int
		alreadyBooked       [][3]int
		groups              []packing.Group
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
			[]packing.Group{createMockGroup(3), createMockGroup(2), createMockGroup(2), createMockGroup(1), createMockGroup(3), createMockGroup(1), createMockGroup(1)},
			1,
			nil,
		},
		{
			"overflow",
			[3]int{2, 2, 2},
			[][3]int{
				{0, 1, 0},
			},
			[]packing.Group{createMockGroup(2), createMockGroup(2), createMockGroup(2), createMockGroup(2)},
			1,
			seating.ErrNotEnoughSpace,
		},
		{
			"not enough seats",
			[3]int{2, 2, 2},
			[][3]int{
				{0, 1, 0},
			},
			[]packing.Group{createMockGroup(3)},
			1,
			seating.ErrNotEnoughSeats,
		},
	}

	for _, test := range tests {
		totalSeats := test.sectionRowSeatSizes[0] * test.sectionRowSeatSizes[1] * test.sectionRowSeatSizes[2]
		bookedSeats := len(test.alreadyBooked)
		remainingSize := totalSeats - bookedSeats - totalSizeOfGroup(test.groups)

		testLayout := createLayout(test.sectionRowSeatSizes[0], test.sectionRowSeatSizes[1], test.sectionRowSeatSizes[2])
		bookSeats(testLayout, test.alreadyBooked)

		filter := layout.NewFilter().WithRank(test.rank).WithAvailable(true)
		t.Run(test.name, func(t *testing.T) {
			l, err := seating.AllocateSeatsInLayout(test.groups, testLayout, filter)

			if !errors.Is(err, test.err) {
				t.Errorf("got %v, want %v", err, test.err)
			}

			if err == nil && remainingSize != len(layout.FilteredSeatsInHall(l, filter)) {
				t.Errorf("got %d, want %d", len(layout.FilteredSeatsInHall(l, filter)), remainingSize)
			}

		})

	}

}

// Compare two matrixes.
func compareMatrix(t *testing.T, got, want [][]int) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	for i, row := range got {
		if len(row) != len(want[i]) {
			t.Errorf("got %v, want %v", got, want)
		}

		for j, col := range row {
			if col != want[i][j] {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	}
}

// totalSizeOfGroups returns the total size of a group.
func totalSizeOfGroup(group []packing.Group) int {
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
