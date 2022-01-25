package packing_test

import (
	"reflect"
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/packing"
)

type mockGroup int

func (g mockGroup) Size() int {
	return int(g)
}

func TestPackGroups(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		groups   []mockGroup         // group sizes
		capacity int                 // bin capacity
		fn       packing.PackingFunc // packing algorithm
		want     [][]mockGroup
	}{
		{
			"best-fit packing",
			[]mockGroup{1, 3, 4, 4, 5, 1, 2, 4},
			8,
			packing.BestFit,
			[][]mockGroup{
				{1, 3, 4},
				{4, 4},
				{5, 1, 2},
			},
		},
		{
			"best-fit packing",
			[]mockGroup{5, 7, 5, 2, 4, 2, 5, 1, 6},
			10,
			packing.BestFit,
			[][]mockGroup{
				{5, 5},
				{7, 2, 1},
				{4, 2},
				{5},
				{6},
			},
		},
		{
			"best-fit-decrased packing",
			[]mockGroup{5, 4, 4, 4, 3, 2, 1, 1},
			8,
			packing.BestFit,
			[][]mockGroup{
				{5, 3},
				{4, 4},
				{4, 2, 1, 1},
			},
		},
		{
			"best-fit-decrased packing",
			[]mockGroup{7, 6, 5, 5, 5, 4, 2, 2, 1},
			10,
			packing.BestFit,
			[][]mockGroup{
				{7, 2, 1},
				{6, 4},
				{5, 5},
				{5, 2},
			},
		},
		{
			"first-fit packing",
			[]mockGroup{1, 3, 4, 4, 5, 1, 2, 4},
			8,
			packing.FirstFit,
			[][]mockGroup{
				{1, 3, 4},
				{4, 1, 2},
				{5},
				{4},
			},
		},
		{
			"first-fit packing",
			[]mockGroup{5, 7, 5, 2, 4, 2, 5, 1, 6},
			10,
			packing.FirstFit,
			[][]mockGroup{
				{5, 5},
				{7, 2, 1},
				{4, 2},
				{5},
				{6},
			},
		},
		{
			"first-fit-decrased packing",
			[]mockGroup{5, 4, 4, 4, 3, 2, 1, 1},
			8,
			packing.FirstFit,
			[][]mockGroup{
				{5, 3},
				{4, 4},
				{4, 2, 1, 1},
			},
		},
		{
			"first-fit-decrased packing",
			[]mockGroup{7, 6, 5, 5, 5, 4, 2, 2, 1},
			10,
			packing.FirstFit,
			[][]mockGroup{
				{7, 2, 1},
				{6, 4},
				{5, 5},
				{5, 2},
			},
		},
		{
			"next-fit packing",
			[]mockGroup{1, 3, 4, 4, 5, 1, 2, 4},
			8,
			packing.NextFit,
			[][]mockGroup{
				{1, 3, 4},
				{4},
				{5, 1, 2},
				{4},
			},
		},
		{
			"next-fit packing",
			[]mockGroup{5, 7, 5, 2, 4, 2, 5, 1, 6},
			10,
			packing.NextFit,
			[][]mockGroup{
				{5},
				{7},
				{5, 2},
				{4, 2},
				{5, 1},
				{6},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			groups := asGroups(test.groups)

			got := packing.PackGroups(groups, test.capacity, test.fn)
			want := make([][]packing.Group, len(test.want), len(test.want))

			for i := range test.want {
				want[i] = asGroups(test.want[i])
			}

			comparePackedGroups(t, got, want)
		})
	}
}

var (
	unorderedGroups = asGroups([]mockGroup{5, 7, 5, 2, 4, 2, 5, 1, 6})
	decrasedGroups  = asGroups([]mockGroup{7, 6, 5, 5, 5, 4, 2, 2, 1})
	capacity        = 10
)

func benchmarksPackGroups(b *testing.B, groups []packing.Group, capacity int, fn packing.PackingFunc) {
	for i := 0; i < b.N; i++ {
		packing.PackGroups(decrasedGroups, capacity, fn)
	}
}

func BenchmarkNextFitPackGroups(b *testing.B) {
	benchmarksPackGroups(b, unorderedGroups, capacity, packing.NextFit)
}

func BenchmarkNextFitPackGroupsDecrased(b *testing.B) {
	benchmarksPackGroups(b, decrasedGroups, capacity, packing.NextFit)
}

func BenchmarkFirstFitPackGroups(b *testing.B) {
	benchmarksPackGroups(b, unorderedGroups, capacity, packing.FirstFit)
}
func BenchmarkFirstFitPackGroupsDecrased(b *testing.B) {
	benchmarksPackGroups(b, decrasedGroups, capacity, packing.FirstFit)
}

func BenchmarkBestFitPackGroups(b *testing.B) {
	benchmarksPackGroups(b, unorderedGroups, capacity, packing.BestFit)
}

func BenchmarkBestFitPackGroupsDecrased(b *testing.B) {
	benchmarksPackGroups(b, decrasedGroups, capacity, packing.BestFit)
}

/*

BenchmarkNextFitPackGroups            	 1478241	       796.1 ns/op	     608 B/op	      10 allocs/op
BenchmarkNextFitPackGroupsDecrased    	 1355222	       867.4 ns/op	     704 B/op	      10 allocs/op
BenchmarkFirstFitPackGroups           	 1235631	       962.2 ns/op	     784 B/op	      11 allocs/op
BenchmarkFirstFitPackGroupsDecrased   	 1191344	       1001 ns/op	     816 B/op	      11 allocs/op
BenchmarkBestFitPackGroups            	 1000000	       1050 ns/op	     784 B/op	      11 allocs/op
BenchmarkBestFitPackGroupsDecrased    	 1157578	       1005 ns/op	     816 B/op	      11 allocs/op

*/

func toMatrix(groups [][]packing.Group) [][]int {
	matrix := make([][]int, 0, len(groups))
	for _, group := range groups {
		sizes := make([]int, 0, len(group))
		for _, g := range group {
			sizes = append(sizes, g.Size())
		}
		matrix = append(matrix, sizes)
	}
	return matrix
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

func comparePackedGroups(t *testing.T, got, want [][]packing.Group) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	for i, row := range got {
		if len(row) != len(want[i]) {
			t.Errorf("got %v, want %v", got, want)
		}

		for j, col := range row {
			if !reflect.DeepEqual(col, want[i][j]) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	}
}

// Compare two groups.
func compareGroups(t *testing.T, got, want []packing.Group) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
	}

	for i, group := range got {
		if want[i].Size() != group.Size() {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

// helper function to convert a slice of mockGroups to a slice of packing.Groups.
func asGroups(set []mockGroup) []packing.Group {
	groups := make([]packing.Group, len(set), len(set))

	for i := range set {
		groups[i] = set[i]
	}

	return groups
}
