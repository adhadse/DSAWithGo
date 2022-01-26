package sorting

// BubbleSort returns a sorted slice using bubble sort algorithm
//
// It has an outer loop looping over whole list index i and
// in the inner loop of j from 0->n-i-1 repetitively creates "bubbles"; item at j and j+1
// exchanging position for incorrectly ordered elements in the "bubble"
func BubbleSort(list []int) []int {
	n := len(list)
	for i := range list {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
