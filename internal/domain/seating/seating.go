package seating

import (
	"fmt"
	"sort"

	"github.com/nozgurozturk/usher/internal/domain/group"
	"github.com/nozgurozturk/usher/internal/domain/layout"
	"github.com/nozgurozturk/usher/internal/domain/packing"
)

var (
	ErrNotEnoughSeats = fmt.Errorf("not enough seats")
	ErrNotEnoughSpace = fmt.Errorf("not enough space")
)

func ReserveSeatsForGroups(g []group.Group, h layout.Hall) (layout.Hall, error) {

	// Copy Layout
	hall := h.Copy()

	availableSeats := layout.ConsecutiveFilteredSeatsInHall(hall, layout.NewFilter().WithAvailable(true))
	sort.SliceStable(availableSeats, func(i, j int) bool {
		return len(availableSeats[i]) < len(availableSeats[j])
	})

	// check Total Size of groups is less than the available seats.
	if err := checkEnoughSpace(availableSeats, g); err != nil {
		return h, err
	}

	// copy groups
	groups := make([]group.Group, len(g))
	copy(groups, g)

	// sort groups by size in descending order
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].Size() < groups[j].Size()
	})

	// Pack groups for the available seats.
	packedListOfGroups := packing.PackGroupsOfUser(groups, len(availableSeats), packing.NextFit)

	for _, packedGroups := range packedListOfGroups {

		for _, group := range packedGroups {

			// Find Closest Seat block for the group.
			closestSeatBlock, _ := findClosestSeatBlock(group, availableSeats)

			// If didn't find any seat block then skip this group.
			if closestSeatBlock == nil {
				continue
			}

			// printInfo(groups, packedListOfGroups, group, closestSeatBlock, hall)
			// Book the seats.
			for i := 0; i < group.Size(); i++ {
				closestSeatBlock[i].Book()
			}

			// Remove group from groups.
			groups = removeGroup(groups, group)

			if len(groups) == 0 {
				break
			}
		}
	}

	if len(groups) > 0 {
		return h, ErrNotEnoughSeats
	}

	return hall, nil
}

// removeGroup removes the given group from the list of groups.
func removeGroup(groups []group.Group, group group.Group) []group.Group {
	for i, g := range groups {
		if g.ID() == group.ID() {
			return append(groups[:i], groups[i+1:]...)
		}
	}
	return groups
}

// checkEnoughSpace checks if the total size of groups is less than the available seats.
func checkEnoughSpace(seats [][]layout.Seat, groups []group.Group) error {

	totalSize := 0
	for _, group := range groups {
		totalSize += group.Size()
	}

	// flatten seats
	flatSeats := make([]layout.Seat, 0)
	for _, row := range seats {
		flatSeats = append(flatSeats, row...)
	}

	if totalSize > len(flatSeats) {
		return ErrNotEnoughSpace
	}

	return nil
}

func findClosestSeatBlock(group group.Group, availableSeats [][]layout.Seat) ([]layout.Seat, int) {

	filteredAvailableSeats := make([][]layout.Seat, 0, len(availableSeats))
	filter := layout.NewFilter().
		WithAvailable(true).
		WithRank(group.RankPreference()).
		WithFeature(layout.SeatFeature(group.SeatPreferences()))

	for _, seatBlock := range availableSeats {
		filteredSeatBlock := layout.FilteredSeatBlock(seatBlock, filter)
		if len(filteredSeatBlock) > 0 {
			filteredAvailableSeats = append(filteredAvailableSeats, filteredSeatBlock)
		}
	}

	maxSize := 0
	for _, seatBlock := range filteredAvailableSeats {
		if len(seatBlock) > maxSize {
			maxSize = len(seatBlock)
		}
	}

	// max value
	minAvailableSeatsInBlock := maxSize
	var closestSeatBlock []layout.Seat

	for _, seatBlock := range filteredAvailableSeats {

		if len(seatBlock)-group.Size() < minAvailableSeatsInBlock && len(seatBlock) >= group.Size() {
			minAvailableSeatsInBlock = len(seatBlock) - group.Size()
			closestSeatBlock = seatBlock
		}

		if len(seatBlock) == group.Size() {
			break
		}
	}

	return closestSeatBlock, minAvailableSeatsInBlock
}

/*
func printInfo(groups []group.Group, packedGroup [][]group.Group, current group.Group, availableSeats []layout.Seat, l layout.Hall) {

	fmt.Println("========")
	if groups != nil {
		fmt.Print("Pool: ")
		for _, group := range groups {
			fmt.Printf("[%d]", group.Size())
		}
		fmt.Println()
	}

	if packedGroup != nil {
		fmt.Print("Packed Groups: ")
		for _, groups := range packedGroup {
			fmt.Print("[")
			for _, group := range groups {
				fmt.Printf("[%d]", group.Size())
			}
			fmt.Print("]")
		}
		fmt.Println()
	}

	if current != nil {
		fmt.Print("Current Group: ")
		fmt.Println(current.String())
	}

	if availableSeats != nil {
		seatLine := ""
		for _, seat := range availableSeats {
			seatLine += seat.String()

		}

		row, col := availableSeats[0].Position()

		fmt.Printf("Available Seats: %s - R:%d C:%d \n", seatLine, row, col)
	}

	fmt.Print(l.String())
}
*/
