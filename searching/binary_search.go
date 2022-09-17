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

package searching

func BinarySearchRecursive(list []int, low, high, key int) int {
	if high >= low {
		mid := low + (high-low)/2
		if list[mid] == key {
			return mid
		} else if list[mid] > key {
			return BinarySearchRecursive(list, low, mid-1, key)
		} else {
			return BinarySearchRecursive(list, mid+1, high, key)
		}
	} else {
		return -1
	}
}

func BinarySearchIterative(list []int, key int) int {
	low := 0
	high := len(list) - 1
	mid := 0
	for high >= low {
		mid = (low + high) / 2
		if list[mid] == key {
			return mid
		} else if list[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
