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

import "testing"

// Binary search only work with sorted arrays
func TestBinarySearchRecursive(t *testing.T) {
	if val := BinarySearchRecursive([]int{1, 2, 45, 99, 450}, 0, 6, 2); val != 1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchRecursive([]int{-45, -1, 0, 2, 45, 99, 450}, 0, 6, 2); val != 3 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchRecursive([]int{-45, -1, 0, 2, 45, 99, 450}, 0, 6, 3); val != -1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
}

func TestBinarySearchIterative(t *testing.T) {
	if val := BinarySearchIterative([]int{1, 2, 45, 99, 450}, 2); val != 1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchIterative([]int{-45, -1, 0, 2, 45, 99, 450}, 2); val != 3 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
	if val := BinarySearchIterative([]int{-45, -1, 0, 2, 45, 99, 450}, 3); val != -1 {
		t.Errorf("Recursive binary search return index other than 1 for key 12: %d", val)
	}
}
