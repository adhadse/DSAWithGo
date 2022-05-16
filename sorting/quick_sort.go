// Copyright 2022 The DSAWithGo Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sorting

// QuickSort returns a sorted slice based on quick
// sort algorithm (Lomuto's partition scheme)
// which is a divide and conquer technique
//
// It finds elements called Pivot which divides slice into two
// parts such that elements in left are smaller than pivot
// and elements in right are greater than pivot
//
// It then adjust pivot, quick sort left part, quick sort right part
//
// Parameters:
//  - list: slice to sort
//  - low: starting index
//  - high: ending index
func QuickSort(list []int, low, high int) []int {
	if low < high {
		partitionIdx := partition(list, low, high)
		QuickSort(list, low, partitionIdx-1)  // Quicksort Left
		QuickSort(list, partitionIdx+1, high) // Quicksort Right
	}
	return list
}

// partition takes last element (index using `high`) as pivot,
// places pivot element at correct position, and places all
// smaller (smaller than pivot) to left (<-) of pivot and all
// greater elements to right (->) and returns Partition Index.
// Parameters:
//  - list: slice to sort
//  - low: starting index
//  - high: ending index
func partition(list []int, low, high int) int {
	pivot := list[high]
	minIdx := low - 1 // Index of smaller element
	for i := low; i < high; i++ {
		if list[i] <= pivot {
			// If current element is smaller than or equal to pivot
			// increment index of smaller element
			// and swap it with item at minIdx
			// i.e., items behind minIdx are smaller than pivot
			minIdx++
			list[minIdx], list[i] = list[i], list[minIdx]
		}
	}
	// Swap list[minIdx+1] and pivot, bringing pivot at appropriate position
	list[minIdx+1], list[high] = list[high], list[minIdx+1]
	return minIdx + 1
}
