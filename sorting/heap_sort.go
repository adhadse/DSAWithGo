package sorting

// HeapSort returns sorted slice in ascending order based on
// binary Heap data structure.
//
// first, it repetitively calls heapify() to create max-heap of unsorted slice
//
// then, swapping first and last element of unsorted slice,i.e.,
// root node at index 0 is the max of heap moved to last of unsorted slice
// creating sorted slice at the ending.
//
// then continuing to repeat above two steps.
//
// The heap is interpreted from array [0, 1, 2, 3, 4] in this way:
// as 0 being root 1 & 2 are its child; 3 & 4 are child of 1.
func HeapSort(list []int) []int {
	end := len(list)
	start := end/2 - 1

	// build a max-heap
	// root contains index of root nodes of subtrees, 0 being the root of heap
	for root := start; root > -1; root-- {
		heapify(list, root, end)
	}

	// extract elements one by one
	for i := end - 1; i > 0; i-- {
		list[i], list[0] = list[0], list[i]
		heapify(list, 0, i)
	}
	return list
}

// heapify creates a max-heap given a root node index for a subtree
// (1 root node and two child nodes)
// by comparing root node element with left and right child nodes (if exists)
// and exchanging root node with the max of left or right child and
// then recursively heapify the affected subtree.
//
// Parameters:
//	- list: complete slice storing heap
//  - root: root index to heapify
//  - end:  size of heap to be heapified out of `list`
func heapify(list []int, root, end int) {
	left := 2*root + 1  // left child index of root node
	right := 2*root + 2 // right child index of root node
	max := root         // initialize largest as root

	// If left child of root exists and is greater than root/max node
	if left < end && list[left] > list[max] {
		max = left
	}
	// If right child of root exists and is greater than root/max node
	if right < end && list[right] > list[max] {
		max = right
	}
	// if max is left or right child, exchange root with max
	if max != root {
		list[root], list[max] = list[max], list[root]
		heapify(list, max, end)
	}
}
