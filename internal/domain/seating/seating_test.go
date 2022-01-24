package seating_test

import (
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/seating"
)

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
