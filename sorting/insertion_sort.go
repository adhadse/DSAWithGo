// Copyright (C) 2022   Anurag Dhadse. All Rights Reserved.
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

// InsertionSort returns sorted slice using Insertion sort algorithm
//
// An outer loop iterate from i:1->end of array index and
// treating element at i as "key" compares with all previous elements in inner loop
// during this comparison, it "pushes" bigger elements in place of "key"
// until the "key" fails to be any smaller; gets inserted in the new created position.
func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && key < list[j] {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}
