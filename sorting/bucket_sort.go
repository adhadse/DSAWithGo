package sorting

// BucketSort returns a sorted slice using Bucket Sort algorithm
// In this, buckets are created to put elements into and then applying
// some sorting algorithm (insertion sort) to sort elements in each bucket.
// At last, we join them to set sorted array.
// Not designed for negative numbers.
func BucketSort(list []int) []int {
	// Find max element and use length of list to determine
	// which value in list goes into which bucket
	max := list[0]
	for i := range list {
		if list[i] > max {
			max = list[i]
		}
	}
	size := max / len(list)

	// Create n = len(list) empty buckets
	bucketList := make([][]int, len(list))

	// Put list elements into different buckets based on size
	for i := range list {
		bucketIdx := list[i] / size
		if bucketIdx < len(list) {
			// put in jth bucket
			bucketList[bucketIdx] = append(bucketList[bucketIdx], list[i])
		} else {
			// put in last bucket
			bucketList[len(bucketList)-1] = append(bucketList[len(bucketList)-1], list[i])
		}
	}

	// Sort elements within the buckets using Insertion Sort
	for i := range list {
		bucketList[i] = InsertionSort(bucketList[i])
	}

	// Concatenate buckets into a single list
	var out []int
	for i := range list {
		out = append(out, bucketList[i]...) // unpack slice into individual elements
	}
	return out
}
