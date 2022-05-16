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

// RadixSort returns a sorted slice by performing
// digit by digit soring starting from least significant
// digit to most significant digit.
// It uses counting sort as a subroutine to sort.
// Not designed for slices involving negative numbers
func RadixSort(list []int) []int {
	max := list[0]

	// find max of all elements
	for i := range list {
		if list[i] > max {
			max = list[i]
		}
	}
	// apply CountSort to nth digits of elements
	exp := 1
	for max/exp > 1 {
		list = CountSortWithExpression(list, exp)
		exp *= 10
	}
	return list
}

// CountSortWithExpression returns sorted slice but applies sort
// according to digits represented by exp 1 for first digits, 10 for second digit and so on:
// For ex when exp=1: 328/1%10 = 8; when exp=10: 328/10%10 = 2
func CountSortWithExpression(list []int, exp int) []int {
	out := make([]int, len(list))
	count := make([]int, 10)

	// Create a count array to store count of individual elements
	for i := range list {
		count[((list[i])/exp)%10]++
	}

	// Adding each number to the right of accumulative-ly
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	// Build sorted slice
	for i := 0; i < len(list); i++ {
		index := list[i] / exp
		outIdx := count[index%10]
		out[outIdx-1] = list[i]
		count[(list[i]/exp)%10]++ // increase index for next occurrence of this element
	}
	return out
}
