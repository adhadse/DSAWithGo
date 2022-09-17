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

// SelectionSort returns a sorted slice by selecting minimum
// element from unsorted portion of slice and repeating the process
func SelectionSort(list []int) []int {
	for i := range list {
		minidx := i
		for j := i + 1; j < len(list); j++ {
			if list[minidx] > list[j] {
				minidx = j
			}
		}
		list[i], list[minidx] = list[minidx], list[i]
	}
	return list
}
