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

// CountSort returns a sorted slice using counting sorting algorithm
//
// The algorithm assume that the range of possible values in list is limited
// (includes negative integers) create a count array for the occurrences of
// elements.
//
// Then adding each number in count array to right of it accumulatively.
// and shifting it to right, inserting 0 at left.
// This result in count storing initial position of elements which
// get used to create sorted slice.
func CountSort(list []int) []int {
	max := list[0]
	min := list[0]
	out := make([]int, len(list))

	// find minimum and maximum
	for _, val := range list {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	// Create a count array to store count of individual
	// elements since we know max min of the range of numbers
	// Adjusted via -min to store count of negative values also
	count := make([]int, max-min+1)
	for i := range list {
		count[list[i]-min]++
	}

	// Adding each number to the right of accumulative-ly
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	// Build sorted slice
	for i := 0; i < len(list); i++ {
		out[count[list[i]-min]-1] = list[i]
		count[list[i]-min]++ // increase index for next occurrence of this element
	}
	return out
}
