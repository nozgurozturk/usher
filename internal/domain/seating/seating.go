package seating

import (
	"fmt"
	"sort"

	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/packing"
)

const (
	ROW_SIZE = 3
	COL_SIZE = 8
)

var (
	ErrOverflow       = fmt.Errorf("overflow")
	ErrNotEnoughSeats = fmt.Errorf("not enough seats")
	ErrNotEnoughSpace = fmt.Errorf("not enough space")
)

// AllocateSeats allocates seats for the given group of users.
// Given a list of “groups of users” per rank (basically sizes, e.g. `(1, 3, 4, 4, 5, 1, 2, 4)` in a specific order,
// try to place the users in their seats, e.g.
// Example:
//
// 		seats := AllocateSeats([]int{1, 3, 4, 4, 5, 1, 2, 4})
//
// 		fmt.Println(seats)
//
// Output:
//
// 		[[1 2 2 2 3 3 3 3]
// 		 [4 4 4 4 5 5 5 5]
// 		 [5 6 7 7 8 8 8 8]]

func AllocateSeats(groupsOfUser []int, isReversed bool) ([][]int, error) {

	totalUser := 0
	for _, groupSize := range groupsOfUser {
		totalUser += groupSize
	}

	if totalUser > ROW_SIZE*COL_SIZE {
		return nil, ErrOverflow
	}

	// nolint:gosimple
	seats := make([][]int, ROW_SIZE, ROW_SIZE)

	for i := 0; i < ROW_SIZE; i++ {
		// nolint:gosimple
		seats[i] = make([]int, COL_SIZE, COL_SIZE)
	}

	row := 0
	col := 0

	for i, groupSize := range groupsOfUser {

		group := i + 1

		for j := 0; j < groupSize; j++ {
			index := col

			// Reverse order for odd row numbers.
			if isReversed && row%2 != 0 {
				index = COL_SIZE - 1 - col
			}

			seats[row][index] = group
			col++

			if col == COL_SIZE {
				col = 0
				row++
			}

			if row == ROW_SIZE {
				break
			}

		}
	}

	return seats, nil
}

// AllocateSeatsInLayout allocates seats for the given group of users and layout.
func AllocateSeatsInLayout(groups []packing.Group, l *layout.Layout, rank int) (*layout.Layout, error) {

	// Find available list of seats and sort by length.
	availableListOfSeats := l.ConsecutiveAvailableSeatsByRank(rank)
	sort.SliceStable(availableListOfSeats, func(i, j int) bool {
		return len(availableListOfSeats[i]) < len(availableListOfSeats[j])
	})

	// check Total Size of groups is less than the available seats.
	if err := checkEnoughSpace(availableListOfSeats, groups); err != nil {
		return nil, err
	}

	gs := make([]packing.Group, len(groups), len(groups))
	copy(gs, groups)

	// Sort groups by size.
	sort.SliceStable(gs, func(i, j int) bool {
		return gs[i].Size() > gs[j].Size()
	})

FindAvailableSeat:
	for _, availableSeats := range availableListOfSeats {

		// Pack groups for the available seats.
		packedListOfGroups := packing.PackGroups(gs, len(availableSeats), packing.NextFit)
		// Sort list of groups by closes size to the available seats.
		sort.SliceStable(packedListOfGroups, func(i, j int) bool {
			return closestToAvailableSeats(packedListOfGroups[i], availableSeats) < closestToAvailableSeats(packedListOfGroups[j], availableSeats)
		})

		for _, packedGroups := range packedListOfGroups {

			// Book the seats.
			seatIndex := 0
			for _, group := range packedGroups {

				if group.Size() > len(availableSeats) {
					continue
				}

				for i := 0; i < group.Size(); i++ {
					availableSeats[seatIndex].Book()
					seatIndex++
				}

				gs = removeGroup(gs, group)
				if len(gs) == 0 {
					break FindAvailableSeat
				}

			}
			if len(availableSeats) >= seatIndex {
				continue FindAvailableSeat
			}
		}

	}

	if len(gs) > 0 {
		return nil, ErrNotEnoughSeats
	}

	return l, nil
}

// removeGroup removes the given group from the list of groups.
func removeGroup(groups []packing.Group, group packing.Group) []packing.Group {
	for i, g := range groups {
		if g.Size() == group.Size() {
			return append(groups[:i], groups[i+1:]...)
		}
	}
	return groups
}

// closestToAvailableSeats finds the closest value total size of list of groups to the available seats.
func closestToAvailableSeats(groups []packing.Group, availableSeats []*layout.Seat) int {

	totalSize := 0
	for _, group := range groups {
		totalSize += group.Size()
	}

	closest := len(availableSeats) - totalSize

	if closest < 0 {
		// max value
		return len(availableSeats)
	}

	return closest
}

// checkEnoughSpace checks if there are enough seats to book the given groups.
func checkEnoughSpace(seats [][]*layout.Seat, groups []packing.Group) error {

	totalSize := 0
	for _, group := range groups {
		totalSize += group.Size()
	}

	// flatten seats
	flatSeats := make([]*layout.Seat, 0)
	for _, row := range seats {
		flatSeats = append(flatSeats, row...)
	}

	if totalSize > len(flatSeats) {
		return ErrNotEnoughSpace
	}

	return nil
}

/*
// printGroups .
func printGroups(groups []packing.Group) {
	for _, group := range groups {
		fmt.Printf("[%d]", group.Size())
	}
}

func printListOfGroupsFor(listOfGroups [][]packing.Group) {
	for _, groups := range listOfGroups {
		fmt.Print("[")
		printGroups(groups)
		fmt.Print("]")
	}
}

// printInfo .
func printInfo(groups []packing.Group, listOfGroups [][]packing.Group, packedGroups []packing.Group, availableSeats []*layout.Seat, l *layout.Layout) {
	return
	fmt.Println("========")
	if groups != nil {
		fmt.Print("Pool: ")
		printGroups(groups)
		fmt.Println()
	}
	if listOfGroups != nil {
		fmt.Print("List of packed groups: ")
		printListOfGroupsFor(listOfGroups)
		fmt.Println()
	}
	if packedGroups != nil {
		fmt.Print("Packed Groups: ")
		printGroups(packedGroups)
		fmt.Println()
	}
	if availableSeats != nil {
		seatLine := ""
		for _, seat := range availableSeats {
			seatLine += seat.String()
		}
		row, col := availableSeats[0].Position()

		fmt.Printf("Available Seats: %s - StartFrom: %d-%d \n", seatLine, row, col)
	}

	fmt.Print(l.String())
}
*/
