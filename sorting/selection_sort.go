package sorting

// SelectionSort returns a sorted slice by selecting minimum
// element from unsorted portion of slice and repeating the process
func SelectionSort(list []int) []int {
	for i := range list {
		minidx := i
		for j := i + 1; j < len(list); j++ {
			if list[minidx] > list[j] {
				minidx = j
			}
		}
		list[i], list[minidx] = list[minidx], list[i]
	}
	return list
}
