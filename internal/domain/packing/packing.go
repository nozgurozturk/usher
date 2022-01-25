package packing

type Group interface {
	Size() int
}

type PackingFunc func([]Group, int) [][]Group

// Packing is the packing of the groups with given groups and function
func PackGroups(groups []Group, capacity int, fn PackingFunc) [][]Group {
	return fn(groups, capacity)
}

// NextFit is the next fit packing algorithm.
// Check the current group fits in the current bin.
// If so, then insert group into bin otherwise create new bin and insert group into bin.
func NextFit(groups []Group, capacity int) [][]Group {

	// Packed groups
	bins := make([][]Group, 0, len(groups))

	// Remaining size of the current bin
	remaining := capacity

	for _, group := range groups {

		// If the group does not fit in the current, create new bin and insert group into bin.
		// Otherwise insert group into bin.
		if group.Size() > remaining || len(bins) == 0 {
			bins = append(bins, []Group{group})
			remaining = capacity - group.Size()
		} else {
			bins[len(bins)-1] = append(bins[len(bins)-1], group)
			remaining -= group.Size()
		}
	}

	return bins

}

// FirstFit is the first fit packaging algorithm.
// Check the bins in order and insert the group into the first bin that can fit the group.
// If no bin can fit the group, create a new bin and insert the group into the bin.
func FirstFit(groups []Group, capacity int) [][]Group {
	// Packed groups
	bins := make([][]Group, 0, len(groups))

	// Remaining size of the groups
	remaining := make([]int, len(groups))

	// Set remaining size of the groups
	for i := range groups {
		remaining[i] = capacity
	}

	for _, group := range groups {
		i := 0
		for j := range bins {
			// If the group size is fit into the current bin, decrease remaining size of bin.
			if remaining[j] >= group.Size() {
				remaining[j] -= group.Size()
				bins[i] = append(bins[i], group)
				break
			}
			i++
		}

		// If the group size is not fit into any bin, create a new bin and insert the group into the bin.
		if len(bins) == i {
			remaining[i] -= group.Size()
			bins = append(bins, []Group{group})
		}
	}

	return bins
}

// BestFit is the best fit packing algorithm.
// Insert next item into bin with the smallest remaining space.
// If no bin can fit the group, create a new bin and insert the group into the bin.
func BestFit(groups []Group, capacity int) [][]Group {

	bins := make([][]Group, 0, len(groups))

	remaining := make([]int, len(groups))

	for i := range groups {
		remaining[i] = capacity
	}

	for _, group := range groups {
		// minRemaining is the smallest remaining size of the bins.
		minRemaining := capacity + 1
		// index of the bin with the smallest remaining space
		binIndex := 0

		for j := range bins {

			if remaining[j] >= group.Size() && remaining[j] < minRemaining {
				minRemaining = remaining[j]
				binIndex = j
			}
		}

		// If the group size is not fit into any bin, create a new bin and insert the group into the bin.
		if minRemaining == capacity+1 {
			remaining[len(bins)] = capacity - group.Size()
			bins = append(bins, []Group{group})
		} else {
			remaining[binIndex] -= group.Size()
			bins[binIndex] = append(bins[binIndex], group)
		}
	}

	return bins
}
