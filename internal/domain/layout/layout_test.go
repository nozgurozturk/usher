package layout_test

import (
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/layout"
)

func TestNewLayout(t *testing.T) {
	t.Parallel()

	layout := layout.NewLayout()

	if layout == nil {
		t.Error("expected layout, got nil")
	}
}

func TestNewSection(t *testing.T) {
	t.Parallel()

	s := layout.NewSection()

	if s == nil {
		t.Error("expected section, got nil")
	}
}

func TestSectionName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{"simple", "section"},
		{"empty", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := layout.NewSection()

			s.SetName(test.want)

			if s.Name() != test.want {
				t.Errorf("expected %s, got %s", test.want, s.Name())
			}
		})
	}
}

func TestAddSection(t *testing.T) {
	t.Parallel()

	l := layout.NewLayout()
	prevSize := len(l.Sections())

	s := layout.NewSection()
	l.AddSection(s)

	if len(l.Sections()) != prevSize+1 {
		t.Errorf("expected %d section, got %d", prevSize+1, len(l.Sections()))
	}
}

func TestAddRow(t *testing.T) {
	t.Parallel()

	s := layout.NewSection()

	prevSize := len(s.Rows())

	s.AddRow(layout.NewRow())

	if len(s.Rows()) != prevSize+1 {
		t.Errorf("expected %d row, got %d", prevSize+1, len(s.Rows()))
	}
}

func TestNewSeat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		rank   int
		number int
	}{
		{"simple", 1, 1},
		{"simple", 2, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := layout.NewSeat(test.rank, 0, test.number)

			if s == nil {
				t.Error("expected seat, got nil")
			}

			if s.Rank() != test.rank {
				t.Errorf("expected rank 1, got %d", s.Rank())
			}

			if s.Number() != test.number {
				t.Errorf("expected number 1, got %d", s.Number())
			}
		})
	}
}

func TestAvailableSeatsByRank(t *testing.T) {
	t.Parallel()

	r1 := layout.NewRow()
	r1.AddSeat(
		layout.NewSeat(1, 0, 1).Book(),
		layout.NewSeat(1, 1, 2),
	)

	r2 := layout.NewRow()
	r2.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 2).Book(),
	)

	r3 := layout.NewRow()
	r3.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(2, 1, 2),
	)

	s1 := layout.NewSection()
	s1.AddRow(r1, r2)

	s2 := layout.NewSection()
	s2.AddRow(r3)

	l := layout.NewLayout()

	l.AddSection(s1, s2)

	tests := []struct {
		name string
		rank int
		want int
	}{
		{"simple rank one", 1, 3},
		{"simple rank two", 2, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := l.AvailableSeatsByRank(test.rank)

			if len(got) != test.want {
				t.Errorf("expected %d available seats, got %d", test.want, len(got))
			}

		})
	}

}

func TestConsecutiveAvailableSeatsByRank(t *testing.T) {
	t.Parallel()

	r1 := layout.NewRow()
	r1.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2),
	)

	r2 := layout.NewRow()
	r2.AddSeat(
		layout.NewSeat(1, 0, 1).Book(),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2).Book(),
	)

	r3 := layout.NewRow()
	r3.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 3).Book(),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6).Book(),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2),
	)

	r4 := layout.NewRow()
	r4.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5).Book(),
		layout.NewSeat(1, 3, 6).Book(),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2),
	)

	r5 := layout.NewRow()
	r5.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6),
		layout.NewSeat(1, 4, 4).Book(),
		layout.NewSeat(1, 5, 2),
	)

	r6 := layout.NewRow()
	r6.AddSeat(
		layout.NewSeat(2, 0, 1),
		layout.NewSeat(2, 1, 3),
		layout.NewSeat(2, 2, 5),
	)

	s1 := layout.NewSection()
	s1.AddRow(r1, r2, r3)

	s2 := layout.NewSection()
	s2.AddRow(r4, r5, r6)

	l := layout.NewLayout()
	l.AddSection(s1, s2)

	tests := []struct {
		name string

		rank int
		want [][]*layout.Seat
	}{
		{
			"simple",
			1,
			[][]*layout.Seat{
				{
					layout.NewSeat(1, 0, 1),
					layout.NewSeat(1, 1, 3),
					layout.NewSeat(1, 2, 5),
					layout.NewSeat(1, 3, 6),
					layout.NewSeat(1, 4, 4),
					layout.NewSeat(1, 5, 2),
				},
				{
					layout.NewSeat(1, 1, 3),
					layout.NewSeat(1, 2, 5),
					layout.NewSeat(1, 3, 6),
					layout.NewSeat(1, 4, 4),
				},
				{
					layout.NewSeat(1, 0, 1),
				},
				{
					layout.NewSeat(1, 2, 5),
				},
				{
					layout.NewSeat(1, 4, 4),
					layout.NewSeat(1, 5, 2),
				},
				{
					layout.NewSeat(1, 0, 1),
					layout.NewSeat(1, 1, 3),
				},
				{
					layout.NewSeat(1, 4, 4),
					layout.NewSeat(1, 5, 2),
				},
				{
					layout.NewSeat(1, 0, 1),
					layout.NewSeat(1, 1, 3),
					layout.NewSeat(1, 2, 5),
					layout.NewSeat(1, 3, 6),
				},
				{
					layout.NewSeat(1, 5, 2),
				},
			},
		},
		{
			"simple",
			2,
			[][]*layout.Seat{
				{
					layout.NewSeat(2, 0, 1),
					layout.NewSeat(2, 1, 3),
					layout.NewSeat(2, 2, 5),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := l.ConsecutiveAvailableSeatsByRank(test.rank)
			compareAvailableSeatLists(t, test.want, got)
		})
	}

}

func compareAvailableSeatLists(t *testing.T, want, got [][]*layout.Seat) {
	if len(got) != len(want) {
		t.Errorf("expected %d seats, got %d", len(want), len(got))
	}

	for i := range want {

		if len(got[i]) != len(want[i]) {
			t.Errorf("expected %d seats, got %d", len(want[i]), len(got[i]))
		}

		for j := range want[i] {
			if got[i][j].Rank() != want[i][j].Rank() {
				t.Errorf("expected rank %d, got %d", want[i][j].Rank(), got[i][j].Rank())
			}

			if got[i][j].Number() != want[i][j].Number() {
				t.Errorf("expected number %d, got %d", want[i][j].Number(), got[i][j].Number())
			}
		}
	}
}

// TestAvailableSeatsByRank_Empty tests that an empty layout returns an empty
// list.
func TestAvailableSeatsByRank_Empty(t *testing.T) {
	l := layout.NewLayout()

	if got := l.ConsecutiveAvailableSeatsByRank(1); len(got) != 0 {
		t.Errorf("expected 0 seats, got %d", len(got))
	}
}

// TestAvailableSeatsByRank_NoRows tests that an empty section returns an
// empty list.
func TestAvailableSeatsByRank_NoRows(t *testing.T) {
	s := layout.NewSection()

	l := layout.NewLayout()
	l.AddSection(s)

	if got := l.ConsecutiveAvailableSeatsByRank(1); len(got) != 0 {
		t.Errorf("expected 0 seats, got %d", len(got))
	}
}

// TestAvailableSeatsByRank_NoSeats tests that a section with no seats returns
// an empty list.
func TestAvailableSeatsByRank_NoSeats(t *testing.T) {

	s := layout.NewSection()
	s.AddRow(layout.NewRow())

	l := layout.NewLayout()
	l.AddSection(s)

	if got := l.ConsecutiveAvailableSeatsByRank(1); len(got) != 0 {
		t.Errorf("expected 0 seats, got %d", len(got))
	}
}

func CreateMockLayout() *layout.Layout {

	r1 := layout.NewRow()
	r1.AddSeat(
		layout.NewSeat(1, 0, 1).Book(),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6),
		layout.NewSeat(1, 4, 4).Book(),
		layout.NewSeat(1, 5, 2),
	)

	r2 := layout.NewRow()
	r2.AddSeat(
		layout.NewSeat(1, 0, 1).Book(),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2).Book(),
	)

	r3 := layout.NewRow()
	r3.AddSeat(
		layout.NewSeat(1, 0, 1).Book(),
		layout.NewSeat(1, 1, 3).Book(),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6).Book(),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2),
	)

	r4 := layout.NewRow()
	r4.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5).Book(),
		layout.NewSeat(1, 3, 6).Book(),
		layout.NewSeat(1, 4, 4),
		layout.NewSeat(1, 5, 2),
	)

	r5 := layout.NewRow()
	r5.AddSeat(
		layout.NewSeat(1, 0, 1),
		layout.NewSeat(1, 1, 3),
		layout.NewSeat(1, 2, 5),
		layout.NewSeat(1, 3, 6),
		layout.NewSeat(1, 4, 4).Book(),
		layout.NewSeat(1, 5, 2).Book(),
	)

	s1 := layout.NewSection()
	s1.AddRow(r1, r2, r3)
	s1.SetName("A")

	s2 := layout.NewSection()
	s2.AddRow(r4, r5)
	s2.SetName("B")

	l := layout.NewLayout()
	l.AddSection(s1, s2)

	return l
}
