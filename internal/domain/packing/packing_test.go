package packing_test

import (
	"fmt"
	"testing"

	"github.com/nozgurozturk/usher/internal/domain/packing"
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

func TestPackGroups(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		groups   []packing.Group     // group sizes
		capacity int                 // bin capacity
		fn       packing.PackingFunc // packing algorithm
		want     [][]packing.Group
	}{
		{
			"best-fit packing",
			[]packing.Group{
				createMockGroup(1), createMockGroup(3), createMockGroup(4), createMockGroup(4), createMockGroup(5), createMockGroup(1), createMockGroup(2), createMockGroup(4)},
			8,
			packing.BestFit,
			[][]packing.Group{
				{createMockGroup(1), createMockGroup(3), createMockGroup(4)},
				{createMockGroup(4), createMockGroup(4)},
				{createMockGroup(5), createMockGroup(1), createMockGroup(2)},
			},
		},
		{
			"best-fit packing",
			[]packing.Group{
				createMockGroup(5), createMockGroup(7), createMockGroup(5), createMockGroup(2), createMockGroup(4), createMockGroup(2), createMockGroup(5), createMockGroup(1), createMockGroup(6)},
			10,
			packing.BestFit,
			[][]packing.Group{
				{createMockGroup(5), createMockGroup(5)},
				{createMockGroup(7), createMockGroup(2), createMockGroup(1)},
				{createMockGroup(4), createMockGroup(2)},
				{createMockGroup(5)},
				{createMockGroup(6)},
			},
		},
		{
			"best-fit-decrased packing",
			[]packing.Group{createMockGroup(5), createMockGroup(4), createMockGroup(4), createMockGroup(4), createMockGroup(3), createMockGroup(2), createMockGroup(1), createMockGroup(1)},
			8,
			packing.BestFit,
			[][]packing.Group{
				{createMockGroup(5), createMockGroup(3)},
				{createMockGroup(4), createMockGroup(4)},
				{createMockGroup(4), createMockGroup(2), createMockGroup(1), createMockGroup(1)},
			},
		},
		{
			"best-fit-decrased packing",
			[]packing.Group{createMockGroup(7), createMockGroup(6), createMockGroup(5), createMockGroup(5), createMockGroup(5), createMockGroup(4), createMockGroup(2), createMockGroup(2), createMockGroup(1)},
			10,
			packing.BestFit,
			[][]packing.Group{
				{createMockGroup(7), createMockGroup(2), createMockGroup(1)},
				{createMockGroup(6), createMockGroup(4)},
				{createMockGroup(5), createMockGroup(5)},
				{createMockGroup(5), createMockGroup(2)},
			},
		},
		{
			"first-fit packing",
			[]packing.Group{createMockGroup(1), createMockGroup(3), createMockGroup(4), createMockGroup(4), createMockGroup(5), createMockGroup(1), createMockGroup(2), createMockGroup(4)},
			8,
			packing.FirstFit,
			[][]packing.Group{
				{createMockGroup(1), createMockGroup(3), createMockGroup(4)},
				{createMockGroup(4), createMockGroup(1), createMockGroup(2)},
				{createMockGroup(5)},
				{createMockGroup(4)},
			},
		},
		{
			"first-fit packing",
			[]packing.Group{createMockGroup(5), createMockGroup(7), createMockGroup(5), createMockGroup(2), createMockGroup(4), createMockGroup(2), createMockGroup(5), createMockGroup(1), createMockGroup(6)},
			10,
			packing.FirstFit,
			[][]packing.Group{
				{createMockGroup(5), createMockGroup(5)},
				{createMockGroup(7), createMockGroup(2), createMockGroup(1)},
				{createMockGroup(4), createMockGroup(2)},
				{createMockGroup(5)},
				{createMockGroup(6)},
			},
		},
		{
			"first-fit-decrased packing",
			[]packing.Group{createMockGroup(5), createMockGroup(4), createMockGroup(4), createMockGroup(4), createMockGroup(3), createMockGroup(2), createMockGroup(1), createMockGroup(1)},
			8,
			packing.FirstFit,
			[][]packing.Group{
				{createMockGroup(5), createMockGroup(3)},
				{createMockGroup(4), createMockGroup(4)},
				{createMockGroup(4), createMockGroup(2), createMockGroup(1), createMockGroup(1)},
			},
		},
		{
			"first-fit-decrased packing",
			[]packing.Group{createMockGroup(7), createMockGroup(6), createMockGroup(5), createMockGroup(5), createMockGroup(5), createMockGroup(4), createMockGroup(2), createMockGroup(2), createMockGroup(1)},
			10,
			packing.FirstFit,
			[][]packing.Group{
				{createMockGroup(7), createMockGroup(2), createMockGroup(1)},
				{createMockGroup(6), createMockGroup(4)},
				{createMockGroup(5), createMockGroup(5)},
				{createMockGroup(5), createMockGroup(2)},
			},
		},
		{
			"next-fit packing",
			[]packing.Group{createMockGroup(1), createMockGroup(3), createMockGroup(4), createMockGroup(4), createMockGroup(5), createMockGroup(1), createMockGroup(2), createMockGroup(4)},
			8,
			packing.NextFit,
			[][]packing.Group{
				{createMockGroup(1), createMockGroup(3), createMockGroup(4)},
				{createMockGroup(4)},
				{createMockGroup(5), createMockGroup(1), createMockGroup(2)},
				{createMockGroup(4)},
			},
		},
		{
			"next-fit packing",
			[]packing.Group{createMockGroup(5), createMockGroup(7), createMockGroup(5), createMockGroup(2), createMockGroup(4), createMockGroup(2), createMockGroup(5), createMockGroup(1), createMockGroup(6)},
			10,
			packing.NextFit,
			[][]packing.Group{
				{createMockGroup(5)},
				{createMockGroup(7)},
				{createMockGroup(5), createMockGroup(2)},
				{createMockGroup(4), createMockGroup(2)},
				{createMockGroup(5), createMockGroup(1)},
				{createMockGroup(6)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			groups := test.groups

			got := packing.PackGroups(groups, test.capacity, test.fn)
			want := make([][]packing.Group, len(test.want))

			for i := range test.want {
				want[i] = test.want[i]
			}

			comparePackedGroups(t, got, want)
		})
	}
}

var (
	unorderedGroups = []packing.Group{createMockGroup(5), createMockGroup(7), createMockGroup(5), createMockGroup(2), createMockGroup(4), createMockGroup(2), createMockGroup(5), createMockGroup(1), createMockGroup(6)}
	decrasedGroups  = []packing.Group{createMockGroup(7), createMockGroup(6), createMockGroup(5), createMockGroup(5), createMockGroup(5), createMockGroup(4), createMockGroup(2), createMockGroup(2), createMockGroup(1)}
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

cpu: Intel(R) Core(TM) i5-4278U CPU @ 2.60GHz
BenchmarkNextFitPackGroups-4            	 1517623	       773.6 ns/op	     464 B/op	      10 allocs/op
BenchmarkNextFitPackGroupsDecrased-4    	 1366284	       776.4 ns/op	     464 B/op	      10 allocs/op
BenchmarkFirstFitPackGroups-4           	 1246713	       935.6 ns/op	     560 B/op	      11 allocs/op
BenchmarkFirstFitPackGroupsDecrased-4   	 1252339	       925.5 ns/op	     560 B/op	      11 allocs/op
BenchmarkBestFitPackGroups-4            	 1192914	       997.5 ns/op	     560 B/op	      11 allocs/op
BenchmarkBestFitPackGroupsDecrased-4    	 1214422	       1090 ns/op	     560 B/op	      11 allocs/op
*/

// comparePackedGroups compares two packed groups and fails if they are not equal
func comparePackedGroups(t *testing.T, got, want [][]packing.Group) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %v, want %v", len(got), len(want))
	}

	for i, row := range got {
		if len(row) != len(want[i]) {
			t.Errorf("got %d, want %d - %d", len(row), len(want[i]), i)
		}

		for j, col := range row {
			if col.Size() != want[i][j].Size() {
				t.Errorf("got %s, want %s", groupstr(got), groupstr(want))
			}
		}
	}
}

func groupstr(listOfGroups [][]packing.Group) string {
	var str string
	for _, groups := range listOfGroups {
		str += "["
		for _, group := range groups {
			str += fmt.Sprintf("[%d]", group.Size())
		}
		str += "]"
	}

	return str
}
