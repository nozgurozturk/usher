package seating

const (
	ROW_SIZE = 3
	COL_SIZE = 8
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

func AllocateSeats(groupsOfUser []int, isReversed bool) [][]int {

	seats := make([][]int, ROW_SIZE, ROW_SIZE)

	for i := 0; i < ROW_SIZE; i++ {
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

	return seats
}
