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

package queue

type PriorityQueue struct {
	queue []int
}

func MakePriorityQueue() PriorityQueue {
	return PriorityQueue{}
}

func (pq *PriorityQueue) Enqueue(data int) {
	if len(pq.queue) == 0 {
		pq.queue = append(pq.queue, data)
	} else {
		pq.queue = append(pq.queue, data)
		for root := len(pq.queue)/2 - 1; root >= 0; root-- {
			pq.heapify(root, len(pq.queue))
		}
	}
}

func (pq *PriorityQueue) Dequeue() int {
	maxPriorityItem := pq.queue[0]
	pq.queue = append(pq.queue[:0], pq.queue[1:]...)
	for root := len(pq.queue)/2 - 1; root >= 0; root-- {
		pq.heapify(root, len(pq.queue))
	}
	return maxPriorityItem
}

func (pq *PriorityQueue) heapify(root, end int) {
	left := 2*root + 1  // left child index of root node
	right := 2*root + 2 // right child index of root node
	max := root         // initialize largest as root

	// If left child of root exists and is greater than root/max node
	if left < end && pq.queue[left] > pq.queue[max] {
		max = left
	}
	// If right child of root exists and is greater than root/max node
	if right < end && pq.queue[right] > pq.queue[max] {
		max = right
	}
	// if max is left or right child, exchange root with max
	if max != root {
		pq.queue[root], pq.queue[max] = pq.queue[max], pq.queue[root]
		pq.heapify(max, end)
	}
}
