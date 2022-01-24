package seating_test

import (
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/seating"
)

func TestAllocateSeats(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		users []int
		want  [][]int
	}{
		{"simple",
			[]int{1, 3, 4, 4, 5, 1, 2, 4},
			[][]int{
				{1, 2, 2, 2, 3, 3, 3, 3},
				{4, 4, 4, 4, 5, 5, 5, 5},
				{5, 6, 7, 7, 8, 8, 8, 8},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := seating.AllocateSeats(test.users)
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
