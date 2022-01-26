package sorting

// InsertionSort returns sorted slice using Insertion sort algorithm
//
// An outer loop iterate from i:1->end of array index and
// treating element at i as "key" compares with all previous elements in inner loop
// during this comparison, it "pushes" bigger elements in place of "key"
// until the "key" fails to be any smaller; gets inserted in the new created position.
func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && key < list[j] {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}
