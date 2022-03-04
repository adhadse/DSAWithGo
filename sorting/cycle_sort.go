package sorting

// CycleSort returns a sorted slice and number of writes/swap performed
func CycleSort(list []int) ([]int, int) {
	writes := 0

	// Loop through the array to find cycles to rotate
	for cycleStart := 0; cycleStart < len(list)-1; cycleStart++ {
		item := list[cycleStart]
		position := cycleStart

		// find position where to put the item
		for i := cycleStart + 1; i < len(list); i++ {
			if list[i] < item {
				position++
			}
		}

		// If item is already there, this is not a cycle
		if position == cycleStart {
			continue
		}

		// Otherwise, put the item there or right after any duplicates
		for item == list[position] {
			position++
			list[position], item = item, list[position]
		}

		// Rotate the rest of the cycle
		for position != cycleStart {
			position = cycleStart
			// find position where to put the element
			for i := cycleStart + 1; i < len(list); i++ {
				if list[i] < item {
					position++
				}
			}
			// ignore all duplicate elements
			for item == list[position] {
				position++
			}
			// put the item to it's right position
			if item != list[position] {
				list[position], item = item, list[position]
				writes++
			}
		}
	}
	return list, writes
}
