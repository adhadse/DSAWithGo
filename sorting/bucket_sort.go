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

// BucketSort returns a sorted slice using Bucket Sort algorithm
// In this, buckets are created to put elements into and then applying
// some sorting algorithm (insertion sort) to sort elements in each bucket.
// At last, we join them to set sorted array.
// Not designed for negative numbers.
func BucketSort(list []int) []int {
	// find max element and use length of list to determine
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
