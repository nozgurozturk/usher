package layout_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/layout"
)

type mockRowMap []string
type mockSectionMap []mockRowMap
type mockHallMap []mockSectionMap

func TestNewSeat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		row, col int
		number   int
		rank     int
		features []layout.SeatFeature
	}{
		{"simple", 1, 1, 1, 1, nil},
		{"with features", 1, 1, 1, 1, []layout.SeatFeature{layout.SeatFeatureHigh, layout.SeatFeatureFront}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			sb := layout.NewSeatBuilder().
				WithPosition(test.row, test.col).
				WithNumber(test.number).
				WithRank(test.rank)

			if len(test.features) > 0 {
				for _, f := range test.features {
					sb.WithFeature(f)
				}
			}

			seat := sb.Build()

			seatRow, seatCol := seat.Position()
			if seatRow != test.row {
				t.Errorf("seat.Position()[0] = %d, want %d", seatRow, test.row)
			}

			if seatCol != test.col {
				t.Errorf("seat.Position()[1] = %d, want %d", seatCol, test.col)
			}

			if seat.Number() != test.number {
				t.Errorf("seat.Number() = %d, want %d", seat.Number(), test.number)
			}

			if seat.Rank() != test.rank {
				t.Errorf("seat.Rank() = %d, want %d", seat.Rank(), test.rank)
			}

			if len(test.features) > 0 {
				for _, f := range test.features {
					if !seat.HasFeature(f) {
						t.Errorf("seat.HasFeature(%v) = false, want true", f)
					}
				}
			}

		})
	}
}

func TestNewRow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rowName string
		order   int
		seats   []layout.Seat
	}{
		{
			"simple", "1", 1, nil,
		},
		{
			"with seats", "1", 1,
			[]layout.Seat{
				layout.NewSeatBuilder().Build(),
				layout.NewSeatBuilder().Build(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			row := layout.NewRowBuilder().
				WithName(test.rowName).
				WithOrder(test.order).
				WithSeat(test.seats...).
				Build()

			if row.Name() != test.rowName {
				t.Errorf("row.Name() = %s, want %s", row.Name(), test.rowName)
			}

			if row.Order() != test.order {
				t.Errorf("row.Order() = %d, want %d", row.Order(), test.order)
			}

			if len(row.Seats()) != len(test.seats) {
				t.Errorf("len(row.Seats()) = %d, want %d", len(row.Seats()), len(test.seats))
			}
		})
	}
}

func TestNewSection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		sectionName string
		rows        []layout.Row
	}{
		{
			"simple", "1", nil,
		},
		{
			"with rows", "1",
			[]layout.Row{
				layout.NewRowBuilder().Build(),
				layout.NewRowBuilder().Build(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			section := layout.NewSectionBuilder().
				WithName(test.sectionName).
				WithRow(test.rows...).
				Build()

			if section.Name() != test.sectionName {
				t.Errorf("section.Name() = %s, want %s", section.Name(), test.sectionName)
			}

			if len(section.Rows()) != len(test.rows) {
				t.Errorf("len(section.Rows()) = %d, want %d", len(section.Rows()), len(test.rows))
			}
		})
	}
}

func TestNewHall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		hallName string
		sections []layout.Section
	}{
		{
			"simple", "main", nil,
		},
		{
			"with sections", "main",
			[]layout.Section{
				layout.NewSectionBuilder().Build(),
				layout.NewSectionBuilder().Build(),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hall := layout.NewHallBuilder().
				WithName(test.hallName).
				WithSection(test.sections...).
				Build()

			if hall.Name() != test.hallName {
				t.Errorf("hall.Name() = %s, want %s", hall.Name(), test.hallName)
			}

			if len(hall.Sections()) != len(test.sections) {
				t.Errorf("len(hall.Sections()) = %d, want %d", len(hall.Sections()), len(test.sections))
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	mockSectionOne := mockSectionMap{
		{"X", "1", "1", "1", "1", "1"}, // X is a seat with booked before // "1" represents a seat rank
		{"X", "1", "1", "1", "1", "X"},
		{"X", "1", "X", "X", "1", "1"},
		{"1", "1", "X", "X", "1", "1"},
		{"1", "1", "1", "1", "X", "X"},
	}

	mockSectionTwo := mockSectionMap{
		{"1", "1", "X", "X", "1", "1"},
		{"1", "1", "1", "1", "X", "X"},
		{"2", "2", "2", "2", "2", "2"},
	}

	mockHall := CreateMockHall("mockHall", mockHallMap{
		mockSectionOne,
		mockSectionTwo, // has feature high
	})

	fmt.Printf("%+v\n", mockHall)

	tests := []struct {
		name   string
		filter layout.Filter
		want   int // number of seats
	}{
		{
			"filter by rank one",
			layout.NewFilter().WithRank(1),
			42,
		},
		{
			"filter by rank two",
			layout.NewFilter().WithRank(2),
			6,
		},
		{
			"filter by availability",
			layout.NewFilter().WithAvailable(true),
			34,
		},
		{
			"filter by feature front",
			layout.NewFilter().WithFeature(layout.SeatFeatureFront),
			12,
		},
		{
			"filter by feature high",
			layout.NewFilter().WithFeature(layout.SeatFeatureHigh),
			18,
		},
		{
			"filter by feature aisle",
			layout.NewFilter().WithFeature(layout.SeatFeatureAisle),
			16,
		},
		{
			"composite filter",
			layout.NewFilter().
				WithRank(1).
				WithAvailable(true).
				WithFeature(layout.SeatFeatureFront),
			9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := layout.FilteredSeatsInHall(mockHall, test.filter)

			if len(got) != test.want {
				t.Errorf("len(got) = %d, want %d", len(got), test.want)
			}
		})
	}
}

func TestConsecutiveAvailableSeats(t *testing.T) {
	t.Parallel()

	mockSectionOne := mockSectionMap{
		{"X", "1", "1", "1", "1", "1"}, // X is a seat with booked before // "1" represents a seat rank
		{"X", "1", "1", "1", "1", "X"},
		{"X", "1", "X", "X", "1", "1"},
		{"1", "1", "X", "X", "1", "1"},
		{"1", "1", "1", "1", "X", "X"},
	}

	mockSectionTwo := mockSectionMap{
		{"1", "1", "X", "X", "1", "1"},
		{"1", "1", "1", "1", "X", "X"},
		{"2", "2", "2", "2", "2", "2"},
	}

	mockHall := CreateMockHall("mockHall", mockHallMap{
		mockSectionOne,
		mockSectionTwo, // has feature high
	})

	tests := []struct {
		name   string
		filter layout.Filter
		want   []int // number of seats
	}{
		{
			"filter by rank one",
			layout.NewFilter().WithRank(1),
			[]int{6, 6, 6, 6, 6, 6, 6},
		},
		{
			"filter by rank two",
			layout.NewFilter().WithRank(2),
			[]int{6},
		},
		{
			"filter by availability",
			layout.NewFilter().WithAvailable(true),
			[]int{5, 4, 1, 2, 2, 2, 4, 2, 2, 4, 6},
		},
		{
			"filter by availability and rank",
			layout.NewFilter().WithAvailable(true).WithRank(1),
			[]int{5, 4, 1, 2, 2, 2, 4, 2, 2, 4},
		},
		{
			"filter by feature front",
			layout.NewFilter().WithFeature(layout.SeatFeatureFront),
			[]int{6, 6},
		},
		{
			"filter by feature high",
			layout.NewFilter().WithFeature(layout.SeatFeatureHigh),
			[]int{6, 6, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := layout.ConsecutiveFilteredSeatsInHall(mockHall, test.filter)

			for i := range got {

				if len(got[i]) != test.want[i] {
					t.Errorf("len(got[%d]) = %d, want %d", i, len(got[i]), test.want[i])
				}
			}

		})
	}
}

func CreateMockHall(name string, hallMap mockHallMap) layout.Hall {

	hallBuilder := layout.NewHallBuilder().WithName(name)

	sections := make([]layout.Section, len(hallMap))

	for sectionIndex, sectionRows := range hallMap {
		sectionName := ""
		if sectionIndex == 0 {
			sectionName = "Main"
		} else {
			sectionName = "Balcony"
		}
		sectionBuilder := layout.NewSectionBuilder().WithName(sectionName)
		rows := make([]layout.Row, len(sectionRows))

		for rowIndex, rowSeats := range sectionRows {
			rowName := strconv.Itoa(rowIndex + 1)
			rowBuilder := layout.NewRowBuilder().WithName(rowName).WithOrder(rowIndex)
			seats := make([]layout.Seat, len(rowSeats))

			for seatIndex, seat := range rowSeats {
				seatBuilder := layout.NewSeatBuilder().WithPosition(rowIndex, seatIndex).WithNumber(seatIndex + 1)

				if rank, err := strconv.Atoi(seat); err == nil {
					seatBuilder.WithRank(rank)
				} else {
					seatBuilder.WithRank(1)
				}

				if rowName == "1" {
					seatBuilder.WithFeature(layout.SeatFeatureFront)
				}

				if seatIndex == 0 || len(rowSeats)-1 == seatIndex {
					seatBuilder.WithFeature(layout.SeatFeatureAisle)
				}

				if sectionIndex > 0 {
					seatBuilder.WithFeature(layout.SeatFeatureHigh)
				}

				s := seatBuilder.Build()

				if seat == "X" {
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
	return hallBuilder.Build()
}
