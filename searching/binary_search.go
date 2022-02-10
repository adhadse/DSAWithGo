package searching

func BinarySearchRecursive(list []int, low, high, key int) int {
	if high >= low {
		mid := low + (high-low)/2
		if list[mid] == key {
			return mid
		} else if list[mid] > key {
			return BinarySearchRecursive(list, low, mid-1, key)
		} else {
			return BinarySearchRecursive(list, mid+1, high, key)
		}
	} else {
		return -1
	}
}

func BinarySearchIterative(list []int, key int) int {
	low := 0
	high := len(list) - 1
	mid := 0
	for high >= low {
		mid = (low + high) / 2
		if list[mid] == key {
			return mid
		} else if list[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
