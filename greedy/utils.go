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

package greedy

func InsertSortCost(list []KnapsackItem) []KnapsackItem {
	for i := 1; i < len(list); i++ {
		key := list[i].cost
		j := i - 1
		for j >= 0 && key < list[j].cost {
			list[j+1] = list[j]
			j--
		}
		list[j+1].cost = key
	}
	return list
}

func InsertSortFrequency(list []*HuffmanNode) []*HuffmanNode {
	for i := 1; i < len(list); i++ {
		item := list[i]
		key := list[i].frequency
		j := i - 1
		for j >= 0 && key < list[j].frequency {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = item
	}
	return list
}

func remove(slice []*HuffmanNode, s *HuffmanNode) []*HuffmanNode {
	for i := range slice {
		if slice[i] == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
