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

func TestCreateRow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		seats []int
		want  []*layout.Seat
	}{
		{"simple",
			[]int{1, 1, 1},
			[]*layout.Seat{
				layout.NewSeat(1, 1),
				layout.NewSeat(1, 2),
				layout.NewSeat(1, 3),
			},
		},
		{"simple with different ranks",
			[]int{1, 3, 2},
			[]*layout.Seat{
				layout.NewSeat(1, 1),
				layout.NewSeat(3, 2),
				layout.NewSeat(2, 3),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := layout.CreateRow(test.seats)

			if r == nil {
				t.Error("expected row, got nil")
			}

			if len(r) != len(test.want) {
				t.Errorf("expected %d seats, got %d", len(test.want), len(r))
			}

			for i := range test.want {
				if r[i].Rank() != test.want[i].Rank() {
					t.Errorf("expected rank %d, got %d", test.want[i].Rank(), r[i].Rank())
				}
			}
		})
	}
}

func TestAddRow(t *testing.T) {
	t.Parallel()

	s := layout.NewSection()

	prevSize := len(s.Rows())

	r := layout.CreateRow([]int{1, 1, 1})

	s.AddRow(r)

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
			s := layout.NewSeat(test.rank, test.number)

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
