package searching

import "testing"

// Binary search only work with sorted arrays
func TestBinarySearchRecursive(t *testing.T) {
	if val := BinarySearchRecursive([]int{1, 2, 45, 99, 450}, 0, 6, 2); val != 1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchRecursive([]int{-45, -1, 0, 2, 45, 99, 450}, 0, 6, 2); val != 3 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchRecursive([]int{-45, -1, 0, 2, 45, 99, 450}, 0, 6, 3); val != -1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
}

func TestBinarySearchIterative(t *testing.T) {
	if val := BinarySearchIterative([]int{1, 2, 45, 99, 450}, 2); val != 1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchIterative([]int{-45, -1, 0, 2, 45, 99, 450}, 2); val != 3 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchIterative([]int{-45, -1, 0, 2, 45, 99, 450}, 3); val != -1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
}
