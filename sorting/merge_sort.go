package sorting

// MergeSort is a Divide and Conquer algorithm which
// divides input slice into two halves, calls itself for the two
// halves, and then merges the two sorted halves returning sorted slice.
func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	middle := len(list) / 2
	left := MergeSort(list[:middle])
	right := MergeSort(list[middle:])
	return merge(left, right)
}

// merge returns a merged slice
func merge(left, right []int) []int {
	i, j := 0, 0
	var result []int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)  // ... unpack a slice and
	result = append(result, right[j:]...) // pass as parameter to append
	return result
}
