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
