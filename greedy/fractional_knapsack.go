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

type KnapsackItem struct {
	weight int
	value  int
	cost   int // value / weight
}

func SolveFractionalKnapsack(weight, value []int, knapsackCapacity int) int {
	var knapsackItems []KnapsackItem
	for i := range weight {
		knapsackItems = append(knapsackItems, KnapsackItem{weight: weight[i],
			value: value[i], cost: value[i] / weight[i]})
	}

	knapsackItems = InsertSortCost(knapsackItems)

	totalValue := 0
	for i := range knapsackItems {
		itemWeight := knapsackItems[i].weight
		itemValue := knapsackItems[i].value
		if knapsackCapacity-itemWeight >= 0 {
			knapsackCapacity -= itemWeight
			totalValue += itemValue
		} else {
			fraction := knapsackCapacity / itemWeight
			totalValue += itemValue * fraction
			knapsackCapacity = knapsackCapacity - (itemWeight * fraction)
			break
		}
	}
	return totalValue
}
