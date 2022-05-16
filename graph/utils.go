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

package graph

// contains check if slice contains element
// or not
func contains(slice []interface{}, element interface{}) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// EdgePriorityQueue for Prim's algorithm
// Check Priority queue in queue package
type EdgePriorityQueue struct {
	queue []Edge
}

func makeEdgePriorityQueue() EdgePriorityQueue {
	return EdgePriorityQueue{}
}

func (pq *EdgePriorityQueue) enqueue(edge Edge) {
	if len(pq.queue) == 0 {
		pq.queue = append(pq.queue, edge)
	} else {
		pq.queue = append(pq.queue, edge)
		for root := len(pq.queue)/2 - 1; root >= 0; root-- {
			pq.heapify(root, len(pq.queue))
		}
	}
}

func (pq *EdgePriorityQueue) dequeue() Edge {
	minPriorityItem := pq.queue[0]
	pq.queue = append(pq.queue[:0], pq.queue[1:]...)
	for root := len(pq.queue)/2 - 1; root >= 0; root-- {
		pq.heapify(root, len(pq.queue))
	}
	return minPriorityItem
}

func (pq *EdgePriorityQueue) heapify(root, end int) {
	left := 2*root + 1
	right := 2*root + 2
	min := root

	if left < end && pq.queue[left].weight < pq.queue[min].weight {
		min = left
	}
	if right < end && pq.queue[right].weight < pq.queue[min].weight {
		min = right
	}
	if min != root {
		pq.queue[root], pq.queue[min] = pq.queue[min], pq.queue[root]
		pq.heapify(min, end)
	}
}

func (pq *EdgePriorityQueue) isEmpty() bool {
	if len(pq.queue) <= 0 {
		return true
	}
	return false
}
