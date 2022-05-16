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

// BubbleSort returns a sorted slice using bubble sort algorithm
//
// It has an outer loop looping over whole list index i and
// in the inner loop of j from 0->n-i-1 repetitively creates "bubbles"; item at j and j+1
// exchanging position for incorrectly ordered elements in the "bubble"
func BubbleSort(list []int) []int {
	n := len(list)
	for i := range list {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
