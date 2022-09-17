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

package greedy

import "testing"

func TestSolveFractionalKnapsack(t *testing.T) {
	weights := []int{10, 40, 20, 30}
	values := []int{60, 40, 100, 120}
	knapsackCapacity := 50
	if val := SolveFractionalKnapsack(weights, values, knapsackCapacity); val != 240 {
		t.Errorf("Fractional Knapsack returned total value other than 240: %d", val)
	}
}

// Should result in
// f -> 00
// c -> 0100
// d -> 0101
// a -> 01100
// b -> 01101
// e -> 0111
func TestHuffmanCode(t *testing.T) {
	characters := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	frequencies := []int{5, 9, 12, 13, 16, 45}

	HuffmanNodes := HuffmanCode(characters, frequencies)
	PrintHuffmanNodes(HuffmanNodes[0], "")
}
